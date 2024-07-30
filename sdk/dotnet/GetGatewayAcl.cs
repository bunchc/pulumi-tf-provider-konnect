// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect
{
    public static class GetGatewayAcl
    {
        public static Task<GetGatewayAclResult> InvokeAsync(GetGatewayAclArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGatewayAclResult>("konnect:index/getGatewayAcl:getGatewayAcl", args ?? new GetGatewayAclArgs(), options.WithDefaults());

        public static Output<GetGatewayAclResult> Invoke(GetGatewayAclInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGatewayAclResult>("konnect:index/getGatewayAcl:getGatewayAcl", args ?? new GetGatewayAclInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGatewayAclArgs : global::Pulumi.InvokeArgs
    {
        [Input("consumerId", required: true)]
        public string ConsumerId { get; set; } = null!;

        [Input("controlPlaneId", required: true)]
        public string ControlPlaneId { get; set; } = null!;

        public GetGatewayAclArgs()
        {
        }
        public static new GetGatewayAclArgs Empty => new GetGatewayAclArgs();
    }

    public sealed class GetGatewayAclInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("consumerId", required: true)]
        public Input<string> ConsumerId { get; set; } = null!;

        [Input("controlPlaneId", required: true)]
        public Input<string> ControlPlaneId { get; set; } = null!;

        public GetGatewayAclInvokeArgs()
        {
        }
        public static new GetGatewayAclInvokeArgs Empty => new GetGatewayAclInvokeArgs();
    }


    [OutputType]
    public sealed class GetGatewayAclResult
    {
        public readonly Outputs.GetGatewayAclConsumerResult Consumer;
        public readonly string ConsumerId;
        public readonly string ControlPlaneId;
        public readonly int CreatedAt;
        public readonly string Group;
        public readonly string Id;
        public readonly ImmutableArray<string> Tags;

        [OutputConstructor]
        private GetGatewayAclResult(
            Outputs.GetGatewayAclConsumerResult consumer,

            string consumerId,

            string controlPlaneId,

            int createdAt,

            string group,

            string id,

            ImmutableArray<string> tags)
        {
            Consumer = consumer;
            ConsumerId = consumerId;
            ControlPlaneId = controlPlaneId;
            CreatedAt = createdAt;
            Group = group;
            Id = id;
            Tags = tags;
        }
    }
}
