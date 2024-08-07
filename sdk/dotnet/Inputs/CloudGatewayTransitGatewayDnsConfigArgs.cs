// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class CloudGatewayTransitGatewayDnsConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("domainProxyLists")]
        private InputList<string>? _domainProxyLists;

        /// <summary>
        /// Internal domain names to proxy for DNS resolution from the listed remote DNS server IP addresses,
        /// for a transit gateway.
        /// </summary>
        public InputList<string> DomainProxyLists
        {
            get => _domainProxyLists ?? (_domainProxyLists = new InputList<string>());
            set => _domainProxyLists = value;
        }

        [Input("remoteDnsServerIpAddresses")]
        private InputList<string>? _remoteDnsServerIpAddresses;

        /// <summary>
        /// Remote DNS Server IP Addresses to connect to for resolving internal DNS via a transit gateway. Requires replacement if changed. ; Not Null
        /// </summary>
        public InputList<string> RemoteDnsServerIpAddresses
        {
            get => _remoteDnsServerIpAddresses ?? (_remoteDnsServerIpAddresses = new InputList<string>());
            set => _remoteDnsServerIpAddresses = value;
        }

        public CloudGatewayTransitGatewayDnsConfigArgs()
        {
        }
        public static new CloudGatewayTransitGatewayDnsConfigArgs Empty => new CloudGatewayTransitGatewayDnsConfigArgs();
    }
}
