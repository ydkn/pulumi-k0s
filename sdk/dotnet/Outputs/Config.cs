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
    public sealed class Config
    {
        public readonly Outputs.Metadata? Metadata;
        public readonly Outputs.K0sSpec? Spec;

        [OutputConstructor]
        private Config(
            Outputs.Metadata? metadata,

            Outputs.K0sSpec? spec)
        {
            Metadata = metadata;
            Spec = spec;
        }
    }
}
