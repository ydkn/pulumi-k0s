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
    public sealed class ImagesCalico
    {
        public readonly Outputs.ContainerImage? Cni;
        public readonly Outputs.ContainerImage? Kubecontrollers;
        public readonly Outputs.ContainerImage? Node;

        [OutputConstructor]
        private ImagesCalico(
            Outputs.ContainerImage? cni,

            Outputs.ContainerImage? kubecontrollers,

            Outputs.ContainerImage? node)
        {
            Cni = cni;
            Kubecontrollers = kubecontrollers;
            Node = node;
        }
    }
}
