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
    public sealed class Images
    {
        public readonly Outputs.ImagesCalico? Calico;
        public readonly Outputs.ContainerImage? Coredns;
        public readonly string? Default_pull_policy;
        public readonly Outputs.ContainerImage? Konnectivity;
        public readonly Outputs.ContainerImage? Kubeproxy;
        public readonly Outputs.ImagesKubeRouter? Kuberouter;
        public readonly Outputs.ContainerImage? Metricsserver;

        [OutputConstructor]
        private Images(
            Outputs.ImagesCalico? calico,

            Outputs.ContainerImage? coredns,

            string? default_pull_policy,

            Outputs.ContainerImage? konnectivity,

            Outputs.ContainerImage? kubeproxy,

            Outputs.ImagesKubeRouter? kuberouter,

            Outputs.ContainerImage? metricsserver)
        {
            Calico = calico;
            Coredns = coredns;
            Default_pull_policy = default_pull_policy;
            Konnectivity = konnectivity;
            Kubeproxy = kubeproxy;
            Kuberouter = kuberouter;
            Metricsserver = metricsserver;
        }
    }
}
