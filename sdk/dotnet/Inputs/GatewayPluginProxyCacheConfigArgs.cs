// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayPluginProxyCacheConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When enabled, respect the Cache-Control behaviors defined in RFC7234.
        /// </summary>
        [Input("cacheControl")]
        public Input<bool>? CacheControl { get; set; }

        /// <summary>
        /// TTL, in seconds, of cache entities.
        /// </summary>
        [Input("cacheTtl")]
        public Input<int>? CacheTtl { get; set; }

        [Input("contentTypes")]
        private InputList<string>? _contentTypes;

        /// <summary>
        /// Upstream response content types considered cacheable. The plugin performs an **exact match** against each specified value.
        /// </summary>
        public InputList<string> ContentTypes
        {
            get => _contentTypes ?? (_contentTypes = new InputList<string>());
            set => _contentTypes = value;
        }

        [Input("ignoreUriCase")]
        public Input<bool>? IgnoreUriCase { get; set; }

        [Input("memory")]
        public Input<Inputs.GatewayPluginProxyCacheConfigMemoryArgs>? Memory { get; set; }

        [Input("requestMethods")]
        private InputList<string>? _requestMethods;

        /// <summary>
        /// Downstream request methods considered cacheable.
        /// </summary>
        public InputList<string> RequestMethods
        {
            get => _requestMethods ?? (_requestMethods = new InputList<string>());
            set => _requestMethods = value;
        }

        [Input("responseCodes")]
        private InputList<int>? _responseCodes;

        /// <summary>
        /// Upstream response status code considered cacheable.
        /// </summary>
        public InputList<int> ResponseCodes
        {
            get => _responseCodes ?? (_responseCodes = new InputList<int>());
            set => _responseCodes = value;
        }

        /// <summary>
        /// Caching related diagnostic headers that should be included in cached responses
        /// </summary>
        [Input("responseHeaders")]
        public Input<Inputs.GatewayPluginProxyCacheConfigResponseHeadersArgs>? ResponseHeaders { get; set; }

        /// <summary>
        /// Number of seconds to keep resources in the storage backend. This value is independent of `cache_ttl` or resource TTLs defined by Cache-Control behaviors.
        /// </summary>
        [Input("storageTtl")]
        public Input<int>? StorageTtl { get; set; }

        /// <summary>
        /// The backing data store in which to hold cache entities. must be one of ["memory"]
        /// </summary>
        [Input("strategy")]
        public Input<string>? Strategy { get; set; }

        [Input("varyHeaders")]
        private InputList<string>? _varyHeaders;

        /// <summary>
        /// Relevant headers considered for the cache key. If undefined, none of the headers are taken into consideration.
        /// </summary>
        public InputList<string> VaryHeaders
        {
            get => _varyHeaders ?? (_varyHeaders = new InputList<string>());
            set => _varyHeaders = value;
        }

        [Input("varyQueryParams")]
        private InputList<string>? _varyQueryParams;

        /// <summary>
        /// Relevant query parameters considered for the cache key. If undefined, all params are taken into consideration.
        /// </summary>
        public InputList<string> VaryQueryParams
        {
            get => _varyQueryParams ?? (_varyQueryParams = new InputList<string>());
            set => _varyQueryParams = value;
        }

        public GatewayPluginProxyCacheConfigArgs()
        {
        }
        public static new GatewayPluginProxyCacheConfigArgs Empty => new GatewayPluginProxyCacheConfigArgs();
    }
}
