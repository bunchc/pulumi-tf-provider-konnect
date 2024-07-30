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
    public sealed class GetGatewayControlPlaneConfigResult
    {
        /// <summary>
        /// Control Plane Endpoint.
        /// </summary>
        public readonly string ControlPlaneEndpoint;
        /// <summary>
        /// Telemetry Endpoint.
        /// </summary>
        public readonly string TelemetryEndpoint;

        [OutputConstructor]
        private GetGatewayControlPlaneConfigResult(
            string controlPlaneEndpoint,

            string telemetryEndpoint)
        {
            ControlPlaneEndpoint = controlPlaneEndpoint;
            TelemetryEndpoint = telemetryEndpoint;
        }
    }
}
