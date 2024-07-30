// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class CloudGatewayCustomDomainStateMetadataArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Reason why the custom domain may be in an erroneous state, reported from backing infrastructure.
        /// </summary>
        [Input("reason")]
        public Input<string>? Reason { get; set; }

        /// <summary>
        /// Reported status of the custom domain from backing infrastructure.
        /// </summary>
        [Input("reportedStatus")]
        public Input<string>? ReportedStatus { get; set; }

        public CloudGatewayCustomDomainStateMetadataArgs()
        {
        }
        public static new CloudGatewayCustomDomainStateMetadataArgs Empty => new CloudGatewayCustomDomainStateMetadataArgs();
    }
}
