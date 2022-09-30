// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Deprecated: use `boundary.BoundaryHostStatic` instead.
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
 * const project = new boundary.BoundaryScope("project", {
 *     description: "My first scope!",
 *     scopeId: org.id,
 *     autoCreateAdminRole: true,
 * });
 * const static = new boundary.BoundaryHostCatalog("static", {
 *     description: "My first host catalog!",
 *     type: "static",
 *     scopeId: project.id,
 * });
 * const example = new boundary.BoundaryHost("example", {
 *     type: "static",
 *     description: "My first host!",
 *     address: "10.0.0.1",
 *     hostCatalogId: static.id,
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import boundary:index/boundaryHost:BoundaryHost foo <my-id>
 * ```
 */
export class BoundaryHost extends pulumi.CustomResource {
    /**
     * Get an existing BoundaryHost resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BoundaryHostState, opts?: pulumi.CustomResourceOptions): BoundaryHost {
        return new BoundaryHost(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'boundary:index/boundaryHost:BoundaryHost';

    /**
     * Returns true if the given object is an instance of BoundaryHost.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BoundaryHost {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BoundaryHost.__pulumiType;
    }

    /**
     * The static address of the host resource as `<IP>` (note: port assignment occurs in the target resource definition, do not add :port here) or a domain name.
     */
    public readonly address!: pulumi.Output<string | undefined>;
    /**
     * The host description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly hostCatalogId!: pulumi.Output<string>;
    /**
     * The host name. Defaults to the resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The type of host
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a BoundaryHost resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BoundaryHostArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BoundaryHostArgs | BoundaryHostState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BoundaryHostState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["hostCatalogId"] = state ? state.hostCatalogId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as BoundaryHostArgs | undefined;
            if ((!args || args.hostCatalogId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostCatalogId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["hostCatalogId"] = args ? args.hostCatalogId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BoundaryHost.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BoundaryHost resources.
 */
export interface BoundaryHostState {
    /**
     * The static address of the host resource as `<IP>` (note: port assignment occurs in the target resource definition, do not add :port here) or a domain name.
     */
    address?: pulumi.Input<string>;
    /**
     * The host description.
     */
    description?: pulumi.Input<string>;
    hostCatalogId?: pulumi.Input<string>;
    /**
     * The host name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The type of host
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BoundaryHost resource.
 */
export interface BoundaryHostArgs {
    /**
     * The static address of the host resource as `<IP>` (note: port assignment occurs in the target resource definition, do not add :port here) or a domain name.
     */
    address?: pulumi.Input<string>;
    /**
     * The host description.
     */
    description?: pulumi.Input<string>;
    hostCatalogId: pulumi.Input<string>;
    /**
     * The host name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The type of host
     */
    type: pulumi.Input<string>;
}
