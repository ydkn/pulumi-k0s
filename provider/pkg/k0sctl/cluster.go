package k0sctl

import (
	"fmt"

	"github.com/k0sproject/dig"
	"github.com/k0sproject/k0sctl/pkg/apis/k0sctl.k0sproject.io/v1beta1"
	"github.com/k0sproject/k0sctl/pkg/apis/k0sctl.k0sproject.io/v1beta1/cluster"
	"github.com/k0sproject/rig"
	"gopkg.in/yaml.v2"
)

type Cluster struct {
	Metadata   *Metadata `json:"metadata,omitempty" yaml:"metadata,omitempty" validate:"required"`
	Spec       *Spec     `json:"spec,omitempty" yaml:"spec,omitempty" validate:"required"`
	Kubeconfig string    `json:"kubeconfig,omitempty" yaml:"-"`
}

func (o *Cluster) K0sCtlObject() *v1beta1.Cluster {
	cluster := &v1beta1.Cluster{
		APIVersion: v1beta1.APIVersion,
		Kind:       "Cluster",
		Metadata:   o.Metadata.K0sCtlObject(),
		Spec:       o.Spec.K0sCtlObject(),
	}

	cfg, err := yaml.Marshal(cluster)
	if err != nil {
		return nil
	}

	if err := yaml.Unmarshal(cfg, cluster); err != nil {
		return nil
	}

	return cluster
}

func (o *Cluster) Check() error {
	if err := o.K0sCtlObject().Validate(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

type Metadata struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty" default:"k0s-cluster" validate:"required"`
}

func (o *Metadata) K0sCtlObject() *v1beta1.ClusterMetadata {
	return &v1beta1.ClusterMetadata{Name: o.Name}
}

type Spec struct {
	Hosts Hosts `json:"hosts,omitempty" yaml:"hosts,omitempty" validate:"required"`
	K0s   *K0s  `json:"k0s,omitempty" yaml:"k0s,omitempty"`
}

func (o *Spec) K0sCtlObject() *cluster.Spec {
	return &cluster.Spec{Hosts: o.Hosts.K0sCtlObject(), K0s: o.K0s.K0sCtlObject()}
}

type Hosts []*Host

func (o *Hosts) K0sCtlObject() cluster.Hosts {
	hosts := make(cluster.Hosts, len(*o))

	for i, host := range *o {
		hosts[i] = host.K0sCtlObject()
	}

	return hosts
}

type Host struct {
	WinRM     *WinRM     `json:"winRM,omitempty" yaml:"winRM,omitempty"`
	SSH       *SSH       `json:"ssh,omitempty" yaml:"ssh,omitempty"`
	Localhost *Localhost `json:"localhost,omitempty" yaml:"localhost,omitempty"`

	Role             string            `json:"role,omitempty" yaml:"role,omitempty" default:"worker" validate:"required"`
	PrivateInterface string            `json:"privateInterface,omitempty" yaml:"privateInterface,omitempty"`
	PrivateAddress   string            `json:"privateAddress,omitempty" yaml:"privateAddress,omitempty"`
	Environment      map[string]string `json:"environment,omitempty" yaml:"environment,flow,omitempty"`
	UploadBinary     bool              `json:"uploadBinary,omitempty" yaml:"uploadBinary,omitempty"`
	K0sBinaryPath    string            `json:"k0sBinaryPath,omitempty" yaml:"k0sBinaryPath,omitempty"`
	InstallFlags     []string          `json:"installFlags,omitempty" yaml:"installFlags,omitempty"`
	Files            []*UploadFile     `json:"files,omitempty" yaml:"files,omitempty"`
	OSIDOverride     string            `json:"os,omitempty" yaml:"os,omitempty"`
	HostnameOverride string            `json:"hostname,omitempty" yaml:"hostname,omitempty"`
	Hooks            Hooks             `json:"hooks,omitempty" yaml:"hooks,omitempty"`
}

func (o *Host) K0sCtlObject() *cluster.Host {
	host := &cluster.Host{
		Role:             o.Role,
		PrivateInterface: o.PrivateInterface,
		PrivateAddress:   o.PrivateAddress,
		Environment:      o.Environment,
		UploadBinary:     o.UploadBinary,
		K0sBinaryPath:    o.K0sBinaryPath,
		InstallFlags:     o.InstallFlags,
		OSIDOverride:     o.OSIDOverride,
		HostnameOverride: o.HostnameOverride,
	}

	if o.WinRM != nil {
		host.WinRM = o.WinRM.K0sCtlObject()
	}

	if o.SSH != nil {
		host.SSH = o.SSH.K0sCtlObject()
	}

	if o.Localhost != nil {
		host.Localhost = o.Localhost.K0sCtlObject()
	}

	if o.Files != nil {
		host.Files = make([]*cluster.UploadFile, len(o.Files))

		for i, file := range o.Files {
			host.Files[i] = file.K0sCtlObject()
		}
	}

	if o.Hooks != nil {
		host.Hooks = o.Hooks.K0sCtlObject()
	}

	return host
}

type WinRM struct {
	Address       string `json:"address,omitempty" yaml:"address,omitempty" validate:"required,hostname|ip"`
	User          string `json:"user,omitempty" yaml:"user,omitempty" default:"Administrator" validate:"omitempty,gt=2"`
	Port          int    `json:"port,omitempty" yaml:"port,omitempty" default:"5985" validate:"gt=0,lte=65535"`
	Password      string `json:"password,omitempty" yaml:"password,omitempty"`
	UseHTTPS      bool   `json:"useHTTPS,omitempty" yaml:"useHTTPS,omitempty" default:"false"`
	Insecure      bool   `json:"insecure,omitempty" yaml:"insecure,omitempty" default:"false"`
	UseNTLM       bool   `json:"useNTLM,omitempty" yaml:"useNTLM,omitempty" default:"false"`
	CACertPath    string `json:"caCertPath,omitempty" yaml:"caCertPath,omitempty" validate:"omitempty,file"`
	CertPath      string `json:"certPath,omitempty" yaml:"certPath,omitempty" validate:"omitempty,file"`
	KeyPath       string `json:"keyPath,omitempty" yaml:"keyPath,omitempty" validate:"omitempty,file"`
	TLSServerName string `json:"tlsServerName,omitempty" yaml:"tlsServerName,omitempty" validate:"omitempty,hostname|ip"`
	Bastion       *SSH   `json:"bastion,omitempty" yaml:"bastion,omitempty"`
}

func (o *WinRM) K0sCtlObject() *rig.WinRM {
	ssh := &rig.WinRM{
		Address:       o.Address,
		User:          o.User,
		Port:          o.Port,
		Password:      o.Password,
		UseHTTPS:      o.UseHTTPS,
		Insecure:      o.Insecure,
		UseNTLM:       o.UseNTLM,
		CACertPath:    o.CACertPath,
		CertPath:      o.CertPath,
		KeyPath:       o.KeyPath,
		TLSServerName: o.TLSServerName,
	}

	if o.Bastion != nil {
		ssh.Bastion = o.Bastion.K0sCtlObject()
	}

	return ssh
}

type SSH struct {
	Address string `json:"address,omitempty" yaml:"address,omitempty" validate:"required,hostname|ip"`
	User    string `json:"user,omitempty" yaml:"user,omitempty" default:"root" validate:"required"`
	Port    int    `json:"port,omitempty" yaml:"port,omitempty" default:"22" validate:"gt=0,lte=65535"`
	KeyPath string `json:"keyPath,omitempty" yaml:"keyPath,omitempty" validate:"omitempty"`
	HostKey string `json:"hostKey,omitempty" yaml:"hostKey,omitempty"`
	Bastion *SSH   `json:"bastion,omitempty" yaml:"bastion,omitempty"`
}

func (o *SSH) K0sCtlObject() *rig.SSH {
	ssh := &rig.SSH{
		Address: o.Address,
		User:    o.User,
		Port:    o.Port,
		KeyPath: o.KeyPath,
		HostKey: o.HostKey,
	}

	if o.Bastion != nil {
		ssh.Bastion = o.Bastion.K0sCtlObject()
	}

	return ssh
}

type Localhost struct {
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty" default:"true" validate:"required,eq=true"`
}

func (o *Localhost) K0sCtlObject() *rig.Localhost {
	return &rig.Localhost{Enabled: o.Enabled}
}

type UploadFile struct {
	Name            string      `json:"name,omitempty" yaml:"name,omitempty"`
	Source          string      `json:"src,omitempty" yaml:"src,omitempty"`
	DestinationDir  string      `json:"dstDir,omitempty" yaml:"dstDir,omitempty"`
	DestinationFile string      `json:"dst,omitempty" yaml:"dst,omitempty"`
	PermMode        interface{} `json:"perm,omitempty" yaml:"perm,omitempty"`
	DirPermMode     interface{} `json:"dirPerm,omitempty" yaml:"dirPerm,omitempty"`
	User            string      `json:"user,omitempty" yaml:"user,omitempty"`
	Group           string      `json:"group,omitempty" yaml:"group,omitempty"`
}

func (o *UploadFile) K0sCtlObject() *cluster.UploadFile {
	return &cluster.UploadFile{
		Name:            o.Name,
		Source:          o.Source,
		DestinationDir:  o.DestinationDir,
		DestinationFile: o.DestinationFile,
		PermMode:        o.PermMode,
		DirPermMode:     o.DirPermMode,
		User:            o.User,
		Group:           o.Group,
	}
}

type Hooks map[string]map[string][]string

func (o *Hooks) K0sCtlObject() cluster.Hooks {
	return cluster.Hooks(*o)
}

type K0s struct {
	Version       string      `json:"version,omitempty" yaml:"version,omitempty"`
	DynamicConfig bool        `json:"dynamicConfig,omitempty" yaml:"dynamicConfig,omitempty" default:"false"`
	Config        dig.Mapping `json:"config,omitempty" yaml:"config,omitempty"`
}

func (o *K0s) K0sCtlObject() *cluster.K0s {
	return &cluster.K0s{Version: o.Version, DynamicConfig: o.DynamicConfig, Config: o.Config}
}
