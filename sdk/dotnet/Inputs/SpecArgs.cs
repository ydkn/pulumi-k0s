// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.K0s.Inputs
{

    public sealed class SpecArgs : global::Pulumi.ResourceArgs
    {
        [Input("hosts", required: true)]
        private InputList<Inputs.HostArgs>? _hosts;
        public InputList<Inputs.HostArgs> Hosts
        {
            get => _hosts ?? (_hosts = new InputList<Inputs.HostArgs>());
            set => _hosts = value;
        }

        [Input("k0s")]
        public Input<Inputs.K0sArgs>? K0s { get; set; }

        public SpecArgs()
        {
        }
        public static new SpecArgs Empty => new SpecArgs();
    }
}
