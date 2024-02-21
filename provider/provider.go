package provider

import (
	"io"

	"github.com/k0sproject/rig"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/sirupsen/logrus"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

const Name string = "k0s"

func Provider() p.Provider {
	// Disable output of k0sctl
	logrus.SetOutput(io.Discard)
	rig.SetLogger(logrus.StandardLogger())

	return infer.Provider(infer.Options{
		Config: infer.Config[Config](),
		Resources: []infer.InferredResource{
			infer.Resource[Cluster, ClusterArgs, ClusterState](),
		},
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"provider": "index",
		},
	})
}
