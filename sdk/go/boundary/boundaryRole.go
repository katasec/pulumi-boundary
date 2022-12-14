// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package boundary

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The role resource allows you to configure a Boundary role.
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
//
//	"github.com/katasec/pulumi-boundary/sdk/go/boundary"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			org, err := boundary.NewBoundaryScope(ctx, "org", &boundary.BoundaryScopeArgs{
//				Description:           pulumi.String("My first scope!"),
//				ScopeId:               pulumi.String("global"),
//				AutoCreateAdminRole:   pulumi.Bool(true),
//				AutoCreateDefaultRole: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = boundary.NewBoundaryRole(ctx, "example", &boundary.BoundaryRoleArgs{
//				Description: pulumi.String("My first role!"),
//				ScopeId:     org.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Usage with a user resource:
//
// ```go
// package main
//
// import (
//
//	"github.com/katasec/pulumi-boundary/sdk/go/boundary"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			org, err := boundary.NewBoundaryScope(ctx, "org", &boundary.BoundaryScopeArgs{
//				Description:           pulumi.String("My first scope!"),
//				ScopeId:               pulumi.String("global"),
//				AutoCreateAdminRole:   pulumi.Bool(true),
//				AutoCreateDefaultRole: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			foo, err := boundary.NewBoundaryUser(ctx, "foo", &boundary.BoundaryUserArgs{
//				ScopeId: org.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			bar, err := boundary.NewBoundaryUser(ctx, "bar", &boundary.BoundaryUserArgs{
//				ScopeId: org.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = boundary.NewBoundaryRole(ctx, "example", &boundary.BoundaryRoleArgs{
//				Description: pulumi.String("My first role!"),
//				PrincipalIds: pulumi.StringArray{
//					foo.ID(),
//					bar.ID(),
//				},
//				ScopeId: org.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Usage with user and grants resource:
//
// ```go
// package main
//
// import (
//
//	"github.com/katasec/pulumi-boundary/sdk/go/boundary"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			org, err := boundary.NewBoundaryScope(ctx, "org", &boundary.BoundaryScopeArgs{
//				Description:           pulumi.String("My first scope!"),
//				ScopeId:               pulumi.String("global"),
//				AutoCreateAdminRole:   pulumi.Bool(true),
//				AutoCreateDefaultRole: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			readonlyBoundaryUser, err := boundary.NewBoundaryUser(ctx, "readonlyBoundaryUser", &boundary.BoundaryUserArgs{
//				Description: pulumi.String("A readonly user"),
//				ScopeId:     org.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = boundary.NewBoundaryRole(ctx, "readonlyBoundaryRole", &boundary.BoundaryRoleArgs{
//				Description: pulumi.String("A readonly role"),
//				PrincipalIds: pulumi.StringArray{
//					readonlyBoundaryUser.ID(),
//				},
//				GrantStrings: pulumi.StringArray{
//					pulumi.String("id=*;type=*;actions=read"),
//				},
//				ScopeId: org.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Usage for a project-specific role:
//
// ```go
// package main
//
// import (
//
//	"github.com/katasec/pulumi-boundary/sdk/go/boundary"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			org, err := boundary.NewBoundaryScope(ctx, "org", &boundary.BoundaryScopeArgs{
//				Description:           pulumi.String("My first scope!"),
//				ScopeId:               pulumi.String("global"),
//				AutoCreateAdminRole:   pulumi.Bool(true),
//				AutoCreateDefaultRole: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			project, err := boundary.NewBoundaryScope(ctx, "project", &boundary.BoundaryScopeArgs{
//				Description:         pulumi.String("My first scope!"),
//				ScopeId:             org.ID(),
//				AutoCreateAdminRole: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			readonlyBoundaryUser, err := boundary.NewBoundaryUser(ctx, "readonlyBoundaryUser", &boundary.BoundaryUserArgs{
//				Description: pulumi.String("A readonly user"),
//				ScopeId:     org.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = boundary.NewBoundaryRole(ctx, "readonlyBoundaryRole", &boundary.BoundaryRoleArgs{
//				Description: pulumi.String("A readonly role"),
//				PrincipalIds: pulumi.StringArray{
//					readonlyBoundaryUser.ID(),
//				},
//				GrantStrings: pulumi.StringArray{
//					pulumi.String("id=*;type=*;actions=read"),
//				},
//				ScopeId: project.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ```sh
//
//	$ pulumi import boundary:index/boundaryRole:BoundaryRole foo <my-id>
//
// ```
type BoundaryRole struct {
	pulumi.CustomResourceState

	// The role description.
	Description  pulumi.StringPtrOutput `pulumi:"description"`
	GrantScopeId pulumi.StringOutput    `pulumi:"grantScopeId"`
	// A list of stringified grants for the role.
	GrantStrings pulumi.StringArrayOutput `pulumi:"grantStrings"`
	// The role name. Defaults to the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of principal (user or group) IDs to add as principals on the role.
	PrincipalIds pulumi.StringArrayOutput `pulumi:"principalIds"`
	// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
	ScopeId pulumi.StringOutput `pulumi:"scopeId"`
}

// NewBoundaryRole registers a new resource with the given unique name, arguments, and options.
func NewBoundaryRole(ctx *pulumi.Context,
	name string, args *BoundaryRoleArgs, opts ...pulumi.ResourceOption) (*BoundaryRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ScopeId == nil {
		return nil, errors.New("invalid value for required argument 'ScopeId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource BoundaryRole
	err := ctx.RegisterResource("boundary:index/boundaryRole:BoundaryRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBoundaryRole gets an existing BoundaryRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBoundaryRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BoundaryRoleState, opts ...pulumi.ResourceOption) (*BoundaryRole, error) {
	var resource BoundaryRole
	err := ctx.ReadResource("boundary:index/boundaryRole:BoundaryRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BoundaryRole resources.
type boundaryRoleState struct {
	// The role description.
	Description  *string `pulumi:"description"`
	GrantScopeId *string `pulumi:"grantScopeId"`
	// A list of stringified grants for the role.
	GrantStrings []string `pulumi:"grantStrings"`
	// The role name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// A list of principal (user or group) IDs to add as principals on the role.
	PrincipalIds []string `pulumi:"principalIds"`
	// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
	ScopeId *string `pulumi:"scopeId"`
}

type BoundaryRoleState struct {
	// The role description.
	Description  pulumi.StringPtrInput
	GrantScopeId pulumi.StringPtrInput
	// A list of stringified grants for the role.
	GrantStrings pulumi.StringArrayInput
	// The role name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// A list of principal (user or group) IDs to add as principals on the role.
	PrincipalIds pulumi.StringArrayInput
	// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
	ScopeId pulumi.StringPtrInput
}

func (BoundaryRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*boundaryRoleState)(nil)).Elem()
}

type boundaryRoleArgs struct {
	// The role description.
	Description  *string `pulumi:"description"`
	GrantScopeId *string `pulumi:"grantScopeId"`
	// A list of stringified grants for the role.
	GrantStrings []string `pulumi:"grantStrings"`
	// The role name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// A list of principal (user or group) IDs to add as principals on the role.
	PrincipalIds []string `pulumi:"principalIds"`
	// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
	ScopeId string `pulumi:"scopeId"`
}

// The set of arguments for constructing a BoundaryRole resource.
type BoundaryRoleArgs struct {
	// The role description.
	Description  pulumi.StringPtrInput
	GrantScopeId pulumi.StringPtrInput
	// A list of stringified grants for the role.
	GrantStrings pulumi.StringArrayInput
	// The role name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// A list of principal (user or group) IDs to add as principals on the role.
	PrincipalIds pulumi.StringArrayInput
	// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
	ScopeId pulumi.StringInput
}

func (BoundaryRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*boundaryRoleArgs)(nil)).Elem()
}

type BoundaryRoleInput interface {
	pulumi.Input

	ToBoundaryRoleOutput() BoundaryRoleOutput
	ToBoundaryRoleOutputWithContext(ctx context.Context) BoundaryRoleOutput
}

func (*BoundaryRole) ElementType() reflect.Type {
	return reflect.TypeOf((**BoundaryRole)(nil)).Elem()
}

func (i *BoundaryRole) ToBoundaryRoleOutput() BoundaryRoleOutput {
	return i.ToBoundaryRoleOutputWithContext(context.Background())
}

func (i *BoundaryRole) ToBoundaryRoleOutputWithContext(ctx context.Context) BoundaryRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoundaryRoleOutput)
}

// BoundaryRoleArrayInput is an input type that accepts BoundaryRoleArray and BoundaryRoleArrayOutput values.
// You can construct a concrete instance of `BoundaryRoleArrayInput` via:
//
//	BoundaryRoleArray{ BoundaryRoleArgs{...} }
type BoundaryRoleArrayInput interface {
	pulumi.Input

	ToBoundaryRoleArrayOutput() BoundaryRoleArrayOutput
	ToBoundaryRoleArrayOutputWithContext(context.Context) BoundaryRoleArrayOutput
}

type BoundaryRoleArray []BoundaryRoleInput

func (BoundaryRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BoundaryRole)(nil)).Elem()
}

func (i BoundaryRoleArray) ToBoundaryRoleArrayOutput() BoundaryRoleArrayOutput {
	return i.ToBoundaryRoleArrayOutputWithContext(context.Background())
}

func (i BoundaryRoleArray) ToBoundaryRoleArrayOutputWithContext(ctx context.Context) BoundaryRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoundaryRoleArrayOutput)
}

// BoundaryRoleMapInput is an input type that accepts BoundaryRoleMap and BoundaryRoleMapOutput values.
// You can construct a concrete instance of `BoundaryRoleMapInput` via:
//
//	BoundaryRoleMap{ "key": BoundaryRoleArgs{...} }
type BoundaryRoleMapInput interface {
	pulumi.Input

	ToBoundaryRoleMapOutput() BoundaryRoleMapOutput
	ToBoundaryRoleMapOutputWithContext(context.Context) BoundaryRoleMapOutput
}

type BoundaryRoleMap map[string]BoundaryRoleInput

func (BoundaryRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BoundaryRole)(nil)).Elem()
}

func (i BoundaryRoleMap) ToBoundaryRoleMapOutput() BoundaryRoleMapOutput {
	return i.ToBoundaryRoleMapOutputWithContext(context.Background())
}

func (i BoundaryRoleMap) ToBoundaryRoleMapOutputWithContext(ctx context.Context) BoundaryRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoundaryRoleMapOutput)
}

type BoundaryRoleOutput struct{ *pulumi.OutputState }

func (BoundaryRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BoundaryRole)(nil)).Elem()
}

func (o BoundaryRoleOutput) ToBoundaryRoleOutput() BoundaryRoleOutput {
	return o
}

func (o BoundaryRoleOutput) ToBoundaryRoleOutputWithContext(ctx context.Context) BoundaryRoleOutput {
	return o
}

// The role description.
func (o BoundaryRoleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BoundaryRole) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o BoundaryRoleOutput) GrantScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v *BoundaryRole) pulumi.StringOutput { return v.GrantScopeId }).(pulumi.StringOutput)
}

// A list of stringified grants for the role.
func (o BoundaryRoleOutput) GrantStrings() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BoundaryRole) pulumi.StringArrayOutput { return v.GrantStrings }).(pulumi.StringArrayOutput)
}

// The role name. Defaults to the resource name.
func (o BoundaryRoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BoundaryRole) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A list of principal (user or group) IDs to add as principals on the role.
func (o BoundaryRoleOutput) PrincipalIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BoundaryRole) pulumi.StringArrayOutput { return v.PrincipalIds }).(pulumi.StringArrayOutput)
}

// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
func (o BoundaryRoleOutput) ScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v *BoundaryRole) pulumi.StringOutput { return v.ScopeId }).(pulumi.StringOutput)
}

type BoundaryRoleArrayOutput struct{ *pulumi.OutputState }

func (BoundaryRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BoundaryRole)(nil)).Elem()
}

func (o BoundaryRoleArrayOutput) ToBoundaryRoleArrayOutput() BoundaryRoleArrayOutput {
	return o
}

func (o BoundaryRoleArrayOutput) ToBoundaryRoleArrayOutputWithContext(ctx context.Context) BoundaryRoleArrayOutput {
	return o
}

func (o BoundaryRoleArrayOutput) Index(i pulumi.IntInput) BoundaryRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BoundaryRole {
		return vs[0].([]*BoundaryRole)[vs[1].(int)]
	}).(BoundaryRoleOutput)
}

type BoundaryRoleMapOutput struct{ *pulumi.OutputState }

func (BoundaryRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BoundaryRole)(nil)).Elem()
}

func (o BoundaryRoleMapOutput) ToBoundaryRoleMapOutput() BoundaryRoleMapOutput {
	return o
}

func (o BoundaryRoleMapOutput) ToBoundaryRoleMapOutputWithContext(ctx context.Context) BoundaryRoleMapOutput {
	return o
}

func (o BoundaryRoleMapOutput) MapIndex(k pulumi.StringInput) BoundaryRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BoundaryRole {
		return vs[0].(map[string]*BoundaryRole)[vs[1].(string)]
	}).(BoundaryRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BoundaryRoleInput)(nil)).Elem(), &BoundaryRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*BoundaryRoleArrayInput)(nil)).Elem(), BoundaryRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BoundaryRoleMapInput)(nil)).Elem(), BoundaryRoleMap{})
	pulumi.RegisterOutputType(BoundaryRoleOutput{})
	pulumi.RegisterOutputType(BoundaryRoleArrayOutput{})
	pulumi.RegisterOutputType(BoundaryRoleMapOutput{})
}
