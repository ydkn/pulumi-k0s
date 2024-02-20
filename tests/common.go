package tests

import (
	"github.com/blang/semver"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	k0s "github.com/ydkn/pulumi-k0s/provider"
)

func urn(typ, name string) resource.URN {
	return resource.NewURN("stack", "project", "", tokens.Type("test:index:"+typ), name)
}

func provider() p.Provider {
	return k0s.Provider()
}

func integrationServer() integration.Server {
	return integration.NewServer(k0s.Name, semver.MustParse("1.0.0"), provider())
}
