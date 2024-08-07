// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect
{
    public static class GetGatewayTarget
    {
        public static Task<GetGatewayTargetResult> InvokeAsync(GetGatewayTargetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGatewayTargetResult>("konnect:index/getGatewayTarget:getGatewayTarget", args ?? new GetGatewayTargetArgs(), options.WithDefaults());

        public static Output<GetGatewayTargetResult> Invoke(GetGatewayTargetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGatewayTargetResult>("konnect:index/getGatewayTarget:getGatewayTarget", args ?? new GetGatewayTargetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGatewayTargetArgs : global::Pulumi.InvokeArgs
    {
        [Input("controlPlaneId", required: true)]
        public string ControlPlaneId { get; set; } = null!;

        [Input("upstreamId", required: true)]
        public string UpstreamId { get; set; } = null!;

        public GetGatewayTargetArgs()
        {
        }
        public static new GetGatewayTargetArgs Empty => new GetGatewayTargetArgs();
    }

    public sealed class GetGatewayTargetInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("controlPlaneId", required: true)]
        public Input<string> ControlPlaneId { get; set; } = null!;

        [Input("upstreamId", required: true)]
        public Input<string> UpstreamId { get; set; } = null!;

        public GetGatewayTargetInvokeArgs()
        {
        }
        public static new GetGatewayTargetInvokeArgs Empty => new GetGatewayTargetInvokeArgs();
    }


    [OutputType]
    public sealed class GetGatewayTargetResult
    {
        public readonly string ControlPlaneId;
        public readonly double CreatedAt;
        public readonly string Id;
        public readonly ImmutableArray<string> Tags;
        public readonly string Target;
        public readonly double UpdatedAt;
        public readonly Outputs.GetGatewayTargetUpstreamResult Upstream;
        public readonly string UpstreamId;
        public readonly int Weight;

        [OutputConstructor]
        private GetGatewayTargetResult(
            string controlPlaneId,

            double createdAt,

            string id,

            ImmutableArray<string> tags,

            string target,

            double updatedAt,

            Outputs.GetGatewayTargetUpstreamResult upstream,

            string upstreamId,

            int weight)
        {
            ControlPlaneId = controlPlaneId;
            CreatedAt = createdAt;
            Id = id;
            Tags = tags;
            Target = target;
            UpdatedAt = updatedAt;
            Upstream = upstream;
            UpstreamId = upstreamId;
            Weight = weight;
        }
    }
}
