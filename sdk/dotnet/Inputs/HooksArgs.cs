// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.K0s.Inputs
{

    public sealed class HooksArgs : Pulumi.ResourceArgs
    {
        [Input("apply")]
        public Input<Inputs.HookArgs>? Apply { get; set; }

        [Input("backup")]
        public Input<Inputs.HookArgs>? Backup { get; set; }

        [Input("reset")]
        public Input<Inputs.HookArgs>? Reset { get; set; }

        public HooksArgs()
        {
        }
    }
}
