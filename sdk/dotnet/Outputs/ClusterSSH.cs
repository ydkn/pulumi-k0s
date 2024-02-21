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
    public sealed class ClusterSSH
    {
        public readonly string Address;
        public readonly Outputs.ClusterSSH? Bastion;
        public readonly string? HostKey;
        public readonly string? Key;
        public readonly int? Port;
        public readonly string? User;

        [OutputConstructor]
        private ClusterSSH(
            string address,

            Outputs.ClusterSSH? bastion,

            string? hostKey,

            string? key,

            int? port,

            string? user)
        {
            Address = address;
            Bastion = bastion;
            HostKey = hostKey;
            Key = key;
            Port = port;
            User = user;
        }
    }
}