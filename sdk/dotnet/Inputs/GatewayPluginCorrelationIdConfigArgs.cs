// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayPluginCorrelationIdConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to echo the header back to downstream (the client).
        /// </summary>
        [Input("echoDownstream")]
        public Input<bool>? EchoDownstream { get; set; }

        /// <summary>
        /// The generator to use for the correlation ID. Accepted values are `uuid`, `uuid#counter`, and `tracker`. See Generators. must be one of ["uuid", "uuid#counter", "tracker"]
        /// </summary>
        [Input("generator")]
        public Input<string>? Generator { get; set; }

        /// <summary>
        /// The HTTP header name to use for the correlation ID.
        /// </summary>
        [Input("headerName")]
        public Input<string>? HeaderName { get; set; }

        public GatewayPluginCorrelationIdConfigArgs()
        {
        }
        public static new GatewayPluginCorrelationIdConfigArgs Empty => new GatewayPluginCorrelationIdConfigArgs();
    }
}