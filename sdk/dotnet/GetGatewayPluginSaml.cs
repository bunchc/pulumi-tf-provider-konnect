// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect
{
    public static class GetGatewayPluginSaml
    {
        public static Task<GetGatewayPluginSamlResult> InvokeAsync(GetGatewayPluginSamlArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGatewayPluginSamlResult>("konnect:index/getGatewayPluginSaml:getGatewayPluginSaml", args ?? new GetGatewayPluginSamlArgs(), options.WithDefaults());

        public static Output<GetGatewayPluginSamlResult> Invoke(GetGatewayPluginSamlInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGatewayPluginSamlResult>("konnect:index/getGatewayPluginSaml:getGatewayPluginSaml", args ?? new GetGatewayPluginSamlInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGatewayPluginSamlArgs : global::Pulumi.InvokeArgs
    {
        [Input("controlPlaneId", required: true)]
        public string ControlPlaneId { get; set; } = null!;

        public GetGatewayPluginSamlArgs()
        {
        }
        public static new GetGatewayPluginSamlArgs Empty => new GetGatewayPluginSamlArgs();
    }

    public sealed class GetGatewayPluginSamlInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("controlPlaneId", required: true)]
        public Input<string> ControlPlaneId { get; set; } = null!;

        public GetGatewayPluginSamlInvokeArgs()
        {
        }
        public static new GetGatewayPluginSamlInvokeArgs Empty => new GetGatewayPluginSamlInvokeArgs();
    }


    [OutputType]
    public sealed class GetGatewayPluginSamlResult
    {
        public readonly Outputs.GetGatewayPluginSamlConfigResult Config;
        public readonly Outputs.GetGatewayPluginSamlConsumerResult Consumer;
        public readonly Outputs.GetGatewayPluginSamlConsumerGroupResult ConsumerGroup;
        public readonly string ControlPlaneId;
        public readonly int CreatedAt;
        public readonly bool Enabled;
        public readonly string Id;
        public readonly string InstanceName;
        public readonly ImmutableArray<string> Protocols;
        public readonly Outputs.GetGatewayPluginSamlRouteResult Route;
        public readonly Outputs.GetGatewayPluginSamlServiceResult Service;
        public readonly ImmutableArray<string> Tags;
        public readonly int UpdatedAt;

        [OutputConstructor]
        private GetGatewayPluginSamlResult(
            Outputs.GetGatewayPluginSamlConfigResult config,

            Outputs.GetGatewayPluginSamlConsumerResult consumer,

            Outputs.GetGatewayPluginSamlConsumerGroupResult consumerGroup,

            string controlPlaneId,

            int createdAt,

            bool enabled,

            string id,

            string instanceName,

            ImmutableArray<string> protocols,

            Outputs.GetGatewayPluginSamlRouteResult route,

            Outputs.GetGatewayPluginSamlServiceResult service,

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
