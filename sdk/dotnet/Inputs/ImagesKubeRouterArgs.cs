// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.K0s.Inputs
{

    public sealed class ImagesKubeRouterArgs : Pulumi.ResourceArgs
    {
        [Input("cni")]
        public Input<Inputs.ContainerImageArgs>? Cni { get; set; }

        [Input("cniInstaller")]
        public Input<Inputs.ContainerImageArgs>? CniInstaller { get; set; }

        public ImagesKubeRouterArgs()
        {
        }
    }
}
