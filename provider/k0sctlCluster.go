package provider

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"dario.cat/mergo"
	"github.com/k0sproject/dig"
	k0sapi "github.com/k0sproject/k0s/pkg/apis/k0s/v1beta1"
	"github.com/k0sproject/k0sctl/cmd"
	"github.com/k0sproject/k0sctl/pkg/apis/k0sctl.k0sproject.io/v1beta1"
	"github.com/k0sproject/k0sctl/pkg/apis/k0sctl.k0sproject.io/v1beta1/cluster"
	"github.com/k0sproject/rig"
	"github.com/k0sproject/version"
	"gopkg.in/yaml.v2"
	yamlJSON "sigs.k8s.io/yaml"
)

type k0sctlCluster ClusterInputs

const (
	apiVersion = v1beta1.APIVersion
	kind       = "Cluster"
)

func PrepareCluster(name string, news *ClusterInputs) error {
	defaultCluster, err := clusterDefaults(name)
	if err != nil {
		return err
	}

	newsCluster := k0sctlCluster(*news)
	newsJSONData, err := json.Marshal(newsCluster)
	if err != nil {
		return err
	}

	newsClusterCopy := k0sctlCluster{}
	if err := json.Unmarshal(newsJSONData, &newsClusterCopy); err != nil {
		return err
	}

	if err := mergo.Merge(&newsCluster, defaultCluster); err != nil {
		return err
	}

	if news.APIVersion == nil || *news.APIVersion == "" {
		version := apiVersion
		news.APIVersion = &version
	}

	if news.Kind == nil || *news.Kind == "" {
		k := kind
		news.Kind = &k
	}

	if news.Metadata == nil {
		news.Metadata = &ClusterMetadata{}
	}

	if news.Metadata.Name == nil {
		news.Metadata.Name = &name
	}

	if news.Spec == nil {
		news.Spec = &ClusterSpec{}
	}

	if news.Spec.K0s == nil {
		news.Spec.K0s = &ClusterK0s{}
	}

	if news.Spec.K0s.Config == nil {
		news.Spec.K0s.Config = &K0s{}
	}

	if news.Spec.K0s.DynamicConfig == nil {
		falseValue := false
		news.Spec.K0s.DynamicConfig = &falseValue
	}

	if news.Spec.K0s.VersionChannel == nil {
		versionChannel := "stable"
		news.Spec.K0s.VersionChannel = &versionChannel
	}

	if news.Spec.K0s.Config.Metadata == nil {
		news.Spec.K0s.Config.Metadata = &K0sMetadata{}
	}

	if news.Spec.K0s.Config.Metadata.Name == nil {
		news.Spec.K0s.Config.Metadata.Name = &name
	}

	if news.Spec.K0s.Config.Spec == nil {
		news.Spec.K0s.Config.Spec = &K0sSpec{}
	}

	if news.Spec.K0s.Config.Spec.Telemetry == nil {
		news.Spec.K0s.Config.Spec.Telemetry = &K0sTelemetry{}
	}

	if newsClusterCopy.Spec != nil && newsClusterCopy.Spec.K0s != nil && newsClusterCopy.Spec.K0s.Config != nil {
		if newsClusterCopy.Spec.K0s.Config.Spec != nil {
			if newsClusterCopy.Spec.K0s.Config.Spec.Telemetry != nil {
				news.Spec.K0s.Config.Spec.Telemetry.Enabled = newsClusterCopy.Spec.K0s.Config.Spec.Telemetry.Enabled
			}
		}
	}

	return nil
}

func (c *k0sctlCluster) APIAddress() string {
	address := "localhost"
	port := 6443

	clt, cleanup, err := c.k0sctl()

	defer cleanup()

	if err != nil {
		return fmt.Sprintf("https://%s:%d", address, port)
	}

	leader := clt.Spec.K0sLeader()
	if leader != nil {
		switch {
		case leader.SSH != nil:
			address = leader.SSH.Address
		case leader.OpenSSH != nil:
			address = leader.OpenSSH.Address
		case leader.WinRM != nil:
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

func (c *k0sctlCluster) k0sctl() (*v1beta1.Cluster, func(), error) {
	prefix := fmt.Sprintf("%s-%d", *c.Metadata.Name, rand.Intn(int(^uint(0)>>1)))

	cleanup := func() {
		_ = cleanupTempFiles(prefix)
	}

	bytes, err := yamlJSON.Marshal(c)
	if err != nil {
		return nil, cleanup, err
	}

	var clt v1beta1.Cluster

	if err := yaml.Unmarshal(bytes, &clt); err != nil {
		return nil, cleanup, err
	}

	// ensure correct types of spec.k0s.config
	if clt.Spec == nil {
		clt.Spec = &cluster.Spec{}
	}

	if clt.Spec.Hosts == nil {
		clt.Spec.Hosts = cluster.Hosts{}
	}

	if clt.Spec.K0s == nil {
		clt.Spec.K0s = &cluster.K0s{}
	}

	if clt.Spec.K0s.Config == nil {
		clt.Spec.K0s.Config = dig.Mapping{}
	}

	if clt.Spec.K0s.Config["metadata"] == nil {
		clt.Spec.K0s.Config["metadata"] = dig.Mapping{}
	}

	if m, ok := clt.Spec.K0s.Config.Dig("metadata").(dig.Mapping); ok {
		m["name"] = *c.Metadata.Name
	}

	if clt.Spec.K0s.Config["spec"] == nil {
		clt.Spec.K0s.Config["spec"] = dig.Mapping{}
	}

	// replace inline values with file paths
	if err := c.k0sctlConvertHostsPaths(prefix, clt.Spec.Hosts); err != nil {
		return nil, cleanup, err
	}

	spec := c.Spec.K0s.Config.Spec
	if spec.Storage != nil && spec.Storage.Etcd != nil && spec.Storage.Etcd.ExternalCluster != nil {
		ec := clt.Spec.K0s.Config.Dig("spec", "storage", "etcd", "externalCluster")

		if ec != nil {
			if ec, ok := ec.(*k0sapi.ExternalCluster); ok {
				if err := c.k0sctlConvertExternalEtcdPaths(prefix, ec); err != nil {
					return nil, cleanup, err
				}
			}
		}
	}

	return &clt, cleanup, nil
}

func (c *k0sctlCluster) k0sctlConvertHostsPaths(prefix string, hosts cluster.Hosts) error {
	for i, host := range hosts {
		clusterHost := c.Spec.Hosts[i]

		if host.SSH != nil {
			if err := c.k0sctlConvertSSHPaths(prefix, clusterHost.SSH, host.SSH); err != nil {
				return err
			}
		}

		if host.OpenSSH != nil {
			if err := c.k0sctlConvertOpenSSHPaths(prefix, clusterHost.OpenSSH, host.OpenSSH); err != nil {
				return err
			}
		}

		if host.WinRM != nil {
			if clusterHost.WinRM.CaCert != nil {
				filename, err := contentToTempFile(prefix, *clusterHost.WinRM.CaCert, true)
				if err != nil {
					return err
				}

				host.WinRM.CACertPath = filename
			}

			if clusterHost.WinRM.Cert != nil {
				filename, err := contentToTempFile(prefix, *clusterHost.WinRM.Cert, true)
				if err != nil {
					return err
				}

				host.WinRM.CertPath = filename
			}

			if clusterHost.WinRM.Key != nil {
				filename, err := contentToTempFile(prefix, *clusterHost.WinRM.Key, true)
				if err != nil {
					return err
				}

				host.WinRM.KeyPath = filename
			}

			if err := c.k0sctlConvertSSHPaths(prefix, clusterHost.WinRM.Bastion, host.WinRM.Bastion); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *k0sctlCluster) k0sctlConvertSSHPaths(prefix string, ssh *ClusterSSH, rigSSH *rig.SSH) error {
	if ssh.Key != nil {
		filename, err := contentToTempFile(prefix, *ssh.Key, true)
		if err != nil {
			return err
		}

		rigSSH.KeyPath = &filename
	}

	if ssh.Bastion != nil {
		return c.k0sctlConvertSSHPaths(prefix, ssh.Bastion, rigSSH.Bastion)
	}

	return nil
}

func (c *k0sctlCluster) k0sctlConvertOpenSSHPaths(prefix string, ssh *ClusterOpenSSH, rigSSH *rig.OpenSSH) error {
	if ssh.Key != nil {
		filename, err := contentToTempFile(prefix, *ssh.Key, true)
		if err != nil {
			return err
		}

		rigSSH.KeyPath = &filename
	}

	return nil
}

func (c *k0sctlCluster) k0sctlConvertExternalEtcdPaths(
	prefix string,
	k0sctlExternalCluster *k0sapi.ExternalCluster,
) error {
	externalCluster := c.Spec.K0s.Config.Spec.Storage.Etcd.ExternalCluster

	if externalCluster == nil || k0sctlExternalCluster == nil {
		return nil
	}

	if externalCluster.CA != nil {
		filename, err := contentToTempFile(prefix, *externalCluster.CA, true)
		if err != nil {
			return err
		}

		k0sctlExternalCluster.CaFile = filename
	}

	if externalCluster.ClientCert != nil {
		filename, err := contentToTempFile(prefix, *externalCluster.ClientCert, true)
		if err != nil {
			return err
		}

		k0sctlExternalCluster.ClientCertFile = filename
	}

	if externalCluster.ClientKey != nil {
		filename, err := contentToTempFile(prefix, *externalCluster.ClientKey, true)
		if err != nil {
			return err
		}

		k0sctlExternalCluster.ClientKeyFile = filename
	}

	return nil
}

// DefaultCluster returns a default cluster configuration.
// see https://github.com/k0sproject/k0sctl/blob/main/cmd/init.go
func clusterDefaults(name string) (*k0sctlCluster, error) {
	k0sVersion, err := version.LatestStable()
	if err != nil {
		return nil, err
	}

	clt := v1beta1.Cluster{
		APIVersion: v1beta1.APIVersion,
		Kind:       "Cluster",
		Metadata:   &v1beta1.ClusterMetadata{Name: name},
		Spec: &cluster.Spec{
			Hosts: []*cluster.Host{},
			K0s: &cluster.K0s{
				Version: k0sVersion,
				Config:  dig.Mapping{},
			},
		},
	}

	if err := yaml.Unmarshal(cmd.DefaultK0sYaml, &clt.Spec.K0s.Config); err != nil {
		return nil, err
	}

	yamlData, err := yaml.Marshal(&clt)
	if err != nil {
		return nil, err
	}

	var cluster k0sctlCluster

	if err := yamlJSON.Unmarshal(yamlData, &cluster); err != nil {
		return nil, err
	}

	return &cluster, nil
}
