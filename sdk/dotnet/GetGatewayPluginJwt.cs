// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect
{
    public static class GetGatewayPluginJwt
    {
        public static Task<GetGatewayPluginJwtResult> InvokeAsync(GetGatewayPluginJwtArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGatewayPluginJwtResult>("konnect:index/getGatewayPluginJwt:getGatewayPluginJwt", args ?? new GetGatewayPluginJwtArgs(), options.WithDefaults());

        public static Output<GetGatewayPluginJwtResult> Invoke(GetGatewayPluginJwtInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGatewayPluginJwtResult>("konnect:index/getGatewayPluginJwt:getGatewayPluginJwt", args ?? new GetGatewayPluginJwtInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGatewayPluginJwtArgs : global::Pulumi.InvokeArgs
    {
        [Input("controlPlaneId", required: true)]
        public string ControlPlaneId { get; set; } = null!;

        public GetGatewayPluginJwtArgs()
        {
        }
        public static new GetGatewayPluginJwtArgs Empty => new GetGatewayPluginJwtArgs();
    }

    public sealed class GetGatewayPluginJwtInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("controlPlaneId", required: true)]
        public Input<string> ControlPlaneId { get; set; } = null!;

        public GetGatewayPluginJwtInvokeArgs()
        {
        }
        public static new GetGatewayPluginJwtInvokeArgs Empty => new GetGatewayPluginJwtInvokeArgs();
    }


    [OutputType]
    public sealed class GetGatewayPluginJwtResult
    {
        public readonly Outputs.GetGatewayPluginJwtConfigResult Config;
        public readonly Outputs.GetGatewayPluginJwtConsumerResult Consumer;
        public readonly Outputs.GetGatewayPluginJwtConsumerGroupResult ConsumerGroup;
        public readonly string ControlPlaneId;
        public readonly int CreatedAt;
        public readonly bool Enabled;
        public readonly string Id;
        public readonly string InstanceName;
        public readonly ImmutableArray<string> Protocols;
        public readonly Outputs.GetGatewayPluginJwtRouteResult Route;
        public readonly Outputs.GetGatewayPluginJwtServiceResult Service;
        public readonly ImmutableArray<string> Tags;
        public readonly int UpdatedAt;

        [OutputConstructor]
        private GetGatewayPluginJwtResult(
            Outputs.GetGatewayPluginJwtConfigResult config,

            Outputs.GetGatewayPluginJwtConsumerResult consumer,

            Outputs.GetGatewayPluginJwtConsumerGroupResult consumerGroup,

            string controlPlaneId,

            int createdAt,

            bool enabled,

            string id,

            string instanceName,

            ImmutableArray<string> protocols,

            Outputs.GetGatewayPluginJwtRouteResult route,

            Outputs.GetGatewayPluginJwtServiceResult service,

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