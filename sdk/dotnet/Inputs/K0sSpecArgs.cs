// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.K0s.Inputs
{

    public sealed class K0sSpecArgs : Pulumi.ResourceArgs
    {
        [Input("api")]
        public Input<Inputs.APIArgs>? Api { get; set; }

        [Input("controllerManager")]
        public Input<Inputs.ControllerManagerArgs>? ControllerManager { get; set; }

        [Input("images")]
        public Input<Inputs.ImagesArgs>? Images { get; set; }

        [Input("installConfig")]
        public Input<Inputs.InstallConfigArgs>? InstallConfig { get; set; }

        [Input("konnectivity")]
        public Input<Inputs.KonnectivityArgs>? Konnectivity { get; set; }

        [Input("network")]
        public Input<Inputs.NetworkArgs>? Network { get; set; }

        [Input("podSecurityPolicy")]
        public Input<Inputs.PodSecurityPolicyArgs>? PodSecurityPolicy { get; set; }

        [Input("scheduler")]
        public Input<Inputs.SchedulerArgs>? Scheduler { get; set; }

        [Input("storage")]
        public Input<Inputs.StorageArgs>? Storage { get; set; }

        [Input("telemetry")]
        public Input<Inputs.TelemetryArgs>? Telemetry { get; set; }

        public K0sSpecArgs()
        {
        }
    }
}
