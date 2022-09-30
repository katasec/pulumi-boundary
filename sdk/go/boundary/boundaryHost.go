// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package boundary

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Deprecated: use `BoundaryHostStatic` instead.
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
//			project, err := boundary.NewBoundaryScope(ctx, "project", &boundary.BoundaryScopeArgs{
//				Description:         pulumi.String("My first scope!"),
//				ScopeId:             org.ID(),
//				AutoCreateAdminRole: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			static, err := boundary.NewBoundaryHostCatalog(ctx, "static", &boundary.BoundaryHostCatalogArgs{
//				Description: pulumi.String("My first host catalog!"),
//				Type:        pulumi.String("static"),
//				ScopeId:     project.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = boundary.NewBoundaryHost(ctx, "example", &boundary.BoundaryHostArgs{
//				Type:          pulumi.String("static"),
//				Description:   pulumi.String("My first host!"),
//				Address:       pulumi.String("10.0.0.1"),
//				HostCatalogId: static.ID(),
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
//	$ pulumi import boundary:index/boundaryHost:BoundaryHost foo <my-id>
//
// ```
type BoundaryHost struct {
	pulumi.CustomResourceState

	// The static address of the host resource as `<IP>` (note: port assignment occurs in the target resource definition, do not add :port here) or a domain name.
	Address pulumi.StringPtrOutput `pulumi:"address"`
	// The host description.
	Description   pulumi.StringPtrOutput `pulumi:"description"`
	HostCatalogId pulumi.StringOutput    `pulumi:"hostCatalogId"`
	// The host name. Defaults to the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of host
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBoundaryHost registers a new resource with the given unique name, arguments, and options.
func NewBoundaryHost(ctx *pulumi.Context,
	name string, args *BoundaryHostArgs, opts ...pulumi.ResourceOption) (*BoundaryHost, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostCatalogId == nil {
		return nil, errors.New("invalid value for required argument 'HostCatalogId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource BoundaryHost
	err := ctx.RegisterResource("boundary:index/boundaryHost:BoundaryHost", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBoundaryHost gets an existing BoundaryHost resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBoundaryHost(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BoundaryHostState, opts ...pulumi.ResourceOption) (*BoundaryHost, error) {
	var resource BoundaryHost
	err := ctx.ReadResource("boundary:index/boundaryHost:BoundaryHost", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BoundaryHost resources.
type boundaryHostState struct {
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

type BoundaryHostState struct {
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

func (BoundaryHostState) ElementType() reflect.Type {
	return reflect.TypeOf((*boundaryHostState)(nil)).Elem()
}

type boundaryHostArgs struct {
	// The static address of the host resource as `<IP>` (note: port assignment occurs in the target resource definition, do not add :port here) or a domain name.
	Address *string `pulumi:"address"`
	// The host description.
	Description   *string `pulumi:"description"`
	HostCatalogId string  `pulumi:"hostCatalogId"`
	// The host name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The type of host
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a BoundaryHost resource.
type BoundaryHostArgs struct {
	// The static address of the host resource as `<IP>` (note: port assignment occurs in the target resource definition, do not add :port here) or a domain name.
	Address pulumi.StringPtrInput
	// The host description.
	Description   pulumi.StringPtrInput
	HostCatalogId pulumi.StringInput
	// The host name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The type of host
	Type pulumi.StringInput
}

func (BoundaryHostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*boundaryHostArgs)(nil)).Elem()
}

type BoundaryHostInput interface {
	pulumi.Input

	ToBoundaryHostOutput() BoundaryHostOutput
	ToBoundaryHostOutputWithContext(ctx context.Context) BoundaryHostOutput
}

func (*BoundaryHost) ElementType() reflect.Type {
	return reflect.TypeOf((**BoundaryHost)(nil)).Elem()
}

func (i *BoundaryHost) ToBoundaryHostOutput() BoundaryHostOutput {
	return i.ToBoundaryHostOutputWithContext(context.Background())
}

func (i *BoundaryHost) ToBoundaryHostOutputWithContext(ctx context.Context) BoundaryHostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoundaryHostOutput)
}

// BoundaryHostArrayInput is an input type that accepts BoundaryHostArray and BoundaryHostArrayOutput values.
// You can construct a concrete instance of `BoundaryHostArrayInput` via:
//
//	BoundaryHostArray{ BoundaryHostArgs{...} }
type BoundaryHostArrayInput interface {
	pulumi.Input

	ToBoundaryHostArrayOutput() BoundaryHostArrayOutput
	ToBoundaryHostArrayOutputWithContext(context.Context) BoundaryHostArrayOutput
}

type BoundaryHostArray []BoundaryHostInput

func (BoundaryHostArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BoundaryHost)(nil)).Elem()
}

func (i BoundaryHostArray) ToBoundaryHostArrayOutput() BoundaryHostArrayOutput {
	return i.ToBoundaryHostArrayOutputWithContext(context.Background())
}

func (i BoundaryHostArray) ToBoundaryHostArrayOutputWithContext(ctx context.Context) BoundaryHostArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoundaryHostArrayOutput)
}

// BoundaryHostMapInput is an input type that accepts BoundaryHostMap and BoundaryHostMapOutput values.
// You can construct a concrete instance of `BoundaryHostMapInput` via:
//
//	BoundaryHostMap{ "key": BoundaryHostArgs{...} }
type BoundaryHostMapInput interface {
	pulumi.Input

	ToBoundaryHostMapOutput() BoundaryHostMapOutput
	ToBoundaryHostMapOutputWithContext(context.Context) BoundaryHostMapOutput
}

type BoundaryHostMap map[string]BoundaryHostInput

func (BoundaryHostMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BoundaryHost)(nil)).Elem()
}

func (i BoundaryHostMap) ToBoundaryHostMapOutput() BoundaryHostMapOutput {
	return i.ToBoundaryHostMapOutputWithContext(context.Background())
}

func (i BoundaryHostMap) ToBoundaryHostMapOutputWithContext(ctx context.Context) BoundaryHostMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoundaryHostMapOutput)
}

type BoundaryHostOutput struct{ *pulumi.OutputState }

func (BoundaryHostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BoundaryHost)(nil)).Elem()
}

func (o BoundaryHostOutput) ToBoundaryHostOutput() BoundaryHostOutput {
	return o
}

func (o BoundaryHostOutput) ToBoundaryHostOutputWithContext(ctx context.Context) BoundaryHostOutput {
	return o
}

// The static address of the host resource as `<IP>` (note: port assignment occurs in the target resource definition, do not add :port here) or a domain name.
func (o BoundaryHostOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BoundaryHost) pulumi.StringPtrOutput { return v.Address }).(pulumi.StringPtrOutput)
}

// The host description.
func (o BoundaryHostOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BoundaryHost) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o BoundaryHostOutput) HostCatalogId() pulumi.StringOutput {
	return o.ApplyT(func(v *BoundaryHost) pulumi.StringOutput { return v.HostCatalogId }).(pulumi.StringOutput)
}

// The host name. Defaults to the resource name.
func (o BoundaryHostOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BoundaryHost) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of host
func (o BoundaryHostOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BoundaryHost) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type BoundaryHostArrayOutput struct{ *pulumi.OutputState }

func (BoundaryHostArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BoundaryHost)(nil)).Elem()
}

func (o BoundaryHostArrayOutput) ToBoundaryHostArrayOutput() BoundaryHostArrayOutput {
	return o
}

func (o BoundaryHostArrayOutput) ToBoundaryHostArrayOutputWithContext(ctx context.Context) BoundaryHostArrayOutput {
	return o
}

func (o BoundaryHostArrayOutput) Index(i pulumi.IntInput) BoundaryHostOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BoundaryHost {
		return vs[0].([]*BoundaryHost)[vs[1].(int)]
	}).(BoundaryHostOutput)
}

type BoundaryHostMapOutput struct{ *pulumi.OutputState }

func (BoundaryHostMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BoundaryHost)(nil)).Elem()
}

func (o BoundaryHostMapOutput) ToBoundaryHostMapOutput() BoundaryHostMapOutput {
	return o
}

func (o BoundaryHostMapOutput) ToBoundaryHostMapOutputWithContext(ctx context.Context) BoundaryHostMapOutput {
	return o
}

func (o BoundaryHostMapOutput) MapIndex(k pulumi.StringInput) BoundaryHostOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BoundaryHost {
		return vs[0].(map[string]*BoundaryHost)[vs[1].(string)]
	}).(BoundaryHostOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BoundaryHostInput)(nil)).Elem(), &BoundaryHost{})
	pulumi.RegisterInputType(reflect.TypeOf((*BoundaryHostArrayInput)(nil)).Elem(), BoundaryHostArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BoundaryHostMapInput)(nil)).Elem(), BoundaryHostMap{})
	pulumi.RegisterOutputType(BoundaryHostOutput{})
	pulumi.RegisterOutputType(BoundaryHostArrayOutput{})
	pulumi.RegisterOutputType(BoundaryHostMapOutput{})
}
