// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.K0s.Inputs
{

    public sealed class DualStackArgs : Pulumi.ResourceArgs
    {
        [Input("IPv6podCIDR")]
        public Input<string>? IPv6podCIDR { get; set; }

        [Input("IPv6serviceCIDR")]
        public Input<string>? IPv6serviceCIDR { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public DualStackArgs()
        {
        }
    }
}
