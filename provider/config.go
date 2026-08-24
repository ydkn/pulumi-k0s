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

const (
	configDefaultSkipDowngradeCheck = false
	configDefaultNoDrain            = false
	configDefaultNoWait             = false
	configDefaultConcurrency        = 30
	configDefaultConcurrentUploads  = 5
)

// Config defines provider-level configuration
type Config struct {
	SkipDowngradeCheck *bool `pulumi:"skipDowngradeCheck,optional"`
	NoDrain            *bool `pulumi:"noDrain,optional"`
	NoWait             *bool `pulumi:"noWait,optional"`
	Concurrency        *int  `pulumi:"concurrency,optional"`
	ConcurrentUploads  *int  `pulumi:"concurrentUploads,optional"`
}

func (c *Config) Annotate(a infer.Annotator) {
	skipDowngradeCheckValue := configDefaultSkipDowngradeCheck
	a.Describe(&c.SkipDowngradeCheck, "Skip downgrade check")
	a.SetDefault(&c.SkipDowngradeCheck, &skipDowngradeCheckValue, "PULUMI_K0S_SKIP_DOWNGRADE_CHECK")

	noDrainValue := configDefaultNoDrain
	a.Describe(&c.NoDrain, "Do not drain worker nodes when upgrading")
	a.SetDefault(&c.NoDrain, &noDrainValue, "PULUMI_K0S_NO_DRAIN")

	noWaitValue := configDefaultNoWait
	a.Describe(&c.NoWait, "Do not wait for worker nodes to join")
	a.SetDefault(&c.NoWait, &noWaitValue, "PULUMI_K0S_NO_WAIT")

	concurrencyValue := configDefaultConcurrency
	a.Describe(&c.Concurrency, "Maximum number of hosts to configure in parallel, set to 0 for unlimited")
	a.SetDefault(&c.Concurrency, &concurrencyValue, "PULUMI_K0S_CONCURRENCY")

	concurrentUploadsValue := configDefaultConcurrentUploads
	a.Describe(&c.ConcurrentUploads, "Maximum number of files to upload in parallel, set to 0 for unlimited")
	a.SetDefault(&c.ConcurrentUploads, &concurrentUploadsValue, "PULUMI_K0S_CONCURRENT_UPLOADS")
}

func (c Config) Diff(_ context.Context, req infer.DiffRequest[Config, Config]) (p.DiffResponse, error) {
	diffResponse := p.DiffResponse{
		DeleteBeforeReplace: false,
		HasChanges:          false,
		DetailedDiff:        map[string]p.PropertyDiff{},
	}

	oldsProps, err := introspect.NewPropertiesMap(req.State)
	if err != nil {
		return p.DiffResponse{}, err
	}

	newsProps, err := introspect.NewPropertiesMap(req.Inputs)
	if err != nil {
		return p.DiffResponse{}, err
	}

	for key := range propertyMapDiff(oldsProps, newsProps, []resource.PropertyKey{}) {
		diffResponse.DetailedDiff[strings.SplitN(string(key), ".", 2)[0]] = p.PropertyDiff{
			Kind:      p.Update,
			InputDiff: true,
		}
	}

	if len(diffResponse.DetailedDiff) > 0 {
		diffResponse.HasChanges = true
	}

	return diffResponse, nil
}

func (c Config) Check(_ context.Context, req infer.CheckRequest) (infer.CheckResponse[Config], error) {
	// Remove "version" from inputs to avoid decode errors
	req.NewInputs = req.NewInputs.Delete("version")

	_, config, err := ende.Decode[Config](req.NewInputs)
	if err != nil {
		return infer.CheckResponse[Config]{}, err
	}

	return infer.CheckResponse[Config]{Inputs: config, Failures: []p.CheckFailure{}}, nil
}
