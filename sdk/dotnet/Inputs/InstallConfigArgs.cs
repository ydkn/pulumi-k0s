// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.K0s.Inputs
{

    public sealed class InstallConfigArgs : Pulumi.ResourceArgs
    {
        [Input("users")]
        public Input<Inputs.InstallConfigUsersArgs>? Users { get; set; }

        public InstallConfigArgs()
        {
        }
    }
}
