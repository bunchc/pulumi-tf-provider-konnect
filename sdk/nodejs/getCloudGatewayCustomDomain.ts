// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getCloudGatewayCustomDomain(opts?: pulumi.InvokeOptions): Promise<GetCloudGatewayCustomDomainResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("konnect:index/getCloudGatewayCustomDomain:getCloudGatewayCustomDomain", {
    }, opts);
}

/**
 * A collection of values returned by getCloudGatewayCustomDomain.
 */
export interface GetCloudGatewayCustomDomainResult {
    readonly certificateId: string;
    readonly controlPlaneGeo: string;
    readonly controlPlaneId: string;
    readonly createdAt: string;
    readonly domain: string;
    readonly entityVersion: number;
    readonly id: string;
    readonly sniId: string;
    readonly state: string;
    readonly stateMetadata: outputs.GetCloudGatewayCustomDomainStateMetadata;
    readonly updatedAt: string;
}
export function getCloudGatewayCustomDomainOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetCloudGatewayCustomDomainResult> {
    return pulumi.output(getCloudGatewayCustomDomain(opts))
}
