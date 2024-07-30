// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect
{
    public static class GetGatewayKey
    {
        public static Task<GetGatewayKeyResult> InvokeAsync(GetGatewayKeyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGatewayKeyResult>("konnect:index/getGatewayKey:getGatewayKey", args ?? new GetGatewayKeyArgs(), options.WithDefaults());

        public static Output<GetGatewayKeyResult> Invoke(GetGatewayKeyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGatewayKeyResult>("konnect:index/getGatewayKey:getGatewayKey", args ?? new GetGatewayKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGatewayKeyArgs : global::Pulumi.InvokeArgs
    {
        [Input("controlPlaneId", required: true)]
        public string ControlPlaneId { get; set; } = null!;

        public GetGatewayKeyArgs()
        {
        }
        public static new GetGatewayKeyArgs Empty => new GetGatewayKeyArgs();
    }

    public sealed class GetGatewayKeyInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("controlPlaneId", required: true)]
        public Input<string> ControlPlaneId { get; set; } = null!;

        public GetGatewayKeyInvokeArgs()
        {
        }
        public static new GetGatewayKeyInvokeArgs Empty => new GetGatewayKeyInvokeArgs();
    }


    [OutputType]
    public sealed class GetGatewayKeyResult
    {
        public readonly string ControlPlaneId;
        public readonly int CreatedAt;
        public readonly string Id;
        public readonly string Jwk;
        public readonly string Kid;
        public readonly string Name;
        public readonly Outputs.GetGatewayKeyPemResult Pem;
        public readonly Outputs.GetGatewayKeySetResult Set;
        public readonly ImmutableArray<string> Tags;
        public readonly int UpdatedAt;

        [OutputConstructor]
        private GetGatewayKeyResult(
            string controlPlaneId,

            int createdAt,

            string id,

            string jwk,

            string kid,

            string name,

            Outputs.GetGatewayKeyPemResult pem,

            Outputs.GetGatewayKeySetResult set,

            ImmutableArray<string> tags,

            int updatedAt)
        {
            ControlPlaneId = controlPlaneId;
            CreatedAt = createdAt;
            Id = id;
            Jwk = jwk;
            Kid = kid;
            Name = name;
            Pem = pem;
            Set = set;
            Tags = tags;
            UpdatedAt = updatedAt;
        }
    }
}
