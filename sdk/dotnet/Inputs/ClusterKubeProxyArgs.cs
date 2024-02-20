// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.K0s.Inputs
{

    public sealed class ClusterKubeProxyArgs : global::Pulumi.ResourceArgs
    {
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        [Input("iptables")]
        public Input<Inputs.ClusterKubeProxyIPTablesArgs>? Iptables { get; set; }

        [Input("ipvs")]
        public Input<Inputs.ClusterKubeProxyIPVSArgs>? Ipvs { get; set; }

        [Input("mode")]
        public Input<string>? Mode { get; set; }

        [Input("nodePortAddresses")]
        public Input<string>? NodePortAddresses { get; set; }

        public ClusterKubeProxyArgs()
        {
        }
        public static new ClusterKubeProxyArgs Empty => new ClusterKubeProxyArgs();
    }
}
