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
    public sealed class GetGatewayPluginRequestTransformerConfigAppendResult
    {
        public readonly ImmutableArray<string> Bodies;
        public readonly ImmutableArray<string> Headers;
        public readonly ImmutableArray<string> Querystrings;

        [OutputConstructor]
        private GetGatewayPluginRequestTransformerConfigAppendResult(
            ImmutableArray<string> bodies,

            ImmutableArray<string> headers,

            ImmutableArray<string> querystrings)
        {
            Bodies = bodies;
            Headers = headers;
            Querystrings = querystrings;
        }
    }
}
