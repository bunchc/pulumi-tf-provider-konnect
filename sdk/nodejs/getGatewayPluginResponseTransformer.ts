// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getGatewayPluginResponseTransformer(args: GetGatewayPluginResponseTransformerArgs, opts?: pulumi.InvokeOptions): Promise<GetGatewayPluginResponseTransformerResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("konnect:index/getGatewayPluginResponseTransformer:getGatewayPluginResponseTransformer", {
        "controlPlaneId": args.controlPlaneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGatewayPluginResponseTransformer.
 */
export interface GetGatewayPluginResponseTransformerArgs {
    controlPlaneId: string;
}

/**
 * A collection of values returned by getGatewayPluginResponseTransformer.
 */
export interface GetGatewayPluginResponseTransformerResult {
    readonly config: outputs.GetGatewayPluginResponseTransformerConfig;
    readonly consumer: outputs.GetGatewayPluginResponseTransformerConsumer;
    readonly consumerGroup: outputs.GetGatewayPluginResponseTransformerConsumerGroup;
    readonly controlPlaneId: string;
    readonly createdAt: number;
    readonly enabled: boolean;
    readonly id: string;
    readonly instanceName: string;
    readonly protocols: string[];
    readonly route: outputs.GetGatewayPluginResponseTransformerRoute;
    readonly service: outputs.GetGatewayPluginResponseTransformerService;
    readonly tags: string[];
    readonly updatedAt: number;
}
export function getGatewayPluginResponseTransformerOutput(args: GetGatewayPluginResponseTransformerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGatewayPluginResponseTransformerResult> {
    return pulumi.output(args).apply((a: any) => getGatewayPluginResponseTransformer(a, opts))
}

/**
 * A collection of arguments for invoking getGatewayPluginResponseTransformer.
 */
export interface GetGatewayPluginResponseTransformerOutputArgs {
    controlPlaneId: pulumi.Input<string>;
}