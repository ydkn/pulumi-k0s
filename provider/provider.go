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

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

var Version string

const Name string = "k0s"

func Provider() p.Provider {
	p, err := infer.NewProviderBuilder().
		WithDisplayName(Name).
		WithDescription("A Pulumi package for creating and managing k0s clusters.").
		WithHomepage("https://github.com/ydkn/pulumi-k0s").
		WithRepository("https://github.com/ydkn/pulumi-k0s").
		WithKeywords("pulumi", "kubernetes", "k0s").
		WithPublisher("Florian Schwab").
		WithLogoURL("https://k0sproject.io/images/k0s-logo.png").
		WithLicense("Apache-2.0").
		WithPluginDownloadURL("https://repo.ydkn.io/pulumi-k0s").
		WithNamespace("ydkn").
		WithConfig(infer.Config(Config{})).
		WithResources(infer.Resource(Cluster{})).
		WithModuleMap(map[tokens.ModuleName]tokens.ModuleName{
			"provider": "index",
		}).Build()
	if err != nil {
		panic(fmt.Errorf("unable to build provider: %w", err))
	}

	return p
}
