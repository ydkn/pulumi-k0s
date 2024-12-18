// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.K0s.Outputs
{

    [OutputType]
    public sealed class ClusterHost
    {
        public readonly ImmutableDictionary<string, string>? Environment;
        public readonly ImmutableArray<Outputs.ClusterFile> Files;
        public readonly Outputs.ClusterHooks? Hooks;
        public readonly string? Hostname;
        public readonly ImmutableArray<string> InstallFlags;
        public readonly string? K0sBinaryPath;
        public readonly Outputs.ClusterLocalhost? Localhost;
        public readonly bool? NoTaints;
        public readonly Outputs.ClusterOpenSSH? OpenSSH;
        public readonly string? Os;
        public readonly string? PrivateAddress;
        public readonly string? PrivateInterface;
        public readonly string Role;
        public readonly Outputs.ClusterSSH? Ssh;
        public readonly bool? UploadBinary;
        public readonly Outputs.ClusterWinRM? WinRM;

        [OutputConstructor]
        private ClusterHost(
            ImmutableDictionary<string, string>? environment,

            ImmutableArray<Outputs.ClusterFile> files,

            Outputs.ClusterHooks? hooks,

            string? hostname,

            ImmutableArray<string> installFlags,

            string? k0sBinaryPath,

            Outputs.ClusterLocalhost? localhost,

            bool? noTaints,

            Outputs.ClusterOpenSSH? openSSH,

            string? os,

            string? privateAddress,

            string? privateInterface,

            string role,

            Outputs.ClusterSSH? ssh,

            bool? uploadBinary,

            Outputs.ClusterWinRM? winRM)
        {
            Environment = environment;
            Files = files;
            Hooks = hooks;
            Hostname = hostname;
            InstallFlags = installFlags;
            K0sBinaryPath = k0sBinaryPath;
            Localhost = localhost;
            NoTaints = noTaints;
            OpenSSH = openSSH;
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
