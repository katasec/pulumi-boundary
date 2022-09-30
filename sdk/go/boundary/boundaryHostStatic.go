// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package boundary

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The static host resource allows you to configure a Boundary static host. Hosts are always part of a project, so a project resource should be used inline or you should have the project ID in hand to successfully configure a host.
type BoundaryHostStatic struct {
	pulumi.CustomResourceState

	// The static address of the host resource as `<IP>` (note: port assignment occurs in the target resource definition, do not add :port here) or a domain name.
	Address pulumi.StringPtrOutput `pulumi:"address"`
	// The host description.
	Description   pulumi.StringPtrOutput `pulumi:"description"`
	HostCatalogId pulumi.StringOutput    `pulumi:"hostCatalogId"`
	// The host name. Defaults to the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of host
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewBoundaryHostStatic registers a new resource with the given unique name, arguments, and options.
func NewBoundaryHostStatic(ctx *pulumi.Context,
	name string, args *BoundaryHostStaticArgs, opts ...pulumi.ResourceOption) (*BoundaryHostStatic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostCatalogId == nil {
		return nil, errors.New("invalid value for required argument 'HostCatalogId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource BoundaryHostStatic
	err := ctx.RegisterResource("boundary:index/boundaryHostStatic:BoundaryHostStatic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBoundaryHostStatic gets an existing BoundaryHostStatic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBoundaryHostStatic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BoundaryHostStaticState, opts ...pulumi.ResourceOption) (*BoundaryHostStatic, error) {
	var resource BoundaryHostStatic
	err := ctx.ReadResource("boundary:index/boundaryHostStatic:BoundaryHostStatic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BoundaryHostStatic resources.
type boundaryHostStaticState struct {
	// The static address of the host resource as `<IP>` (note: port assignment occurs in the target resource definition, do not add :port here) or a domain name.
	Address *string `pulumi:"address"`
	// The host description.
	Description   *string `pulumi:"description"`
	HostCatalogId *string `pulumi:"hostCatalogId"`
	// The host name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The type of host
	Type *string `pulumi:"type"`
}

type BoundaryHostStaticState struct {
	// The static address of the host resource as `<IP>` (note: port assignment occurs in the target resource definition, do not add :port here) or a domain name.
	Address pulumi.StringPtrInput
	// The host description.
	Description   pulumi.StringPtrInput
	HostCatalogId pulumi.StringPtrInput
	// The host name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The type of host
	Type pulumi.StringPtrInput
}

func (BoundaryHostStaticState) ElementType() reflect.Type {
	return reflect.TypeOf((*boundaryHostStaticState)(nil)).Elem()
}

type boundaryHostStaticArgs struct {
	// The static address of the host resource as `<IP>` (note: port assignment occurs in the target resource definition, do not add :port here) or a domain name.
	Address *string `pulumi:"address"`
	// The host description.
	Description   *string `pulumi:"description"`
	HostCatalogId string  `pulumi:"hostCatalogId"`
	// The host name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The type of host
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a BoundaryHostStatic resource.
type BoundaryHostStaticArgs struct {
	// The static address of the host resource as `<IP>` (note: port assignment occurs in the target resource definition, do not add :port here) or a domain name.
	Address pulumi.StringPtrInput
	// The host description.
	Description   pulumi.StringPtrInput
	HostCatalogId pulumi.StringInput
	// The host name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The type of host
	Type pulumi.StringPtrInput
}

func (BoundaryHostStaticArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*boundaryHostStaticArgs)(nil)).Elem()
}

type BoundaryHostStaticInput interface {
	pulumi.Input

	ToBoundaryHostStaticOutput() BoundaryHostStaticOutput
	ToBoundaryHostStaticOutputWithContext(ctx context.Context) BoundaryHostStaticOutput
}

func (*BoundaryHostStatic) ElementType() reflect.Type {
	return reflect.TypeOf((**BoundaryHostStatic)(nil)).Elem()
}

func (i *BoundaryHostStatic) ToBoundaryHostStaticOutput() BoundaryHostStaticOutput {
	return i.ToBoundaryHostStaticOutputWithContext(context.Background())
}

func (i *BoundaryHostStatic) ToBoundaryHostStaticOutputWithContext(ctx context.Context) BoundaryHostStaticOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoundaryHostStaticOutput)
}

// BoundaryHostStaticArrayInput is an input type that accepts BoundaryHostStaticArray and BoundaryHostStaticArrayOutput values.
// You can construct a concrete instance of `BoundaryHostStaticArrayInput` via:
//
//	BoundaryHostStaticArray{ BoundaryHostStaticArgs{...} }
type BoundaryHostStaticArrayInput interface {
	pulumi.Input

	ToBoundaryHostStaticArrayOutput() BoundaryHostStaticArrayOutput
	ToBoundaryHostStaticArrayOutputWithContext(context.Context) BoundaryHostStaticArrayOutput
}

type BoundaryHostStaticArray []BoundaryHostStaticInput

func (BoundaryHostStaticArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BoundaryHostStatic)(nil)).Elem()
}

func (i BoundaryHostStaticArray) ToBoundaryHostStaticArrayOutput() BoundaryHostStaticArrayOutput {
	return i.ToBoundaryHostStaticArrayOutputWithContext(context.Background())
}

func (i BoundaryHostStaticArray) ToBoundaryHostStaticArrayOutputWithContext(ctx context.Context) BoundaryHostStaticArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoundaryHostStaticArrayOutput)
}

// BoundaryHostStaticMapInput is an input type that accepts BoundaryHostStaticMap and BoundaryHostStaticMapOutput values.
// You can construct a concrete instance of `BoundaryHostStaticMapInput` via:
//
//	BoundaryHostStaticMap{ "key": BoundaryHostStaticArgs{...} }
type BoundaryHostStaticMapInput interface {
	pulumi.Input

	ToBoundaryHostStaticMapOutput() BoundaryHostStaticMapOutput
	ToBoundaryHostStaticMapOutputWithContext(context.Context) BoundaryHostStaticMapOutput
}

type BoundaryHostStaticMap map[string]BoundaryHostStaticInput

func (BoundaryHostStaticMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BoundaryHostStatic)(nil)).Elem()
}

func (i BoundaryHostStaticMap) ToBoundaryHostStaticMapOutput() BoundaryHostStaticMapOutput {
	return i.ToBoundaryHostStaticMapOutputWithContext(context.Background())
}

func (i BoundaryHostStaticMap) ToBoundaryHostStaticMapOutputWithContext(ctx context.Context) BoundaryHostStaticMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoundaryHostStaticMapOutput)
}

type BoundaryHostStaticOutput struct{ *pulumi.OutputState }

func (BoundaryHostStaticOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BoundaryHostStatic)(nil)).Elem()
}

func (o BoundaryHostStaticOutput) ToBoundaryHostStaticOutput() BoundaryHostStaticOutput {
	return o
}

func (o BoundaryHostStaticOutput) ToBoundaryHostStaticOutputWithContext(ctx context.Context) BoundaryHostStaticOutput {
	return o
}

// The static address of the host resource as `<IP>` (note: port assignment occurs in the target resource definition, do not add :port here) or a domain name.
func (o BoundaryHostStaticOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BoundaryHostStatic) pulumi.StringPtrOutput { return v.Address }).(pulumi.StringPtrOutput)
}

// The host description.
func (o BoundaryHostStaticOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BoundaryHostStatic) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o BoundaryHostStaticOutput) HostCatalogId() pulumi.StringOutput {
	return o.ApplyT(func(v *BoundaryHostStatic) pulumi.StringOutput { return v.HostCatalogId }).(pulumi.StringOutput)
}

// The host name. Defaults to the resource name.
func (o BoundaryHostStaticOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BoundaryHostStatic) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of host
func (o BoundaryHostStaticOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BoundaryHostStatic) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type BoundaryHostStaticArrayOutput struct{ *pulumi.OutputState }

func (BoundaryHostStaticArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BoundaryHostStatic)(nil)).Elem()
}

func (o BoundaryHostStaticArrayOutput) ToBoundaryHostStaticArrayOutput() BoundaryHostStaticArrayOutput {
	return o
}

func (o BoundaryHostStaticArrayOutput) ToBoundaryHostStaticArrayOutputWithContext(ctx context.Context) BoundaryHostStaticArrayOutput {
	return o
}

func (o BoundaryHostStaticArrayOutput) Index(i pulumi.IntInput) BoundaryHostStaticOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BoundaryHostStatic {
		return vs[0].([]*BoundaryHostStatic)[vs[1].(int)]
	}).(BoundaryHostStaticOutput)
}

type BoundaryHostStaticMapOutput struct{ *pulumi.OutputState }

func (BoundaryHostStaticMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BoundaryHostStatic)(nil)).Elem()
}

func (o BoundaryHostStaticMapOutput) ToBoundaryHostStaticMapOutput() BoundaryHostStaticMapOutput {
	return o
}

func (o BoundaryHostStaticMapOutput) ToBoundaryHostStaticMapOutputWithContext(ctx context.Context) BoundaryHostStaticMapOutput {
	return o
}

func (o BoundaryHostStaticMapOutput) MapIndex(k pulumi.StringInput) BoundaryHostStaticOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BoundaryHostStatic {
		return vs[0].(map[string]*BoundaryHostStatic)[vs[1].(string)]
	}).(BoundaryHostStaticOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BoundaryHostStaticInput)(nil)).Elem(), &BoundaryHostStatic{})
	pulumi.RegisterInputType(reflect.TypeOf((*BoundaryHostStaticArrayInput)(nil)).Elem(), BoundaryHostStaticArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BoundaryHostStaticMapInput)(nil)).Elem(), BoundaryHostStaticMap{})
	pulumi.RegisterOutputType(BoundaryHostStaticOutput{})
	pulumi.RegisterOutputType(BoundaryHostStaticArrayOutput{})
	pulumi.RegisterOutputType(BoundaryHostStaticMapOutput{})
}
