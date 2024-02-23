// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.K0s.Inputs
{

    public sealed class K0sSpecArgs : global::Pulumi.ResourceArgs
    {
        [Input("api")]
        public Input<Inputs.K0sAPIArgs>? Api { get; set; }

        [Input("controllerManager")]
        public Input<Inputs.K0sControllerManagerArgs>? ControllerManager { get; set; }

        [Input("featureGates")]
        private InputList<Inputs.K0sFeatureGateArgs>? _featureGates;
        public InputList<Inputs.K0sFeatureGateArgs> FeatureGates
        {
            get => _featureGates ?? (_featureGates = new InputList<Inputs.K0sFeatureGateArgs>());
            set => _featureGates = value;
        }

        [Input("images")]
        public Input<Inputs.K0sImagesArgs>? Images { get; set; }

        [Input("installConfig")]
        public Input<Inputs.K0sInstallConfigArgs>? InstallConfig { get; set; }

        [Input("konnectivity")]
        public Input<Inputs.K0sKonnectivityArgs>? Konnectivity { get; set; }

        [Input("network")]
        public Input<Inputs.K0sNetworkArgs>? Network { get; set; }

        [Input("podSecurityPolicy")]
        public Input<Inputs.K0sPodSecurityPolicyArgs>? PodSecurityPolicy { get; set; }

        [Input("scheduler")]
        public Input<Inputs.K0sSchedulerArgs>? Scheduler { get; set; }

        [Input("storage")]
        public Input<Inputs.K0sStorageArgs>? Storage { get; set; }

        [Input("telemetry")]
        public Input<Inputs.K0sTelemetryArgs>? Telemetry { get; set; }

        [Input("workerProfiles")]
        private InputList<Inputs.K0sWorkerProfileArgs>? _workerProfiles;
        public InputList<Inputs.K0sWorkerProfileArgs> WorkerProfiles
        {
            get => _workerProfiles ?? (_workerProfiles = new InputList<Inputs.K0sWorkerProfileArgs>());
            set => _workerProfiles = value;
        }

        public K0sSpecArgs()
        {
        }
        public static new K0sSpecArgs Empty => new K0sSpecArgs();
    }
}