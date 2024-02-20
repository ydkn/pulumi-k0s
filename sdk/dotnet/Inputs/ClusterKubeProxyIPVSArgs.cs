// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.K0s.Inputs
{

    public sealed class ClusterKubeProxyIPVSArgs : global::Pulumi.ResourceArgs
    {
        [Input("excludeCIDRs")]
        public Input<string>? ExcludeCIDRs { get; set; }

        [Input("minSyncPeriod")]
        public Input<string>? MinSyncPeriod { get; set; }

        [Input("scheduler")]
        public Input<string>? Scheduler { get; set; }

        [Input("strictARP")]
        public Input<bool>? StrictARP { get; set; }

        [Input("syncPeriod")]
        public Input<string>? SyncPeriod { get; set; }

        [Input("tcpFinTimeout")]
        public Input<string>? TcpFinTimeout { get; set; }

        [Input("tcpTimeout")]
        public Input<string>? TcpTimeout { get; set; }

        [Input("udpTimeout")]
        public Input<string>? UdpTimeout { get; set; }

        public ClusterKubeProxyIPVSArgs()
        {
        }
        public static new ClusterKubeProxyIPVSArgs Empty => new ClusterKubeProxyIPVSArgs();
    }
}
