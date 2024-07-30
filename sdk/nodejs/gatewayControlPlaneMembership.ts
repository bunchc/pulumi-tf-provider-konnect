// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * GatewayControlPlaneMembership Resource
 */
export class GatewayControlPlaneMembership extends pulumi.CustomResource {
    /**
     * Get an existing GatewayControlPlaneMembership resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GatewayControlPlaneMembershipState, opts?: pulumi.CustomResourceOptions): GatewayControlPlaneMembership {
        return new GatewayControlPlaneMembership(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'konnect:index/gatewayControlPlaneMembership:GatewayControlPlaneMembership';

    /**
     * Returns true if the given object is an instance of GatewayControlPlaneMembership.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GatewayControlPlaneMembership {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GatewayControlPlaneMembership.__pulumiType;
    }

    /**
     * Requires replacement if changed.
     */
    public readonly members!: pulumi.Output<outputs.GatewayControlPlaneMembershipMember[] | undefined>;

    /**
     * Create a GatewayControlPlaneMembership resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: GatewayControlPlaneMembershipArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GatewayControlPlaneMembershipArgs | GatewayControlPlaneMembershipState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GatewayControlPlaneMembershipState | undefined;
            resourceInputs["members"] = state ? state.members : undefined;
        } else {
            const args = argsOrState as GatewayControlPlaneMembershipArgs | undefined;
            resourceInputs["members"] = args ? args.members : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GatewayControlPlaneMembership.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GatewayControlPlaneMembership resources.
 */
export interface GatewayControlPlaneMembershipState {
    /**
     * Requires replacement if changed.
     */
    members?: pulumi.Input<pulumi.Input<inputs.GatewayControlPlaneMembershipMember>[]>;
}

/**
 * The set of arguments for constructing a GatewayControlPlaneMembership resource.
 */
export interface GatewayControlPlaneMembershipArgs {
    /**
     * Requires replacement if changed.
     */
    members?: pulumi.Input<pulumi.Input<inputs.GatewayControlPlaneMembershipMember>[]>;
}
