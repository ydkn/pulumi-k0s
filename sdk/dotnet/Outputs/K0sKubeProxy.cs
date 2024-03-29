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
    public sealed class K0sKubeProxy
    {
        public readonly bool? Disabled;
        public readonly Outputs.K0sKubeProxyIPTables? Iptables;
        public readonly Outputs.K0sKubeProxyIPVS? Ipvs;
        public readonly string? Mode;
        public readonly string? NodePortAddresses;

        [OutputConstructor]
        private K0sKubeProxy(
            bool? disabled,

            Outputs.K0sKubeProxyIPTables? iptables,

            Outputs.K0sKubeProxyIPVS? ipvs,

            string? mode,

            string? nodePortAddresses)
        {
            Disabled = disabled;
            Iptables = iptables;
            Ipvs = ipvs;
            Mode = mode;
            NodePortAddresses = nodePortAddresses;
        }
    }
}
