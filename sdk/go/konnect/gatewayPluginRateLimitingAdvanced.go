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

// GatewayPluginRateLimitingAdvanced Resource
type GatewayPluginRateLimitingAdvanced struct {
	pulumi.CustomResourceState

	Config GatewayPluginRateLimitingAdvancedConfigOutput `pulumi:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      GatewayPluginRateLimitingAdvancedConsumerOutput      `pulumi:"consumer"`
	ConsumerGroup GatewayPluginRateLimitingAdvancedConsumerGroupOutput `pulumi:"consumerGroup"`
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
	Route GatewayPluginRateLimitingAdvancedRouteOutput `pulumi:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service GatewayPluginRateLimitingAdvancedServiceOutput `pulumi:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Unix epoch when the resource was last updated.
	UpdatedAt pulumi.IntOutput `pulumi:"updatedAt"`
}

// NewGatewayPluginRateLimitingAdvanced registers a new resource with the given unique name, arguments, and options.
func NewGatewayPluginRateLimitingAdvanced(ctx *pulumi.Context,
	name string, args *GatewayPluginRateLimitingAdvancedArgs, opts ...pulumi.ResourceOption) (*GatewayPluginRateLimitingAdvanced, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ControlPlaneId == nil {
		return nil, errors.New("invalid value for required argument 'ControlPlaneId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GatewayPluginRateLimitingAdvanced
	err := ctx.RegisterResource("konnect:index/gatewayPluginRateLimitingAdvanced:GatewayPluginRateLimitingAdvanced", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayPluginRateLimitingAdvanced gets an existing GatewayPluginRateLimitingAdvanced resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayPluginRateLimitingAdvanced(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayPluginRateLimitingAdvancedState, opts ...pulumi.ResourceOption) (*GatewayPluginRateLimitingAdvanced, error) {
	var resource GatewayPluginRateLimitingAdvanced
	err := ctx.ReadResource("konnect:index/gatewayPluginRateLimitingAdvanced:GatewayPluginRateLimitingAdvanced", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayPluginRateLimitingAdvanced resources.
type gatewayPluginRateLimitingAdvancedState struct {
	Config *GatewayPluginRateLimitingAdvancedConfig `pulumi:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *GatewayPluginRateLimitingAdvancedConsumer      `pulumi:"consumer"`
	ConsumerGroup *GatewayPluginRateLimitingAdvancedConsumerGroup `pulumi:"consumerGroup"`
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
	Route *GatewayPluginRateLimitingAdvancedRoute `pulumi:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *GatewayPluginRateLimitingAdvancedService `pulumi:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `pulumi:"tags"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int `pulumi:"updatedAt"`
}

type GatewayPluginRateLimitingAdvancedState struct {
	Config GatewayPluginRateLimitingAdvancedConfigPtrInput
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      GatewayPluginRateLimitingAdvancedConsumerPtrInput
	ConsumerGroup GatewayPluginRateLimitingAdvancedConsumerGroupPtrInput
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
	Route GatewayPluginRateLimitingAdvancedRoutePtrInput
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service GatewayPluginRateLimitingAdvancedServicePtrInput
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags pulumi.StringArrayInput
	// Unix epoch when the resource was last updated.
	UpdatedAt pulumi.IntPtrInput
}

func (GatewayPluginRateLimitingAdvancedState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayPluginRateLimitingAdvancedState)(nil)).Elem()
}

type gatewayPluginRateLimitingAdvancedArgs struct {
	Config *GatewayPluginRateLimitingAdvancedConfig `pulumi:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *GatewayPluginRateLimitingAdvancedConsumer      `pulumi:"consumer"`
	ConsumerGroup *GatewayPluginRateLimitingAdvancedConsumerGroup `pulumi:"consumerGroup"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId string `pulumi:"controlPlaneId"`
	// Whether the plugin is applied.
	Enabled      *bool   `pulumi:"enabled"`
	InstanceName *string `pulumi:"instanceName"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []string `pulumi:"protocols"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *GatewayPluginRateLimitingAdvancedRoute `pulumi:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *GatewayPluginRateLimitingAdvancedService `pulumi:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a GatewayPluginRateLimitingAdvanced resource.
type GatewayPluginRateLimitingAdvancedArgs struct {
	Config GatewayPluginRateLimitingAdvancedConfigPtrInput
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      GatewayPluginRateLimitingAdvancedConsumerPtrInput
	ConsumerGroup GatewayPluginRateLimitingAdvancedConsumerGroupPtrInput
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId pulumi.StringInput
	// Whether the plugin is applied.
	Enabled      pulumi.BoolPtrInput
	InstanceName pulumi.StringPtrInput
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols pulumi.StringArrayInput
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route GatewayPluginRateLimitingAdvancedRoutePtrInput
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service GatewayPluginRateLimitingAdvancedServicePtrInput
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags pulumi.StringArrayInput
}

func (GatewayPluginRateLimitingAdvancedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayPluginRateLimitingAdvancedArgs)(nil)).Elem()
}

type GatewayPluginRateLimitingAdvancedInput interface {
	pulumi.Input

	ToGatewayPluginRateLimitingAdvancedOutput() GatewayPluginRateLimitingAdvancedOutput
	ToGatewayPluginRateLimitingAdvancedOutputWithContext(ctx context.Context) GatewayPluginRateLimitingAdvancedOutput
}

func (*GatewayPluginRateLimitingAdvanced) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayPluginRateLimitingAdvanced)(nil)).Elem()
}

func (i *GatewayPluginRateLimitingAdvanced) ToGatewayPluginRateLimitingAdvancedOutput() GatewayPluginRateLimitingAdvancedOutput {
	return i.ToGatewayPluginRateLimitingAdvancedOutputWithContext(context.Background())
}

func (i *GatewayPluginRateLimitingAdvanced) ToGatewayPluginRateLimitingAdvancedOutputWithContext(ctx context.Context) GatewayPluginRateLimitingAdvancedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayPluginRateLimitingAdvancedOutput)
}

// GatewayPluginRateLimitingAdvancedArrayInput is an input type that accepts GatewayPluginRateLimitingAdvancedArray and GatewayPluginRateLimitingAdvancedArrayOutput values.
// You can construct a concrete instance of `GatewayPluginRateLimitingAdvancedArrayInput` via:
//
//	GatewayPluginRateLimitingAdvancedArray{ GatewayPluginRateLimitingAdvancedArgs{...} }
type GatewayPluginRateLimitingAdvancedArrayInput interface {
	pulumi.Input

	ToGatewayPluginRateLimitingAdvancedArrayOutput() GatewayPluginRateLimitingAdvancedArrayOutput
	ToGatewayPluginRateLimitingAdvancedArrayOutputWithContext(context.Context) GatewayPluginRateLimitingAdvancedArrayOutput
}

type GatewayPluginRateLimitingAdvancedArray []GatewayPluginRateLimitingAdvancedInput

func (GatewayPluginRateLimitingAdvancedArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayPluginRateLimitingAdvanced)(nil)).Elem()
}

func (i GatewayPluginRateLimitingAdvancedArray) ToGatewayPluginRateLimitingAdvancedArrayOutput() GatewayPluginRateLimitingAdvancedArrayOutput {
	return i.ToGatewayPluginRateLimitingAdvancedArrayOutputWithContext(context.Background())
}

func (i GatewayPluginRateLimitingAdvancedArray) ToGatewayPluginRateLimitingAdvancedArrayOutputWithContext(ctx context.Context) GatewayPluginRateLimitingAdvancedArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayPluginRateLimitingAdvancedArrayOutput)
}

// GatewayPluginRateLimitingAdvancedMapInput is an input type that accepts GatewayPluginRateLimitingAdvancedMap and GatewayPluginRateLimitingAdvancedMapOutput values.
// You can construct a concrete instance of `GatewayPluginRateLimitingAdvancedMapInput` via:
//
//	GatewayPluginRateLimitingAdvancedMap{ "key": GatewayPluginRateLimitingAdvancedArgs{...} }
type GatewayPluginRateLimitingAdvancedMapInput interface {
	pulumi.Input

	ToGatewayPluginRateLimitingAdvancedMapOutput() GatewayPluginRateLimitingAdvancedMapOutput
	ToGatewayPluginRateLimitingAdvancedMapOutputWithContext(context.Context) GatewayPluginRateLimitingAdvancedMapOutput
}

type GatewayPluginRateLimitingAdvancedMap map[string]GatewayPluginRateLimitingAdvancedInput

func (GatewayPluginRateLimitingAdvancedMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayPluginRateLimitingAdvanced)(nil)).Elem()
}

func (i GatewayPluginRateLimitingAdvancedMap) ToGatewayPluginRateLimitingAdvancedMapOutput() GatewayPluginRateLimitingAdvancedMapOutput {
	return i.ToGatewayPluginRateLimitingAdvancedMapOutputWithContext(context.Background())
}

func (i GatewayPluginRateLimitingAdvancedMap) ToGatewayPluginRateLimitingAdvancedMapOutputWithContext(ctx context.Context) GatewayPluginRateLimitingAdvancedMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayPluginRateLimitingAdvancedMapOutput)
}

type GatewayPluginRateLimitingAdvancedOutput struct{ *pulumi.OutputState }

func (GatewayPluginRateLimitingAdvancedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayPluginRateLimitingAdvanced)(nil)).Elem()
}

func (o GatewayPluginRateLimitingAdvancedOutput) ToGatewayPluginRateLimitingAdvancedOutput() GatewayPluginRateLimitingAdvancedOutput {
	return o
}

func (o GatewayPluginRateLimitingAdvancedOutput) ToGatewayPluginRateLimitingAdvancedOutputWithContext(ctx context.Context) GatewayPluginRateLimitingAdvancedOutput {
	return o
}

func (o GatewayPluginRateLimitingAdvancedOutput) Config() GatewayPluginRateLimitingAdvancedConfigOutput {
	return o.ApplyT(func(v *GatewayPluginRateLimitingAdvanced) GatewayPluginRateLimitingAdvancedConfigOutput {
		return v.Config
	}).(GatewayPluginRateLimitingAdvancedConfigOutput)
}

// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
func (o GatewayPluginRateLimitingAdvancedOutput) Consumer() GatewayPluginRateLimitingAdvancedConsumerOutput {
	return o.ApplyT(func(v *GatewayPluginRateLimitingAdvanced) GatewayPluginRateLimitingAdvancedConsumerOutput {
		return v.Consumer
	}).(GatewayPluginRateLimitingAdvancedConsumerOutput)
}

func (o GatewayPluginRateLimitingAdvancedOutput) ConsumerGroup() GatewayPluginRateLimitingAdvancedConsumerGroupOutput {
	return o.ApplyT(func(v *GatewayPluginRateLimitingAdvanced) GatewayPluginRateLimitingAdvancedConsumerGroupOutput {
		return v.ConsumerGroup
	}).(GatewayPluginRateLimitingAdvancedConsumerGroupOutput)
}

// The UUID of your control plane. This variable is available in the Konnect manager.
func (o GatewayPluginRateLimitingAdvancedOutput) ControlPlaneId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayPluginRateLimitingAdvanced) pulumi.StringOutput { return v.ControlPlaneId }).(pulumi.StringOutput)
}

// Unix epoch when the resource was created.
func (o GatewayPluginRateLimitingAdvancedOutput) CreatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayPluginRateLimitingAdvanced) pulumi.IntOutput { return v.CreatedAt }).(pulumi.IntOutput)
}

// Whether the plugin is applied.
func (o GatewayPluginRateLimitingAdvancedOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *GatewayPluginRateLimitingAdvanced) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

func (o GatewayPluginRateLimitingAdvancedOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayPluginRateLimitingAdvanced) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
func (o GatewayPluginRateLimitingAdvancedOutput) Protocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayPluginRateLimitingAdvanced) pulumi.StringArrayOutput { return v.Protocols }).(pulumi.StringArrayOutput)
}

// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
func (o GatewayPluginRateLimitingAdvancedOutput) Route() GatewayPluginRateLimitingAdvancedRouteOutput {
	return o.ApplyT(func(v *GatewayPluginRateLimitingAdvanced) GatewayPluginRateLimitingAdvancedRouteOutput {
		return v.Route
	}).(GatewayPluginRateLimitingAdvancedRouteOutput)
}

// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
func (o GatewayPluginRateLimitingAdvancedOutput) Service() GatewayPluginRateLimitingAdvancedServiceOutput {
	return o.ApplyT(func(v *GatewayPluginRateLimitingAdvanced) GatewayPluginRateLimitingAdvancedServiceOutput {
		return v.Service
	}).(GatewayPluginRateLimitingAdvancedServiceOutput)
}

// An optional set of strings associated with the Plugin for grouping and filtering.
func (o GatewayPluginRateLimitingAdvancedOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayPluginRateLimitingAdvanced) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Unix epoch when the resource was last updated.
func (o GatewayPluginRateLimitingAdvancedOutput) UpdatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayPluginRateLimitingAdvanced) pulumi.IntOutput { return v.UpdatedAt }).(pulumi.IntOutput)
}

type GatewayPluginRateLimitingAdvancedArrayOutput struct{ *pulumi.OutputState }

func (GatewayPluginRateLimitingAdvancedArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayPluginRateLimitingAdvanced)(nil)).Elem()
}

func (o GatewayPluginRateLimitingAdvancedArrayOutput) ToGatewayPluginRateLimitingAdvancedArrayOutput() GatewayPluginRateLimitingAdvancedArrayOutput {
	return o
}

func (o GatewayPluginRateLimitingAdvancedArrayOutput) ToGatewayPluginRateLimitingAdvancedArrayOutputWithContext(ctx context.Context) GatewayPluginRateLimitingAdvancedArrayOutput {
	return o
}

func (o GatewayPluginRateLimitingAdvancedArrayOutput) Index(i pulumi.IntInput) GatewayPluginRateLimitingAdvancedOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GatewayPluginRateLimitingAdvanced {
		return vs[0].([]*GatewayPluginRateLimitingAdvanced)[vs[1].(int)]
	}).(GatewayPluginRateLimitingAdvancedOutput)
}

type GatewayPluginRateLimitingAdvancedMapOutput struct{ *pulumi.OutputState }

func (GatewayPluginRateLimitingAdvancedMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayPluginRateLimitingAdvanced)(nil)).Elem()
}

func (o GatewayPluginRateLimitingAdvancedMapOutput) ToGatewayPluginRateLimitingAdvancedMapOutput() GatewayPluginRateLimitingAdvancedMapOutput {
	return o
}

func (o GatewayPluginRateLimitingAdvancedMapOutput) ToGatewayPluginRateLimitingAdvancedMapOutputWithContext(ctx context.Context) GatewayPluginRateLimitingAdvancedMapOutput {
	return o
}

func (o GatewayPluginRateLimitingAdvancedMapOutput) MapIndex(k pulumi.StringInput) GatewayPluginRateLimitingAdvancedOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GatewayPluginRateLimitingAdvanced {
		return vs[0].(map[string]*GatewayPluginRateLimitingAdvanced)[vs[1].(string)]
	}).(GatewayPluginRateLimitingAdvancedOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayPluginRateLimitingAdvancedInput)(nil)).Elem(), &GatewayPluginRateLimitingAdvanced{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayPluginRateLimitingAdvancedArrayInput)(nil)).Elem(), GatewayPluginRateLimitingAdvancedArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayPluginRateLimitingAdvancedMapInput)(nil)).Elem(), GatewayPluginRateLimitingAdvancedMap{})
	pulumi.RegisterOutputType(GatewayPluginRateLimitingAdvancedOutput{})
	pulumi.RegisterOutputType(GatewayPluginRateLimitingAdvancedArrayOutput{})
	pulumi.RegisterOutputType(GatewayPluginRateLimitingAdvancedMapOutput{})
}
