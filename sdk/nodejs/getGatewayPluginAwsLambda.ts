// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getGatewayPluginAwsLambda(args: GetGatewayPluginAwsLambdaArgs, opts?: pulumi.InvokeOptions): Promise<GetGatewayPluginAwsLambdaResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("konnect:index/getGatewayPluginAwsLambda:getGatewayPluginAwsLambda", {
        "controlPlaneId": args.controlPlaneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGatewayPluginAwsLambda.
 */
export interface GetGatewayPluginAwsLambdaArgs {
    controlPlaneId: string;
}

/**
 * A collection of values returned by getGatewayPluginAwsLambda.
 */
export interface GetGatewayPluginAwsLambdaResult {
    readonly config: outputs.GetGatewayPluginAwsLambdaConfig;
    readonly consumer: outputs.GetGatewayPluginAwsLambdaConsumer;
    readonly consumerGroup: outputs.GetGatewayPluginAwsLambdaConsumerGroup;
    readonly controlPlaneId: string;
    readonly createdAt: number;
    readonly enabled: boolean;
    readonly id: string;
    readonly instanceName: string;
    readonly protocols: string[];
    readonly route: outputs.GetGatewayPluginAwsLambdaRoute;
    readonly service: outputs.GetGatewayPluginAwsLambdaService;
    readonly tags: string[];
    readonly updatedAt: number;
}
export function getGatewayPluginAwsLambdaOutput(args: GetGatewayPluginAwsLambdaOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGatewayPluginAwsLambdaResult> {
    return pulumi.output(args).apply((a: any) => getGatewayPluginAwsLambda(a, opts))
}

/**
 * A collection of arguments for invoking getGatewayPluginAwsLambda.
 */
export interface GetGatewayPluginAwsLambdaOutputArgs {
    controlPlaneId: pulumi.Input<string>;
}