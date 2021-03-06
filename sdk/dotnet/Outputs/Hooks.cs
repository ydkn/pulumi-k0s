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
    public sealed class Hooks
    {
        public readonly Outputs.Hook? Apply;
        public readonly Outputs.Hook? Backup;
        public readonly Outputs.Hook? Reset;

        [OutputConstructor]
        private Hooks(
            Outputs.Hook? apply,

            Outputs.Hook? backup,

            Outputs.Hook? reset)
        {
            Apply = apply;
            Backup = backup;
            Reset = reset;
        }
    }
}
