// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayUpstreamHealthchecksActiveUnhealthyArgs : global::Pulumi.ResourceArgs
    {
        [Input("httpFailures")]
        public Input<int>? HttpFailures { get; set; }

        [Input("httpStatuses")]
        private InputList<int>? _httpStatuses;
        public InputList<int> HttpStatuses
        {
            get => _httpStatuses ?? (_httpStatuses = new InputList<int>());
            set => _httpStatuses = value;
        }

        [Input("interval")]
        public Input<double>? Interval { get; set; }

        [Input("tcpFailures")]
        public Input<int>? TcpFailures { get; set; }

        [Input("timeouts")]
        public Input<int>? Timeouts { get; set; }

        public GatewayUpstreamHealthchecksActiveUnhealthyArgs()
        {
        }
        public static new GatewayUpstreamHealthchecksActiveUnhealthyArgs Empty => new GatewayUpstreamHealthchecksActiveUnhealthyArgs();
    }
}
