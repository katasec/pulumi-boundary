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
    /// The host_set_plugin resource allows you to configure a Boundary host set. Host sets are always part of a host catalog, so a host catalog resource should be used inline or you should have the host catalog ID in hand to successfully configure a host set.
    /// </summary>
    [BoundaryResourceType("boundary:index/boundaryHostSetPlugin:BoundaryHostSetPlugin")]
    public partial class BoundaryHostSetPlugin : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The attributes for the host set. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host set.
        /// </summary>
        [Output("attributesJson")]
        public Output<string?> AttributesJson { get; private set; } = null!;

        /// <summary>
        /// The host set description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The catalog for the host set.
        /// </summary>
        [Output("hostCatalogId")]
        public Output<string> HostCatalogId { get; private set; } = null!;

        /// <summary>
        /// The host set name. Defaults to the resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ordered list of preferred endpoints.
        /// </summary>
        [Output("preferredEndpoints")]
        public Output<ImmutableArray<string>> PreferredEndpoints { get; private set; } = null!;

        /// <summary>
        /// The value to set for the sync interval seconds.
        /// </summary>
        [Output("syncIntervalSeconds")]
        public Output<int?> SyncIntervalSeconds { get; private set; } = null!;

        /// <summary>
        /// The type of host set
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a BoundaryHostSetPlugin resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BoundaryHostSetPlugin(string name, BoundaryHostSetPluginArgs args, CustomResourceOptions? options = null)
            : base("boundary:index/boundaryHostSetPlugin:BoundaryHostSetPlugin", name, args ?? new BoundaryHostSetPluginArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BoundaryHostSetPlugin(string name, Input<string> id, BoundaryHostSetPluginState? state = null, CustomResourceOptions? options = null)
            : base("boundary:index/boundaryHostSetPlugin:BoundaryHostSetPlugin", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BoundaryHostSetPlugin resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BoundaryHostSetPlugin Get(string name, Input<string> id, BoundaryHostSetPluginState? state = null, CustomResourceOptions? options = null)
        {
            return new BoundaryHostSetPlugin(name, id, state, options);
        }
    }

    public sealed class BoundaryHostSetPluginArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The attributes for the host set. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host set.
        /// </summary>
        [Input("attributesJson")]
        public Input<string>? AttributesJson { get; set; }

        /// <summary>
        /// The host set description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The catalog for the host set.
        /// </summary>
        [Input("hostCatalogId", required: true)]
        public Input<string> HostCatalogId { get; set; } = null!;

        /// <summary>
        /// The host set name. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("preferredEndpoints")]
        private InputList<string>? _preferredEndpoints;

        /// <summary>
        /// The ordered list of preferred endpoints.
        /// </summary>
        public InputList<string> PreferredEndpoints
        {
            get => _preferredEndpoints ?? (_preferredEndpoints = new InputList<string>());
            set => _preferredEndpoints = value;
        }

        /// <summary>
        /// The value to set for the sync interval seconds.
        /// </summary>
        [Input("syncIntervalSeconds")]
        public Input<int>? SyncIntervalSeconds { get; set; }

        /// <summary>
        /// The type of host set
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public BoundaryHostSetPluginArgs()
        {
        }
        public static new BoundaryHostSetPluginArgs Empty => new BoundaryHostSetPluginArgs();
    }

    public sealed class BoundaryHostSetPluginState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The attributes for the host set. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host set.
        /// </summary>
        [Input("attributesJson")]
        public Input<string>? AttributesJson { get; set; }

        /// <summary>
        /// The host set description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The catalog for the host set.
        /// </summary>
        [Input("hostCatalogId")]
        public Input<string>? HostCatalogId { get; set; }

        /// <summary>
        /// The host set name. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("preferredEndpoints")]
        private InputList<string>? _preferredEndpoints;

        /// <summary>
        /// The ordered list of preferred endpoints.
        /// </summary>
        public InputList<string> PreferredEndpoints
        {
            get => _preferredEndpoints ?? (_preferredEndpoints = new InputList<string>());
            set => _preferredEndpoints = value;
        }

        /// <summary>
        /// The value to set for the sync interval seconds.
        /// </summary>
        [Input("syncIntervalSeconds")]
        public Input<int>? SyncIntervalSeconds { get; set; }

        /// <summary>
        /// The type of host set
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public BoundaryHostSetPluginState()
        {
        }
        public static new BoundaryHostSetPluginState Empty => new BoundaryHostSetPluginState();
    }
}