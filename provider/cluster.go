package provider

import (
	"strings"

	"github.com/TwiN/deepmerge"
	"github.com/k0sproject/dig"
	"github.com/k0sproject/rig"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/ydkn/pulumi-k0s/provider/internal/introspect"
	"github.com/ydkn/pulumi-k0s/provider/internal/k0sctl"
	"gopkg.in/yaml.v2"
)

type ClusterArgs struct {
	APIVersion *string          `pulumi:"apiVersion,optional" yaml:"apiVersion,omitempty"`
	Kind       *string          `pulumi:"kind,optional" yaml:"kind,omitempty"`
	Metadata   *ClusterMetadata `pulumi:"metadata,optional" yaml:"metadata,omitempty"`
	Spec       *ClusterSpec     `pulumi:"spec,optional" yaml:"spec,omitempty"`
}

type ClusterMetadata struct {
	Name *string `pulumi:"name" yaml:"name"`
}

type ClusterSpec struct {
	Hosts []ClusterHost `pulumi:"hosts" yaml:"hosts"`
	K0s   *ClusterK0s   `pulumi:"k0s,optional" yaml:"k0s,omitempty"`
}

type ClusterHost struct {
	Role             *string           `pulumi:"role" yaml:"role"`
	PrivateInterface *string           `pulumi:"privateInterface,optional" yaml:"privateInterface,omitempty"`
	PrivateAddress   *string           `pulumi:"privateAddress,optional" yaml:"privateAddress,omitempty"`
	Environment      map[string]string `pulumi:"environment,optional" yaml:"environment,omitempty"`
	UploadBinary     *bool             `pulumi:"uploadBinary,optional" yaml:"uploadBinary,omitempty"`
	K0sBinaryPath    *string           `pulumi:"k0sBinaryPath,optional" yaml:"k0sBinaryPath,omitempty"`
	InstallFlags     []string          `pulumi:"installFlags,optional" yaml:"installFlags,omitempty"`
	Files            []ClusterFile     `pulumi:"files,optional" yaml:"files,omitempty"`
	OS               *string           `pulumi:"os,optional" yaml:"os,omitempty"`
	Hostname         *string           `pulumi:"hostname,optional" yaml:"hostname,omitempty"`
	Hooks            *ClusterHooks     `pulumi:"hooks,optional" yaml:"hooks,omitempty"`
	WinRM            *ClusterWinRM     `pulumi:"winRM,optional" yaml:"winRM,omitempty"`
	SSH              *ClusterSSH       `pulumi:"ssh,optional" yaml:"ssh,omitempty"`
	Localhost        *ClusterLocalhost `pulumi:"localhost,optional" yaml:"localhost,omitempty"`
	NoTaints         *bool             `pulumi:"noTaints,optional" yaml:"noTaints,omitempty"`
}

type ClusterFile struct {
	Name                 *string `pulumi:"name,optional" yaml:"name,omitempty"`
	Source               *string `pulumi:"src,optional" yaml:"src,omitempty"`
	DestinationDirectory *string `pulumi:"dstDir,optional" yaml:"dstDir,omitempty"`
	Destination          *string `pulumi:"dst,optional" yaml:"dst,omitempty"`
	Permissions          *string `pulumi:"perm,optional" yaml:"perm,omitempty"`
	DirectoryPermissions *string `pulumi:"dirPerm,optional" yaml:"dirPerm,omitempty"`
	User                 *string `pulumi:"user,optional" yaml:"user,omitempty"`
	Group                *string `pulumi:"group,optional" yaml:"group,omitempty"`
}

type ClusterHooks struct {
	Apply  *ClusterHook `pulumi:"apply,optional" yaml:"apply,omitempty"`
	Backup *ClusterHook `pulumi:"backup,optional" yaml:"backup,omitempty"`
	Reset  *ClusterHook `pulumi:"reset,optional" yaml:"reset,omitempty"`
}

type ClusterHook struct {
	Before []string `pulumi:"before,optional" yaml:"before,omitempty"`
	After  []string `pulumi:"after,optional" yaml:"after,omitempty"`
}

type ClusterWinRM struct {
	Address       *string     `pulumi:"address" yaml:"address"`
	Port          *int        `pulumi:"port,optional" yaml:"port,omitempty"`
	User          *string     `pulumi:"user,optional" yaml:"user,omitempty"`
	Password      *string     `pulumi:"password,optional" yaml:"password,omitempty" provider:"secret"`
	UseHTTPS      *bool       `pulumi:"useHTTPS,optional" yaml:"useHTTPS,omitempty"`
	Insecure      *bool       `pulumi:"insecure,optional" yaml:"insecure,omitempty"`
	UseNTLM       *bool       `pulumi:"useNTLM,optional" yaml:"useNTLM,omitempty"`
	CaCert        *string     `pulumi:"caCert,optional" yaml:"caCert,omitempty" provider:"secret"`
	Cert          *string     `pulumi:"cert,optional" yaml:"cert,omitempty" provider:"secret"`
	Key           *string     `pulumi:"key,optional" yaml:"key,omitempty" provider:"secret"`
	TLSServerName *string     `pulumi:"tlsServerName,optional" yaml:"tlsServerName,omitempty"`
	Bastion       *ClusterSSH `pulumi:"bastion,optional" yaml:"bastion,omitempty"`
}

type ClusterSSH struct {
	Address *string     `pulumi:"address" yaml:"address,omitempty"`
	Port    *int        `pulumi:"port,optional" yaml:"port,omitempty"`
	User    *string     `pulumi:"user,optional" yaml:"user,omitempty"`
	Key     *string     `pulumi:"key,optional" yaml:"key,omitempty" provider:"secret"`
	HostKey *string     `pulumi:"hostKey,optional" yaml:"hostKey,omitempty"`
	Bastion *ClusterSSH `pulumi:"bastion,optional" yaml:"bastion,omitempty"`
}

type ClusterLocalhost struct {
	Enabled *bool `pulumi:"enabled,optional" yaml:"enabled,omitempty"`
}

type ClusterK0s struct {
	Version        *string           `pulumi:"version,optional" yaml:"version,omitempty"`
	VersionChannel *string           `pulumi:"versionChannel,optional" yaml:"versionChannel,omitempty"`
	DynamicConfig  *bool             `pulumi:"dynamicConfig,optional" yaml:"dynamicConfig,omitempty"`
	Config         *ClusterK0sConfig `pulumi:"config,optional" yaml:"config,omitempty"`
}

type ClusterK0sConfig struct {
	Metadata *ClusterMetadata `pulumi:"metadata,optional" yaml:"metadata,omitempty"`
	Spec     *ClusterK0sSpec  `pulumi:"spec,optional" yaml:"spec,omitempty"`
}

type ClusterK0sSpec struct {
	API               *ClusterAPI               `pulumi:"api,optional" yaml:"api,omitempty"`
	Images            *ClusterImages            `pulumi:"images,optional" yaml:"images,omitempty"`
	InstallConfig     *ClusterInstallConfig     `pulumi:"installConfig,optional" yaml:"installConfig,omitempty"`
	Konnectivity      *ClusterKonnectivity      `pulumi:"konnectivity,optional" yaml:"konnectivity,omitempty"`
	Network           *ClusterNetwork           `pulumi:"network,optional" yaml:"network,omitempty"`
	PodSecurityPolicy *ClusterPodSecurityPolicy `pulumi:"podSecurityPolicy,optional" yaml:"podSecurityPolicy,omitempty"`
	ControllerManager *ClusterControllerManager `pulumi:"controllerManager,optional" yaml:"controllerManager,omitempty"`
	Scheduler         *ClusterScheduler         `pulumi:"scheduler,optional" yaml:"scheduler,omitempty"`
	Storage           *ClusterStorage           `pulumi:"storage,optional" yaml:"storage,omitempty"`
	WorkerProfiles    []ClusterWorkerProfile    `pulumi:"workerProfiles,optional" yaml:"workerProfiles,omitempty"`
	FeatureGates      []ClusterFeatureGate      `pulumi:"featureGates,optional" yaml:"featureGates,omitempty"`
	Telemetry         *ClusterTelemetry         `pulumi:"telemetry,optional" yaml:"telemetry,omitempty"`
}

type ClusterAPI struct {
	Address         *string           `pulumi:"address,optional" yaml:"address,omitempty"`
	Port            *int              `pulumi:"port,optional" yaml:"port,omitempty"`
	K0sApiPort      *int              `pulumi:"k0sApiPort,optional" yaml:"k0sApiPort,omitempty"`
	ExternalAddress *string           `pulumi:"externalAddress,optional" yaml:"externalAddress,omitempty"`
	SANs            []string          `pulumi:"sans,optional" yaml:"sans,omitempty"`
	ExtraArgs       map[string]string `pulumi:"extraArgs,optional" yaml:"extraArgs,omitempty"`
}

type ClusterImages struct {
	DefaultPullPolicy *string                 `pulumi:"default_pull_policy,optional" yaml:"default_pull_policy,omitempty"`
	Repository        *string                 `pulumi:"repository,optional" yaml:"repository,omitempty"`
	Konnectivity      *ClusterImage           `pulumi:"konnectivity,optional" yaml:"konnectivity,omitempty"`
	MetricsServer     *ClusterImage           `pulumi:"metricsserver,optional" yaml:"metricsserver,omitempty"`
	Kubeproxy         *ClusterImage           `pulumi:"kubeproxy,optional" yaml:"kubeproxy,omitempty"`
	CoreDNS           *ClusterImage           `pulumi:"coredns,optional" yaml:"coredns,omitempty"`
	Pause             *ClusterImage           `pulumi:"pause,optional" yaml:"pause,omitempty"`
	Calico            *ClusterCalicoImage     `pulumi:"calico,optional" yaml:"calico,omitempty"`
	KubeRouter        *ClusterKubeRouterImage `pulumi:"kuberouter,optional" yaml:"kuberouter,omitempty"`
}

type ClusterImage struct {
	Image   *string `pulumi:"image,optional" yaml:"image,omitempty"`
	Version *string `pulumi:"version,optional" yaml:"version,omitempty"`
}

type ClusterCalicoImage struct {
	CNI             *ClusterImage `pulumi:"cni,optional" yaml:"cni,omitempty"`
	FlexVolume      *ClusterImage `pulumi:"flexvolume,optional" yaml:"flexvolume,omitempty"`
	Node            *ClusterImage `pulumi:"node,optional" yaml:"node,omitempty"`
	KubeControllers *ClusterImage `pulumi:"kubecontrollers,optional" yaml:"kubecontrollers,omitempty"`
}

type ClusterKubeRouterImage struct {
	CNI          *ClusterImage `pulumi:"cni,optional" yaml:"cni,omitempty"`
	CNIInstaller *ClusterImage `pulumi:"cniInstaller,optional" yaml:"cniInstaller,omitempty"`
}

type ClusterInstallConfig struct {
	Users *ClusterInstallConfigUser `pulumi:"users,optional" yaml:"users,omitempty"`
}

type ClusterInstallConfigUser struct {
	EtcdUser          *string `pulumi:"etcdUser,optional" yaml:"etcdUser,omitempty"`
	KineUser          *string `pulumi:"kineUser,optional" yaml:"kineUser,omitempty"`
	KonnectivityUser  *string `pulumi:"konnectivityUser,optional" yaml:"konnectivityUser,omitempty"`
	KubeAPIServerUser *string `pulumi:"kubeAPIserverUser,optional" yaml:"kubeAPIserverUser,omitempty"`
	KubeSchedulerUser *string `pulumi:"kubeSchedulerUser,optional" yaml:"kubeSchedulerUser,omitempty"`
}

type ClusterKonnectivity struct {
	AdminPort *int `pulumi:"adminPort,optional" yaml:"adminPort,omitempty"`
	AgentPort *int `pulumi:"agentPort,optional" yaml:"agentPort,omitempty"`
}

type ClusterNetwork struct {
	Provider               *string                        `pulumi:"provider,optional" yaml:"provider,omitempty"`
	PodCIDR                *string                        `pulumi:"podCIDR,optional" yaml:"podCIDR,omitempty"`
	ServiceCIDR            *string                        `pulumi:"serviceCIDR,optional" yaml:"serviceCIDR,omitempty"`
	ClusterDomain          *string                        `pulumi:"clusterDomain,optional" yaml:"clusterDomain,omitempty"`
	DualStack              *ClusterDualStack              `pulumi:"dualStack,optional" yaml:"dualStack,omitempty"`
	Calico                 *ClusterCalico                 `pulumi:"calico,optional" yaml:"calico,omitempty"`
	KubeRouter             *ClusterKubeRouter             `pulumi:"kuberouter,optional" yaml:"kuberouter,omitempty"`
	KubeProxy              *ClusterKubeProxy              `pulumi:"kubeProxy,optional" yaml:"kubeProxy,omitempty"`
	NodeLocalLoadBalancing *ClusterNodeLocalLoadBalancing `pulumi:"nodeLocalLoadBalancing,optional" yaml:"nodeLocalLoadBalancing,omitempty"`
}

type ClusterCalico struct {
	Mode                  *string           `pulumi:"mode,optional" yaml:"mode,omitempty"`
	Overlay               *string           `pulumi:"overlay,optional" yaml:"overlay,omitempty"`
	VXLANPort             *int              `pulumi:"vxlanPort,optional" yaml:"vxlanPort,omitempty"`
	VXLANVNI              *int              `pulumi:"vxlanVNI,optional" yaml:"vxlanVNI,omitempty"`
	MTU                   *int              `pulumi:"mtu,optional" yaml:"mtu,omitempty"`
	Wireguard             *bool             `pulumi:"wireguard,optional" yaml:"wireguard,omitempty"`
	FlexVolumeDriverPath  *string           `pulumi:"flexVolumeDriverPath,optional" yaml:"flexVolumeDriverPath,omitempty"`
	IPAutodetectionMethod *string           `pulumi:"ipAutodetectionMethod,optional" yaml:"ipAutodetectionMethod,omitempty"`
	EnvVars               map[string]string `pulumi:"envVars,optional" yaml:"envVars,omitempty"`
}

type ClusterDualStack struct {
	Enabled         *bool   `pulumi:"enabled,optional" yaml:"enabled,omitempty"`
	IPv6PodCIDR     *string `pulumi:"IPv6podCIDR,optional" yaml:"IPv6podCIDR,omitempty"`
	IPv6ServiceCIDR *string `pulumi:"IPv6serviceCIDR,optional" yaml:"IPv6serviceCIDR,omitempty"`
}

type ClusterKubeRouter struct {
	AutoMTU     *bool             `pulumi:"autoMTU,optional" yaml:"autoMTU,omitempty"`
	MTU         *int              `pulumi:"mtu,optional" yaml:"mtu,omitempty"`
	MetricsPort *int              `pulumi:"metricsPort,optional" yaml:"metricsPort,omitempty"`
	Hairpin     *string           `pulumi:"hairpin,optional" yaml:"hairpin,omitempty"`
	IPMasq      *bool             `pulumi:"ipMasq,optional" yaml:"ipMasq,omitempty"`
	ExtraArgs   map[string]string `pulumi:"extraArgs,optional" yaml:"extraArgs,omitempty"`
}

type ClusterKubeProxy struct {
	Disabled          *bool                     `pulumi:"disabled,optional" yaml:"disabled,omitempty"`
	Mode              *string                   `pulumi:"mode,optional" yaml:"mode,omitempty"`
	IPTables          *ClusterKubeProxyIPTables `pulumi:"iptables,optional" yaml:"iptables,omitempty"`
	IPVS              *ClusterKubeProxyIPVS     `pulumi:"ipvs,optional" yaml:"ipvs,omitempty"`
	NodePortAddresses *string                   `pulumi:"nodePortAddresses,optional" yaml:"nodePortAddresses,omitempty"`
}

type ClusterKubeProxyIPTables struct {
	MasqueradeAll *bool   `pulumi:"masqueradeAll,optional" yaml:"masqueradeAll,omitempty"`
	MasqueradeBit *int    `pulumi:"masqueradeBit,optional" yaml:"masqueradeBit,omitempty"`
	MinSyncPeriod *string `pulumi:"minSyncPeriod,optional" yaml:"minSyncPeriod,omitempty"`
	SyncPeriod    *string `pulumi:"syncPeriod,optional" yaml:"syncPeriod,omitempty"`
}

type ClusterKubeProxyIPVS struct {
	ExcludeCIDRs  *string `pulumi:"excludeCIDRs,optional" yaml:"excludeCIDRs,omitempty"`
	MinSyncPeriod *string `pulumi:"minSyncPeriod,optional" yaml:"minSyncPeriod,omitempty"`
	Scheduler     *string `pulumi:"scheduler,optional" yaml:"scheduler,omitempty"`
	StrictARP     *bool   `pulumi:"strictARP,optional" yaml:"strictARP,omitempty"`
	SyncPeriod    *string `pulumi:"syncPeriod,optional" yaml:"syncPeriod,omitempty"`
	TCPFinTimeout *string `pulumi:"tcpFinTimeout,optional" yaml:"tcpFinTimeout,omitempty"`
	TCPTimeout    *string `pulumi:"tcpTimeout,optional" yaml:"tcpTimeout,omitempty"`
	UDPTimeout    *string `pulumi:"udpTimeout,optional" yaml:"udpTimeout,omitempty"`
}

type ClusterNodeLocalLoadBalancing struct {
	Enabled    *bool              `pulumi:"enabled,optional" yaml:"enabled,omitempty"`
	Type       *string            `pulumi:"type,optional" yaml:"type,omitempty"`
	EnvoyProxy *ClusterEnvoyProxy `pulumi:"envoyProxy,optional" yaml:"envoyProxy,omitempty"`
}

type ClusterEnvoyProxy struct {
	Image                      *string `pulumi:"image,optional" yaml:"image,omitempty"`
	ImagePullPolicy            *string `pulumi:"imagePullPolicy,optional" yaml:"imagePullPolicy,omitempty"`
	APIServerBindPort          *int    `pulumi:"apiServerBindPort,optional" yaml:"apiServerBindPort,omitempty"`
	KonnectivityServerBindPort *int    `pulumi:"konnectivityServerBindPort,optional" yaml:"konnectivityServerBindPort,omitempty"`
}

type ClusterPodSecurityPolicy struct {
	DefaultPolicy *string `pulumi:"defaultPolicy,optional" yaml:"defaultPolicy,omitempty"`
}

type ClusterControllerManager struct {
	ExtraArgs map[string]string `pulumi:"extraArgs,optional" yaml:"extraArgs,omitempty"`
}

type ClusterScheduler struct {
	ExtraArgs map[string]string `pulumi:"extraArgs,optional" yaml:"extraArgs,omitempty"`
}

type ClusterStorage struct {
	Type *string      `pulumi:"type,optional" yaml:"type,omitempty"`
	Etcd *ClusterEtcd `pulumi:"etcd,optional" yaml:"etcd,omitempty"`
	Kine *ClusterKine `pulumi:"kine,optional" yaml:"kine,omitempty"`
}

type ClusterEtcd struct {
	PeerAddress     *string                     `pulumi:"peerAddress,optional" yaml:"peerAddress,omitempty"`
	ExtraArgs       map[string]string           `pulumi:"extraArgs,optional" yaml:"extraArgs,omitempty"`
	ExternalCluster *ClusterEtcdExternalCluster `pulumi:"externalCluster,optional" yaml:"externalCluster,omitempty"`
}

type ClusterEtcdExternalCluster struct {
	Endpoints  []string `pulumi:"endpoints" yaml:"endpoints"`
	EtcdPrefix *string  `pulumi:"etcdPrefix,optional" yaml:"etcdPrefix,omitempty"`
	CA         *string  `pulumi:"ca,optional" yaml:"ca,omitempty" provider:"secret"`
	ClientCert *string  `pulumi:"clientCert,optional" yaml:"clientCert,omitempty" provider:"secret"`
	ClientKey  *string  `pulumi:"clientKey,optional" yaml:"clientKey,omitempty" provider:"secret"`
}

type ClusterKine struct {
	DataSource *string `pulumi:"dataSource" yaml:"dataSource,omitempty" provider:"secret"`
}

type ClusterTelemetry struct {
	Enabled *bool `pulumi:"enabled,optional" yaml:"enabled,omitempty"`
}

type ClusterWorkerProfile struct {
	Name   *string        `pulumi:"name" yaml:"name"`
	Values map[string]any `pulumi:"values" yaml:"values"`
}

type ClusterFeatureGate struct {
	Enabled    *bool    `pulumi:"enabled,optional" yaml:"enabled,omitempty"`
	Name       *string  `pulumi:"name" yaml:"name"`
	Components []string `pulumi:"components,optional" yaml:"components,omitempty"`
}

type Cluster struct{}

type ClusterState struct {
	ClusterArgs

	Kubeconfig string `pulumi:"kubeconfig" provider:"secret"`
}

func (c Cluster) Diff(ctx p.Context, name string, state ClusterState, args ClusterArgs) (p.DiffResponse, error) {
	defer func() {
		_ = cleanupTempFiles(name)
	}()

	diffResponse := p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          false,
		DetailedDiff:        map[string]p.PropertyDiff{},
	}

	_, cluster, err := clusterArgsToK0sCtlCluster(name, &args)
	if err != nil {
		return diffResponse, err
	}

	if args.Spec != nil && args.Spec.K0s != nil && args.Spec.K0s.Version == nil {
		version := cluster.Spec.K0s.Version.String()
		args.Spec.K0s.Version = &version
	}

	stateProps, err := introspect.NewPropertiesMap(state)
	if err != nil {
		return p.DiffResponse{}, err
	}

	argsProps, err := introspect.NewPropertiesMap(args)
	if err != nil {
		return p.DiffResponse{}, err
	}

	for key := range propertyMapDiff(stateProps, argsProps, []resource.PropertyKey{"kubeconfig"}) {
		diffResponse.DetailedDiff[string(key)] = p.PropertyDiff{
			Kind:      p.Update,
			InputDiff: true,
		}
	}

	if len(diffResponse.DetailedDiff) > 0 {
		diffResponse.HasChanges = true
	}

	return diffResponse, nil
}

func (c Cluster) Create(ctx p.Context, name string, args ClusterArgs, preview bool) (string, ClusterState, error) {
	newState, err := clusterApply(ctx, name, nil, &args, preview)
	if err != nil {
		return name, *newState, err
	}

	return name, *newState, nil
}

func (c *Cluster) Read(
	ctx p.Context,
	name string,
	args ClusterArgs,
	state ClusterState,
) (string, ClusterArgs, ClusterState, error) {
	newArgs, cluster, err := clusterArgsToK0sCtlCluster(name, &args)
	if err != nil {
		return name, *newArgs, state, err
	}

	state.ClusterArgs = *newArgs
	state.Kubeconfig = cluster.Metadata.Kubeconfig

	return name, *newArgs, state, nil
}

func (c Cluster) Update(
	ctx p.Context,
	name string,
	state ClusterState,
	args ClusterArgs,
	preview bool,
) (ClusterState, error) {
	newState, err := clusterApply(ctx, name, &state, &args, preview)
	if err != nil {
		return *newState, err
	}

	return *newState, nil
}

func (c Cluster) Delete(ctx p.Context, name string, state ClusterState) error {
	defer func() {
		_ = cleanupTempFiles(name)
	}()

	_, cluster, err := clusterArgsToK0sCtlCluster(name, &ClusterArgs{Metadata: state.Metadata, Spec: state.Spec})
	if err != nil {
		return err
	}

	if _, err = k0sctl.Reset(cluster); err != nil {
		return err
	}

	state.Kubeconfig = ""

	return nil
}

func clusterApply(
	ctx p.Context,
	name string,
	state *ClusterState,
	clusterArgs *ClusterArgs,
	preview bool,
) (*ClusterState, error) {
	defer func() {
		_ = cleanupTempFiles(name)
	}()

	applyConfig := k0sctl.ApplyConfig{
		SkipDowngradeCheck: false,
		NoDrain:            false,
		RestoreFrom:        "",
	}

	config := infer.GetConfig[Config](ctx)

	if config.SkipDowngradeCheck != nil {
		applyConfig.SkipDowngradeCheck = strings.ToLower(*config.SkipDowngradeCheck) == "true"
	}

	if config.SkipDowngradeCheck != nil {
		applyConfig.NoDrain = strings.ToLower(*config.NoDrain) == "true"
	}

	if state == nil {
		state = &ClusterState{}
	}

	newClusterArgs, cluster, err := clusterArgsToK0sCtlCluster(name, clusterArgs)
	if err != nil {
		return state, err
	}

	state.ClusterArgs = *newClusterArgs

	if preview {
		return state, nil
	}

	cluster, err = k0sctl.Apply(cluster, applyConfig)
	if err != nil {
		return state, err
	}

	state.Kubeconfig = cluster.Metadata.Kubeconfig

	return state, nil
}

func clusterArgsToK0sCtlCluster(name string, clusterArgs *ClusterArgs) (*ClusterArgs, *k0sctl.Cluster, error) {
	defaultCluster, err := k0sctl.DefaultClusterConfig()
	if err != nil {
		return clusterArgs, nil, err
	}

	defaultClusterBytes, err := yaml.Marshal(defaultCluster)
	if err != nil {
		return clusterArgs, nil, err
	}

	argsBytes, err := yaml.Marshal(clusterArgs)
	if err != nil {
		return clusterArgs, nil, err
	}

	clusterBytes, err := deepmerge.YAML(defaultClusterBytes, argsBytes, deepmerge.Config{
		PreventMultipleDefinitionsOfKeysWithPrimitiveValue: false,
	})
	if err != nil {
		return clusterArgs, nil, err
	}

	var cluster *k0sctl.Cluster

	if err := yaml.Unmarshal(clusterBytes, &cluster); err != nil {
		return clusterArgs, nil, err
	}

	if clusterArgs.APIVersion == nil || *clusterArgs.APIVersion == "" {
		clusterArgs.APIVersion = &cluster.APIVersion
	}

	if clusterArgs.Kind == nil || *clusterArgs.Kind == "" {
		clusterArgs.Kind = &cluster.Kind
	}

	if clusterArgs.Metadata == nil {
		clusterArgs.Metadata = &ClusterMetadata{}
	}

	if clusterArgs.Metadata.Name == nil {
		clusterArgs.Metadata.Name = &name
	}

	if clusterArgs.Spec == nil {
		clusterArgs.Spec = &ClusterSpec{}
	}

	if clusterArgs.Spec.K0s == nil {
		clusterArgs.Spec.K0s = &ClusterK0s{}
	}

	if clusterArgs.Spec.K0s.Version == nil {
		version := cluster.Spec.K0s.Version.String()
		clusterArgs.Spec.K0s.Version = &version
	}

	if clusterArgs.Spec.K0s.Config == nil {
		clusterArgs.Spec.K0s.Config = &ClusterK0sConfig{}
	}

	if clusterArgs.Spec.K0s.Config.Metadata == nil {
		clusterArgs.Spec.K0s.Config.Metadata = &ClusterMetadata{}
	}

	if clusterArgs.Spec.K0s.Config.Metadata.Name == nil {
		clusterArgs.Spec.K0s.Config.Metadata.Name = clusterArgs.Metadata.Name
	}

	if _, ok := cluster.Spec.K0s.Config["metadata"]; !ok {
		cluster.Spec.K0s.Config["metadata"] = dig.Mapping{}
	}

	if metadata, ok := cluster.Spec.K0s.Config["metadata"]; ok {
		if metadataMap, ok := metadata.(dig.Mapping); ok {
			metadataMap["name"] = cluster.Metadata.Name
		}
	}

	if err := clusterReplaceHosts(name, clusterArgs.Spec.Hosts, cluster); err != nil {
		return clusterArgs, nil, err
	}

	if err := clusterExternalEtcdReplace(name, clusterArgs, cluster); err != nil {
		return clusterArgs, nil, err
	}

	return clusterArgs, cluster, err
}

func clusterReplaceHosts(name string, hosts []ClusterHost, cluster *k0sctl.Cluster) error {
	for i, host := range hosts {
		if host.SSH != nil {
			if err := clusterHostsReplaceSSHKey(name, host.SSH, cluster.Spec.Hosts[i].SSH); err != nil {
				return err
			}
		}

		if host.WinRM != nil {
			if host.WinRM.CaCert != nil {
				filename, err := contentToTempFile(name, *host.WinRM.CaCert, true)
				if err != nil {
					return err
				}

				cluster.Spec.Hosts[i].WinRM.CACertPath = filename
			}

			if host.WinRM.Cert != nil {
				filename, err := contentToTempFile(name, *host.WinRM.Cert, true)
				if err != nil {
					return err
				}

				cluster.Spec.Hosts[i].WinRM.CertPath = filename
			}

			if host.WinRM.Key != nil {
				filename, err := contentToTempFile(name, *host.WinRM.Key, true)
				if err != nil {
					return err
				}

				cluster.Spec.Hosts[i].WinRM.KeyPath = filename
			}

			if err := clusterHostsReplaceSSHKey(name, host.WinRM.Bastion, cluster.Spec.Hosts[i].WinRM.Bastion); err != nil {
				return err
			}
		}
	}

	return nil
}

func clusterHostsReplaceSSHKey(name string, sshArgs *ClusterSSH, k0sctlSSHArgs *rig.SSH) error {
	if sshArgs.Key != nil {
		filename, err := contentToTempFile(name, *sshArgs.Key, true)
		if err != nil {
			return err
		}

		k0sctlSSHArgs.KeyPath = &filename
	}

	if sshArgs.Bastion != nil {
		return clusterHostsReplaceSSHKey(name, sshArgs.Bastion, k0sctlSSHArgs.Bastion)
	}

	return nil
}

func clusterExternalEtcdReplace(name string, clusterArgs *ClusterArgs, cluster *k0sctl.Cluster) error {
	if clusterArgs.Spec == nil || clusterArgs.Spec.K0s == nil || clusterArgs.Spec.K0s.Config == nil {
		return nil
	}

	if clusterArgs.Spec.K0s.Config.Spec == nil || clusterArgs.Spec.K0s.Config.Spec.Storage == nil {
		return nil
	}

	if clusterArgs.Spec.K0s.Config.Spec.Storage == nil && clusterArgs.Spec.K0s.Config.Spec.Storage.Etcd == nil {
		return nil
	}

	externalCluster := clusterArgs.Spec.K0s.Config.Spec.Storage.Etcd.ExternalCluster
	k0sctlExternalCluster := cluster.Spec.K0s.Config.Dig("spec", "storage", "etcd", "externalCluster")

	if externalCluster == nil || k0sctlExternalCluster == nil {
		return nil
	}

	k0sctlExternalClusterMap, ok := k0sctlExternalCluster.(dig.Mapping)
	if !ok {
		return nil
	}

	if externalCluster.CA != nil {
		filename, err := contentToTempFile(name, *externalCluster.CA, true)
		if err != nil {
			return err
		}

		k0sctlExternalClusterMap["caFile"] = &filename

		delete(k0sctlExternalClusterMap, "ca")
	}

	if externalCluster.ClientCert != nil {
		filename, err := contentToTempFile(name, *externalCluster.ClientCert, true)
		if err != nil {
			return err
		}

		k0sctlExternalClusterMap["clientCertFile"] = &filename

		delete(k0sctlExternalClusterMap, "clientCert")
	}

	if externalCluster.ClientKey != nil {
		filename, err := contentToTempFile(name, *externalCluster.ClientKey, true)
		if err != nil {
			return err
		}

		k0sctlExternalClusterMap["clientKeyFile"] = &filename

		delete(k0sctlExternalClusterMap, "clientKey")
	}

	return nil
}
