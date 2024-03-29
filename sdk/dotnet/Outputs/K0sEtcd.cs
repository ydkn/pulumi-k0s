// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.K0s.Outputs
{

    [OutputType]
    public sealed class K0sEtcd
    {
        public readonly Outputs.K0sEtcdExternalCluster? ExternalCluster;
        public readonly ImmutableDictionary<string, string>? ExtraArgs;
        public readonly string? PeerAddress;

        [OutputConstructor]
        private K0sEtcd(
            Outputs.K0sEtcdExternalCluster? externalCluster,

            ImmutableDictionary<string, string>? extraArgs,

            string? peerAddress)
        {
            ExternalCluster = externalCluster;
            ExtraArgs = extraArgs;
            PeerAddress = peerAddress;
        }
    }
}
