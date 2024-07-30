// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect
{
    public static class GetGatewayPluginAcl
    {
        public static Task<GetGatewayPluginAclResult> InvokeAsync(GetGatewayPluginAclArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGatewayPluginAclResult>("konnect:index/getGatewayPluginAcl:getGatewayPluginAcl", args ?? new GetGatewayPluginAclArgs(), options.WithDefaults());

        public static Output<GetGatewayPluginAclResult> Invoke(GetGatewayPluginAclInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGatewayPluginAclResult>("konnect:index/getGatewayPluginAcl:getGatewayPluginAcl", args ?? new GetGatewayPluginAclInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGatewayPluginAclArgs : global::Pulumi.InvokeArgs
    {
        [Input("controlPlaneId", required: true)]
        public string ControlPlaneId { get; set; } = null!;

        public GetGatewayPluginAclArgs()
        {
        }
        public static new GetGatewayPluginAclArgs Empty => new GetGatewayPluginAclArgs();
    }

    public sealed class GetGatewayPluginAclInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("controlPlaneId", required: true)]
        public Input<string> ControlPlaneId { get; set; } = null!;

        public GetGatewayPluginAclInvokeArgs()
        {
        }
        public static new GetGatewayPluginAclInvokeArgs Empty => new GetGatewayPluginAclInvokeArgs();
    }


    [OutputType]
    public sealed class GetGatewayPluginAclResult
    {
        public readonly Outputs.GetGatewayPluginAclConfigResult Config;
        public readonly Outputs.GetGatewayPluginAclConsumerResult Consumer;
        public readonly Outputs.GetGatewayPluginAclConsumerGroupResult ConsumerGroup;
        public readonly string ControlPlaneId;
        public readonly int CreatedAt;
        public readonly bool Enabled;
        public readonly string Id;
        public readonly string InstanceName;
        public readonly ImmutableArray<string> Protocols;
        public readonly Outputs.GetGatewayPluginAclRouteResult Route;
        public readonly Outputs.GetGatewayPluginAclServiceResult Service;
        public readonly ImmutableArray<string> Tags;
        public readonly int UpdatedAt;

        [OutputConstructor]
        private GetGatewayPluginAclResult(
            Outputs.GetGatewayPluginAclConfigResult config,

            Outputs.GetGatewayPluginAclConsumerResult consumer,

            Outputs.GetGatewayPluginAclConsumerGroupResult consumerGroup,

            string controlPlaneId,

            int createdAt,

            bool enabled,

            string id,

            string instanceName,

            ImmutableArray<string> protocols,

            Outputs.GetGatewayPluginAclRouteResult route,

            Outputs.GetGatewayPluginAclServiceResult service,

            ImmutableArray<string> tags,

            int updatedAt)
        {
            Config = config;
            Consumer = consumer;
            ConsumerGroup = consumerGroup;
            ControlPlaneId = controlPlaneId;
            CreatedAt = createdAt;
            Enabled = enabled;
            Id = id;
            InstanceName = instanceName;
            Protocols = protocols;
            Route = route;
            Service = service;
            Tags = tags;
            UpdatedAt = updatedAt;
        }
    }
}