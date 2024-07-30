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
    public sealed class GetGatewayPluginOpentelemetryConfigPropagationResult
    {
        /// <summary>
        /// Header names to clear after context extraction. This allows to extract the context from a certain header and then remove it from the request, useful when extraction and injection are performed on different header formats and the original header should not be sent to the upstream. If left empty, no headers are cleared.
        /// </summary>
        public readonly ImmutableArray<string> Clears;
        /// <summary>
        /// The default header format to use when extractors did not match any format in the incoming headers and `inject` is configured with the value: `preserve`. This can happen when no tracing header was found in the request, or the incoming tracing header formats were not included in `extract`. must be one of ["b3", "gcp", "b3-single", "jaeger", "aws", "ot", "w3c", "datadog"]
        /// </summary>
        public readonly string DefaultFormat;
        /// <summary>
        /// Header formats used to extract tracing context from incoming requests. If multiple values are specified, the first one found will be used for extraction. If left empty, Kong will not extract any tracing context information from incoming requests and generate a trace with no parent and a new trace ID.
        /// </summary>
        public readonly ImmutableArray<string> Extracts;
        /// <summary>
        /// Header formats used to inject tracing context. The value `preserve` will use the same header format as the incoming request. If multiple values are specified, all of them will be used during injection. If left empty, Kong will not inject any tracing context information in outgoing requests.
        /// </summary>
        public readonly ImmutableArray<string> Injects;

        [OutputConstructor]
        private GetGatewayPluginOpentelemetryConfigPropagationResult(
            ImmutableArray<string> clears,

            string defaultFormat,

            ImmutableArray<string> extracts,

            ImmutableArray<string> injects)
        {
            Clears = clears;
            DefaultFormat = defaultFormat;
            Extracts = extracts;
            Injects = injects;
        }
    }
}
