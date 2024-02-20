import pulumi
import pulumi_k0s as k0s

my_provider = k0s.Provider("myProvider", no_drain="true")
my_cluster = k0s.Cluster("myCluster", spec=k0s.ClusterSpecArgs(
    hosts=[k0s.ClusterHostArgs(
        role="controller+worker",
        localhost=k0s.ClusterLocalhostArgs(
            enabled=True,
        ),
    )],
),
opts=pulumi.ResourceOptions(provider=my_provider))
pulumi.export("output", {
    "value": my_cluster.kubeconfig,
})
