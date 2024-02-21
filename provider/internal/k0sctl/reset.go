package k0sctl

import (
	"github.com/k0sproject/k0sctl/phase"
	"github.com/k0sproject/k0sctl/pkg/apis/k0sctl.k0sproject.io/v1beta1"
)

func Reset(cluster *Cluster) (*Cluster, error) {
	clt := v1beta1.Cluster(*cluster)

	manager := phase.Manager{Config: &clt}

	lockPhase := &phase.Lock{}

	manager.AddPhase(
		&phase.Connect{},
		&phase.DetectOS{},
		lockPhase,
		&phase.PrepareHosts{},
		&phase.GatherK0sFacts{},
		&phase.RunHooks{Stage: "before", Action: "reset"},
		&phase.ResetWorkers{NoDrain: true, NoDelete: true},
		&phase.ResetControllers{NoDrain: true, NoDelete: true, NoLeave: true},
		&phase.ResetLeader{},
		&phase.RunHooks{Stage: "after", Action: "reset"},
		&phase.Unlock{Cancel: lockPhase.Cancel},
		&phase.Disconnect{},
	)

	if err := manager.Run(); err != nil {
		return nil, err
	}

	newCluster := Cluster(*manager.Config)

	return &newCluster, nil
}
