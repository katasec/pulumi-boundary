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
    /// The account resource allows you to configure a Boundary account.
    /// </summary>
    [BoundaryResourceType("boundary:index/boundaryAccountOidc:BoundaryAccountOidc")]
    public partial class BoundaryAccountOidc : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The resource ID for the auth method.
        /// </summary>
        [Output("authMethodId")]
        public Output<string> AuthMethodId { get; private set; } = null!;

        /// <summary>
        /// The account description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The OIDC issuer.
        /// </summary>
        [Output("issuer")]
        public Output<string?> Issuer { get; private set; } = null!;

        /// <summary>
        /// The account name. Defaults to the resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The OIDC subject.
        /// </summary>
        [Output("subject")]
        public Output<string?> Subject { get; private set; } = null!;


        /// <summary>
        /// Create a BoundaryAccountOidc resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BoundaryAccountOidc(string name, BoundaryAccountOidcArgs args, CustomResourceOptions? options = null)
            : base("boundary:index/boundaryAccountOidc:BoundaryAccountOidc", name, args ?? new BoundaryAccountOidcArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BoundaryAccountOidc(string name, Input<string> id, BoundaryAccountOidcState? state = null, CustomResourceOptions? options = null)
            : base("boundary:index/boundaryAccountOidc:BoundaryAccountOidc", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/katasec/pulumi-boundary/releases/download/v${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BoundaryAccountOidc resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BoundaryAccountOidc Get(string name, Input<string> id, BoundaryAccountOidcState? state = null, CustomResourceOptions? options = null)
        {
            return new BoundaryAccountOidc(name, id, state, options);
        }
    }

    public sealed class BoundaryAccountOidcArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource ID for the auth method.
        /// </summary>
        [Input("authMethodId", required: true)]
        public Input<string> AuthMethodId { get; set; } = null!;

        /// <summary>
        /// The account description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The OIDC issuer.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        /// <summary>
        /// The account name. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The OIDC subject.
        /// </summary>
        [Input("subject")]
        public Input<string>? Subject { get; set; }

        public BoundaryAccountOidcArgs()
        {
        }
        public static new BoundaryAccountOidcArgs Empty => new BoundaryAccountOidcArgs();
    }

    public sealed class BoundaryAccountOidcState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource ID for the auth method.
        /// </summary>
        [Input("authMethodId")]
        public Input<string>? AuthMethodId { get; set; }

        /// <summary>
        /// The account description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The OIDC issuer.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        /// <summary>
        /// The account name. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The OIDC subject.
        /// </summary>
        [Input("subject")]
        public Input<string>? Subject { get; set; }

        public BoundaryAccountOidcState()
        {
        }
        public static new BoundaryAccountOidcState Empty => new BoundaryAccountOidcState();
    }
}