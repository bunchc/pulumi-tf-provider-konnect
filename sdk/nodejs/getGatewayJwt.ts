// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getGatewayJwt(args: GetGatewayJwtArgs, opts?: pulumi.InvokeOptions): Promise<GetGatewayJwtResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("konnect:index/getGatewayJwt:getGatewayJwt", {
        "consumerId": args.consumerId,
        "controlPlaneId": args.controlPlaneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGatewayJwt.
 */
export interface GetGatewayJwtArgs {
    consumerId: string;
    controlPlaneId: string;
}

/**
 * A collection of values returned by getGatewayJwt.
 */
export interface GetGatewayJwtResult {
    readonly algorithm: string;
    readonly consumer: outputs.GetGatewayJwtConsumer;
    readonly consumerId: string;
    readonly controlPlaneId: string;
    readonly createdAt: number;
    readonly id: string;
    readonly key: string;
    readonly rsaPublicKey: string;
    readonly secret: string;
    readonly tags: string[];
}
export function getGatewayJwtOutput(args: GetGatewayJwtOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGatewayJwtResult> {
    return pulumi.output(args).apply((a: any) => getGatewayJwt(a, opts))
}

/**
 * A collection of arguments for invoking getGatewayJwt.
 */
export interface GetGatewayJwtOutputArgs {
    consumerId: pulumi.Input<string>;
    controlPlaneId: pulumi.Input<string>;
}
