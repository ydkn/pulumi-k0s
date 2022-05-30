package k0sctl

import (
	"github.com/k0sproject/k0sctl/phase"
	"gopkg.in/yaml.v2"
)

type ApplyConfig struct {
	SkipDowngradeCheck bool
	NoDrain            bool
}

func Apply(cluster *Cluster, config ApplyConfig) (*Cluster, error) {
	manager := phase.Manager{Config: cluster.K0sCtlObject()}

	lockPhase := &phase.Lock{}

	manager.AddPhase(
		&phase.Connect{},
		&phase.DetectOS{},
		lockPhase,
		&phase.PrepareHosts{},
		&phase.GatherFacts{},
		&phase.DownloadBinaries{},
		&phase.UploadFiles{},
		&phase.ValidateHosts{},
		&phase.GatherK0sFacts{},
		&phase.ValidateFacts{SkipDowngradeCheck: config.SkipDowngradeCheck},
		&phase.UploadBinaries{},
		&phase.DownloadK0s{},
		&phase.RunHooks{Stage: "before", Action: "apply"},
		&phase.PrepareArm{},
		&phase.ConfigureK0s{},
		&phase.Restore{RestoreFrom: ""},
		&phase.InitializeK0s{},
		&phase.InstallControllers{},
		&phase.InstallWorkers{},
		&phase.UpgradeControllers{},
		&phase.UpgradeWorkers{NoDrain: config.NoDrain},
		&phase.RunHooks{Stage: "after", Action: "apply"},
		&phase.Unlock{Cancel: lockPhase.Cancel},
		&phase.Disconnect{},
	)

	if err := manager.Run(); err != nil {
		return nil, err
	}

	data, err := yaml.Marshal(manager.Config)
	if err != nil {
		return nil, err
	}

	var output Cluster

	if err := yaml.Unmarshal(data, &output); err != nil {
		return nil, err
	}

	return &output, nil
}
