// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package boundary

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The SSH private key credential resource allows you to configure a credential using a username, private key and optional passphrase.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"io/ioutil"
//
//	"github.com/katasec/pulumi-boundary/sdk/go/boundary"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := ioutil.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			org, err := boundary.NewBoundaryScope(ctx, "org", &boundary.BoundaryScopeArgs{
//				Description:           pulumi.String("global scope"),
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
//			exampleBoundaryCredentialStoreStatic, err := boundary.NewBoundaryCredentialStoreStatic(ctx, "exampleBoundaryCredentialStoreStatic", &boundary.BoundaryCredentialStoreStaticArgs{
//				Description: pulumi.String("My first static credential store!"),
//				ScopeId:     project.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = boundary.NewBoundaryCredentialSshPrivateKey(ctx, "exampleBoundaryCredentialSshPrivateKey", &boundary.BoundaryCredentialSshPrivateKeyArgs{
//				Description:          pulumi.String("My first ssh private key credential!"),
//				CredentialStoreId:    exampleBoundaryCredentialStoreStatic.ID(),
//				Username:             pulumi.String("my-username"),
//				PrivateKey:           readFileOrPanic("~/.ssh/id_rsa"),
//				PrivateKeyPassphrase: pulumi.String("optional-passphrase"),
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
//	$ pulumi import boundary:index/boundaryCredentialSshPrivateKey:BoundaryCredentialSshPrivateKey example_ssh_private_key <my-id>
//
// ```
type BoundaryCredentialSshPrivateKey struct {
	pulumi.CustomResourceState

	// ID of the credential store this credential belongs to.
	CredentialStoreId pulumi.StringOutput `pulumi:"credentialStoreId"`
	// The description of the credential.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the credential. Defaults to the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The private key associated with the credential.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// The private key hmac.
	PrivateKeyHmac pulumi.StringOutput `pulumi:"privateKeyHmac"`
	// The passphrase of the private key associated with the credential.
	PrivateKeyPassphrase pulumi.StringPtrOutput `pulumi:"privateKeyPassphrase"`
	// The private key passphrase hmac.
	PrivateKeyPassphraseHmac pulumi.StringOutput `pulumi:"privateKeyPassphraseHmac"`
	// The username associated with the credential.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewBoundaryCredentialSshPrivateKey registers a new resource with the given unique name, arguments, and options.
func NewBoundaryCredentialSshPrivateKey(ctx *pulumi.Context,
	name string, args *BoundaryCredentialSshPrivateKeyArgs, opts ...pulumi.ResourceOption) (*BoundaryCredentialSshPrivateKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CredentialStoreId == nil {
		return nil, errors.New("invalid value for required argument 'CredentialStoreId'")
	}
	if args.PrivateKey == nil {
		return nil, errors.New("invalid value for required argument 'PrivateKey'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource BoundaryCredentialSshPrivateKey
	err := ctx.RegisterResource("boundary:index/boundaryCredentialSshPrivateKey:BoundaryCredentialSshPrivateKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBoundaryCredentialSshPrivateKey gets an existing BoundaryCredentialSshPrivateKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBoundaryCredentialSshPrivateKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BoundaryCredentialSshPrivateKeyState, opts ...pulumi.ResourceOption) (*BoundaryCredentialSshPrivateKey, error) {
	var resource BoundaryCredentialSshPrivateKey
	err := ctx.ReadResource("boundary:index/boundaryCredentialSshPrivateKey:BoundaryCredentialSshPrivateKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BoundaryCredentialSshPrivateKey resources.
type boundaryCredentialSshPrivateKeyState struct {
	// ID of the credential store this credential belongs to.
	CredentialStoreId *string `pulumi:"credentialStoreId"`
	// The description of the credential.
	Description *string `pulumi:"description"`
	// The name of the credential. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The private key associated with the credential.
	PrivateKey *string `pulumi:"privateKey"`
	// The private key hmac.
	PrivateKeyHmac *string `pulumi:"privateKeyHmac"`
	// The passphrase of the private key associated with the credential.
	PrivateKeyPassphrase *string `pulumi:"privateKeyPassphrase"`
	// The private key passphrase hmac.
	PrivateKeyPassphraseHmac *string `pulumi:"privateKeyPassphraseHmac"`
	// The username associated with the credential.
	Username *string `pulumi:"username"`
}

type BoundaryCredentialSshPrivateKeyState struct {
	// ID of the credential store this credential belongs to.
	CredentialStoreId pulumi.StringPtrInput
	// The description of the credential.
	Description pulumi.StringPtrInput
	// The name of the credential. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The private key associated with the credential.
	PrivateKey pulumi.StringPtrInput
	// The private key hmac.
	PrivateKeyHmac pulumi.StringPtrInput
	// The passphrase of the private key associated with the credential.
	PrivateKeyPassphrase pulumi.StringPtrInput
	// The private key passphrase hmac.
	PrivateKeyPassphraseHmac pulumi.StringPtrInput
	// The username associated with the credential.
	Username pulumi.StringPtrInput
}

func (BoundaryCredentialSshPrivateKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*boundaryCredentialSshPrivateKeyState)(nil)).Elem()
}

type boundaryCredentialSshPrivateKeyArgs struct {
	// ID of the credential store this credential belongs to.
	CredentialStoreId string `pulumi:"credentialStoreId"`
	// The description of the credential.
	Description *string `pulumi:"description"`
	// The name of the credential. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The private key associated with the credential.
	PrivateKey string `pulumi:"privateKey"`
	// The passphrase of the private key associated with the credential.
	PrivateKeyPassphrase *string `pulumi:"privateKeyPassphrase"`
	// The username associated with the credential.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a BoundaryCredentialSshPrivateKey resource.
type BoundaryCredentialSshPrivateKeyArgs struct {
	// ID of the credential store this credential belongs to.
	CredentialStoreId pulumi.StringInput
	// The description of the credential.
	Description pulumi.StringPtrInput
	// The name of the credential. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The private key associated with the credential.
	PrivateKey pulumi.StringInput
	// The passphrase of the private key associated with the credential.
	PrivateKeyPassphrase pulumi.StringPtrInput
	// The username associated with the credential.
	Username pulumi.StringInput
}

func (BoundaryCredentialSshPrivateKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*boundaryCredentialSshPrivateKeyArgs)(nil)).Elem()
}

type BoundaryCredentialSshPrivateKeyInput interface {
	pulumi.Input

	ToBoundaryCredentialSshPrivateKeyOutput() BoundaryCredentialSshPrivateKeyOutput
	ToBoundaryCredentialSshPrivateKeyOutputWithContext(ctx context.Context) BoundaryCredentialSshPrivateKeyOutput
}

func (*BoundaryCredentialSshPrivateKey) ElementType() reflect.Type {
	return reflect.TypeOf((**BoundaryCredentialSshPrivateKey)(nil)).Elem()
}

func (i *BoundaryCredentialSshPrivateKey) ToBoundaryCredentialSshPrivateKeyOutput() BoundaryCredentialSshPrivateKeyOutput {
	return i.ToBoundaryCredentialSshPrivateKeyOutputWithContext(context.Background())
}

func (i *BoundaryCredentialSshPrivateKey) ToBoundaryCredentialSshPrivateKeyOutputWithContext(ctx context.Context) BoundaryCredentialSshPrivateKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoundaryCredentialSshPrivateKeyOutput)
}

// BoundaryCredentialSshPrivateKeyArrayInput is an input type that accepts BoundaryCredentialSshPrivateKeyArray and BoundaryCredentialSshPrivateKeyArrayOutput values.
// You can construct a concrete instance of `BoundaryCredentialSshPrivateKeyArrayInput` via:
//
//	BoundaryCredentialSshPrivateKeyArray{ BoundaryCredentialSshPrivateKeyArgs{...} }
type BoundaryCredentialSshPrivateKeyArrayInput interface {
	pulumi.Input

	ToBoundaryCredentialSshPrivateKeyArrayOutput() BoundaryCredentialSshPrivateKeyArrayOutput
	ToBoundaryCredentialSshPrivateKeyArrayOutputWithContext(context.Context) BoundaryCredentialSshPrivateKeyArrayOutput
}

type BoundaryCredentialSshPrivateKeyArray []BoundaryCredentialSshPrivateKeyInput

func (BoundaryCredentialSshPrivateKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BoundaryCredentialSshPrivateKey)(nil)).Elem()
}

func (i BoundaryCredentialSshPrivateKeyArray) ToBoundaryCredentialSshPrivateKeyArrayOutput() BoundaryCredentialSshPrivateKeyArrayOutput {
	return i.ToBoundaryCredentialSshPrivateKeyArrayOutputWithContext(context.Background())
}

func (i BoundaryCredentialSshPrivateKeyArray) ToBoundaryCredentialSshPrivateKeyArrayOutputWithContext(ctx context.Context) BoundaryCredentialSshPrivateKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoundaryCredentialSshPrivateKeyArrayOutput)
}

// BoundaryCredentialSshPrivateKeyMapInput is an input type that accepts BoundaryCredentialSshPrivateKeyMap and BoundaryCredentialSshPrivateKeyMapOutput values.
// You can construct a concrete instance of `BoundaryCredentialSshPrivateKeyMapInput` via:
//
//	BoundaryCredentialSshPrivateKeyMap{ "key": BoundaryCredentialSshPrivateKeyArgs{...} }
type BoundaryCredentialSshPrivateKeyMapInput interface {
	pulumi.Input

	ToBoundaryCredentialSshPrivateKeyMapOutput() BoundaryCredentialSshPrivateKeyMapOutput
	ToBoundaryCredentialSshPrivateKeyMapOutputWithContext(context.Context) BoundaryCredentialSshPrivateKeyMapOutput
}

type BoundaryCredentialSshPrivateKeyMap map[string]BoundaryCredentialSshPrivateKeyInput

func (BoundaryCredentialSshPrivateKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BoundaryCredentialSshPrivateKey)(nil)).Elem()
}

func (i BoundaryCredentialSshPrivateKeyMap) ToBoundaryCredentialSshPrivateKeyMapOutput() BoundaryCredentialSshPrivateKeyMapOutput {
	return i.ToBoundaryCredentialSshPrivateKeyMapOutputWithContext(context.Background())
}

func (i BoundaryCredentialSshPrivateKeyMap) ToBoundaryCredentialSshPrivateKeyMapOutputWithContext(ctx context.Context) BoundaryCredentialSshPrivateKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoundaryCredentialSshPrivateKeyMapOutput)
}

type BoundaryCredentialSshPrivateKeyOutput struct{ *pulumi.OutputState }

func (BoundaryCredentialSshPrivateKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BoundaryCredentialSshPrivateKey)(nil)).Elem()
}

func (o BoundaryCredentialSshPrivateKeyOutput) ToBoundaryCredentialSshPrivateKeyOutput() BoundaryCredentialSshPrivateKeyOutput {
	return o
}

func (o BoundaryCredentialSshPrivateKeyOutput) ToBoundaryCredentialSshPrivateKeyOutputWithContext(ctx context.Context) BoundaryCredentialSshPrivateKeyOutput {
	return o
}

// ID of the credential store this credential belongs to.
func (o BoundaryCredentialSshPrivateKeyOutput) CredentialStoreId() pulumi.StringOutput {
	return o.ApplyT(func(v *BoundaryCredentialSshPrivateKey) pulumi.StringOutput { return v.CredentialStoreId }).(pulumi.StringOutput)
}

// The description of the credential.
func (o BoundaryCredentialSshPrivateKeyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BoundaryCredentialSshPrivateKey) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the credential. Defaults to the resource name.
func (o BoundaryCredentialSshPrivateKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BoundaryCredentialSshPrivateKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The private key associated with the credential.
func (o BoundaryCredentialSshPrivateKeyOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *BoundaryCredentialSshPrivateKey) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

// The private key hmac.
func (o BoundaryCredentialSshPrivateKeyOutput) PrivateKeyHmac() pulumi.StringOutput {
	return o.ApplyT(func(v *BoundaryCredentialSshPrivateKey) pulumi.StringOutput { return v.PrivateKeyHmac }).(pulumi.StringOutput)
}

// The passphrase of the private key associated with the credential.
func (o BoundaryCredentialSshPrivateKeyOutput) PrivateKeyPassphrase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BoundaryCredentialSshPrivateKey) pulumi.StringPtrOutput { return v.PrivateKeyPassphrase }).(pulumi.StringPtrOutput)
}

// The private key passphrase hmac.
func (o BoundaryCredentialSshPrivateKeyOutput) PrivateKeyPassphraseHmac() pulumi.StringOutput {
	return o.ApplyT(func(v *BoundaryCredentialSshPrivateKey) pulumi.StringOutput { return v.PrivateKeyPassphraseHmac }).(pulumi.StringOutput)
}

// The username associated with the credential.
func (o BoundaryCredentialSshPrivateKeyOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *BoundaryCredentialSshPrivateKey) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type BoundaryCredentialSshPrivateKeyArrayOutput struct{ *pulumi.OutputState }

func (BoundaryCredentialSshPrivateKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BoundaryCredentialSshPrivateKey)(nil)).Elem()
}

func (o BoundaryCredentialSshPrivateKeyArrayOutput) ToBoundaryCredentialSshPrivateKeyArrayOutput() BoundaryCredentialSshPrivateKeyArrayOutput {
	return o
}

func (o BoundaryCredentialSshPrivateKeyArrayOutput) ToBoundaryCredentialSshPrivateKeyArrayOutputWithContext(ctx context.Context) BoundaryCredentialSshPrivateKeyArrayOutput {
	return o
}

func (o BoundaryCredentialSshPrivateKeyArrayOutput) Index(i pulumi.IntInput) BoundaryCredentialSshPrivateKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BoundaryCredentialSshPrivateKey {
		return vs[0].([]*BoundaryCredentialSshPrivateKey)[vs[1].(int)]
	}).(BoundaryCredentialSshPrivateKeyOutput)
}

type BoundaryCredentialSshPrivateKeyMapOutput struct{ *pulumi.OutputState }

func (BoundaryCredentialSshPrivateKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BoundaryCredentialSshPrivateKey)(nil)).Elem()
}

func (o BoundaryCredentialSshPrivateKeyMapOutput) ToBoundaryCredentialSshPrivateKeyMapOutput() BoundaryCredentialSshPrivateKeyMapOutput {
	return o
}

func (o BoundaryCredentialSshPrivateKeyMapOutput) ToBoundaryCredentialSshPrivateKeyMapOutputWithContext(ctx context.Context) BoundaryCredentialSshPrivateKeyMapOutput {
	return o
}

func (o BoundaryCredentialSshPrivateKeyMapOutput) MapIndex(k pulumi.StringInput) BoundaryCredentialSshPrivateKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BoundaryCredentialSshPrivateKey {
		return vs[0].(map[string]*BoundaryCredentialSshPrivateKey)[vs[1].(string)]
	}).(BoundaryCredentialSshPrivateKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BoundaryCredentialSshPrivateKeyInput)(nil)).Elem(), &BoundaryCredentialSshPrivateKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*BoundaryCredentialSshPrivateKeyArrayInput)(nil)).Elem(), BoundaryCredentialSshPrivateKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BoundaryCredentialSshPrivateKeyMapInput)(nil)).Elem(), BoundaryCredentialSshPrivateKeyMap{})
	pulumi.RegisterOutputType(BoundaryCredentialSshPrivateKeyOutput{})
	pulumi.RegisterOutputType(BoundaryCredentialSshPrivateKeyArrayOutput{})
	pulumi.RegisterOutputType(BoundaryCredentialSshPrivateKeyMapOutput{})
}
