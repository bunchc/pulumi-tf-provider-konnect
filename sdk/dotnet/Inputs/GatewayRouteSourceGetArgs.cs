// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayRouteSourceGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        public GatewayRouteSourceGetArgs()
        {
        }
        public static new GatewayRouteSourceGetArgs Empty => new GatewayRouteSourceGetArgs();
    }
}