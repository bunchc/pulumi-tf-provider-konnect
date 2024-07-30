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
    public sealed class GetGatewayPluginPreFunctionConfigResult
    {
        public readonly ImmutableArray<string> Accesses;
        public readonly ImmutableArray<string> BodyFilters;
        public readonly ImmutableArray<string> Certificates;
        public readonly ImmutableArray<string> HeaderFilters;
        public readonly ImmutableArray<string> Logs;
        public readonly ImmutableArray<string> Rewrites;
        public readonly ImmutableArray<string> WsClientFrames;
        public readonly ImmutableArray<string> WsCloses;
        public readonly ImmutableArray<string> WsHandshakes;
        public readonly ImmutableArray<string> WsUpstreamFrames;

        [OutputConstructor]
        private GetGatewayPluginPreFunctionConfigResult(
            ImmutableArray<string> accesses,

            ImmutableArray<string> bodyFilters,

            ImmutableArray<string> certificates,

            ImmutableArray<string> headerFilters,

            ImmutableArray<string> logs,

            ImmutableArray<string> rewrites,

            ImmutableArray<string> wsClientFrames,

            ImmutableArray<string> wsCloses,

            ImmutableArray<string> wsHandshakes,

            ImmutableArray<string> wsUpstreamFrames)
        {
            Accesses = accesses;
            BodyFilters = bodyFilters;
            Certificates = certificates;
            HeaderFilters = headerFilters;
            Logs = logs;
            Rewrites = rewrites;
            WsClientFrames = wsClientFrames;
            WsCloses = wsCloses;
            WsHandshakes = wsHandshakes;
            WsUpstreamFrames = wsUpstreamFrames;
        }
    }
}
