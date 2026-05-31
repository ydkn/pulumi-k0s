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
	"context"
	"strings"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/ydkn/pulumi-k0s/provider/internal/introspect"
	"github.com/ydkn/pulumi-k0s/provider/internal/upstream/ende"
)

// Random is the controller for the resource.
//
// Each resource has a controlling struct.
// Resource behavior is determined by implementing methods on the controlling struct.
// The `Create` method is mandatory, but other methods are optional.
// - Check: Remap inputs before they are typed.
// - Diff: Change how instances of a resource are compared.
// - Update: Mutate a resource in place.
// - Read: Get the state of a resource from the backing provider.
// - Delete: Custom logic when the resource is deleted.
// - Annotate: Describe fields and set defaults for a resource.
// - WireDependencies: Control how outputs and secrets flows through values.
type Cluster struct{}

func (c *Cluster) Annotate(a infer.Annotator) {
	a.Describe(&c, "The k0s cluster resource.")
}

func (c Cluster) Check(
	ctx context.Context,
	req infer.CheckRequest,
) (infer.CheckResponse[ClusterArgs], error) {
	res := infer.CheckResponse[ClusterArgs]{Failures: []p.CheckFailure{}}

	_, args, decodeErr := ende.Decode[ClusterArgs](req.NewInputs)
	if decodeErr != nil {
		return res, decodeErr
	}

	if err := args.FillDefaults(nil); err != nil {
		return res, err
	}

	if err := NewK0sctl(&args).Validate(); err != nil {
		res.Failures = append(res.Failures, p.CheckFailure{Reason: err.Error()})
	}

	res.Inputs = args

	return res, nil
}

func (c Cluster) Diff(
	ctx context.Context,
	req infer.DiffRequest[ClusterArgs, ClusterState],
) (infer.DiffResponse, error) {
	res := infer.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          false,
		DetailedDiff:        map[string]p.PropertyDiff{},
	}

	if err := req.Inputs.FillDefaults(nil); err != nil {
		return res, err
	}

	oldsProps, err := introspect.NewPropertiesMap(req.State)
	if err != nil {
		return res, err
	}

	newsProps, err := introspect.NewPropertiesMap(req.Inputs)
	if err != nil {
		return res, err
	}

	for key := range propertyMapDiff(oldsProps, newsProps, []resource.PropertyKey{"kubeconfig"}) {
		res.DetailedDiff[strings.SplitN(string(key), ".", 2)[0]] = p.PropertyDiff{
			Kind:      p.Update,
			InputDiff: true,
		}
	}

	if len(res.DetailedDiff) > 0 {
		res.HasChanges = true
	}

	return res, nil
}

func (c Cluster) Read(
	ctx context.Context,
	req infer.ReadRequest[ClusterArgs, ClusterState],
) (infer.ReadResponse[ClusterArgs, ClusterState], error) {
	state := ClusterState{ClusterArgs: req.Inputs}
	res := infer.ReadResponse[ClusterArgs, ClusterState]{ID: *req.State.Metadata.Name, Inputs: req.Inputs, State: state}

	if err := req.Inputs.FillDefaults(nil); err != nil {
		return res, err
	}

	manager := NewK0sctl(&req.Inputs)

	if err := manager.Kubeconfig(ctx); err != nil {
		return res, err
	}

	if manager.kubeconfig != nil {
		state.Kubeconfig = *manager.kubeconfig
	}

	return infer.ReadResponse[ClusterArgs, ClusterState]{ID: *req.State.Metadata.Name, Inputs: req.Inputs, State: state},
		nil
}

func (c Cluster) Create(
	ctx context.Context,
	req infer.CreateRequest[ClusterArgs],
) (infer.CreateResponse[ClusterState], error) {
	config := infer.GetConfig[Config](ctx)
	state := ClusterState{ClusterArgs: req.Inputs}
	res := infer.CreateResponse[ClusterState]{ID: req.Name, Output: state}

	if err := req.Inputs.FillDefaults(&req.Name); err != nil {
		return res, err
	}

	if req.DryRun {
		return res, nil
	}

	manager := NewK0sctl(&req.Inputs)

	if err := manager.Apply(ctx, &config); err != nil {
		return res, err
	}

	state.ClusterArgs = req.Inputs

	if manager.kubeconfig != nil {
		state.Kubeconfig = *manager.kubeconfig
	}

	return infer.CreateResponse[ClusterState]{ID: req.Name, Output: state}, nil
}

func (c Cluster) Update(
	ctx context.Context,
	req infer.UpdateRequest[ClusterArgs, ClusterState],
) (infer.UpdateResponse[ClusterState], error) {
	config := infer.GetConfig[Config](ctx)
	state := req.State
	res := infer.UpdateResponse[ClusterState]{Output: state}

	if err := req.Inputs.FillDefaults(nil); err != nil {
		return res, err
	}

	manager := NewK0sctl(&req.Inputs)

	if err := manager.Apply(ctx, &config); err != nil {
		return res, err
	}

	state.ClusterArgs = req.Inputs

	if manager.kubeconfig != nil {
		state.Kubeconfig = *manager.kubeconfig
	}

	return infer.UpdateResponse[ClusterState]{Output: state}, nil
}

func (c Cluster) Delete(ctx context.Context, req infer.DeleteRequest[ClusterState]) error {
	if err := req.State.ClusterArgs.FillDefaults(nil); err != nil {
		return err
	}

	if err := NewK0sctl(&req.State.ClusterArgs).Reset(ctx); err != nil {
		return err
	}

	return nil
}
