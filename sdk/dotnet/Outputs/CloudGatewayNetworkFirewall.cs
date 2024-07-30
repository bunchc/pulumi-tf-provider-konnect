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
    public sealed class CloudGatewayNetworkFirewall
    {
        /// <summary>
        /// List of allowed CIDR blocks to access a network.
        /// </summary>
        public readonly ImmutableArray<string> AllowedCidrBlocks;
        /// <summary>
        /// List of denied CIDR blocks to access a network.
        /// </summary>
        public readonly ImmutableArray<string> DeniedCidrBlocks;

        [OutputConstructor]
        private CloudGatewayNetworkFirewall(
            ImmutableArray<string> allowedCidrBlocks,

            ImmutableArray<string> deniedCidrBlocks)
        {
            AllowedCidrBlocks = allowedCidrBlocks;
            DeniedCidrBlocks = deniedCidrBlocks;
        }
    }
}
