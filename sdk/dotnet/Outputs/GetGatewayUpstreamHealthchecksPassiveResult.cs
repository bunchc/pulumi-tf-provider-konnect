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
    public sealed class GetGatewayUpstreamHealthchecksPassiveResult
    {
        public readonly Outputs.GetGatewayUpstreamHealthchecksPassiveHealthyResult Healthy;
        /// <summary>
        /// must be one of ["tcp", "http", "https", "grpc", "grpcs"]
        /// </summary>
        public readonly string Type;
        public readonly Outputs.GetGatewayUpstreamHealthchecksPassiveUnhealthyResult Unhealthy;

        [OutputConstructor]
        private GetGatewayUpstreamHealthchecksPassiveResult(
            Outputs.GetGatewayUpstreamHealthchecksPassiveHealthyResult healthy,

            string type,

            Outputs.GetGatewayUpstreamHealthchecksPassiveUnhealthyResult unhealthy)
        {
            Healthy = healthy;
            Type = type;
            Unhealthy = unhealthy;
        }
    }
}