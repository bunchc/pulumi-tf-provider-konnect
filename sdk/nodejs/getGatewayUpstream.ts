// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getGatewayUpstream(args: GetGatewayUpstreamArgs, opts?: pulumi.InvokeOptions): Promise<GetGatewayUpstreamResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("konnect:index/getGatewayUpstream:getGatewayUpstream", {
        "controlPlaneId": args.controlPlaneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGatewayUpstream.
 */
export interface GetGatewayUpstreamArgs {
    controlPlaneId: string;
}

/**
 * A collection of values returned by getGatewayUpstream.
 */
export interface GetGatewayUpstreamResult {
    readonly algorithm: string;
    readonly clientCertificate: outputs.GetGatewayUpstreamClientCertificate;
    readonly controlPlaneId: string;
    readonly createdAt: number;
    readonly hashFallback: string;
    readonly hashFallbackHeader: string;
    readonly hashFallbackQueryArg: string;
    readonly hashFallbackUriCapture: string;
    readonly hashOn: string;
    readonly hashOnCookie: string;
    readonly hashOnCookiePath: string;
    readonly hashOnHeader: string;
    readonly hashOnQueryArg: string;
    readonly hashOnUriCapture: string;
    readonly healthchecks: outputs.GetGatewayUpstreamHealthchecks;
    readonly hostHeader: string;
    readonly id: string;
    readonly name: string;
    readonly slots: number;
    readonly tags: string[];
    readonly updatedAt: number;
    readonly useSrvName: boolean;
}
export function getGatewayUpstreamOutput(args: GetGatewayUpstreamOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGatewayUpstreamResult> {
    return pulumi.output(args).apply((a: any) => getGatewayUpstream(a, opts))
}

/**
 * A collection of arguments for invoking getGatewayUpstream.
 */
export interface GetGatewayUpstreamOutputArgs {
    controlPlaneId: pulumi.Input<string>;
}