// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getGatewaySni(args: GetGatewaySniArgs, opts?: pulumi.InvokeOptions): Promise<GetGatewaySniResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("konnect:index/getGatewaySni:getGatewaySni", {
        "controlPlaneId": args.controlPlaneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGatewaySni.
 */
export interface GetGatewaySniArgs {
    controlPlaneId: string;
}

/**
 * A collection of values returned by getGatewaySni.
 */
export interface GetGatewaySniResult {
    readonly certificate: outputs.GetGatewaySniCertificate;
    readonly controlPlaneId: string;
    readonly createdAt: number;
    readonly id: string;
    readonly name: string;
    readonly tags: string[];
    readonly updatedAt: number;
}
export function getGatewaySniOutput(args: GetGatewaySniOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGatewaySniResult> {
    return pulumi.output(args).apply((a: any) => getGatewaySni(a, opts))
}

/**
 * A collection of arguments for invoking getGatewaySni.
 */
export interface GetGatewaySniOutputArgs {
    controlPlaneId: pulumi.Input<string>;
}