// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.K0s.Inputs
{

    public sealed class K0sKubeRouterArgs : global::Pulumi.ResourceArgs
    {
        [Input("autoMTU")]
        public Input<bool>? AutoMTU { get; set; }

        [Input("extraArgs")]
        private InputMap<string>? _extraArgs;
        public InputMap<string> ExtraArgs
        {
            get => _extraArgs ?? (_extraArgs = new InputMap<string>());
            set => _extraArgs = value;
        }

        [Input("hairpin")]
        public Input<string>? Hairpin { get; set; }

        [Input("ipMasq")]
        public Input<bool>? IpMasq { get; set; }

        [Input("metricsPort")]
        public Input<int>? MetricsPort { get; set; }

        [Input("mtu")]
        public Input<int>? Mtu { get; set; }

        public K0sKubeRouterArgs()
        {
        }
        public static new K0sKubeRouterArgs Empty => new K0sKubeRouterArgs();
    }
}
