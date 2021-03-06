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
    public sealed class API
    {
        public readonly string? Address;
        public readonly string? ExternalAddress;
        public readonly ImmutableDictionary<string, string>? ExtraArgs;
        public readonly double? K0sApiPort;
        public readonly double? Port;
        public readonly ImmutableArray<string> Sans;

        [OutputConstructor]
        private API(
            string? address,

            string? externalAddress,

            ImmutableDictionary<string, string>? extraArgs,

            double? k0sApiPort,

            double? port,

            ImmutableArray<string> sans)
        {
            Address = address;
            ExternalAddress = externalAddress;
            ExtraArgs = extraArgs;
            K0sApiPort = k0sApiPort;
            Port = port;
            Sans = sans;
        }
    }
}
