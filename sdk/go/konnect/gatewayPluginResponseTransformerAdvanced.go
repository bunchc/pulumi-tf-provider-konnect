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

// GatewayPluginResponseTransformerAdvanced Resource
type GatewayPluginResponseTransformerAdvanced struct {
	pulumi.CustomResourceState

	Config GatewayPluginResponseTransformerAdvancedConfigOutput `pulumi:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      GatewayPluginResponseTransformerAdvancedConsumerOutput      `pulumi:"consumer"`
	ConsumerGroup GatewayPluginResponseTransformerAdvancedConsumerGroupOutput `pulumi:"consumerGroup"`
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
	Route GatewayPluginResponseTransformerAdvancedRouteOutput `pulumi:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service GatewayPluginResponseTransformerAdvancedServiceOutput `pulumi:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Unix epoch when the resource was last updated.
	UpdatedAt pulumi.IntOutput `pulumi:"updatedAt"`
}

// NewGatewayPluginResponseTransformerAdvanced registers a new resource with the given unique name, arguments, and options.
func NewGatewayPluginResponseTransformerAdvanced(ctx *pulumi.Context,
	name string, args *GatewayPluginResponseTransformerAdvancedArgs, opts ...pulumi.ResourceOption) (*GatewayPluginResponseTransformerAdvanced, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ControlPlaneId == nil {
		return nil, errors.New("invalid value for required argument 'ControlPlaneId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GatewayPluginResponseTransformerAdvanced
	err := ctx.RegisterResource("konnect:index/gatewayPluginResponseTransformerAdvanced:GatewayPluginResponseTransformerAdvanced", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayPluginResponseTransformerAdvanced gets an existing GatewayPluginResponseTransformerAdvanced resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayPluginResponseTransformerAdvanced(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayPluginResponseTransformerAdvancedState, opts ...pulumi.ResourceOption) (*GatewayPluginResponseTransformerAdvanced, error) {
	var resource GatewayPluginResponseTransformerAdvanced
	err := ctx.ReadResource("konnect:index/gatewayPluginResponseTransformerAdvanced:GatewayPluginResponseTransformerAdvanced", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayPluginResponseTransformerAdvanced resources.
type gatewayPluginResponseTransformerAdvancedState struct {
	Config *GatewayPluginResponseTransformerAdvancedConfig `pulumi:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *GatewayPluginResponseTransformerAdvancedConsumer      `pulumi:"consumer"`
	ConsumerGroup *GatewayPluginResponseTransformerAdvancedConsumerGroup `pulumi:"consumerGroup"`
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
	Route *GatewayPluginResponseTransformerAdvancedRoute `pulumi:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *GatewayPluginResponseTransformerAdvancedService `pulumi:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `pulumi:"tags"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int `pulumi:"updatedAt"`
}

type GatewayPluginResponseTransformerAdvancedState struct {
	Config GatewayPluginResponseTransformerAdvancedConfigPtrInput
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      GatewayPluginResponseTransformerAdvancedConsumerPtrInput
	ConsumerGroup GatewayPluginResponseTransformerAdvancedConsumerGroupPtrInput
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
	Route GatewayPluginResponseTransformerAdvancedRoutePtrInput
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service GatewayPluginResponseTransformerAdvancedServicePtrInput
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags pulumi.StringArrayInput
	// Unix epoch when the resource was last updated.
	UpdatedAt pulumi.IntPtrInput
}

func (GatewayPluginResponseTransformerAdvancedState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayPluginResponseTransformerAdvancedState)(nil)).Elem()
}

type gatewayPluginResponseTransformerAdvancedArgs struct {
	Config *GatewayPluginResponseTransformerAdvancedConfig `pulumi:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *GatewayPluginResponseTransformerAdvancedConsumer      `pulumi:"consumer"`
	ConsumerGroup *GatewayPluginResponseTransformerAdvancedConsumerGroup `pulumi:"consumerGroup"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId string `pulumi:"controlPlaneId"`
	// Whether the plugin is applied.
	Enabled      *bool   `pulumi:"enabled"`
	InstanceName *string `pulumi:"instanceName"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []string `pulumi:"protocols"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *GatewayPluginResponseTransformerAdvancedRoute `pulumi:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *GatewayPluginResponseTransformerAdvancedService `pulumi:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a GatewayPluginResponseTransformerAdvanced resource.
type GatewayPluginResponseTransformerAdvancedArgs struct {
	Config GatewayPluginResponseTransformerAdvancedConfigPtrInput
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      GatewayPluginResponseTransformerAdvancedConsumerPtrInput
	ConsumerGroup GatewayPluginResponseTransformerAdvancedConsumerGroupPtrInput
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId pulumi.StringInput
	// Whether the plugin is applied.
	Enabled      pulumi.BoolPtrInput
	InstanceName pulumi.StringPtrInput
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols pulumi.StringArrayInput
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route GatewayPluginResponseTransformerAdvancedRoutePtrInput
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service GatewayPluginResponseTransformerAdvancedServicePtrInput
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags pulumi.StringArrayInput
}

func (GatewayPluginResponseTransformerAdvancedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayPluginResponseTransformerAdvancedArgs)(nil)).Elem()
}

type GatewayPluginResponseTransformerAdvancedInput interface {
	pulumi.Input

	ToGatewayPluginResponseTransformerAdvancedOutput() GatewayPluginResponseTransformerAdvancedOutput
	ToGatewayPluginResponseTransformerAdvancedOutputWithContext(ctx context.Context) GatewayPluginResponseTransformerAdvancedOutput
}

func (*GatewayPluginResponseTransformerAdvanced) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayPluginResponseTransformerAdvanced)(nil)).Elem()
}

func (i *GatewayPluginResponseTransformerAdvanced) ToGatewayPluginResponseTransformerAdvancedOutput() GatewayPluginResponseTransformerAdvancedOutput {
	return i.ToGatewayPluginResponseTransformerAdvancedOutputWithContext(context.Background())
}

func (i *GatewayPluginResponseTransformerAdvanced) ToGatewayPluginResponseTransformerAdvancedOutputWithContext(ctx context.Context) GatewayPluginResponseTransformerAdvancedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayPluginResponseTransformerAdvancedOutput)
}

// GatewayPluginResponseTransformerAdvancedArrayInput is an input type that accepts GatewayPluginResponseTransformerAdvancedArray and GatewayPluginResponseTransformerAdvancedArrayOutput values.
// You can construct a concrete instance of `GatewayPluginResponseTransformerAdvancedArrayInput` via:
//
//	GatewayPluginResponseTransformerAdvancedArray{ GatewayPluginResponseTransformerAdvancedArgs{...} }
type GatewayPluginResponseTransformerAdvancedArrayInput interface {
	pulumi.Input

	ToGatewayPluginResponseTransformerAdvancedArrayOutput() GatewayPluginResponseTransformerAdvancedArrayOutput
	ToGatewayPluginResponseTransformerAdvancedArrayOutputWithContext(context.Context) GatewayPluginResponseTransformerAdvancedArrayOutput
}

type GatewayPluginResponseTransformerAdvancedArray []GatewayPluginResponseTransformerAdvancedInput

func (GatewayPluginResponseTransformerAdvancedArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayPluginResponseTransformerAdvanced)(nil)).Elem()
}

func (i GatewayPluginResponseTransformerAdvancedArray) ToGatewayPluginResponseTransformerAdvancedArrayOutput() GatewayPluginResponseTransformerAdvancedArrayOutput {
	return i.ToGatewayPluginResponseTransformerAdvancedArrayOutputWithContext(context.Background())
}

func (i GatewayPluginResponseTransformerAdvancedArray) ToGatewayPluginResponseTransformerAdvancedArrayOutputWithContext(ctx context.Context) GatewayPluginResponseTransformerAdvancedArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayPluginResponseTransformerAdvancedArrayOutput)
}

// GatewayPluginResponseTransformerAdvancedMapInput is an input type that accepts GatewayPluginResponseTransformerAdvancedMap and GatewayPluginResponseTransformerAdvancedMapOutput values.
// You can construct a concrete instance of `GatewayPluginResponseTransformerAdvancedMapInput` via:
//
//	GatewayPluginResponseTransformerAdvancedMap{ "key": GatewayPluginResponseTransformerAdvancedArgs{...} }
type GatewayPluginResponseTransformerAdvancedMapInput interface {
	pulumi.Input

	ToGatewayPluginResponseTransformerAdvancedMapOutput() GatewayPluginResponseTransformerAdvancedMapOutput
	ToGatewayPluginResponseTransformerAdvancedMapOutputWithContext(context.Context) GatewayPluginResponseTransformerAdvancedMapOutput
}

type GatewayPluginResponseTransformerAdvancedMap map[string]GatewayPluginResponseTransformerAdvancedInput

func (GatewayPluginResponseTransformerAdvancedMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayPluginResponseTransformerAdvanced)(nil)).Elem()
}

func (i GatewayPluginResponseTransformerAdvancedMap) ToGatewayPluginResponseTransformerAdvancedMapOutput() GatewayPluginResponseTransformerAdvancedMapOutput {
	return i.ToGatewayPluginResponseTransformerAdvancedMapOutputWithContext(context.Background())
}

func (i GatewayPluginResponseTransformerAdvancedMap) ToGatewayPluginResponseTransformerAdvancedMapOutputWithContext(ctx context.Context) GatewayPluginResponseTransformerAdvancedMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayPluginResponseTransformerAdvancedMapOutput)
}

type GatewayPluginResponseTransformerAdvancedOutput struct{ *pulumi.OutputState }

func (GatewayPluginResponseTransformerAdvancedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayPluginResponseTransformerAdvanced)(nil)).Elem()
}

func (o GatewayPluginResponseTransformerAdvancedOutput) ToGatewayPluginResponseTransformerAdvancedOutput() GatewayPluginResponseTransformerAdvancedOutput {
	return o
}

func (o GatewayPluginResponseTransformerAdvancedOutput) ToGatewayPluginResponseTransformerAdvancedOutputWithContext(ctx context.Context) GatewayPluginResponseTransformerAdvancedOutput {
	return o
}

func (o GatewayPluginResponseTransformerAdvancedOutput) Config() GatewayPluginResponseTransformerAdvancedConfigOutput {
	return o.ApplyT(func(v *GatewayPluginResponseTransformerAdvanced) GatewayPluginResponseTransformerAdvancedConfigOutput {
		return v.Config
	}).(GatewayPluginResponseTransformerAdvancedConfigOutput)
}

// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
func (o GatewayPluginResponseTransformerAdvancedOutput) Consumer() GatewayPluginResponseTransformerAdvancedConsumerOutput {
	return o.ApplyT(func(v *GatewayPluginResponseTransformerAdvanced) GatewayPluginResponseTransformerAdvancedConsumerOutput {
		return v.Consumer
	}).(GatewayPluginResponseTransformerAdvancedConsumerOutput)
}

func (o GatewayPluginResponseTransformerAdvancedOutput) ConsumerGroup() GatewayPluginResponseTransformerAdvancedConsumerGroupOutput {
	return o.ApplyT(func(v *GatewayPluginResponseTransformerAdvanced) GatewayPluginResponseTransformerAdvancedConsumerGroupOutput {
		return v.ConsumerGroup
	}).(GatewayPluginResponseTransformerAdvancedConsumerGroupOutput)
}

// The UUID of your control plane. This variable is available in the Konnect manager.
func (o GatewayPluginResponseTransformerAdvancedOutput) ControlPlaneId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayPluginResponseTransformerAdvanced) pulumi.StringOutput { return v.ControlPlaneId }).(pulumi.StringOutput)
}

// Unix epoch when the resource was created.
func (o GatewayPluginResponseTransformerAdvancedOutput) CreatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayPluginResponseTransformerAdvanced) pulumi.IntOutput { return v.CreatedAt }).(pulumi.IntOutput)
}

// Whether the plugin is applied.
func (o GatewayPluginResponseTransformerAdvancedOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *GatewayPluginResponseTransformerAdvanced) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

func (o GatewayPluginResponseTransformerAdvancedOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayPluginResponseTransformerAdvanced) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
func (o GatewayPluginResponseTransformerAdvancedOutput) Protocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayPluginResponseTransformerAdvanced) pulumi.StringArrayOutput { return v.Protocols }).(pulumi.StringArrayOutput)
}

// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
func (o GatewayPluginResponseTransformerAdvancedOutput) Route() GatewayPluginResponseTransformerAdvancedRouteOutput {
	return o.ApplyT(func(v *GatewayPluginResponseTransformerAdvanced) GatewayPluginResponseTransformerAdvancedRouteOutput {
		return v.Route
	}).(GatewayPluginResponseTransformerAdvancedRouteOutput)
}

// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
func (o GatewayPluginResponseTransformerAdvancedOutput) Service() GatewayPluginResponseTransformerAdvancedServiceOutput {
	return o.ApplyT(func(v *GatewayPluginResponseTransformerAdvanced) GatewayPluginResponseTransformerAdvancedServiceOutput {
		return v.Service
	}).(GatewayPluginResponseTransformerAdvancedServiceOutput)
}

// An optional set of strings associated with the Plugin for grouping and filtering.
func (o GatewayPluginResponseTransformerAdvancedOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayPluginResponseTransformerAdvanced) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Unix epoch when the resource was last updated.
func (o GatewayPluginResponseTransformerAdvancedOutput) UpdatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayPluginResponseTransformerAdvanced) pulumi.IntOutput { return v.UpdatedAt }).(pulumi.IntOutput)
}

type GatewayPluginResponseTransformerAdvancedArrayOutput struct{ *pulumi.OutputState }

func (GatewayPluginResponseTransformerAdvancedArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayPluginResponseTransformerAdvanced)(nil)).Elem()
}

func (o GatewayPluginResponseTransformerAdvancedArrayOutput) ToGatewayPluginResponseTransformerAdvancedArrayOutput() GatewayPluginResponseTransformerAdvancedArrayOutput {
	return o
}

func (o GatewayPluginResponseTransformerAdvancedArrayOutput) ToGatewayPluginResponseTransformerAdvancedArrayOutputWithContext(ctx context.Context) GatewayPluginResponseTransformerAdvancedArrayOutput {
	return o
}

func (o GatewayPluginResponseTransformerAdvancedArrayOutput) Index(i pulumi.IntInput) GatewayPluginResponseTransformerAdvancedOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GatewayPluginResponseTransformerAdvanced {
		return vs[0].([]*GatewayPluginResponseTransformerAdvanced)[vs[1].(int)]
	}).(GatewayPluginResponseTransformerAdvancedOutput)
}

type GatewayPluginResponseTransformerAdvancedMapOutput struct{ *pulumi.OutputState }

func (GatewayPluginResponseTransformerAdvancedMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayPluginResponseTransformerAdvanced)(nil)).Elem()
}

func (o GatewayPluginResponseTransformerAdvancedMapOutput) ToGatewayPluginResponseTransformerAdvancedMapOutput() GatewayPluginResponseTransformerAdvancedMapOutput {
	return o
}

func (o GatewayPluginResponseTransformerAdvancedMapOutput) ToGatewayPluginResponseTransformerAdvancedMapOutputWithContext(ctx context.Context) GatewayPluginResponseTransformerAdvancedMapOutput {
	return o
}

func (o GatewayPluginResponseTransformerAdvancedMapOutput) MapIndex(k pulumi.StringInput) GatewayPluginResponseTransformerAdvancedOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GatewayPluginResponseTransformerAdvanced {
		return vs[0].(map[string]*GatewayPluginResponseTransformerAdvanced)[vs[1].(string)]
	}).(GatewayPluginResponseTransformerAdvancedOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayPluginResponseTransformerAdvancedInput)(nil)).Elem(), &GatewayPluginResponseTransformerAdvanced{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayPluginResponseTransformerAdvancedArrayInput)(nil)).Elem(), GatewayPluginResponseTransformerAdvancedArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayPluginResponseTransformerAdvancedMapInput)(nil)).Elem(), GatewayPluginResponseTransformerAdvancedMap{})
	pulumi.RegisterOutputType(GatewayPluginResponseTransformerAdvancedOutput{})
	pulumi.RegisterOutputType(GatewayPluginResponseTransformerAdvancedArrayOutput{})
	pulumi.RegisterOutputType(GatewayPluginResponseTransformerAdvancedMapOutput{})
}
