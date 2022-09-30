// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Katasec.Boundary
{
    /// <summary>
    /// The credential store for Vault resource allows you to configure a Boundary credential store for Vault.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Boundary = Katasec.Boundary;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var org = new Boundary.BoundaryScope("org", new()
    ///     {
    ///         Description = "My first scope!",
    ///         ScopeId = "global",
    ///         AutoCreateAdminRole = true,
    ///         AutoCreateDefaultRole = true,
    ///     });
    /// 
    ///     var project = new Boundary.BoundaryScope("project", new()
    ///     {
    ///         Description = "My first scope!",
    ///         ScopeId = org.Id,
    ///         AutoCreateAdminRole = true,
    ///     });
    /// 
    ///     var example = new Boundary.BoundaryCredentialStoreVault("example", new()
    ///     {
    ///         Description = "My first Vault credential store!",
    ///         Address = "http://127.0.0.1:8200",
    ///         Token = "s.0ufRo6XEGU2jOqnIr7OlFYP5",
    ///         ScopeId = project.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import boundary:index/boundaryCredentialStoreVault:BoundaryCredentialStoreVault foo &lt;my-id&gt;
    /// ```
    /// </summary>
    [BoundaryResourceType("boundary:index/boundaryCredentialStoreVault:BoundaryCredentialStoreVault")]
    public partial class BoundaryCredentialStoreVault : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The address to Vault server. This should be a complete URL such as 'https://127.0.0.1:8200'
        /// </summary>
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        /// <summary>
        /// A PEM-encoded CA certificate to verify the Vault server's TLS certificate.
        /// </summary>
        [Output("caCert")]
        public Output<string?> CaCert { get; private set; } = null!;

        /// <summary>
        /// A PEM-encoded client certificate to use for TLS authentication to the Vault server.
        /// </summary>
        [Output("clientCertificate")]
        public Output<string?> ClientCertificate { get; private set; } = null!;

        /// <summary>
        /// A PEM-encoded private key matching the client certificate from 'client_certificate'.
        /// </summary>
        [Output("clientCertificateKey")]
        public Output<string?> ClientCertificateKey { get; private set; } = null!;

        /// <summary>
        /// The Vault client certificate key hmac.
        /// </summary>
        [Output("clientCertificateKeyHmac")]
        public Output<string> ClientCertificateKeyHmac { get; private set; } = null!;

        /// <summary>
        /// The Vault credential store description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The Vault credential store name. Defaults to the resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The namespace within Vault to use.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// The scope for this credential store.
        /// </summary>
        [Output("scopeId")]
        public Output<string> ScopeId { get; private set; } = null!;

        /// <summary>
        /// Name to use as the SNI host when connecting to Vault via TLS.
        /// </summary>
        [Output("tlsServerName")]
        public Output<string?> TlsServerName { get; private set; } = null!;

        /// <summary>
        /// Whether or not to skip TLS verification.
        /// </summary>
        [Output("tlsSkipVerify")]
        public Output<bool?> TlsSkipVerify { get; private set; } = null!;

        /// <summary>
        /// A token used for accessing Vault.
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        /// <summary>
        /// The Vault token hmac.
        /// </summary>
        [Output("tokenHmac")]
        public Output<string> TokenHmac { get; private set; } = null!;


        /// <summary>
        /// Create a BoundaryCredentialStoreVault resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BoundaryCredentialStoreVault(string name, BoundaryCredentialStoreVaultArgs args, CustomResourceOptions? options = null)
            : base("boundary:index/boundaryCredentialStoreVault:BoundaryCredentialStoreVault", name, args ?? new BoundaryCredentialStoreVaultArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BoundaryCredentialStoreVault(string name, Input<string> id, BoundaryCredentialStoreVaultState? state = null, CustomResourceOptions? options = null)
            : base("boundary:index/boundaryCredentialStoreVault:BoundaryCredentialStoreVault", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/katasec/pulumi-boundary/releases/downloadv${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BoundaryCredentialStoreVault resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BoundaryCredentialStoreVault Get(string name, Input<string> id, BoundaryCredentialStoreVaultState? state = null, CustomResourceOptions? options = null)
        {
            return new BoundaryCredentialStoreVault(name, id, state, options);
        }
    }

    public sealed class BoundaryCredentialStoreVaultArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The address to Vault server. This should be a complete URL such as 'https://127.0.0.1:8200'
        /// </summary>
        [Input("address", required: true)]
        public Input<string> Address { get; set; } = null!;

        /// <summary>
        /// A PEM-encoded CA certificate to verify the Vault server's TLS certificate.
        /// </summary>
        [Input("caCert")]
        public Input<string>? CaCert { get; set; }

        /// <summary>
        /// A PEM-encoded client certificate to use for TLS authentication to the Vault server.
        /// </summary>
        [Input("clientCertificate")]
        public Input<string>? ClientCertificate { get; set; }

        /// <summary>
        /// A PEM-encoded private key matching the client certificate from 'client_certificate'.
        /// </summary>
        [Input("clientCertificateKey")]
        public Input<string>? ClientCertificateKey { get; set; }

        /// <summary>
        /// The Vault credential store description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Vault credential store name. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace within Vault to use.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The scope for this credential store.
        /// </summary>
        [Input("scopeId", required: true)]
        public Input<string> ScopeId { get; set; } = null!;

        /// <summary>
        /// Name to use as the SNI host when connecting to Vault via TLS.
        /// </summary>
        [Input("tlsServerName")]
        public Input<string>? TlsServerName { get; set; }

        /// <summary>
        /// Whether or not to skip TLS verification.
        /// </summary>
        [Input("tlsSkipVerify")]
        public Input<bool>? TlsSkipVerify { get; set; }

        /// <summary>
        /// A token used for accessing Vault.
        /// </summary>
        [Input("token", required: true)]
        public Input<string> Token { get; set; } = null!;

        public BoundaryCredentialStoreVaultArgs()
        {
        }
        public static new BoundaryCredentialStoreVaultArgs Empty => new BoundaryCredentialStoreVaultArgs();
    }

    public sealed class BoundaryCredentialStoreVaultState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The address to Vault server. This should be a complete URL such as 'https://127.0.0.1:8200'
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// A PEM-encoded CA certificate to verify the Vault server's TLS certificate.
        /// </summary>
        [Input("caCert")]
        public Input<string>? CaCert { get; set; }

        /// <summary>
        /// A PEM-encoded client certificate to use for TLS authentication to the Vault server.
        /// </summary>
        [Input("clientCertificate")]
        public Input<string>? ClientCertificate { get; set; }

        /// <summary>
        /// A PEM-encoded private key matching the client certificate from 'client_certificate'.
        /// </summary>
        [Input("clientCertificateKey")]
        public Input<string>? ClientCertificateKey { get; set; }

        /// <summary>
        /// The Vault client certificate key hmac.
        /// </summary>
        [Input("clientCertificateKeyHmac")]
        public Input<string>? ClientCertificateKeyHmac { get; set; }

        /// <summary>
        /// The Vault credential store description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Vault credential store name. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace within Vault to use.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The scope for this credential store.
        /// </summary>
        [Input("scopeId")]
        public Input<string>? ScopeId { get; set; }

        /// <summary>
        /// Name to use as the SNI host when connecting to Vault via TLS.
        /// </summary>
        [Input("tlsServerName")]
        public Input<string>? TlsServerName { get; set; }

        /// <summary>
        /// Whether or not to skip TLS verification.
        /// </summary>
        [Input("tlsSkipVerify")]
        public Input<bool>? TlsSkipVerify { get; set; }

        /// <summary>
        /// A token used for accessing Vault.
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        /// <summary>
        /// The Vault token hmac.
        /// </summary>
        [Input("tokenHmac")]
        public Input<string>? TokenHmac { get; set; }

        public BoundaryCredentialStoreVaultState()
        {
        }
        public static new BoundaryCredentialStoreVaultState Empty => new BoundaryCredentialStoreVaultState();
    }
}
