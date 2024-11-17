// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.K0s
{
    [K0sResourceType("pulumi:providers:k0s")]
    public partial class Provider : global::Pulumi.ProviderResource
    {
        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs? args = null, CustomResourceOptions? options = null)
            : base("k0s", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://repo.ydkn.io/pulumi-k0s",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class ProviderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum number of hosts to configure in parallel, set to 0 for unlimited
        /// </summary>
        [Input("concurrency", json: true)]
        public Input<int>? Concurrency { get; set; }

        /// <summary>
        /// Maximum number of files to upload in parallel, set to 0 for unlimited
        /// </summary>
        [Input("concurrentUploads", json: true)]
        public Input<int>? ConcurrentUploads { get; set; }

        /// <summary>
        /// Do not drain worker nodes when upgrading
        /// </summary>
        [Input("noDrain", json: true)]
        public Input<bool>? NoDrain { get; set; }

        /// <summary>
        /// Do not wait for worker nodes to join
        /// </summary>
        [Input("noWait", json: true)]
        public Input<bool>? NoWait { get; set; }

        /// <summary>
        /// Skip downgrade check
        /// </summary>
        [Input("skipDowngradeCheck", json: true)]
        public Input<bool>? SkipDowngradeCheck { get; set; }

        public ProviderArgs()
        {
            Concurrency = Utilities.GetEnvInt32("PULUMI_K0S_CONCURRENCY") ?? 30;
            ConcurrentUploads = Utilities.GetEnvInt32("PULUMI_K0S_CONCURRENT_UPLOADS") ?? 5;
            NoDrain = Utilities.GetEnvBoolean("PULUMI_K0S_NO_DRAIN") ?? false;
            NoWait = Utilities.GetEnvBoolean("PULUMI_K0S_NO_WAIT") ?? false;
            SkipDowngradeCheck = Utilities.GetEnvBoolean("PULUMI_K0S_SKIP_DOWNGRADE_CHECK") ?? false;
        }
        public static new ProviderArgs Empty => new ProviderArgs();
    }
}
