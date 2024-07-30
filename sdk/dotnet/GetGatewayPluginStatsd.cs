// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect
{
    public static class GetGatewayPluginStatsd
    {
        public static Task<GetGatewayPluginStatsdResult> InvokeAsync(GetGatewayPluginStatsdArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGatewayPluginStatsdResult>("konnect:index/getGatewayPluginStatsd:getGatewayPluginStatsd", args ?? new GetGatewayPluginStatsdArgs(), options.WithDefaults());

        public static Output<GetGatewayPluginStatsdResult> Invoke(GetGatewayPluginStatsdInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGatewayPluginStatsdResult>("konnect:index/getGatewayPluginStatsd:getGatewayPluginStatsd", args ?? new GetGatewayPluginStatsdInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGatewayPluginStatsdArgs : global::Pulumi.InvokeArgs
    {
        [Input("controlPlaneId", required: true)]
        public string ControlPlaneId { get; set; } = null!;

        public GetGatewayPluginStatsdArgs()
        {
        }
        public static new GetGatewayPluginStatsdArgs Empty => new GetGatewayPluginStatsdArgs();
    }

    public sealed class GetGatewayPluginStatsdInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("controlPlaneId", required: true)]
        public Input<string> ControlPlaneId { get; set; } = null!;

        public GetGatewayPluginStatsdInvokeArgs()
        {
        }
        public static new GetGatewayPluginStatsdInvokeArgs Empty => new GetGatewayPluginStatsdInvokeArgs();
    }


    [OutputType]
    public sealed class GetGatewayPluginStatsdResult
    {
        public readonly Outputs.GetGatewayPluginStatsdConfigResult Config;
        public readonly Outputs.GetGatewayPluginStatsdConsumerResult Consumer;
        public readonly Outputs.GetGatewayPluginStatsdConsumerGroupResult ConsumerGroup;
        public readonly string ControlPlaneId;
        public readonly int CreatedAt;
        public readonly bool Enabled;
        public readonly string Id;
        public readonly string InstanceName;
        public readonly ImmutableArray<string> Protocols;
        public readonly Outputs.GetGatewayPluginStatsdRouteResult Route;
        public readonly Outputs.GetGatewayPluginStatsdServiceResult Service;
        public readonly ImmutableArray<string> Tags;
        public readonly int UpdatedAt;

        [OutputConstructor]
        private GetGatewayPluginStatsdResult(
            Outputs.GetGatewayPluginStatsdConfigResult config,

            Outputs.GetGatewayPluginStatsdConsumerResult consumer,

            Outputs.GetGatewayPluginStatsdConsumerGroupResult consumerGroup,

            string controlPlaneId,

            int createdAt,

            bool enabled,

            string id,

            string instanceName,

            ImmutableArray<string> protocols,

            Outputs.GetGatewayPluginStatsdRouteResult route,

            Outputs.GetGatewayPluginStatsdServiceResult service,

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
