// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect
{
    public static class GetGatewayPluginAiPromptDecorator
    {
        public static Task<GetGatewayPluginAiPromptDecoratorResult> InvokeAsync(GetGatewayPluginAiPromptDecoratorArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGatewayPluginAiPromptDecoratorResult>("konnect:index/getGatewayPluginAiPromptDecorator:getGatewayPluginAiPromptDecorator", args ?? new GetGatewayPluginAiPromptDecoratorArgs(), options.WithDefaults());

        public static Output<GetGatewayPluginAiPromptDecoratorResult> Invoke(GetGatewayPluginAiPromptDecoratorInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGatewayPluginAiPromptDecoratorResult>("konnect:index/getGatewayPluginAiPromptDecorator:getGatewayPluginAiPromptDecorator", args ?? new GetGatewayPluginAiPromptDecoratorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGatewayPluginAiPromptDecoratorArgs : global::Pulumi.InvokeArgs
    {
        [Input("controlPlaneId", required: true)]
        public string ControlPlaneId { get; set; } = null!;

        public GetGatewayPluginAiPromptDecoratorArgs()
        {
        }
        public static new GetGatewayPluginAiPromptDecoratorArgs Empty => new GetGatewayPluginAiPromptDecoratorArgs();
    }

    public sealed class GetGatewayPluginAiPromptDecoratorInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("controlPlaneId", required: true)]
        public Input<string> ControlPlaneId { get; set; } = null!;

        public GetGatewayPluginAiPromptDecoratorInvokeArgs()
        {
        }
        public static new GetGatewayPluginAiPromptDecoratorInvokeArgs Empty => new GetGatewayPluginAiPromptDecoratorInvokeArgs();
    }


    [OutputType]
    public sealed class GetGatewayPluginAiPromptDecoratorResult
    {
        public readonly Outputs.GetGatewayPluginAiPromptDecoratorConfigResult Config;
        public readonly Outputs.GetGatewayPluginAiPromptDecoratorConsumerResult Consumer;
        public readonly Outputs.GetGatewayPluginAiPromptDecoratorConsumerGroupResult ConsumerGroup;
        public readonly string ControlPlaneId;
        public readonly int CreatedAt;
        public readonly bool Enabled;
        public readonly string Id;
        public readonly string InstanceName;
        public readonly ImmutableArray<string> Protocols;
        public readonly Outputs.GetGatewayPluginAiPromptDecoratorRouteResult Route;
        public readonly Outputs.GetGatewayPluginAiPromptDecoratorServiceResult Service;
        public readonly ImmutableArray<string> Tags;
        public readonly int UpdatedAt;

        [OutputConstructor]
        private GetGatewayPluginAiPromptDecoratorResult(
            Outputs.GetGatewayPluginAiPromptDecoratorConfigResult config,

            Outputs.GetGatewayPluginAiPromptDecoratorConsumerResult consumer,

            Outputs.GetGatewayPluginAiPromptDecoratorConsumerGroupResult consumerGroup,

            string controlPlaneId,

            int createdAt,

            bool enabled,

            string id,

            string instanceName,

            ImmutableArray<string> protocols,

            Outputs.GetGatewayPluginAiPromptDecoratorRouteResult route,

            Outputs.GetGatewayPluginAiPromptDecoratorServiceResult service,

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