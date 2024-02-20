package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/ydkn/pulumi-k0s/sdk/go/k0s"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myProvider, err := k0s.NewProvider(ctx, "myProvider", &k0s.ProviderArgs{
			NoDrain: pulumi.String("true"),
		})
		if err != nil {
			return err
		}
		myCluster, err := k0s.NewCluster(ctx, "myCluster", &k0s.ClusterArgs{
			Spec: &k0s.ClusterSpecArgs{
				Hosts: []k0s.ClusterHostArgs{
					{
						Role: pulumi.String("controller+worker"),
						Localhost: {
							Enabled: pulumi.Bool(true),
						},
					},
				},
			},
		}, pulumi.Provider(myProvider))
		if err != nil {
			return err
		}
		ctx.Export("output", map[string]interface{}{
			"value": myCluster.Kubeconfig,
		})
		return nil
	})
}
