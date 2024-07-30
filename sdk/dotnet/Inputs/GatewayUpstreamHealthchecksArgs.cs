// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayUpstreamHealthchecksArgs : global::Pulumi.ResourceArgs
    {
        [Input("active")]
        public Input<Inputs.GatewayUpstreamHealthchecksActiveArgs>? Active { get; set; }

        [Input("passive")]
        public Input<Inputs.GatewayUpstreamHealthchecksPassiveArgs>? Passive { get; set; }

        [Input("threshold")]
        public Input<double>? Threshold { get; set; }

        public GatewayUpstreamHealthchecksArgs()
        {
        }
        public static new GatewayUpstreamHealthchecksArgs Empty => new GatewayUpstreamHealthchecksArgs();
    }
}