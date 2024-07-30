// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class ServerlessCloudGatewayControlPlaneArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the serverless cloud gateway CP. Requires replacement if changed.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The prefix of the serverless cloud gateway CP. Requires replacement if changed.
        /// </summary>
        [Input("prefix", required: true)]
        public Input<string> Prefix { get; set; } = null!;

        /// <summary>
        /// The control plane region. Requires replacement if changed. ; must be one of ["us", "eu", "au"]
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public ServerlessCloudGatewayControlPlaneArgs()
        {
        }
        public static new ServerlessCloudGatewayControlPlaneArgs Empty => new ServerlessCloudGatewayControlPlaneArgs();
    }
}