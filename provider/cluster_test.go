package provider

import (
	"testing"
)

func TestMergeConfigs(t *testing.T) {
	name := "my-cluster"
	role := "controller+worker"
	address := "1.2.3.4"
	user := "foo"
	falseValue := false

	args := &ClusterArgs{
		Metadata: &ClusterMetadata{Name: &name},
		Spec: &ClusterSpec{
			Hosts: []ClusterHost{
				{Role: &role, SSH: &ClusterSSH{Address: &address, User: &user}},
			},
			K0s: &ClusterK0s{
				Config: &ClusterK0sConfig{
					Spec: &ClusterK0sSpec{
						Telemetry: &ClusterTelemetry{Enabled: &falseValue},
					},
				},
			},
		},
	}

	_, cluster, err := clusterArgsToK0sCtlCluster(name, args)
	if err != nil {
		t.Errorf("clusterArgsToK0sCtlCluster error: %s", err.Error())

		return
	}

	if cluster.Metadata.Name != name {
		t.Errorf("cluster.Metadata.Name: want: %s; got: %s", name, *args.Metadata.Name)
	}

	if len(cluster.Spec.Hosts) != 1 {
		t.Errorf("len(cluster.Spec.Hosts): want: %d; got: %d", len(args.Spec.Hosts), len(cluster.Spec.Hosts))

		return
	}

	if cluster.Spec.Hosts[0].Role != role {
		t.Errorf("cluster.Spec.Hosts[0].Role: want: %s; got: %s", *args.Spec.Hosts[0].Role, cluster.Spec.Hosts[0].Role)
	}

	if cluster.Spec.Hosts[0].SSH == nil {
		t.Errorf("cluster.Spec.Hosts[0].SSH is nil")
	}

	if cluster.Spec.Hosts[0].SSH.Address != address {
		t.Errorf("cluster.Spec.Hosts[0].SSH.Address: want: %s; got: %s", *args.Spec.Hosts[0].SSH.Address, cluster.Spec.Hosts[0].SSH.Address)
	}

	if cluster.Spec.Hosts[0].SSH.User != user {
		t.Errorf("cluster.Spec.Hosts[0].SSH.Address: want: %s; got: %s", *args.Spec.Hosts[0].SSH.User, cluster.Spec.Hosts[0].SSH.User)
	}

	if cluster.Spec.K0s.Config == nil {
		t.Errorf("cluster.Spec.K0s.Config is nil")
	}

	telemetryEnabled, ok := cluster.Spec.K0s.Config.Dig("spec", "telemetry", "enabled").(bool)
	if !ok || telemetryEnabled != falseValue {
		t.Errorf("cluster.Spec.K0s.Config.Spec.Telemetry.Enabled: want: %t; got: %t", *args.Spec.K0s.Config.Spec.Telemetry.Enabled, telemetryEnabled)
	}
}
