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

// GatewayVault Resource
type GatewayVault struct {
	pulumi.CustomResourceState

	// The configuration properties for the Vault which can be found on the vaults' documentation page. Parsed as JSON.
	Config pulumi.StringOutput `pulumi:"config"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId pulumi.StringOutput `pulumi:"controlPlaneId"`
	// Unix epoch when the resource was created.
	CreatedAt pulumi.IntOutput `pulumi:"createdAt"`
	// The description of the Vault entity.
	Description pulumi.StringOutput `pulumi:"description"`
	// The name of the Vault that's going to be added. Currently, the Vault implementation must be installed in every Kong instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// The unique prefix (or identifier) for this Vault configuration. The prefix is used to load the right Vault configuration and implementation when referencing secrets with the other entities.
	Prefix pulumi.StringOutput `pulumi:"prefix"`
	// An optional set of strings associated with the Vault for grouping and filtering.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Unix epoch when the resource was last updated.
	UpdatedAt pulumi.IntOutput `pulumi:"updatedAt"`
}

// NewGatewayVault registers a new resource with the given unique name, arguments, and options.
func NewGatewayVault(ctx *pulumi.Context,
	name string, args *GatewayVaultArgs, opts ...pulumi.ResourceOption) (*GatewayVault, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ControlPlaneId == nil {
		return nil, errors.New("invalid value for required argument 'ControlPlaneId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GatewayVault
	err := ctx.RegisterResource("konnect:index/gatewayVault:GatewayVault", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayVault gets an existing GatewayVault resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayVault(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayVaultState, opts ...pulumi.ResourceOption) (*GatewayVault, error) {
	var resource GatewayVault
	err := ctx.ReadResource("konnect:index/gatewayVault:GatewayVault", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayVault resources.
type gatewayVaultState struct {
	// The configuration properties for the Vault which can be found on the vaults' documentation page. Parsed as JSON.
	Config *string `pulumi:"config"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId *string `pulumi:"controlPlaneId"`
	// Unix epoch when the resource was created.
	CreatedAt *int `pulumi:"createdAt"`
	// The description of the Vault entity.
	Description *string `pulumi:"description"`
	// The name of the Vault that's going to be added. Currently, the Vault implementation must be installed in every Kong instance.
	Name *string `pulumi:"name"`
	// The unique prefix (or identifier) for this Vault configuration. The prefix is used to load the right Vault configuration and implementation when referencing secrets with the other entities.
	Prefix *string `pulumi:"prefix"`
	// An optional set of strings associated with the Vault for grouping and filtering.
	Tags []string `pulumi:"tags"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int `pulumi:"updatedAt"`
}

type GatewayVaultState struct {
	// The configuration properties for the Vault which can be found on the vaults' documentation page. Parsed as JSON.
	Config pulumi.StringPtrInput
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId pulumi.StringPtrInput
	// Unix epoch when the resource was created.
	CreatedAt pulumi.IntPtrInput
	// The description of the Vault entity.
	Description pulumi.StringPtrInput
	// The name of the Vault that's going to be added. Currently, the Vault implementation must be installed in every Kong instance.
	Name pulumi.StringPtrInput
	// The unique prefix (or identifier) for this Vault configuration. The prefix is used to load the right Vault configuration and implementation when referencing secrets with the other entities.
	Prefix pulumi.StringPtrInput
	// An optional set of strings associated with the Vault for grouping and filtering.
	Tags pulumi.StringArrayInput
	// Unix epoch when the resource was last updated.
	UpdatedAt pulumi.IntPtrInput
}

func (GatewayVaultState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayVaultState)(nil)).Elem()
}

type gatewayVaultArgs struct {
	// The configuration properties for the Vault which can be found on the vaults' documentation page. Parsed as JSON.
	Config *string `pulumi:"config"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId string `pulumi:"controlPlaneId"`
	// The description of the Vault entity.
	Description *string `pulumi:"description"`
	// The name of the Vault that's going to be added. Currently, the Vault implementation must be installed in every Kong instance.
	Name *string `pulumi:"name"`
	// The unique prefix (or identifier) for this Vault configuration. The prefix is used to load the right Vault configuration and implementation when referencing secrets with the other entities.
	Prefix *string `pulumi:"prefix"`
	// An optional set of strings associated with the Vault for grouping and filtering.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a GatewayVault resource.
type GatewayVaultArgs struct {
	// The configuration properties for the Vault which can be found on the vaults' documentation page. Parsed as JSON.
	Config pulumi.StringPtrInput
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId pulumi.StringInput
	// The description of the Vault entity.
	Description pulumi.StringPtrInput
	// The name of the Vault that's going to be added. Currently, the Vault implementation must be installed in every Kong instance.
	Name pulumi.StringPtrInput
	// The unique prefix (or identifier) for this Vault configuration. The prefix is used to load the right Vault configuration and implementation when referencing secrets with the other entities.
	Prefix pulumi.StringPtrInput
	// An optional set of strings associated with the Vault for grouping and filtering.
	Tags pulumi.StringArrayInput
}

func (GatewayVaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayVaultArgs)(nil)).Elem()
}

type GatewayVaultInput interface {
	pulumi.Input

	ToGatewayVaultOutput() GatewayVaultOutput
	ToGatewayVaultOutputWithContext(ctx context.Context) GatewayVaultOutput
}

func (*GatewayVault) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayVault)(nil)).Elem()
}

func (i *GatewayVault) ToGatewayVaultOutput() GatewayVaultOutput {
	return i.ToGatewayVaultOutputWithContext(context.Background())
}

func (i *GatewayVault) ToGatewayVaultOutputWithContext(ctx context.Context) GatewayVaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayVaultOutput)
}

// GatewayVaultArrayInput is an input type that accepts GatewayVaultArray and GatewayVaultArrayOutput values.
// You can construct a concrete instance of `GatewayVaultArrayInput` via:
//
//	GatewayVaultArray{ GatewayVaultArgs{...} }
type GatewayVaultArrayInput interface {
	pulumi.Input

	ToGatewayVaultArrayOutput() GatewayVaultArrayOutput
	ToGatewayVaultArrayOutputWithContext(context.Context) GatewayVaultArrayOutput
}

type GatewayVaultArray []GatewayVaultInput

func (GatewayVaultArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayVault)(nil)).Elem()
}

func (i GatewayVaultArray) ToGatewayVaultArrayOutput() GatewayVaultArrayOutput {
	return i.ToGatewayVaultArrayOutputWithContext(context.Background())
}

func (i GatewayVaultArray) ToGatewayVaultArrayOutputWithContext(ctx context.Context) GatewayVaultArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayVaultArrayOutput)
}

// GatewayVaultMapInput is an input type that accepts GatewayVaultMap and GatewayVaultMapOutput values.
// You can construct a concrete instance of `GatewayVaultMapInput` via:
//
//	GatewayVaultMap{ "key": GatewayVaultArgs{...} }
type GatewayVaultMapInput interface {
	pulumi.Input

	ToGatewayVaultMapOutput() GatewayVaultMapOutput
	ToGatewayVaultMapOutputWithContext(context.Context) GatewayVaultMapOutput
}

type GatewayVaultMap map[string]GatewayVaultInput

func (GatewayVaultMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayVault)(nil)).Elem()
}

func (i GatewayVaultMap) ToGatewayVaultMapOutput() GatewayVaultMapOutput {
	return i.ToGatewayVaultMapOutputWithContext(context.Background())
}

func (i GatewayVaultMap) ToGatewayVaultMapOutputWithContext(ctx context.Context) GatewayVaultMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayVaultMapOutput)
}

type GatewayVaultOutput struct{ *pulumi.OutputState }

func (GatewayVaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayVault)(nil)).Elem()
}

func (o GatewayVaultOutput) ToGatewayVaultOutput() GatewayVaultOutput {
	return o
}

func (o GatewayVaultOutput) ToGatewayVaultOutputWithContext(ctx context.Context) GatewayVaultOutput {
	return o
}

// The configuration properties for the Vault which can be found on the vaults' documentation page. Parsed as JSON.
func (o GatewayVaultOutput) Config() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayVault) pulumi.StringOutput { return v.Config }).(pulumi.StringOutput)
}

// The UUID of your control plane. This variable is available in the Konnect manager.
func (o GatewayVaultOutput) ControlPlaneId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayVault) pulumi.StringOutput { return v.ControlPlaneId }).(pulumi.StringOutput)
}

// Unix epoch when the resource was created.
func (o GatewayVaultOutput) CreatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayVault) pulumi.IntOutput { return v.CreatedAt }).(pulumi.IntOutput)
}

// The description of the Vault entity.
func (o GatewayVaultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayVault) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The name of the Vault that's going to be added. Currently, the Vault implementation must be installed in every Kong instance.
func (o GatewayVaultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayVault) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The unique prefix (or identifier) for this Vault configuration. The prefix is used to load the right Vault configuration and implementation when referencing secrets with the other entities.
func (o GatewayVaultOutput) Prefix() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayVault) pulumi.StringOutput { return v.Prefix }).(pulumi.StringOutput)
}

// An optional set of strings associated with the Vault for grouping and filtering.
func (o GatewayVaultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayVault) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Unix epoch when the resource was last updated.
func (o GatewayVaultOutput) UpdatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayVault) pulumi.IntOutput { return v.UpdatedAt }).(pulumi.IntOutput)
}

type GatewayVaultArrayOutput struct{ *pulumi.OutputState }

func (GatewayVaultArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayVault)(nil)).Elem()
}

func (o GatewayVaultArrayOutput) ToGatewayVaultArrayOutput() GatewayVaultArrayOutput {
	return o
}

func (o GatewayVaultArrayOutput) ToGatewayVaultArrayOutputWithContext(ctx context.Context) GatewayVaultArrayOutput {
	return o
}

func (o GatewayVaultArrayOutput) Index(i pulumi.IntInput) GatewayVaultOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GatewayVault {
		return vs[0].([]*GatewayVault)[vs[1].(int)]
	}).(GatewayVaultOutput)
}

type GatewayVaultMapOutput struct{ *pulumi.OutputState }

func (GatewayVaultMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayVault)(nil)).Elem()
}

func (o GatewayVaultMapOutput) ToGatewayVaultMapOutput() GatewayVaultMapOutput {
	return o
}

func (o GatewayVaultMapOutput) ToGatewayVaultMapOutputWithContext(ctx context.Context) GatewayVaultMapOutput {
	return o
}

func (o GatewayVaultMapOutput) MapIndex(k pulumi.StringInput) GatewayVaultOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GatewayVault {
		return vs[0].(map[string]*GatewayVault)[vs[1].(string)]
	}).(GatewayVaultOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayVaultInput)(nil)).Elem(), &GatewayVault{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayVaultArrayInput)(nil)).Elem(), GatewayVaultArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayVaultMapInput)(nil)).Elem(), GatewayVaultMap{})
	pulumi.RegisterOutputType(GatewayVaultOutput{})
	pulumi.RegisterOutputType(GatewayVaultArrayOutput{})
	pulumi.RegisterOutputType(GatewayVaultMapOutput{})
}