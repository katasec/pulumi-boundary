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
    /// The target resource allows you to configure a Boundary target.
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import boundary:index/boundaryTarget:BoundaryTarget foo &lt;my-id&gt;
    /// ```
    /// </summary>
    [BoundaryResourceType("boundary:index/boundaryTarget:BoundaryTarget")]
    public partial class BoundaryTarget : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of brokered credential source ID's.
        /// </summary>
        [Output("brokeredCredentialSourceIds")]
        public Output<ImmutableArray<string>> BrokeredCredentialSourceIds { get; private set; } = null!;

        /// <summary>
        /// The default port for this target.
        /// </summary>
        [Output("defaultPort")]
        public Output<int?> DefaultPort { get; private set; } = null!;

        /// <summary>
        /// The target description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A list of host source ID's.
        /// </summary>
        [Output("hostSourceIds")]
        public Output<ImmutableArray<string>> HostSourceIds { get; private set; } = null!;

        /// <summary>
        /// A list of injected application credential source ID's.
        /// </summary>
        [Output("injectedApplicationCredentialSourceIds")]
        public Output<ImmutableArray<string>> InjectedApplicationCredentialSourceIds { get; private set; } = null!;

        /// <summary>
        /// The target name. Defaults to the resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
        /// </summary>
        [Output("scopeId")]
        public Output<string> ScopeId { get; private set; } = null!;

        [Output("sessionConnectionLimit")]
        public Output<int> SessionConnectionLimit { get; private set; } = null!;

        [Output("sessionMaxSeconds")]
        public Output<int> SessionMaxSeconds { get; private set; } = null!;

        /// <summary>
        /// The target resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Boolean expression to filter the workers for this target
        /// </summary>
        [Output("workerFilter")]
        public Output<string?> WorkerFilter { get; private set; } = null!;


        /// <summary>
        /// Create a BoundaryTarget resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BoundaryTarget(string name, BoundaryTargetArgs args, CustomResourceOptions? options = null)
            : base("boundary:index/boundaryTarget:BoundaryTarget", name, args ?? new BoundaryTargetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BoundaryTarget(string name, Input<string> id, BoundaryTargetState? state = null, CustomResourceOptions? options = null)
            : base("boundary:index/boundaryTarget:BoundaryTarget", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BoundaryTarget resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BoundaryTarget Get(string name, Input<string> id, BoundaryTargetState? state = null, CustomResourceOptions? options = null)
        {
            return new BoundaryTarget(name, id, state, options);
        }
    }

    public sealed class BoundaryTargetArgs : global::Pulumi.ResourceArgs
    {
        [Input("brokeredCredentialSourceIds")]
        private InputList<string>? _brokeredCredentialSourceIds;

        /// <summary>
        /// A list of brokered credential source ID's.
        /// </summary>
        public InputList<string> BrokeredCredentialSourceIds
        {
            get => _brokeredCredentialSourceIds ?? (_brokeredCredentialSourceIds = new InputList<string>());
            set => _brokeredCredentialSourceIds = value;
        }

        /// <summary>
        /// The default port for this target.
        /// </summary>
        [Input("defaultPort")]
        public Input<int>? DefaultPort { get; set; }

        /// <summary>
        /// The target description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("hostSourceIds")]
        private InputList<string>? _hostSourceIds;

        /// <summary>
        /// A list of host source ID's.
        /// </summary>
        public InputList<string> HostSourceIds
        {
            get => _hostSourceIds ?? (_hostSourceIds = new InputList<string>());
            set => _hostSourceIds = value;
        }

        [Input("injectedApplicationCredentialSourceIds")]
        private InputList<string>? _injectedApplicationCredentialSourceIds;

        /// <summary>
        /// A list of injected application credential source ID's.
        /// </summary>
        public InputList<string> InjectedApplicationCredentialSourceIds
        {
            get => _injectedApplicationCredentialSourceIds ?? (_injectedApplicationCredentialSourceIds = new InputList<string>());
            set => _injectedApplicationCredentialSourceIds = value;
        }

        /// <summary>
        /// The target name. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
        /// </summary>
        [Input("scopeId", required: true)]
        public Input<string> ScopeId { get; set; } = null!;

        [Input("sessionConnectionLimit")]
        public Input<int>? SessionConnectionLimit { get; set; }

        [Input("sessionMaxSeconds")]
        public Input<int>? SessionMaxSeconds { get; set; }

        /// <summary>
        /// The target resource type.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Boolean expression to filter the workers for this target
        /// </summary>
        [Input("workerFilter")]
        public Input<string>? WorkerFilter { get; set; }

        public BoundaryTargetArgs()
        {
        }
        public static new BoundaryTargetArgs Empty => new BoundaryTargetArgs();
    }

    public sealed class BoundaryTargetState : global::Pulumi.ResourceArgs
    {
        [Input("brokeredCredentialSourceIds")]
        private InputList<string>? _brokeredCredentialSourceIds;

        /// <summary>
        /// A list of brokered credential source ID's.
        /// </summary>
        public InputList<string> BrokeredCredentialSourceIds
        {
            get => _brokeredCredentialSourceIds ?? (_brokeredCredentialSourceIds = new InputList<string>());
            set => _brokeredCredentialSourceIds = value;
        }

        /// <summary>
        /// The default port for this target.
        /// </summary>
        [Input("defaultPort")]
        public Input<int>? DefaultPort { get; set; }

        /// <summary>
        /// The target description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("hostSourceIds")]
        private InputList<string>? _hostSourceIds;

        /// <summary>
        /// A list of host source ID's.
        /// </summary>
        public InputList<string> HostSourceIds
        {
            get => _hostSourceIds ?? (_hostSourceIds = new InputList<string>());
            set => _hostSourceIds = value;
        }

        [Input("injectedApplicationCredentialSourceIds")]
        private InputList<string>? _injectedApplicationCredentialSourceIds;

        /// <summary>
        /// A list of injected application credential source ID's.
        /// </summary>
        public InputList<string> InjectedApplicationCredentialSourceIds
        {
            get => _injectedApplicationCredentialSourceIds ?? (_injectedApplicationCredentialSourceIds = new InputList<string>());
            set => _injectedApplicationCredentialSourceIds = value;
        }

        /// <summary>
        /// The target name. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
        /// </summary>
        [Input("scopeId")]
        public Input<string>? ScopeId { get; set; }

        [Input("sessionConnectionLimit")]
        public Input<int>? SessionConnectionLimit { get; set; }

        [Input("sessionMaxSeconds")]
        public Input<int>? SessionMaxSeconds { get; set; }

        /// <summary>
        /// The target resource type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Boolean expression to filter the workers for this target
        /// </summary>
        [Input("workerFilter")]
        public Input<string>? WorkerFilter { get; set; }

        public BoundaryTargetState()
        {
        }
        public static new BoundaryTargetState Empty => new BoundaryTargetState();
    }
}
