// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The account resource allows you to configure a Boundary account.
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
 * const jeff = new boundary.BoundaryAccountPassword("jeff", {
 *     authMethodId: password.id,
 *     type: "password",
 *     loginName: "jeff",
 *     password: `$uper$ecure`,
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import boundary:index/boundaryAccountPassword:BoundaryAccountPassword foo <my-id>
 * ```
 */
export class BoundaryAccountPassword extends pulumi.CustomResource {
    /**
     * Get an existing BoundaryAccountPassword resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BoundaryAccountPasswordState, opts?: pulumi.CustomResourceOptions): BoundaryAccountPassword {
        return new BoundaryAccountPassword(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'boundary:index/boundaryAccountPassword:BoundaryAccountPassword';

    /**
     * Returns true if the given object is an instance of BoundaryAccountPassword.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BoundaryAccountPassword {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BoundaryAccountPassword.__pulumiType;
    }

    /**
     * The resource ID for the auth method.
     */
    public readonly authMethodId!: pulumi.Output<string>;
    /**
     * The account description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The login name for this account.
     */
    public readonly loginName!: pulumi.Output<string | undefined>;
    /**
     * The account name. Defaults to the resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The account password. Only set on create, changes will not be reflected when updating account.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * The resource type.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a BoundaryAccountPassword resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BoundaryAccountPasswordArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BoundaryAccountPasswordArgs | BoundaryAccountPasswordState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BoundaryAccountPasswordState | undefined;
            resourceInputs["authMethodId"] = state ? state.authMethodId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["loginName"] = state ? state.loginName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as BoundaryAccountPasswordArgs | undefined;
            if ((!args || args.authMethodId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authMethodId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["authMethodId"] = args ? args.authMethodId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["loginName"] = args ? args.loginName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BoundaryAccountPassword.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BoundaryAccountPassword resources.
 */
export interface BoundaryAccountPasswordState {
    /**
     * The resource ID for the auth method.
     */
    authMethodId?: pulumi.Input<string>;
    /**
     * The account description.
     */
    description?: pulumi.Input<string>;
    /**
     * The login name for this account.
     */
    loginName?: pulumi.Input<string>;
    /**
     * The account name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The account password. Only set on create, changes will not be reflected when updating account.
     */
    password?: pulumi.Input<string>;
    /**
     * The resource type.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BoundaryAccountPassword resource.
 */
export interface BoundaryAccountPasswordArgs {
    /**
     * The resource ID for the auth method.
     */
    authMethodId: pulumi.Input<string>;
    /**
     * The account description.
     */
    description?: pulumi.Input<string>;
    /**
     * The login name for this account.
     */
    loginName?: pulumi.Input<string>;
    /**
     * The account name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The account password. Only set on create, changes will not be reflected when updating account.
     */
    password?: pulumi.Input<string>;
    /**
     * The resource type.
     */
    type: pulumi.Input<string>;
}