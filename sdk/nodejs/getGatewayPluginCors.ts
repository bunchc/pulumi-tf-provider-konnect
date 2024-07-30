// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getGatewayPluginCors(args: GetGatewayPluginCorsArgs, opts?: pulumi.InvokeOptions): Promise<GetGatewayPluginCorsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("konnect:index/getGatewayPluginCors:getGatewayPluginCors", {
        "controlPlaneId": args.controlPlaneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGatewayPluginCors.
 */
export interface GetGatewayPluginCorsArgs {
    controlPlaneId: string;
}

/**
 * A collection of values returned by getGatewayPluginCors.
 */
export interface GetGatewayPluginCorsResult {
    readonly config: outputs.GetGatewayPluginCorsConfig;
    readonly consumer: outputs.GetGatewayPluginCorsConsumer;
    readonly consumerGroup: outputs.GetGatewayPluginCorsConsumerGroup;
    readonly controlPlaneId: string;
    readonly createdAt: number;
    readonly enabled: boolean;
    readonly id: string;
    readonly instanceName: string;
    readonly protocols: string[];
    readonly route: outputs.GetGatewayPluginCorsRoute;
    readonly service: outputs.GetGatewayPluginCorsService;
    readonly tags: string[];
    readonly updatedAt: number;
}
export function getGatewayPluginCorsOutput(args: GetGatewayPluginCorsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGatewayPluginCorsResult> {
    return pulumi.output(args).apply((a: any) => getGatewayPluginCors(a, opts))
}

/**
 * A collection of arguments for invoking getGatewayPluginCors.
 */
export interface GetGatewayPluginCorsOutputArgs {
    controlPlaneId: pulumi.Input<string>;
}