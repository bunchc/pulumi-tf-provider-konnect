// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayUpstreamHealthchecksActiveHealthyArgs : global::Pulumi.ResourceArgs
    {
        [Input("httpStatuses")]
        private InputList<int>? _httpStatuses;
        public InputList<int> HttpStatuses
        {
            get => _httpStatuses ?? (_httpStatuses = new InputList<int>());
            set => _httpStatuses = value;
        }

        [Input("interval")]
        public Input<double>? Interval { get; set; }

        [Input("successes")]
        public Input<int>? Successes { get; set; }

        public GatewayUpstreamHealthchecksActiveHealthyArgs()
        {
        }
        public static new GatewayUpstreamHealthchecksActiveHealthyArgs Empty => new GatewayUpstreamHealthchecksActiveHealthyArgs();
    }
}
