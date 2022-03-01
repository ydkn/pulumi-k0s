package k0sctl

import (
	"encoding/json"

	"github.com/creasty/defaults"
	"github.com/k0sproject/dig"
	"github.com/k0sproject/k0sctl/cmd"
	"github.com/k0sproject/k0sctl/integration/github"
	"github.com/k0sproject/k0sctl/pkg/apis/k0sctl.k0sproject.io/v1beta1"
	"github.com/k0sproject/k0sctl/pkg/apis/k0sctl.k0sproject.io/v1beta1/cluster"
	"gopkg.in/yaml.v2"
)

// DefaultCluster returns a default cluster configuration.
// see https://github.com/k0sproject/k0sctl/blob/main/cmd/init.go
func DefaultClusterConfig() (*Cluster, error) {
	k0sVersion, err := latestK0sVersion()
	if err != nil {
		return nil, err
	}

	cfg := v1beta1.Cluster{
		APIVersion: v1beta1.APIVersion,
		Kind:       "Cluster",
		Metadata:   &v1beta1.ClusterMetadata{Name: "k0s-cluster"},
		Spec: &cluster.Spec{
			Hosts: make(cluster.Hosts, 0),
			K0s:   &cluster.K0s{Version: k0sVersion},
		},
	}

	if err := defaults.Set(&cfg); err != nil {
		return nil, err
	}

	cfg.Spec.K0s.Config = dig.Mapping{}
	if err := yaml.Unmarshal(cmd.DefaultK0sYaml, &cfg.Spec.K0s.Config); err != nil {
		return nil, err
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return nil, err
	}

	var cluster Cluster

	if err := json.Unmarshal(data, &cluster); err != nil {
		return nil, err
	}

	return &cluster, nil
}

func latestK0sVersion() (string, error) {
	return github.LatestK0sVersion(false)
}
