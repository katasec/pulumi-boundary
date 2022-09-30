// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The auth method resource allows you to configure a Boundary auth_method.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as boundary from "@katasec/boundary";
 *
 * const org = new boundary.BoundaryScope("org", {
 *     description: "My first scope!",
 *     scopeId: "global",
 *     autoCreateAdminRole: true,
 *     autoCreateDefaultRole: true,
 * });
 * const password = new boundary.BoundaryAuthMethod("password", {
 *     scopeId: org.id,
 *     type: "password",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import boundary:index/boundaryAuthMethod:BoundaryAuthMethod foo <my-id>
 * ```
 */
export class BoundaryAuthMethod extends pulumi.CustomResource {
    /**
     * Get an existing BoundaryAuthMethod resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BoundaryAuthMethodState, opts?: pulumi.CustomResourceOptions): BoundaryAuthMethod {
        return new BoundaryAuthMethod(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'boundary:index/boundaryAuthMethod:BoundaryAuthMethod';

    /**
     * Returns true if the given object is an instance of BoundaryAuthMethod.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BoundaryAuthMethod {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BoundaryAuthMethod.__pulumiType;
    }

    /**
     * The auth method description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The minimum login name length.
     *
     * @deprecated Will be removed in favor of using attributes parameter
     */
    public readonly minLoginNameLength!: pulumi.Output<number>;
    /**
     * The minimum password length.
     *
     * @deprecated Will be removed in favor of using attributes parameter
     */
    public readonly minPasswordLength!: pulumi.Output<number>;
    /**
     * The auth method name. Defaults to the resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The scope ID.
     */
    public readonly scopeId!: pulumi.Output<string>;
    /**
     * The resource type.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a BoundaryAuthMethod resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BoundaryAuthMethodArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BoundaryAuthMethodArgs | BoundaryAuthMethodState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BoundaryAuthMethodState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["minLoginNameLength"] = state ? state.minLoginNameLength : undefined;
            resourceInputs["minPasswordLength"] = state ? state.minPasswordLength : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["scopeId"] = state ? state.scopeId : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as BoundaryAuthMethodArgs | undefined;
            if ((!args || args.scopeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scopeId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["minLoginNameLength"] = args ? args.minLoginNameLength : undefined;
            resourceInputs["minPasswordLength"] = args ? args.minPasswordLength : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["scopeId"] = args ? args.scopeId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BoundaryAuthMethod.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BoundaryAuthMethod resources.
 */
export interface BoundaryAuthMethodState {
    /**
     * The auth method description.
     */
    description?: pulumi.Input<string>;
    /**
     * The minimum login name length.
     *
     * @deprecated Will be removed in favor of using attributes parameter
     */
    minLoginNameLength?: pulumi.Input<number>;
    /**
     * The minimum password length.
     *
     * @deprecated Will be removed in favor of using attributes parameter
     */
    minPasswordLength?: pulumi.Input<number>;
    /**
     * The auth method name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The scope ID.
     */
    scopeId?: pulumi.Input<string>;
    /**
     * The resource type.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BoundaryAuthMethod resource.
 */
export interface BoundaryAuthMethodArgs {
    /**
     * The auth method description.
     */
    description?: pulumi.Input<string>;
    /**
     * The minimum login name length.
     *
     * @deprecated Will be removed in favor of using attributes parameter
     */
    minLoginNameLength?: pulumi.Input<number>;
    /**
     * The minimum password length.
     *
     * @deprecated Will be removed in favor of using attributes parameter
     */
    minPasswordLength?: pulumi.Input<number>;
    /**
     * The auth method name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The scope ID.
     */
    scopeId: pulumi.Input<string>;
    /**
     * The resource type.
     */
    type: pulumi.Input<string>;
}
