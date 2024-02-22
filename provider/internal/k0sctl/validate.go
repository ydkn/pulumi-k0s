package k0sctl

import (
	"github.com/k0sproject/k0sctl/pkg/apis/k0sctl.k0sproject.io/v1beta1"
)

func Validate(cluster *Cluster) error {
	clt := v1beta1.Cluster(*cluster)

	return clt.Validate()
}
