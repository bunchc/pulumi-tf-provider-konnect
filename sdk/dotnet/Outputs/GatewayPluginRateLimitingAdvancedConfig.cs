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
    public sealed class GatewayPluginRateLimitingAdvancedConfig
    {
        /// <summary>
        /// List of consumer groups allowed to override the rate limiting settings for the given Route or Service. Required if `enforce_consumer_groups` is set to `true`.
        /// </summary>
        public readonly ImmutableArray<string> ConsumerGroups;
        /// <summary>
        /// The shared dictionary where counters are stored. When the plugin is configured to synchronize counter data externally (that is `config.strategy` is `cluster` or `redis` and `config.sync_rate` isn't `-1`), this dictionary serves as a buffer to populate counters in the data store on each synchronization cycle.
        /// </summary>
        public readonly string? DictionaryName;
        /// <summary>
        /// If set to `true`, this doesn't count denied requests (status = `429`). If set to `false`, all requests, including denied ones, are counted. This parameter only affects the `sliding` window_type.
        /// </summary>
        public readonly bool? DisablePenalty;
        /// <summary>
        /// Determines if consumer groups are allowed to override the rate limiting settings for the given Route or Service. Flipping `enforce_consumer_groups` from `true` to `false` disables the group override, but does not clear the list of consumer groups. You can then flip `enforce_consumer_groups` to `true` to re-enforce the groups.
        /// </summary>
        public readonly bool? EnforceConsumerGroups;
        /// <summary>
        /// Set a custom error code to return when the rate limit is exceeded.
        /// </summary>
        public readonly double? ErrorCode;
        /// <summary>
        /// Set a custom error message to return when the rate limit is exceeded.
        /// </summary>
        public readonly string? ErrorMessage;
        /// <summary>
        /// A string representing an HTTP header name.
        /// </summary>
        public readonly string? HeaderName;
        /// <summary>
        /// Optionally hide informative response headers that would otherwise provide information about the current status of limits and counters.
        /// </summary>
        public readonly bool? HideClientHeaders;
        /// <summary>
        /// The type of identifier used to generate the rate limit key. Defines the scope used to increment the rate limiting counters. Can be `ip`, `credential`, `consumer`, `service`, `header`, `path` or `consumer-group`. must be one of ["ip", "credential", "consumer", "service", "header", "path", "consumer-group"]
        /// </summary>
        public readonly string? Identifier;
        /// <summary>
        /// One or more requests-per-window limits to apply. There must be a matching number of window limits and sizes specified.
        /// </summary>
        public readonly ImmutableArray<double> Limits;
        /// <summary>
        /// The rate limiting library namespace to use for this plugin instance. Counter data and sync configuration is isolated in each namespace. NOTE: For the plugin instances sharing the same namespace, all the configurations that are required for synchronizing counters, e.g. `strategy`, `redis`, `sync_rate`, `window_size`, `dictionary_name`, need to be the same.
        /// </summary>
        public readonly string? Namespace;
        /// <summary>
        /// A string representing a URL path, such as /path/to/resource. Must start with a forward slash (/) and must not contain empty segments (i.e., two consecutive forward slashes).
        /// </summary>
        public readonly string? Path;
        public readonly Outputs.GatewayPluginRateLimitingAdvancedConfigRedis? Redis;
        /// <summary>
        /// The upper bound of a jitter (random delay) in seconds to be added to the `Retry-After` header of denied requests (status = `429`) in order to prevent all the clients from coming back at the same time. The lower bound of the jitter is `0`; in this case, the `Retry-After` header is equal to the `RateLimit-Reset` header.
        /// </summary>
        public readonly double? RetryAfterJitterMax;
        /// <summary>
        /// The rate-limiting strategy to use for retrieving and incrementing the limits. Available values are: `local` and `cluster`. must be one of ["cluster", "redis", "local"]
        /// </summary>
        public readonly string? Strategy;
        /// <summary>
        /// How often to sync counter data to the central data store. A value of 0 results in synchronous behavior; a value of -1 ignores sync behavior entirely and only stores counters in node memory. A value greater than 0 will sync the counters in the specified number of seconds. The minimum allowed interval is 0.02 seconds (20ms).
        /// </summary>
        public readonly double? SyncRate;
        /// <summary>
        /// One or more window sizes to apply a limit to (defined in seconds). There must be a matching number of window limits and sizes specified.
        /// </summary>
        public readonly ImmutableArray<double> WindowSizes;
        /// <summary>
        /// Sets the time window type to either `sliding` (default) or `fixed`. Sliding windows apply the rate limiting logic while taking into account previous hit rates (from the window that immediately precedes the current) using a dynamic weight. Fixed windows consist of buckets that are statically assigned to a definitive time range, each request is mapped to only one fixed window based on its timestamp and will affect only that window's counters. must be one of ["fixed", "sliding"]
        /// </summary>
        public readonly string? WindowType;

        [OutputConstructor]
        private GatewayPluginRateLimitingAdvancedConfig(
            ImmutableArray<string> consumerGroups,

            string? dictionaryName,

            bool? disablePenalty,

            bool? enforceConsumerGroups,

            double? errorCode,

            string? errorMessage,

            string? headerName,

            bool? hideClientHeaders,

            string? identifier,

            ImmutableArray<double> limits,

            string? @namespace,

            string? path,

            Outputs.GatewayPluginRateLimitingAdvancedConfigRedis? redis,

            double? retryAfterJitterMax,

            string? strategy,

            double? syncRate,

            ImmutableArray<double> windowSizes,

            string? windowType)
        {
            ConsumerGroups = consumerGroups;
            DictionaryName = dictionaryName;
            DisablePenalty = disablePenalty;
            EnforceConsumerGroups = enforceConsumerGroups;
            ErrorCode = errorCode;
            ErrorMessage = errorMessage;
            HeaderName = headerName;
            HideClientHeaders = hideClientHeaders;
            Identifier = identifier;
            Limits = limits;
            Namespace = @namespace;
            Path = path;
            Redis = redis;
            RetryAfterJitterMax = retryAfterJitterMax;
            Strategy = strategy;
            SyncRate = syncRate;
            WindowSizes = windowSizes;
            WindowType = windowType;
        }
    }
}
