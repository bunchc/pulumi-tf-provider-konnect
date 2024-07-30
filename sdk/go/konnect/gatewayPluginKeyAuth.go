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

// GatewayPluginKeyAuth Resource
type GatewayPluginKeyAuth struct {
	pulumi.CustomResourceState

	Config GatewayPluginKeyAuthConfigOutput `pulumi:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      GatewayPluginKeyAuthConsumerOutput      `pulumi:"consumer"`
	ConsumerGroup GatewayPluginKeyAuthConsumerGroupOutput `pulumi:"consumerGroup"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId pulumi.StringOutput `pulumi:"controlPlaneId"`
	// Unix epoch when the resource was created.
	CreatedAt pulumi.IntOutput `pulumi:"createdAt"`
	// Whether the plugin is applied.
	Enabled      pulumi.BoolOutput   `pulumi:"enabled"`
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols pulumi.StringArrayOutput `pulumi:"protocols"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route GatewayPluginKeyAuthRouteOutput `pulumi:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service GatewayPluginKeyAuthServiceOutput `pulumi:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Unix epoch when the resource was last updated.
	UpdatedAt pulumi.IntOutput `pulumi:"updatedAt"`
}

// NewGatewayPluginKeyAuth registers a new resource with the given unique name, arguments, and options.
func NewGatewayPluginKeyAuth(ctx *pulumi.Context,
	name string, args *GatewayPluginKeyAuthArgs, opts ...pulumi.ResourceOption) (*GatewayPluginKeyAuth, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ControlPlaneId == nil {
		return nil, errors.New("invalid value for required argument 'ControlPlaneId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GatewayPluginKeyAuth
	err := ctx.RegisterResource("konnect:index/gatewayPluginKeyAuth:GatewayPluginKeyAuth", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayPluginKeyAuth gets an existing GatewayPluginKeyAuth resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayPluginKeyAuth(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayPluginKeyAuthState, opts ...pulumi.ResourceOption) (*GatewayPluginKeyAuth, error) {
	var resource GatewayPluginKeyAuth
	err := ctx.ReadResource("konnect:index/gatewayPluginKeyAuth:GatewayPluginKeyAuth", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayPluginKeyAuth resources.
type gatewayPluginKeyAuthState struct {
	Config *GatewayPluginKeyAuthConfig `pulumi:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *GatewayPluginKeyAuthConsumer      `pulumi:"consumer"`
	ConsumerGroup *GatewayPluginKeyAuthConsumerGroup `pulumi:"consumerGroup"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId *string `pulumi:"controlPlaneId"`
	// Unix epoch when the resource was created.
	CreatedAt *int `pulumi:"createdAt"`
	// Whether the plugin is applied.
	Enabled      *bool   `pulumi:"enabled"`
	InstanceName *string `pulumi:"instanceName"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []string `pulumi:"protocols"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *GatewayPluginKeyAuthRoute `pulumi:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *GatewayPluginKeyAuthService `pulumi:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `pulumi:"tags"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int `pulumi:"updatedAt"`
}

type GatewayPluginKeyAuthState struct {
	Config GatewayPluginKeyAuthConfigPtrInput
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      GatewayPluginKeyAuthConsumerPtrInput
	ConsumerGroup GatewayPluginKeyAuthConsumerGroupPtrInput
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId pulumi.StringPtrInput
	// Unix epoch when the resource was created.
	CreatedAt pulumi.IntPtrInput
	// Whether the plugin is applied.
	Enabled      pulumi.BoolPtrInput
	InstanceName pulumi.StringPtrInput
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols pulumi.StringArrayInput
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route GatewayPluginKeyAuthRoutePtrInput
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service GatewayPluginKeyAuthServicePtrInput
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags pulumi.StringArrayInput
	// Unix epoch when the resource was last updated.
	UpdatedAt pulumi.IntPtrInput
}

func (GatewayPluginKeyAuthState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayPluginKeyAuthState)(nil)).Elem()
}

type gatewayPluginKeyAuthArgs struct {
	Config *GatewayPluginKeyAuthConfig `pulumi:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *GatewayPluginKeyAuthConsumer      `pulumi:"consumer"`
	ConsumerGroup *GatewayPluginKeyAuthConsumerGroup `pulumi:"consumerGroup"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId string `pulumi:"controlPlaneId"`
	// Whether the plugin is applied.
	Enabled      *bool   `pulumi:"enabled"`
	InstanceName *string `pulumi:"instanceName"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []string `pulumi:"protocols"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *GatewayPluginKeyAuthRoute `pulumi:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *GatewayPluginKeyAuthService `pulumi:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a GatewayPluginKeyAuth resource.
type GatewayPluginKeyAuthArgs struct {
	Config GatewayPluginKeyAuthConfigPtrInput
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      GatewayPluginKeyAuthConsumerPtrInput
	ConsumerGroup GatewayPluginKeyAuthConsumerGroupPtrInput
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId pulumi.StringInput
	// Whether the plugin is applied.
	Enabled      pulumi.BoolPtrInput
	InstanceName pulumi.StringPtrInput
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols pulumi.StringArrayInput
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route GatewayPluginKeyAuthRoutePtrInput
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service GatewayPluginKeyAuthServicePtrInput
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags pulumi.StringArrayInput
}

func (GatewayPluginKeyAuthArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayPluginKeyAuthArgs)(nil)).Elem()
}

type GatewayPluginKeyAuthInput interface {
	pulumi.Input

	ToGatewayPluginKeyAuthOutput() GatewayPluginKeyAuthOutput
	ToGatewayPluginKeyAuthOutputWithContext(ctx context.Context) GatewayPluginKeyAuthOutput
}

func (*GatewayPluginKeyAuth) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayPluginKeyAuth)(nil)).Elem()
}

func (i *GatewayPluginKeyAuth) ToGatewayPluginKeyAuthOutput() GatewayPluginKeyAuthOutput {
	return i.ToGatewayPluginKeyAuthOutputWithContext(context.Background())
}

func (i *GatewayPluginKeyAuth) ToGatewayPluginKeyAuthOutputWithContext(ctx context.Context) GatewayPluginKeyAuthOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayPluginKeyAuthOutput)
}

// GatewayPluginKeyAuthArrayInput is an input type that accepts GatewayPluginKeyAuthArray and GatewayPluginKeyAuthArrayOutput values.
// You can construct a concrete instance of `GatewayPluginKeyAuthArrayInput` via:
//
//	GatewayPluginKeyAuthArray{ GatewayPluginKeyAuthArgs{...} }
type GatewayPluginKeyAuthArrayInput interface {
	pulumi.Input

	ToGatewayPluginKeyAuthArrayOutput() GatewayPluginKeyAuthArrayOutput
	ToGatewayPluginKeyAuthArrayOutputWithContext(context.Context) GatewayPluginKeyAuthArrayOutput
}

type GatewayPluginKeyAuthArray []GatewayPluginKeyAuthInput

func (GatewayPluginKeyAuthArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayPluginKeyAuth)(nil)).Elem()
}

func (i GatewayPluginKeyAuthArray) ToGatewayPluginKeyAuthArrayOutput() GatewayPluginKeyAuthArrayOutput {
	return i.ToGatewayPluginKeyAuthArrayOutputWithContext(context.Background())
}

func (i GatewayPluginKeyAuthArray) ToGatewayPluginKeyAuthArrayOutputWithContext(ctx context.Context) GatewayPluginKeyAuthArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayPluginKeyAuthArrayOutput)
}

// GatewayPluginKeyAuthMapInput is an input type that accepts GatewayPluginKeyAuthMap and GatewayPluginKeyAuthMapOutput values.
// You can construct a concrete instance of `GatewayPluginKeyAuthMapInput` via:
//
//	GatewayPluginKeyAuthMap{ "key": GatewayPluginKeyAuthArgs{...} }
type GatewayPluginKeyAuthMapInput interface {
	pulumi.Input

	ToGatewayPluginKeyAuthMapOutput() GatewayPluginKeyAuthMapOutput
	ToGatewayPluginKeyAuthMapOutputWithContext(context.Context) GatewayPluginKeyAuthMapOutput
}

type GatewayPluginKeyAuthMap map[string]GatewayPluginKeyAuthInput

func (GatewayPluginKeyAuthMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayPluginKeyAuth)(nil)).Elem()
}

func (i GatewayPluginKeyAuthMap) ToGatewayPluginKeyAuthMapOutput() GatewayPluginKeyAuthMapOutput {
	return i.ToGatewayPluginKeyAuthMapOutputWithContext(context.Background())
}

func (i GatewayPluginKeyAuthMap) ToGatewayPluginKeyAuthMapOutputWithContext(ctx context.Context) GatewayPluginKeyAuthMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayPluginKeyAuthMapOutput)
}

type GatewayPluginKeyAuthOutput struct{ *pulumi.OutputState }

func (GatewayPluginKeyAuthOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayPluginKeyAuth)(nil)).Elem()
}

func (o GatewayPluginKeyAuthOutput) ToGatewayPluginKeyAuthOutput() GatewayPluginKeyAuthOutput {
	return o
}

func (o GatewayPluginKeyAuthOutput) ToGatewayPluginKeyAuthOutputWithContext(ctx context.Context) GatewayPluginKeyAuthOutput {
	return o
}

func (o GatewayPluginKeyAuthOutput) Config() GatewayPluginKeyAuthConfigOutput {
	return o.ApplyT(func(v *GatewayPluginKeyAuth) GatewayPluginKeyAuthConfigOutput { return v.Config }).(GatewayPluginKeyAuthConfigOutput)
}

// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
func (o GatewayPluginKeyAuthOutput) Consumer() GatewayPluginKeyAuthConsumerOutput {
	return o.ApplyT(func(v *GatewayPluginKeyAuth) GatewayPluginKeyAuthConsumerOutput { return v.Consumer }).(GatewayPluginKeyAuthConsumerOutput)
}

func (o GatewayPluginKeyAuthOutput) ConsumerGroup() GatewayPluginKeyAuthConsumerGroupOutput {
	return o.ApplyT(func(v *GatewayPluginKeyAuth) GatewayPluginKeyAuthConsumerGroupOutput { return v.ConsumerGroup }).(GatewayPluginKeyAuthConsumerGroupOutput)
}

// The UUID of your control plane. This variable is available in the Konnect manager.
func (o GatewayPluginKeyAuthOutput) ControlPlaneId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayPluginKeyAuth) pulumi.StringOutput { return v.ControlPlaneId }).(pulumi.StringOutput)
}

// Unix epoch when the resource was created.
func (o GatewayPluginKeyAuthOutput) CreatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayPluginKeyAuth) pulumi.IntOutput { return v.CreatedAt }).(pulumi.IntOutput)
}

// Whether the plugin is applied.
func (o GatewayPluginKeyAuthOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *GatewayPluginKeyAuth) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

func (o GatewayPluginKeyAuthOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayPluginKeyAuth) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
func (o GatewayPluginKeyAuthOutput) Protocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayPluginKeyAuth) pulumi.StringArrayOutput { return v.Protocols }).(pulumi.StringArrayOutput)
}

// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
func (o GatewayPluginKeyAuthOutput) Route() GatewayPluginKeyAuthRouteOutput {
	return o.ApplyT(func(v *GatewayPluginKeyAuth) GatewayPluginKeyAuthRouteOutput { return v.Route }).(GatewayPluginKeyAuthRouteOutput)
}

// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
func (o GatewayPluginKeyAuthOutput) Service() GatewayPluginKeyAuthServiceOutput {
	return o.ApplyT(func(v *GatewayPluginKeyAuth) GatewayPluginKeyAuthServiceOutput { return v.Service }).(GatewayPluginKeyAuthServiceOutput)
}

// An optional set of strings associated with the Plugin for grouping and filtering.
func (o GatewayPluginKeyAuthOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayPluginKeyAuth) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Unix epoch when the resource was last updated.
func (o GatewayPluginKeyAuthOutput) UpdatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayPluginKeyAuth) pulumi.IntOutput { return v.UpdatedAt }).(pulumi.IntOutput)
}

type GatewayPluginKeyAuthArrayOutput struct{ *pulumi.OutputState }

func (GatewayPluginKeyAuthArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayPluginKeyAuth)(nil)).Elem()
}

func (o GatewayPluginKeyAuthArrayOutput) ToGatewayPluginKeyAuthArrayOutput() GatewayPluginKeyAuthArrayOutput {
	return o
}

func (o GatewayPluginKeyAuthArrayOutput) ToGatewayPluginKeyAuthArrayOutputWithContext(ctx context.Context) GatewayPluginKeyAuthArrayOutput {
	return o
}

func (o GatewayPluginKeyAuthArrayOutput) Index(i pulumi.IntInput) GatewayPluginKeyAuthOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GatewayPluginKeyAuth {
		return vs[0].([]*GatewayPluginKeyAuth)[vs[1].(int)]
	}).(GatewayPluginKeyAuthOutput)
}

type GatewayPluginKeyAuthMapOutput struct{ *pulumi.OutputState }

func (GatewayPluginKeyAuthMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayPluginKeyAuth)(nil)).Elem()
}

func (o GatewayPluginKeyAuthMapOutput) ToGatewayPluginKeyAuthMapOutput() GatewayPluginKeyAuthMapOutput {
	return o
}

func (o GatewayPluginKeyAuthMapOutput) ToGatewayPluginKeyAuthMapOutputWithContext(ctx context.Context) GatewayPluginKeyAuthMapOutput {
	return o
}

func (o GatewayPluginKeyAuthMapOutput) MapIndex(k pulumi.StringInput) GatewayPluginKeyAuthOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GatewayPluginKeyAuth {
		return vs[0].(map[string]*GatewayPluginKeyAuth)[vs[1].(string)]
	}).(GatewayPluginKeyAuthOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayPluginKeyAuthInput)(nil)).Elem(), &GatewayPluginKeyAuth{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayPluginKeyAuthArrayInput)(nil)).Elem(), GatewayPluginKeyAuthArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayPluginKeyAuthMapInput)(nil)).Elem(), GatewayPluginKeyAuthMap{})
	pulumi.RegisterOutputType(GatewayPluginKeyAuthOutput{})
	pulumi.RegisterOutputType(GatewayPluginKeyAuthArrayOutput{})
	pulumi.RegisterOutputType(GatewayPluginKeyAuthMapOutput{})
}