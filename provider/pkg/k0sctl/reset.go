package k0sctl

import "github.com/k0sproject/k0sctl/phase"

func Reset(cluster *Cluster) error {
	manager := phase.Manager{Config: cluster.K0sCtlObject()}

	manager.AddPhase(
		&phase.Connect{},
		&phase.DetectOS{},
		&phase.PrepareHosts{},
		&phase.GatherK0sFacts{},
		&phase.RunHooks{Stage: "before", Action: "reset"},
		&phase.Reset{},
		&phase.RunHooks{Stage: "after", Action: "reset"},
		&phase.Disconnect{},
	)

	return manager.Run()
}
