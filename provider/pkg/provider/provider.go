package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/imdario/mergo"
	"github.com/k0sproject/rig"
	"github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	log "github.com/sirupsen/logrus"
	"github.com/ydkn/pulumi-k0s/provider/pkg/k0sctl"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"

	pbempty "github.com/golang/protobuf/ptypes/empty"
)

type k0sProvider struct {
	host    *provider.HostClient
	name    string
	version string

	skipDowngradeCheck bool
	noDrain            bool
}

func makeProvider(host *provider.HostClient, name, version string) (pulumirpc.ResourceProviderServer, error) {
	// Disable output of k0sctl
	log.SetOutput(io.Discard)
	rig.SetLogger(log.StandardLogger())

	// Return the new provider
	return &k0sProvider{
		host:    host,
		name:    name,
		version: version,
	}, nil
}

// Call dynamically executes a method in the provider associated with a component resource.
func (p *k0sProvider) Call(ctx context.Context, req *pulumirpc.CallRequest) (*pulumirpc.CallResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Call is not yet implemented")
}

// Construct creates a new component resource.
func (p *k0sProvider) Construct(ctx context.Context, req *pulumirpc.ConstructRequest) (*pulumirpc.ConstructResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Construct is not yet implemented")
}

// CheckConfig validates the configuration for this provider.
func (p *k0sProvider) CheckConfig(ctx context.Context, req *pulumirpc.CheckRequest) (*pulumirpc.CheckResponse, error) {
	return &pulumirpc.CheckResponse{Inputs: req.GetNews()}, nil
}

// DiffConfig diffs the configuration for this provider.
func (p *k0sProvider) DiffConfig(ctx context.Context, req *pulumirpc.DiffRequest) (*pulumirpc.DiffResponse, error) {
	olds, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	news, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	// Calculate the difference between old and new inputs.
	diff := olds.Diff(news, func(key resource.PropertyKey) bool {
		return strings.HasPrefix(string(key), "__")
	})

	changes := make([]string, 0)

	if diff != nil {
		for _, k := range diff.ChangedKeys() {
			changes = append(changes, string(k))
		}
	}

	changeType := pulumirpc.DiffResponse_DIFF_NONE
	if len(changes) > 0 {
		changeType = pulumirpc.DiffResponse_DIFF_SOME
	}

	return &pulumirpc.DiffResponse{
		Diffs:           changes,
		Changes:         changeType,
		HasDetailedDiff: false,
	}, nil
}

// Configure configures the resource provider with "globals" that control its behavior.
func (p *k0sProvider) Configure(_ context.Context, req *pulumirpc.ConfigureRequest) (*pulumirpc.ConfigureResponse, error) {
	const trueStr = "true"

	vars := req.GetVariables()

	skipDowngradeCheck := func() bool {
		// If the provider flag is set, use that value to determine behavior. This will override the ENV var.
		if enabled, exists := vars["k0s:config:skipDowngradeCheck"]; exists {
			return enabled == trueStr
		}
		// If the provider flag is not set, fall back to the ENV var.
		if enabled, exists := os.LookupEnv("PULUMI_K0S_SKIP_DOWNGRADE_CHECK"); exists {
			return enabled == trueStr
		}
		// Default to false.
		return false
	}
	if skipDowngradeCheck() {
		p.skipDowngradeCheck = true
	}

	noDrain := func() bool {
		// If the provider flag is set, use that value to determine behavior. This will override the ENV var.
		if enabled, exists := vars["k0s:config:noDrain"]; exists {
			return enabled == trueStr
		}
		// If the provider flag is not set, fall back to the ENV var.
		if enabled, exists := os.LookupEnv("PULUMI_K0S_NO_DRAIN"); exists {
			return enabled == trueStr
		}
		// Default to false.
		return false
	}
	if noDrain() {
		p.noDrain = true
	}

	return &pulumirpc.ConfigureResponse{
		AcceptSecrets:   req.GetAcceptSecrets(),
		SupportsPreview: true,
		AcceptResources: req.GetAcceptResources(),
		AcceptOutputs:   true,
	}, nil
}

// Invoke dynamically executes a built-in function in the provider.
func (p *k0sProvider) Invoke(_ context.Context, req *pulumirpc.InvokeRequest) (*pulumirpc.InvokeResponse, error) {
	tok := req.GetTok()
	return nil, fmt.Errorf("Unknown Invoke token '%s'", tok)
}

// StreamInvoke dynamically executes a built-in function in the provider. The result is streamed
// back as a series of messages.
func (p *k0sProvider) StreamInvoke(req *pulumirpc.InvokeRequest, server pulumirpc.ResourceProvider_StreamInvokeServer) error {
	tok := req.GetTok()
	return fmt.Errorf("Unknown StreamInvoke token '%s'", tok)
}

// Check validates that the given property bag is valid for a resource of the given type and returns
// the inputs that should be passed to successive calls to Diff, Create, or Update for this
// resource. As a rule, the provider inputs returned by a call to Check should preserve the original
// representation of the properties as present in the program inputs. Though this rule is not
// required for correctness, violations thereof can negatively impact the end-user experience, as
// the provider inputs are using for detecting and rendering diffs.
func (p *k0sProvider) Check(ctx context.Context, req *pulumirpc.CheckRequest) (*pulumirpc.CheckResponse, error) {
	urn := resource.URN(req.GetUrn())
	ty := urn.Type()
	if ty != "k0s:index:Cluster" {
		return nil, fmt.Errorf("Unknown resource type '%s'", ty)
	}

	news, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	cluster, err := propertiesToCluster(news)
	if err != nil {
		return nil, err
	}

	var failures []*pulumirpc.CheckFailure

	if err := cluster.Check(); err != nil {
		failures = []*pulumirpc.CheckFailure{{Reason: err.Error()}}
	}

	return &pulumirpc.CheckResponse{Inputs: req.News, Failures: failures}, nil
}

// Diff checks what impacts a hypothetical update will have on the resource's properties.
func (p *k0sProvider) Diff(ctx context.Context, req *pulumirpc.DiffRequest) (*pulumirpc.DiffResponse, error) {
	urn := resource.URN(req.GetUrn())
	ty := urn.Type()
	if ty != "k0s:index:Cluster" {
		return nil, fmt.Errorf("Unknown resource type '%s'", ty)
	}

	olds, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	oldInputs := parseCheckpointObject(olds)

	oldCluster, err := propertiesToCluster(olds)
	if err != nil {
		return nil, err
	}

	news, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	newCluster, err := propertiesToCluster(news)
	if err != nil {
		return nil, err
	}

	// Calculate the difference between old and new inputs.
	inputsDiff := oldInputs.Diff(news, func(key resource.PropertyKey) bool {
		return strings.HasPrefix(string(key), "__")
	})

	// Calculate the difference between old and new clusters.
	clusterDiff := prepareClusterForDiff(oldCluster).Diff(prepareClusterForDiff(newCluster),
		func(key resource.PropertyKey) bool {
			return strings.HasPrefix(string(key), "__")
		})

	changes := make([]string, 0)

	for _, diff := range []*resource.ObjectDiff{inputsDiff, clusterDiff} {
		if diff == nil {
			continue
		}

		for _, k := range diff.ChangedKeys() {
			changes = append(changes, string(k))
		}
	}

	changeType := pulumirpc.DiffResponse_DIFF_NONE
	if len(changes) > 0 {
		changeType = pulumirpc.DiffResponse_DIFF_SOME
	}

	return &pulumirpc.DiffResponse{
		Diffs:           changes,
		Changes:         changeType,
		HasDetailedDiff: false,
	}, nil
}

// Create allocates a new instance of the provided resource and returns its unique ID afterwards.
func (p *k0sProvider) Create(ctx context.Context, req *pulumirpc.CreateRequest) (*pulumirpc.CreateResponse, error) {
	urn := resource.URN(req.GetUrn())
	ty := urn.Type()
	if ty != "k0s:index:Cluster" {
		return nil, fmt.Errorf("Unknown resource type '%s'", ty)
	}

	inputs, err := plugin.UnmarshalProperties(req.GetProperties(),
		plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	cluster, err := propertiesToCluster(inputs)
	if err != nil {
		return nil, err
	}

	cluster, err = k0sctl.Apply(cluster,
		k0sctl.ApplyConfig{SkipDowngradeCheck: p.skipDowngradeCheck, NoDrain: p.noDrain})
	if err != nil {
		return nil, err
	}

	cluster, err = k0sctl.KubeConfig(cluster, "")
	if err != nil {
		return nil, err
	}

	outputProperties, err := plugin.MarshalProperties(
		checkpointObject(inputs, cluster),
		plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true},
	)
	if err != nil {
		return nil, err
	}

	return &pulumirpc.CreateResponse{
		Id:         urn.Name().String(),
		Properties: outputProperties,
	}, nil
}

// Read the current live state associated with a resource.
func (p *k0sProvider) Read(ctx context.Context, req *pulumirpc.ReadRequest) (*pulumirpc.ReadResponse, error) {
	urn := resource.URN(req.GetUrn())
	ty := urn.Type()
	if ty != "k0s:index:Cluster" {
		return nil, fmt.Errorf("Unknown resource type '%s'", ty)
	}

	return nil, status.Error(codes.Unimplemented, "Read is not yet implemented for 'k0s:index:Cluster'")
}

// Update updates an existing resource with new values.
func (p *k0sProvider) Update(ctx context.Context, req *pulumirpc.UpdateRequest) (*pulumirpc.UpdateResponse, error) {
	urn := resource.URN(req.GetUrn())
	ty := urn.Type()
	if ty != "k0s:index:Cluster" {
		return nil, fmt.Errorf("Unknown resource type '%s'", ty)
	}

	inputs, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	cluster, err := propertiesToCluster(inputs)
	if err != nil {
		return nil, err
	}

	cluster, err = k0sctl.Apply(cluster,
		k0sctl.ApplyConfig{SkipDowngradeCheck: p.skipDowngradeCheck, NoDrain: p.noDrain})
	if err != nil {
		return nil, err
	}

	cluster, err = k0sctl.KubeConfig(cluster, "")
	if err != nil {
		return nil, err
	}

	outputProperties, err := plugin.MarshalProperties(
		checkpointObject(inputs, cluster),
		plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true},
	)
	if err != nil {
		return nil, err
	}

	return &pulumirpc.UpdateResponse{
		Properties: outputProperties,
	}, nil
}

// Delete tears down an existing resource with the given ID.  If it fails, the resource is assumed
// to still exist.
func (p *k0sProvider) Delete(ctx context.Context, req *pulumirpc.DeleteRequest) (*pbempty.Empty, error) {
	urn := resource.URN(req.GetUrn())
	ty := urn.Type()
	if ty != "k0s:index:Cluster" {
		return nil, fmt.Errorf("Unknown resource type '%s'", ty)
	}

	inputs, err := plugin.UnmarshalProperties(req.GetProperties(),
		plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	cluster, err := propertiesToCluster(inputs)
	if err != nil {
		return nil, err
	}

	err = k0sctl.Reset(cluster)
	if err != nil {
		return nil, err
	}

	return &pbempty.Empty{}, nil
}

// GetPluginInfo returns generic information about this plugin, like its version.
func (p *k0sProvider) GetPluginInfo(context.Context, *pbempty.Empty) (*pulumirpc.PluginInfo, error) {
	return &pulumirpc.PluginInfo{
		Version: p.version,
	}, nil
}

// GetSchema returns the JSON-serialized schema for the provider.
func (p *k0sProvider) GetSchema(ctx context.Context, req *pulumirpc.GetSchemaRequest) (*pulumirpc.GetSchemaResponse, error) {
	if v := req.GetVersion(); v != 0 {
		return nil, fmt.Errorf("unsupported schema version %d", v)
	}

	return &pulumirpc.GetSchemaResponse{}, nil
}

// Cancel signals the provider to gracefully shut down and abort any ongoing resource operations.
// Operations aborted in this way will return an error (e.g., `Update` and `Create` will either a
// creation error or an initialization error). Since Cancel is advisory and non-blocking, it is up
// to the host to decide how long to wait after Cancel is called before (e.g.)
// hard-closing any gRPC connection.
func (p *k0sProvider) Cancel(context.Context, *pbempty.Empty) (*pbempty.Empty, error) {
	return &pbempty.Empty{}, nil
}

func checkpointObject(inputs resource.PropertyMap, cluster *k0sctl.Cluster) resource.PropertyMap {
	outputs := resource.NewPropertyMap(cluster)

	outputs["__inputs"] = resource.MakeSecret(resource.NewObjectProperty(inputs))

	return outputs
}

func parseCheckpointObject(obj resource.PropertyMap) resource.PropertyMap {
	if inputs, ok := obj["__inputs"]; ok {
		return inputs.ObjectValue()
	}

	return nil
}

func propertiesToCluster(propertyMap resource.PropertyMap) (*k0sctl.Cluster, error) {
	defaultCluster, err := k0sctl.DefaultClusterConfig()
	if err != nil {
		return nil, err
	}

	propBytes, err := json.Marshal(propertyMap.Mappable())
	if err != nil {
		return nil, err
	}

	var cluster *k0sctl.Cluster

	if err := json.Unmarshal(propBytes, &cluster); err != nil {
		return nil, err
	}

	cluster, err = mergeConfigs(defaultCluster, cluster)
	if err != nil {
		return nil, err
	}

	for _, host := range cluster.Spec.Hosts {
		if host.Environment == nil {
			host.Environment = map[string]string{}
		}
	}

	return cluster, err
}

func prepareClusterForDiff(cluster *k0sctl.Cluster) resource.PropertyMap {
	if cluster.Spec != nil {
		cluster.Spec.Hosts = nil
	}

	return resource.NewPropertyMap(cluster)
}

func mergeConfigs(configs ...*k0sctl.Cluster) (*k0sctl.Cluster, error) {
	if len(configs) == 0 {
		return nil, fmt.Errorf("no configs provided")
	}

	var config = configs[0]

	for i, c := range configs {
		if i == 0 {
			continue
		}

		if err := mergo.Merge(config, c, mergo.WithOverride); err != nil {
			return nil, err
		}
	}

	return config, nil
}
