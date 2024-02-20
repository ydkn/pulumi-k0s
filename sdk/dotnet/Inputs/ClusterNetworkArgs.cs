// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.K0s.Inputs
{

    public sealed class ClusterNetworkArgs : global::Pulumi.ResourceArgs
    {
        [Input("calico")]
        public Input<Inputs.ClusterCalicoArgs>? Calico { get; set; }

        [Input("clusterDomain")]
        public Input<string>? ClusterDomain { get; set; }

        [Input("dualStack")]
        public Input<Inputs.ClusterDualStackArgs>? DualStack { get; set; }

        [Input("kubeProxy")]
        public Input<Inputs.ClusterKubeProxyArgs>? KubeProxy { get; set; }

        [Input("kuberouter")]
        public Input<Inputs.ClusterKubeRouterArgs>? Kuberouter { get; set; }

        [Input("nodeLocalLoadBalancing")]
        public Input<Inputs.ClusterNodeLocalLoadBalancingArgs>? NodeLocalLoadBalancing { get; set; }

        [Input("podCIDR")]
        public Input<string>? PodCIDR { get; set; }

        [Input("provider")]
        public Input<string>? Provider { get; set; }

        [Input("serviceCIDR")]
        public Input<string>? ServiceCIDR { get; set; }

        public ClusterNetworkArgs()
        {
        }
        public static new ClusterNetworkArgs Empty => new ClusterNetworkArgs();
    }
}
