// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package boundary

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The hostSetPlugin resource allows you to configure a Boundary host set. Host sets are always part of a host catalog, so a host catalog resource should be used inline or you should have the host catalog ID in hand to successfully configure a host set.
type BoundaryHostSetPlugin struct {
	pulumi.CustomResourceState

	// The attributes for the host set. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host set.
	AttributesJson pulumi.StringPtrOutput `pulumi:"attributesJson"`
	// The host set description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The catalog for the host set.
	HostCatalogId pulumi.StringOutput `pulumi:"hostCatalogId"`
	// The host set name. Defaults to the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ordered list of preferred endpoints.
	PreferredEndpoints pulumi.StringArrayOutput `pulumi:"preferredEndpoints"`
	// The value to set for the sync interval seconds.
	SyncIntervalSeconds pulumi.IntPtrOutput `pulumi:"syncIntervalSeconds"`
	// The type of host set
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewBoundaryHostSetPlugin registers a new resource with the given unique name, arguments, and options.
func NewBoundaryHostSetPlugin(ctx *pulumi.Context,
	name string, args *BoundaryHostSetPluginArgs, opts ...pulumi.ResourceOption) (*BoundaryHostSetPlugin, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostCatalogId == nil {
		return nil, errors.New("invalid value for required argument 'HostCatalogId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource BoundaryHostSetPlugin
	err := ctx.RegisterResource("boundary:index/boundaryHostSetPlugin:BoundaryHostSetPlugin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBoundaryHostSetPlugin gets an existing BoundaryHostSetPlugin resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBoundaryHostSetPlugin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BoundaryHostSetPluginState, opts ...pulumi.ResourceOption) (*BoundaryHostSetPlugin, error) {
	var resource BoundaryHostSetPlugin
	err := ctx.ReadResource("boundary:index/boundaryHostSetPlugin:BoundaryHostSetPlugin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BoundaryHostSetPlugin resources.
type boundaryHostSetPluginState struct {
	// The attributes for the host set. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host set.
	AttributesJson *string `pulumi:"attributesJson"`
	// The host set description.
	Description *string `pulumi:"description"`
	// The catalog for the host set.
	HostCatalogId *string `pulumi:"hostCatalogId"`
	// The host set name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The ordered list of preferred endpoints.
	PreferredEndpoints []string `pulumi:"preferredEndpoints"`
	// The value to set for the sync interval seconds.
	SyncIntervalSeconds *int `pulumi:"syncIntervalSeconds"`
	// The type of host set
	Type *string `pulumi:"type"`
}

type BoundaryHostSetPluginState struct {
	// The attributes for the host set. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host set.
	AttributesJson pulumi.StringPtrInput
	// The host set description.
	Description pulumi.StringPtrInput
	// The catalog for the host set.
	HostCatalogId pulumi.StringPtrInput
	// The host set name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The ordered list of preferred endpoints.
	PreferredEndpoints pulumi.StringArrayInput
	// The value to set for the sync interval seconds.
	SyncIntervalSeconds pulumi.IntPtrInput
	// The type of host set
	Type pulumi.StringPtrInput
}

func (BoundaryHostSetPluginState) ElementType() reflect.Type {
	return reflect.TypeOf((*boundaryHostSetPluginState)(nil)).Elem()
}

type boundaryHostSetPluginArgs struct {
	// The attributes for the host set. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host set.
	AttributesJson *string `pulumi:"attributesJson"`
	// The host set description.
	Description *string `pulumi:"description"`
	// The catalog for the host set.
	HostCatalogId string `pulumi:"hostCatalogId"`
	// The host set name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The ordered list of preferred endpoints.
	PreferredEndpoints []string `pulumi:"preferredEndpoints"`
	// The value to set for the sync interval seconds.
	SyncIntervalSeconds *int `pulumi:"syncIntervalSeconds"`
	// The type of host set
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a BoundaryHostSetPlugin resource.
type BoundaryHostSetPluginArgs struct {
	// The attributes for the host set. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host set.
	AttributesJson pulumi.StringPtrInput
	// The host set description.
	Description pulumi.StringPtrInput
	// The catalog for the host set.
	HostCatalogId pulumi.StringInput
	// The host set name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The ordered list of preferred endpoints.
	PreferredEndpoints pulumi.StringArrayInput
	// The value to set for the sync interval seconds.
	SyncIntervalSeconds pulumi.IntPtrInput
	// The type of host set
	Type pulumi.StringPtrInput
}

func (BoundaryHostSetPluginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*boundaryHostSetPluginArgs)(nil)).Elem()
}

type BoundaryHostSetPluginInput interface {
	pulumi.Input

	ToBoundaryHostSetPluginOutput() BoundaryHostSetPluginOutput
	ToBoundaryHostSetPluginOutputWithContext(ctx context.Context) BoundaryHostSetPluginOutput
}

func (*BoundaryHostSetPlugin) ElementType() reflect.Type {
	return reflect.TypeOf((**BoundaryHostSetPlugin)(nil)).Elem()
}

func (i *BoundaryHostSetPlugin) ToBoundaryHostSetPluginOutput() BoundaryHostSetPluginOutput {
	return i.ToBoundaryHostSetPluginOutputWithContext(context.Background())
}

func (i *BoundaryHostSetPlugin) ToBoundaryHostSetPluginOutputWithContext(ctx context.Context) BoundaryHostSetPluginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoundaryHostSetPluginOutput)
}

// BoundaryHostSetPluginArrayInput is an input type that accepts BoundaryHostSetPluginArray and BoundaryHostSetPluginArrayOutput values.
// You can construct a concrete instance of `BoundaryHostSetPluginArrayInput` via:
//
//	BoundaryHostSetPluginArray{ BoundaryHostSetPluginArgs{...} }
type BoundaryHostSetPluginArrayInput interface {
	pulumi.Input

	ToBoundaryHostSetPluginArrayOutput() BoundaryHostSetPluginArrayOutput
	ToBoundaryHostSetPluginArrayOutputWithContext(context.Context) BoundaryHostSetPluginArrayOutput
}

type BoundaryHostSetPluginArray []BoundaryHostSetPluginInput

func (BoundaryHostSetPluginArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BoundaryHostSetPlugin)(nil)).Elem()
}

func (i BoundaryHostSetPluginArray) ToBoundaryHostSetPluginArrayOutput() BoundaryHostSetPluginArrayOutput {
	return i.ToBoundaryHostSetPluginArrayOutputWithContext(context.Background())
}

func (i BoundaryHostSetPluginArray) ToBoundaryHostSetPluginArrayOutputWithContext(ctx context.Context) BoundaryHostSetPluginArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoundaryHostSetPluginArrayOutput)
}

// BoundaryHostSetPluginMapInput is an input type that accepts BoundaryHostSetPluginMap and BoundaryHostSetPluginMapOutput values.
// You can construct a concrete instance of `BoundaryHostSetPluginMapInput` via:
//
//	BoundaryHostSetPluginMap{ "key": BoundaryHostSetPluginArgs{...} }
type BoundaryHostSetPluginMapInput interface {
	pulumi.Input

	ToBoundaryHostSetPluginMapOutput() BoundaryHostSetPluginMapOutput
	ToBoundaryHostSetPluginMapOutputWithContext(context.Context) BoundaryHostSetPluginMapOutput
}

type BoundaryHostSetPluginMap map[string]BoundaryHostSetPluginInput

func (BoundaryHostSetPluginMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BoundaryHostSetPlugin)(nil)).Elem()
}

func (i BoundaryHostSetPluginMap) ToBoundaryHostSetPluginMapOutput() BoundaryHostSetPluginMapOutput {
	return i.ToBoundaryHostSetPluginMapOutputWithContext(context.Background())
}

func (i BoundaryHostSetPluginMap) ToBoundaryHostSetPluginMapOutputWithContext(ctx context.Context) BoundaryHostSetPluginMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoundaryHostSetPluginMapOutput)
}

type BoundaryHostSetPluginOutput struct{ *pulumi.OutputState }

func (BoundaryHostSetPluginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BoundaryHostSetPlugin)(nil)).Elem()
}

func (o BoundaryHostSetPluginOutput) ToBoundaryHostSetPluginOutput() BoundaryHostSetPluginOutput {
	return o
}

func (o BoundaryHostSetPluginOutput) ToBoundaryHostSetPluginOutputWithContext(ctx context.Context) BoundaryHostSetPluginOutput {
	return o
}

// The attributes for the host set. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host set.
func (o BoundaryHostSetPluginOutput) AttributesJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BoundaryHostSetPlugin) pulumi.StringPtrOutput { return v.AttributesJson }).(pulumi.StringPtrOutput)
}

// The host set description.
func (o BoundaryHostSetPluginOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BoundaryHostSetPlugin) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The catalog for the host set.
func (o BoundaryHostSetPluginOutput) HostCatalogId() pulumi.StringOutput {
	return o.ApplyT(func(v *BoundaryHostSetPlugin) pulumi.StringOutput { return v.HostCatalogId }).(pulumi.StringOutput)
}

// The host set name. Defaults to the resource name.
func (o BoundaryHostSetPluginOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BoundaryHostSetPlugin) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ordered list of preferred endpoints.
func (o BoundaryHostSetPluginOutput) PreferredEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BoundaryHostSetPlugin) pulumi.StringArrayOutput { return v.PreferredEndpoints }).(pulumi.StringArrayOutput)
}

// The value to set for the sync interval seconds.
func (o BoundaryHostSetPluginOutput) SyncIntervalSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BoundaryHostSetPlugin) pulumi.IntPtrOutput { return v.SyncIntervalSeconds }).(pulumi.IntPtrOutput)
}

// The type of host set
func (o BoundaryHostSetPluginOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BoundaryHostSetPlugin) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type BoundaryHostSetPluginArrayOutput struct{ *pulumi.OutputState }

func (BoundaryHostSetPluginArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BoundaryHostSetPlugin)(nil)).Elem()
}

func (o BoundaryHostSetPluginArrayOutput) ToBoundaryHostSetPluginArrayOutput() BoundaryHostSetPluginArrayOutput {
	return o
}

func (o BoundaryHostSetPluginArrayOutput) ToBoundaryHostSetPluginArrayOutputWithContext(ctx context.Context) BoundaryHostSetPluginArrayOutput {
	return o
}

func (o BoundaryHostSetPluginArrayOutput) Index(i pulumi.IntInput) BoundaryHostSetPluginOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BoundaryHostSetPlugin {
		return vs[0].([]*BoundaryHostSetPlugin)[vs[1].(int)]
	}).(BoundaryHostSetPluginOutput)
}

type BoundaryHostSetPluginMapOutput struct{ *pulumi.OutputState }

func (BoundaryHostSetPluginMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BoundaryHostSetPlugin)(nil)).Elem()
}

func (o BoundaryHostSetPluginMapOutput) ToBoundaryHostSetPluginMapOutput() BoundaryHostSetPluginMapOutput {
	return o
}

func (o BoundaryHostSetPluginMapOutput) ToBoundaryHostSetPluginMapOutputWithContext(ctx context.Context) BoundaryHostSetPluginMapOutput {
	return o
}

func (o BoundaryHostSetPluginMapOutput) MapIndex(k pulumi.StringInput) BoundaryHostSetPluginOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BoundaryHostSetPlugin {
		return vs[0].(map[string]*BoundaryHostSetPlugin)[vs[1].(string)]
	}).(BoundaryHostSetPluginOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BoundaryHostSetPluginInput)(nil)).Elem(), &BoundaryHostSetPlugin{})
	pulumi.RegisterInputType(reflect.TypeOf((*BoundaryHostSetPluginArrayInput)(nil)).Elem(), BoundaryHostSetPluginArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BoundaryHostSetPluginMapInput)(nil)).Elem(), BoundaryHostSetPluginMap{})
	pulumi.RegisterOutputType(BoundaryHostSetPluginOutput{})
	pulumi.RegisterOutputType(BoundaryHostSetPluginArrayOutput{})
	pulumi.RegisterOutputType(BoundaryHostSetPluginMapOutput{})
}
