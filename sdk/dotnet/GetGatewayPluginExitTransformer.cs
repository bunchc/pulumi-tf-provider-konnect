// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect
{
    public static class GetGatewayPluginExitTransformer
    {
        public static Task<GetGatewayPluginExitTransformerResult> InvokeAsync(GetGatewayPluginExitTransformerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGatewayPluginExitTransformerResult>("konnect:index/getGatewayPluginExitTransformer:getGatewayPluginExitTransformer", args ?? new GetGatewayPluginExitTransformerArgs(), options.WithDefaults());

        public static Output<GetGatewayPluginExitTransformerResult> Invoke(GetGatewayPluginExitTransformerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGatewayPluginExitTransformerResult>("konnect:index/getGatewayPluginExitTransformer:getGatewayPluginExitTransformer", args ?? new GetGatewayPluginExitTransformerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGatewayPluginExitTransformerArgs : global::Pulumi.InvokeArgs
    {
        [Input("controlPlaneId", required: true)]
        public string ControlPlaneId { get; set; } = null!;

        public GetGatewayPluginExitTransformerArgs()
        {
        }
        public static new GetGatewayPluginExitTransformerArgs Empty => new GetGatewayPluginExitTransformerArgs();
    }

    public sealed class GetGatewayPluginExitTransformerInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("controlPlaneId", required: true)]
        public Input<string> ControlPlaneId { get; set; } = null!;

        public GetGatewayPluginExitTransformerInvokeArgs()
        {
        }
        public static new GetGatewayPluginExitTransformerInvokeArgs Empty => new GetGatewayPluginExitTransformerInvokeArgs();
    }


    [OutputType]
    public sealed class GetGatewayPluginExitTransformerResult
    {
        public readonly Outputs.GetGatewayPluginExitTransformerConfigResult Config;
        public readonly Outputs.GetGatewayPluginExitTransformerConsumerResult Consumer;
        public readonly Outputs.GetGatewayPluginExitTransformerConsumerGroupResult ConsumerGroup;
        public readonly string ControlPlaneId;
        public readonly int CreatedAt;
        public readonly bool Enabled;
        public readonly string Id;
        public readonly string InstanceName;
        public readonly ImmutableArray<string> Protocols;
        public readonly Outputs.GetGatewayPluginExitTransformerRouteResult Route;
        public readonly Outputs.GetGatewayPluginExitTransformerServiceResult Service;
        public readonly ImmutableArray<string> Tags;
        public readonly int UpdatedAt;

        [OutputConstructor]
        private GetGatewayPluginExitTransformerResult(
            Outputs.GetGatewayPluginExitTransformerConfigResult config,

            Outputs.GetGatewayPluginExitTransformerConsumerResult consumer,

            Outputs.GetGatewayPluginExitTransformerConsumerGroupResult consumerGroup,

            string controlPlaneId,

            int createdAt,

            bool enabled,

            string id,

            string instanceName,

            ImmutableArray<string> protocols,

            Outputs.GetGatewayPluginExitTransformerRouteResult route,

            Outputs.GetGatewayPluginExitTransformerServiceResult service,

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