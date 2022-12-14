// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The OIDC auth method resource allows you to configure a Boundary auth_method_oidc.
 */
export class BoundaryAuthMethodOidc extends pulumi.CustomResource {
    /**
     * Get an existing BoundaryAuthMethodOidc resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BoundaryAuthMethodOidcState, opts?: pulumi.CustomResourceOptions): BoundaryAuthMethodOidc {
        return new BoundaryAuthMethodOidc(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'boundary:index/boundaryAuthMethodOidc:BoundaryAuthMethodOidc';

    /**
     * Returns true if the given object is an instance of BoundaryAuthMethodOidc.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BoundaryAuthMethodOidc {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BoundaryAuthMethodOidc.__pulumiType;
    }

    /**
     * Account claim maps for the toClaim of sub.
     */
    public readonly accountClaimMaps!: pulumi.Output<string[] | undefined>;
    /**
     * Audiences for which the provider responses will be allowed
     */
    public readonly allowedAudiences!: pulumi.Output<string[] | undefined>;
    /**
     * The API prefix to use when generating callback URLs for the provider. Should be set to an address at which the provider can reach back to the controller.
     */
    public readonly apiUrlPrefix!: pulumi.Output<string | undefined>;
    /**
     * The URL that should be provided to the IdP for callbacks.
     */
    public readonly callbackUrl!: pulumi.Output<string>;
    /**
     * Claims scopes.
     */
    public readonly claimsScopes!: pulumi.Output<string[] | undefined>;
    /**
     * The client ID assigned to this auth method from the provider.
     */
    public readonly clientId!: pulumi.Output<string | undefined>;
    /**
     * The secret key assigned to this auth method from the provider. Once set, only the hash will be kept and the original value can be removed from configuration.
     */
    public readonly clientSecret!: pulumi.Output<string | undefined>;
    /**
     * The HMAC of the client secret returned by the Boundary controller, which is used for comparison after initial setting of the value.
     */
    public readonly clientSecretHmac!: pulumi.Output<string>;
    /**
     * The auth method description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Disables validation logic ensuring that the OIDC provider's information from its discovery endpoint matches the information here. The validation is only performed at create or update time.
     */
    public readonly disableDiscoveredConfigValidation!: pulumi.Output<boolean | undefined>;
    /**
     * A list of CA certificates to trust when validating the IdP's token signatures.
     */
    public readonly idpCaCerts!: pulumi.Output<string[] | undefined>;
    /**
     * When true, makes this auth method the primary auth method for the scope in which it resides. The primary auth method for a scope means the the user will be automatically created when they login using an OIDC account.
     */
    public readonly isPrimaryForScope!: pulumi.Output<boolean | undefined>;
    /**
     * The issuer corresponding to the provider, which must match the issuer field in generated tokens.
     */
    public readonly issuer!: pulumi.Output<string | undefined>;
    /**
     * The max age to provide to the provider, indicating how much time is allowed to have passed since the last authentication before the user is challenged again.
     */
    public readonly maxAge!: pulumi.Output<number | undefined>;
    /**
     * The auth method name. Defaults to the resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The scope ID.
     */
    public readonly scopeId!: pulumi.Output<string>;
    /**
     * Allowed signing algorithms for the provider's issued tokens.
     */
    public readonly signingAlgorithms!: pulumi.Output<string[] | undefined>;
    /**
     * Can be one of 'inactive', 'active-private', or 'active-public'. Currently automatically set to active-public.
     */
    public readonly state!: pulumi.Output<string>;
    /**
     * The type of auth method; hardcoded.
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a BoundaryAuthMethodOidc resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BoundaryAuthMethodOidcArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BoundaryAuthMethodOidcArgs | BoundaryAuthMethodOidcState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BoundaryAuthMethodOidcState | undefined;
            resourceInputs["accountClaimMaps"] = state ? state.accountClaimMaps : undefined;
            resourceInputs["allowedAudiences"] = state ? state.allowedAudiences : undefined;
            resourceInputs["apiUrlPrefix"] = state ? state.apiUrlPrefix : undefined;
            resourceInputs["callbackUrl"] = state ? state.callbackUrl : undefined;
            resourceInputs["claimsScopes"] = state ? state.claimsScopes : undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clientSecret"] = state ? state.clientSecret : undefined;
            resourceInputs["clientSecretHmac"] = state ? state.clientSecretHmac : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["disableDiscoveredConfigValidation"] = state ? state.disableDiscoveredConfigValidation : undefined;
            resourceInputs["idpCaCerts"] = state ? state.idpCaCerts : undefined;
            resourceInputs["isPrimaryForScope"] = state ? state.isPrimaryForScope : undefined;
            resourceInputs["issuer"] = state ? state.issuer : undefined;
            resourceInputs["maxAge"] = state ? state.maxAge : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["scopeId"] = state ? state.scopeId : undefined;
            resourceInputs["signingAlgorithms"] = state ? state.signingAlgorithms : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as BoundaryAuthMethodOidcArgs | undefined;
            if ((!args || args.scopeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scopeId'");
            }
            resourceInputs["accountClaimMaps"] = args ? args.accountClaimMaps : undefined;
            resourceInputs["allowedAudiences"] = args ? args.allowedAudiences : undefined;
            resourceInputs["apiUrlPrefix"] = args ? args.apiUrlPrefix : undefined;
            resourceInputs["callbackUrl"] = args ? args.callbackUrl : undefined;
            resourceInputs["claimsScopes"] = args ? args.claimsScopes : undefined;
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clientSecret"] = args ? args.clientSecret : undefined;
            resourceInputs["clientSecretHmac"] = args ? args.clientSecretHmac : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["disableDiscoveredConfigValidation"] = args ? args.disableDiscoveredConfigValidation : undefined;
            resourceInputs["idpCaCerts"] = args ? args.idpCaCerts : undefined;
            resourceInputs["isPrimaryForScope"] = args ? args.isPrimaryForScope : undefined;
            resourceInputs["issuer"] = args ? args.issuer : undefined;
            resourceInputs["maxAge"] = args ? args.maxAge : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["scopeId"] = args ? args.scopeId : undefined;
            resourceInputs["signingAlgorithms"] = args ? args.signingAlgorithms : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BoundaryAuthMethodOidc.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BoundaryAuthMethodOidc resources.
 */
export interface BoundaryAuthMethodOidcState {
    /**
     * Account claim maps for the toClaim of sub.
     */
    accountClaimMaps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Audiences for which the provider responses will be allowed
     */
    allowedAudiences?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The API prefix to use when generating callback URLs for the provider. Should be set to an address at which the provider can reach back to the controller.
     */
    apiUrlPrefix?: pulumi.Input<string>;
    /**
     * The URL that should be provided to the IdP for callbacks.
     */
    callbackUrl?: pulumi.Input<string>;
    /**
     * Claims scopes.
     */
    claimsScopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The client ID assigned to this auth method from the provider.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The secret key assigned to this auth method from the provider. Once set, only the hash will be kept and the original value can be removed from configuration.
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * The HMAC of the client secret returned by the Boundary controller, which is used for comparison after initial setting of the value.
     */
    clientSecretHmac?: pulumi.Input<string>;
    /**
     * The auth method description.
     */
    description?: pulumi.Input<string>;
    /**
     * Disables validation logic ensuring that the OIDC provider's information from its discovery endpoint matches the information here. The validation is only performed at create or update time.
     */
    disableDiscoveredConfigValidation?: pulumi.Input<boolean>;
    /**
     * A list of CA certificates to trust when validating the IdP's token signatures.
     */
    idpCaCerts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When true, makes this auth method the primary auth method for the scope in which it resides. The primary auth method for a scope means the the user will be automatically created when they login using an OIDC account.
     */
    isPrimaryForScope?: pulumi.Input<boolean>;
    /**
     * The issuer corresponding to the provider, which must match the issuer field in generated tokens.
     */
    issuer?: pulumi.Input<string>;
    /**
     * The max age to provide to the provider, indicating how much time is allowed to have passed since the last authentication before the user is challenged again.
     */
    maxAge?: pulumi.Input<number>;
    /**
     * The auth method name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The scope ID.
     */
    scopeId?: pulumi.Input<string>;
    /**
     * Allowed signing algorithms for the provider's issued tokens.
     */
    signingAlgorithms?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Can be one of 'inactive', 'active-private', or 'active-public'. Currently automatically set to active-public.
     */
    state?: pulumi.Input<string>;
    /**
     * The type of auth method; hardcoded.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BoundaryAuthMethodOidc resource.
 */
export interface BoundaryAuthMethodOidcArgs {
    /**
     * Account claim maps for the toClaim of sub.
     */
    accountClaimMaps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Audiences for which the provider responses will be allowed
     */
    allowedAudiences?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The API prefix to use when generating callback URLs for the provider. Should be set to an address at which the provider can reach back to the controller.
     */
    apiUrlPrefix?: pulumi.Input<string>;
    /**
     * The URL that should be provided to the IdP for callbacks.
     */
    callbackUrl?: pulumi.Input<string>;
    /**
     * Claims scopes.
     */
    claimsScopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The client ID assigned to this auth method from the provider.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The secret key assigned to this auth method from the provider. Once set, only the hash will be kept and the original value can be removed from configuration.
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * The HMAC of the client secret returned by the Boundary controller, which is used for comparison after initial setting of the value.
     */
    clientSecretHmac?: pulumi.Input<string>;
    /**
     * The auth method description.
     */
    description?: pulumi.Input<string>;
    /**
     * Disables validation logic ensuring that the OIDC provider's information from its discovery endpoint matches the information here. The validation is only performed at create or update time.
     */
    disableDiscoveredConfigValidation?: pulumi.Input<boolean>;
    /**
     * A list of CA certificates to trust when validating the IdP's token signatures.
     */
    idpCaCerts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When true, makes this auth method the primary auth method for the scope in which it resides. The primary auth method for a scope means the the user will be automatically created when they login using an OIDC account.
     */
    isPrimaryForScope?: pulumi.Input<boolean>;
    /**
     * The issuer corresponding to the provider, which must match the issuer field in generated tokens.
     */
    issuer?: pulumi.Input<string>;
    /**
     * The max age to provide to the provider, indicating how much time is allowed to have passed since the last authentication before the user is challenged again.
     */
    maxAge?: pulumi.Input<number>;
    /**
     * The auth method name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The scope ID.
     */
    scopeId: pulumi.Input<string>;
    /**
     * Allowed signing algorithms for the provider's issued tokens.
     */
    signingAlgorithms?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Can be one of 'inactive', 'active-private', or 'active-public'. Currently automatically set to active-public.
     */
    state?: pulumi.Input<string>;
    /**
     * The type of auth method; hardcoded.
     */
    type?: pulumi.Input<string>;
}
