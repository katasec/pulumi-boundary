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
    /// The SSH private key credential resource allows you to configure a credential using a username, private key and optional passphrase.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.IO;
    /// using Pulumi;
    /// using Boundary = Katasec.Boundary;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var org = new Boundary.BoundaryScope("org", new()
    ///     {
    ///         Description = "global scope",
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
    ///     var exampleBoundaryCredentialStoreStatic = new Boundary.BoundaryCredentialStoreStatic("exampleBoundaryCredentialStoreStatic", new()
    ///     {
    ///         Description = "My first static credential store!",
    ///         ScopeId = project.Id,
    ///     });
    /// 
    ///     var exampleBoundaryCredentialSshPrivateKey = new Boundary.BoundaryCredentialSshPrivateKey("exampleBoundaryCredentialSshPrivateKey", new()
    ///     {
    ///         Description = "My first ssh private key credential!",
    ///         CredentialStoreId = exampleBoundaryCredentialStoreStatic.Id,
    ///         Username = "my-username",
    ///         PrivateKey = File.ReadAllText("~/.ssh/id_rsa"),
    ///         PrivateKeyPassphrase = "optional-passphrase",
    ///     });
    /// 
    ///     // change to the passphrase of the Private Key if required
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import boundary:index/boundaryCredentialSshPrivateKey:BoundaryCredentialSshPrivateKey example_ssh_private_key &lt;my-id&gt;
    /// ```
    /// </summary>
    [BoundaryResourceType("boundary:index/boundaryCredentialSshPrivateKey:BoundaryCredentialSshPrivateKey")]
    public partial class BoundaryCredentialSshPrivateKey : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the credential store this credential belongs to.
        /// </summary>
        [Output("credentialStoreId")]
        public Output<string> CredentialStoreId { get; private set; } = null!;

        /// <summary>
        /// The description of the credential.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the credential. Defaults to the resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The private key associated with the credential.
        /// </summary>
        [Output("privateKey")]
        public Output<string> PrivateKey { get; private set; } = null!;

        /// <summary>
        /// The private key hmac.
        /// </summary>
        [Output("privateKeyHmac")]
        public Output<string> PrivateKeyHmac { get; private set; } = null!;

        /// <summary>
        /// The passphrase of the private key associated with the credential.
        /// </summary>
        [Output("privateKeyPassphrase")]
        public Output<string?> PrivateKeyPassphrase { get; private set; } = null!;

        /// <summary>
        /// The private key passphrase hmac.
        /// </summary>
        [Output("privateKeyPassphraseHmac")]
        public Output<string> PrivateKeyPassphraseHmac { get; private set; } = null!;

        /// <summary>
        /// The username associated with the credential.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a BoundaryCredentialSshPrivateKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BoundaryCredentialSshPrivateKey(string name, BoundaryCredentialSshPrivateKeyArgs args, CustomResourceOptions? options = null)
            : base("boundary:index/boundaryCredentialSshPrivateKey:BoundaryCredentialSshPrivateKey", name, args ?? new BoundaryCredentialSshPrivateKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BoundaryCredentialSshPrivateKey(string name, Input<string> id, BoundaryCredentialSshPrivateKeyState? state = null, CustomResourceOptions? options = null)
            : base("boundary:index/boundaryCredentialSshPrivateKey:BoundaryCredentialSshPrivateKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BoundaryCredentialSshPrivateKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BoundaryCredentialSshPrivateKey Get(string name, Input<string> id, BoundaryCredentialSshPrivateKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new BoundaryCredentialSshPrivateKey(name, id, state, options);
        }
    }

    public sealed class BoundaryCredentialSshPrivateKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the credential store this credential belongs to.
        /// </summary>
        [Input("credentialStoreId", required: true)]
        public Input<string> CredentialStoreId { get; set; } = null!;

        /// <summary>
        /// The description of the credential.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the credential. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The private key associated with the credential.
        /// </summary>
        [Input("privateKey", required: true)]
        public Input<string> PrivateKey { get; set; } = null!;

        /// <summary>
        /// The passphrase of the private key associated with the credential.
        /// </summary>
        [Input("privateKeyPassphrase")]
        public Input<string>? PrivateKeyPassphrase { get; set; }

        /// <summary>
        /// The username associated with the credential.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public BoundaryCredentialSshPrivateKeyArgs()
        {
        }
        public static new BoundaryCredentialSshPrivateKeyArgs Empty => new BoundaryCredentialSshPrivateKeyArgs();
    }

    public sealed class BoundaryCredentialSshPrivateKeyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the credential store this credential belongs to.
        /// </summary>
        [Input("credentialStoreId")]
        public Input<string>? CredentialStoreId { get; set; }

        /// <summary>
        /// The description of the credential.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the credential. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The private key associated with the credential.
        /// </summary>
        [Input("privateKey")]
        public Input<string>? PrivateKey { get; set; }

        /// <summary>
        /// The private key hmac.
        /// </summary>
        [Input("privateKeyHmac")]
        public Input<string>? PrivateKeyHmac { get; set; }

        /// <summary>
        /// The passphrase of the private key associated with the credential.
        /// </summary>
        [Input("privateKeyPassphrase")]
        public Input<string>? PrivateKeyPassphrase { get; set; }

        /// <summary>
        /// The private key passphrase hmac.
        /// </summary>
        [Input("privateKeyPassphraseHmac")]
        public Input<string>? PrivateKeyPassphraseHmac { get; set; }

        /// <summary>
        /// The username associated with the credential.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public BoundaryCredentialSshPrivateKeyState()
        {
        }
        public static new BoundaryCredentialSshPrivateKeyState Empty => new BoundaryCredentialSshPrivateKeyState();
    }
}
