// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect
{
    public static class GetGatewayPluginPostFunction
    {
        public static Task<GetGatewayPluginPostFunctionResult> InvokeAsync(GetGatewayPluginPostFunctionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGatewayPluginPostFunctionResult>("konnect:index/getGatewayPluginPostFunction:getGatewayPluginPostFunction", args ?? new GetGatewayPluginPostFunctionArgs(), options.WithDefaults());

        public static Output<GetGatewayPluginPostFunctionResult> Invoke(GetGatewayPluginPostFunctionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGatewayPluginPostFunctionResult>("konnect:index/getGatewayPluginPostFunction:getGatewayPluginPostFunction", args ?? new GetGatewayPluginPostFunctionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGatewayPluginPostFunctionArgs : global::Pulumi.InvokeArgs
    {
        [Input("controlPlaneId", required: true)]
        public string ControlPlaneId { get; set; } = null!;

        public GetGatewayPluginPostFunctionArgs()
        {
        }
        public static new GetGatewayPluginPostFunctionArgs Empty => new GetGatewayPluginPostFunctionArgs();
    }

    public sealed class GetGatewayPluginPostFunctionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("controlPlaneId", required: true)]
        public Input<string> ControlPlaneId { get; set; } = null!;

        public GetGatewayPluginPostFunctionInvokeArgs()
        {
        }
        public static new GetGatewayPluginPostFunctionInvokeArgs Empty => new GetGatewayPluginPostFunctionInvokeArgs();
    }


    [OutputType]
    public sealed class GetGatewayPluginPostFunctionResult
    {
        public readonly Outputs.GetGatewayPluginPostFunctionConfigResult Config;
        public readonly Outputs.GetGatewayPluginPostFunctionConsumerResult Consumer;
        public readonly Outputs.GetGatewayPluginPostFunctionConsumerGroupResult ConsumerGroup;
        public readonly string ControlPlaneId;
        public readonly int CreatedAt;
        public readonly bool Enabled;
        public readonly string Id;
        public readonly string InstanceName;
        public readonly ImmutableArray<string> Protocols;
        public readonly Outputs.GetGatewayPluginPostFunctionRouteResult Route;
        public readonly Outputs.GetGatewayPluginPostFunctionServiceResult Service;
        public readonly ImmutableArray<string> Tags;
        public readonly int UpdatedAt;

        [OutputConstructor]
        private GetGatewayPluginPostFunctionResult(
            Outputs.GetGatewayPluginPostFunctionConfigResult config,

            Outputs.GetGatewayPluginPostFunctionConsumerResult consumer,

            Outputs.GetGatewayPluginPostFunctionConsumerGroupResult consumerGroup,

            string controlPlaneId,

            int createdAt,

            bool enabled,

            string id,

            string instanceName,

            ImmutableArray<string> protocols,

            Outputs.GetGatewayPluginPostFunctionRouteResult route,

            Outputs.GetGatewayPluginPostFunctionServiceResult service,

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