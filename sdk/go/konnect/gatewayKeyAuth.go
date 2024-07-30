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

// GatewayKeyAuth Resource
type GatewayKeyAuth struct {
	pulumi.CustomResourceState

	Consumer GatewayKeyAuthConsumerOutput `pulumi:"consumer"`
	// Consumer ID for nested entities. Requires replacement if changed.
	ConsumerId pulumi.StringOutput `pulumi:"consumerId"`
	// The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.
	ControlPlaneId pulumi.StringOutput `pulumi:"controlPlaneId"`
	// Unix epoch when the resource was created.
	CreatedAt pulumi.IntOutput `pulumi:"createdAt"`
	// Requires replacement if changed.
	Key pulumi.StringOutput `pulumi:"key"`
	// Requires replacement if changed.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewGatewayKeyAuth registers a new resource with the given unique name, arguments, and options.
func NewGatewayKeyAuth(ctx *pulumi.Context,
	name string, args *GatewayKeyAuthArgs, opts ...pulumi.ResourceOption) (*GatewayKeyAuth, error) {
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
	var resource GatewayKeyAuth
	err := ctx.RegisterResource("konnect:index/gatewayKeyAuth:GatewayKeyAuth", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayKeyAuth gets an existing GatewayKeyAuth resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayKeyAuth(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayKeyAuthState, opts ...pulumi.ResourceOption) (*GatewayKeyAuth, error) {
	var resource GatewayKeyAuth
	err := ctx.ReadResource("konnect:index/gatewayKeyAuth:GatewayKeyAuth", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayKeyAuth resources.
type gatewayKeyAuthState struct {
	Consumer *GatewayKeyAuthConsumer `pulumi:"consumer"`
	// Consumer ID for nested entities. Requires replacement if changed.
	ConsumerId *string `pulumi:"consumerId"`
	// The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.
	ControlPlaneId *string `pulumi:"controlPlaneId"`
	// Unix epoch when the resource was created.
	CreatedAt *int `pulumi:"createdAt"`
	// Requires replacement if changed.
	Key *string `pulumi:"key"`
	// Requires replacement if changed.
	Tags []string `pulumi:"tags"`
}

type GatewayKeyAuthState struct {
	Consumer GatewayKeyAuthConsumerPtrInput
	// Consumer ID for nested entities. Requires replacement if changed.
	ConsumerId pulumi.StringPtrInput
	// The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.
	ControlPlaneId pulumi.StringPtrInput
	// Unix epoch when the resource was created.
	CreatedAt pulumi.IntPtrInput
	// Requires replacement if changed.
	Key pulumi.StringPtrInput
	// Requires replacement if changed.
	Tags pulumi.StringArrayInput
}

func (GatewayKeyAuthState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayKeyAuthState)(nil)).Elem()
}

type gatewayKeyAuthArgs struct {
	// Consumer ID for nested entities. Requires replacement if changed.
	ConsumerId string `pulumi:"consumerId"`
	// The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.
	ControlPlaneId string `pulumi:"controlPlaneId"`
	// Requires replacement if changed.
	Key *string `pulumi:"key"`
	// Requires replacement if changed.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a GatewayKeyAuth resource.
type GatewayKeyAuthArgs struct {
	// Consumer ID for nested entities. Requires replacement if changed.
	ConsumerId pulumi.StringInput
	// The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.
	ControlPlaneId pulumi.StringInput
	// Requires replacement if changed.
	Key pulumi.StringPtrInput
	// Requires replacement if changed.
	Tags pulumi.StringArrayInput
}

func (GatewayKeyAuthArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayKeyAuthArgs)(nil)).Elem()
}

type GatewayKeyAuthInput interface {
	pulumi.Input

	ToGatewayKeyAuthOutput() GatewayKeyAuthOutput
	ToGatewayKeyAuthOutputWithContext(ctx context.Context) GatewayKeyAuthOutput
}

func (*GatewayKeyAuth) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayKeyAuth)(nil)).Elem()
}

func (i *GatewayKeyAuth) ToGatewayKeyAuthOutput() GatewayKeyAuthOutput {
	return i.ToGatewayKeyAuthOutputWithContext(context.Background())
}

func (i *GatewayKeyAuth) ToGatewayKeyAuthOutputWithContext(ctx context.Context) GatewayKeyAuthOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayKeyAuthOutput)
}

// GatewayKeyAuthArrayInput is an input type that accepts GatewayKeyAuthArray and GatewayKeyAuthArrayOutput values.
// You can construct a concrete instance of `GatewayKeyAuthArrayInput` via:
//
//	GatewayKeyAuthArray{ GatewayKeyAuthArgs{...} }
type GatewayKeyAuthArrayInput interface {
	pulumi.Input

	ToGatewayKeyAuthArrayOutput() GatewayKeyAuthArrayOutput
	ToGatewayKeyAuthArrayOutputWithContext(context.Context) GatewayKeyAuthArrayOutput
}

type GatewayKeyAuthArray []GatewayKeyAuthInput

func (GatewayKeyAuthArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayKeyAuth)(nil)).Elem()
}

func (i GatewayKeyAuthArray) ToGatewayKeyAuthArrayOutput() GatewayKeyAuthArrayOutput {
	return i.ToGatewayKeyAuthArrayOutputWithContext(context.Background())
}

func (i GatewayKeyAuthArray) ToGatewayKeyAuthArrayOutputWithContext(ctx context.Context) GatewayKeyAuthArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayKeyAuthArrayOutput)
}

// GatewayKeyAuthMapInput is an input type that accepts GatewayKeyAuthMap and GatewayKeyAuthMapOutput values.
// You can construct a concrete instance of `GatewayKeyAuthMapInput` via:
//
//	GatewayKeyAuthMap{ "key": GatewayKeyAuthArgs{...} }
type GatewayKeyAuthMapInput interface {
	pulumi.Input

	ToGatewayKeyAuthMapOutput() GatewayKeyAuthMapOutput
	ToGatewayKeyAuthMapOutputWithContext(context.Context) GatewayKeyAuthMapOutput
}

type GatewayKeyAuthMap map[string]GatewayKeyAuthInput

func (GatewayKeyAuthMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayKeyAuth)(nil)).Elem()
}

func (i GatewayKeyAuthMap) ToGatewayKeyAuthMapOutput() GatewayKeyAuthMapOutput {
	return i.ToGatewayKeyAuthMapOutputWithContext(context.Background())
}

func (i GatewayKeyAuthMap) ToGatewayKeyAuthMapOutputWithContext(ctx context.Context) GatewayKeyAuthMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayKeyAuthMapOutput)
}

type GatewayKeyAuthOutput struct{ *pulumi.OutputState }

func (GatewayKeyAuthOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayKeyAuth)(nil)).Elem()
}

func (o GatewayKeyAuthOutput) ToGatewayKeyAuthOutput() GatewayKeyAuthOutput {
	return o
}

func (o GatewayKeyAuthOutput) ToGatewayKeyAuthOutputWithContext(ctx context.Context) GatewayKeyAuthOutput {
	return o
}

func (o GatewayKeyAuthOutput) Consumer() GatewayKeyAuthConsumerOutput {
	return o.ApplyT(func(v *GatewayKeyAuth) GatewayKeyAuthConsumerOutput { return v.Consumer }).(GatewayKeyAuthConsumerOutput)
}

// Consumer ID for nested entities. Requires replacement if changed.
func (o GatewayKeyAuthOutput) ConsumerId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayKeyAuth) pulumi.StringOutput { return v.ConsumerId }).(pulumi.StringOutput)
}

// The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.
func (o GatewayKeyAuthOutput) ControlPlaneId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayKeyAuth) pulumi.StringOutput { return v.ControlPlaneId }).(pulumi.StringOutput)
}

// Unix epoch when the resource was created.
func (o GatewayKeyAuthOutput) CreatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayKeyAuth) pulumi.IntOutput { return v.CreatedAt }).(pulumi.IntOutput)
}

// Requires replacement if changed.
func (o GatewayKeyAuthOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayKeyAuth) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Requires replacement if changed.
func (o GatewayKeyAuthOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayKeyAuth) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

type GatewayKeyAuthArrayOutput struct{ *pulumi.OutputState }

func (GatewayKeyAuthArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayKeyAuth)(nil)).Elem()
}

func (o GatewayKeyAuthArrayOutput) ToGatewayKeyAuthArrayOutput() GatewayKeyAuthArrayOutput {
	return o
}

func (o GatewayKeyAuthArrayOutput) ToGatewayKeyAuthArrayOutputWithContext(ctx context.Context) GatewayKeyAuthArrayOutput {
	return o
}

func (o GatewayKeyAuthArrayOutput) Index(i pulumi.IntInput) GatewayKeyAuthOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GatewayKeyAuth {
		return vs[0].([]*GatewayKeyAuth)[vs[1].(int)]
	}).(GatewayKeyAuthOutput)
}

type GatewayKeyAuthMapOutput struct{ *pulumi.OutputState }

func (GatewayKeyAuthMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayKeyAuth)(nil)).Elem()
}

func (o GatewayKeyAuthMapOutput) ToGatewayKeyAuthMapOutput() GatewayKeyAuthMapOutput {
	return o
}

func (o GatewayKeyAuthMapOutput) ToGatewayKeyAuthMapOutputWithContext(ctx context.Context) GatewayKeyAuthMapOutput {
	return o
}

func (o GatewayKeyAuthMapOutput) MapIndex(k pulumi.StringInput) GatewayKeyAuthOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GatewayKeyAuth {
		return vs[0].(map[string]*GatewayKeyAuth)[vs[1].(string)]
	}).(GatewayKeyAuthOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayKeyAuthInput)(nil)).Elem(), &GatewayKeyAuth{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayKeyAuthArrayInput)(nil)).Elem(), GatewayKeyAuthArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayKeyAuthMapInput)(nil)).Elem(), GatewayKeyAuthMap{})
	pulumi.RegisterOutputType(GatewayKeyAuthOutput{})
	pulumi.RegisterOutputType(GatewayKeyAuthArrayOutput{})
	pulumi.RegisterOutputType(GatewayKeyAuthMapOutput{})
}
