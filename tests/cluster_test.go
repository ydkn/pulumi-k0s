package tests

import (
	"testing"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClusterCreatePreview(t *testing.T) {
	prov := integrationServer()
	name := "test-cluster"
	validPropertyKeys := []resource.PropertyKey{"apiVersion", "kind", "metadata", "spec", "kubeconfig"}

	response, err := prov.Create(p.CreateRequest{
		Urn:        urn("Cluster", name),
		Properties: resource.PropertyMap{},
		Preview:    true,
	})
	require.NoError(t, err)
	assert.Equal(t, name, response.ID)

	for k := range response.Properties {
		found := false

		for _, pKey := range validPropertyKeys {
			if k == pKey {
				found = true
			}
		}

		if !found {
			t.Errorf("unexpected property key: %s", k)
			t.Fail()
		}
	}
}
