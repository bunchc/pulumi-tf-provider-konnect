// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect
{
    public static class GetPortalList
    {
        public static Task<GetPortalListResult> InvokeAsync(GetPortalListArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPortalListResult>("konnect:index/getPortalList:getPortalList", args ?? new GetPortalListArgs(), options.WithDefaults());

        public static Output<GetPortalListResult> Invoke(GetPortalListInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPortalListResult>("konnect:index/getPortalList:getPortalList", args ?? new GetPortalListInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPortalListArgs : global::Pulumi.InvokeArgs
    {
        [Input("pageNumber")]
        public int? PageNumber { get; set; }

        [Input("pageSize")]
        public int? PageSize { get; set; }

        [Input("sort")]
        public string? Sort { get; set; }

        public GetPortalListArgs()
        {
        }
        public static new GetPortalListArgs Empty => new GetPortalListArgs();
    }

    public sealed class GetPortalListInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("pageNumber")]
        public Input<int>? PageNumber { get; set; }

        [Input("pageSize")]
        public Input<int>? PageSize { get; set; }

        [Input("sort")]
        public Input<string>? Sort { get; set; }

        public GetPortalListInvokeArgs()
        {
        }
        public static new GetPortalListInvokeArgs Empty => new GetPortalListInvokeArgs();
    }


    [OutputType]
    public sealed class GetPortalListResult
    {
        public readonly ImmutableArray<Outputs.GetPortalListDataResult> Datas;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly Outputs.GetPortalListMetaResult Meta;
        public readonly int? PageNumber;
        public readonly int? PageSize;
        public readonly string? Sort;

        [OutputConstructor]
        private GetPortalListResult(
            ImmutableArray<Outputs.GetPortalListDataResult> datas,

            string id,

            Outputs.GetPortalListMetaResult meta,

            int? pageNumber,

            int? pageSize,

            string? sort)
        {
            Datas = datas;
            Id = id;
            Meta = meta;
            PageNumber = pageNumber;
            PageSize = pageSize;
            Sort = sort;
        }
    }
}