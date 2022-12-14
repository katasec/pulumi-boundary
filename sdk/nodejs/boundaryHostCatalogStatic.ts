// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The static host catalog resource allows you to configure a Boundary static-type host catalog. Host catalogs are always part of a project, so a project resource should be used inline or you should have the project ID in hand to successfully configure a host catalog.
 */
export class BoundaryHostCatalogStatic extends pulumi.CustomResource {
    /**
     * Get an existing BoundaryHostCatalogStatic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BoundaryHostCatalogStaticState, opts?: pulumi.CustomResourceOptions): BoundaryHostCatalogStatic {
        return new BoundaryHostCatalogStatic(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'boundary:index/boundaryHostCatalogStatic:BoundaryHostCatalogStatic';

    /**
     * Returns true if the given object is an instance of BoundaryHostCatalogStatic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BoundaryHostCatalogStatic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BoundaryHostCatalogStatic.__pulumiType;
    }

    /**
     * The host catalog description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The host catalog name. Defaults to the resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The scope ID in which the resource is created.
     */
    public readonly scopeId!: pulumi.Output<string>;

    /**
     * Create a BoundaryHostCatalogStatic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BoundaryHostCatalogStaticArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BoundaryHostCatalogStaticArgs | BoundaryHostCatalogStaticState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BoundaryHostCatalogStaticState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["scopeId"] = state ? state.scopeId : undefined;
        } else {
            const args = argsOrState as BoundaryHostCatalogStaticArgs | undefined;
            if ((!args || args.scopeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scopeId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["scopeId"] = args ? args.scopeId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BoundaryHostCatalogStatic.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BoundaryHostCatalogStatic resources.
 */
export interface BoundaryHostCatalogStaticState {
    /**
     * The host catalog description.
     */
    description?: pulumi.Input<string>;
    /**
     * The host catalog name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The scope ID in which the resource is created.
     */
    scopeId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BoundaryHostCatalogStatic resource.
 */
export interface BoundaryHostCatalogStaticArgs {
    /**
     * The host catalog description.
     */
    description?: pulumi.Input<string>;
    /**
     * The host catalog name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The scope ID in which the resource is created.
     */
    scopeId: pulumi.Input<string>;
}
