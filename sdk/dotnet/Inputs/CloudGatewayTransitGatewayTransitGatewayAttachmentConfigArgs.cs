// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class CloudGatewayTransitGatewayTransitGatewayAttachmentConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Requires replacement if changed.
        /// </summary>
        [Input("awsTransitGatewayAttachmentConfig")]
        public Input<Inputs.CloudGatewayTransitGatewayTransitGatewayAttachmentConfigAwsTransitGatewayAttachmentConfigArgs>? AwsTransitGatewayAttachmentConfig { get; set; }

        public CloudGatewayTransitGatewayTransitGatewayAttachmentConfigArgs()
        {
        }
        public static new CloudGatewayTransitGatewayTransitGatewayAttachmentConfigArgs Empty => new CloudGatewayTransitGatewayTransitGatewayAttachmentConfigArgs();
    }
}
