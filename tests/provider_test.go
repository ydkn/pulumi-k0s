package tests

import (
	"testing"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/require"
)

func TestProviderConfigure(t *testing.T) {
	prov := integrationServer()

	err := prov.Configure(p.ConfigureRequest{
		Args: resource.PropertyMap{
			"noDrain":            resource.NewBoolProperty(true),
			"skipDowngradeCheck": resource.NewBoolProperty(true),
		},
	})
	require.NoError(t, err)
}
