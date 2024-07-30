// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect
{
    public static class GetApiProductSpecification
    {
        public static Task<GetApiProductSpecificationResult> InvokeAsync(GetApiProductSpecificationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApiProductSpecificationResult>("konnect:index/getApiProductSpecification:getApiProductSpecification", args ?? new GetApiProductSpecificationArgs(), options.WithDefaults());

        public static Output<GetApiProductSpecificationResult> Invoke(GetApiProductSpecificationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApiProductSpecificationResult>("konnect:index/getApiProductSpecification:getApiProductSpecification", args ?? new GetApiProductSpecificationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApiProductSpecificationArgs : global::Pulumi.InvokeArgs
    {
        [Input("apiProductId", required: true)]
        public string ApiProductId { get; set; } = null!;

        [Input("apiProductVersionId", required: true)]
        public string ApiProductVersionId { get; set; } = null!;

        public GetApiProductSpecificationArgs()
        {
        }
        public static new GetApiProductSpecificationArgs Empty => new GetApiProductSpecificationArgs();
    }

    public sealed class GetApiProductSpecificationInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("apiProductId", required: true)]
        public Input<string> ApiProductId { get; set; } = null!;

        [Input("apiProductVersionId", required: true)]
        public Input<string> ApiProductVersionId { get; set; } = null!;

        public GetApiProductSpecificationInvokeArgs()
        {
        }
        public static new GetApiProductSpecificationInvokeArgs Empty => new GetApiProductSpecificationInvokeArgs();
    }


    [OutputType]
    public sealed class GetApiProductSpecificationResult
    {
        public readonly string ApiProductId;
        public readonly string ApiProductVersionId;
        public readonly string Content;
        public readonly string CreatedAt;
        public readonly string Id;
        public readonly string Name;
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetApiProductSpecificationResult(
            string apiProductId,

            string apiProductVersionId,

            string content,

            string createdAt,

            string id,

            string name,

            string updatedAt)
        {
            ApiProductId = apiProductId;
            ApiProductVersionId = apiProductVersionId;
            Content = content;
            CreatedAt = createdAt;
            Id = id;
            Name = name;
            UpdatedAt = updatedAt;
        }
    }
}