// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The hostSetPlugin resource allows you to configure a Boundary host set. Host sets are always part of a host catalog, so a host catalog resource should be used inline or you should have the host catalog ID in hand to successfully configure a host set.
 */
export class BoundaryHostSetPlugin extends pulumi.CustomResource {
    /**
     * Get an existing BoundaryHostSetPlugin resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BoundaryHostSetPluginState, opts?: pulumi.CustomResourceOptions): BoundaryHostSetPlugin {
        return new BoundaryHostSetPlugin(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'boundary:index/boundaryHostSetPlugin:BoundaryHostSetPlugin';

    /**
     * Returns true if the given object is an instance of BoundaryHostSetPlugin.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BoundaryHostSetPlugin {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BoundaryHostSetPlugin.__pulumiType;
    }

    /**
     * The attributes for the host set. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host set.
     */
    public readonly attributesJson!: pulumi.Output<string | undefined>;
    /**
     * The host set description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The catalog for the host set.
     */
    public readonly hostCatalogId!: pulumi.Output<string>;
    /**
     * The host set name. Defaults to the resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ordered list of preferred endpoints.
     */
    public readonly preferredEndpoints!: pulumi.Output<string[] | undefined>;
    /**
     * The value to set for the sync interval seconds.
     */
    public readonly syncIntervalSeconds!: pulumi.Output<number | undefined>;
    /**
     * The type of host set
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a BoundaryHostSetPlugin resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BoundaryHostSetPluginArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BoundaryHostSetPluginArgs | BoundaryHostSetPluginState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BoundaryHostSetPluginState | undefined;
            resourceInputs["attributesJson"] = state ? state.attributesJson : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["hostCatalogId"] = state ? state.hostCatalogId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["preferredEndpoints"] = state ? state.preferredEndpoints : undefined;
            resourceInputs["syncIntervalSeconds"] = state ? state.syncIntervalSeconds : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as BoundaryHostSetPluginArgs | undefined;
            if ((!args || args.hostCatalogId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostCatalogId'");
            }
            resourceInputs["attributesJson"] = args ? args.attributesJson : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["hostCatalogId"] = args ? args.hostCatalogId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["preferredEndpoints"] = args ? args.preferredEndpoints : undefined;
            resourceInputs["syncIntervalSeconds"] = args ? args.syncIntervalSeconds : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BoundaryHostSetPlugin.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BoundaryHostSetPlugin resources.
 */
export interface BoundaryHostSetPluginState {
    /**
     * The attributes for the host set. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host set.
     */
    attributesJson?: pulumi.Input<string>;
    /**
     * The host set description.
     */
    description?: pulumi.Input<string>;
    /**
     * The catalog for the host set.
     */
    hostCatalogId?: pulumi.Input<string>;
    /**
     * The host set name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The ordered list of preferred endpoints.
     */
    preferredEndpoints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The value to set for the sync interval seconds.
     */
    syncIntervalSeconds?: pulumi.Input<number>;
    /**
     * The type of host set
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BoundaryHostSetPlugin resource.
 */
export interface BoundaryHostSetPluginArgs {
    /**
     * The attributes for the host set. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host set.
     */
    attributesJson?: pulumi.Input<string>;
    /**
     * The host set description.
     */
    description?: pulumi.Input<string>;
    /**
     * The catalog for the host set.
     */
    hostCatalogId: pulumi.Input<string>;
    /**
     * The host set name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The ordered list of preferred endpoints.
     */
    preferredEndpoints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The value to set for the sync interval seconds.
     */
    syncIntervalSeconds?: pulumi.Input<number>;
    /**
     * The type of host set
     */
    type?: pulumi.Input<string>;
}
