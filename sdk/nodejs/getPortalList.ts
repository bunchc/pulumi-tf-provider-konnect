// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getPortalList(args?: GetPortalListArgs, opts?: pulumi.InvokeOptions): Promise<GetPortalListResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("konnect:index/getPortalList:getPortalList", {
        "pageNumber": args.pageNumber,
        "pageSize": args.pageSize,
        "sort": args.sort,
    }, opts);
}

/**
 * A collection of arguments for invoking getPortalList.
 */
export interface GetPortalListArgs {
    pageNumber?: number;
    pageSize?: number;
    sort?: string;
}

/**
 * A collection of values returned by getPortalList.
 */
export interface GetPortalListResult {
    readonly datas: outputs.GetPortalListData[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly meta: outputs.GetPortalListMeta;
    readonly pageNumber?: number;
    readonly pageSize?: number;
    readonly sort?: string;
}
export function getPortalListOutput(args?: GetPortalListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPortalListResult> {
    return pulumi.output(args).apply((a: any) => getPortalList(a, opts))
}

/**
 * A collection of arguments for invoking getPortalList.
 */
export interface GetPortalListOutputArgs {
    pageNumber?: pulumi.Input<number>;
    pageSize?: pulumi.Input<number>;
    sort?: pulumi.Input<string>;
}
