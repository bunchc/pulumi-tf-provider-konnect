// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * APIProductDocument Resource
 */
export class ApiProductDocument extends pulumi.CustomResource {
    /**
     * Get an existing ApiProductDocument resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApiProductDocumentState, opts?: pulumi.CustomResourceOptions): ApiProductDocument {
        return new ApiProductDocument(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'konnect:index/apiProductDocument:ApiProductDocument';

    /**
     * Returns true if the given object is an instance of ApiProductDocument.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApiProductDocument {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApiProductDocument.__pulumiType;
    }

    /**
     * The API product identifier
     */
    public readonly apiProductId!: pulumi.Output<string>;
    /**
     * Can be markdown string content or base64 encoded string
     */
    public readonly content!: pulumi.Output<string>;
    /**
     * An ISO-8601 timestamp representation of entity creation date.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * metadata of the document
     */
    public readonly metadata!: pulumi.Output<outputs.ApiProductDocumentMetadata>;
    /**
     * parent document id
     */
    public readonly parentDocumentId!: pulumi.Output<string>;
    /**
     * document slug. must be unique accross documents belonging to an api product
     */
    public readonly slug!: pulumi.Output<string>;
    /**
     * document publish status. must be one of ["published", "unpublished"]
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * document title
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * An ISO-8601 timestamp representation of entity update date.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a ApiProductDocument resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiProductDocumentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApiProductDocumentArgs | ApiProductDocumentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApiProductDocumentState | undefined;
            resourceInputs["apiProductId"] = state ? state.apiProductId : undefined;
            resourceInputs["content"] = state ? state.content : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["parentDocumentId"] = state ? state.parentDocumentId : undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["title"] = state ? state.title : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as ApiProductDocumentArgs | undefined;
            if ((!args || args.apiProductId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiProductId'");
            }
            if ((!args || args.slug === undefined) && !opts.urn) {
                throw new Error("Missing required property 'slug'");
            }
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            if ((!args || args.title === undefined) && !opts.urn) {
                throw new Error("Missing required property 'title'");
            }
            resourceInputs["apiProductId"] = args ? args.apiProductId : undefined;
            resourceInputs["content"] = args ? args.content : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["parentDocumentId"] = args ? args.parentDocumentId : undefined;
            resourceInputs["slug"] = args ? args.slug : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApiProductDocument.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApiProductDocument resources.
 */
export interface ApiProductDocumentState {
    /**
     * The API product identifier
     */
    apiProductId?: pulumi.Input<string>;
    /**
     * Can be markdown string content or base64 encoded string
     */
    content?: pulumi.Input<string>;
    /**
     * An ISO-8601 timestamp representation of entity creation date.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * metadata of the document
     */
    metadata?: pulumi.Input<inputs.ApiProductDocumentMetadata>;
    /**
     * parent document id
     */
    parentDocumentId?: pulumi.Input<string>;
    /**
     * document slug. must be unique accross documents belonging to an api product
     */
    slug?: pulumi.Input<string>;
    /**
     * document publish status. must be one of ["published", "unpublished"]
     */
    status?: pulumi.Input<string>;
    /**
     * document title
     */
    title?: pulumi.Input<string>;
    /**
     * An ISO-8601 timestamp representation of entity update date.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApiProductDocument resource.
 */
export interface ApiProductDocumentArgs {
    /**
     * The API product identifier
     */
    apiProductId: pulumi.Input<string>;
    /**
     * Can be markdown string content or base64 encoded string
     */
    content?: pulumi.Input<string>;
    /**
     * metadata of the document
     */
    metadata?: pulumi.Input<inputs.ApiProductDocumentMetadata>;
    /**
     * parent document id
     */
    parentDocumentId?: pulumi.Input<string>;
    /**
     * document slug. must be unique accross documents belonging to an api product
     */
    slug: pulumi.Input<string>;
    /**
     * document publish status. must be one of ["published", "unpublished"]
     */
    status: pulumi.Input<string>;
    /**
     * document title
     */
    title: pulumi.Input<string>;
}