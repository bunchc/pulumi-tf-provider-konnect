// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayControlPlaneProxyUrlArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Hostname of the proxy URL.
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        /// <summary>
        /// Port of the proxy URL.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// Protocol of the proxy URL.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        public GatewayControlPlaneProxyUrlArgs()
        {
        }
        public static new GatewayControlPlaneProxyUrlArgs Empty => new GatewayControlPlaneProxyUrlArgs();
    }
}
