// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getGatewayPluginJwt(args: GetGatewayPluginJwtArgs, opts?: pulumi.InvokeOptions): Promise<GetGatewayPluginJwtResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("konnect:index/getGatewayPluginJwt:getGatewayPluginJwt", {
        "controlPlaneId": args.controlPlaneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGatewayPluginJwt.
 */
export interface GetGatewayPluginJwtArgs {
    controlPlaneId: string;
}

/**
 * A collection of values returned by getGatewayPluginJwt.
 */
export interface GetGatewayPluginJwtResult {
    readonly config: outputs.GetGatewayPluginJwtConfig;
    readonly consumer: outputs.GetGatewayPluginJwtConsumer;
    readonly consumerGroup: outputs.GetGatewayPluginJwtConsumerGroup;
    readonly controlPlaneId: string;
    readonly createdAt: number;
    readonly enabled: boolean;
    readonly id: string;
    readonly instanceName: string;
    readonly protocols: string[];
    readonly route: outputs.GetGatewayPluginJwtRoute;
    readonly service: outputs.GetGatewayPluginJwtService;
    readonly tags: string[];
    readonly updatedAt: number;
}
export function getGatewayPluginJwtOutput(args: GetGatewayPluginJwtOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGatewayPluginJwtResult> {
    return pulumi.output(args).apply((a: any) => getGatewayPluginJwt(a, opts))
}

/**
 * A collection of arguments for invoking getGatewayPluginJwt.
 */
export interface GetGatewayPluginJwtOutputArgs {
    controlPlaneId: pulumi.Input<string>;
}