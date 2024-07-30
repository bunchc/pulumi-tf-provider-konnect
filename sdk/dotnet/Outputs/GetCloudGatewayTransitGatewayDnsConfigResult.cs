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
    public sealed class GetCloudGatewayTransitGatewayDnsConfigResult
    {
        /// <summary>
        /// Internal domain names to proxy for DNS resolution from the listed remote DNS server IP addresses,
        /// for a transit gateway.
        /// </summary>
        public readonly ImmutableArray<string> DomainProxyLists;
        /// <summary>
        /// Remote DNS Server IP Addresses to connect to for resolving internal DNS via a transit gateway.
        /// </summary>
        public readonly ImmutableArray<string> RemoteDnsServerIpAddresses;

        [OutputConstructor]
        private GetCloudGatewayTransitGatewayDnsConfigResult(
            ImmutableArray<string> domainProxyLists,

            ImmutableArray<string> remoteDnsServerIpAddresses)
        {
            DomainProxyLists = domainProxyLists;
            RemoteDnsServerIpAddresses = remoteDnsServerIpAddresses;
        }
    }
}