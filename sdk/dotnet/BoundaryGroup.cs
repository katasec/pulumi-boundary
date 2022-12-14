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
    /// The group resource allows you to configure a Boundary group.
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
    ///     var foo = new Boundary.BoundaryUser("foo", new()
    ///     {
    ///         Description = "foo user",
    ///         ScopeId = org.Id,
    ///     });
    /// 
    ///     var example = new Boundary.BoundaryGroup("example", new()
    ///     {
    ///         Description = "My first group!",
    ///         MemberIds = new[]
    ///         {
    ///             foo.Id,
    ///         },
    ///         ScopeId = org.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Usage for project-specific group:
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
    ///     var foo = new Boundary.BoundaryUser("foo", new()
    ///     {
    ///         Description = "foo user",
    ///         ScopeId = org.Id,
    ///     });
    /// 
    ///     var example = new Boundary.BoundaryGroup("example", new()
    ///     {
    ///         Description = "My first group!",
    ///         MemberIds = new[]
    ///         {
    ///             foo.Id,
    ///         },
    ///         ScopeId = project.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import boundary:index/boundaryGroup:BoundaryGroup foo &lt;my-id&gt;
    /// ```
    /// </summary>
    [BoundaryResourceType("boundary:index/boundaryGroup:BoundaryGroup")]
    public partial class BoundaryGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The group description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Resource IDs for group members, these are most likely boundary users.
        /// </summary>
        [Output("memberIds")]
        public Output<ImmutableArray<string>> MemberIds { get; private set; } = null!;

        /// <summary>
        /// The group name. Defaults to the resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
        /// </summary>
        [Output("scopeId")]
        public Output<string> ScopeId { get; private set; } = null!;


        /// <summary>
        /// Create a BoundaryGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BoundaryGroup(string name, BoundaryGroupArgs args, CustomResourceOptions? options = null)
            : base("boundary:index/boundaryGroup:BoundaryGroup", name, args ?? new BoundaryGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BoundaryGroup(string name, Input<string> id, BoundaryGroupState? state = null, CustomResourceOptions? options = null)
            : base("boundary:index/boundaryGroup:BoundaryGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BoundaryGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BoundaryGroup Get(string name, Input<string> id, BoundaryGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new BoundaryGroup(name, id, state, options);
        }
    }

    public sealed class BoundaryGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The group description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("memberIds")]
        private InputList<string>? _memberIds;

        /// <summary>
        /// Resource IDs for group members, these are most likely boundary users.
        /// </summary>
        public InputList<string> MemberIds
        {
            get => _memberIds ?? (_memberIds = new InputList<string>());
            set => _memberIds = value;
        }

        /// <summary>
        /// The group name. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
        /// </summary>
        [Input("scopeId", required: true)]
        public Input<string> ScopeId { get; set; } = null!;

        public BoundaryGroupArgs()
        {
        }
        public static new BoundaryGroupArgs Empty => new BoundaryGroupArgs();
    }

    public sealed class BoundaryGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The group description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("memberIds")]
        private InputList<string>? _memberIds;

        /// <summary>
        /// Resource IDs for group members, these are most likely boundary users.
        /// </summary>
        public InputList<string> MemberIds
        {
            get => _memberIds ?? (_memberIds = new InputList<string>());
            set => _memberIds = value;
        }

        /// <summary>
        /// The group name. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
        /// </summary>
        [Input("scopeId")]
        public Input<string>? ScopeId { get; set; }

        public BoundaryGroupState()
        {
        }
        public static new BoundaryGroupState Empty => new BoundaryGroupState();
    }
}
