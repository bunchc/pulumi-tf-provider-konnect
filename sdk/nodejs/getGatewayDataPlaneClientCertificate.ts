// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getGatewayDataPlaneClientCertificate(args: GetGatewayDataPlaneClientCertificateArgs, opts?: pulumi.InvokeOptions): Promise<GetGatewayDataPlaneClientCertificateResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("konnect:index/getGatewayDataPlaneClientCertificate:getGatewayDataPlaneClientCertificate", {
        "controlPlaneId": args.controlPlaneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGatewayDataPlaneClientCertificate.
 */
export interface GetGatewayDataPlaneClientCertificateArgs {
    controlPlaneId: string;
}

/**
 * A collection of values returned by getGatewayDataPlaneClientCertificate.
 */
export interface GetGatewayDataPlaneClientCertificateResult {
    readonly cert: string;
    readonly controlPlaneId: string;
    readonly createdAt: number;
    readonly id: string;
    readonly updatedAt: number;
}
export function getGatewayDataPlaneClientCertificateOutput(args: GetGatewayDataPlaneClientCertificateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGatewayDataPlaneClientCertificateResult> {
    return pulumi.output(args).apply((a: any) => getGatewayDataPlaneClientCertificate(a, opts))
}

/**
 * A collection of arguments for invoking getGatewayDataPlaneClientCertificate.
 */
export interface GetGatewayDataPlaneClientCertificateOutputArgs {
    controlPlaneId: pulumi.Input<string>;
}
