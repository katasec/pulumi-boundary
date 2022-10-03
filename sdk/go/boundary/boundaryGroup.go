// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package boundary

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The group resource allows you to configure a Boundary group.
//
// ## Example Usage
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
//				Description: pulumi.String("foo user"),
//				ScopeId:     org.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = boundary.NewBoundaryGroup(ctx, "example", &boundary.BoundaryGroupArgs{
//				Description: pulumi.String("My first group!"),
//				MemberIds: pulumi.StringArray{
//					foo.ID(),
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
// Usage for project-specific group:
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
//			foo, err := boundary.NewBoundaryUser(ctx, "foo", &boundary.BoundaryUserArgs{
//				Description: pulumi.String("foo user"),
//				ScopeId:     org.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = boundary.NewBoundaryGroup(ctx, "example", &boundary.BoundaryGroupArgs{
//				Description: pulumi.String("My first group!"),
//				MemberIds: pulumi.StringArray{
//					foo.ID(),
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
//	$ pulumi import boundary:index/boundaryGroup:BoundaryGroup foo <my-id>
//
// ```
type BoundaryGroup struct {
	pulumi.CustomResourceState

	// The group description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Resource IDs for group members, these are most likely boundary users.
	MemberIds pulumi.StringArrayOutput `pulumi:"memberIds"`
	// The group name. Defaults to the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
	ScopeId pulumi.StringOutput `pulumi:"scopeId"`
}

// NewBoundaryGroup registers a new resource with the given unique name, arguments, and options.
func NewBoundaryGroup(ctx *pulumi.Context,
	name string, args *BoundaryGroupArgs, opts ...pulumi.ResourceOption) (*BoundaryGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ScopeId == nil {
		return nil, errors.New("invalid value for required argument 'ScopeId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource BoundaryGroup
	err := ctx.RegisterResource("boundary:index/boundaryGroup:BoundaryGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBoundaryGroup gets an existing BoundaryGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBoundaryGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BoundaryGroupState, opts ...pulumi.ResourceOption) (*BoundaryGroup, error) {
	var resource BoundaryGroup
	err := ctx.ReadResource("boundary:index/boundaryGroup:BoundaryGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BoundaryGroup resources.
type boundaryGroupState struct {
	// The group description.
	Description *string `pulumi:"description"`
	// Resource IDs for group members, these are most likely boundary users.
	MemberIds []string `pulumi:"memberIds"`
	// The group name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
	ScopeId *string `pulumi:"scopeId"`
}

type BoundaryGroupState struct {
	// The group description.
	Description pulumi.StringPtrInput
	// Resource IDs for group members, these are most likely boundary users.
	MemberIds pulumi.StringArrayInput
	// The group name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
	ScopeId pulumi.StringPtrInput
}

func (BoundaryGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*boundaryGroupState)(nil)).Elem()
}

type boundaryGroupArgs struct {
	// The group description.
	Description *string `pulumi:"description"`
	// Resource IDs for group members, these are most likely boundary users.
	MemberIds []string `pulumi:"memberIds"`
	// The group name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
	ScopeId string `pulumi:"scopeId"`
}

// The set of arguments for constructing a BoundaryGroup resource.
type BoundaryGroupArgs struct {
	// The group description.
	Description pulumi.StringPtrInput
	// Resource IDs for group members, these are most likely boundary users.
	MemberIds pulumi.StringArrayInput
	// The group name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
	ScopeId pulumi.StringInput
}

func (BoundaryGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*boundaryGroupArgs)(nil)).Elem()
}

type BoundaryGroupInput interface {
	pulumi.Input

	ToBoundaryGroupOutput() BoundaryGroupOutput
	ToBoundaryGroupOutputWithContext(ctx context.Context) BoundaryGroupOutput
}

func (*BoundaryGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**BoundaryGroup)(nil)).Elem()
}

func (i *BoundaryGroup) ToBoundaryGroupOutput() BoundaryGroupOutput {
	return i.ToBoundaryGroupOutputWithContext(context.Background())
}

func (i *BoundaryGroup) ToBoundaryGroupOutputWithContext(ctx context.Context) BoundaryGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoundaryGroupOutput)
}

// BoundaryGroupArrayInput is an input type that accepts BoundaryGroupArray and BoundaryGroupArrayOutput values.
// You can construct a concrete instance of `BoundaryGroupArrayInput` via:
//
//	BoundaryGroupArray{ BoundaryGroupArgs{...} }
type BoundaryGroupArrayInput interface {
	pulumi.Input

	ToBoundaryGroupArrayOutput() BoundaryGroupArrayOutput
	ToBoundaryGroupArrayOutputWithContext(context.Context) BoundaryGroupArrayOutput
}

type BoundaryGroupArray []BoundaryGroupInput

func (BoundaryGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BoundaryGroup)(nil)).Elem()
}

func (i BoundaryGroupArray) ToBoundaryGroupArrayOutput() BoundaryGroupArrayOutput {
	return i.ToBoundaryGroupArrayOutputWithContext(context.Background())
}

func (i BoundaryGroupArray) ToBoundaryGroupArrayOutputWithContext(ctx context.Context) BoundaryGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoundaryGroupArrayOutput)
}

// BoundaryGroupMapInput is an input type that accepts BoundaryGroupMap and BoundaryGroupMapOutput values.
// You can construct a concrete instance of `BoundaryGroupMapInput` via:
//
//	BoundaryGroupMap{ "key": BoundaryGroupArgs{...} }
type BoundaryGroupMapInput interface {
	pulumi.Input

	ToBoundaryGroupMapOutput() BoundaryGroupMapOutput
	ToBoundaryGroupMapOutputWithContext(context.Context) BoundaryGroupMapOutput
}

type BoundaryGroupMap map[string]BoundaryGroupInput

func (BoundaryGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BoundaryGroup)(nil)).Elem()
}

func (i BoundaryGroupMap) ToBoundaryGroupMapOutput() BoundaryGroupMapOutput {
	return i.ToBoundaryGroupMapOutputWithContext(context.Background())
}

func (i BoundaryGroupMap) ToBoundaryGroupMapOutputWithContext(ctx context.Context) BoundaryGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoundaryGroupMapOutput)
}

type BoundaryGroupOutput struct{ *pulumi.OutputState }

func (BoundaryGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BoundaryGroup)(nil)).Elem()
}

func (o BoundaryGroupOutput) ToBoundaryGroupOutput() BoundaryGroupOutput {
	return o
}

func (o BoundaryGroupOutput) ToBoundaryGroupOutputWithContext(ctx context.Context) BoundaryGroupOutput {
	return o
}

// The group description.
func (o BoundaryGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BoundaryGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Resource IDs for group members, these are most likely boundary users.
func (o BoundaryGroupOutput) MemberIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BoundaryGroup) pulumi.StringArrayOutput { return v.MemberIds }).(pulumi.StringArrayOutput)
}

// The group name. Defaults to the resource name.
func (o BoundaryGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BoundaryGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
func (o BoundaryGroupOutput) ScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v *BoundaryGroup) pulumi.StringOutput { return v.ScopeId }).(pulumi.StringOutput)
}

type BoundaryGroupArrayOutput struct{ *pulumi.OutputState }

func (BoundaryGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BoundaryGroup)(nil)).Elem()
}

func (o BoundaryGroupArrayOutput) ToBoundaryGroupArrayOutput() BoundaryGroupArrayOutput {
	return o
}

func (o BoundaryGroupArrayOutput) ToBoundaryGroupArrayOutputWithContext(ctx context.Context) BoundaryGroupArrayOutput {
	return o
}

func (o BoundaryGroupArrayOutput) Index(i pulumi.IntInput) BoundaryGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BoundaryGroup {
		return vs[0].([]*BoundaryGroup)[vs[1].(int)]
	}).(BoundaryGroupOutput)
}

type BoundaryGroupMapOutput struct{ *pulumi.OutputState }

func (BoundaryGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BoundaryGroup)(nil)).Elem()
}

func (o BoundaryGroupMapOutput) ToBoundaryGroupMapOutput() BoundaryGroupMapOutput {
	return o
}

func (o BoundaryGroupMapOutput) ToBoundaryGroupMapOutputWithContext(ctx context.Context) BoundaryGroupMapOutput {
	return o
}

func (o BoundaryGroupMapOutput) MapIndex(k pulumi.StringInput) BoundaryGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BoundaryGroup {
		return vs[0].(map[string]*BoundaryGroup)[vs[1].(string)]
	}).(BoundaryGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BoundaryGroupInput)(nil)).Elem(), &BoundaryGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*BoundaryGroupArrayInput)(nil)).Elem(), BoundaryGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BoundaryGroupMapInput)(nil)).Elem(), BoundaryGroupMap{})
	pulumi.RegisterOutputType(BoundaryGroupOutput{})
	pulumi.RegisterOutputType(BoundaryGroupArrayOutput{})
	pulumi.RegisterOutputType(BoundaryGroupMapOutput{})
}