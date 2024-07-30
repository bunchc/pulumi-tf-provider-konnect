// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayPluginStatsdConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowStatusCodes")]
        private InputList<string>? _allowStatusCodes;

        /// <summary>
        /// List of status code ranges that are allowed to be logged in metrics.
        /// </summary>
        public InputList<string> AllowStatusCodes
        {
            get => _allowStatusCodes ?? (_allowStatusCodes = new InputList<string>());
            set => _allowStatusCodes = value;
        }

        /// <summary>
        /// must be one of ["consumer*id", "custom*id", "username"]
        /// </summary>
        [Input("consumerIdentifierDefault")]
        public Input<string>? ConsumerIdentifierDefault { get; set; }

        [Input("flushTimeout")]
        public Input<double>? FlushTimeout { get; set; }

        /// <summary>
        /// The IP address or hostname of StatsD server to send data to.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        [Input("hostnameInPrefix")]
        public Input<bool>? HostnameInPrefix { get; set; }

        [Input("metrics")]
        private InputList<Inputs.GatewayPluginStatsdConfigMetricArgs>? _metrics;

        /// <summary>
        /// List of metrics to be logged.
        /// </summary>
        public InputList<Inputs.GatewayPluginStatsdConfigMetricArgs> Metrics
        {
            get => _metrics ?? (_metrics = new InputList<Inputs.GatewayPluginStatsdConfigMetricArgs>());
            set => _metrics = value;
        }

        /// <summary>
        /// The port of StatsD server to send data to.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// String to prefix to each metric's name.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        [Input("queue")]
        public Input<Inputs.GatewayPluginStatsdConfigQueueArgs>? Queue { get; set; }

        [Input("queueSize")]
        public Input<int>? QueueSize { get; set; }

        [Input("retryCount")]
        public Input<int>? RetryCount { get; set; }

        /// <summary>
        /// must be one of ["service*id", "service*name", "service*host", "service*name*or*host"]
        /// </summary>
        [Input("serviceIdentifierDefault")]
        public Input<string>? ServiceIdentifierDefault { get; set; }

        /// <summary>
        /// must be one of ["dogstatsd", "influxdb", "librato", "signalfx"]
        /// </summary>
        [Input("tagStyle")]
        public Input<string>? TagStyle { get; set; }

        [Input("udpPacketSize")]
        public Input<double>? UdpPacketSize { get; set; }

        [Input("useTcp")]
        public Input<bool>? UseTcp { get; set; }

        /// <summary>
        /// must be one of ["workspace*id", "workspace*name"]
        /// </summary>
        [Input("workspaceIdentifierDefault")]
        public Input<string>? WorkspaceIdentifierDefault { get; set; }

        public GatewayPluginStatsdConfigArgs()
        {
        }
        public static new GatewayPluginStatsdConfigArgs Empty => new GatewayPluginStatsdConfigArgs();
    }
}
