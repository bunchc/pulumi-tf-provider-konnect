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

// GatewayPluginProxyCache Resource
type GatewayPluginProxyCache struct {
	pulumi.CustomResourceState

	Config GatewayPluginProxyCacheConfigOutput `pulumi:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      GatewayPluginProxyCacheConsumerOutput      `pulumi:"consumer"`
	ConsumerGroup GatewayPluginProxyCacheConsumerGroupOutput `pulumi:"consumerGroup"`
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
	Route GatewayPluginProxyCacheRouteOutput `pulumi:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service GatewayPluginProxyCacheServiceOutput `pulumi:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Unix epoch when the resource was last updated.
	UpdatedAt pulumi.IntOutput `pulumi:"updatedAt"`
}

// NewGatewayPluginProxyCache registers a new resource with the given unique name, arguments, and options.
func NewGatewayPluginProxyCache(ctx *pulumi.Context,
	name string, args *GatewayPluginProxyCacheArgs, opts ...pulumi.ResourceOption) (*GatewayPluginProxyCache, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ControlPlaneId == nil {
		return nil, errors.New("invalid value for required argument 'ControlPlaneId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GatewayPluginProxyCache
	err := ctx.RegisterResource("konnect:index/gatewayPluginProxyCache:GatewayPluginProxyCache", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayPluginProxyCache gets an existing GatewayPluginProxyCache resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayPluginProxyCache(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayPluginProxyCacheState, opts ...pulumi.ResourceOption) (*GatewayPluginProxyCache, error) {
	var resource GatewayPluginProxyCache
	err := ctx.ReadResource("konnect:index/gatewayPluginProxyCache:GatewayPluginProxyCache", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayPluginProxyCache resources.
type gatewayPluginProxyCacheState struct {
	Config *GatewayPluginProxyCacheConfig `pulumi:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *GatewayPluginProxyCacheConsumer      `pulumi:"consumer"`
	ConsumerGroup *GatewayPluginProxyCacheConsumerGroup `pulumi:"consumerGroup"`
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
	Route *GatewayPluginProxyCacheRoute `pulumi:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *GatewayPluginProxyCacheService `pulumi:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `pulumi:"tags"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int `pulumi:"updatedAt"`
}

type GatewayPluginProxyCacheState struct {
	Config GatewayPluginProxyCacheConfigPtrInput
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      GatewayPluginProxyCacheConsumerPtrInput
	ConsumerGroup GatewayPluginProxyCacheConsumerGroupPtrInput
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
	Route GatewayPluginProxyCacheRoutePtrInput
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service GatewayPluginProxyCacheServicePtrInput
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags pulumi.StringArrayInput
	// Unix epoch when the resource was last updated.
	UpdatedAt pulumi.IntPtrInput
}

func (GatewayPluginProxyCacheState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayPluginProxyCacheState)(nil)).Elem()
}

type gatewayPluginProxyCacheArgs struct {
	Config *GatewayPluginProxyCacheConfig `pulumi:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *GatewayPluginProxyCacheConsumer      `pulumi:"consumer"`
	ConsumerGroup *GatewayPluginProxyCacheConsumerGroup `pulumi:"consumerGroup"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId string `pulumi:"controlPlaneId"`
	// Whether the plugin is applied.
	Enabled      *bool   `pulumi:"enabled"`
	InstanceName *string `pulumi:"instanceName"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []string `pulumi:"protocols"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *GatewayPluginProxyCacheRoute `pulumi:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *GatewayPluginProxyCacheService `pulumi:"service"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a GatewayPluginProxyCache resource.
type GatewayPluginProxyCacheArgs struct {
	Config GatewayPluginProxyCacheConfigPtrInput
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      GatewayPluginProxyCacheConsumerPtrInput
	ConsumerGroup GatewayPluginProxyCacheConsumerGroupPtrInput
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId pulumi.StringInput
	// Whether the plugin is applied.
	Enabled      pulumi.BoolPtrInput
	InstanceName pulumi.StringPtrInput
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols pulumi.StringArrayInput
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route GatewayPluginProxyCacheRoutePtrInput
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service GatewayPluginProxyCacheServicePtrInput
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags pulumi.StringArrayInput
}

func (GatewayPluginProxyCacheArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayPluginProxyCacheArgs)(nil)).Elem()
}

type GatewayPluginProxyCacheInput interface {
	pulumi.Input

	ToGatewayPluginProxyCacheOutput() GatewayPluginProxyCacheOutput
	ToGatewayPluginProxyCacheOutputWithContext(ctx context.Context) GatewayPluginProxyCacheOutput
}

func (*GatewayPluginProxyCache) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayPluginProxyCache)(nil)).Elem()
}

func (i *GatewayPluginProxyCache) ToGatewayPluginProxyCacheOutput() GatewayPluginProxyCacheOutput {
	return i.ToGatewayPluginProxyCacheOutputWithContext(context.Background())
}

func (i *GatewayPluginProxyCache) ToGatewayPluginProxyCacheOutputWithContext(ctx context.Context) GatewayPluginProxyCacheOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayPluginProxyCacheOutput)
}

// GatewayPluginProxyCacheArrayInput is an input type that accepts GatewayPluginProxyCacheArray and GatewayPluginProxyCacheArrayOutput values.
// You can construct a concrete instance of `GatewayPluginProxyCacheArrayInput` via:
//
//	GatewayPluginProxyCacheArray{ GatewayPluginProxyCacheArgs{...} }
type GatewayPluginProxyCacheArrayInput interface {
	pulumi.Input

	ToGatewayPluginProxyCacheArrayOutput() GatewayPluginProxyCacheArrayOutput
	ToGatewayPluginProxyCacheArrayOutputWithContext(context.Context) GatewayPluginProxyCacheArrayOutput
}

type GatewayPluginProxyCacheArray []GatewayPluginProxyCacheInput

func (GatewayPluginProxyCacheArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayPluginProxyCache)(nil)).Elem()
}

func (i GatewayPluginProxyCacheArray) ToGatewayPluginProxyCacheArrayOutput() GatewayPluginProxyCacheArrayOutput {
	return i.ToGatewayPluginProxyCacheArrayOutputWithContext(context.Background())
}

func (i GatewayPluginProxyCacheArray) ToGatewayPluginProxyCacheArrayOutputWithContext(ctx context.Context) GatewayPluginProxyCacheArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayPluginProxyCacheArrayOutput)
}

// GatewayPluginProxyCacheMapInput is an input type that accepts GatewayPluginProxyCacheMap and GatewayPluginProxyCacheMapOutput values.
// You can construct a concrete instance of `GatewayPluginProxyCacheMapInput` via:
//
//	GatewayPluginProxyCacheMap{ "key": GatewayPluginProxyCacheArgs{...} }
type GatewayPluginProxyCacheMapInput interface {
	pulumi.Input

	ToGatewayPluginProxyCacheMapOutput() GatewayPluginProxyCacheMapOutput
	ToGatewayPluginProxyCacheMapOutputWithContext(context.Context) GatewayPluginProxyCacheMapOutput
}

type GatewayPluginProxyCacheMap map[string]GatewayPluginProxyCacheInput

func (GatewayPluginProxyCacheMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayPluginProxyCache)(nil)).Elem()
}

func (i GatewayPluginProxyCacheMap) ToGatewayPluginProxyCacheMapOutput() GatewayPluginProxyCacheMapOutput {
	return i.ToGatewayPluginProxyCacheMapOutputWithContext(context.Background())
}

func (i GatewayPluginProxyCacheMap) ToGatewayPluginProxyCacheMapOutputWithContext(ctx context.Context) GatewayPluginProxyCacheMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayPluginProxyCacheMapOutput)
}

type GatewayPluginProxyCacheOutput struct{ *pulumi.OutputState }

func (GatewayPluginProxyCacheOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayPluginProxyCache)(nil)).Elem()
}

func (o GatewayPluginProxyCacheOutput) ToGatewayPluginProxyCacheOutput() GatewayPluginProxyCacheOutput {
	return o
}

func (o GatewayPluginProxyCacheOutput) ToGatewayPluginProxyCacheOutputWithContext(ctx context.Context) GatewayPluginProxyCacheOutput {
	return o
}

func (o GatewayPluginProxyCacheOutput) Config() GatewayPluginProxyCacheConfigOutput {
	return o.ApplyT(func(v *GatewayPluginProxyCache) GatewayPluginProxyCacheConfigOutput { return v.Config }).(GatewayPluginProxyCacheConfigOutput)
}

// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
func (o GatewayPluginProxyCacheOutput) Consumer() GatewayPluginProxyCacheConsumerOutput {
	return o.ApplyT(func(v *GatewayPluginProxyCache) GatewayPluginProxyCacheConsumerOutput { return v.Consumer }).(GatewayPluginProxyCacheConsumerOutput)
}

func (o GatewayPluginProxyCacheOutput) ConsumerGroup() GatewayPluginProxyCacheConsumerGroupOutput {
	return o.ApplyT(func(v *GatewayPluginProxyCache) GatewayPluginProxyCacheConsumerGroupOutput { return v.ConsumerGroup }).(GatewayPluginProxyCacheConsumerGroupOutput)
}

// The UUID of your control plane. This variable is available in the Konnect manager.
func (o GatewayPluginProxyCacheOutput) ControlPlaneId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayPluginProxyCache) pulumi.StringOutput { return v.ControlPlaneId }).(pulumi.StringOutput)
}

// Unix epoch when the resource was created.
func (o GatewayPluginProxyCacheOutput) CreatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayPluginProxyCache) pulumi.IntOutput { return v.CreatedAt }).(pulumi.IntOutput)
}

// Whether the plugin is applied.
func (o GatewayPluginProxyCacheOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *GatewayPluginProxyCache) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

func (o GatewayPluginProxyCacheOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayPluginProxyCache) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
func (o GatewayPluginProxyCacheOutput) Protocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayPluginProxyCache) pulumi.StringArrayOutput { return v.Protocols }).(pulumi.StringArrayOutput)
}

// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
func (o GatewayPluginProxyCacheOutput) Route() GatewayPluginProxyCacheRouteOutput {
	return o.ApplyT(func(v *GatewayPluginProxyCache) GatewayPluginProxyCacheRouteOutput { return v.Route }).(GatewayPluginProxyCacheRouteOutput)
}

// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
func (o GatewayPluginProxyCacheOutput) Service() GatewayPluginProxyCacheServiceOutput {
	return o.ApplyT(func(v *GatewayPluginProxyCache) GatewayPluginProxyCacheServiceOutput { return v.Service }).(GatewayPluginProxyCacheServiceOutput)
}

// An optional set of strings associated with the Plugin for grouping and filtering.
func (o GatewayPluginProxyCacheOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayPluginProxyCache) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Unix epoch when the resource was last updated.
func (o GatewayPluginProxyCacheOutput) UpdatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayPluginProxyCache) pulumi.IntOutput { return v.UpdatedAt }).(pulumi.IntOutput)
}

type GatewayPluginProxyCacheArrayOutput struct{ *pulumi.OutputState }

func (GatewayPluginProxyCacheArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayPluginProxyCache)(nil)).Elem()
}

func (o GatewayPluginProxyCacheArrayOutput) ToGatewayPluginProxyCacheArrayOutput() GatewayPluginProxyCacheArrayOutput {
	return o
}

func (o GatewayPluginProxyCacheArrayOutput) ToGatewayPluginProxyCacheArrayOutputWithContext(ctx context.Context) GatewayPluginProxyCacheArrayOutput {
	return o
}

func (o GatewayPluginProxyCacheArrayOutput) Index(i pulumi.IntInput) GatewayPluginProxyCacheOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GatewayPluginProxyCache {
		return vs[0].([]*GatewayPluginProxyCache)[vs[1].(int)]
	}).(GatewayPluginProxyCacheOutput)
}

type GatewayPluginProxyCacheMapOutput struct{ *pulumi.OutputState }

func (GatewayPluginProxyCacheMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayPluginProxyCache)(nil)).Elem()
}

func (o GatewayPluginProxyCacheMapOutput) ToGatewayPluginProxyCacheMapOutput() GatewayPluginProxyCacheMapOutput {
	return o
}

func (o GatewayPluginProxyCacheMapOutput) ToGatewayPluginProxyCacheMapOutputWithContext(ctx context.Context) GatewayPluginProxyCacheMapOutput {
	return o
}

func (o GatewayPluginProxyCacheMapOutput) MapIndex(k pulumi.StringInput) GatewayPluginProxyCacheOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GatewayPluginProxyCache {
		return vs[0].(map[string]*GatewayPluginProxyCache)[vs[1].(string)]
	}).(GatewayPluginProxyCacheOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayPluginProxyCacheInput)(nil)).Elem(), &GatewayPluginProxyCache{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayPluginProxyCacheArrayInput)(nil)).Elem(), GatewayPluginProxyCacheArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayPluginProxyCacheMapInput)(nil)).Elem(), GatewayPluginProxyCacheMap{})
	pulumi.RegisterOutputType(GatewayPluginProxyCacheOutput{})
	pulumi.RegisterOutputType(GatewayPluginProxyCacheArrayOutput{})
	pulumi.RegisterOutputType(GatewayPluginProxyCacheMapOutput{})
}