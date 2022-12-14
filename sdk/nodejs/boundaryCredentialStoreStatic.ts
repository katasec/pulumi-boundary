// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The static credential store resource allows you to configure a Boundary static credential store.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as boundary from "@katasec/boundary";
 *
 * const org = new boundary.BoundaryScope("org", {
 *     description: "global scope",
 *     scopeId: "global",
 *     autoCreateAdminRole: true,
 *     autoCreateDefaultRole: true,
 * });
 * const project = new boundary.BoundaryScope("project", {
 *     description: "My first scope!",
 *     scopeId: org.id,
 *     autoCreateAdminRole: true,
 * });
 * const example = new boundary.BoundaryCredentialStoreStatic("example", {
 *     description: "My first static credential store!",
 *     scopeId: project.id,
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import boundary:index/boundaryCredentialStoreStatic:BoundaryCredentialStoreStatic example_static_credential_store <my-id>
 * ```
 */
export class BoundaryCredentialStoreStatic extends pulumi.CustomResource {
    /**
     * Get an existing BoundaryCredentialStoreStatic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BoundaryCredentialStoreStaticState, opts?: pulumi.CustomResourceOptions): BoundaryCredentialStoreStatic {
        return new BoundaryCredentialStoreStatic(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'boundary:index/boundaryCredentialStoreStatic:BoundaryCredentialStoreStatic';

    /**
     * Returns true if the given object is an instance of BoundaryCredentialStoreStatic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BoundaryCredentialStoreStatic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BoundaryCredentialStoreStatic.__pulumiType;
    }

    /**
     * The static credential store description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The static credential store name. Defaults to the resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The scope for this credential store.
     */
    public readonly scopeId!: pulumi.Output<string>;

    /**
     * Create a BoundaryCredentialStoreStatic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BoundaryCredentialStoreStaticArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BoundaryCredentialStoreStaticArgs | BoundaryCredentialStoreStaticState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BoundaryCredentialStoreStaticState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["scopeId"] = state ? state.scopeId : undefined;
        } else {
            const args = argsOrState as BoundaryCredentialStoreStaticArgs | undefined;
            if ((!args || args.scopeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scopeId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["scopeId"] = args ? args.scopeId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BoundaryCredentialStoreStatic.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BoundaryCredentialStoreStatic resources.
 */
export interface BoundaryCredentialStoreStaticState {
    /**
     * The static credential store description.
     */
    description?: pulumi.Input<string>;
    /**
     * The static credential store name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The scope for this credential store.
     */
    scopeId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BoundaryCredentialStoreStatic resource.
 */
export interface BoundaryCredentialStoreStaticArgs {
    /**
     * The static credential store description.
     */
    description?: pulumi.Input<string>;
    /**
     * The static credential store name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The scope for this credential store.
     */
    scopeId: pulumi.Input<string>;
}
