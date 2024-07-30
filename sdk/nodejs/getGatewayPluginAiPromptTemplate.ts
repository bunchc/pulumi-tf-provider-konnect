// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getGatewayPluginAiPromptTemplate(args: GetGatewayPluginAiPromptTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetGatewayPluginAiPromptTemplateResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("konnect:index/getGatewayPluginAiPromptTemplate:getGatewayPluginAiPromptTemplate", {
        "controlPlaneId": args.controlPlaneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGatewayPluginAiPromptTemplate.
 */
export interface GetGatewayPluginAiPromptTemplateArgs {
    controlPlaneId: string;
}

/**
 * A collection of values returned by getGatewayPluginAiPromptTemplate.
 */
export interface GetGatewayPluginAiPromptTemplateResult {
    readonly config: outputs.GetGatewayPluginAiPromptTemplateConfig;
    readonly consumer: outputs.GetGatewayPluginAiPromptTemplateConsumer;
    readonly consumerGroup: outputs.GetGatewayPluginAiPromptTemplateConsumerGroup;
    readonly controlPlaneId: string;
    readonly createdAt: number;
    readonly enabled: boolean;
    readonly id: string;
    readonly instanceName: string;
    readonly protocols: string[];
    readonly route: outputs.GetGatewayPluginAiPromptTemplateRoute;
    readonly service: outputs.GetGatewayPluginAiPromptTemplateService;
    readonly tags: string[];
    readonly updatedAt: number;
}
export function getGatewayPluginAiPromptTemplateOutput(args: GetGatewayPluginAiPromptTemplateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGatewayPluginAiPromptTemplateResult> {
    return pulumi.output(args).apply((a: any) => getGatewayPluginAiPromptTemplate(a, opts))
}

/**
 * A collection of arguments for invoking getGatewayPluginAiPromptTemplate.
 */
export interface GetGatewayPluginAiPromptTemplateOutputArgs {
    controlPlaneId: pulumi.Input<string>;
}
