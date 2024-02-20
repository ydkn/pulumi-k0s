package k0sctl

import (
	"fmt"

	"github.com/k0sproject/k0sctl/pkg/apis/k0sctl.k0sproject.io/v1beta1"
)

type Cluster v1beta1.Cluster

func (c *Cluster) APIAddress() string {
	address := "localhost"
	port := 6443

	clt := v1beta1.Cluster(*c)

	leader := clt.Spec.K0sLeader()
	if leader != nil {
		if leader.SSH != nil {
			address = leader.SSH.Address
		}

		if leader.WinRM != nil {
			address = leader.WinRM.Address
		}
	}

	if clt.Spec != nil && clt.Spec.K0s != nil && clt.Spec.K0s.Config != nil {
		config := clt.Spec.K0s.Config

		externalAddress := config.Dig("spec", "api", "externalAddress")
		if externalAddress != nil {
			externalAddressString, ok := externalAddress.(string)
			if ok {
				address = externalAddressString
			}
		}

		apiPort := config.Dig("spec", "api", "port")
		if apiPort != nil {
			apiPortInt, ok := apiPort.(int)
			if ok {
				port = apiPortInt
			}
		}
	}

	return fmt.Sprintf("https://%s:%d", address, port)
}
