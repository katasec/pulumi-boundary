// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package boundary

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Deprecated: use `BoundaryHostCatalogStatic` instead.
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
//				ScopeId:               pulumi.Any(boundary_scope.Global.Id),
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
//			_, err = boundary.NewBoundaryHostCatalog(ctx, "example", &boundary.BoundaryHostCatalogArgs{
//				Description: pulumi.String("My first host catalog!"),
//				Type:        pulumi.String("Static"),
//				ScopeId:     project.ID(),
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
//	$ pulumi import boundary:index/boundaryHostCatalog:BoundaryHostCatalog foo <my-id>
//
// ```
type BoundaryHostCatalog struct {
	pulumi.CustomResourceState

	// The host catalog description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The host catalog name. Defaults to the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The scope ID in which the resource is created.
	ScopeId pulumi.StringOutput `pulumi:"scopeId"`
	// The host catalog type. Only `static` is supported.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBoundaryHostCatalog registers a new resource with the given unique name, arguments, and options.
func NewBoundaryHostCatalog(ctx *pulumi.Context,
	name string, args *BoundaryHostCatalogArgs, opts ...pulumi.ResourceOption) (*BoundaryHostCatalog, error) {
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
	var resource BoundaryHostCatalog
	err := ctx.RegisterResource("boundary:index/boundaryHostCatalog:BoundaryHostCatalog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBoundaryHostCatalog gets an existing BoundaryHostCatalog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBoundaryHostCatalog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BoundaryHostCatalogState, opts ...pulumi.ResourceOption) (*BoundaryHostCatalog, error) {
	var resource BoundaryHostCatalog
	err := ctx.ReadResource("boundary:index/boundaryHostCatalog:BoundaryHostCatalog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BoundaryHostCatalog resources.
type boundaryHostCatalogState struct {
	// The host catalog description.
	Description *string `pulumi:"description"`
	// The host catalog name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The scope ID in which the resource is created.
	ScopeId *string `pulumi:"scopeId"`
	// The host catalog type. Only `static` is supported.
	Type *string `pulumi:"type"`
}

type BoundaryHostCatalogState struct {
	// The host catalog description.
	Description pulumi.StringPtrInput
	// The host catalog name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The scope ID in which the resource is created.
	ScopeId pulumi.StringPtrInput
	// The host catalog type. Only `static` is supported.
	Type pulumi.StringPtrInput
}

func (BoundaryHostCatalogState) ElementType() reflect.Type {
	return reflect.TypeOf((*boundaryHostCatalogState)(nil)).Elem()
}

type boundaryHostCatalogArgs struct {
	// The host catalog description.
	Description *string `pulumi:"description"`
	// The host catalog name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The scope ID in which the resource is created.
	ScopeId string `pulumi:"scopeId"`
	// The host catalog type. Only `static` is supported.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a BoundaryHostCatalog resource.
type BoundaryHostCatalogArgs struct {
	// The host catalog description.
	Description pulumi.StringPtrInput
	// The host catalog name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The scope ID in which the resource is created.
	ScopeId pulumi.StringInput
	// The host catalog type. Only `static` is supported.
	Type pulumi.StringInput
}

func (BoundaryHostCatalogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*boundaryHostCatalogArgs)(nil)).Elem()
}

type BoundaryHostCatalogInput interface {
	pulumi.Input

	ToBoundaryHostCatalogOutput() BoundaryHostCatalogOutput
	ToBoundaryHostCatalogOutputWithContext(ctx context.Context) BoundaryHostCatalogOutput
}

func (*BoundaryHostCatalog) ElementType() reflect.Type {
	return reflect.TypeOf((**BoundaryHostCatalog)(nil)).Elem()
}

func (i *BoundaryHostCatalog) ToBoundaryHostCatalogOutput() BoundaryHostCatalogOutput {
	return i.ToBoundaryHostCatalogOutputWithContext(context.Background())
}

func (i *BoundaryHostCatalog) ToBoundaryHostCatalogOutputWithContext(ctx context.Context) BoundaryHostCatalogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoundaryHostCatalogOutput)
}

// BoundaryHostCatalogArrayInput is an input type that accepts BoundaryHostCatalogArray and BoundaryHostCatalogArrayOutput values.
// You can construct a concrete instance of `BoundaryHostCatalogArrayInput` via:
//
//	BoundaryHostCatalogArray{ BoundaryHostCatalogArgs{...} }
type BoundaryHostCatalogArrayInput interface {
	pulumi.Input

	ToBoundaryHostCatalogArrayOutput() BoundaryHostCatalogArrayOutput
	ToBoundaryHostCatalogArrayOutputWithContext(context.Context) BoundaryHostCatalogArrayOutput
}

type BoundaryHostCatalogArray []BoundaryHostCatalogInput

func (BoundaryHostCatalogArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BoundaryHostCatalog)(nil)).Elem()
}

func (i BoundaryHostCatalogArray) ToBoundaryHostCatalogArrayOutput() BoundaryHostCatalogArrayOutput {
	return i.ToBoundaryHostCatalogArrayOutputWithContext(context.Background())
}

func (i BoundaryHostCatalogArray) ToBoundaryHostCatalogArrayOutputWithContext(ctx context.Context) BoundaryHostCatalogArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoundaryHostCatalogArrayOutput)
}

// BoundaryHostCatalogMapInput is an input type that accepts BoundaryHostCatalogMap and BoundaryHostCatalogMapOutput values.
// You can construct a concrete instance of `BoundaryHostCatalogMapInput` via:
//
//	BoundaryHostCatalogMap{ "key": BoundaryHostCatalogArgs{...} }
type BoundaryHostCatalogMapInput interface {
	pulumi.Input

	ToBoundaryHostCatalogMapOutput() BoundaryHostCatalogMapOutput
	ToBoundaryHostCatalogMapOutputWithContext(context.Context) BoundaryHostCatalogMapOutput
}

type BoundaryHostCatalogMap map[string]BoundaryHostCatalogInput

func (BoundaryHostCatalogMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BoundaryHostCatalog)(nil)).Elem()
}

func (i BoundaryHostCatalogMap) ToBoundaryHostCatalogMapOutput() BoundaryHostCatalogMapOutput {
	return i.ToBoundaryHostCatalogMapOutputWithContext(context.Background())
}

func (i BoundaryHostCatalogMap) ToBoundaryHostCatalogMapOutputWithContext(ctx context.Context) BoundaryHostCatalogMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoundaryHostCatalogMapOutput)
}

type BoundaryHostCatalogOutput struct{ *pulumi.OutputState }

func (BoundaryHostCatalogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BoundaryHostCatalog)(nil)).Elem()
}

func (o BoundaryHostCatalogOutput) ToBoundaryHostCatalogOutput() BoundaryHostCatalogOutput {
	return o
}

func (o BoundaryHostCatalogOutput) ToBoundaryHostCatalogOutputWithContext(ctx context.Context) BoundaryHostCatalogOutput {
	return o
}

// The host catalog description.
func (o BoundaryHostCatalogOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BoundaryHostCatalog) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The host catalog name. Defaults to the resource name.
func (o BoundaryHostCatalogOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BoundaryHostCatalog) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The scope ID in which the resource is created.
func (o BoundaryHostCatalogOutput) ScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v *BoundaryHostCatalog) pulumi.StringOutput { return v.ScopeId }).(pulumi.StringOutput)
}

// The host catalog type. Only `static` is supported.
func (o BoundaryHostCatalogOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BoundaryHostCatalog) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type BoundaryHostCatalogArrayOutput struct{ *pulumi.OutputState }

func (BoundaryHostCatalogArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BoundaryHostCatalog)(nil)).Elem()
}

func (o BoundaryHostCatalogArrayOutput) ToBoundaryHostCatalogArrayOutput() BoundaryHostCatalogArrayOutput {
	return o
}

func (o BoundaryHostCatalogArrayOutput) ToBoundaryHostCatalogArrayOutputWithContext(ctx context.Context) BoundaryHostCatalogArrayOutput {
	return o
}

func (o BoundaryHostCatalogArrayOutput) Index(i pulumi.IntInput) BoundaryHostCatalogOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BoundaryHostCatalog {
		return vs[0].([]*BoundaryHostCatalog)[vs[1].(int)]
	}).(BoundaryHostCatalogOutput)
}

type BoundaryHostCatalogMapOutput struct{ *pulumi.OutputState }

func (BoundaryHostCatalogMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BoundaryHostCatalog)(nil)).Elem()
}

func (o BoundaryHostCatalogMapOutput) ToBoundaryHostCatalogMapOutput() BoundaryHostCatalogMapOutput {
	return o
}

func (o BoundaryHostCatalogMapOutput) ToBoundaryHostCatalogMapOutputWithContext(ctx context.Context) BoundaryHostCatalogMapOutput {
	return o
}

func (o BoundaryHostCatalogMapOutput) MapIndex(k pulumi.StringInput) BoundaryHostCatalogOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BoundaryHostCatalog {
		return vs[0].(map[string]*BoundaryHostCatalog)[vs[1].(string)]
	}).(BoundaryHostCatalogOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BoundaryHostCatalogInput)(nil)).Elem(), &BoundaryHostCatalog{})
	pulumi.RegisterInputType(reflect.TypeOf((*BoundaryHostCatalogArrayInput)(nil)).Elem(), BoundaryHostCatalogArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BoundaryHostCatalogMapInput)(nil)).Elem(), BoundaryHostCatalogMap{})
	pulumi.RegisterOutputType(BoundaryHostCatalogOutput{})
	pulumi.RegisterOutputType(BoundaryHostCatalogArrayOutput{})
	pulumi.RegisterOutputType(BoundaryHostCatalogMapOutput{})
}