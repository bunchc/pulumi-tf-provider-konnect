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
    public sealed class GetGatewayPluginResponseTransformerConfigAppendResult
    {
        public readonly ImmutableArray<string> Headers;
        /// <summary>
        /// List of JSON type names. Specify the types of the JSON values returned when appending
        /// JSON properties. Each string element can be one of: boolean, number, or string.
        /// </summary>
        public readonly ImmutableArray<string> JsonTypes;
        public readonly ImmutableArray<string> Jsons;

        [OutputConstructor]
        private GetGatewayPluginResponseTransformerConfigAppendResult(
            ImmutableArray<string> headers,

            ImmutableArray<string> jsonTypes,

            ImmutableArray<string> jsons)
        {
            Headers = headers;
            JsonTypes = jsonTypes;
            Jsons = jsons;
        }
    }
}