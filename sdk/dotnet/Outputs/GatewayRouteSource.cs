// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Outputs
{

    [OutputType]
    public sealed class GatewayRouteSource
    {
        public readonly string? Ip;
        public readonly int? Port;

        [OutputConstructor]
        private GatewayRouteSource(
            string? ip,

            int? port)
        {
            Ip = ip;
            Port = port;
        }
    }
}
