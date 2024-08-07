// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayPluginOpentelemetryConfigGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The delay, in seconds, between two consecutive batches.
        /// </summary>
        [Input("batchFlushDelay")]
        public Input<int>? BatchFlushDelay { get; set; }

        /// <summary>
        /// The number of spans to be sent in a single batch.
        /// </summary>
        [Input("batchSpanCount")]
        public Input<int>? BatchSpanCount { get; set; }

        /// <summary>
        /// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
        /// </summary>
        [Input("connectTimeout")]
        public Input<int>? ConnectTimeout { get; set; }

        /// <summary>
        /// A string representing a URL, such as https://example.com/path/to/resource?q=search.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// must be one of ["preserve", "ignore", "b3", "b3-single", "w3c", "jaeger", "ot", "aws", "gcp", "datadog"]
        /// </summary>
        [Input("headerType")]
        public Input<string>? HeaderType { get; set; }

        [Input("headers")]
        private InputMap<string>? _headers;

        /// <summary>
        /// The custom headers to be added in the HTTP request sent to the OTLP server. This setting is useful for adding the authentication headers (token) for the APM backend.
        /// </summary>
        public InputMap<string> Headers
        {
            get => _headers ?? (_headers = new InputMap<string>());
            set => _headers = value;
        }

        [Input("httpResponseHeaderForTraceid")]
        public Input<string>? HttpResponseHeaderForTraceid { get; set; }

        [Input("propagation")]
        public Input<Inputs.GatewayPluginOpentelemetryConfigPropagationGetArgs>? Propagation { get; set; }

        [Input("queue")]
        public Input<Inputs.GatewayPluginOpentelemetryConfigQueueGetArgs>? Queue { get; set; }

        /// <summary>
        /// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
        /// </summary>
        [Input("readTimeout")]
        public Input<int>? ReadTimeout { get; set; }

        [Input("resourceAttributes")]
        private InputMap<string>? _resourceAttributes;
        public InputMap<string> ResourceAttributes
        {
            get => _resourceAttributes ?? (_resourceAttributes = new InputMap<string>());
            set => _resourceAttributes = value;
        }

        /// <summary>
        /// Tracing sampling rate for configuring the probability-based sampler. When set, this value supersedes the global `tracing_sampling_rate` setting from kong.conf.
        /// </summary>
        [Input("samplingRate")]
        public Input<double>? SamplingRate { get; set; }

        /// <summary>
        /// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
        /// </summary>
        [Input("sendTimeout")]
        public Input<int>? SendTimeout { get; set; }

        public GatewayPluginOpentelemetryConfigGetArgs()
        {
        }
        public static new GatewayPluginOpentelemetryConfigGetArgs Empty => new GatewayPluginOpentelemetryConfigGetArgs();
    }
}
