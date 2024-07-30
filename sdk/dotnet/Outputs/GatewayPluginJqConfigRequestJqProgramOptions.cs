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
    public sealed class GatewayPluginJqConfigRequestJqProgramOptions
    {
        public readonly bool? AsciiOutput;
        public readonly bool? CompactOutput;
        public readonly bool? JoinOutput;
        public readonly bool? RawOutput;
        public readonly bool? SortKeys;

        [OutputConstructor]
        private GatewayPluginJqConfigRequestJqProgramOptions(
            bool? asciiOutput,

            bool? compactOutput,

            bool? joinOutput,

            bool? rawOutput,

            bool? sortKeys)
        {
            AsciiOutput = asciiOutput;
            CompactOutput = compactOutput;
            JoinOutput = joinOutput;
            RawOutput = rawOutput;
            SortKeys = sortKeys;
        }
    }
}
