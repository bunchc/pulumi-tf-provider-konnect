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

// CloudGatewayNetwork Resource
type CloudGatewayNetwork struct {
	pulumi.CustomResourceState

	// List of availability zones that the network is attached to. Requires replacement if changed.
	AvailabilityZones pulumi.StringArrayOutput `pulumi:"availabilityZones"`
	// CIDR block configuration for the network. Requires replacement if changed.
	CidrBlock pulumi.StringOutput `pulumi:"cidrBlock"`
	// Requires replacement if changed.
	CloudGatewayProviderAccountId pulumi.StringOutput `pulumi:"cloudGatewayProviderAccountId"`
	// The number of configurations that reference this network.
	ConfigurationReferenceCount pulumi.IntOutput `pulumi:"configurationReferenceCount"`
	// An RFC-3339 timestamp representation of network creation date.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Whether DDOS protection is enabled for the network. Requires replacement if changed.
	DdosProtection pulumi.BoolOutput `pulumi:"ddosProtection"`
	// Whether the network is a default network or not. Default networks are Networks that are created
	// automatically by Konnect when an organization is linked to a provider account.
	Default pulumi.BoolOutput `pulumi:"default"`
	// Monotonically-increasing version count of the network, to indicate the order of updates to the network.
	EntityVersion pulumi.IntOutput `pulumi:"entityVersion"`
	// Firewall configuration for a network.
	Firewall CloudGatewayNetworkFirewallOutput `pulumi:"firewall"`
	// Human-readable name of the network.
	Name pulumi.StringOutput `pulumi:"name"`
	// Metadata describing attributes returned by cloud-provider for the network.
	ProviderMetadata CloudGatewayNetworkProviderMetadataOutput `pulumi:"providerMetadata"`
	// Region ID for cloud provider region. Requires replacement if changed.
	Region pulumi.StringOutput `pulumi:"region"`
	// State of the network. must be one of ["created", "initializing", "offline", "ready", "terminating", "terminated"]
	State pulumi.StringOutput `pulumi:"state"`
	// The number of transit gateways attached to this network.
	TransitGatewayCount pulumi.IntOutput `pulumi:"transitGatewayCount"`
	// An RFC-3339 timestamp representation of network update date.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewCloudGatewayNetwork registers a new resource with the given unique name, arguments, and options.
func NewCloudGatewayNetwork(ctx *pulumi.Context,
	name string, args *CloudGatewayNetworkArgs, opts ...pulumi.ResourceOption) (*CloudGatewayNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityZones == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityZones'")
	}
	if args.CidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'CidrBlock'")
	}
	if args.CloudGatewayProviderAccountId == nil {
		return nil, errors.New("invalid value for required argument 'CloudGatewayProviderAccountId'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CloudGatewayNetwork
	err := ctx.RegisterResource("konnect:index/cloudGatewayNetwork:CloudGatewayNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudGatewayNetwork gets an existing CloudGatewayNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudGatewayNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudGatewayNetworkState, opts ...pulumi.ResourceOption) (*CloudGatewayNetwork, error) {
	var resource CloudGatewayNetwork
	err := ctx.ReadResource("konnect:index/cloudGatewayNetwork:CloudGatewayNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudGatewayNetwork resources.
type cloudGatewayNetworkState struct {
	// List of availability zones that the network is attached to. Requires replacement if changed.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// CIDR block configuration for the network. Requires replacement if changed.
	CidrBlock *string `pulumi:"cidrBlock"`
	// Requires replacement if changed.
	CloudGatewayProviderAccountId *string `pulumi:"cloudGatewayProviderAccountId"`
	// The number of configurations that reference this network.
	ConfigurationReferenceCount *int `pulumi:"configurationReferenceCount"`
	// An RFC-3339 timestamp representation of network creation date.
	CreatedAt *string `pulumi:"createdAt"`
	// Whether DDOS protection is enabled for the network. Requires replacement if changed.
	DdosProtection *bool `pulumi:"ddosProtection"`
	// Whether the network is a default network or not. Default networks are Networks that are created
	// automatically by Konnect when an organization is linked to a provider account.
	Default *bool `pulumi:"default"`
	// Monotonically-increasing version count of the network, to indicate the order of updates to the network.
	EntityVersion *int `pulumi:"entityVersion"`
	// Firewall configuration for a network.
	Firewall *CloudGatewayNetworkFirewall `pulumi:"firewall"`
	// Human-readable name of the network.
	Name *string `pulumi:"name"`
	// Metadata describing attributes returned by cloud-provider for the network.
	ProviderMetadata *CloudGatewayNetworkProviderMetadata `pulumi:"providerMetadata"`
	// Region ID for cloud provider region. Requires replacement if changed.
	Region *string `pulumi:"region"`
	// State of the network. must be one of ["created", "initializing", "offline", "ready", "terminating", "terminated"]
	State *string `pulumi:"state"`
	// The number of transit gateways attached to this network.
	TransitGatewayCount *int `pulumi:"transitGatewayCount"`
	// An RFC-3339 timestamp representation of network update date.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type CloudGatewayNetworkState struct {
	// List of availability zones that the network is attached to. Requires replacement if changed.
	AvailabilityZones pulumi.StringArrayInput
	// CIDR block configuration for the network. Requires replacement if changed.
	CidrBlock pulumi.StringPtrInput
	// Requires replacement if changed.
	CloudGatewayProviderAccountId pulumi.StringPtrInput
	// The number of configurations that reference this network.
	ConfigurationReferenceCount pulumi.IntPtrInput
	// An RFC-3339 timestamp representation of network creation date.
	CreatedAt pulumi.StringPtrInput
	// Whether DDOS protection is enabled for the network. Requires replacement if changed.
	DdosProtection pulumi.BoolPtrInput
	// Whether the network is a default network or not. Default networks are Networks that are created
	// automatically by Konnect when an organization is linked to a provider account.
	Default pulumi.BoolPtrInput
	// Monotonically-increasing version count of the network, to indicate the order of updates to the network.
	EntityVersion pulumi.IntPtrInput
	// Firewall configuration for a network.
	Firewall CloudGatewayNetworkFirewallPtrInput
	// Human-readable name of the network.
	Name pulumi.StringPtrInput
	// Metadata describing attributes returned by cloud-provider for the network.
	ProviderMetadata CloudGatewayNetworkProviderMetadataPtrInput
	// Region ID for cloud provider region. Requires replacement if changed.
	Region pulumi.StringPtrInput
	// State of the network. must be one of ["created", "initializing", "offline", "ready", "terminating", "terminated"]
	State pulumi.StringPtrInput
	// The number of transit gateways attached to this network.
	TransitGatewayCount pulumi.IntPtrInput
	// An RFC-3339 timestamp representation of network update date.
	UpdatedAt pulumi.StringPtrInput
}

func (CloudGatewayNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudGatewayNetworkState)(nil)).Elem()
}

type cloudGatewayNetworkArgs struct {
	// List of availability zones that the network is attached to. Requires replacement if changed.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// CIDR block configuration for the network. Requires replacement if changed.
	CidrBlock string `pulumi:"cidrBlock"`
	// Requires replacement if changed.
	CloudGatewayProviderAccountId string `pulumi:"cloudGatewayProviderAccountId"`
	// Whether DDOS protection is enabled for the network. Requires replacement if changed.
	DdosProtection *bool `pulumi:"ddosProtection"`
	// Firewall configuration for a network.
	Firewall *CloudGatewayNetworkFirewall `pulumi:"firewall"`
	// Human-readable name of the network.
	Name *string `pulumi:"name"`
	// Region ID for cloud provider region. Requires replacement if changed.
	Region string `pulumi:"region"`
}

// The set of arguments for constructing a CloudGatewayNetwork resource.
type CloudGatewayNetworkArgs struct {
	// List of availability zones that the network is attached to. Requires replacement if changed.
	AvailabilityZones pulumi.StringArrayInput
	// CIDR block configuration for the network. Requires replacement if changed.
	CidrBlock pulumi.StringInput
	// Requires replacement if changed.
	CloudGatewayProviderAccountId pulumi.StringInput
	// Whether DDOS protection is enabled for the network. Requires replacement if changed.
	DdosProtection pulumi.BoolPtrInput
	// Firewall configuration for a network.
	Firewall CloudGatewayNetworkFirewallPtrInput
	// Human-readable name of the network.
	Name pulumi.StringPtrInput
	// Region ID for cloud provider region. Requires replacement if changed.
	Region pulumi.StringInput
}

func (CloudGatewayNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudGatewayNetworkArgs)(nil)).Elem()
}

type CloudGatewayNetworkInput interface {
	pulumi.Input

	ToCloudGatewayNetworkOutput() CloudGatewayNetworkOutput
	ToCloudGatewayNetworkOutputWithContext(ctx context.Context) CloudGatewayNetworkOutput
}

func (*CloudGatewayNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudGatewayNetwork)(nil)).Elem()
}

func (i *CloudGatewayNetwork) ToCloudGatewayNetworkOutput() CloudGatewayNetworkOutput {
	return i.ToCloudGatewayNetworkOutputWithContext(context.Background())
}

func (i *CloudGatewayNetwork) ToCloudGatewayNetworkOutputWithContext(ctx context.Context) CloudGatewayNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudGatewayNetworkOutput)
}

// CloudGatewayNetworkArrayInput is an input type that accepts CloudGatewayNetworkArray and CloudGatewayNetworkArrayOutput values.
// You can construct a concrete instance of `CloudGatewayNetworkArrayInput` via:
//
//	CloudGatewayNetworkArray{ CloudGatewayNetworkArgs{...} }
type CloudGatewayNetworkArrayInput interface {
	pulumi.Input

	ToCloudGatewayNetworkArrayOutput() CloudGatewayNetworkArrayOutput
	ToCloudGatewayNetworkArrayOutputWithContext(context.Context) CloudGatewayNetworkArrayOutput
}

type CloudGatewayNetworkArray []CloudGatewayNetworkInput

func (CloudGatewayNetworkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudGatewayNetwork)(nil)).Elem()
}

func (i CloudGatewayNetworkArray) ToCloudGatewayNetworkArrayOutput() CloudGatewayNetworkArrayOutput {
	return i.ToCloudGatewayNetworkArrayOutputWithContext(context.Background())
}

func (i CloudGatewayNetworkArray) ToCloudGatewayNetworkArrayOutputWithContext(ctx context.Context) CloudGatewayNetworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudGatewayNetworkArrayOutput)
}

// CloudGatewayNetworkMapInput is an input type that accepts CloudGatewayNetworkMap and CloudGatewayNetworkMapOutput values.
// You can construct a concrete instance of `CloudGatewayNetworkMapInput` via:
//
//	CloudGatewayNetworkMap{ "key": CloudGatewayNetworkArgs{...} }
type CloudGatewayNetworkMapInput interface {
	pulumi.Input

	ToCloudGatewayNetworkMapOutput() CloudGatewayNetworkMapOutput
	ToCloudGatewayNetworkMapOutputWithContext(context.Context) CloudGatewayNetworkMapOutput
}

type CloudGatewayNetworkMap map[string]CloudGatewayNetworkInput

func (CloudGatewayNetworkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudGatewayNetwork)(nil)).Elem()
}

func (i CloudGatewayNetworkMap) ToCloudGatewayNetworkMapOutput() CloudGatewayNetworkMapOutput {
	return i.ToCloudGatewayNetworkMapOutputWithContext(context.Background())
}

func (i CloudGatewayNetworkMap) ToCloudGatewayNetworkMapOutputWithContext(ctx context.Context) CloudGatewayNetworkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudGatewayNetworkMapOutput)
}

type CloudGatewayNetworkOutput struct{ *pulumi.OutputState }

func (CloudGatewayNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudGatewayNetwork)(nil)).Elem()
}

func (o CloudGatewayNetworkOutput) ToCloudGatewayNetworkOutput() CloudGatewayNetworkOutput {
	return o
}

func (o CloudGatewayNetworkOutput) ToCloudGatewayNetworkOutputWithContext(ctx context.Context) CloudGatewayNetworkOutput {
	return o
}

// List of availability zones that the network is attached to. Requires replacement if changed.
func (o CloudGatewayNetworkOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CloudGatewayNetwork) pulumi.StringArrayOutput { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

// CIDR block configuration for the network. Requires replacement if changed.
func (o CloudGatewayNetworkOutput) CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudGatewayNetwork) pulumi.StringOutput { return v.CidrBlock }).(pulumi.StringOutput)
}

// Requires replacement if changed.
func (o CloudGatewayNetworkOutput) CloudGatewayProviderAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudGatewayNetwork) pulumi.StringOutput { return v.CloudGatewayProviderAccountId }).(pulumi.StringOutput)
}

// The number of configurations that reference this network.
func (o CloudGatewayNetworkOutput) ConfigurationReferenceCount() pulumi.IntOutput {
	return o.ApplyT(func(v *CloudGatewayNetwork) pulumi.IntOutput { return v.ConfigurationReferenceCount }).(pulumi.IntOutput)
}

// An RFC-3339 timestamp representation of network creation date.
func (o CloudGatewayNetworkOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudGatewayNetwork) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Whether DDOS protection is enabled for the network. Requires replacement if changed.
func (o CloudGatewayNetworkOutput) DdosProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v *CloudGatewayNetwork) pulumi.BoolOutput { return v.DdosProtection }).(pulumi.BoolOutput)
}

// Whether the network is a default network or not. Default networks are Networks that are created
// automatically by Konnect when an organization is linked to a provider account.
func (o CloudGatewayNetworkOutput) Default() pulumi.BoolOutput {
	return o.ApplyT(func(v *CloudGatewayNetwork) pulumi.BoolOutput { return v.Default }).(pulumi.BoolOutput)
}

// Monotonically-increasing version count of the network, to indicate the order of updates to the network.
func (o CloudGatewayNetworkOutput) EntityVersion() pulumi.IntOutput {
	return o.ApplyT(func(v *CloudGatewayNetwork) pulumi.IntOutput { return v.EntityVersion }).(pulumi.IntOutput)
}

// Firewall configuration for a network.
func (o CloudGatewayNetworkOutput) Firewall() CloudGatewayNetworkFirewallOutput {
	return o.ApplyT(func(v *CloudGatewayNetwork) CloudGatewayNetworkFirewallOutput { return v.Firewall }).(CloudGatewayNetworkFirewallOutput)
}

// Human-readable name of the network.
func (o CloudGatewayNetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudGatewayNetwork) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Metadata describing attributes returned by cloud-provider for the network.
func (o CloudGatewayNetworkOutput) ProviderMetadata() CloudGatewayNetworkProviderMetadataOutput {
	return o.ApplyT(func(v *CloudGatewayNetwork) CloudGatewayNetworkProviderMetadataOutput { return v.ProviderMetadata }).(CloudGatewayNetworkProviderMetadataOutput)
}

// Region ID for cloud provider region. Requires replacement if changed.
func (o CloudGatewayNetworkOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudGatewayNetwork) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// State of the network. must be one of ["created", "initializing", "offline", "ready", "terminating", "terminated"]
func (o CloudGatewayNetworkOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudGatewayNetwork) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The number of transit gateways attached to this network.
func (o CloudGatewayNetworkOutput) TransitGatewayCount() pulumi.IntOutput {
	return o.ApplyT(func(v *CloudGatewayNetwork) pulumi.IntOutput { return v.TransitGatewayCount }).(pulumi.IntOutput)
}

// An RFC-3339 timestamp representation of network update date.
func (o CloudGatewayNetworkOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudGatewayNetwork) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type CloudGatewayNetworkArrayOutput struct{ *pulumi.OutputState }

func (CloudGatewayNetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudGatewayNetwork)(nil)).Elem()
}

func (o CloudGatewayNetworkArrayOutput) ToCloudGatewayNetworkArrayOutput() CloudGatewayNetworkArrayOutput {
	return o
}

func (o CloudGatewayNetworkArrayOutput) ToCloudGatewayNetworkArrayOutputWithContext(ctx context.Context) CloudGatewayNetworkArrayOutput {
	return o
}

func (o CloudGatewayNetworkArrayOutput) Index(i pulumi.IntInput) CloudGatewayNetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudGatewayNetwork {
		return vs[0].([]*CloudGatewayNetwork)[vs[1].(int)]
	}).(CloudGatewayNetworkOutput)
}

type CloudGatewayNetworkMapOutput struct{ *pulumi.OutputState }

func (CloudGatewayNetworkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudGatewayNetwork)(nil)).Elem()
}

func (o CloudGatewayNetworkMapOutput) ToCloudGatewayNetworkMapOutput() CloudGatewayNetworkMapOutput {
	return o
}

func (o CloudGatewayNetworkMapOutput) ToCloudGatewayNetworkMapOutputWithContext(ctx context.Context) CloudGatewayNetworkMapOutput {
	return o
}

func (o CloudGatewayNetworkMapOutput) MapIndex(k pulumi.StringInput) CloudGatewayNetworkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudGatewayNetwork {
		return vs[0].(map[string]*CloudGatewayNetwork)[vs[1].(string)]
	}).(CloudGatewayNetworkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudGatewayNetworkInput)(nil)).Elem(), &CloudGatewayNetwork{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudGatewayNetworkArrayInput)(nil)).Elem(), CloudGatewayNetworkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudGatewayNetworkMapInput)(nil)).Elem(), CloudGatewayNetworkMap{})
	pulumi.RegisterOutputType(CloudGatewayNetworkOutput{})
	pulumi.RegisterOutputType(CloudGatewayNetworkArrayOutput{})
	pulumi.RegisterOutputType(CloudGatewayNetworkMapOutput{})
}