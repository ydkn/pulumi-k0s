// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.K0s.Inputs
{

    public sealed class K0sArgs : global::Pulumi.ResourceArgs
    {
        [Input("config")]
        public Input<Inputs.ConfigArgs>? Config { get; set; }

        [Input("dynamicConfig")]
        public Input<bool>? DynamicConfig { get; set; }

        [Input("version")]
        public Input<string>? Version { get; set; }

        public K0sArgs()
        {
        }
        public static new K0sArgs Empty => new K0sArgs();
    }
}
