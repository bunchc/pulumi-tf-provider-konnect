// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayPluginStatsdConfigQueueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time in seconds before the initial retry is made for a failing batch.
        /// </summary>
        [Input("initialRetryDelay")]
        public Input<double>? InitialRetryDelay { get; set; }

        /// <summary>
        /// Maximum number of entries that can be processed at a time.
        /// </summary>
        [Input("maxBatchSize")]
        public Input<int>? MaxBatchSize { get; set; }

        /// <summary>
        /// Maximum number of bytes that can be waiting on a queue, requires string content.
        /// </summary>
        [Input("maxBytes")]
        public Input<int>? MaxBytes { get; set; }

        /// <summary>
        /// Maximum number of (fractional) seconds to elapse after the first entry was queued before the queue starts calling the handler.
        /// </summary>
        [Input("maxCoalescingDelay")]
        public Input<double>? MaxCoalescingDelay { get; set; }

        /// <summary>
        /// Maximum number of entries that can be waiting on the queue.
        /// </summary>
        [Input("maxEntries")]
        public Input<int>? MaxEntries { get; set; }

        /// <summary>
        /// Maximum time in seconds between retries, caps exponential backoff.
        /// </summary>
        [Input("maxRetryDelay")]
        public Input<double>? MaxRetryDelay { get; set; }

        /// <summary>
        /// Time in seconds before the queue gives up calling a failed handler for a batch.
        /// </summary>
        [Input("maxRetryTime")]
        public Input<double>? MaxRetryTime { get; set; }

        public GatewayPluginStatsdConfigQueueArgs()
        {
        }
        public static new GatewayPluginStatsdConfigQueueArgs Empty => new GatewayPluginStatsdConfigQueueArgs();
    }
}
