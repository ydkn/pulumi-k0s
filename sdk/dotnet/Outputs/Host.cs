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
    public sealed class Host
    {
        public readonly ImmutableDictionary<string, string>? Environment;
        public readonly ImmutableArray<Outputs.UploadFile> Files;
        public readonly Outputs.Hooks? Hooks;
        public readonly string? Hostname;
        public readonly ImmutableArray<string> InstallFlags;
        public readonly string? K0sBinaryPath;
        public readonly Outputs.Localhost? Localhost;
        public readonly string? Os;
        public readonly string? PrivateAddress;
        public readonly string? PrivateInterface;
        public readonly string Role;
        public readonly Outputs.SSH? Ssh;
        public readonly bool? UploadBinary;
        public readonly Outputs.WinRM? WinRM;

        [OutputConstructor]
        private Host(
            ImmutableDictionary<string, string>? environment,

            ImmutableArray<Outputs.UploadFile> files,

            Outputs.Hooks? hooks,

            string? hostname,

            ImmutableArray<string> installFlags,

            string? k0sBinaryPath,

            Outputs.Localhost? localhost,

            string? os,

            string? privateAddress,

            string? privateInterface,

            string role,

            Outputs.SSH? ssh,

            bool? uploadBinary,

            Outputs.WinRM? winRM)
        {
            Environment = environment;
            Files = files;
            Hooks = hooks;
            Hostname = hostname;
            InstallFlags = installFlags;
            K0sBinaryPath = k0sBinaryPath;
            Localhost = localhost;
            Os = os;
            PrivateAddress = privateAddress;
            PrivateInterface = privateInterface;
            Role = role;
            Ssh = ssh;
            UploadBinary = uploadBinary;
            WinRM = winRM;
        }
    }
}
