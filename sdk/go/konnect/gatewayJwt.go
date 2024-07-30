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

// GatewayJWT Resource
type GatewayJwt struct {
	pulumi.CustomResourceState

	// Requires replacement if changed. ; must be one of ["HS256", "HS384", "HS512", "RS256", "RS384", "RS512", "ES256", "ES384", "ES512", "PS256", "PS384", "PS512", "EdDSA"]
	Algorithm pulumi.StringOutput      `pulumi:"algorithm"`
	Consumer  GatewayJwtConsumerOutput `pulumi:"consumer"`
	// Consumer ID for nested entities. Requires replacement if changed.
	ConsumerId pulumi.StringOutput `pulumi:"consumerId"`
	// The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.
	ControlPlaneId pulumi.StringOutput `pulumi:"controlPlaneId"`
	// Unix epoch when the resource was created.
	CreatedAt pulumi.IntOutput `pulumi:"createdAt"`
	// Requires replacement if changed.
	Key pulumi.StringOutput `pulumi:"key"`
	// Requires replacement if changed.
	RsaPublicKey pulumi.StringOutput `pulumi:"rsaPublicKey"`
	// Requires replacement if changed.
	Secret pulumi.StringOutput `pulumi:"secret"`
	// Requires replacement if changed.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewGatewayJwt registers a new resource with the given unique name, arguments, and options.
func NewGatewayJwt(ctx *pulumi.Context,
	name string, args *GatewayJwtArgs, opts ...pulumi.ResourceOption) (*GatewayJwt, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConsumerId == nil {
		return nil, errors.New("invalid value for required argument 'ConsumerId'")
	}
	if args.ControlPlaneId == nil {
		return nil, errors.New("invalid value for required argument 'ControlPlaneId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GatewayJwt
	err := ctx.RegisterResource("konnect:index/gatewayJwt:GatewayJwt", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayJwt gets an existing GatewayJwt resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayJwt(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayJwtState, opts ...pulumi.ResourceOption) (*GatewayJwt, error) {
	var resource GatewayJwt
	err := ctx.ReadResource("konnect:index/gatewayJwt:GatewayJwt", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayJwt resources.
type gatewayJwtState struct {
	// Requires replacement if changed. ; must be one of ["HS256", "HS384", "HS512", "RS256", "RS384", "RS512", "ES256", "ES384", "ES512", "PS256", "PS384", "PS512", "EdDSA"]
	Algorithm *string             `pulumi:"algorithm"`
	Consumer  *GatewayJwtConsumer `pulumi:"consumer"`
	// Consumer ID for nested entities. Requires replacement if changed.
	ConsumerId *string `pulumi:"consumerId"`
	// The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.
	ControlPlaneId *string `pulumi:"controlPlaneId"`
	// Unix epoch when the resource was created.
	CreatedAt *int `pulumi:"createdAt"`
	// Requires replacement if changed.
	Key *string `pulumi:"key"`
	// Requires replacement if changed.
	RsaPublicKey *string `pulumi:"rsaPublicKey"`
	// Requires replacement if changed.
	Secret *string `pulumi:"secret"`
	// Requires replacement if changed.
	Tags []string `pulumi:"tags"`
}

type GatewayJwtState struct {
	// Requires replacement if changed. ; must be one of ["HS256", "HS384", "HS512", "RS256", "RS384", "RS512", "ES256", "ES384", "ES512", "PS256", "PS384", "PS512", "EdDSA"]
	Algorithm pulumi.StringPtrInput
	Consumer  GatewayJwtConsumerPtrInput
	// Consumer ID for nested entities. Requires replacement if changed.
	ConsumerId pulumi.StringPtrInput
	// The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.
	ControlPlaneId pulumi.StringPtrInput
	// Unix epoch when the resource was created.
	CreatedAt pulumi.IntPtrInput
	// Requires replacement if changed.
	Key pulumi.StringPtrInput
	// Requires replacement if changed.
	RsaPublicKey pulumi.StringPtrInput
	// Requires replacement if changed.
	Secret pulumi.StringPtrInput
	// Requires replacement if changed.
	Tags pulumi.StringArrayInput
}

func (GatewayJwtState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayJwtState)(nil)).Elem()
}

type gatewayJwtArgs struct {
	// Requires replacement if changed. ; must be one of ["HS256", "HS384", "HS512", "RS256", "RS384", "RS512", "ES256", "ES384", "ES512", "PS256", "PS384", "PS512", "EdDSA"]
	Algorithm *string `pulumi:"algorithm"`
	// Consumer ID for nested entities. Requires replacement if changed.
	ConsumerId string `pulumi:"consumerId"`
	// The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.
	ControlPlaneId string `pulumi:"controlPlaneId"`
	// Requires replacement if changed.
	Key *string `pulumi:"key"`
	// Requires replacement if changed.
	RsaPublicKey *string `pulumi:"rsaPublicKey"`
	// Requires replacement if changed.
	Secret *string `pulumi:"secret"`
	// Requires replacement if changed.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a GatewayJwt resource.
type GatewayJwtArgs struct {
	// Requires replacement if changed. ; must be one of ["HS256", "HS384", "HS512", "RS256", "RS384", "RS512", "ES256", "ES384", "ES512", "PS256", "PS384", "PS512", "EdDSA"]
	Algorithm pulumi.StringPtrInput
	// Consumer ID for nested entities. Requires replacement if changed.
	ConsumerId pulumi.StringInput
	// The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.
	ControlPlaneId pulumi.StringInput
	// Requires replacement if changed.
	Key pulumi.StringPtrInput
	// Requires replacement if changed.
	RsaPublicKey pulumi.StringPtrInput
	// Requires replacement if changed.
	Secret pulumi.StringPtrInput
	// Requires replacement if changed.
	Tags pulumi.StringArrayInput
}

func (GatewayJwtArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayJwtArgs)(nil)).Elem()
}

type GatewayJwtInput interface {
	pulumi.Input

	ToGatewayJwtOutput() GatewayJwtOutput
	ToGatewayJwtOutputWithContext(ctx context.Context) GatewayJwtOutput
}

func (*GatewayJwt) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayJwt)(nil)).Elem()
}

func (i *GatewayJwt) ToGatewayJwtOutput() GatewayJwtOutput {
	return i.ToGatewayJwtOutputWithContext(context.Background())
}

func (i *GatewayJwt) ToGatewayJwtOutputWithContext(ctx context.Context) GatewayJwtOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayJwtOutput)
}

// GatewayJwtArrayInput is an input type that accepts GatewayJwtArray and GatewayJwtArrayOutput values.
// You can construct a concrete instance of `GatewayJwtArrayInput` via:
//
//	GatewayJwtArray{ GatewayJwtArgs{...} }
type GatewayJwtArrayInput interface {
	pulumi.Input

	ToGatewayJwtArrayOutput() GatewayJwtArrayOutput
	ToGatewayJwtArrayOutputWithContext(context.Context) GatewayJwtArrayOutput
}

type GatewayJwtArray []GatewayJwtInput

func (GatewayJwtArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayJwt)(nil)).Elem()
}

func (i GatewayJwtArray) ToGatewayJwtArrayOutput() GatewayJwtArrayOutput {
	return i.ToGatewayJwtArrayOutputWithContext(context.Background())
}

func (i GatewayJwtArray) ToGatewayJwtArrayOutputWithContext(ctx context.Context) GatewayJwtArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayJwtArrayOutput)
}

// GatewayJwtMapInput is an input type that accepts GatewayJwtMap and GatewayJwtMapOutput values.
// You can construct a concrete instance of `GatewayJwtMapInput` via:
//
//	GatewayJwtMap{ "key": GatewayJwtArgs{...} }
type GatewayJwtMapInput interface {
	pulumi.Input

	ToGatewayJwtMapOutput() GatewayJwtMapOutput
	ToGatewayJwtMapOutputWithContext(context.Context) GatewayJwtMapOutput
}

type GatewayJwtMap map[string]GatewayJwtInput

func (GatewayJwtMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayJwt)(nil)).Elem()
}

func (i GatewayJwtMap) ToGatewayJwtMapOutput() GatewayJwtMapOutput {
	return i.ToGatewayJwtMapOutputWithContext(context.Background())
}

func (i GatewayJwtMap) ToGatewayJwtMapOutputWithContext(ctx context.Context) GatewayJwtMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayJwtMapOutput)
}

type GatewayJwtOutput struct{ *pulumi.OutputState }

func (GatewayJwtOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayJwt)(nil)).Elem()
}

func (o GatewayJwtOutput) ToGatewayJwtOutput() GatewayJwtOutput {
	return o
}

func (o GatewayJwtOutput) ToGatewayJwtOutputWithContext(ctx context.Context) GatewayJwtOutput {
	return o
}

// Requires replacement if changed. ; must be one of ["HS256", "HS384", "HS512", "RS256", "RS384", "RS512", "ES256", "ES384", "ES512", "PS256", "PS384", "PS512", "EdDSA"]
func (o GatewayJwtOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayJwt) pulumi.StringOutput { return v.Algorithm }).(pulumi.StringOutput)
}

func (o GatewayJwtOutput) Consumer() GatewayJwtConsumerOutput {
	return o.ApplyT(func(v *GatewayJwt) GatewayJwtConsumerOutput { return v.Consumer }).(GatewayJwtConsumerOutput)
}

// Consumer ID for nested entities. Requires replacement if changed.
func (o GatewayJwtOutput) ConsumerId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayJwt) pulumi.StringOutput { return v.ConsumerId }).(pulumi.StringOutput)
}

// The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.
func (o GatewayJwtOutput) ControlPlaneId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayJwt) pulumi.StringOutput { return v.ControlPlaneId }).(pulumi.StringOutput)
}

// Unix epoch when the resource was created.
func (o GatewayJwtOutput) CreatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayJwt) pulumi.IntOutput { return v.CreatedAt }).(pulumi.IntOutput)
}

// Requires replacement if changed.
func (o GatewayJwtOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayJwt) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Requires replacement if changed.
func (o GatewayJwtOutput) RsaPublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayJwt) pulumi.StringOutput { return v.RsaPublicKey }).(pulumi.StringOutput)
}

// Requires replacement if changed.
func (o GatewayJwtOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayJwt) pulumi.StringOutput { return v.Secret }).(pulumi.StringOutput)
}

// Requires replacement if changed.
func (o GatewayJwtOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayJwt) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

type GatewayJwtArrayOutput struct{ *pulumi.OutputState }

func (GatewayJwtArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayJwt)(nil)).Elem()
}

func (o GatewayJwtArrayOutput) ToGatewayJwtArrayOutput() GatewayJwtArrayOutput {
	return o
}

func (o GatewayJwtArrayOutput) ToGatewayJwtArrayOutputWithContext(ctx context.Context) GatewayJwtArrayOutput {
	return o
}

func (o GatewayJwtArrayOutput) Index(i pulumi.IntInput) GatewayJwtOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GatewayJwt {
		return vs[0].([]*GatewayJwt)[vs[1].(int)]
	}).(GatewayJwtOutput)
}

type GatewayJwtMapOutput struct{ *pulumi.OutputState }

func (GatewayJwtMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayJwt)(nil)).Elem()
}

func (o GatewayJwtMapOutput) ToGatewayJwtMapOutput() GatewayJwtMapOutput {
	return o
}

func (o GatewayJwtMapOutput) ToGatewayJwtMapOutputWithContext(ctx context.Context) GatewayJwtMapOutput {
	return o
}

func (o GatewayJwtMapOutput) MapIndex(k pulumi.StringInput) GatewayJwtOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GatewayJwt {
		return vs[0].(map[string]*GatewayJwt)[vs[1].(string)]
	}).(GatewayJwtOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayJwtInput)(nil)).Elem(), &GatewayJwt{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayJwtArrayInput)(nil)).Elem(), GatewayJwtArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayJwtMapInput)(nil)).Elem(), GatewayJwtMap{})
	pulumi.RegisterOutputType(GatewayJwtOutput{})
	pulumi.RegisterOutputType(GatewayJwtArrayOutput{})
	pulumi.RegisterOutputType(GatewayJwtMapOutput{})
}
