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

// GatewayPluginCORS Resource
type GatewayPluginCors struct {
	pulumi.CustomResourceState

	Config GatewayPluginCorsConfigOutput `pulumi:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      GatewayPluginCorsConsumerOutput      `pulumi:"consumer"`
	ConsumerGroup GatewayPluginCorsConsumerGroupOutput `pulumi:"consumerGroup"`
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
	Route GatewayPluginCorsRouteOutput `pulumi:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service GatewayPluginCorsServiceOutput `pulumi:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Unix epoch when the resource was last updated.
	UpdatedAt pulumi.IntOutput `pulumi:"updatedAt"`
}

// NewGatewayPluginCors registers a new resource with the given unique name, arguments, and options.
func NewGatewayPluginCors(ctx *pulumi.Context,
	name string, args *GatewayPluginCorsArgs, opts ...pulumi.ResourceOption) (*GatewayPluginCors, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ControlPlaneId == nil {
		return nil, errors.New("invalid value for required argument 'ControlPlaneId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GatewayPluginCors
	err := ctx.RegisterResource("konnect:index/gatewayPluginCors:GatewayPluginCors", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayPluginCors gets an existing GatewayPluginCors resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayPluginCors(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayPluginCorsState, opts ...pulumi.ResourceOption) (*GatewayPluginCors, error) {
	var resource GatewayPluginCors
	err := ctx.ReadResource("konnect:index/gatewayPluginCors:GatewayPluginCors", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayPluginCors resources.
type gatewayPluginCorsState struct {
	Config *GatewayPluginCorsConfig `pulumi:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *GatewayPluginCorsConsumer      `pulumi:"consumer"`
	ConsumerGroup *GatewayPluginCorsConsumerGroup `pulumi:"consumerGroup"`
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
	Route *GatewayPluginCorsRoute `pulumi:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *GatewayPluginCorsService `pulumi:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `pulumi:"tags"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int `pulumi:"updatedAt"`
}

type GatewayPluginCorsState struct {
	Config GatewayPluginCorsConfigPtrInput
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      GatewayPluginCorsConsumerPtrInput
	ConsumerGroup GatewayPluginCorsConsumerGroupPtrInput
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
	Route GatewayPluginCorsRoutePtrInput
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service GatewayPluginCorsServicePtrInput
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags pulumi.StringArrayInput
	// Unix epoch when the resource was last updated.
	UpdatedAt pulumi.IntPtrInput
}

func (GatewayPluginCorsState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayPluginCorsState)(nil)).Elem()
}

type gatewayPluginCorsArgs struct {
	Config *GatewayPluginCorsConfig `pulumi:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *GatewayPluginCorsConsumer      `pulumi:"consumer"`
	ConsumerGroup *GatewayPluginCorsConsumerGroup `pulumi:"consumerGroup"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId string `pulumi:"controlPlaneId"`
	// Whether the plugin is applied.
	Enabled      *bool   `pulumi:"enabled"`
	InstanceName *string `pulumi:"instanceName"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []string `pulumi:"protocols"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *GatewayPluginCorsRoute `pulumi:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *GatewayPluginCorsService `pulumi:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a GatewayPluginCors resource.
type GatewayPluginCorsArgs struct {
	Config GatewayPluginCorsConfigPtrInput
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      GatewayPluginCorsConsumerPtrInput
	ConsumerGroup GatewayPluginCorsConsumerGroupPtrInput
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId pulumi.StringInput
	// Whether the plugin is applied.
	Enabled      pulumi.BoolPtrInput
	InstanceName pulumi.StringPtrInput
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols pulumi.StringArrayInput
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route GatewayPluginCorsRoutePtrInput
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service GatewayPluginCorsServicePtrInput
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags pulumi.StringArrayInput
}

func (GatewayPluginCorsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayPluginCorsArgs)(nil)).Elem()
}

type GatewayPluginCorsInput interface {
	pulumi.Input

	ToGatewayPluginCorsOutput() GatewayPluginCorsOutput
	ToGatewayPluginCorsOutputWithContext(ctx context.Context) GatewayPluginCorsOutput
}

func (*GatewayPluginCors) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayPluginCors)(nil)).Elem()
}

func (i *GatewayPluginCors) ToGatewayPluginCorsOutput() GatewayPluginCorsOutput {
	return i.ToGatewayPluginCorsOutputWithContext(context.Background())
}

func (i *GatewayPluginCors) ToGatewayPluginCorsOutputWithContext(ctx context.Context) GatewayPluginCorsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayPluginCorsOutput)
}

// GatewayPluginCorsArrayInput is an input type that accepts GatewayPluginCorsArray and GatewayPluginCorsArrayOutput values.
// You can construct a concrete instance of `GatewayPluginCorsArrayInput` via:
//
//	GatewayPluginCorsArray{ GatewayPluginCorsArgs{...} }
type GatewayPluginCorsArrayInput interface {
	pulumi.Input

	ToGatewayPluginCorsArrayOutput() GatewayPluginCorsArrayOutput
	ToGatewayPluginCorsArrayOutputWithContext(context.Context) GatewayPluginCorsArrayOutput
}

type GatewayPluginCorsArray []GatewayPluginCorsInput

func (GatewayPluginCorsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayPluginCors)(nil)).Elem()
}

func (i GatewayPluginCorsArray) ToGatewayPluginCorsArrayOutput() GatewayPluginCorsArrayOutput {
	return i.ToGatewayPluginCorsArrayOutputWithContext(context.Background())
}

func (i GatewayPluginCorsArray) ToGatewayPluginCorsArrayOutputWithContext(ctx context.Context) GatewayPluginCorsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayPluginCorsArrayOutput)
}

// GatewayPluginCorsMapInput is an input type that accepts GatewayPluginCorsMap and GatewayPluginCorsMapOutput values.
// You can construct a concrete instance of `GatewayPluginCorsMapInput` via:
//
//	GatewayPluginCorsMap{ "key": GatewayPluginCorsArgs{...} }
type GatewayPluginCorsMapInput interface {
	pulumi.Input

	ToGatewayPluginCorsMapOutput() GatewayPluginCorsMapOutput
	ToGatewayPluginCorsMapOutputWithContext(context.Context) GatewayPluginCorsMapOutput
}

type GatewayPluginCorsMap map[string]GatewayPluginCorsInput

func (GatewayPluginCorsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayPluginCors)(nil)).Elem()
}

func (i GatewayPluginCorsMap) ToGatewayPluginCorsMapOutput() GatewayPluginCorsMapOutput {
	return i.ToGatewayPluginCorsMapOutputWithContext(context.Background())
}

func (i GatewayPluginCorsMap) ToGatewayPluginCorsMapOutputWithContext(ctx context.Context) GatewayPluginCorsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayPluginCorsMapOutput)
}

type GatewayPluginCorsOutput struct{ *pulumi.OutputState }

func (GatewayPluginCorsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayPluginCors)(nil)).Elem()
}

func (o GatewayPluginCorsOutput) ToGatewayPluginCorsOutput() GatewayPluginCorsOutput {
	return o
}

func (o GatewayPluginCorsOutput) ToGatewayPluginCorsOutputWithContext(ctx context.Context) GatewayPluginCorsOutput {
	return o
}

func (o GatewayPluginCorsOutput) Config() GatewayPluginCorsConfigOutput {
	return o.ApplyT(func(v *GatewayPluginCors) GatewayPluginCorsConfigOutput { return v.Config }).(GatewayPluginCorsConfigOutput)
}

// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
func (o GatewayPluginCorsOutput) Consumer() GatewayPluginCorsConsumerOutput {
	return o.ApplyT(func(v *GatewayPluginCors) GatewayPluginCorsConsumerOutput { return v.Consumer }).(GatewayPluginCorsConsumerOutput)
}

func (o GatewayPluginCorsOutput) ConsumerGroup() GatewayPluginCorsConsumerGroupOutput {
	return o.ApplyT(func(v *GatewayPluginCors) GatewayPluginCorsConsumerGroupOutput { return v.ConsumerGroup }).(GatewayPluginCorsConsumerGroupOutput)
}

// The UUID of your control plane. This variable is available in the Konnect manager.
func (o GatewayPluginCorsOutput) ControlPlaneId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayPluginCors) pulumi.StringOutput { return v.ControlPlaneId }).(pulumi.StringOutput)
}

// Unix epoch when the resource was created.
func (o GatewayPluginCorsOutput) CreatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayPluginCors) pulumi.IntOutput { return v.CreatedAt }).(pulumi.IntOutput)
}

// Whether the plugin is applied.
func (o GatewayPluginCorsOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *GatewayPluginCors) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

func (o GatewayPluginCorsOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayPluginCors) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
func (o GatewayPluginCorsOutput) Protocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayPluginCors) pulumi.StringArrayOutput { return v.Protocols }).(pulumi.StringArrayOutput)
}

// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
func (o GatewayPluginCorsOutput) Route() GatewayPluginCorsRouteOutput {
	return o.ApplyT(func(v *GatewayPluginCors) GatewayPluginCorsRouteOutput { return v.Route }).(GatewayPluginCorsRouteOutput)
}

// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
func (o GatewayPluginCorsOutput) Service() GatewayPluginCorsServiceOutput {
	return o.ApplyT(func(v *GatewayPluginCors) GatewayPluginCorsServiceOutput { return v.Service }).(GatewayPluginCorsServiceOutput)
}

// An optional set of strings associated with the Plugin for grouping and filtering.
func (o GatewayPluginCorsOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayPluginCors) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Unix epoch when the resource was last updated.
func (o GatewayPluginCorsOutput) UpdatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayPluginCors) pulumi.IntOutput { return v.UpdatedAt }).(pulumi.IntOutput)
}

type GatewayPluginCorsArrayOutput struct{ *pulumi.OutputState }

func (GatewayPluginCorsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayPluginCors)(nil)).Elem()
}

func (o GatewayPluginCorsArrayOutput) ToGatewayPluginCorsArrayOutput() GatewayPluginCorsArrayOutput {
	return o
}

func (o GatewayPluginCorsArrayOutput) ToGatewayPluginCorsArrayOutputWithContext(ctx context.Context) GatewayPluginCorsArrayOutput {
	return o
}

func (o GatewayPluginCorsArrayOutput) Index(i pulumi.IntInput) GatewayPluginCorsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GatewayPluginCors {
		return vs[0].([]*GatewayPluginCors)[vs[1].(int)]
	}).(GatewayPluginCorsOutput)
}

type GatewayPluginCorsMapOutput struct{ *pulumi.OutputState }

func (GatewayPluginCorsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayPluginCors)(nil)).Elem()
}

func (o GatewayPluginCorsMapOutput) ToGatewayPluginCorsMapOutput() GatewayPluginCorsMapOutput {
	return o
}

func (o GatewayPluginCorsMapOutput) ToGatewayPluginCorsMapOutputWithContext(ctx context.Context) GatewayPluginCorsMapOutput {
	return o
}

func (o GatewayPluginCorsMapOutput) MapIndex(k pulumi.StringInput) GatewayPluginCorsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GatewayPluginCors {
		return vs[0].(map[string]*GatewayPluginCors)[vs[1].(string)]
	}).(GatewayPluginCorsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayPluginCorsInput)(nil)).Elem(), &GatewayPluginCors{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayPluginCorsArrayInput)(nil)).Elem(), GatewayPluginCorsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayPluginCorsMapInput)(nil)).Elem(), GatewayPluginCorsMap{})
	pulumi.RegisterOutputType(GatewayPluginCorsOutput{})
	pulumi.RegisterOutputType(GatewayPluginCorsArrayOutput{})
	pulumi.RegisterOutputType(GatewayPluginCorsMapOutput{})
}
