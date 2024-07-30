// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package konnect

import (
	"context"
	"reflect"

	"errors"
	"github.com/bunchc/pulumi-tf-provider-konnect/sdk/go/konnect/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// GatewayKey Resource
type GatewayKey struct {
	pulumi.CustomResourceState

	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId pulumi.StringOutput `pulumi:"controlPlaneId"`
	// Unix epoch when the resource was created.
	CreatedAt pulumi.IntOutput `pulumi:"createdAt"`
	// A JSON Web Key represented as a string.
	Jwk pulumi.StringOutput `pulumi:"jwk"`
	// A unique identifier for a key.
	Kid pulumi.StringOutput `pulumi:"kid"`
	// The name to associate with the given keys.
	Name pulumi.StringOutput `pulumi:"name"`
	// A keypair in PEM format.
	Pem GatewayKeyPemOutput `pulumi:"pem"`
	// The id (an UUID) of the key-set with which to associate the key.
	Set GatewayKeySetTypeOutput `pulumi:"set"`
	// An optional set of strings associated with the Key for grouping and filtering.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Unix epoch when the resource was last updated.
	UpdatedAt pulumi.IntOutput `pulumi:"updatedAt"`
}

// NewGatewayKey registers a new resource with the given unique name, arguments, and options.
func NewGatewayKey(ctx *pulumi.Context,
	name string, args *GatewayKeyArgs, opts ...pulumi.ResourceOption) (*GatewayKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ControlPlaneId == nil {
		return nil, errors.New("invalid value for required argument 'ControlPlaneId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GatewayKey
	err := ctx.RegisterResource("konnect:index/gatewayKey:GatewayKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayKey gets an existing GatewayKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayKeyState, opts ...pulumi.ResourceOption) (*GatewayKey, error) {
	var resource GatewayKey
	err := ctx.ReadResource("konnect:index/gatewayKey:GatewayKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayKey resources.
type gatewayKeyState struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId *string `pulumi:"controlPlaneId"`
	// Unix epoch when the resource was created.
	CreatedAt *int `pulumi:"createdAt"`
	// A JSON Web Key represented as a string.
	Jwk *string `pulumi:"jwk"`
	// A unique identifier for a key.
	Kid *string `pulumi:"kid"`
	// The name to associate with the given keys.
	Name *string `pulumi:"name"`
	// A keypair in PEM format.
	Pem *GatewayKeyPem `pulumi:"pem"`
	// The id (an UUID) of the key-set with which to associate the key.
	Set *GatewayKeySetType `pulumi:"set"`
	// An optional set of strings associated with the Key for grouping and filtering.
	Tags []string `pulumi:"tags"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int `pulumi:"updatedAt"`
}

type GatewayKeyState struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId pulumi.StringPtrInput
	// Unix epoch when the resource was created.
	CreatedAt pulumi.IntPtrInput
	// A JSON Web Key represented as a string.
	Jwk pulumi.StringPtrInput
	// A unique identifier for a key.
	Kid pulumi.StringPtrInput
	// The name to associate with the given keys.
	Name pulumi.StringPtrInput
	// A keypair in PEM format.
	Pem GatewayKeyPemPtrInput
	// The id (an UUID) of the key-set with which to associate the key.
	Set GatewayKeySetTypePtrInput
	// An optional set of strings associated with the Key for grouping and filtering.
	Tags pulumi.StringArrayInput
	// Unix epoch when the resource was last updated.
	UpdatedAt pulumi.IntPtrInput
}

func (GatewayKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayKeyState)(nil)).Elem()
}

type gatewayKeyArgs struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId string `pulumi:"controlPlaneId"`
	// A JSON Web Key represented as a string.
	Jwk *string `pulumi:"jwk"`
	// A unique identifier for a key.
	Kid *string `pulumi:"kid"`
	// The name to associate with the given keys.
	Name *string `pulumi:"name"`
	// A keypair in PEM format.
	Pem *GatewayKeyPem `pulumi:"pem"`
	// The id (an UUID) of the key-set with which to associate the key.
	Set *GatewayKeySetType `pulumi:"set"`
	// An optional set of strings associated with the Key for grouping and filtering.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a GatewayKey resource.
type GatewayKeyArgs struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId pulumi.StringInput
	// A JSON Web Key represented as a string.
	Jwk pulumi.StringPtrInput
	// A unique identifier for a key.
	Kid pulumi.StringPtrInput
	// The name to associate with the given keys.
	Name pulumi.StringPtrInput
	// A keypair in PEM format.
	Pem GatewayKeyPemPtrInput
	// The id (an UUID) of the key-set with which to associate the key.
	Set GatewayKeySetTypePtrInput
	// An optional set of strings associated with the Key for grouping and filtering.
	Tags pulumi.StringArrayInput
}

func (GatewayKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayKeyArgs)(nil)).Elem()
}

type GatewayKeyInput interface {
	pulumi.Input

	ToGatewayKeyOutput() GatewayKeyOutput
	ToGatewayKeyOutputWithContext(ctx context.Context) GatewayKeyOutput
}

func (*GatewayKey) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayKey)(nil)).Elem()
}

func (i *GatewayKey) ToGatewayKeyOutput() GatewayKeyOutput {
	return i.ToGatewayKeyOutputWithContext(context.Background())
}

func (i *GatewayKey) ToGatewayKeyOutputWithContext(ctx context.Context) GatewayKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayKeyOutput)
}

// GatewayKeyArrayInput is an input type that accepts GatewayKeyArray and GatewayKeyArrayOutput values.
// You can construct a concrete instance of `GatewayKeyArrayInput` via:
//
//	GatewayKeyArray{ GatewayKeyArgs{...} }
type GatewayKeyArrayInput interface {
	pulumi.Input

	ToGatewayKeyArrayOutput() GatewayKeyArrayOutput
	ToGatewayKeyArrayOutputWithContext(context.Context) GatewayKeyArrayOutput
}

type GatewayKeyArray []GatewayKeyInput

func (GatewayKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayKey)(nil)).Elem()
}

func (i GatewayKeyArray) ToGatewayKeyArrayOutput() GatewayKeyArrayOutput {
	return i.ToGatewayKeyArrayOutputWithContext(context.Background())
}

func (i GatewayKeyArray) ToGatewayKeyArrayOutputWithContext(ctx context.Context) GatewayKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayKeyArrayOutput)
}

// GatewayKeyMapInput is an input type that accepts GatewayKeyMap and GatewayKeyMapOutput values.
// You can construct a concrete instance of `GatewayKeyMapInput` via:
//
//	GatewayKeyMap{ "key": GatewayKeyArgs{...} }
type GatewayKeyMapInput interface {
	pulumi.Input

	ToGatewayKeyMapOutput() GatewayKeyMapOutput
	ToGatewayKeyMapOutputWithContext(context.Context) GatewayKeyMapOutput
}

type GatewayKeyMap map[string]GatewayKeyInput

func (GatewayKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayKey)(nil)).Elem()
}

func (i GatewayKeyMap) ToGatewayKeyMapOutput() GatewayKeyMapOutput {
	return i.ToGatewayKeyMapOutputWithContext(context.Background())
}

func (i GatewayKeyMap) ToGatewayKeyMapOutputWithContext(ctx context.Context) GatewayKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayKeyMapOutput)
}

type GatewayKeyOutput struct{ *pulumi.OutputState }

func (GatewayKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayKey)(nil)).Elem()
}

func (o GatewayKeyOutput) ToGatewayKeyOutput() GatewayKeyOutput {
	return o
}

func (o GatewayKeyOutput) ToGatewayKeyOutputWithContext(ctx context.Context) GatewayKeyOutput {
	return o
}

// The UUID of your control plane. This variable is available in the Konnect manager.
func (o GatewayKeyOutput) ControlPlaneId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayKey) pulumi.StringOutput { return v.ControlPlaneId }).(pulumi.StringOutput)
}

// Unix epoch when the resource was created.
func (o GatewayKeyOutput) CreatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayKey) pulumi.IntOutput { return v.CreatedAt }).(pulumi.IntOutput)
}

// A JSON Web Key represented as a string.
func (o GatewayKeyOutput) Jwk() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayKey) pulumi.StringOutput { return v.Jwk }).(pulumi.StringOutput)
}

// A unique identifier for a key.
func (o GatewayKeyOutput) Kid() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayKey) pulumi.StringOutput { return v.Kid }).(pulumi.StringOutput)
}

// The name to associate with the given keys.
func (o GatewayKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A keypair in PEM format.
func (o GatewayKeyOutput) Pem() GatewayKeyPemOutput {
	return o.ApplyT(func(v *GatewayKey) GatewayKeyPemOutput { return v.Pem }).(GatewayKeyPemOutput)
}

// The id (an UUID) of the key-set with which to associate the key.
func (o GatewayKeyOutput) Set() GatewayKeySetTypeOutput {
	return o.ApplyT(func(v *GatewayKey) GatewayKeySetTypeOutput { return v.Set }).(GatewayKeySetTypeOutput)
}

// An optional set of strings associated with the Key for grouping and filtering.
func (o GatewayKeyOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayKey) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Unix epoch when the resource was last updated.
func (o GatewayKeyOutput) UpdatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayKey) pulumi.IntOutput { return v.UpdatedAt }).(pulumi.IntOutput)
}

type GatewayKeyArrayOutput struct{ *pulumi.OutputState }

func (GatewayKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayKey)(nil)).Elem()
}

func (o GatewayKeyArrayOutput) ToGatewayKeyArrayOutput() GatewayKeyArrayOutput {
	return o
}

func (o GatewayKeyArrayOutput) ToGatewayKeyArrayOutputWithContext(ctx context.Context) GatewayKeyArrayOutput {
	return o
}

func (o GatewayKeyArrayOutput) Index(i pulumi.IntInput) GatewayKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GatewayKey {
		return vs[0].([]*GatewayKey)[vs[1].(int)]
	}).(GatewayKeyOutput)
}

type GatewayKeyMapOutput struct{ *pulumi.OutputState }

func (GatewayKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayKey)(nil)).Elem()
}

func (o GatewayKeyMapOutput) ToGatewayKeyMapOutput() GatewayKeyMapOutput {
	return o
}

func (o GatewayKeyMapOutput) ToGatewayKeyMapOutputWithContext(ctx context.Context) GatewayKeyMapOutput {
	return o
}

func (o GatewayKeyMapOutput) MapIndex(k pulumi.StringInput) GatewayKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GatewayKey {
		return vs[0].(map[string]*GatewayKey)[vs[1].(string)]
	}).(GatewayKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayKeyInput)(nil)).Elem(), &GatewayKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayKeyArrayInput)(nil)).Elem(), GatewayKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayKeyMapInput)(nil)).Elem(), GatewayKeyMap{})
	pulumi.RegisterOutputType(GatewayKeyOutput{})
	pulumi.RegisterOutputType(GatewayKeyArrayOutput{})
	pulumi.RegisterOutputType(GatewayKeyMapOutput{})
}
