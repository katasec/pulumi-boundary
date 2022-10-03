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
    /// The scope resource allows you to configure a Boundary scope.
    /// 
    /// ## Example Usage
    /// 
    /// Creating the global scope:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Boundary = Katasec.Boundary;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @global = new Boundary.BoundaryScope("global", new()
    ///     {
    ///         GlobalScope = true,
    ///         ScopeId = "global",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Creating an organization scope within global:
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
    ///         ScopeId = boundary_scope.Global.Id,
    ///         AutoCreateAdminRole = true,
    ///         AutoCreateDefaultRole = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Creating an project scope within an organization:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Boundary = Katasec.Boundary;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var project = new Boundary.BoundaryScope("project", new()
    ///     {
    ///         Description = "My first scope!",
    ///         ScopeId = boundary_scope.Org.Id,
    ///         AutoCreateAdminRole = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Creating an organization scope with a managed role for administration (auto create role set false):
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
    ///         ScopeId = boundary_scope.Global.Id,
    ///     });
    /// 
    ///     var orgAdmin = new Boundary.BoundaryRole("orgAdmin", new()
    ///     {
    ///         ScopeId = boundary_scope.Global.Id,
    ///         GrantScopeId = org.Id,
    ///         GrantStrings = new[]
    ///         {
    ///             "id=*;type=*;actions=*",
    ///         },
    ///         PrincipalIds = new[]
    ///         {
    ///             "u_auth",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import boundary:index/boundaryScope:BoundaryScope foo &lt;my-id&gt;
    /// ```
    /// </summary>
    [BoundaryResourceType("boundary:index/boundaryScope:BoundaryScope")]
    public partial class BoundaryScope : global::Pulumi.CustomResource
    {
        /// <summary>
        /// If set, when a new scope is created, the provider will not disable the functionality that automatically creates a role
        /// in the new scope and gives permissions to manage the scope to the provider's user. Marking this true makes for simpler
        /// HCL but results in role resources that are unmanaged by Terraform.
        /// </summary>
        [Output("autoCreateAdminRole")]
        public Output<bool?> AutoCreateAdminRole { get; private set; } = null!;

        /// <summary>
        /// Only relevant when creating an org scope. If set, when a new scope is created, the provider will not disable the
        /// functionality that automatically creates a role in the new scope and gives listing of scopes and auth methods and the
        /// ability to authenticate to the anonymous user. Marking this true makes for simpler HCL but results in role resources
        /// that are unmanaged by Terraform.
        /// </summary>
        [Output("autoCreateDefaultRole")]
        public Output<bool?> AutoCreateDefaultRole { get; private set; } = null!;

        /// <summary>
        /// The scope description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Indicates that the scope containing this value is the global scope, which triggers some specialized behavior to allow it to be imported and managed.
        /// </summary>
        [Output("globalScope")]
        public Output<bool?> GlobalScope { get; private set; } = null!;

        /// <summary>
        /// The scope name. Defaults to the resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The scope ID containing the sub scope resource.
        /// </summary>
        [Output("scopeId")]
        public Output<string> ScopeId { get; private set; } = null!;


        /// <summary>
        /// Create a BoundaryScope resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BoundaryScope(string name, BoundaryScopeArgs args, CustomResourceOptions? options = null)
            : base("boundary:index/boundaryScope:BoundaryScope", name, args ?? new BoundaryScopeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BoundaryScope(string name, Input<string> id, BoundaryScopeState? state = null, CustomResourceOptions? options = null)
            : base("boundary:index/boundaryScope:BoundaryScope", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BoundaryScope resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BoundaryScope Get(string name, Input<string> id, BoundaryScopeState? state = null, CustomResourceOptions? options = null)
        {
            return new BoundaryScope(name, id, state, options);
        }
    }

    public sealed class BoundaryScopeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set, when a new scope is created, the provider will not disable the functionality that automatically creates a role
        /// in the new scope and gives permissions to manage the scope to the provider's user. Marking this true makes for simpler
        /// HCL but results in role resources that are unmanaged by Terraform.
        /// </summary>
        [Input("autoCreateAdminRole")]
        public Input<bool>? AutoCreateAdminRole { get; set; }

        /// <summary>
        /// Only relevant when creating an org scope. If set, when a new scope is created, the provider will not disable the
        /// functionality that automatically creates a role in the new scope and gives listing of scopes and auth methods and the
        /// ability to authenticate to the anonymous user. Marking this true makes for simpler HCL but results in role resources
        /// that are unmanaged by Terraform.
        /// </summary>
        [Input("autoCreateDefaultRole")]
        public Input<bool>? AutoCreateDefaultRole { get; set; }

        /// <summary>
        /// The scope description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates that the scope containing this value is the global scope, which triggers some specialized behavior to allow it to be imported and managed.
        /// </summary>
        [Input("globalScope")]
        public Input<bool>? GlobalScope { get; set; }

        /// <summary>
        /// The scope name. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The scope ID containing the sub scope resource.
        /// </summary>
        [Input("scopeId", required: true)]
        public Input<string> ScopeId { get; set; } = null!;

        public BoundaryScopeArgs()
        {
        }
        public static new BoundaryScopeArgs Empty => new BoundaryScopeArgs();
    }

    public sealed class BoundaryScopeState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set, when a new scope is created, the provider will not disable the functionality that automatically creates a role
        /// in the new scope and gives permissions to manage the scope to the provider's user. Marking this true makes for simpler
        /// HCL but results in role resources that are unmanaged by Terraform.
        /// </summary>
        [Input("autoCreateAdminRole")]
        public Input<bool>? AutoCreateAdminRole { get; set; }

        /// <summary>
        /// Only relevant when creating an org scope. If set, when a new scope is created, the provider will not disable the
        /// functionality that automatically creates a role in the new scope and gives listing of scopes and auth methods and the
        /// ability to authenticate to the anonymous user. Marking this true makes for simpler HCL but results in role resources
        /// that are unmanaged by Terraform.
        /// </summary>
        [Input("autoCreateDefaultRole")]
        public Input<bool>? AutoCreateDefaultRole { get; set; }

        /// <summary>
        /// The scope description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates that the scope containing this value is the global scope, which triggers some specialized behavior to allow it to be imported and managed.
        /// </summary>
        [Input("globalScope")]
        public Input<bool>? GlobalScope { get; set; }

        /// <summary>
        /// The scope name. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The scope ID containing the sub scope resource.
        /// </summary>
        [Input("scopeId")]
        public Input<string>? ScopeId { get; set; }

        public BoundaryScopeState()
        {
        }
        public static new BoundaryScopeState Empty => new BoundaryScopeState();
    }
}
