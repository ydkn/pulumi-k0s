// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.K0s.Inputs
{

    public sealed class ClusterHooksArgs : global::Pulumi.ResourceArgs
    {
        [Input("apply")]
        public Input<Inputs.ClusterHookArgs>? Apply { get; set; }

        [Input("backup")]
        public Input<Inputs.ClusterHookArgs>? Backup { get; set; }

        [Input("reset")]
        public Input<Inputs.ClusterHookArgs>? Reset { get; set; }

        public ClusterHooksArgs()
        {
        }
        public static new ClusterHooksArgs Empty => new ClusterHooksArgs();
    }
}
