// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.K0s.Inputs
{

    public sealed class UploadFileArgs : global::Pulumi.ResourceArgs
    {
        [Input("dirPerm")]
        public Input<double>? DirPerm { get; set; }

        [Input("dst")]
        public Input<string>? Dst { get; set; }

        [Input("dstDir")]
        public Input<string>? DstDir { get; set; }

        [Input("group")]
        public Input<string>? Group { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("perm")]
        public Input<double>? Perm { get; set; }

        [Input("src")]
        public Input<string>? Src { get; set; }

        [Input("user")]
        public Input<string>? User { get; set; }

        public UploadFileArgs()
        {
        }
        public static new UploadFileArgs Empty => new UploadFileArgs();
    }
}
