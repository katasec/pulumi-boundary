// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package boundary

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The user resource allows you to configure a Boundary user.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
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
//			password, err := boundary.NewBoundaryAuthMethod(ctx, "password", &boundary.BoundaryAuthMethodArgs{
//				ScopeId: org.ID(),
//				Type:    pulumi.String("password"),
//			})
//			if err != nil {
//				return err
//			}
//			jeffBoundaryAccountPassword, err := boundary.NewBoundaryAccountPassword(ctx, "jeffBoundaryAccountPassword", &boundary.BoundaryAccountPasswordArgs{
//				AuthMethodId: password.ID(),
//				Type:         pulumi.String("password"),
//				LoginName:    pulumi.String("jeff"),
//				Password:     pulumi.String(fmt.Sprintf("$uper$ecure")),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = boundary.NewBoundaryUser(ctx, "jeffBoundaryUser", &boundary.BoundaryUserArgs{
//				Description: pulumi.String("Jeff's user resource"),
//				AccountIds: pulumi.StringArray{
//					jeffBoundaryAccountPassword.ID(),
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
// ## Import
//
// ```sh
//
//	$ pulumi import boundary:index/boundaryUser:BoundaryUser foo <my-id>
//
// ```
type BoundaryUser struct {
	pulumi.CustomResourceState

	// Account ID's to associate with this user resource.
	AccountIds pulumi.StringArrayOutput `pulumi:"accountIds"`
	// The user description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The username. Defaults to the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
	ScopeId pulumi.StringOutput `pulumi:"scopeId"`
}

// NewBoundaryUser registers a new resource with the given unique name, arguments, and options.
func NewBoundaryUser(ctx *pulumi.Context,
	name string, args *BoundaryUserArgs, opts ...pulumi.ResourceOption) (*BoundaryUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ScopeId == nil {
		return nil, errors.New("invalid value for required argument 'ScopeId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource BoundaryUser
	err := ctx.RegisterResource("boundary:index/boundaryUser:BoundaryUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBoundaryUser gets an existing BoundaryUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBoundaryUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BoundaryUserState, opts ...pulumi.ResourceOption) (*BoundaryUser, error) {
	var resource BoundaryUser
	err := ctx.ReadResource("boundary:index/boundaryUser:BoundaryUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BoundaryUser resources.
type boundaryUserState struct {
	// Account ID's to associate with this user resource.
	AccountIds []string `pulumi:"accountIds"`
	// The user description.
	Description *string `pulumi:"description"`
	// The username. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
	ScopeId *string `pulumi:"scopeId"`
}

type BoundaryUserState struct {
	// Account ID's to associate with this user resource.
	AccountIds pulumi.StringArrayInput
	// The user description.
	Description pulumi.StringPtrInput
	// The username. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
	ScopeId pulumi.StringPtrInput
}

func (BoundaryUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*boundaryUserState)(nil)).Elem()
}

type boundaryUserArgs struct {
	// Account ID's to associate with this user resource.
	AccountIds []string `pulumi:"accountIds"`
	// The user description.
	Description *string `pulumi:"description"`
	// The username. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
	ScopeId string `pulumi:"scopeId"`
}

// The set of arguments for constructing a BoundaryUser resource.
type BoundaryUserArgs struct {
	// Account ID's to associate with this user resource.
	AccountIds pulumi.StringArrayInput
	// The user description.
	Description pulumi.StringPtrInput
	// The username. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
	ScopeId pulumi.StringInput
}

func (BoundaryUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*boundaryUserArgs)(nil)).Elem()
}

type BoundaryUserInput interface {
	pulumi.Input

	ToBoundaryUserOutput() BoundaryUserOutput
	ToBoundaryUserOutputWithContext(ctx context.Context) BoundaryUserOutput
}

func (*BoundaryUser) ElementType() reflect.Type {
	return reflect.TypeOf((**BoundaryUser)(nil)).Elem()
}

func (i *BoundaryUser) ToBoundaryUserOutput() BoundaryUserOutput {
	return i.ToBoundaryUserOutputWithContext(context.Background())
}

func (i *BoundaryUser) ToBoundaryUserOutputWithContext(ctx context.Context) BoundaryUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoundaryUserOutput)
}

// BoundaryUserArrayInput is an input type that accepts BoundaryUserArray and BoundaryUserArrayOutput values.
// You can construct a concrete instance of `BoundaryUserArrayInput` via:
//
//	BoundaryUserArray{ BoundaryUserArgs{...} }
type BoundaryUserArrayInput interface {
	pulumi.Input

	ToBoundaryUserArrayOutput() BoundaryUserArrayOutput
	ToBoundaryUserArrayOutputWithContext(context.Context) BoundaryUserArrayOutput
}

type BoundaryUserArray []BoundaryUserInput

func (BoundaryUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BoundaryUser)(nil)).Elem()
}

func (i BoundaryUserArray) ToBoundaryUserArrayOutput() BoundaryUserArrayOutput {
	return i.ToBoundaryUserArrayOutputWithContext(context.Background())
}

func (i BoundaryUserArray) ToBoundaryUserArrayOutputWithContext(ctx context.Context) BoundaryUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoundaryUserArrayOutput)
}

// BoundaryUserMapInput is an input type that accepts BoundaryUserMap and BoundaryUserMapOutput values.
// You can construct a concrete instance of `BoundaryUserMapInput` via:
//
//	BoundaryUserMap{ "key": BoundaryUserArgs{...} }
type BoundaryUserMapInput interface {
	pulumi.Input

	ToBoundaryUserMapOutput() BoundaryUserMapOutput
	ToBoundaryUserMapOutputWithContext(context.Context) BoundaryUserMapOutput
}

type BoundaryUserMap map[string]BoundaryUserInput

func (BoundaryUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BoundaryUser)(nil)).Elem()
}

func (i BoundaryUserMap) ToBoundaryUserMapOutput() BoundaryUserMapOutput {
	return i.ToBoundaryUserMapOutputWithContext(context.Background())
}

func (i BoundaryUserMap) ToBoundaryUserMapOutputWithContext(ctx context.Context) BoundaryUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoundaryUserMapOutput)
}

type BoundaryUserOutput struct{ *pulumi.OutputState }

func (BoundaryUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BoundaryUser)(nil)).Elem()
}

func (o BoundaryUserOutput) ToBoundaryUserOutput() BoundaryUserOutput {
	return o
}

func (o BoundaryUserOutput) ToBoundaryUserOutputWithContext(ctx context.Context) BoundaryUserOutput {
	return o
}

// Account ID's to associate with this user resource.
func (o BoundaryUserOutput) AccountIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BoundaryUser) pulumi.StringArrayOutput { return v.AccountIds }).(pulumi.StringArrayOutput)
}

// The user description.
func (o BoundaryUserOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BoundaryUser) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The username. Defaults to the resource name.
func (o BoundaryUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BoundaryUser) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
func (o BoundaryUserOutput) ScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v *BoundaryUser) pulumi.StringOutput { return v.ScopeId }).(pulumi.StringOutput)
}

type BoundaryUserArrayOutput struct{ *pulumi.OutputState }

func (BoundaryUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BoundaryUser)(nil)).Elem()
}

func (o BoundaryUserArrayOutput) ToBoundaryUserArrayOutput() BoundaryUserArrayOutput {
	return o
}

func (o BoundaryUserArrayOutput) ToBoundaryUserArrayOutputWithContext(ctx context.Context) BoundaryUserArrayOutput {
	return o
}

func (o BoundaryUserArrayOutput) Index(i pulumi.IntInput) BoundaryUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BoundaryUser {
		return vs[0].([]*BoundaryUser)[vs[1].(int)]
	}).(BoundaryUserOutput)
}

type BoundaryUserMapOutput struct{ *pulumi.OutputState }

func (BoundaryUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BoundaryUser)(nil)).Elem()
}

func (o BoundaryUserMapOutput) ToBoundaryUserMapOutput() BoundaryUserMapOutput {
	return o
}

func (o BoundaryUserMapOutput) ToBoundaryUserMapOutputWithContext(ctx context.Context) BoundaryUserMapOutput {
	return o
}

func (o BoundaryUserMapOutput) MapIndex(k pulumi.StringInput) BoundaryUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BoundaryUser {
		return vs[0].(map[string]*BoundaryUser)[vs[1].(string)]
	}).(BoundaryUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BoundaryUserInput)(nil)).Elem(), &BoundaryUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*BoundaryUserArrayInput)(nil)).Elem(), BoundaryUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BoundaryUserMapInput)(nil)).Elem(), BoundaryUserMap{})
	pulumi.RegisterOutputType(BoundaryUserOutput{})
	pulumi.RegisterOutputType(BoundaryUserArrayOutput{})
	pulumi.RegisterOutputType(BoundaryUserMapOutput{})
}
