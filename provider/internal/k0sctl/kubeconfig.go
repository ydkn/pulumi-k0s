package k0sctl

import (
	"github.com/k0sproject/k0sctl/phase"
	"github.com/k0sproject/k0sctl/pkg/apis/k0sctl.k0sproject.io/v1beta1"
	k0sctlCluster "github.com/k0sproject/k0sctl/pkg/apis/k0sctl.k0sproject.io/v1beta1/cluster"
)

func Kubeconfig(cluster *Cluster) (*Cluster, error) {
	clt := v1beta1.Cluster(*cluster)

	// Change so that the internal config has only single controller host as we do not need to connect to all nodes
	clt.Spec.Hosts = k0sctlCluster.Hosts{clt.Spec.K0sLeader()}

	manager := phase.Manager{Config: &clt}

	manager.AddPhase(
		&phase.Connect{},
		&phase.DetectOS{},
		&phase.GetKubeconfig{APIAddress: cluster.APIAddress()},
		&phase.Disconnect{},
	)

	if err := manager.Run(); err != nil {
		return nil, err
	}

	newCluster := Cluster(*manager.Config)

	return &newCluster, nil
}
