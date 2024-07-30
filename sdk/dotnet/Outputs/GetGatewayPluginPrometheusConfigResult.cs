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
    public sealed class GetGatewayPluginPrometheusConfigResult
    {
        /// <summary>
        /// A boolean value that determines if bandwidth metrics should be collected. If enabled, `bandwidth_bytes` and `stream_sessions_total` metrics will be exported.
        /// </summary>
        public readonly bool BandwidthMetrics;
        /// <summary>
        /// A boolean value that determines if latency metrics should be collected. If enabled, `kong_latency_ms`, `upstream_latency_ms` and `request_latency_ms` metrics will be exported.
        /// </summary>
        public readonly bool LatencyMetrics;
        /// <summary>
        /// A boolean value that determines if per-consumer metrics should be collected. If enabled, the `kong_http_requests_total` and `kong_bandwidth_bytes` metrics fill in the consumer label when available.
        /// </summary>
        public readonly bool PerConsumer;
        /// <summary>
        /// A boolean value that determines if status code metrics should be collected. If enabled, `http_requests_total`, `stream_sessions_total` metrics will be exported.
        /// </summary>
        public readonly bool StatusCodeMetrics;
        /// <summary>
        /// A boolean value that determines if upstream metrics should be collected. If enabled, `upstream_target_health` metric will be exported.
        /// </summary>
        public readonly bool UpstreamHealthMetrics;

        [OutputConstructor]
        private GetGatewayPluginPrometheusConfigResult(
            bool bandwidthMetrics,

            bool latencyMetrics,

            bool perConsumer,

            bool statusCodeMetrics,

            bool upstreamHealthMetrics)
        {
            BandwidthMetrics = bandwidthMetrics;
            LatencyMetrics = latencyMetrics;
            PerConsumer = perConsumer;
            StatusCodeMetrics = statusCodeMetrics;
            UpstreamHealthMetrics = upstreamHealthMetrics;
        }
    }
}
