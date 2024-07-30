// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayPluginRateLimitingAdvancedConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("consumerGroups")]
        private InputList<string>? _consumerGroups;

        /// <summary>
        /// List of consumer groups allowed to override the rate limiting settings for the given Route or Service. Required if `enforce_consumer_groups` is set to `true`.
        /// </summary>
        public InputList<string> ConsumerGroups
        {
            get => _consumerGroups ?? (_consumerGroups = new InputList<string>());
            set => _consumerGroups = value;
        }

        /// <summary>
        /// The shared dictionary where counters are stored. When the plugin is configured to synchronize counter data externally (that is `config.strategy` is `cluster` or `redis` and `config.sync_rate` isn't `-1`), this dictionary serves as a buffer to populate counters in the data store on each synchronization cycle.
        /// </summary>
        [Input("dictionaryName")]
        public Input<string>? DictionaryName { get; set; }

        /// <summary>
        /// If set to `true`, this doesn't count denied requests (status = `429`). If set to `false`, all requests, including denied ones, are counted. This parameter only affects the `sliding` window_type.
        /// </summary>
        [Input("disablePenalty")]
        public Input<bool>? DisablePenalty { get; set; }

        /// <summary>
        /// Determines if consumer groups are allowed to override the rate limiting settings for the given Route or Service. Flipping `enforce_consumer_groups` from `true` to `false` disables the group override, but does not clear the list of consumer groups. You can then flip `enforce_consumer_groups` to `true` to re-enforce the groups.
        /// </summary>
        [Input("enforceConsumerGroups")]
        public Input<bool>? EnforceConsumerGroups { get; set; }

        /// <summary>
        /// Set a custom error code to return when the rate limit is exceeded.
        /// </summary>
        [Input("errorCode")]
        public Input<double>? ErrorCode { get; set; }

        /// <summary>
        /// Set a custom error message to return when the rate limit is exceeded.
        /// </summary>
        [Input("errorMessage")]
        public Input<string>? ErrorMessage { get; set; }

        /// <summary>
        /// A string representing an HTTP header name.
        /// </summary>
        [Input("headerName")]
        public Input<string>? HeaderName { get; set; }

        /// <summary>
        /// Optionally hide informative response headers that would otherwise provide information about the current status of limits and counters.
        /// </summary>
        [Input("hideClientHeaders")]
        public Input<bool>? HideClientHeaders { get; set; }

        /// <summary>
        /// The type of identifier used to generate the rate limit key. Defines the scope used to increment the rate limiting counters. Can be `ip`, `credential`, `consumer`, `service`, `header`, `path` or `consumer-group`. must be one of ["ip", "credential", "consumer", "service", "header", "path", "consumer-group"]
        /// </summary>
        [Input("identifier")]
        public Input<string>? Identifier { get; set; }

        [Input("limits")]
        private InputList<double>? _limits;

        /// <summary>
        /// One or more requests-per-window limits to apply. There must be a matching number of window limits and sizes specified.
        /// </summary>
        public InputList<double> Limits
        {
            get => _limits ?? (_limits = new InputList<double>());
            set => _limits = value;
        }

        /// <summary>
        /// The rate limiting library namespace to use for this plugin instance. Counter data and sync configuration is isolated in each namespace. NOTE: For the plugin instances sharing the same namespace, all the configurations that are required for synchronizing counters, e.g. `strategy`, `redis`, `sync_rate`, `window_size`, `dictionary_name`, need to be the same.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// A string representing a URL path, such as /path/to/resource. Must start with a forward slash (/) and must not contain empty segments (i.e., two consecutive forward slashes).
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("redis")]
        public Input<Inputs.GatewayPluginRateLimitingAdvancedConfigRedisArgs>? Redis { get; set; }

        /// <summary>
        /// The upper bound of a jitter (random delay) in seconds to be added to the `Retry-After` header of denied requests (status = `429`) in order to prevent all the clients from coming back at the same time. The lower bound of the jitter is `0`; in this case, the `Retry-After` header is equal to the `RateLimit-Reset` header.
        /// </summary>
        [Input("retryAfterJitterMax")]
        public Input<double>? RetryAfterJitterMax { get; set; }

        /// <summary>
        /// The rate-limiting strategy to use for retrieving and incrementing the limits. Available values are: `local` and `cluster`. must be one of ["cluster", "redis", "local"]
        /// </summary>
        [Input("strategy")]
        public Input<string>? Strategy { get; set; }

        /// <summary>
        /// How often to sync counter data to the central data store. A value of 0 results in synchronous behavior; a value of -1 ignores sync behavior entirely and only stores counters in node memory. A value greater than 0 will sync the counters in the specified number of seconds. The minimum allowed interval is 0.02 seconds (20ms).
        /// </summary>
        [Input("syncRate")]
        public Input<double>? SyncRate { get; set; }

        [Input("windowSizes")]
        private InputList<double>? _windowSizes;

        /// <summary>
        /// One or more window sizes to apply a limit to (defined in seconds). There must be a matching number of window limits and sizes specified.
        /// </summary>
        public InputList<double> WindowSizes
        {
            get => _windowSizes ?? (_windowSizes = new InputList<double>());
            set => _windowSizes = value;
        }

        /// <summary>
        /// Sets the time window type to either `sliding` (default) or `fixed`. Sliding windows apply the rate limiting logic while taking into account previous hit rates (from the window that immediately precedes the current) using a dynamic weight. Fixed windows consist of buckets that are statically assigned to a definitive time range, each request is mapped to only one fixed window based on its timestamp and will affect only that window's counters. must be one of ["fixed", "sliding"]
        /// </summary>
        [Input("windowType")]
        public Input<string>? WindowType { get; set; }

        public GatewayPluginRateLimitingAdvancedConfigArgs()
        {
        }
        public static new GatewayPluginRateLimitingAdvancedConfigArgs Empty => new GatewayPluginRateLimitingAdvancedConfigArgs();
    }
}