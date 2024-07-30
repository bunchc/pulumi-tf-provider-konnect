// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getGatewayPluginStatsd(args: GetGatewayPluginStatsdArgs, opts?: pulumi.InvokeOptions): Promise<GetGatewayPluginStatsdResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("konnect:index/getGatewayPluginStatsd:getGatewayPluginStatsd", {
        "controlPlaneId": args.controlPlaneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGatewayPluginStatsd.
 */
export interface GetGatewayPluginStatsdArgs {
    controlPlaneId: string;
}

/**
 * A collection of values returned by getGatewayPluginStatsd.
 */
export interface GetGatewayPluginStatsdResult {
    readonly config: outputs.GetGatewayPluginStatsdConfig;
    readonly consumer: outputs.GetGatewayPluginStatsdConsumer;
    readonly consumerGroup: outputs.GetGatewayPluginStatsdConsumerGroup;
    readonly controlPlaneId: string;
    readonly createdAt: number;
    readonly enabled: boolean;
    readonly id: string;
    readonly instanceName: string;
    readonly protocols: string[];
    readonly route: outputs.GetGatewayPluginStatsdRoute;
    readonly service: outputs.GetGatewayPluginStatsdService;
    readonly tags: string[];
    readonly updatedAt: number;
}
export function getGatewayPluginStatsdOutput(args: GetGatewayPluginStatsdOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGatewayPluginStatsdResult> {
    return pulumi.output(args).apply((a: any) => getGatewayPluginStatsd(a, opts))
}

/**
 * A collection of arguments for invoking getGatewayPluginStatsd.
 */
export interface GetGatewayPluginStatsdOutputArgs {
    controlPlaneId: pulumi.Input<string>;
}
