// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getCloudGatewayConfiguration(opts?: pulumi.InvokeOptions): Promise<GetCloudGatewayConfigurationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("konnect:index/getCloudGatewayConfiguration:getCloudGatewayConfiguration", {
    }, opts);
}

/**
 * A collection of values returned by getCloudGatewayConfiguration.
 */
export interface GetCloudGatewayConfigurationResult {
    readonly apiAccess: string;
    readonly controlPlaneGeo: string;
    readonly controlPlaneId: string;
    readonly createdAt: string;
    readonly dataplaneGroupConfigs: outputs.GetCloudGatewayConfigurationDataplaneGroupConfig[];
    readonly dataplaneGroups: outputs.GetCloudGatewayConfigurationDataplaneGroup[];
    readonly entityVersion: number;
    readonly id: string;
    readonly updatedAt: string;
    readonly version: string;
}
export function getCloudGatewayConfigurationOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetCloudGatewayConfigurationResult> {
    return pulumi.output(getCloudGatewayConfiguration(opts))
}