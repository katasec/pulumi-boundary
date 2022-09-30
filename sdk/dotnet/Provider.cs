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
    /// The provider type for the boundary package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// </summary>
    [BoundaryResourceType("pulumi:providers:boundary")]
    public partial class Provider : global::Pulumi.ProviderResource
    {
        /// <summary>
        /// The base url of the Boundary API, e.g. "http://127.0.0.1:9200". If not set, it will be read from the "BOUNDARY_ADDR" env
        /// var.
        /// </summary>
        [Output("addr")]
        public Output<string> Addr { get; private set; } = null!;

        /// <summary>
        /// The auth method ID e.g. ampw_1234567890
        /// </summary>
        [Output("authMethodId")]
        public Output<string?> AuthMethodId { get; private set; } = null!;

        /// <summary>
        /// The auth method login name for password-style auth methods
        /// </summary>
        [Output("passwordAuthMethodLoginName")]
        public Output<string?> PasswordAuthMethodLoginName { get; private set; } = null!;

        /// <summary>
        /// The auth method password for password-style auth methods
        /// </summary>
        [Output("passwordAuthMethodPassword")]
        public Output<string?> PasswordAuthMethodPassword { get; private set; } = null!;

        /// <summary>
        /// Can be a heredoc string or a path on disk. If set, the string/file will be parsed as HCL and used with the recovery KMS
        /// mechanism. While this is set, it will override any other authentication information; the KMS mechanism will always be
        /// used. See Boundary's KMS docs for examples: https://boundaryproject.io/docs/configuration/kms
        /// </summary>
        [Output("recoveryKmsHcl")]
        public Output<string?> RecoveryKmsHcl { get; private set; } = null!;

        /// <summary>
        /// The Boundary token to use, as a string or path on disk containing just the string. If set, the token read here will be
        /// used in place of authenticating with the auth method specified in "auth_method_id", although the recovery KMS mechanism
        /// will still override this. Can also be set with the BOUNDARY_TOKEN environment variable.
        /// </summary>
        [Output("token")]
        public Output<string?> Token { get; private set; } = null!;


        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs args, CustomResourceOptions? options = null)
            : base("boundary", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
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
    }

    public sealed class ProviderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The base url of the Boundary API, e.g. "http://127.0.0.1:9200". If not set, it will be read from the "BOUNDARY_ADDR" env
        /// var.
        /// </summary>
        [Input("addr", required: true)]
        public Input<string> Addr { get; set; } = null!;

        /// <summary>
        /// The auth method ID e.g. ampw_1234567890
        /// </summary>
        [Input("authMethodId")]
        public Input<string>? AuthMethodId { get; set; }

        /// <summary>
        /// The auth method login name for password-style auth methods
        /// </summary>
        [Input("passwordAuthMethodLoginName")]
        public Input<string>? PasswordAuthMethodLoginName { get; set; }

        /// <summary>
        /// The auth method password for password-style auth methods
        /// </summary>
        [Input("passwordAuthMethodPassword")]
        public Input<string>? PasswordAuthMethodPassword { get; set; }

        /// <summary>
        /// Can be a heredoc string or a path on disk. If set, the string/file will be parsed as HCL and used with the recovery KMS
        /// mechanism. While this is set, it will override any other authentication information; the KMS mechanism will always be
        /// used. See Boundary's KMS docs for examples: https://boundaryproject.io/docs/configuration/kms
        /// </summary>
        [Input("recoveryKmsHcl")]
        public Input<string>? RecoveryKmsHcl { get; set; }

        /// <summary>
        /// When set to true, does not validate the Boundary API endpoint certificate
        /// </summary>
        [Input("tlsInsecure", json: true)]
        public Input<bool>? TlsInsecure { get; set; }

        /// <summary>
        /// The Boundary token to use, as a string or path on disk containing just the string. If set, the token read here will be
        /// used in place of authenticating with the auth method specified in "auth_method_id", although the recovery KMS mechanism
        /// will still override this. Can also be set with the BOUNDARY_TOKEN environment variable.
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        public ProviderArgs()
        {
        }
        public static new ProviderArgs Empty => new ProviderArgs();
    }
}
