// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class CloudGatewayConfigurationDataplaneGroupConfigAutoscaleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Object that describes the autopilot autoscaling strategy.
        /// </summary>
        [Input("configurationDataPlaneGroupAutoscaleAutopilot")]
        public Input<Inputs.CloudGatewayConfigurationDataplaneGroupConfigAutoscaleConfigurationDataPlaneGroupAutoscaleAutopilotGetArgs>? ConfigurationDataPlaneGroupAutoscaleAutopilot { get; set; }

        /// <summary>
        /// Object that describes the static autoscaling strategy.
        /// </summary>
        [Input("configurationDataPlaneGroupAutoscaleStatic")]
        public Input<Inputs.CloudGatewayConfigurationDataplaneGroupConfigAutoscaleConfigurationDataPlaneGroupAutoscaleStaticGetArgs>? ConfigurationDataPlaneGroupAutoscaleStatic { get; set; }

        public CloudGatewayConfigurationDataplaneGroupConfigAutoscaleGetArgs()
        {
        }
        public static new CloudGatewayConfigurationDataplaneGroupConfigAutoscaleGetArgs Empty => new CloudGatewayConfigurationDataplaneGroupConfigAutoscaleGetArgs();
    }
}
