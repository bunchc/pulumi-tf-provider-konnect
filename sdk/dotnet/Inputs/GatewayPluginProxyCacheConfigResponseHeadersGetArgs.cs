// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayPluginProxyCacheConfigResponseHeadersGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("age")]
        public Input<bool>? Age { get; set; }

        [Input("xCacheKey")]
        public Input<bool>? XCacheKey { get; set; }

        [Input("xCacheStatus")]
        public Input<bool>? XCacheStatus { get; set; }

        public GatewayPluginProxyCacheConfigResponseHeadersGetArgs()
        {
        }
        public static new GatewayPluginProxyCacheConfigResponseHeadersGetArgs Empty => new GatewayPluginProxyCacheConfigResponseHeadersGetArgs();
    }
}
