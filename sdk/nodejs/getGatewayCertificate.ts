// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getGatewayCertificate(args: GetGatewayCertificateArgs, opts?: pulumi.InvokeOptions): Promise<GetGatewayCertificateResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("konnect:index/getGatewayCertificate:getGatewayCertificate", {
        "controlPlaneId": args.controlPlaneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGatewayCertificate.
 */
export interface GetGatewayCertificateArgs {
    controlPlaneId: string;
}

/**
 * A collection of values returned by getGatewayCertificate.
 */
export interface GetGatewayCertificateResult {
    readonly cert: string;
    readonly certAlt: string;
    readonly controlPlaneId: string;
    readonly createdAt: number;
    readonly id: string;
    readonly key: string;
    readonly keyAlt: string;
    readonly tags: string[];
    readonly updatedAt: number;
}
export function getGatewayCertificateOutput(args: GetGatewayCertificateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGatewayCertificateResult> {
    return pulumi.output(args).apply((a: any) => getGatewayCertificate(a, opts))
}

/**
 * A collection of arguments for invoking getGatewayCertificate.
 */
export interface GetGatewayCertificateOutputArgs {
    controlPlaneId: pulumi.Input<string>;
}
