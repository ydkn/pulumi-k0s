using System.Collections.Generic;
using System.Linq;
using Pulumi;
using K0s = Pulumi.K0s;

return await Deployment.RunAsync(() => 
{
    var myProvider = new K0s.Provider("myProvider", new()
    {
        NoDrain = true,
    });

    var myCluster = new K0s.Cluster("myCluster", new()
    {
        Spec = new K0s.Inputs.ClusterSpecArgs
        {
            Hosts = new[]
            {
                new K0s.Inputs.ClusterHostArgs
                {
                    Role = "controller+worker",
                    Localhost = new K0s.Inputs.ClusterLocalhostArgs
                    {
                        Enabled = true,
                    },
                },
            },
        },
    }, new CustomResourceOptions
    {
        Provider = myProvider,
    });

    return new Dictionary<string, object?>
    {
        ["output"] = 
        {
            { "value", myCluster.Kubeconfig },
        },
    };
});

