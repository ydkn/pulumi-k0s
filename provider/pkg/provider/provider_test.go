package provider

import (
	"testing"

	"github.com/ydkn/pulumi-k0s/provider/pkg/k0sctl"
)

func TestMergeConfigs(t *testing.T) {
	config0, _ := k0sctl.DefaultClusterConfig()

	config1 := &k0sctl.Cluster{
		Metadata: &k0sctl.Metadata{Name: "my-cluster"},
		Spec: &k0sctl.Spec{
			Hosts: k0sctl.Hosts{
				{Role: "controller", SSH: &k0sctl.SSH{Address: "0.0.0.0", User: "foo"}},
			},
			K0s: &k0sctl.K0s{
				Version: "2",
				Config: map[string]interface{}{
					"spec": map[string]interface{}{
						"api": map[string]interface{}{"k0sApiPort": 123},
					},
				},
			},
		},
	}

	config, err := mergeConfigs(config0, config1)
	if err != nil {
		t.Errorf("mergeConfigs(-1) error: %s", err.Error())
	}

	if config.Metadata.Name != "my-cluster" {
		t.Errorf("mergeConfigs(-1) error: metadata.Name; want: my-cluster; got: %s", config.Metadata.Name)
	}

	if len(config.Spec.Hosts) != 1 {
		t.Errorf("mergeConfigs(-1) error: len(spec.Hosts); want: 1; got: %d", len(config.Spec.Hosts))
	}

	if config.Spec.K0s.Version != "2" {
		t.Errorf("mergeConfigs(-1) error: spec.K0s.Version; want: 2; got: %s", config.Spec.K0s.Version)
	}
}
