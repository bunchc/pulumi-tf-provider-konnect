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

// GatewayKeySet Resource
type GatewayKeySet struct {
	pulumi.CustomResourceState

	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId pulumi.StringOutput `pulumi:"controlPlaneId"`
	// Unix epoch when the resource was created.
	CreatedAt pulumi.IntOutput         `pulumi:"createdAt"`
	Name      pulumi.StringOutput      `pulumi:"name"`
	Tags      pulumi.StringArrayOutput `pulumi:"tags"`
	// Unix epoch when the resource was last updated.
	UpdatedAt pulumi.IntOutput `pulumi:"updatedAt"`
}

// NewGatewayKeySet registers a new resource with the given unique name, arguments, and options.
func NewGatewayKeySet(ctx *pulumi.Context,
	name string, args *GatewayKeySetArgs, opts ...pulumi.ResourceOption) (*GatewayKeySet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ControlPlaneId == nil {
		return nil, errors.New("invalid value for required argument 'ControlPlaneId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GatewayKeySet
	err := ctx.RegisterResource("konnect:index/gatewayKeySet:GatewayKeySet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayKeySet gets an existing GatewayKeySet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayKeySet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayKeySetState, opts ...pulumi.ResourceOption) (*GatewayKeySet, error) {
	var resource GatewayKeySet
	err := ctx.ReadResource("konnect:index/gatewayKeySet:GatewayKeySet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayKeySet resources.
type gatewayKeySetState struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId *string `pulumi:"controlPlaneId"`
	// Unix epoch when the resource was created.
	CreatedAt *int     `pulumi:"createdAt"`
	Name      *string  `pulumi:"name"`
	Tags      []string `pulumi:"tags"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int `pulumi:"updatedAt"`
}

type GatewayKeySetState struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId pulumi.StringPtrInput
	// Unix epoch when the resource was created.
	CreatedAt pulumi.IntPtrInput
	Name      pulumi.StringPtrInput
	Tags      pulumi.StringArrayInput
	// Unix epoch when the resource was last updated.
	UpdatedAt pulumi.IntPtrInput
}

func (GatewayKeySetState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayKeySetState)(nil)).Elem()
}

type gatewayKeySetArgs struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId string   `pulumi:"controlPlaneId"`
	Name           *string  `pulumi:"name"`
	Tags           []string `pulumi:"tags"`
}

// The set of arguments for constructing a GatewayKeySet resource.
type GatewayKeySetArgs struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId pulumi.StringInput
	Name           pulumi.StringPtrInput
	Tags           pulumi.StringArrayInput
}

func (GatewayKeySetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayKeySetArgs)(nil)).Elem()
}

type GatewayKeySetInput interface {
	pulumi.Input

	ToGatewayKeySetOutput() GatewayKeySetOutput
	ToGatewayKeySetOutputWithContext(ctx context.Context) GatewayKeySetOutput
}

func (*GatewayKeySet) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayKeySet)(nil)).Elem()
}

func (i *GatewayKeySet) ToGatewayKeySetOutput() GatewayKeySetOutput {
	return i.ToGatewayKeySetOutputWithContext(context.Background())
}

func (i *GatewayKeySet) ToGatewayKeySetOutputWithContext(ctx context.Context) GatewayKeySetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayKeySetOutput)
}

// GatewayKeySetArrayInput is an input type that accepts GatewayKeySetArray and GatewayKeySetArrayOutput values.
// You can construct a concrete instance of `GatewayKeySetArrayInput` via:
//
//	GatewayKeySetArray{ GatewayKeySetArgs{...} }
type GatewayKeySetArrayInput interface {
	pulumi.Input

	ToGatewayKeySetArrayOutput() GatewayKeySetArrayOutput
	ToGatewayKeySetArrayOutputWithContext(context.Context) GatewayKeySetArrayOutput
}

type GatewayKeySetArray []GatewayKeySetInput

func (GatewayKeySetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayKeySet)(nil)).Elem()
}

func (i GatewayKeySetArray) ToGatewayKeySetArrayOutput() GatewayKeySetArrayOutput {
	return i.ToGatewayKeySetArrayOutputWithContext(context.Background())
}

func (i GatewayKeySetArray) ToGatewayKeySetArrayOutputWithContext(ctx context.Context) GatewayKeySetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayKeySetArrayOutput)
}

// GatewayKeySetMapInput is an input type that accepts GatewayKeySetMap and GatewayKeySetMapOutput values.
// You can construct a concrete instance of `GatewayKeySetMapInput` via:
//
//	GatewayKeySetMap{ "key": GatewayKeySetArgs{...} }
type GatewayKeySetMapInput interface {
	pulumi.Input

	ToGatewayKeySetMapOutput() GatewayKeySetMapOutput
	ToGatewayKeySetMapOutputWithContext(context.Context) GatewayKeySetMapOutput
}

type GatewayKeySetMap map[string]GatewayKeySetInput

func (GatewayKeySetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayKeySet)(nil)).Elem()
}

func (i GatewayKeySetMap) ToGatewayKeySetMapOutput() GatewayKeySetMapOutput {
	return i.ToGatewayKeySetMapOutputWithContext(context.Background())
}

func (i GatewayKeySetMap) ToGatewayKeySetMapOutputWithContext(ctx context.Context) GatewayKeySetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayKeySetMapOutput)
}

type GatewayKeySetOutput struct{ *pulumi.OutputState }

func (GatewayKeySetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayKeySet)(nil)).Elem()
}

func (o GatewayKeySetOutput) ToGatewayKeySetOutput() GatewayKeySetOutput {
	return o
}

func (o GatewayKeySetOutput) ToGatewayKeySetOutputWithContext(ctx context.Context) GatewayKeySetOutput {
	return o
}

// The UUID of your control plane. This variable is available in the Konnect manager.
func (o GatewayKeySetOutput) ControlPlaneId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayKeySet) pulumi.StringOutput { return v.ControlPlaneId }).(pulumi.StringOutput)
}

// Unix epoch when the resource was created.
func (o GatewayKeySetOutput) CreatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayKeySet) pulumi.IntOutput { return v.CreatedAt }).(pulumi.IntOutput)
}

func (o GatewayKeySetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayKeySet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GatewayKeySetOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayKeySet) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Unix epoch when the resource was last updated.
func (o GatewayKeySetOutput) UpdatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayKeySet) pulumi.IntOutput { return v.UpdatedAt }).(pulumi.IntOutput)
}

type GatewayKeySetArrayOutput struct{ *pulumi.OutputState }

func (GatewayKeySetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayKeySet)(nil)).Elem()
}

func (o GatewayKeySetArrayOutput) ToGatewayKeySetArrayOutput() GatewayKeySetArrayOutput {
	return o
}

func (o GatewayKeySetArrayOutput) ToGatewayKeySetArrayOutputWithContext(ctx context.Context) GatewayKeySetArrayOutput {
	return o
}

func (o GatewayKeySetArrayOutput) Index(i pulumi.IntInput) GatewayKeySetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GatewayKeySet {
		return vs[0].([]*GatewayKeySet)[vs[1].(int)]
	}).(GatewayKeySetOutput)
}

type GatewayKeySetMapOutput struct{ *pulumi.OutputState }

func (GatewayKeySetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayKeySet)(nil)).Elem()
}

func (o GatewayKeySetMapOutput) ToGatewayKeySetMapOutput() GatewayKeySetMapOutput {
	return o
}

func (o GatewayKeySetMapOutput) ToGatewayKeySetMapOutputWithContext(ctx context.Context) GatewayKeySetMapOutput {
	return o
}

func (o GatewayKeySetMapOutput) MapIndex(k pulumi.StringInput) GatewayKeySetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GatewayKeySet {
		return vs[0].(map[string]*GatewayKeySet)[vs[1].(string)]
	}).(GatewayKeySetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayKeySetInput)(nil)).Elem(), &GatewayKeySet{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayKeySetArrayInput)(nil)).Elem(), GatewayKeySetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayKeySetMapInput)(nil)).Elem(), GatewayKeySetMap{})
	pulumi.RegisterOutputType(GatewayKeySetOutput{})
	pulumi.RegisterOutputType(GatewayKeySetArrayOutput{})
	pulumi.RegisterOutputType(GatewayKeySetMapOutput{})
}