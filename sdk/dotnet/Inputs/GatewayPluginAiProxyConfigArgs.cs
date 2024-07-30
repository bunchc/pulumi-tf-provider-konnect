// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayPluginAiProxyConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("auth")]
        public Input<Inputs.GatewayPluginAiProxyConfigAuthArgs>? Auth { get; set; }

        [Input("logging")]
        public Input<Inputs.GatewayPluginAiProxyConfigLoggingArgs>? Logging { get; set; }

        [Input("model")]
        public Input<Inputs.GatewayPluginAiProxyConfigModelArgs>? Model { get; set; }

        /// <summary>
        /// Whether to 'optionally allow', 'deny', or 'always' (force) the streaming of answers via server sent events. must be one of ["allow", "deny", "always"]
        /// </summary>
        [Input("responseStreaming")]
        public Input<string>? ResponseStreaming { get; set; }

        /// <summary>
        /// The model's operation implementation, for this provider. Set to `preserve` to pass through without transformation. must be one of ["llm/v1/chat", "llm/v1/completions", "preserve"]
        /// </summary>
        [Input("routeType")]
        public Input<string>? RouteType { get; set; }

        public GatewayPluginAiProxyConfigArgs()
        {
        }
        public static new GatewayPluginAiProxyConfigArgs Empty => new GatewayPluginAiProxyConfigArgs();
    }
}
