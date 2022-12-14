// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The role resource allows you to configure a Boundary role.
 *
 * ## Example Usage
 *
 * Basic usage:
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
 * const example = new boundary.BoundaryRole("example", {
 *     description: "My first role!",
 *     scopeId: org.id,
 * });
 * ```
 *
 * Usage with a user resource:
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
 * const foo = new boundary.BoundaryUser("foo", {scopeId: org.id});
 * const bar = new boundary.BoundaryUser("bar", {scopeId: org.id});
 * const example = new boundary.BoundaryRole("example", {
 *     description: "My first role!",
 *     principalIds: [
 *         foo.id,
 *         bar.id,
 *     ],
 *     scopeId: org.id,
 * });
 * ```
 *
 * Usage with user and grants resource:
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
 * const readonlyBoundaryUser = new boundary.BoundaryUser("readonlyBoundaryUser", {
 *     description: "A readonly user",
 *     scopeId: org.id,
 * });
 * const readonlyBoundaryRole = new boundary.BoundaryRole("readonlyBoundaryRole", {
 *     description: "A readonly role",
 *     principalIds: [readonlyBoundaryUser.id],
 *     grantStrings: ["id=*;type=*;actions=read"],
 *     scopeId: org.id,
 * });
 * ```
 *
 * Usage for a project-specific role:
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
 * const project = new boundary.BoundaryScope("project", {
 *     description: "My first scope!",
 *     scopeId: org.id,
 *     autoCreateAdminRole: true,
 * });
 * const readonlyBoundaryUser = new boundary.BoundaryUser("readonlyBoundaryUser", {
 *     description: "A readonly user",
 *     scopeId: org.id,
 * });
 * const readonlyBoundaryRole = new boundary.BoundaryRole("readonlyBoundaryRole", {
 *     description: "A readonly role",
 *     principalIds: [readonlyBoundaryUser.id],
 *     grantStrings: ["id=*;type=*;actions=read"],
 *     scopeId: project.id,
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import boundary:index/boundaryRole:BoundaryRole foo <my-id>
 * ```
 */
export class BoundaryRole extends pulumi.CustomResource {
    /**
     * Get an existing BoundaryRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BoundaryRoleState, opts?: pulumi.CustomResourceOptions): BoundaryRole {
        return new BoundaryRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'boundary:index/boundaryRole:BoundaryRole';

    /**
     * Returns true if the given object is an instance of BoundaryRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BoundaryRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BoundaryRole.__pulumiType;
    }

    /**
     * The role description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly grantScopeId!: pulumi.Output<string>;
    /**
     * A list of stringified grants for the role.
     */
    public readonly grantStrings!: pulumi.Output<string[] | undefined>;
    /**
     * The role name. Defaults to the resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of principal (user or group) IDs to add as principals on the role.
     */
    public readonly principalIds!: pulumi.Output<string[] | undefined>;
    /**
     * The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
     */
    public readonly scopeId!: pulumi.Output<string>;

    /**
     * Create a BoundaryRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BoundaryRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BoundaryRoleArgs | BoundaryRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BoundaryRoleState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["grantScopeId"] = state ? state.grantScopeId : undefined;
            resourceInputs["grantStrings"] = state ? state.grantStrings : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["principalIds"] = state ? state.principalIds : undefined;
            resourceInputs["scopeId"] = state ? state.scopeId : undefined;
        } else {
            const args = argsOrState as BoundaryRoleArgs | undefined;
            if ((!args || args.scopeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scopeId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["grantScopeId"] = args ? args.grantScopeId : undefined;
            resourceInputs["grantStrings"] = args ? args.grantStrings : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["principalIds"] = args ? args.principalIds : undefined;
            resourceInputs["scopeId"] = args ? args.scopeId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BoundaryRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BoundaryRole resources.
 */
export interface BoundaryRoleState {
    /**
     * The role description.
     */
    description?: pulumi.Input<string>;
    grantScopeId?: pulumi.Input<string>;
    /**
     * A list of stringified grants for the role.
     */
    grantStrings?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The role name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of principal (user or group) IDs to add as principals on the role.
     */
    principalIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
     */
    scopeId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BoundaryRole resource.
 */
export interface BoundaryRoleArgs {
    /**
     * The role description.
     */
    description?: pulumi.Input<string>;
    grantScopeId?: pulumi.Input<string>;
    /**
     * A list of stringified grants for the role.
     */
    grantStrings?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The role name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of principal (user or group) IDs to add as principals on the role.
     */
    principalIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
     */
    scopeId: pulumi.Input<string>;
}
