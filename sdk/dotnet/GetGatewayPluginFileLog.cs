// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect
{
    public static class GetGatewayPluginFileLog
    {
        public static Task<GetGatewayPluginFileLogResult> InvokeAsync(GetGatewayPluginFileLogArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGatewayPluginFileLogResult>("konnect:index/getGatewayPluginFileLog:getGatewayPluginFileLog", args ?? new GetGatewayPluginFileLogArgs(), options.WithDefaults());

        public static Output<GetGatewayPluginFileLogResult> Invoke(GetGatewayPluginFileLogInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGatewayPluginFileLogResult>("konnect:index/getGatewayPluginFileLog:getGatewayPluginFileLog", args ?? new GetGatewayPluginFileLogInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGatewayPluginFileLogArgs : global::Pulumi.InvokeArgs
    {
        [Input("controlPlaneId", required: true)]
        public string ControlPlaneId { get; set; } = null!;

        public GetGatewayPluginFileLogArgs()
        {
        }
        public static new GetGatewayPluginFileLogArgs Empty => new GetGatewayPluginFileLogArgs();
    }

    public sealed class GetGatewayPluginFileLogInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("controlPlaneId", required: true)]
        public Input<string> ControlPlaneId { get; set; } = null!;

        public GetGatewayPluginFileLogInvokeArgs()
        {
        }
        public static new GetGatewayPluginFileLogInvokeArgs Empty => new GetGatewayPluginFileLogInvokeArgs();
    }


    [OutputType]
    public sealed class GetGatewayPluginFileLogResult
    {
        public readonly Outputs.GetGatewayPluginFileLogConfigResult Config;
        public readonly Outputs.GetGatewayPluginFileLogConsumerResult Consumer;
        public readonly Outputs.GetGatewayPluginFileLogConsumerGroupResult ConsumerGroup;
        public readonly string ControlPlaneId;
        public readonly int CreatedAt;
        public readonly bool Enabled;
        public readonly string Id;
        public readonly string InstanceName;
        public readonly ImmutableArray<string> Protocols;
        public readonly Outputs.GetGatewayPluginFileLogRouteResult Route;
        public readonly Outputs.GetGatewayPluginFileLogServiceResult Service;
        public readonly ImmutableArray<string> Tags;
        public readonly int UpdatedAt;

        [OutputConstructor]
        private GetGatewayPluginFileLogResult(
            Outputs.GetGatewayPluginFileLogConfigResult config,

            Outputs.GetGatewayPluginFileLogConsumerResult consumer,

            Outputs.GetGatewayPluginFileLogConsumerGroupResult consumerGroup,

            string controlPlaneId,

            int createdAt,

            bool enabled,

            string id,

            string instanceName,

            ImmutableArray<string> protocols,

            Outputs.GetGatewayPluginFileLogRouteResult route,

            Outputs.GetGatewayPluginFileLogServiceResult service,

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
