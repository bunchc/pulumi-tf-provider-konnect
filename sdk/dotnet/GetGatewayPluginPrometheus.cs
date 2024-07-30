// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect
{
    public static class GetGatewayPluginPrometheus
    {
        public static Task<GetGatewayPluginPrometheusResult> InvokeAsync(GetGatewayPluginPrometheusArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGatewayPluginPrometheusResult>("konnect:index/getGatewayPluginPrometheus:getGatewayPluginPrometheus", args ?? new GetGatewayPluginPrometheusArgs(), options.WithDefaults());

        public static Output<GetGatewayPluginPrometheusResult> Invoke(GetGatewayPluginPrometheusInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGatewayPluginPrometheusResult>("konnect:index/getGatewayPluginPrometheus:getGatewayPluginPrometheus", args ?? new GetGatewayPluginPrometheusInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGatewayPluginPrometheusArgs : global::Pulumi.InvokeArgs
    {
        [Input("controlPlaneId", required: true)]
        public string ControlPlaneId { get; set; } = null!;

        public GetGatewayPluginPrometheusArgs()
        {
        }
        public static new GetGatewayPluginPrometheusArgs Empty => new GetGatewayPluginPrometheusArgs();
    }

    public sealed class GetGatewayPluginPrometheusInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("controlPlaneId", required: true)]
        public Input<string> ControlPlaneId { get; set; } = null!;

        public GetGatewayPluginPrometheusInvokeArgs()
        {
        }
        public static new GetGatewayPluginPrometheusInvokeArgs Empty => new GetGatewayPluginPrometheusInvokeArgs();
    }


    [OutputType]
    public sealed class GetGatewayPluginPrometheusResult
    {
        public readonly Outputs.GetGatewayPluginPrometheusConfigResult Config;
        public readonly Outputs.GetGatewayPluginPrometheusConsumerResult Consumer;
        public readonly Outputs.GetGatewayPluginPrometheusConsumerGroupResult ConsumerGroup;
        public readonly string ControlPlaneId;
        public readonly int CreatedAt;
        public readonly bool Enabled;
        public readonly string Id;
        public readonly string InstanceName;
        public readonly ImmutableArray<string> Protocols;
        public readonly Outputs.GetGatewayPluginPrometheusRouteResult Route;
        public readonly Outputs.GetGatewayPluginPrometheusServiceResult Service;
        public readonly ImmutableArray<string> Tags;
        public readonly int UpdatedAt;

        [OutputConstructor]
        private GetGatewayPluginPrometheusResult(
            Outputs.GetGatewayPluginPrometheusConfigResult config,

            Outputs.GetGatewayPluginPrometheusConsumerResult consumer,

            Outputs.GetGatewayPluginPrometheusConsumerGroupResult consumerGroup,

            string controlPlaneId,

            int createdAt,

            bool enabled,

            string id,

            string instanceName,

            ImmutableArray<string> protocols,

            Outputs.GetGatewayPluginPrometheusRouteResult route,

            Outputs.GetGatewayPluginPrometheusServiceResult service,

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