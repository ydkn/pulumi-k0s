package k0sctl

import (
	"github.com/k0sproject/k0sctl/phase"
	"github.com/k0sproject/k0sctl/pkg/apis/k0sctl.k0sproject.io/v1beta1"
)

type ApplyConfig struct {
	SkipDowngradeCheck bool
	NoDrain            bool
	RestoreFrom        string
}

func Apply(cluster *Cluster, config ApplyConfig) (*Cluster, error) {
	clt := v1beta1.Cluster(*cluster)

	manager := phase.Manager{Config: &clt}

	lockPhase := &phase.Lock{}

	manager.AddPhase(
		&phase.DefaultK0sVersion{},
		&phase.Connect{},
		&phase.DetectOS{},
		lockPhase,
		&phase.PrepareHosts{},
		&phase.GatherFacts{},
		&phase.ValidateHosts{},
		&phase.GatherK0sFacts{},
		&phase.ValidateFacts{SkipDowngradeCheck: config.SkipDowngradeCheck},
		&phase.DownloadBinaries{},
		&phase.UploadK0s{},
		&phase.DownloadK0s{},
		&phase.UploadFiles{},
		&phase.InstallBinaries{},
		&phase.PrepareArm{},
		&phase.ConfigureK0s{},
		&phase.Restore{RestoreFrom: config.RestoreFrom},
		&phase.RunHooks{Stage: "before", Action: "apply"},
		&phase.InitializeK0s{},
		&phase.InstallControllers{},
		&phase.InstallWorkers{},
		&phase.UpgradeControllers{},
		&phase.UpgradeWorkers{NoDrain: config.NoDrain},
		&phase.ResetWorkers{NoDrain: config.NoDrain},
		&phase.ResetControllers{NoDrain: config.NoDrain},
		&phase.RunHooks{Stage: "after", Action: "apply"},
		&phase.GetKubeconfig{APIAddress: cluster.APIAddress()},
		&phase.Unlock{Cancel: lockPhase.Cancel},
		&phase.Disconnect{},
	)

	if err := manager.Run(); err != nil {
		return nil, err
	}

	newCluster := Cluster(*manager.Config)

	return &newCluster, nil
}
