// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.K0s.Inputs
{

    public sealed class ImagesCalicoArgs : Pulumi.ResourceArgs
    {
        [Input("cni")]
        public Input<Inputs.ContainerImageArgs>? Cni { get; set; }

        [Input("kubecontrollers")]
        public Input<Inputs.ContainerImageArgs>? Kubecontrollers { get; set; }

        [Input("node")]
        public Input<Inputs.ContainerImageArgs>? Node { get; set; }

        public ImagesCalicoArgs()
        {
        }
    }
}
