// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The credential store for Vault resource allows you to configure a Boundary credential store for Vault.
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
 * const example = new boundary.BoundaryCredentialStoreVault("example", {
 *     description: "My first Vault credential store!",
 *     address: "http://127.0.0.1:8200",
 *     token: "s.0ufRo6XEGU2jOqnIr7OlFYP5",
 *     scopeId: project.id,
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import boundary:index/boundaryCredentialStoreVault:BoundaryCredentialStoreVault foo <my-id>
 * ```
 */
export class BoundaryCredentialStoreVault extends pulumi.CustomResource {
    /**
     * Get an existing BoundaryCredentialStoreVault resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BoundaryCredentialStoreVaultState, opts?: pulumi.CustomResourceOptions): BoundaryCredentialStoreVault {
        return new BoundaryCredentialStoreVault(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'boundary:index/boundaryCredentialStoreVault:BoundaryCredentialStoreVault';

    /**
     * Returns true if the given object is an instance of BoundaryCredentialStoreVault.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BoundaryCredentialStoreVault {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BoundaryCredentialStoreVault.__pulumiType;
    }

    /**
     * The address to Vault server. This should be a complete URL such as 'https://127.0.0.1:8200'
     */
    public readonly address!: pulumi.Output<string>;
    /**
     * A PEM-encoded CA certificate to verify the Vault server's TLS certificate.
     */
    public readonly caCert!: pulumi.Output<string | undefined>;
    /**
     * A PEM-encoded client certificate to use for TLS authentication to the Vault server.
     */
    public readonly clientCertificate!: pulumi.Output<string | undefined>;
    /**
     * A PEM-encoded private key matching the client certificate from 'client_certificate'.
     */
    public readonly clientCertificateKey!: pulumi.Output<string | undefined>;
    /**
     * The Vault client certificate key hmac.
     */
    public /*out*/ readonly clientCertificateKeyHmac!: pulumi.Output<string>;
    /**
     * The Vault credential store description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Vault credential store name. Defaults to the resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The namespace within Vault to use.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * The scope for this credential store.
     */
    public readonly scopeId!: pulumi.Output<string>;
    /**
     * Name to use as the SNI host when connecting to Vault via TLS.
     */
    public readonly tlsServerName!: pulumi.Output<string | undefined>;
    /**
     * Whether or not to skip TLS verification.
     */
    public readonly tlsSkipVerify!: pulumi.Output<boolean | undefined>;
    /**
     * A token used for accessing Vault.
     */
    public readonly token!: pulumi.Output<string>;
    /**
     * The Vault token hmac.
     */
    public /*out*/ readonly tokenHmac!: pulumi.Output<string>;

    /**
     * Create a BoundaryCredentialStoreVault resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BoundaryCredentialStoreVaultArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BoundaryCredentialStoreVaultArgs | BoundaryCredentialStoreVaultState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BoundaryCredentialStoreVaultState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["caCert"] = state ? state.caCert : undefined;
            resourceInputs["clientCertificate"] = state ? state.clientCertificate : undefined;
            resourceInputs["clientCertificateKey"] = state ? state.clientCertificateKey : undefined;
            resourceInputs["clientCertificateKeyHmac"] = state ? state.clientCertificateKeyHmac : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["scopeId"] = state ? state.scopeId : undefined;
            resourceInputs["tlsServerName"] = state ? state.tlsServerName : undefined;
            resourceInputs["tlsSkipVerify"] = state ? state.tlsSkipVerify : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["tokenHmac"] = state ? state.tokenHmac : undefined;
        } else {
            const args = argsOrState as BoundaryCredentialStoreVaultArgs | undefined;
            if ((!args || args.address === undefined) && !opts.urn) {
                throw new Error("Missing required property 'address'");
            }
            if ((!args || args.scopeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scopeId'");
            }
            if ((!args || args.token === undefined) && !opts.urn) {
                throw new Error("Missing required property 'token'");
            }
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["caCert"] = args ? args.caCert : undefined;
            resourceInputs["clientCertificate"] = args ? args.clientCertificate : undefined;
            resourceInputs["clientCertificateKey"] = args ? args.clientCertificateKey : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["scopeId"] = args ? args.scopeId : undefined;
            resourceInputs["tlsServerName"] = args ? args.tlsServerName : undefined;
            resourceInputs["tlsSkipVerify"] = args ? args.tlsSkipVerify : undefined;
            resourceInputs["token"] = args ? args.token : undefined;
            resourceInputs["clientCertificateKeyHmac"] = undefined /*out*/;
            resourceInputs["tokenHmac"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BoundaryCredentialStoreVault.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BoundaryCredentialStoreVault resources.
 */
export interface BoundaryCredentialStoreVaultState {
    /**
     * The address to Vault server. This should be a complete URL such as 'https://127.0.0.1:8200'
     */
    address?: pulumi.Input<string>;
    /**
     * A PEM-encoded CA certificate to verify the Vault server's TLS certificate.
     */
    caCert?: pulumi.Input<string>;
    /**
     * A PEM-encoded client certificate to use for TLS authentication to the Vault server.
     */
    clientCertificate?: pulumi.Input<string>;
    /**
     * A PEM-encoded private key matching the client certificate from 'client_certificate'.
     */
    clientCertificateKey?: pulumi.Input<string>;
    /**
     * The Vault client certificate key hmac.
     */
    clientCertificateKeyHmac?: pulumi.Input<string>;
    /**
     * The Vault credential store description.
     */
    description?: pulumi.Input<string>;
    /**
     * The Vault credential store name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace within Vault to use.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The scope for this credential store.
     */
    scopeId?: pulumi.Input<string>;
    /**
     * Name to use as the SNI host when connecting to Vault via TLS.
     */
    tlsServerName?: pulumi.Input<string>;
    /**
     * Whether or not to skip TLS verification.
     */
    tlsSkipVerify?: pulumi.Input<boolean>;
    /**
     * A token used for accessing Vault.
     */
    token?: pulumi.Input<string>;
    /**
     * The Vault token hmac.
     */
    tokenHmac?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BoundaryCredentialStoreVault resource.
 */
export interface BoundaryCredentialStoreVaultArgs {
    /**
     * The address to Vault server. This should be a complete URL such as 'https://127.0.0.1:8200'
     */
    address: pulumi.Input<string>;
    /**
     * A PEM-encoded CA certificate to verify the Vault server's TLS certificate.
     */
    caCert?: pulumi.Input<string>;
    /**
     * A PEM-encoded client certificate to use for TLS authentication to the Vault server.
     */
    clientCertificate?: pulumi.Input<string>;
    /**
     * A PEM-encoded private key matching the client certificate from 'client_certificate'.
     */
    clientCertificateKey?: pulumi.Input<string>;
    /**
     * The Vault credential store description.
     */
    description?: pulumi.Input<string>;
    /**
     * The Vault credential store name. Defaults to the resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace within Vault to use.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The scope for this credential store.
     */
    scopeId: pulumi.Input<string>;
    /**
     * Name to use as the SNI host when connecting to Vault via TLS.
     */
    tlsServerName?: pulumi.Input<string>;
    /**
     * Whether or not to skip TLS verification.
     */
    tlsSkipVerify?: pulumi.Input<boolean>;
    /**
     * A token used for accessing Vault.
     */
    token: pulumi.Input<string>;
}