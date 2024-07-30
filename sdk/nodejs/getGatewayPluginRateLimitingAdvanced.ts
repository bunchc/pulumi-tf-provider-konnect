// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getGatewayPluginRateLimitingAdvanced(args: GetGatewayPluginRateLimitingAdvancedArgs, opts?: pulumi.InvokeOptions): Promise<GetGatewayPluginRateLimitingAdvancedResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("konnect:index/getGatewayPluginRateLimitingAdvanced:getGatewayPluginRateLimitingAdvanced", {
        "controlPlaneId": args.controlPlaneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGatewayPluginRateLimitingAdvanced.
 */
export interface GetGatewayPluginRateLimitingAdvancedArgs {
    controlPlaneId: string;
}

/**
 * A collection of values returned by getGatewayPluginRateLimitingAdvanced.
 */
export interface GetGatewayPluginRateLimitingAdvancedResult {
    readonly config: outputs.GetGatewayPluginRateLimitingAdvancedConfig;
    readonly consumer: outputs.GetGatewayPluginRateLimitingAdvancedConsumer;
    readonly consumerGroup: outputs.GetGatewayPluginRateLimitingAdvancedConsumerGroup;
    readonly controlPlaneId: string;
    readonly createdAt: number;
    readonly enabled: boolean;
    readonly id: string;
    readonly instanceName: string;
    readonly protocols: string[];
    readonly route: outputs.GetGatewayPluginRateLimitingAdvancedRoute;
    readonly service: outputs.GetGatewayPluginRateLimitingAdvancedService;
    readonly tags: string[];
    readonly updatedAt: number;
}
export function getGatewayPluginRateLimitingAdvancedOutput(args: GetGatewayPluginRateLimitingAdvancedOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGatewayPluginRateLimitingAdvancedResult> {
    return pulumi.output(args).apply((a: any) => getGatewayPluginRateLimitingAdvanced(a, opts))
}

/**
 * A collection of arguments for invoking getGatewayPluginRateLimitingAdvanced.
 */
export interface GetGatewayPluginRateLimitingAdvancedOutputArgs {
    controlPlaneId: pulumi.Input<string>;
}