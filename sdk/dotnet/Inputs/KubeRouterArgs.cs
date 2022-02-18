// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.K0s.Inputs
{

    public sealed class KubeRouterArgs : Pulumi.ResourceArgs
    {
        [Input("autoMTU")]
        public Input<bool>? AutoMTU { get; set; }

        [Input("mtu")]
        public Input<double>? Mtu { get; set; }

        [Input("peerRouterASNs")]
        public Input<string>? PeerRouterASNs { get; set; }

        [Input("peerRouterIPs")]
        public Input<string>? PeerRouterIPs { get; set; }

        public KubeRouterArgs()
        {
        }
    }
}
