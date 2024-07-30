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

// PortalProductVersion Resource
type PortalProductVersion struct {
	pulumi.CustomResourceState

	// Whether the application registration on this portal for the api product version is enabled
	ApplicationRegistrationEnabled pulumi.BoolOutput `pulumi:"applicationRegistrationEnabled"`
	// A list of authentication strategies
	AuthStrategies PortalProductVersionAuthStrategyArrayOutput `pulumi:"authStrategies"`
	// A list of authentication strategy IDs
	AuthStrategyIds pulumi.StringArrayOutput `pulumi:"authStrategyIds"`
	// Whether the application registration auto approval on this portal for the api product version is enabled
	AutoApproveRegistration pulumi.BoolOutput `pulumi:"autoApproveRegistration"`
	// An ISO-8601 timestamp representation of entity creation date.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Whether the api product version on the portal is deprecated
	Deprecated pulumi.BoolOutput `pulumi:"deprecated"`
	// Whether to notify developers who are affected by this change
	NotifyDevelopers pulumi.BoolPtrOutput `pulumi:"notifyDevelopers"`
	// portal identifier
	PortalId pulumi.StringOutput `pulumi:"portalId"`
	// API product version identifier
	ProductVersionId pulumi.StringOutput `pulumi:"productVersionId"`
	// Publication status of the API product version on the portal. must be one of ["published", "unpublished"]
	PublishStatus pulumi.StringOutput `pulumi:"publishStatus"`
	// An ISO-8601 timestamp representation of entity update date.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewPortalProductVersion registers a new resource with the given unique name, arguments, and options.
func NewPortalProductVersion(ctx *pulumi.Context,
	name string, args *PortalProductVersionArgs, opts ...pulumi.ResourceOption) (*PortalProductVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationRegistrationEnabled == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationRegistrationEnabled'")
	}
	if args.AuthStrategyIds == nil {
		return nil, errors.New("invalid value for required argument 'AuthStrategyIds'")
	}
	if args.AutoApproveRegistration == nil {
		return nil, errors.New("invalid value for required argument 'AutoApproveRegistration'")
	}
	if args.Deprecated == nil {
		return nil, errors.New("invalid value for required argument 'Deprecated'")
	}
	if args.PortalId == nil {
		return nil, errors.New("invalid value for required argument 'PortalId'")
	}
	if args.ProductVersionId == nil {
		return nil, errors.New("invalid value for required argument 'ProductVersionId'")
	}
	if args.PublishStatus == nil {
		return nil, errors.New("invalid value for required argument 'PublishStatus'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PortalProductVersion
	err := ctx.RegisterResource("konnect:index/portalProductVersion:PortalProductVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPortalProductVersion gets an existing PortalProductVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPortalProductVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PortalProductVersionState, opts ...pulumi.ResourceOption) (*PortalProductVersion, error) {
	var resource PortalProductVersion
	err := ctx.ReadResource("konnect:index/portalProductVersion:PortalProductVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PortalProductVersion resources.
type portalProductVersionState struct {
	// Whether the application registration on this portal for the api product version is enabled
	ApplicationRegistrationEnabled *bool `pulumi:"applicationRegistrationEnabled"`
	// A list of authentication strategies
	AuthStrategies []PortalProductVersionAuthStrategy `pulumi:"authStrategies"`
	// A list of authentication strategy IDs
	AuthStrategyIds []string `pulumi:"authStrategyIds"`
	// Whether the application registration auto approval on this portal for the api product version is enabled
	AutoApproveRegistration *bool `pulumi:"autoApproveRegistration"`
	// An ISO-8601 timestamp representation of entity creation date.
	CreatedAt *string `pulumi:"createdAt"`
	// Whether the api product version on the portal is deprecated
	Deprecated *bool `pulumi:"deprecated"`
	// Whether to notify developers who are affected by this change
	NotifyDevelopers *bool `pulumi:"notifyDevelopers"`
	// portal identifier
	PortalId *string `pulumi:"portalId"`
	// API product version identifier
	ProductVersionId *string `pulumi:"productVersionId"`
	// Publication status of the API product version on the portal. must be one of ["published", "unpublished"]
	PublishStatus *string `pulumi:"publishStatus"`
	// An ISO-8601 timestamp representation of entity update date.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type PortalProductVersionState struct {
	// Whether the application registration on this portal for the api product version is enabled
	ApplicationRegistrationEnabled pulumi.BoolPtrInput
	// A list of authentication strategies
	AuthStrategies PortalProductVersionAuthStrategyArrayInput
	// A list of authentication strategy IDs
	AuthStrategyIds pulumi.StringArrayInput
	// Whether the application registration auto approval on this portal for the api product version is enabled
	AutoApproveRegistration pulumi.BoolPtrInput
	// An ISO-8601 timestamp representation of entity creation date.
	CreatedAt pulumi.StringPtrInput
	// Whether the api product version on the portal is deprecated
	Deprecated pulumi.BoolPtrInput
	// Whether to notify developers who are affected by this change
	NotifyDevelopers pulumi.BoolPtrInput
	// portal identifier
	PortalId pulumi.StringPtrInput
	// API product version identifier
	ProductVersionId pulumi.StringPtrInput
	// Publication status of the API product version on the portal. must be one of ["published", "unpublished"]
	PublishStatus pulumi.StringPtrInput
	// An ISO-8601 timestamp representation of entity update date.
	UpdatedAt pulumi.StringPtrInput
}

func (PortalProductVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*portalProductVersionState)(nil)).Elem()
}

type portalProductVersionArgs struct {
	// Whether the application registration on this portal for the api product version is enabled
	ApplicationRegistrationEnabled bool `pulumi:"applicationRegistrationEnabled"`
	// A list of authentication strategy IDs
	AuthStrategyIds []string `pulumi:"authStrategyIds"`
	// Whether the application registration auto approval on this portal for the api product version is enabled
	AutoApproveRegistration bool `pulumi:"autoApproveRegistration"`
	// Whether the api product version on the portal is deprecated
	Deprecated bool `pulumi:"deprecated"`
	// Whether to notify developers who are affected by this change
	NotifyDevelopers *bool `pulumi:"notifyDevelopers"`
	// portal identifier
	PortalId string `pulumi:"portalId"`
	// API product version identifier
	ProductVersionId string `pulumi:"productVersionId"`
	// Publication status of the API product version on the portal. must be one of ["published", "unpublished"]
	PublishStatus string `pulumi:"publishStatus"`
}

// The set of arguments for constructing a PortalProductVersion resource.
type PortalProductVersionArgs struct {
	// Whether the application registration on this portal for the api product version is enabled
	ApplicationRegistrationEnabled pulumi.BoolInput
	// A list of authentication strategy IDs
	AuthStrategyIds pulumi.StringArrayInput
	// Whether the application registration auto approval on this portal for the api product version is enabled
	AutoApproveRegistration pulumi.BoolInput
	// Whether the api product version on the portal is deprecated
	Deprecated pulumi.BoolInput
	// Whether to notify developers who are affected by this change
	NotifyDevelopers pulumi.BoolPtrInput
	// portal identifier
	PortalId pulumi.StringInput
	// API product version identifier
	ProductVersionId pulumi.StringInput
	// Publication status of the API product version on the portal. must be one of ["published", "unpublished"]
	PublishStatus pulumi.StringInput
}

func (PortalProductVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*portalProductVersionArgs)(nil)).Elem()
}

type PortalProductVersionInput interface {
	pulumi.Input

	ToPortalProductVersionOutput() PortalProductVersionOutput
	ToPortalProductVersionOutputWithContext(ctx context.Context) PortalProductVersionOutput
}

func (*PortalProductVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**PortalProductVersion)(nil)).Elem()
}

func (i *PortalProductVersion) ToPortalProductVersionOutput() PortalProductVersionOutput {
	return i.ToPortalProductVersionOutputWithContext(context.Background())
}

func (i *PortalProductVersion) ToPortalProductVersionOutputWithContext(ctx context.Context) PortalProductVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortalProductVersionOutput)
}

// PortalProductVersionArrayInput is an input type that accepts PortalProductVersionArray and PortalProductVersionArrayOutput values.
// You can construct a concrete instance of `PortalProductVersionArrayInput` via:
//
//	PortalProductVersionArray{ PortalProductVersionArgs{...} }
type PortalProductVersionArrayInput interface {
	pulumi.Input

	ToPortalProductVersionArrayOutput() PortalProductVersionArrayOutput
	ToPortalProductVersionArrayOutputWithContext(context.Context) PortalProductVersionArrayOutput
}

type PortalProductVersionArray []PortalProductVersionInput

func (PortalProductVersionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PortalProductVersion)(nil)).Elem()
}

func (i PortalProductVersionArray) ToPortalProductVersionArrayOutput() PortalProductVersionArrayOutput {
	return i.ToPortalProductVersionArrayOutputWithContext(context.Background())
}

func (i PortalProductVersionArray) ToPortalProductVersionArrayOutputWithContext(ctx context.Context) PortalProductVersionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortalProductVersionArrayOutput)
}

// PortalProductVersionMapInput is an input type that accepts PortalProductVersionMap and PortalProductVersionMapOutput values.
// You can construct a concrete instance of `PortalProductVersionMapInput` via:
//
//	PortalProductVersionMap{ "key": PortalProductVersionArgs{...} }
type PortalProductVersionMapInput interface {
	pulumi.Input

	ToPortalProductVersionMapOutput() PortalProductVersionMapOutput
	ToPortalProductVersionMapOutputWithContext(context.Context) PortalProductVersionMapOutput
}

type PortalProductVersionMap map[string]PortalProductVersionInput

func (PortalProductVersionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PortalProductVersion)(nil)).Elem()
}

func (i PortalProductVersionMap) ToPortalProductVersionMapOutput() PortalProductVersionMapOutput {
	return i.ToPortalProductVersionMapOutputWithContext(context.Background())
}

func (i PortalProductVersionMap) ToPortalProductVersionMapOutputWithContext(ctx context.Context) PortalProductVersionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortalProductVersionMapOutput)
}

type PortalProductVersionOutput struct{ *pulumi.OutputState }

func (PortalProductVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PortalProductVersion)(nil)).Elem()
}

func (o PortalProductVersionOutput) ToPortalProductVersionOutput() PortalProductVersionOutput {
	return o
}

func (o PortalProductVersionOutput) ToPortalProductVersionOutputWithContext(ctx context.Context) PortalProductVersionOutput {
	return o
}

// Whether the application registration on this portal for the api product version is enabled
func (o PortalProductVersionOutput) ApplicationRegistrationEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *PortalProductVersion) pulumi.BoolOutput { return v.ApplicationRegistrationEnabled }).(pulumi.BoolOutput)
}

// A list of authentication strategies
func (o PortalProductVersionOutput) AuthStrategies() PortalProductVersionAuthStrategyArrayOutput {
	return o.ApplyT(func(v *PortalProductVersion) PortalProductVersionAuthStrategyArrayOutput { return v.AuthStrategies }).(PortalProductVersionAuthStrategyArrayOutput)
}

// A list of authentication strategy IDs
func (o PortalProductVersionOutput) AuthStrategyIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PortalProductVersion) pulumi.StringArrayOutput { return v.AuthStrategyIds }).(pulumi.StringArrayOutput)
}

// Whether the application registration auto approval on this portal for the api product version is enabled
func (o PortalProductVersionOutput) AutoApproveRegistration() pulumi.BoolOutput {
	return o.ApplyT(func(v *PortalProductVersion) pulumi.BoolOutput { return v.AutoApproveRegistration }).(pulumi.BoolOutput)
}

// An ISO-8601 timestamp representation of entity creation date.
func (o PortalProductVersionOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *PortalProductVersion) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Whether the api product version on the portal is deprecated
func (o PortalProductVersionOutput) Deprecated() pulumi.BoolOutput {
	return o.ApplyT(func(v *PortalProductVersion) pulumi.BoolOutput { return v.Deprecated }).(pulumi.BoolOutput)
}

// Whether to notify developers who are affected by this change
func (o PortalProductVersionOutput) NotifyDevelopers() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PortalProductVersion) pulumi.BoolPtrOutput { return v.NotifyDevelopers }).(pulumi.BoolPtrOutput)
}

// portal identifier
func (o PortalProductVersionOutput) PortalId() pulumi.StringOutput {
	return o.ApplyT(func(v *PortalProductVersion) pulumi.StringOutput { return v.PortalId }).(pulumi.StringOutput)
}

// API product version identifier
func (o PortalProductVersionOutput) ProductVersionId() pulumi.StringOutput {
	return o.ApplyT(func(v *PortalProductVersion) pulumi.StringOutput { return v.ProductVersionId }).(pulumi.StringOutput)
}

// Publication status of the API product version on the portal. must be one of ["published", "unpublished"]
func (o PortalProductVersionOutput) PublishStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *PortalProductVersion) pulumi.StringOutput { return v.PublishStatus }).(pulumi.StringOutput)
}

// An ISO-8601 timestamp representation of entity update date.
func (o PortalProductVersionOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *PortalProductVersion) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type PortalProductVersionArrayOutput struct{ *pulumi.OutputState }

func (PortalProductVersionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PortalProductVersion)(nil)).Elem()
}

func (o PortalProductVersionArrayOutput) ToPortalProductVersionArrayOutput() PortalProductVersionArrayOutput {
	return o
}

func (o PortalProductVersionArrayOutput) ToPortalProductVersionArrayOutputWithContext(ctx context.Context) PortalProductVersionArrayOutput {
	return o
}

func (o PortalProductVersionArrayOutput) Index(i pulumi.IntInput) PortalProductVersionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PortalProductVersion {
		return vs[0].([]*PortalProductVersion)[vs[1].(int)]
	}).(PortalProductVersionOutput)
}

type PortalProductVersionMapOutput struct{ *pulumi.OutputState }

func (PortalProductVersionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PortalProductVersion)(nil)).Elem()
}

func (o PortalProductVersionMapOutput) ToPortalProductVersionMapOutput() PortalProductVersionMapOutput {
	return o
}

func (o PortalProductVersionMapOutput) ToPortalProductVersionMapOutputWithContext(ctx context.Context) PortalProductVersionMapOutput {
	return o
}

func (o PortalProductVersionMapOutput) MapIndex(k pulumi.StringInput) PortalProductVersionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PortalProductVersion {
		return vs[0].(map[string]*PortalProductVersion)[vs[1].(string)]
	}).(PortalProductVersionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PortalProductVersionInput)(nil)).Elem(), &PortalProductVersion{})
	pulumi.RegisterInputType(reflect.TypeOf((*PortalProductVersionArrayInput)(nil)).Elem(), PortalProductVersionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PortalProductVersionMapInput)(nil)).Elem(), PortalProductVersionMap{})
	pulumi.RegisterOutputType(PortalProductVersionOutput{})
	pulumi.RegisterOutputType(PortalProductVersionArrayOutput{})
	pulumi.RegisterOutputType(PortalProductVersionMapOutput{})
}