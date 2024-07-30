// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class CloudGatewayNetworkProviderMetadataGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("subnetIds")]
        private InputList<string>? _subnetIds;
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public CloudGatewayNetworkProviderMetadataGetArgs()
        {
        }
        public static new CloudGatewayNetworkProviderMetadataGetArgs Empty => new CloudGatewayNetworkProviderMetadataGetArgs();
    }
}