// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayPluginJqConfigRequestJqProgramOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("asciiOutput")]
        public Input<bool>? AsciiOutput { get; set; }

        [Input("compactOutput")]
        public Input<bool>? CompactOutput { get; set; }

        [Input("joinOutput")]
        public Input<bool>? JoinOutput { get; set; }

        [Input("rawOutput")]
        public Input<bool>? RawOutput { get; set; }

        [Input("sortKeys")]
        public Input<bool>? SortKeys { get; set; }

        public GatewayPluginJqConfigRequestJqProgramOptionsArgs()
        {
        }
        public static new GatewayPluginJqConfigRequestJqProgramOptionsArgs Empty => new GatewayPluginJqConfigRequestJqProgramOptionsArgs();
    }
}