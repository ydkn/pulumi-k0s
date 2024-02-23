package provider

import (
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/ydkn/pulumi-k0s/provider/internal/introspect"
)

type ClusterInputs struct {
	APIVersion *string          `pulumi:"apiVersion,optional" json:"apiVersion,omitempty"`
	Kind       *string          `pulumi:"kind,optional" json:"kind,omitempty"`
	Metadata   *ClusterMetadata `pulumi:"metadata,optional" json:"metadata,omitempty"`
	Spec       *ClusterSpec     `pulumi:"spec,optional" json:"spec,omitempty"`
	Kubeconfig *string          `json:"-"`
}

type ClusterMetadata struct {
	Name *string `pulumi:"name" json:"name"`
}

type ClusterSpec struct {
	Hosts []ClusterHost `pulumi:"hosts" json:"hosts"`
	K0s   *ClusterK0s   `pulumi:"k0s,optional" json:"k0s,omitempty"`
}

type ClusterHost struct {
	Role             *string           `pulumi:"role" json:"role"`
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

type ClusterWinRM struct {
	Address       *string     `pulumi:"address" json:"address"`
	Port          *int        `pulumi:"port,optional" json:"port,omitempty"`
	User          *string     `pulumi:"user,optional" json:"user,omitempty"`
	Password      *string     `pulumi:"password,optional" provider:"secret" json:"password,omitempty"`
	UseHTTPS      *bool       `pulumi:"useHTTPS,optional" json:"useHTTPS,omitempty"`
	Insecure      *bool       `pulumi:"insecure,optional" json:"insecure,omitempty"`
	UseNTLM       *bool       `pulumi:"useNTLM,optional" json:"useNTLM,omitempty"`
	CaCert        *string     `pulumi:"caCert,optional" provider:"secret" json:"caCert,omitempty"`
	Cert          *string     `pulumi:"cert,optional" provider:"secret" json:"cert,omitempty"`
	Key           *string     `pulumi:"key,optional" provider:"secret" json:"key,omitempty"`
	TLSServerName *string     `pulumi:"tlsServerName,optional" json:"tlsServerName,omitempty"`
	Bastion       *ClusterSSH `pulumi:"bastion,optional" json:"bastion,omitempty"`
}

type ClusterSSH struct {
	Address *string     `pulumi:"address" json:"address,omitempty"`
	Port    *int        `pulumi:"port,optional" json:"port,omitempty"`
	User    *string     `pulumi:"user,optional" json:"user,omitempty"`
	Key     *string     `pulumi:"key,optional" provider:"secret" json:"-"`
	HostKey *string     `pulumi:"hostKey,optional" json:"hostKey,omitempty"`
	Bastion *ClusterSSH `pulumi:"bastion,optional" json:"bastion,omitempty"`
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

type K0s struct {
	APIVersion *string      `pulumi:"apiVersion,optional" json:"apiVersion,omitempty"`
	Kind       *string      `pulumi:"kind,optional" json:"kind,omitempty"`
	Metadata   *K0sMetadata `pulumi:"metadata,optional" json:"metadata,omitempty"`
	Spec       *K0sSpec     `pulumi:"spec,optional" json:"spec,omitempty"`
}

type K0sMetadata struct {
	Name *string `pulumi:"name" json:"name"`
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
	WorkerProfiles    []K0sWorkerProfile    `pulumi:"workerProfiles,optional" json:"workerProfiles,omitempty"`
	FeatureGates      []K0sFeatureGate      `pulumi:"featureGates,optional" json:"featureGates,omitempty"`
	Telemetry         *K0sTelemetry         `pulumi:"telemetry,optional" json:"telemetry,omitempty"`
}

type K0sAPI struct {
	Address         *string           `pulumi:"address,optional" json:"address,omitempty"`
	Port            *int              `pulumi:"port,optional" json:"port,omitempty"`
	K0sApiPort      *int              `pulumi:"k0sApiPort,optional" json:"k0sApiPort,omitempty"`
	ExternalAddress *string           `pulumi:"externalAddress,optional" json:"externalAddress,omitempty"`
	SANs            []string          `pulumi:"sans,optional" json:"sans,omitempty"`
	ExtraArgs       map[string]string `pulumi:"extraArgs,optional" json:"extraArgs,omitempty"`
}

type K0sImages struct {
	DefaultPullPolicy *string             `pulumi:"default_pull_policy,optional" json:"default_pull_policy,omitempty"`
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

type K0sInstallConfigUser struct {
	EtcdUser          *string `pulumi:"etcdUser,optional" json:"etcdUser,omitempty"`
	KineUser          *string `pulumi:"kineUser,optional" json:"kineUser,omitempty"`
	KonnectivityUser  *string `pulumi:"konnectivityUser,optional" json:"konnectivityUser,omitempty"`
	KubeAPIServerUser *string `pulumi:"kubeAPIserverUser,optional" json:"kubeAPIserverUser,omitempty"`
	KubeSchedulerUser *string `pulumi:"kubeSchedulerUser,optional" json:"kubeSchedulerUser,omitempty"`
}

type K0sKonnectivity struct {
	AdminPort *int `pulumi:"adminPort,optional" json:"adminPort,omitempty"`
	AgentPort *int `pulumi:"agentPort,optional" json:"agentPort,omitempty"`
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

type K0sKubeProxy struct {
	Disabled          *bool                 `pulumi:"disabled,optional" json:"disabled,omitempty"`
	Mode              *string               `pulumi:"mode,optional" json:"mode,omitempty"`
	IPTables          *K0sKubeProxyIPTables `pulumi:"iptables,optional" json:"iptables,omitempty"`
	IPVS              *K0sKubeProxyIPVS     `pulumi:"ipvs,optional" json:"ipvs,omitempty"`
	NodePortAddresses *string               `pulumi:"nodePortAddresses,optional" json:"nodePortAddresses,omitempty"`
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
	DataSource *string `pulumi:"dataSource" json:"dataSource,omitempty" provider:"secret"`
}

type K0sTelemetry struct {
	Enabled *bool `pulumi:"enabled,optional" json:"enabled,omitempty"`
}

type K0sWorkerProfile struct {
	Name   *string        `pulumi:"name" json:"name"`
	Values map[string]any `pulumi:"values" json:"values"`
}

type K0sFeatureGate struct {
	Enabled    *bool    `pulumi:"enabled,optional" json:"enabled,omitempty"`
	Name       *string  `pulumi:"name" json:"name"`
	Components []string `pulumi:"components,optional" json:"components,omitempty"`
}

type ClusterOutputs struct {
	ClusterInputs
	Kubeconfig *string `pulumi:"kubeconfig" provider:"secret" json:"-"`
}

type Cluster struct{}

func (c Cluster) Check(
	ctx p.Context,
	name string,
	olds ClusterOutputs,
	news ClusterInputs,
) (ClusterInputs, []p.CheckFailure, error) {
	failures := []p.CheckFailure{}

	manager, err := c.newManager(name, &news)
	if err != nil {
		return news, failures, err
	}

	if err := manager.Validate(); err != nil {
		failures = append(failures, p.CheckFailure{Reason: err.Error()})
	}

	return news, failures, nil
}

func (c Cluster) Diff(ctx p.Context, name string, olds ClusterOutputs, news ClusterInputs) (p.DiffResponse, error) {
	diffResponse := p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          false,
		DetailedDiff:        map[string]p.PropertyDiff{},
	}

	if err := PrepareCluster(name, &news); err != nil {
		return diffResponse, err
	}

	oldsProps, err := introspect.NewPropertiesMap(olds)
	if err != nil {
		return p.DiffResponse{}, err
	}

	newsProps, err := introspect.NewPropertiesMap(news)
	if err != nil {
		return p.DiffResponse{}, err
	}

	for key := range propertyMapDiff(oldsProps, newsProps, []resource.PropertyKey{"kubeconfig"}) {
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

func (c *Cluster) Read(
	ctx p.Context,
	name string,
	news ClusterInputs,
	olds ClusterOutputs,
) (string, ClusterInputs, ClusterOutputs, error) {
	manager, err := c.newManager(name, &news)
	if err != nil {
		return name, news, olds, err
	}

	if err := manager.Kubeconfig(); err != nil {
		return name, news, olds, err
	}

	olds.ClusterInputs = news

	if news.Kubeconfig != nil {
		olds.Kubeconfig = news.Kubeconfig
	}

	return name, news, olds, nil
}

func (c Cluster) Create(
	ctx p.Context,
	name string,
	news ClusterInputs,
	preview bool,
) (string, ClusterOutputs, error) {
	config := infer.GetConfig[Config](ctx)
	olds := ClusterOutputs{ClusterInputs: news}

	manager, err := c.newManager(name, &news)
	if err != nil {
		return name, olds, err
	}

	if preview {
		olds.ClusterInputs = news

		return name, olds, err
	}

	if err := manager.Apply(&config); err != nil {
		return name, olds, err
	}

	if news.Kubeconfig != nil {
		olds.Kubeconfig = news.Kubeconfig
	}

	return name, olds, nil
}

func (c Cluster) Update(
	ctx p.Context,
	name string,
	olds ClusterOutputs,
	news ClusterInputs,
	preview bool,
) (ClusterOutputs, error) {
	config := infer.GetConfig[Config](ctx)

	manager, err := c.newManager(name, &news)
	if err != nil {
		return olds, err
	}

	if preview {
		olds.ClusterInputs = news

		return olds, nil
	}

	if err := manager.Apply(&config); err != nil {
		return olds, err
	}

	olds.ClusterInputs = news

	if news.Kubeconfig != nil {
		olds.Kubeconfig = news.Kubeconfig
	}

	return olds, nil
}

func (c Cluster) Delete(ctx p.Context, name string, olds ClusterOutputs) error {
	manager, err := c.newManager(name, &olds.ClusterInputs)
	if err != nil {
		return err
	}

	if err := manager.Reset(); err != nil {
		return err
	}

	olds.Kubeconfig = nil

	return nil
}

func (c Cluster) newManager(name string, news *ClusterInputs) (*K0sctl, error) {
	if err := PrepareCluster(name, news); err != nil {
		return nil, err
	}

	return NewK0sctl(news), nil
}
