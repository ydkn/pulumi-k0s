// Copyright 2025, Florian Schwab.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"fmt"
	"math/rand"

	"github.com/k0sproject/dig"
	k0sapi "github.com/k0sproject/k0s/pkg/apis/k0s/v1beta1"
	"github.com/k0sproject/k0sctl/pkg/apis/k0sctl.k0sproject.io/v1beta1"
	"github.com/k0sproject/k0sctl/pkg/apis/k0sctl.k0sproject.io/v1beta1/cluster"
	"github.com/k0sproject/rig/v2/protocol/openssh"
	"github.com/k0sproject/rig/v2/protocol/ssh"
	"github.com/k0sproject/version"
	"github.com/pulumi/pulumi-go-provider/infer"
	"gopkg.in/yaml.v2"
	yamlJSON "sigs.k8s.io/yaml"
)

const (
	defaultAPIVersion            = v1beta1.APIVersion
	defaultKind                  = "Cluster"
	defaultClusterName           = "k0s"
	defaultVersionChannel        = "stable"
	defaultDynamicConfig         = false
	defaultK0sAPIVersion         = "k0s.k0sproject.io/v1beta1"
	defaultK0sKind               = "Cluster"
	defaultK0sName               = "k0s"
	defaultAPIPort               = 6443
	defaultAPIK0sAPIPort         = 9443
	defaultEtcdUser              = "etcd"
	defaultKineUser              = "kube-apiserver"
	defaultKonnectivityUser      = "konnectivity-server"
	defaultKubeAPIserverUser     = "kube-apiserver"
	defaultKubeSchedulerUser     = "kube-scheduler"
	defaultKonnectivityAgentPort = 8132
	defaultKonnectivityAdminPort = 8133
	defaultNetworkPodCIDR        = "10.244.0.0/16"
	defaultNetworkServiceCIDR    = "10.96.0.0/12"
	defaultNetworkProvider       = "kuberouter"
	defaultKubeRouterMTU         = 0
	defaultKubeRouterAutoMTU     = true
	defaultKubeProxyDisabled     = false
	defaultKubeProxyMode         = "iptables"
	defaultPodSecurityPolicy     = "00-k0s-privileged"
	defaultStorageType           = "etcd"
	defaultTelemetryEnabled      = true
)

type ClusterArgs struct {
	APIVersion *string          `pulumi:"apiVersion,optional" json:"apiVersion,omitempty"`
	Kind       *string          `pulumi:"kind,optional" json:"kind,omitempty"`
	Metadata   *ClusterMetadata `pulumi:"metadata,optional" json:"metadata,omitempty"`
	Spec       *ClusterSpec     `pulumi:"spec,optional" json:"spec,omitempty"`
}

func (f *ClusterArgs) Annotate(a infer.Annotator) {
	apiVersion := defaultAPIVersion
	a.Describe(&f.APIVersion, "The API version of the cluster resource.")
	a.SetDefault(&f.APIVersion, &apiVersion)

	kind := defaultKind
	a.Describe(&f.Kind, "The kind of the cluster resource.")
	a.SetDefault(&f.Kind, &kind)

	a.Describe(&f.Metadata, "Metadata of the cluster resource.")
	//a.SetDefault(&f.Metadata, &ClusterMetadata{})

	a.Describe(&f.Spec, "Specification of the cluster resource.")
	//a.SetDefault(&f.Spec, &ClusterSpec{})
}

type ClusterMetadata struct {
	Name *string `pulumi:"name,optional" json:"name,omitempty"`
}

func (f *ClusterMetadata) Annotate(a infer.Annotator) {
	name := defaultClusterName
	a.Describe(&f.Name, "The name of the cluster.")
	a.SetDefault(&f.Name, &name)
}

type ClusterSpec struct {
	Hosts []*ClusterHost `pulumi:"hosts" json:"hosts"`
	K0s   *ClusterK0s    `pulumi:"k0s,optional" json:"k0s,omitempty"`
}

func (f *ClusterSpec) Annotate(a infer.Annotator) {
	a.Describe(&f.Hosts, "The hosts that will form the cluster.")
	//a.SetDefault(&f.Hosts, []*ClusterHost{})

	a.Describe(&f.K0s, "K0s configuration.")
	//a.SetDefault(&f.K0s, &ClusterK0s{})
}

type ClusterHost struct {
	Role             string            `pulumi:"role" json:"role"`
	PrivateInterface *string           `pulumi:"privateInterface,optional" json:"privateInterface,omitempty"`
	PrivateAddress   *string           `pulumi:"privateAddress,optional" json:"privateAddress,omitempty"`
	Environment      map[string]string `pulumi:"environment,optional" json:"environment,omitempty"`
	UploadBinary     *bool             `pulumi:"uploadBinary,optional" json:"uploadBinary,omitempty"`
	K0sBinaryPath    *string           `pulumi:"k0sBinaryPath,optional" json:"k0sBinaryPath,omitempty"`
	InstallFlags     []string          `pulumi:"installFlags,optional" json:"installFlags,omitempty"`
	Files            []ClusterFile     `pulumi:"files,optional" json:"files,omitempty"`
	OS               *string           `pulumi:"os,optional" json:"os,omitempty"`
	Hostname         *string           `pulumi:"hostname,optional" json:"hostname,omitempty"`
	Hooks            *ClusterHooks     `pulumi:"hooks,optional" json:"hooks,omitempty"`
	WinRM            *ClusterWinRM     `pulumi:"winRM,optional" json:"winRM,omitempty"`
	SSH              *ClusterSSH       `pulumi:"ssh,optional" json:"ssh,omitempty"`
	OpenSSH          *ClusterOpenSSH   `pulumi:"openSSH,optional" json:"openSSH,omitempty"`
	Localhost        *ClusterLocalhost `pulumi:"localhost,optional" json:"localhost,omitempty"`
	NoTaints         *bool             `pulumi:"noTaints,optional" json:"noTaints,omitempty"`
}

type ClusterFile struct {
	Name                 *string `pulumi:"name,optional" json:"name,omitempty"`
	Source               *string `pulumi:"src,optional" json:"src,omitempty"`
	DestinationDirectory *string `pulumi:"dstDir,optional" json:"dstDir,omitempty"`
	Destination          *string `pulumi:"dst,optional" json:"dst,omitempty"`
	Permissions          *string `pulumi:"perm,optional" json:"perm,omitempty"`
	DirectoryPermissions *string `pulumi:"dirPerm,optional" json:"dirPerm,omitempty"`
	User                 *string `pulumi:"user,optional" json:"user,omitempty"`
	Group                *string `pulumi:"group,optional" json:"group,omitempty"`
}

type ClusterHooks struct {
	Apply  *ClusterHook `pulumi:"apply,optional" json:"apply,omitempty"`
	Backup *ClusterHook `pulumi:"backup,optional" json:"backup,omitempty"`
	Reset  *ClusterHook `pulumi:"reset,optional" json:"reset,omitempty"`
}

type ClusterHook struct {
	Before []string `pulumi:"before,optional" json:"before,omitempty"`
	After  []string `pulumi:"after,optional" json:"after,omitempty"`
}

type ClusterBastion struct {
	Address string  `pulumi:"address" json:"address,omitempty"`
	Port    *int    `pulumi:"port,optional" json:"port,omitempty"`
	User    *string `pulumi:"user,optional" json:"user,omitempty"`
	Key     *string `pulumi:"key,optional" provider:"secret" json:"-"`
	HostKey *string `pulumi:"hostKey,optional" json:"hostKey,omitempty"`
}

type ClusterWinRM struct {
	Address       string          `pulumi:"address" json:"address"`
	Port          *int            `pulumi:"port,optional" json:"port,omitempty"`
	User          *string         `pulumi:"user,optional" json:"user,omitempty"`
	Password      *string         `pulumi:"password,optional" provider:"secret" json:"password,omitempty"`
	UseHTTPS      *bool           `pulumi:"useHTTPS,optional" json:"useHTTPS,omitempty"`
	Insecure      *bool           `pulumi:"insecure,optional" json:"insecure,omitempty"`
	UseNTLM       *bool           `pulumi:"useNTLM,optional" json:"useNTLM,omitempty"`
	CaCert        *string         `pulumi:"caCert,optional" provider:"secret" json:"caCert,omitempty"`
	Cert          *string         `pulumi:"cert,optional" provider:"secret" json:"cert,omitempty"`
	Key           *string         `pulumi:"key,optional" provider:"secret" json:"key,omitempty"`
	TLSServerName *string         `pulumi:"tlsServerName,optional" json:"tlsServerName,omitempty"`
	Bastion       *ClusterBastion `pulumi:"bastion,optional" json:"bastion,omitempty"`
}

type ClusterSSH struct {
	Address string          `pulumi:"address" json:"address,omitempty"`
	Port    *int            `pulumi:"port,optional" json:"port,omitempty"`
	User    *string         `pulumi:"user,optional" json:"user,omitempty"`
	Key     *string         `pulumi:"key,optional" provider:"secret" json:"-"`
	HostKey *string         `pulumi:"hostKey,optional" json:"hostKey,omitempty"`
	Bastion *ClusterBastion `pulumi:"bastion,optional" json:"bastion,omitempty"`
}

type ClusterOpenSSH struct {
	Address             string            `pulumi:"address" json:"address,omitempty"`
	Port                *int              `pulumi:"port,optional" json:"port,omitempty"`
	User                *string           `pulumi:"user,optional" json:"user,omitempty"`
	Key                 *string           `pulumi:"key,optional" provider:"secret" json:"-"`
	ConfigPath          *string           `pulumi:"configPath,optional" json:"configPath,omitempty"`
	Options             map[string]string `pulumi:"options,optional" json:"options,omitempty"`
	DisableMultiplexing *bool             `pulumi:"disableMultiplexing,optional" json:"disableMultiplexing,omitempty"`
}

type ClusterLocalhost struct {
	Enabled *bool `pulumi:"enabled,optional" json:"enabled,omitempty"`
}

type ClusterK0s struct {
	Version        *string `pulumi:"version,optional" json:"version,omitempty"`
	VersionChannel *string `pulumi:"versionChannel,optional" json:"versionChannel,omitempty"`
	DynamicConfig  *bool   `pulumi:"dynamicConfig,optional" json:"dynamicConfig,omitempty"`
	Config         *K0s    `pulumi:"config,optional" json:"config,omitempty"`
}

func (f *ClusterK0s) Annotate(a infer.Annotator) {
	a.Describe(&f.Version, "The hosts that will form the cluster.")

	versionChannel := defaultVersionChannel
	a.Describe(&f.VersionChannel, "The k0s version channel to use.")
	a.SetDefault(&f.VersionChannel, &versionChannel)

	dynamicConfig := defaultDynamicConfig
	a.Describe(&f.DynamicConfig, "Whether to use dynamic configuration.")
	a.SetDefault(&f.DynamicConfig, &dynamicConfig)

	a.Describe(&f.Config, "K0s configuration.")
	//a.SetDefault(&f.Config, &K0s{})
}

type K0s struct {
	APIVersion *string      `pulumi:"apiVersion,optional" json:"apiVersion,omitempty"`
	Kind       *string      `pulumi:"kind,optional" json:"kind,omitempty"`
	Metadata   *K0sMetadata `pulumi:"metadata,optional" json:"metadata,omitempty"`
	Spec       *K0sSpec     `pulumi:"spec,optional" json:"spec,omitempty"`
}

func (f *K0s) Annotate(a infer.Annotator) {
	apiVersion := defaultK0sAPIVersion
	a.Describe(&f.APIVersion, "The API version of the k0s configuration.")
	a.SetDefault(&f.APIVersion, &apiVersion)

	kind := defaultK0sKind
	a.Describe(&f.Kind, "The kind of the k0s configuration.")
	a.SetDefault(&f.Kind, &kind)

	a.Describe(&f.Metadata, "K0s configuration metadata.")
	//a.SetDefault(&f.Metadata, &K0sMetadata{})

	a.Describe(&f.Spec, "K0s configuration specification.")
	//a.SetDefault(&f.Spec, &K0sSpec{})
}

type K0sMetadata struct {
	Name *string `pulumi:"name,optional" json:"name,omitempty"`
}

func (f *K0sMetadata) Annotate(a infer.Annotator) {
	name := defaultK0sName
	a.Describe(&f.Name, "The name of the k0s cluster.")
	a.SetDefault(&f.Name, &name)
}

type K0sSpec struct {
	API               *K0sAPI               `pulumi:"api,optional" json:"api,omitempty"`
	Images            *K0sImages            `pulumi:"images,optional" json:"images,omitempty"`
	InstallConfig     *K0sInstallConfig     `pulumi:"installConfig,optional" json:"installConfig,omitempty"`
	Konnectivity      *K0sKonnectivity      `pulumi:"konnectivity,optional" json:"konnectivity,omitempty"`
	Network           *K0sNetwork           `pulumi:"network,optional" json:"network,omitempty"`
	PodSecurityPolicy *K0sPodSecurityPolicy `pulumi:"podSecurityPolicy,optional" json:"podSecurityPolicy,omitempty"`
	ControllerManager *K0sControllerManager `pulumi:"controllerManager,optional" json:"controllerManager,omitempty"`
	Scheduler         *K0sScheduler         `pulumi:"scheduler,optional" json:"scheduler,omitempty"`
	Storage           *K0sStorage           `pulumi:"storage,optional" json:"storage,omitempty"`
	WorkerProfiles    []*K0sWorkerProfile   `pulumi:"workerProfiles,optional" json:"workerProfiles,omitempty"`
	FeatureGates      []*K0sFeatureGate     `pulumi:"featureGates,optional" json:"featureGates,omitempty"`
	Telemetry         *K0sTelemetry         `pulumi:"telemetry,optional" json:"telemetry,omitempty"`
}

func (f *K0sSpec) Annotate(a infer.Annotator) {
	a.Describe(&f.API, "K0s API configuration.")
	//a.SetDefault(&f.API, &K0sAPI{})

	a.Describe(&f.Images, "K0s images configuration.")
	//a.SetDefault(&f.Images, &K0sImages{})

	a.Describe(&f.InstallConfig, "K0s install configuration.")
	//a.SetDefault(&f.InstallConfig, &K0sInstallConfig{})

	a.Describe(&f.Konnectivity, "K0s konnectivity configuration.")
	//a.SetDefault(&f.Konnectivity, &K0sKonnectivity{})

	a.Describe(&f.Network, "K0s network configuration.")
	//a.SetDefault(&f.Network, &K0sNetwork{})

	a.Describe(&f.PodSecurityPolicy, "K0s pod security policy configuration.")
	//a.SetDefault(&f.PodSecurityPolicy, &K0sPodSecurityPolicy{})

	a.Describe(&f.ControllerManager, "K0s controller manager configuration.")
	//a.SetDefault(&f.ControllerManager, &K0sControllerManager{})

	a.Describe(&f.Scheduler, "K0s scheduler configuration.")
	//a.SetDefault(&f.Scheduler, &K0sScheduler{})

	a.Describe(&f.Storage, "K0s storage configuration.")
	//a.SetDefault(&f.Storage, &K0sStorage{})

	a.Describe(&f.WorkerProfiles, "K0s worker profiles configuration.")
	//a.SetDefault(&f.WorkerProfiles, []*K0sWorkerProfile{})

	a.Describe(&f.FeatureGates, "K0s feature gates configuration.")
	//a.SetDefault(&f.FeatureGates, []*K0sFeatureGate{})

	a.Describe(&f.Telemetry, "K0s telemetry configuration.")
	//a.SetDefault(&f.Telemetry, &K0sTelemetry{})
}

type K0sAPI struct {
	Address         *string           `pulumi:"address,optional" json:"address,omitempty"`
	Port            *int              `pulumi:"port,optional" json:"port,omitempty"`
	K0sApiPort      *int              `pulumi:"k0sApiPort,optional" json:"k0sApiPort,omitempty"`
	ExternalAddress *string           `pulumi:"externalAddress,optional" json:"externalAddress,omitempty"`
	SANs            []string          `pulumi:"sans,optional" json:"sans,omitempty"`
	ExtraArgs       map[string]string `pulumi:"extraArgs,optional" json:"extraArgs,omitempty"`
}

func (f *K0sAPI) Annotate(a infer.Annotator) {
	port := defaultAPIPort
	a.Describe(&f.Port, "The port the Kubernetes API will listen on.")
	a.SetDefault(&f.Port, &port)

	k0sApiPort := defaultAPIK0sAPIPort
	a.Describe(&f.K0sApiPort, "The port the k0s API will listen on.")
	a.SetDefault(&f.K0sApiPort, &k0sApiPort)
}

type K0sImages struct {
	DefaultPullPolicy *string             `pulumi:"defaultPullPolicy,optional" json:"default_pull_policy,omitempty"`
	Repository        *string             `pulumi:"repository,optional" json:"repository,omitempty"`
	Konnectivity      *K0sImage           `pulumi:"konnectivity,optional" json:"konnectivity,omitempty"`
	MetricsServer     *K0sImage           `pulumi:"metricsserver,optional" json:"metricsserver,omitempty"`
	Kubeproxy         *K0sImage           `pulumi:"kubeproxy,optional" json:"kubeproxy,omitempty"`
	CoreDNS           *K0sImage           `pulumi:"coredns,optional" json:"coredns,omitempty"`
	Pause             *K0sImage           `pulumi:"pause,optional" json:"pause,omitempty"`
	Calico            *K0sCalicoImage     `pulumi:"calico,optional" json:"calico,omitempty"`
	KubeRouter        *K0sKubeRouterImage `pulumi:"kuberouter,optional" json:"kuberouter,omitempty"`
}

type K0sImage struct {
	Image   *string `pulumi:"image,optional" json:"image,omitempty"`
	Version *string `pulumi:"version,optional" json:"version,omitempty"`
}

type K0sCalicoImage struct {
	CNI             *K0sImage `pulumi:"cni,optional" json:"cni,omitempty"`
	FlexVolume      *K0sImage `pulumi:"flexvolume,optional" json:"flexvolume,omitempty"`
	Node            *K0sImage `pulumi:"node,optional" json:"node,omitempty"`
	KubeControllers *K0sImage `pulumi:"kubecontrollers,optional" json:"kubecontrollers,omitempty"`
}

type K0sKubeRouterImage struct {
	CNI          *K0sImage `pulumi:"cni,optional" json:"cni,omitempty"`
	CNIInstaller *K0sImage `pulumi:"cniInstaller,optional" json:"cniInstaller,omitempty"`
}

type K0sInstallConfig struct {
	Users *K0sInstallConfigUser `pulumi:"users,optional" json:"users,omitempty"`
}

func (f *K0sInstallConfig) Annotate(a infer.Annotator) {
	a.Describe(&f.Users, "K0s install configuration users.")
	//a.SetDefault(&f.Users, &K0sInstallConfigUser{})
}

type K0sInstallConfigUser struct {
	EtcdUser          *string `pulumi:"etcdUser,optional" json:"etcdUser,omitempty"`
	KineUser          *string `pulumi:"kineUser,optional" json:"kineUser,omitempty"`
	KonnectivityUser  *string `pulumi:"konnectivityUser,optional" json:"konnectivityUser,omitempty"`
	KubeAPIServerUser *string `pulumi:"kubeAPIserverUser,optional" json:"kubeAPIserverUser,omitempty"`
	KubeSchedulerUser *string `pulumi:"kubeSchedulerUser,optional" json:"kubeSchedulerUser,omitempty"`
}

func (f *K0sInstallConfigUser) Annotate(a infer.Annotator) {
	etcdUser := defaultEtcdUser
	a.Describe(&f.EtcdUser, "The user the etcd process will run as.")
	a.SetDefault(&f.EtcdUser, &etcdUser)

	kineUser := defaultKineUser
	a.Describe(&f.KineUser, "The user the kine process will run as.")
	a.SetDefault(&f.KineUser, &kineUser)

	konnectivityUser := defaultKonnectivityUser
	a.Describe(&f.KonnectivityUser, "The user the konnectivity process will run as.")
	a.SetDefault(&f.KonnectivityUser, &konnectivityUser)

	kubeAPIserverUser := defaultKubeAPIserverUser
	a.Describe(&f.KubeAPIServerUser, "The user the kube-apiserver process will run as.")
	a.SetDefault(&f.KubeAPIServerUser, &kubeAPIserverUser)

	kubeSchedulerUser := defaultKubeSchedulerUser
	a.Describe(&f.KubeSchedulerUser, "The user the kube-scheduler process will run as.")
	a.SetDefault(&f.KubeSchedulerUser, &kubeSchedulerUser)
}

type K0sKonnectivity struct {
	AdminPort *int `pulumi:"adminPort,optional" json:"adminPort,omitempty"`
	AgentPort *int `pulumi:"agentPort,optional" json:"agentPort,omitempty"`
}

func (f *K0sKonnectivity) Annotate(a infer.Annotator) {
	agentPort := defaultKonnectivityAgentPort
	a.Describe(&f.AgentPort, "The port the konnectivity agent will listen on.")
	a.SetDefault(&f.AgentPort, &agentPort)

	adminPort := defaultKonnectivityAdminPort
	a.Describe(&f.AdminPort, "The port the konnectivity admin server will listen on.")
	a.SetDefault(&f.AdminPort, &adminPort)
}

type K0sNetwork struct {
	Provider               *string                    `pulumi:"provider,optional" json:"provider,omitempty"`
	PodCIDR                *string                    `pulumi:"podCIDR,optional" json:"podCIDR,omitempty"`
	ServiceCIDR            *string                    `pulumi:"serviceCIDR,optional" json:"serviceCIDR,omitempty"`
	ClusterDomain          *string                    `pulumi:"clusterDomain,optional" json:"clusterDomain,omitempty"`
	DualStack              *K0sDualStack              `pulumi:"dualStack,optional" json:"dualStack,omitempty"`
	Calico                 *K0sCalico                 `pulumi:"calico,optional" json:"calico,omitempty"`
	KubeRouter             *K0sKubeRouter             `pulumi:"kuberouter,optional" json:"kuberouter,omitempty"`
	KubeProxy              *K0sKubeProxy              `pulumi:"kubeProxy,optional" json:"kubeProxy,omitempty"`
	NodeLocalLoadBalancing *K0sNodeLocalLoadBalancing `pulumi:"nodeLocalLoadBalancing,optional" json:"nodeLocalLoadBalancing,omitempty"`
}

func (f *K0sNetwork) Annotate(a infer.Annotator) {
	podCIDR := defaultNetworkPodCIDR
	a.Describe(&f.PodCIDR, "The CIDR from which Pod IPs are allocated.")
	a.SetDefault(&f.PodCIDR, &podCIDR)

	serviceCIDR := defaultNetworkServiceCIDR
	a.Describe(&f.ServiceCIDR, "The CIDR from which Service IPs are allocated.")
	a.SetDefault(&f.ServiceCIDR, &serviceCIDR)

	provider := defaultNetworkProvider
	a.Describe(&f.Provider, "The network provider to use.")
	a.SetDefault(&f.Provider, &provider)

	a.Describe(&f.DualStack, "K0s dual-stack configuration.")
	//a.SetDefault(&f.DualStack, &K0sDualStack{})

	a.Describe(&f.Calico, "K0s calico configuration.")
	//a.SetDefault(&f.Calico, &K0sCalico{})

	a.Describe(&f.KubeRouter, "K0s kube-router configuration.")
	//a.SetDefault(&f.KubeRouter, &K0sKubeRouter{})

	a.Describe(&f.KubeProxy, "K0s kube-proxy configuration.")
	//a.SetDefault(&f.KubeProxy, &K0sKubeProxy{})

	a.Describe(&f.NodeLocalLoadBalancing, "K0s node local load balancing configuration.")
	//a.SetDefault(&f.NodeLocalLoadBalancing, &K0sNodeLocalLoadBalancing{})
}

type K0sCalico struct {
	Mode                  *string           `pulumi:"mode,optional" json:"mode,omitempty"`
	Overlay               *string           `pulumi:"overlay,optional" json:"overlay,omitempty"`
	VXLANPort             *int              `pulumi:"vxlanPort,optional" json:"vxlanPort,omitempty"`
	VXLANVNI              *int              `pulumi:"vxlanVNI,optional" json:"vxlanVNI,omitempty"`
	MTU                   *int              `pulumi:"mtu,optional" json:"mtu,omitempty"`
	Wireguard             *bool             `pulumi:"wireguard,optional" json:"wireguard,omitempty"`
	FlexVolumeDriverPath  *string           `pulumi:"flexVolumeDriverPath,optional" json:"flexVolumeDriverPath,omitempty"`
	IPAutodetectionMethod *string           `pulumi:"ipAutodetectionMethod,optional" json:"ipAutodetectionMethod,omitempty"`
	EnvVars               map[string]string `pulumi:"envVars,optional" json:"envVars,omitempty"`
}

type K0sDualStack struct {
	Enabled         *bool   `pulumi:"enabled,optional" json:"enabled,omitempty"`
	IPv6PodCIDR     *string `pulumi:"IPv6podCIDR,optional" json:"IPv6podCIDR,omitempty"`
	IPv6ServiceCIDR *string `pulumi:"IPv6serviceCIDR,optional" json:"IPv6serviceCIDR,omitempty"`
}

type K0sKubeRouter struct {
	AutoMTU     *bool             `pulumi:"autoMTU,optional" json:"autoMTU,omitempty"`
	MTU         *int              `pulumi:"mtu,optional" json:"mtu,omitempty"`
	MetricsPort *int              `pulumi:"metricsPort,optional" json:"metricsPort,omitempty"`
	Hairpin     *string           `pulumi:"hairpin,optional" json:"hairpin,omitempty"`
	IPMasq      *bool             `pulumi:"ipMasq,optional" json:"ipMasq,omitempty"`
	ExtraArgs   map[string]string `pulumi:"extraArgs,optional" json:"extraArgs,omitempty"`
}

func (f *K0sKubeRouter) Annotate(a infer.Annotator) {
	mtu := defaultKubeRouterMTU
	a.Describe(&f.MTU, "The MTU to use for the kube-router network interfaces.")
	a.SetDefault(&f.MTU, &mtu)

	autoMTU := defaultKubeRouterAutoMTU
	a.Describe(&f.AutoMTU, "Automatically detect and set the MTU based on the host interface.")
	a.SetDefault(&f.AutoMTU, &autoMTU)
}

type K0sKubeProxy struct {
	Disabled          *bool                 `pulumi:"disabled,optional" json:"disabled,omitempty"`
	Mode              *string               `pulumi:"mode,optional" json:"mode,omitempty"`
	IPTables          *K0sKubeProxyIPTables `pulumi:"iptables,optional" json:"iptables,omitempty"`
	IPVS              *K0sKubeProxyIPVS     `pulumi:"ipvs,optional" json:"ipvs,omitempty"`
	NodePortAddresses *string               `pulumi:"nodePortAddresses,optional" json:"nodePortAddresses,omitempty"`
}

func (f *K0sKubeProxy) Annotate(a infer.Annotator) {
	disabled := defaultKubeProxyDisabled
	a.Describe(&f.Disabled, "Disable kube-proxy.")
	a.SetDefault(&f.Disabled, &disabled)

	mode := defaultKubeProxyMode
	a.Describe(&f.Mode, "The kube-proxy mode to use. One of 'iptables' or 'ipvs'.")
	a.SetDefault(&f.Mode, &mode)
}

type K0sKubeProxyIPTables struct {
	MasqueradeAll *bool   `pulumi:"masqueradeAll,optional" json:"masqueradeAll,omitempty"`
	MasqueradeBit *int    `pulumi:"masqueradeBit,optional" json:"masqueradeBit,omitempty"`
	MinSyncPeriod *string `pulumi:"minSyncPeriod,optional" json:"minSyncPeriod,omitempty"`
	SyncPeriod    *string `pulumi:"syncPeriod,optional" json:"syncPeriod,omitempty"`
}

type K0sKubeProxyIPVS struct {
	ExcludeCIDRs  *string `pulumi:"excludeCIDRs,optional" json:"excludeCIDRs,omitempty"`
	MinSyncPeriod *string `pulumi:"minSyncPeriod,optional" json:"minSyncPeriod,omitempty"`
	Scheduler     *string `pulumi:"scheduler,optional" json:"scheduler,omitempty"`
	StrictARP     *bool   `pulumi:"strictARP,optional" json:"strictARP,omitempty"`
	SyncPeriod    *string `pulumi:"syncPeriod,optional" json:"syncPeriod,omitempty"`
	TCPFinTimeout *string `pulumi:"tcpFinTimeout,optional" json:"tcpFinTimeout,omitempty"`
	TCPTimeout    *string `pulumi:"tcpTimeout,optional" json:"tcpTimeout,omitempty"`
	UDPTimeout    *string `pulumi:"udpTimeout,optional" json:"udpTimeout,omitempty"`
}

type K0sNodeLocalLoadBalancing struct {
	Enabled    *bool          `pulumi:"enabled,optional" json:"enabled,omitempty"`
	Type       *string        `pulumi:"type,optional" json:"type,omitempty"`
	EnvoyProxy *K0sEnvoyProxy `pulumi:"envoyProxy,optional" json:"envoyProxy,omitempty"`
}

type K0sEnvoyProxy struct {
	Image                      *string `pulumi:"image,optional" json:"image,omitempty"`
	ImagePullPolicy            *string `pulumi:"imagePullPolicy,optional" json:"imagePullPolicy,omitempty"`
	APIServerBindPort          *int    `pulumi:"apiServerBindPort,optional" json:"apiServerBindPort,omitempty"`
	KonnectivityServerBindPort *int    `pulumi:"konnectivityServerBindPort,optional" json:"konnectivityServerBindPort,omitempty"`
}

type K0sPodSecurityPolicy struct {
	DefaultPolicy *string `pulumi:"defaultPolicy,optional" json:"defaultPolicy,omitempty"`
}

func (f *K0sPodSecurityPolicy) Annotate(a infer.Annotator) {
	defaultPolicy := defaultPodSecurityPolicy
	a.Describe(&f.DefaultPolicy, "The default Pod Security Policy to use.")
	a.SetDefault(&f.DefaultPolicy, &defaultPolicy)
}

type K0sControllerManager struct {
	ExtraArgs map[string]string `pulumi:"extraArgs,optional" json:"extraArgs,omitempty"`
}

type K0sScheduler struct {
	ExtraArgs map[string]string `pulumi:"extraArgs,optional" json:"extraArgs,omitempty"`
}

type K0sStorage struct {
	Type *string  `pulumi:"type,optional" json:"type,omitempty"`
	Etcd *K0sEtcd `pulumi:"etcd,optional" json:"etcd,omitempty"`
	Kine *K0sKine `pulumi:"kine,optional" json:"kine,omitempty"`
}

func (f *K0sStorage) Annotate(a infer.Annotator) {
	storageType := defaultStorageType
	a.Describe(&f.Type, "The storage type to use. One of 'etcd' or 'kine'.")
	a.SetDefault(&f.Type, &storageType)
}

type K0sEtcd struct {
	PeerAddress     *string                 `pulumi:"peerAddress,optional" json:"peerAddress,omitempty"`
	ExtraArgs       map[string]string       `pulumi:"extraArgs,optional" json:"extraArgs,omitempty"`
	ExternalCluster *K0sEtcdExternalCluster `pulumi:"externalCluster,optional" json:"externalCluster,omitempty"`
}

type K0sEtcdExternalCluster struct {
	Endpoints  []string `pulumi:"endpoints" json:"endpoints"`
	EtcdPrefix *string  `pulumi:"etcdPrefix,optional" json:"etcdPrefix,omitempty"`
	CA         *string  `pulumi:"ca,optional" provider:"secret" json:"ca,omitempty"`
	ClientCert *string  `pulumi:"clientCert,optional" provider:"secret" json:"clientCert,omitempty"`
	ClientKey  *string  `pulumi:"clientKey,optional" provider:"secret" json:"clientKey,omitempty"`
}

type K0sKine struct {
	DataSource string `pulumi:"dataSource" json:"dataSource,omitempty" provider:"secret"`
}

type K0sTelemetry struct {
	Enabled *bool `pulumi:"enabled,optional" json:"enabled,omitempty"`
}

func (f *K0sTelemetry) Annotate(a infer.Annotator) {
	enabled := defaultTelemetryEnabled
	a.Describe(&f.Enabled, "Enable or disable telemetry.")
	a.SetDefault(&f.Enabled, &enabled)
}

type K0sWorkerProfile struct {
	Name   string            `pulumi:"name" json:"name"`
	Values map[string]string `pulumi:"values" json:"values"`
}

type K0sFeatureGate struct {
	Enabled    *bool    `pulumi:"enabled,optional" json:"enabled,omitempty"`
	Name       string   `pulumi:"name" json:"name"`
	Components []string `pulumi:"components,optional" json:"components,omitempty"`
}

func (args *ClusterArgs) FillDefaults(name *string) error {
	if args.APIVersion == nil {
		apiVersion := defaultAPIVersion
		args.APIVersion = &apiVersion
	}

	if args.Kind == nil {
		kind := defaultKind
		args.Kind = &kind
	}

	if args.Metadata == nil {
		args.Metadata = &ClusterMetadata{}
	}

	if args.Metadata.Name == nil {
		args.Metadata.Name = name
	}

	if args.Metadata.Name == nil {
		clusterName := defaultClusterName
		args.Metadata.Name = &clusterName
	}

	if args.Spec == nil {
		args.Spec = &ClusterSpec{}
	}

	if args.Spec.Hosts == nil {
		args.Spec.Hosts = []*ClusterHost{}
	}

	if args.Spec.K0s == nil {
		args.Spec.K0s = &ClusterK0s{}
	}

	if args.Spec.K0s.Version == nil {
		k0sVersion, err := version.LatestStable()
		if err != nil {
			return err
		}

		k0sVersionStr := k0sVersion.String()
		args.Spec.K0s.Version = &k0sVersionStr
	}

	if args.Spec.K0s.VersionChannel == nil {
		versionChannel := defaultVersionChannel
		args.Spec.K0s.VersionChannel = &versionChannel
	}

	if args.Spec.K0s.DynamicConfig == nil {
		dynamicConfig := defaultDynamicConfig
		args.Spec.K0s.DynamicConfig = &dynamicConfig
	}

	if args.Spec.K0s.Config == nil {
		args.Spec.K0s.Config = &K0s{}
	}

	if args.Spec.K0s.Config.APIVersion == nil {
		apiVersion := defaultK0sAPIVersion
		args.Spec.K0s.Config.APIVersion = &apiVersion
	}

	if args.Spec.K0s.Config.Kind == nil {
		kind := defaultK0sKind
		args.Spec.K0s.Config.Kind = &kind
	}

	if args.Spec.K0s.Config.Metadata == nil {
		args.Spec.K0s.Config.Metadata = &K0sMetadata{}
	}

	if args.Spec.K0s.Config.Metadata.Name == nil {
		args.Spec.K0s.Config.Metadata.Name = name
	}

	if args.Spec.K0s.Config.Metadata.Name == nil {
		k0sName := defaultK0sName
		args.Spec.K0s.Config.Metadata.Name = &k0sName
	}

	if args.Spec.K0s.Config.Spec == nil {
		args.Spec.K0s.Config.Spec = &K0sSpec{}
	}

	if args.Spec.K0s.Config.Spec.API == nil {
		args.Spec.K0s.Config.Spec.API = &K0sAPI{}
	}

	if args.Spec.K0s.Config.Spec.API.Port == nil {
		port := defaultAPIPort
		args.Spec.K0s.Config.Spec.API.Port = &port
	}

	if args.Spec.K0s.Config.Spec.API.K0sApiPort == nil {
		k0sApiPort := defaultAPIK0sAPIPort
		args.Spec.K0s.Config.Spec.API.K0sApiPort = &k0sApiPort
	}

	if args.Spec.K0s.Config.Spec.InstallConfig == nil {
		args.Spec.K0s.Config.Spec.InstallConfig = &K0sInstallConfig{}
	}

	if args.Spec.K0s.Config.Spec.InstallConfig.Users == nil {
		args.Spec.K0s.Config.Spec.InstallConfig.Users = &K0sInstallConfigUser{}
	}

	if args.Spec.K0s.Config.Spec.InstallConfig.Users.EtcdUser == nil {
		etcdUser := defaultEtcdUser
		args.Spec.K0s.Config.Spec.InstallConfig.Users.EtcdUser = &etcdUser
	}

	if args.Spec.K0s.Config.Spec.InstallConfig.Users.KineUser == nil {
		kineUser := defaultKineUser
		args.Spec.K0s.Config.Spec.InstallConfig.Users.KineUser = &kineUser
	}

	if args.Spec.K0s.Config.Spec.InstallConfig.Users.KonnectivityUser == nil {
		konnectivityUser := defaultKonnectivityUser
		args.Spec.K0s.Config.Spec.InstallConfig.Users.KonnectivityUser = &konnectivityUser
	}

	if args.Spec.K0s.Config.Spec.InstallConfig.Users.KubeAPIServerUser == nil {
		kubeAPIserverUser := defaultKubeAPIserverUser
		args.Spec.K0s.Config.Spec.InstallConfig.Users.KubeAPIServerUser = &kubeAPIserverUser
	}

	if args.Spec.K0s.Config.Spec.InstallConfig.Users.KubeSchedulerUser == nil {
		kubeSchedulerUser := defaultKubeSchedulerUser
		args.Spec.K0s.Config.Spec.InstallConfig.Users.KubeSchedulerUser = &kubeSchedulerUser
	}

	if args.Spec.K0s.Config.Spec.Konnectivity == nil {
		args.Spec.K0s.Config.Spec.Konnectivity = &K0sKonnectivity{}
	}

	if args.Spec.K0s.Config.Spec.Konnectivity.AgentPort == nil {
		agentPort := defaultKonnectivityAgentPort
		args.Spec.K0s.Config.Spec.Konnectivity.AgentPort = &agentPort
	}

	if args.Spec.K0s.Config.Spec.Konnectivity.AdminPort == nil {
		adminPort := defaultKonnectivityAdminPort
		args.Spec.K0s.Config.Spec.Konnectivity.AdminPort = &adminPort
	}

	if args.Spec.K0s.Config.Spec.Network == nil {
		args.Spec.K0s.Config.Spec.Network = &K0sNetwork{}
	}

	if args.Spec.K0s.Config.Spec.Network.PodCIDR == nil {
		podCIDR := defaultNetworkPodCIDR
		args.Spec.K0s.Config.Spec.Network.PodCIDR = &podCIDR
	}

	if args.Spec.K0s.Config.Spec.Network.ServiceCIDR == nil {
		serviceCIDR := defaultNetworkServiceCIDR
		args.Spec.K0s.Config.Spec.Network.ServiceCIDR = &serviceCIDR
	}

	if args.Spec.K0s.Config.Spec.Network.Provider == nil {
		provider := defaultNetworkProvider
		args.Spec.K0s.Config.Spec.Network.Provider = &provider
	}

	if args.Spec.K0s.Config.Spec.Network.KubeRouter == nil {
		args.Spec.K0s.Config.Spec.Network.KubeRouter = &K0sKubeRouter{}
	}

	if args.Spec.K0s.Config.Spec.Network.KubeRouter.MTU == nil {
		mtu := defaultKubeRouterMTU
		args.Spec.K0s.Config.Spec.Network.KubeRouter.MTU = &mtu
	}

	if args.Spec.K0s.Config.Spec.Network.KubeRouter.AutoMTU == nil {
		autoMTU := defaultKubeRouterAutoMTU
		args.Spec.K0s.Config.Spec.Network.KubeRouter.AutoMTU = &autoMTU
	}

	if args.Spec.K0s.Config.Spec.Network.KubeProxy == nil {
		args.Spec.K0s.Config.Spec.Network.KubeProxy = &K0sKubeProxy{}
	}

	if args.Spec.K0s.Config.Spec.Network.KubeProxy.Disabled == nil {
		disabled := defaultKubeProxyDisabled
		args.Spec.K0s.Config.Spec.Network.KubeProxy.Disabled = &disabled
	}

	if args.Spec.K0s.Config.Spec.Network.KubeProxy.Mode == nil {
		mode := defaultKubeProxyMode
		args.Spec.K0s.Config.Spec.Network.KubeProxy.Mode = &mode
	}

	if args.Spec.K0s.Config.Spec.PodSecurityPolicy == nil {
		args.Spec.K0s.Config.Spec.PodSecurityPolicy = &K0sPodSecurityPolicy{}
	}

	if args.Spec.K0s.Config.Spec.PodSecurityPolicy.DefaultPolicy == nil {
		defaultPolicy := defaultPodSecurityPolicy
		args.Spec.K0s.Config.Spec.PodSecurityPolicy.DefaultPolicy = &defaultPolicy
	}

	if args.Spec.K0s.Config.Spec.Storage == nil {
		args.Spec.K0s.Config.Spec.Storage = &K0sStorage{}
	}

	if args.Spec.K0s.Config.Spec.Storage.Type == nil {
		storageType := defaultStorageType
		args.Spec.K0s.Config.Spec.Storage.Type = &storageType
	}

	if args.Spec.K0s.Config.Spec.Telemetry == nil {
		args.Spec.K0s.Config.Spec.Telemetry = &K0sTelemetry{}
	}

	if args.Spec.K0s.Config.Spec.Telemetry.Enabled == nil {
		enabled := defaultTelemetryEnabled
		args.Spec.K0s.Config.Spec.Telemetry.Enabled = &enabled
	}

	return nil
}

func (args *ClusterArgs) APIAddress() string {
	address := "localhost"
	port := 6443

	clt, cleanup, err := args.k0sctl()

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

func (args *ClusterArgs) k0sctl() (*v1beta1.Cluster, func(), error) {
	prefix := fmt.Sprintf("%s-%d", *args.Metadata.Name, rand.Intn(int(^uint(0)>>1)))

	cleanup := func() {
		_ = cleanupTempFiles(prefix)
	}

	bytes, err := yamlJSON.Marshal(args)
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
		m["name"] = args.Metadata.Name
	}

	if clt.Spec.K0s.Config["spec"] == nil {
		clt.Spec.K0s.Config["spec"] = dig.Mapping{}
	}

	// replace inline values with file paths
	if err := args.k0sctlConvertHostsPaths(prefix, clt.Spec.Hosts); err != nil {
		return nil, cleanup, err
	}

	ec := clt.Spec.K0s.Config.Dig("spec", "storage", "etcd", "externalCluster")
	if ec != nil {
		if ec, ok := ec.(*k0sapi.ExternalCluster); ok {
			if err := args.k0sctlConvertExternalEtcdPaths(prefix, ec); err != nil {
				return nil, cleanup, err
			}
		}
	}

	return &clt, cleanup, nil
}

func (args *ClusterArgs) k0sctlConvertHostsPaths(prefix string, hosts cluster.Hosts) error {
	for i, host := range hosts {
		clusterHost := args.Spec.Hosts[i]

		if host.SSH != nil {
			if err := args.k0sctlConvertSSHPaths(prefix, clusterHost.SSH, host.SSH); err != nil {
				return err
			}
		}

		if host.OpenSSH != nil {
			if err := args.k0sctlConvertOpenSSHPaths(prefix, clusterHost.OpenSSH, host.OpenSSH); err != nil {
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

			if err := args.k0sctlConvertBastionPaths(prefix, clusterHost.WinRM.Bastion, host.WinRM.Bastion); err != nil {
				return err
			}
		}
	}

	return nil
}

func (args *ClusterArgs) k0sctlConvertSSHPaths(prefix string, ssh *ClusterSSH, rigSSH *ssh.Config) error {
	if ssh != nil && ssh.Key != nil {
		filename, err := contentToTempFile(prefix, *ssh.Key, true)
		if err != nil {
			return err
		}

		rigSSH.KeyPath = &filename
	}

	return args.k0sctlConvertBastionPaths(prefix, ssh.Bastion, rigSSH.Bastion)
}

func (args *ClusterArgs) k0sctlConvertBastionPaths(prefix string, ssh *ClusterBastion, rigSSH *ssh.Config) error {
	if ssh != nil && ssh.Key != nil {
		filename, err := contentToTempFile(prefix, *ssh.Key, true)
		if err != nil {
			return err
		}

		rigSSH.KeyPath = &filename
	}

	return nil
}

func (args *ClusterArgs) k0sctlConvertOpenSSHPaths(prefix string, ssh *ClusterOpenSSH, rigSSH *openssh.Config) error {
	if ssh != nil && ssh.Key != nil {
		filename, err := contentToTempFile(prefix, *ssh.Key, true)
		if err != nil {
			return err
		}

		rigSSH.KeyPath = &filename
	}

	return nil
}

func (args *ClusterArgs) k0sctlConvertExternalEtcdPaths(
	prefix string,
	k0sctlExternalCluster *k0sapi.ExternalCluster,
) error {
	externalCluster := args.Spec.K0s.Config.Spec.Storage.Etcd.ExternalCluster

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
