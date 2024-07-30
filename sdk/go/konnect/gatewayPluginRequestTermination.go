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

// GatewayPluginRequestTermination Resource
type GatewayPluginRequestTermination struct {
	pulumi.CustomResourceState

	Config GatewayPluginRequestTerminationConfigOutput `pulumi:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      GatewayPluginRequestTerminationConsumerOutput      `pulumi:"consumer"`
	ConsumerGroup GatewayPluginRequestTerminationConsumerGroupOutput `pulumi:"consumerGroup"`
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
	Route GatewayPluginRequestTerminationRouteOutput `pulumi:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service GatewayPluginRequestTerminationServiceOutput `pulumi:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Unix epoch when the resource was last updated.
	UpdatedAt pulumi.IntOutput `pulumi:"updatedAt"`
}

// NewGatewayPluginRequestTermination registers a new resource with the given unique name, arguments, and options.
func NewGatewayPluginRequestTermination(ctx *pulumi.Context,
	name string, args *GatewayPluginRequestTerminationArgs, opts ...pulumi.ResourceOption) (*GatewayPluginRequestTermination, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ControlPlaneId == nil {
		return nil, errors.New("invalid value for required argument 'ControlPlaneId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GatewayPluginRequestTermination
	err := ctx.RegisterResource("konnect:index/gatewayPluginRequestTermination:GatewayPluginRequestTermination", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayPluginRequestTermination gets an existing GatewayPluginRequestTermination resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayPluginRequestTermination(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayPluginRequestTerminationState, opts ...pulumi.ResourceOption) (*GatewayPluginRequestTermination, error) {
	var resource GatewayPluginRequestTermination
	err := ctx.ReadResource("konnect:index/gatewayPluginRequestTermination:GatewayPluginRequestTermination", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayPluginRequestTermination resources.
type gatewayPluginRequestTerminationState struct {
	Config *GatewayPluginRequestTerminationConfig `pulumi:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *GatewayPluginRequestTerminationConsumer      `pulumi:"consumer"`
	ConsumerGroup *GatewayPluginRequestTerminationConsumerGroup `pulumi:"consumerGroup"`
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
	Route *GatewayPluginRequestTerminationRoute `pulumi:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *GatewayPluginRequestTerminationService `pulumi:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `pulumi:"tags"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int `pulumi:"updatedAt"`
}

type GatewayPluginRequestTerminationState struct {
	Config GatewayPluginRequestTerminationConfigPtrInput
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      GatewayPluginRequestTerminationConsumerPtrInput
	ConsumerGroup GatewayPluginRequestTerminationConsumerGroupPtrInput
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
	Route GatewayPluginRequestTerminationRoutePtrInput
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service GatewayPluginRequestTerminationServicePtrInput
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags pulumi.StringArrayInput
	// Unix epoch when the resource was last updated.
	UpdatedAt pulumi.IntPtrInput
}

func (GatewayPluginRequestTerminationState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayPluginRequestTerminationState)(nil)).Elem()
}

type gatewayPluginRequestTerminationArgs struct {
	Config *GatewayPluginRequestTerminationConfig `pulumi:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *GatewayPluginRequestTerminationConsumer      `pulumi:"consumer"`
	ConsumerGroup *GatewayPluginRequestTerminationConsumerGroup `pulumi:"consumerGroup"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId string `pulumi:"controlPlaneId"`
	// Whether the plugin is applied.
	Enabled      *bool   `pulumi:"enabled"`
	InstanceName *string `pulumi:"instanceName"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []string `pulumi:"protocols"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *GatewayPluginRequestTerminationRoute `pulumi:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *GatewayPluginRequestTerminationService `pulumi:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a GatewayPluginRequestTermination resource.
type GatewayPluginRequestTerminationArgs struct {
	Config GatewayPluginRequestTerminationConfigPtrInput
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      GatewayPluginRequestTerminationConsumerPtrInput
	ConsumerGroup GatewayPluginRequestTerminationConsumerGroupPtrInput
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId pulumi.StringInput
	// Whether the plugin is applied.
	Enabled      pulumi.BoolPtrInput
	InstanceName pulumi.StringPtrInput
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols pulumi.StringArrayInput
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route GatewayPluginRequestTerminationRoutePtrInput
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service GatewayPluginRequestTerminationServicePtrInput
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags pulumi.StringArrayInput
}

func (GatewayPluginRequestTerminationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayPluginRequestTerminationArgs)(nil)).Elem()
}

type GatewayPluginRequestTerminationInput interface {
	pulumi.Input

	ToGatewayPluginRequestTerminationOutput() GatewayPluginRequestTerminationOutput
	ToGatewayPluginRequestTerminationOutputWithContext(ctx context.Context) GatewayPluginRequestTerminationOutput
}

func (*GatewayPluginRequestTermination) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayPluginRequestTermination)(nil)).Elem()
}

func (i *GatewayPluginRequestTermination) ToGatewayPluginRequestTerminationOutput() GatewayPluginRequestTerminationOutput {
	return i.ToGatewayPluginRequestTerminationOutputWithContext(context.Background())
}

func (i *GatewayPluginRequestTermination) ToGatewayPluginRequestTerminationOutputWithContext(ctx context.Context) GatewayPluginRequestTerminationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayPluginRequestTerminationOutput)
}

// GatewayPluginRequestTerminationArrayInput is an input type that accepts GatewayPluginRequestTerminationArray and GatewayPluginRequestTerminationArrayOutput values.
// You can construct a concrete instance of `GatewayPluginRequestTerminationArrayInput` via:
//
//	GatewayPluginRequestTerminationArray{ GatewayPluginRequestTerminationArgs{...} }
type GatewayPluginRequestTerminationArrayInput interface {
	pulumi.Input

	ToGatewayPluginRequestTerminationArrayOutput() GatewayPluginRequestTerminationArrayOutput
	ToGatewayPluginRequestTerminationArrayOutputWithContext(context.Context) GatewayPluginRequestTerminationArrayOutput
}

type GatewayPluginRequestTerminationArray []GatewayPluginRequestTerminationInput

func (GatewayPluginRequestTerminationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayPluginRequestTermination)(nil)).Elem()
}

func (i GatewayPluginRequestTerminationArray) ToGatewayPluginRequestTerminationArrayOutput() GatewayPluginRequestTerminationArrayOutput {
	return i.ToGatewayPluginRequestTerminationArrayOutputWithContext(context.Background())
}

func (i GatewayPluginRequestTerminationArray) ToGatewayPluginRequestTerminationArrayOutputWithContext(ctx context.Context) GatewayPluginRequestTerminationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayPluginRequestTerminationArrayOutput)
}

// GatewayPluginRequestTerminationMapInput is an input type that accepts GatewayPluginRequestTerminationMap and GatewayPluginRequestTerminationMapOutput values.
// You can construct a concrete instance of `GatewayPluginRequestTerminationMapInput` via:
//
//	GatewayPluginRequestTerminationMap{ "key": GatewayPluginRequestTerminationArgs{...} }
type GatewayPluginRequestTerminationMapInput interface {
	pulumi.Input

	ToGatewayPluginRequestTerminationMapOutput() GatewayPluginRequestTerminationMapOutput
	ToGatewayPluginRequestTerminationMapOutputWithContext(context.Context) GatewayPluginRequestTerminationMapOutput
}

type GatewayPluginRequestTerminationMap map[string]GatewayPluginRequestTerminationInput

func (GatewayPluginRequestTerminationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayPluginRequestTermination)(nil)).Elem()
}

func (i GatewayPluginRequestTerminationMap) ToGatewayPluginRequestTerminationMapOutput() GatewayPluginRequestTerminationMapOutput {
	return i.ToGatewayPluginRequestTerminationMapOutputWithContext(context.Background())
}

func (i GatewayPluginRequestTerminationMap) ToGatewayPluginRequestTerminationMapOutputWithContext(ctx context.Context) GatewayPluginRequestTerminationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayPluginRequestTerminationMapOutput)
}

type GatewayPluginRequestTerminationOutput struct{ *pulumi.OutputState }

func (GatewayPluginRequestTerminationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayPluginRequestTermination)(nil)).Elem()
}

func (o GatewayPluginRequestTerminationOutput) ToGatewayPluginRequestTerminationOutput() GatewayPluginRequestTerminationOutput {
	return o
}

func (o GatewayPluginRequestTerminationOutput) ToGatewayPluginRequestTerminationOutputWithContext(ctx context.Context) GatewayPluginRequestTerminationOutput {
	return o
}

func (o GatewayPluginRequestTerminationOutput) Config() GatewayPluginRequestTerminationConfigOutput {
	return o.ApplyT(func(v *GatewayPluginRequestTermination) GatewayPluginRequestTerminationConfigOutput { return v.Config }).(GatewayPluginRequestTerminationConfigOutput)
}

// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
func (o GatewayPluginRequestTerminationOutput) Consumer() GatewayPluginRequestTerminationConsumerOutput {
	return o.ApplyT(func(v *GatewayPluginRequestTermination) GatewayPluginRequestTerminationConsumerOutput {
		return v.Consumer
	}).(GatewayPluginRequestTerminationConsumerOutput)
}

func (o GatewayPluginRequestTerminationOutput) ConsumerGroup() GatewayPluginRequestTerminationConsumerGroupOutput {
	return o.ApplyT(func(v *GatewayPluginRequestTermination) GatewayPluginRequestTerminationConsumerGroupOutput {
		return v.ConsumerGroup
	}).(GatewayPluginRequestTerminationConsumerGroupOutput)
}

// The UUID of your control plane. This variable is available in the Konnect manager.
func (o GatewayPluginRequestTerminationOutput) ControlPlaneId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayPluginRequestTermination) pulumi.StringOutput { return v.ControlPlaneId }).(pulumi.StringOutput)
}

// Unix epoch when the resource was created.
func (o GatewayPluginRequestTerminationOutput) CreatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayPluginRequestTermination) pulumi.IntOutput { return v.CreatedAt }).(pulumi.IntOutput)
}

// Whether the plugin is applied.
func (o GatewayPluginRequestTerminationOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *GatewayPluginRequestTermination) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

func (o GatewayPluginRequestTerminationOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayPluginRequestTermination) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
func (o GatewayPluginRequestTerminationOutput) Protocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayPluginRequestTermination) pulumi.StringArrayOutput { return v.Protocols }).(pulumi.StringArrayOutput)
}

// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
func (o GatewayPluginRequestTerminationOutput) Route() GatewayPluginRequestTerminationRouteOutput {
	return o.ApplyT(func(v *GatewayPluginRequestTermination) GatewayPluginRequestTerminationRouteOutput { return v.Route }).(GatewayPluginRequestTerminationRouteOutput)
}

// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
func (o GatewayPluginRequestTerminationOutput) Service() GatewayPluginRequestTerminationServiceOutput {
	return o.ApplyT(func(v *GatewayPluginRequestTermination) GatewayPluginRequestTerminationServiceOutput {
		return v.Service
	}).(GatewayPluginRequestTerminationServiceOutput)
}

// An optional set of strings associated with the Plugin for grouping and filtering.
func (o GatewayPluginRequestTerminationOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayPluginRequestTermination) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Unix epoch when the resource was last updated.
func (o GatewayPluginRequestTerminationOutput) UpdatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayPluginRequestTermination) pulumi.IntOutput { return v.UpdatedAt }).(pulumi.IntOutput)
}

type GatewayPluginRequestTerminationArrayOutput struct{ *pulumi.OutputState }

func (GatewayPluginRequestTerminationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayPluginRequestTermination)(nil)).Elem()
}

func (o GatewayPluginRequestTerminationArrayOutput) ToGatewayPluginRequestTerminationArrayOutput() GatewayPluginRequestTerminationArrayOutput {
	return o
}

func (o GatewayPluginRequestTerminationArrayOutput) ToGatewayPluginRequestTerminationArrayOutputWithContext(ctx context.Context) GatewayPluginRequestTerminationArrayOutput {
	return o
}

func (o GatewayPluginRequestTerminationArrayOutput) Index(i pulumi.IntInput) GatewayPluginRequestTerminationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GatewayPluginRequestTermination {
		return vs[0].([]*GatewayPluginRequestTermination)[vs[1].(int)]
	}).(GatewayPluginRequestTerminationOutput)
}

type GatewayPluginRequestTerminationMapOutput struct{ *pulumi.OutputState }

func (GatewayPluginRequestTerminationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayPluginRequestTermination)(nil)).Elem()
}

func (o GatewayPluginRequestTerminationMapOutput) ToGatewayPluginRequestTerminationMapOutput() GatewayPluginRequestTerminationMapOutput {
	return o
}

func (o GatewayPluginRequestTerminationMapOutput) ToGatewayPluginRequestTerminationMapOutputWithContext(ctx context.Context) GatewayPluginRequestTerminationMapOutput {
	return o
}

func (o GatewayPluginRequestTerminationMapOutput) MapIndex(k pulumi.StringInput) GatewayPluginRequestTerminationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GatewayPluginRequestTermination {
		return vs[0].(map[string]*GatewayPluginRequestTermination)[vs[1].(string)]
	}).(GatewayPluginRequestTerminationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayPluginRequestTerminationInput)(nil)).Elem(), &GatewayPluginRequestTermination{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayPluginRequestTerminationArrayInput)(nil)).Elem(), GatewayPluginRequestTerminationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayPluginRequestTerminationMapInput)(nil)).Elem(), GatewayPluginRequestTerminationMap{})
	pulumi.RegisterOutputType(GatewayPluginRequestTerminationOutput{})
	pulumi.RegisterOutputType(GatewayPluginRequestTerminationArrayOutput{})
	pulumi.RegisterOutputType(GatewayPluginRequestTerminationMapOutput{})
}