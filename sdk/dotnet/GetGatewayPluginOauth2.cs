// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect
{
    public static class GetGatewayPluginOauth2
    {
        public static Task<GetGatewayPluginOauth2Result> InvokeAsync(GetGatewayPluginOauth2Args args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGatewayPluginOauth2Result>("konnect:index/getGatewayPluginOauth2:getGatewayPluginOauth2", args ?? new GetGatewayPluginOauth2Args(), options.WithDefaults());

        public static Output<GetGatewayPluginOauth2Result> Invoke(GetGatewayPluginOauth2InvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGatewayPluginOauth2Result>("konnect:index/getGatewayPluginOauth2:getGatewayPluginOauth2", args ?? new GetGatewayPluginOauth2InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGatewayPluginOauth2Args : global::Pulumi.InvokeArgs
    {
        [Input("controlPlaneId", required: true)]
        public string ControlPlaneId { get; set; } = null!;

        public GetGatewayPluginOauth2Args()
        {
        }
        public static new GetGatewayPluginOauth2Args Empty => new GetGatewayPluginOauth2Args();
    }

    public sealed class GetGatewayPluginOauth2InvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("controlPlaneId", required: true)]
        public Input<string> ControlPlaneId { get; set; } = null!;

        public GetGatewayPluginOauth2InvokeArgs()
        {
        }
        public static new GetGatewayPluginOauth2InvokeArgs Empty => new GetGatewayPluginOauth2InvokeArgs();
    }


    [OutputType]
    public sealed class GetGatewayPluginOauth2Result
    {
        public readonly Outputs.GetGatewayPluginOauth2ConfigResult Config;
        public readonly Outputs.GetGatewayPluginOauth2ConsumerResult Consumer;
        public readonly Outputs.GetGatewayPluginOauth2ConsumerGroupResult ConsumerGroup;
        public readonly string ControlPlaneId;
        public readonly int CreatedAt;
        public readonly bool Enabled;
        public readonly string Id;
        public readonly string InstanceName;
        public readonly ImmutableArray<string> Protocols;
        public readonly Outputs.GetGatewayPluginOauth2RouteResult Route;
        public readonly Outputs.GetGatewayPluginOauth2ServiceResult Service;
        public readonly ImmutableArray<string> Tags;
        public readonly int UpdatedAt;

        [OutputConstructor]
        private GetGatewayPluginOauth2Result(
            Outputs.GetGatewayPluginOauth2ConfigResult config,

            Outputs.GetGatewayPluginOauth2ConsumerResult consumer,

            Outputs.GetGatewayPluginOauth2ConsumerGroupResult consumerGroup,

            string controlPlaneId,

            int createdAt,

            bool enabled,

            string id,

            string instanceName,

            ImmutableArray<string> protocols,

            Outputs.GetGatewayPluginOauth2RouteResult route,

            Outputs.GetGatewayPluginOauth2ServiceResult service,

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