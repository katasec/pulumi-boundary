// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package boundary

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The auth method resource allows you to configure a Boundary auth_method.
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
//			_, err = boundary.NewBoundaryAuthMethod(ctx, "password", &boundary.BoundaryAuthMethodArgs{
//				ScopeId: org.ID(),
//				Type:    pulumi.String("password"),
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
//	$ pulumi import boundary:index/boundaryAuthMethod:BoundaryAuthMethod foo <my-id>
//
// ```
type BoundaryAuthMethod struct {
	pulumi.CustomResourceState

	// The auth method description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The minimum login name length.
	//
	// Deprecated: Will be removed in favor of using attributes parameter
	MinLoginNameLength pulumi.IntOutput `pulumi:"minLoginNameLength"`
	// The minimum password length.
	//
	// Deprecated: Will be removed in favor of using attributes parameter
	MinPasswordLength pulumi.IntOutput `pulumi:"minPasswordLength"`
	// The auth method name. Defaults to the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The scope ID.
	ScopeId pulumi.StringOutput `pulumi:"scopeId"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBoundaryAuthMethod registers a new resource with the given unique name, arguments, and options.
func NewBoundaryAuthMethod(ctx *pulumi.Context,
	name string, args *BoundaryAuthMethodArgs, opts ...pulumi.ResourceOption) (*BoundaryAuthMethod, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ScopeId == nil {
		return nil, errors.New("invalid value for required argument 'ScopeId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource BoundaryAuthMethod
	err := ctx.RegisterResource("boundary:index/boundaryAuthMethod:BoundaryAuthMethod", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBoundaryAuthMethod gets an existing BoundaryAuthMethod resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBoundaryAuthMethod(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BoundaryAuthMethodState, opts ...pulumi.ResourceOption) (*BoundaryAuthMethod, error) {
	var resource BoundaryAuthMethod
	err := ctx.ReadResource("boundary:index/boundaryAuthMethod:BoundaryAuthMethod", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BoundaryAuthMethod resources.
type boundaryAuthMethodState struct {
	// The auth method description.
	Description *string `pulumi:"description"`
	// The minimum login name length.
	//
	// Deprecated: Will be removed in favor of using attributes parameter
	MinLoginNameLength *int `pulumi:"minLoginNameLength"`
	// The minimum password length.
	//
	// Deprecated: Will be removed in favor of using attributes parameter
	MinPasswordLength *int `pulumi:"minPasswordLength"`
	// The auth method name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The scope ID.
	ScopeId *string `pulumi:"scopeId"`
	// The resource type.
	Type *string `pulumi:"type"`
}

type BoundaryAuthMethodState struct {
	// The auth method description.
	Description pulumi.StringPtrInput
	// The minimum login name length.
	//
	// Deprecated: Will be removed in favor of using attributes parameter
	MinLoginNameLength pulumi.IntPtrInput
	// The minimum password length.
	//
	// Deprecated: Will be removed in favor of using attributes parameter
	MinPasswordLength pulumi.IntPtrInput
	// The auth method name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The scope ID.
	ScopeId pulumi.StringPtrInput
	// The resource type.
	Type pulumi.StringPtrInput
}

func (BoundaryAuthMethodState) ElementType() reflect.Type {
	return reflect.TypeOf((*boundaryAuthMethodState)(nil)).Elem()
}

type boundaryAuthMethodArgs struct {
	// The auth method description.
	Description *string `pulumi:"description"`
	// The minimum login name length.
	//
	// Deprecated: Will be removed in favor of using attributes parameter
	MinLoginNameLength *int `pulumi:"minLoginNameLength"`
	// The minimum password length.
	//
	// Deprecated: Will be removed in favor of using attributes parameter
	MinPasswordLength *int `pulumi:"minPasswordLength"`
	// The auth method name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The scope ID.
	ScopeId string `pulumi:"scopeId"`
	// The resource type.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a BoundaryAuthMethod resource.
type BoundaryAuthMethodArgs struct {
	// The auth method description.
	Description pulumi.StringPtrInput
	// The minimum login name length.
	//
	// Deprecated: Will be removed in favor of using attributes parameter
	MinLoginNameLength pulumi.IntPtrInput
	// The minimum password length.
	//
	// Deprecated: Will be removed in favor of using attributes parameter
	MinPasswordLength pulumi.IntPtrInput
	// The auth method name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The scope ID.
	ScopeId pulumi.StringInput
	// The resource type.
	Type pulumi.StringInput
}

func (BoundaryAuthMethodArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*boundaryAuthMethodArgs)(nil)).Elem()
}

type BoundaryAuthMethodInput interface {
	pulumi.Input

	ToBoundaryAuthMethodOutput() BoundaryAuthMethodOutput
	ToBoundaryAuthMethodOutputWithContext(ctx context.Context) BoundaryAuthMethodOutput
}

func (*BoundaryAuthMethod) ElementType() reflect.Type {
	return reflect.TypeOf((**BoundaryAuthMethod)(nil)).Elem()
}

func (i *BoundaryAuthMethod) ToBoundaryAuthMethodOutput() BoundaryAuthMethodOutput {
	return i.ToBoundaryAuthMethodOutputWithContext(context.Background())
}

func (i *BoundaryAuthMethod) ToBoundaryAuthMethodOutputWithContext(ctx context.Context) BoundaryAuthMethodOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoundaryAuthMethodOutput)
}

// BoundaryAuthMethodArrayInput is an input type that accepts BoundaryAuthMethodArray and BoundaryAuthMethodArrayOutput values.
// You can construct a concrete instance of `BoundaryAuthMethodArrayInput` via:
//
//	BoundaryAuthMethodArray{ BoundaryAuthMethodArgs{...} }
type BoundaryAuthMethodArrayInput interface {
	pulumi.Input

	ToBoundaryAuthMethodArrayOutput() BoundaryAuthMethodArrayOutput
	ToBoundaryAuthMethodArrayOutputWithContext(context.Context) BoundaryAuthMethodArrayOutput
}

type BoundaryAuthMethodArray []BoundaryAuthMethodInput

func (BoundaryAuthMethodArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BoundaryAuthMethod)(nil)).Elem()
}

func (i BoundaryAuthMethodArray) ToBoundaryAuthMethodArrayOutput() BoundaryAuthMethodArrayOutput {
	return i.ToBoundaryAuthMethodArrayOutputWithContext(context.Background())
}

func (i BoundaryAuthMethodArray) ToBoundaryAuthMethodArrayOutputWithContext(ctx context.Context) BoundaryAuthMethodArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoundaryAuthMethodArrayOutput)
}

// BoundaryAuthMethodMapInput is an input type that accepts BoundaryAuthMethodMap and BoundaryAuthMethodMapOutput values.
// You can construct a concrete instance of `BoundaryAuthMethodMapInput` via:
//
//	BoundaryAuthMethodMap{ "key": BoundaryAuthMethodArgs{...} }
type BoundaryAuthMethodMapInput interface {
	pulumi.Input

	ToBoundaryAuthMethodMapOutput() BoundaryAuthMethodMapOutput
	ToBoundaryAuthMethodMapOutputWithContext(context.Context) BoundaryAuthMethodMapOutput
}

type BoundaryAuthMethodMap map[string]BoundaryAuthMethodInput

func (BoundaryAuthMethodMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BoundaryAuthMethod)(nil)).Elem()
}

func (i BoundaryAuthMethodMap) ToBoundaryAuthMethodMapOutput() BoundaryAuthMethodMapOutput {
	return i.ToBoundaryAuthMethodMapOutputWithContext(context.Background())
}

func (i BoundaryAuthMethodMap) ToBoundaryAuthMethodMapOutputWithContext(ctx context.Context) BoundaryAuthMethodMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoundaryAuthMethodMapOutput)
}

type BoundaryAuthMethodOutput struct{ *pulumi.OutputState }

func (BoundaryAuthMethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BoundaryAuthMethod)(nil)).Elem()
}

func (o BoundaryAuthMethodOutput) ToBoundaryAuthMethodOutput() BoundaryAuthMethodOutput {
	return o
}

func (o BoundaryAuthMethodOutput) ToBoundaryAuthMethodOutputWithContext(ctx context.Context) BoundaryAuthMethodOutput {
	return o
}

// The auth method description.
func (o BoundaryAuthMethodOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BoundaryAuthMethod) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The minimum login name length.
//
// Deprecated: Will be removed in favor of using attributes parameter
func (o BoundaryAuthMethodOutput) MinLoginNameLength() pulumi.IntOutput {
	return o.ApplyT(func(v *BoundaryAuthMethod) pulumi.IntOutput { return v.MinLoginNameLength }).(pulumi.IntOutput)
}

// The minimum password length.
//
// Deprecated: Will be removed in favor of using attributes parameter
func (o BoundaryAuthMethodOutput) MinPasswordLength() pulumi.IntOutput {
	return o.ApplyT(func(v *BoundaryAuthMethod) pulumi.IntOutput { return v.MinPasswordLength }).(pulumi.IntOutput)
}

// The auth method name. Defaults to the resource name.
func (o BoundaryAuthMethodOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BoundaryAuthMethod) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The scope ID.
func (o BoundaryAuthMethodOutput) ScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v *BoundaryAuthMethod) pulumi.StringOutput { return v.ScopeId }).(pulumi.StringOutput)
}

// The resource type.
func (o BoundaryAuthMethodOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BoundaryAuthMethod) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type BoundaryAuthMethodArrayOutput struct{ *pulumi.OutputState }

func (BoundaryAuthMethodArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BoundaryAuthMethod)(nil)).Elem()
}

func (o BoundaryAuthMethodArrayOutput) ToBoundaryAuthMethodArrayOutput() BoundaryAuthMethodArrayOutput {
	return o
}

func (o BoundaryAuthMethodArrayOutput) ToBoundaryAuthMethodArrayOutputWithContext(ctx context.Context) BoundaryAuthMethodArrayOutput {
	return o
}

func (o BoundaryAuthMethodArrayOutput) Index(i pulumi.IntInput) BoundaryAuthMethodOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BoundaryAuthMethod {
		return vs[0].([]*BoundaryAuthMethod)[vs[1].(int)]
	}).(BoundaryAuthMethodOutput)
}

type BoundaryAuthMethodMapOutput struct{ *pulumi.OutputState }

func (BoundaryAuthMethodMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BoundaryAuthMethod)(nil)).Elem()
}

func (o BoundaryAuthMethodMapOutput) ToBoundaryAuthMethodMapOutput() BoundaryAuthMethodMapOutput {
	return o
}

func (o BoundaryAuthMethodMapOutput) ToBoundaryAuthMethodMapOutputWithContext(ctx context.Context) BoundaryAuthMethodMapOutput {
	return o
}

func (o BoundaryAuthMethodMapOutput) MapIndex(k pulumi.StringInput) BoundaryAuthMethodOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BoundaryAuthMethod {
		return vs[0].(map[string]*BoundaryAuthMethod)[vs[1].(string)]
	}).(BoundaryAuthMethodOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BoundaryAuthMethodInput)(nil)).Elem(), &BoundaryAuthMethod{})
	pulumi.RegisterInputType(reflect.TypeOf((*BoundaryAuthMethodArrayInput)(nil)).Elem(), BoundaryAuthMethodArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BoundaryAuthMethodMapInput)(nil)).Elem(), BoundaryAuthMethodMap{})
	pulumi.RegisterOutputType(BoundaryAuthMethodOutput{})
	pulumi.RegisterOutputType(BoundaryAuthMethodArrayOutput{})
	pulumi.RegisterOutputType(BoundaryAuthMethodMapOutput{})
}
