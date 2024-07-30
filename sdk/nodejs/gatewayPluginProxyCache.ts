// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * GatewayPluginProxyCache Resource
 */
export class GatewayPluginProxyCache extends pulumi.CustomResource {
    /**
     * Get an existing GatewayPluginProxyCache resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GatewayPluginProxyCacheState, opts?: pulumi.CustomResourceOptions): GatewayPluginProxyCache {
        return new GatewayPluginProxyCache(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'konnect:index/gatewayPluginProxyCache:GatewayPluginProxyCache';

    /**
     * Returns true if the given object is an instance of GatewayPluginProxyCache.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GatewayPluginProxyCache {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GatewayPluginProxyCache.__pulumiType;
    }

    public readonly config!: pulumi.Output<outputs.GatewayPluginProxyCacheConfig>;
    /**
     * If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
     */
    public readonly consumer!: pulumi.Output<outputs.GatewayPluginProxyCacheConsumer>;
    public readonly consumerGroup!: pulumi.Output<outputs.GatewayPluginProxyCacheConsumerGroup>;
    /**
     * The UUID of your control plane. This variable is available in the Konnect manager.
     */
    public readonly controlPlaneId!: pulumi.Output<string>;
    /**
     * Unix epoch when the resource was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<number>;
    /**
     * Whether the plugin is applied.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
     */
    public readonly protocols!: pulumi.Output<string[]>;
    /**
     * If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
     */
    public readonly route!: pulumi.Output<outputs.GatewayPluginProxyCacheRoute>;
    /**
     * If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
     */
    public readonly service!: pulumi.Output<outputs.GatewayPluginProxyCacheService>;
    /**
     * An optional set of strings associated with the Plugin for grouping and filtering.
     */
    public readonly tags!: pulumi.Output<string[]>;
    /**
     * Unix epoch when the resource was last updated.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<number>;

    /**
     * Create a GatewayPluginProxyCache resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GatewayPluginProxyCacheArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GatewayPluginProxyCacheArgs | GatewayPluginProxyCacheState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GatewayPluginProxyCacheState | undefined;
            resourceInputs["config"] = state ? state.config : undefined;
            resourceInputs["consumer"] = state ? state.consumer : undefined;
            resourceInputs["consumerGroup"] = state ? state.consumerGroup : undefined;
            resourceInputs["controlPlaneId"] = state ? state.controlPlaneId : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["protocols"] = state ? state.protocols : undefined;
            resourceInputs["route"] = state ? state.route : undefined;
            resourceInputs["service"] = state ? state.service : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as GatewayPluginProxyCacheArgs | undefined;
            if ((!args || args.controlPlaneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'controlPlaneId'");
            }
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["consumer"] = args ? args.consumer : undefined;
            resourceInputs["consumerGroup"] = args ? args.consumerGroup : undefined;
            resourceInputs["controlPlaneId"] = args ? args.controlPlaneId : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["protocols"] = args ? args.protocols : undefined;
            resourceInputs["route"] = args ? args.route : undefined;
            resourceInputs["service"] = args ? args.service : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GatewayPluginProxyCache.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GatewayPluginProxyCache resources.
 */
export interface GatewayPluginProxyCacheState {
    config?: pulumi.Input<inputs.GatewayPluginProxyCacheConfig>;
    /**
     * If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
     */
    consumer?: pulumi.Input<inputs.GatewayPluginProxyCacheConsumer>;
    consumerGroup?: pulumi.Input<inputs.GatewayPluginProxyCacheConsumerGroup>;
    /**
     * The UUID of your control plane. This variable is available in the Konnect manager.
     */
    controlPlaneId?: pulumi.Input<string>;
    /**
     * Unix epoch when the resource was created.
     */
    createdAt?: pulumi.Input<number>;
    /**
     * Whether the plugin is applied.
     */
    enabled?: pulumi.Input<boolean>;
    instanceName?: pulumi.Input<string>;
    /**
     * A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
     */
    protocols?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
     */
    route?: pulumi.Input<inputs.GatewayPluginProxyCacheRoute>;
    /**
     * If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
     */
    service?: pulumi.Input<inputs.GatewayPluginProxyCacheService>;
    /**
     * An optional set of strings associated with the Plugin for grouping and filtering.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Unix epoch when the resource was last updated.
     */
    updatedAt?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a GatewayPluginProxyCache resource.
 */
export interface GatewayPluginProxyCacheArgs {
    config?: pulumi.Input<inputs.GatewayPluginProxyCacheConfig>;
    /**
     * If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
     */
    consumer?: pulumi.Input<inputs.GatewayPluginProxyCacheConsumer>;
    consumerGroup?: pulumi.Input<inputs.GatewayPluginProxyCacheConsumerGroup>;
    /**
     * The UUID of your control plane. This variable is available in the Konnect manager.
     */
    controlPlaneId: pulumi.Input<string>;
    /**
     * Whether the plugin is applied.
     */
    enabled?: pulumi.Input<boolean>;
    instanceName?: pulumi.Input<string>;
    /**
     * A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
     */
    protocols?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
     */
    route?: pulumi.Input<inputs.GatewayPluginProxyCacheRoute>;
    /**
     * If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
     */
    service?: pulumi.Input<inputs.GatewayPluginProxyCacheService>;
    /**
     * An optional set of strings associated with the Plugin for grouping and filtering.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}
