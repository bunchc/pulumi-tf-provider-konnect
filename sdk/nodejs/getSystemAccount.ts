// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getSystemAccount(opts?: pulumi.InvokeOptions): Promise<GetSystemAccountResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("konnect:index/getSystemAccount:getSystemAccount", {
    }, opts);
}

/**
 * A collection of values returned by getSystemAccount.
 */
export interface GetSystemAccountResult {
    readonly createdAt: string;
    readonly description: string;
    readonly id: string;
    readonly konnectManaged: boolean;
    readonly name: string;
    readonly updatedAt: string;
}
export function getSystemAccountOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemAccountResult> {
    return pulumi.output(getSystemAccount(opts))
}
