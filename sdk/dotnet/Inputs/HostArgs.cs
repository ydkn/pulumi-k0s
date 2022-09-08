// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.K0s.Inputs
{

    public sealed class HostArgs : global::Pulumi.ResourceArgs
    {
        [Input("environment")]
        private InputMap<string>? _environment;
        public InputMap<string> Environment
        {
            get => _environment ?? (_environment = new InputMap<string>());
            set => _environment = value;
        }

        [Input("files")]
        private InputList<Inputs.UploadFileArgs>? _files;
        public InputList<Inputs.UploadFileArgs> Files
        {
            get => _files ?? (_files = new InputList<Inputs.UploadFileArgs>());
            set => _files = value;
        }

        [Input("hooks")]
        public Input<Inputs.HooksArgs>? Hooks { get; set; }

        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        [Input("installFlags")]
        private InputList<string>? _installFlags;
        public InputList<string> InstallFlags
        {
            get => _installFlags ?? (_installFlags = new InputList<string>());
            set => _installFlags = value;
        }

        [Input("k0sBinaryPath")]
        public Input<string>? K0sBinaryPath { get; set; }

        [Input("localhost")]
        public Input<Inputs.LocalhostArgs>? Localhost { get; set; }

        [Input("noTaints")]
        public Input<bool>? NoTaints { get; set; }

        [Input("os")]
        public Input<string>? Os { get; set; }

        [Input("privateAddress")]
        public Input<string>? PrivateAddress { get; set; }

        [Input("privateInterface")]
        public Input<string>? PrivateInterface { get; set; }

        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        [Input("ssh")]
        public Input<Inputs.SSHArgs>? Ssh { get; set; }

        [Input("uploadBinary")]
        public Input<bool>? UploadBinary { get; set; }

        [Input("winRM")]
        public Input<Inputs.WinRMArgs>? WinRM { get; set; }

        public HostArgs()
        {
        }
        public static new HostArgs Empty => new HostArgs();
    }
}
