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
    public sealed class GatewayPluginSamlConfigSessionRedisClusterNode
    {
        /// <summary>
        /// A string representing a host name, such as example.com.
        /// </summary>
        public readonly string? Ip;
        /// <summary>
        /// An integer representing a port number between 0 and 65535, inclusive.
        /// </summary>
        public readonly int? Port;

        [OutputConstructor]
        private GatewayPluginSamlConfigSessionRedisClusterNode(
            string? ip,

            int? port)
        {
            Ip = ip;
            Port = port;
        }
    }
}
