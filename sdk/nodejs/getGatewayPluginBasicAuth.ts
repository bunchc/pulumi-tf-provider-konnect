// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getGatewayPluginBasicAuth(args: GetGatewayPluginBasicAuthArgs, opts?: pulumi.InvokeOptions): Promise<GetGatewayPluginBasicAuthResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("konnect:index/getGatewayPluginBasicAuth:getGatewayPluginBasicAuth", {
        "controlPlaneId": args.controlPlaneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGatewayPluginBasicAuth.
 */
export interface GetGatewayPluginBasicAuthArgs {
    controlPlaneId: string;
}

/**
 * A collection of values returned by getGatewayPluginBasicAuth.
 */
export interface GetGatewayPluginBasicAuthResult {
    readonly config: outputs.GetGatewayPluginBasicAuthConfig;
    readonly consumer: outputs.GetGatewayPluginBasicAuthConsumer;
    readonly consumerGroup: outputs.GetGatewayPluginBasicAuthConsumerGroup;
    readonly controlPlaneId: string;
    readonly createdAt: number;
    readonly enabled: boolean;
    readonly id: string;
    readonly instanceName: string;
    readonly protocols: string[];
    readonly route: outputs.GetGatewayPluginBasicAuthRoute;
    readonly service: outputs.GetGatewayPluginBasicAuthService;
    readonly tags: string[];
    readonly updatedAt: number;
}
export function getGatewayPluginBasicAuthOutput(args: GetGatewayPluginBasicAuthOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGatewayPluginBasicAuthResult> {
    return pulumi.output(args).apply((a: any) => getGatewayPluginBasicAuth(a, opts))
}

/**
 * A collection of arguments for invoking getGatewayPluginBasicAuth.
 */
export interface GetGatewayPluginBasicAuthOutputArgs {
    controlPlaneId: pulumi.Input<string>;
}