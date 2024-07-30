// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * APIProductVersion Resource
 */
export class ApiProductVersion extends pulumi.CustomResource {
    /**
     * Get an existing ApiProductVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApiProductVersionState, opts?: pulumi.CustomResourceOptions): ApiProductVersion {
        return new ApiProductVersion(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'konnect:index/apiProductVersion:ApiProductVersion';

    /**
     * Returns true if the given object is an instance of ApiProductVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApiProductVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApiProductVersion.__pulumiType;
    }

    /**
     * The API Product ID
     */
    public readonly apiProductId!: pulumi.Output<string>;
    /**
     * An ISO-8601 timestamp representation of entity creation date.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Indicates if the version of the API product is deprecated. Applies deprecation or removes deprecation from all related portal product versions. This field is deprecated: Use [PortalProductVersion.deprecated](https://docs.konghq.com/konnect/api/portal-management/latest/#/Portal%20Product%20Versions/create-portal-product-version) instead.
     */
    public readonly deprecated!: pulumi.Output<boolean>;
    public readonly gatewayService!: pulumi.Output<outputs.ApiProductVersionGatewayService>;
    /**
     * The version name of the API product version.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The list of portals which this API product version is configured for
     */
    public /*out*/ readonly portals!: pulumi.Output<outputs.ApiProductVersionPortal[]>;
    /**
     * An ISO-8601 timestamp representation of entity update date.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a ApiProductVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiProductVersionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApiProductVersionArgs | ApiProductVersionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApiProductVersionState | undefined;
            resourceInputs["apiProductId"] = state ? state.apiProductId : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["deprecated"] = state ? state.deprecated : undefined;
            resourceInputs["gatewayService"] = state ? state.gatewayService : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["portals"] = state ? state.portals : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as ApiProductVersionArgs | undefined;
            if ((!args || args.apiProductId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiProductId'");
            }
            resourceInputs["apiProductId"] = args ? args.apiProductId : undefined;
            resourceInputs["deprecated"] = args ? args.deprecated : undefined;
            resourceInputs["gatewayService"] = args ? args.gatewayService : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["portals"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApiProductVersion.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApiProductVersion resources.
 */
export interface ApiProductVersionState {
    /**
     * The API Product ID
     */
    apiProductId?: pulumi.Input<string>;
    /**
     * An ISO-8601 timestamp representation of entity creation date.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Indicates if the version of the API product is deprecated. Applies deprecation or removes deprecation from all related portal product versions. This field is deprecated: Use [PortalProductVersion.deprecated](https://docs.konghq.com/konnect/api/portal-management/latest/#/Portal%20Product%20Versions/create-portal-product-version) instead.
     */
    deprecated?: pulumi.Input<boolean>;
    gatewayService?: pulumi.Input<inputs.ApiProductVersionGatewayService>;
    /**
     * The version name of the API product version.
     */
    name?: pulumi.Input<string>;
    /**
     * The list of portals which this API product version is configured for
     */
    portals?: pulumi.Input<pulumi.Input<inputs.ApiProductVersionPortal>[]>;
    /**
     * An ISO-8601 timestamp representation of entity update date.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApiProductVersion resource.
 */
export interface ApiProductVersionArgs {
    /**
     * The API Product ID
     */
    apiProductId: pulumi.Input<string>;
    /**
     * Indicates if the version of the API product is deprecated. Applies deprecation or removes deprecation from all related portal product versions. This field is deprecated: Use [PortalProductVersion.deprecated](https://docs.konghq.com/konnect/api/portal-management/latest/#/Portal%20Product%20Versions/create-portal-product-version) instead.
     */
    deprecated?: pulumi.Input<boolean>;
    gatewayService?: pulumi.Input<inputs.ApiProductVersionGatewayService>;
    /**
     * The version name of the API product version.
     */
    name?: pulumi.Input<string>;
}
