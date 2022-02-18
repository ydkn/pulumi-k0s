// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.K0s.Outputs
{

    [OutputType]
    public sealed class KubeRouter
    {
        public readonly bool? AutoMTU;
        public readonly double? Mtu;
        public readonly string? PeerRouterASNs;
        public readonly string? PeerRouterIPs;

        [OutputConstructor]
        private KubeRouter(
            bool? autoMTU,

            double? mtu,

            string? peerRouterASNs,

            string? peerRouterIPs)
        {
            AutoMTU = autoMTU;
            Mtu = mtu;
            PeerRouterASNs = peerRouterASNs;
            PeerRouterIPs = peerRouterIPs;
        }
    }
}
