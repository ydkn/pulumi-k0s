package k0sctl

import (
	"github.com/k0sproject/k0sctl/phase"
	k0sctlCluster "github.com/k0sproject/k0sctl/pkg/apis/k0sctl.k0sproject.io/v1beta1/cluster"
)

func KubeConfig(cluster *Cluster, apiAddress string) (*Cluster, error) {
	c := cluster.K0sCtlObject()

	// Change so that the internal config has only single controller host as we do not need to connect to all nodes
	c.Spec.Hosts = k0sctlCluster.Hosts{c.Spec.K0sLeader()}
	manager := phase.Manager{Config: c}

	kubeconfig := &getKubeconfigPhase{APIAddress: apiAddress}

	manager.AddPhase(
		&phase.Connect{},
		&phase.DetectOS{},
		kubeconfig,
		&phase.Disconnect{},
	)

	if err := manager.Run(); err != nil {
		return nil, err
	}

	cluster.KubeConfig = kubeconfig.Kubeconfig()

	return cluster, nil
}
