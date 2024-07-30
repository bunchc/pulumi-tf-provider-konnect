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

// SystemAccountRole Resource
type SystemAccountRole struct {
	pulumi.CustomResourceState

	// ID of the system account. Requires replacement if changed.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The ID of the entity. Requires replacement if changed.
	EntityId pulumi.StringOutput `pulumi:"entityId"`
	// The region of the team. Requires replacement if changed. ; must be one of ["us", "eu", "au", "*"]
	EntityRegion pulumi.StringOutput `pulumi:"entityRegion"`
	// The type of entity. Requires replacement if changed. ; must be one of ["API Products", "Application Auth Strategies", "Audit Logs", "Control Planes", "DCR Providers", "Identity", "Mesh Control Planes", "Networks", "Portals", "Service Hub"]
	EntityTypeName pulumi.StringOutput `pulumi:"entityTypeName"`
	// The desired role. Requires replacement if changed. ; must be one of ["Admin", "Appearance Maintainer", "Application Registration", "Certificate Admin", "Cloud Gateway Cluster Admin", "Cloud Gateway Cluster Viewer", "Consumer Admin", "Creator", "Deployer", "Discovery Admin", "Discovery Viewer", "Gateway Service Admin", "Integration Admin", "Integration Viewer", "Key Admin", "Maintainer", "Network Admin", "Network Creator", "Network Viewer", "Plugin Admin", "Plugins Admin", "Product Publisher", "Publisher", "Route Admin", "SNI Admin", "Service Admin", "Service Creator", "Service Viewer", "Upstream Admin", "Vault Admin", "Viewer"]
	RoleName pulumi.StringOutput `pulumi:"roleName"`
}

// NewSystemAccountRole registers a new resource with the given unique name, arguments, and options.
func NewSystemAccountRole(ctx *pulumi.Context,
	name string, args *SystemAccountRoleArgs, opts ...pulumi.ResourceOption) (*SystemAccountRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountId == nil {
		return nil, errors.New("invalid value for required argument 'AccountId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemAccountRole
	err := ctx.RegisterResource("konnect:index/systemAccountRole:SystemAccountRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemAccountRole gets an existing SystemAccountRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemAccountRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemAccountRoleState, opts ...pulumi.ResourceOption) (*SystemAccountRole, error) {
	var resource SystemAccountRole
	err := ctx.ReadResource("konnect:index/systemAccountRole:SystemAccountRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemAccountRole resources.
type systemAccountRoleState struct {
	// ID of the system account. Requires replacement if changed.
	AccountId *string `pulumi:"accountId"`
	// The ID of the entity. Requires replacement if changed.
	EntityId *string `pulumi:"entityId"`
	// The region of the team. Requires replacement if changed. ; must be one of ["us", "eu", "au", "*"]
	EntityRegion *string `pulumi:"entityRegion"`
	// The type of entity. Requires replacement if changed. ; must be one of ["API Products", "Application Auth Strategies", "Audit Logs", "Control Planes", "DCR Providers", "Identity", "Mesh Control Planes", "Networks", "Portals", "Service Hub"]
	EntityTypeName *string `pulumi:"entityTypeName"`
	// The desired role. Requires replacement if changed. ; must be one of ["Admin", "Appearance Maintainer", "Application Registration", "Certificate Admin", "Cloud Gateway Cluster Admin", "Cloud Gateway Cluster Viewer", "Consumer Admin", "Creator", "Deployer", "Discovery Admin", "Discovery Viewer", "Gateway Service Admin", "Integration Admin", "Integration Viewer", "Key Admin", "Maintainer", "Network Admin", "Network Creator", "Network Viewer", "Plugin Admin", "Plugins Admin", "Product Publisher", "Publisher", "Route Admin", "SNI Admin", "Service Admin", "Service Creator", "Service Viewer", "Upstream Admin", "Vault Admin", "Viewer"]
	RoleName *string `pulumi:"roleName"`
}

type SystemAccountRoleState struct {
	// ID of the system account. Requires replacement if changed.
	AccountId pulumi.StringPtrInput
	// The ID of the entity. Requires replacement if changed.
	EntityId pulumi.StringPtrInput
	// The region of the team. Requires replacement if changed. ; must be one of ["us", "eu", "au", "*"]
	EntityRegion pulumi.StringPtrInput
	// The type of entity. Requires replacement if changed. ; must be one of ["API Products", "Application Auth Strategies", "Audit Logs", "Control Planes", "DCR Providers", "Identity", "Mesh Control Planes", "Networks", "Portals", "Service Hub"]
	EntityTypeName pulumi.StringPtrInput
	// The desired role. Requires replacement if changed. ; must be one of ["Admin", "Appearance Maintainer", "Application Registration", "Certificate Admin", "Cloud Gateway Cluster Admin", "Cloud Gateway Cluster Viewer", "Consumer Admin", "Creator", "Deployer", "Discovery Admin", "Discovery Viewer", "Gateway Service Admin", "Integration Admin", "Integration Viewer", "Key Admin", "Maintainer", "Network Admin", "Network Creator", "Network Viewer", "Plugin Admin", "Plugins Admin", "Product Publisher", "Publisher", "Route Admin", "SNI Admin", "Service Admin", "Service Creator", "Service Viewer", "Upstream Admin", "Vault Admin", "Viewer"]
	RoleName pulumi.StringPtrInput
}

func (SystemAccountRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAccountRoleState)(nil)).Elem()
}

type systemAccountRoleArgs struct {
	// ID of the system account. Requires replacement if changed.
	AccountId string `pulumi:"accountId"`
	// The ID of the entity. Requires replacement if changed.
	EntityId *string `pulumi:"entityId"`
	// The region of the team. Requires replacement if changed. ; must be one of ["us", "eu", "au", "*"]
	EntityRegion *string `pulumi:"entityRegion"`
	// The type of entity. Requires replacement if changed. ; must be one of ["API Products", "Application Auth Strategies", "Audit Logs", "Control Planes", "DCR Providers", "Identity", "Mesh Control Planes", "Networks", "Portals", "Service Hub"]
	EntityTypeName *string `pulumi:"entityTypeName"`
	// The desired role. Requires replacement if changed. ; must be one of ["Admin", "Appearance Maintainer", "Application Registration", "Certificate Admin", "Cloud Gateway Cluster Admin", "Cloud Gateway Cluster Viewer", "Consumer Admin", "Creator", "Deployer", "Discovery Admin", "Discovery Viewer", "Gateway Service Admin", "Integration Admin", "Integration Viewer", "Key Admin", "Maintainer", "Network Admin", "Network Creator", "Network Viewer", "Plugin Admin", "Plugins Admin", "Product Publisher", "Publisher", "Route Admin", "SNI Admin", "Service Admin", "Service Creator", "Service Viewer", "Upstream Admin", "Vault Admin", "Viewer"]
	RoleName *string `pulumi:"roleName"`
}

// The set of arguments for constructing a SystemAccountRole resource.
type SystemAccountRoleArgs struct {
	// ID of the system account. Requires replacement if changed.
	AccountId pulumi.StringInput
	// The ID of the entity. Requires replacement if changed.
	EntityId pulumi.StringPtrInput
	// The region of the team. Requires replacement if changed. ; must be one of ["us", "eu", "au", "*"]
	EntityRegion pulumi.StringPtrInput
	// The type of entity. Requires replacement if changed. ; must be one of ["API Products", "Application Auth Strategies", "Audit Logs", "Control Planes", "DCR Providers", "Identity", "Mesh Control Planes", "Networks", "Portals", "Service Hub"]
	EntityTypeName pulumi.StringPtrInput
	// The desired role. Requires replacement if changed. ; must be one of ["Admin", "Appearance Maintainer", "Application Registration", "Certificate Admin", "Cloud Gateway Cluster Admin", "Cloud Gateway Cluster Viewer", "Consumer Admin", "Creator", "Deployer", "Discovery Admin", "Discovery Viewer", "Gateway Service Admin", "Integration Admin", "Integration Viewer", "Key Admin", "Maintainer", "Network Admin", "Network Creator", "Network Viewer", "Plugin Admin", "Plugins Admin", "Product Publisher", "Publisher", "Route Admin", "SNI Admin", "Service Admin", "Service Creator", "Service Viewer", "Upstream Admin", "Vault Admin", "Viewer"]
	RoleName pulumi.StringPtrInput
}

func (SystemAccountRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAccountRoleArgs)(nil)).Elem()
}

type SystemAccountRoleInput interface {
	pulumi.Input

	ToSystemAccountRoleOutput() SystemAccountRoleOutput
	ToSystemAccountRoleOutputWithContext(ctx context.Context) SystemAccountRoleOutput
}

func (*SystemAccountRole) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAccountRole)(nil)).Elem()
}

func (i *SystemAccountRole) ToSystemAccountRoleOutput() SystemAccountRoleOutput {
	return i.ToSystemAccountRoleOutputWithContext(context.Background())
}

func (i *SystemAccountRole) ToSystemAccountRoleOutputWithContext(ctx context.Context) SystemAccountRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAccountRoleOutput)
}

// SystemAccountRoleArrayInput is an input type that accepts SystemAccountRoleArray and SystemAccountRoleArrayOutput values.
// You can construct a concrete instance of `SystemAccountRoleArrayInput` via:
//
//	SystemAccountRoleArray{ SystemAccountRoleArgs{...} }
type SystemAccountRoleArrayInput interface {
	pulumi.Input

	ToSystemAccountRoleArrayOutput() SystemAccountRoleArrayOutput
	ToSystemAccountRoleArrayOutputWithContext(context.Context) SystemAccountRoleArrayOutput
}

type SystemAccountRoleArray []SystemAccountRoleInput

func (SystemAccountRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAccountRole)(nil)).Elem()
}

func (i SystemAccountRoleArray) ToSystemAccountRoleArrayOutput() SystemAccountRoleArrayOutput {
	return i.ToSystemAccountRoleArrayOutputWithContext(context.Background())
}

func (i SystemAccountRoleArray) ToSystemAccountRoleArrayOutputWithContext(ctx context.Context) SystemAccountRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAccountRoleArrayOutput)
}

// SystemAccountRoleMapInput is an input type that accepts SystemAccountRoleMap and SystemAccountRoleMapOutput values.
// You can construct a concrete instance of `SystemAccountRoleMapInput` via:
//
//	SystemAccountRoleMap{ "key": SystemAccountRoleArgs{...} }
type SystemAccountRoleMapInput interface {
	pulumi.Input

	ToSystemAccountRoleMapOutput() SystemAccountRoleMapOutput
	ToSystemAccountRoleMapOutputWithContext(context.Context) SystemAccountRoleMapOutput
}

type SystemAccountRoleMap map[string]SystemAccountRoleInput

func (SystemAccountRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAccountRole)(nil)).Elem()
}

func (i SystemAccountRoleMap) ToSystemAccountRoleMapOutput() SystemAccountRoleMapOutput {
	return i.ToSystemAccountRoleMapOutputWithContext(context.Background())
}

func (i SystemAccountRoleMap) ToSystemAccountRoleMapOutputWithContext(ctx context.Context) SystemAccountRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAccountRoleMapOutput)
}

type SystemAccountRoleOutput struct{ *pulumi.OutputState }

func (SystemAccountRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAccountRole)(nil)).Elem()
}

func (o SystemAccountRoleOutput) ToSystemAccountRoleOutput() SystemAccountRoleOutput {
	return o
}

func (o SystemAccountRoleOutput) ToSystemAccountRoleOutputWithContext(ctx context.Context) SystemAccountRoleOutput {
	return o
}

// ID of the system account. Requires replacement if changed.
func (o SystemAccountRoleOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAccountRole) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// The ID of the entity. Requires replacement if changed.
func (o SystemAccountRoleOutput) EntityId() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAccountRole) pulumi.StringOutput { return v.EntityId }).(pulumi.StringOutput)
}

// The region of the team. Requires replacement if changed. ; must be one of ["us", "eu", "au", "*"]
func (o SystemAccountRoleOutput) EntityRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAccountRole) pulumi.StringOutput { return v.EntityRegion }).(pulumi.StringOutput)
}

// The type of entity. Requires replacement if changed. ; must be one of ["API Products", "Application Auth Strategies", "Audit Logs", "Control Planes", "DCR Providers", "Identity", "Mesh Control Planes", "Networks", "Portals", "Service Hub"]
func (o SystemAccountRoleOutput) EntityTypeName() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAccountRole) pulumi.StringOutput { return v.EntityTypeName }).(pulumi.StringOutput)
}

// The desired role. Requires replacement if changed. ; must be one of ["Admin", "Appearance Maintainer", "Application Registration", "Certificate Admin", "Cloud Gateway Cluster Admin", "Cloud Gateway Cluster Viewer", "Consumer Admin", "Creator", "Deployer", "Discovery Admin", "Discovery Viewer", "Gateway Service Admin", "Integration Admin", "Integration Viewer", "Key Admin", "Maintainer", "Network Admin", "Network Creator", "Network Viewer", "Plugin Admin", "Plugins Admin", "Product Publisher", "Publisher", "Route Admin", "SNI Admin", "Service Admin", "Service Creator", "Service Viewer", "Upstream Admin", "Vault Admin", "Viewer"]
func (o SystemAccountRoleOutput) RoleName() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAccountRole) pulumi.StringOutput { return v.RoleName }).(pulumi.StringOutput)
}

type SystemAccountRoleArrayOutput struct{ *pulumi.OutputState }

func (SystemAccountRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAccountRole)(nil)).Elem()
}

func (o SystemAccountRoleArrayOutput) ToSystemAccountRoleArrayOutput() SystemAccountRoleArrayOutput {
	return o
}

func (o SystemAccountRoleArrayOutput) ToSystemAccountRoleArrayOutputWithContext(ctx context.Context) SystemAccountRoleArrayOutput {
	return o
}

func (o SystemAccountRoleArrayOutput) Index(i pulumi.IntInput) SystemAccountRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemAccountRole {
		return vs[0].([]*SystemAccountRole)[vs[1].(int)]
	}).(SystemAccountRoleOutput)
}

type SystemAccountRoleMapOutput struct{ *pulumi.OutputState }

func (SystemAccountRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAccountRole)(nil)).Elem()
}

func (o SystemAccountRoleMapOutput) ToSystemAccountRoleMapOutput() SystemAccountRoleMapOutput {
	return o
}

func (o SystemAccountRoleMapOutput) ToSystemAccountRoleMapOutputWithContext(ctx context.Context) SystemAccountRoleMapOutput {
	return o
}

func (o SystemAccountRoleMapOutput) MapIndex(k pulumi.StringInput) SystemAccountRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemAccountRole {
		return vs[0].(map[string]*SystemAccountRole)[vs[1].(string)]
	}).(SystemAccountRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAccountRoleInput)(nil)).Elem(), &SystemAccountRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAccountRoleArrayInput)(nil)).Elem(), SystemAccountRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAccountRoleMapInput)(nil)).Elem(), SystemAccountRoleMap{})
	pulumi.RegisterOutputType(SystemAccountRoleOutput{})
	pulumi.RegisterOutputType(SystemAccountRoleArrayOutput{})
	pulumi.RegisterOutputType(SystemAccountRoleMapOutput{})
}