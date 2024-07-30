// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package konnect

import (
	"context"
	"reflect"

	"github.com/bunchc/pulumi-tf-provider-konnect/sdk/go/konnect/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// GatewayControlPlane Resource
type GatewayControlPlane struct {
	pulumi.CustomResourceState

	// The auth type value of the cluster associated with the Runtime Group. must be one of ["pinned*client*certs", "pki*client*certs"]
	AuthType pulumi.StringPtrOutput `pulumi:"authType"`
	// Whether this control-plane can be used for cloud-gateways. Requires replacement if changed.
	CloudGateway pulumi.BoolPtrOutput `pulumi:"cloudGateway"`
	// The ClusterType value of the cluster associated with the Control Plane. Requires replacement if changed. ; must be one of ["CLUSTER*TYPE*CONTROL*PLANE", "CLUSTER*TYPE*HYBRID", "CLUSTER*TYPE*K8S*INGRESS*CONTROLLER", "CLUSTER*TYPE*CONTROL*PLANE*GROUP", "CLUSTER*TYPE_SERVERLESS"]
	ClusterType pulumi.StringPtrOutput `pulumi:"clusterType"`
	// CP configuration object for related access endpoints.
	Config GatewayControlPlaneConfigOutput `pulumi:"config"`
	// The description of the control plane in Konnect.
	Description pulumi.StringOutput `pulumi:"description"`
	// Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The name of the control plane.
	Name pulumi.StringOutput `pulumi:"name"`
	// Array of proxy URLs associated with reaching the data-planes connected to a control-plane.
	ProxyUrls GatewayControlPlaneProxyUrlArrayOutput `pulumi:"proxyUrls"`
}

// NewGatewayControlPlane registers a new resource with the given unique name, arguments, and options.
func NewGatewayControlPlane(ctx *pulumi.Context,
	name string, args *GatewayControlPlaneArgs, opts ...pulumi.ResourceOption) (*GatewayControlPlane, error) {
	if args == nil {
		args = &GatewayControlPlaneArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GatewayControlPlane
	err := ctx.RegisterResource("konnect:index/gatewayControlPlane:GatewayControlPlane", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayControlPlane gets an existing GatewayControlPlane resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayControlPlane(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayControlPlaneState, opts ...pulumi.ResourceOption) (*GatewayControlPlane, error) {
	var resource GatewayControlPlane
	err := ctx.ReadResource("konnect:index/gatewayControlPlane:GatewayControlPlane", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayControlPlane resources.
type gatewayControlPlaneState struct {
	// The auth type value of the cluster associated with the Runtime Group. must be one of ["pinned*client*certs", "pki*client*certs"]
	AuthType *string `pulumi:"authType"`
	// Whether this control-plane can be used for cloud-gateways. Requires replacement if changed.
	CloudGateway *bool `pulumi:"cloudGateway"`
	// The ClusterType value of the cluster associated with the Control Plane. Requires replacement if changed. ; must be one of ["CLUSTER*TYPE*CONTROL*PLANE", "CLUSTER*TYPE*HYBRID", "CLUSTER*TYPE*K8S*INGRESS*CONTROLLER", "CLUSTER*TYPE*CONTROL*PLANE*GROUP", "CLUSTER*TYPE_SERVERLESS"]
	ClusterType *string `pulumi:"clusterType"`
	// CP configuration object for related access endpoints.
	Config *GatewayControlPlaneConfig `pulumi:"config"`
	// The description of the control plane in Konnect.
	Description *string `pulumi:"description"`
	// Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
	Labels map[string]string `pulumi:"labels"`
	// The name of the control plane.
	Name *string `pulumi:"name"`
	// Array of proxy URLs associated with reaching the data-planes connected to a control-plane.
	ProxyUrls []GatewayControlPlaneProxyUrl `pulumi:"proxyUrls"`
}

type GatewayControlPlaneState struct {
	// The auth type value of the cluster associated with the Runtime Group. must be one of ["pinned*client*certs", "pki*client*certs"]
	AuthType pulumi.StringPtrInput
	// Whether this control-plane can be used for cloud-gateways. Requires replacement if changed.
	CloudGateway pulumi.BoolPtrInput
	// The ClusterType value of the cluster associated with the Control Plane. Requires replacement if changed. ; must be one of ["CLUSTER*TYPE*CONTROL*PLANE", "CLUSTER*TYPE*HYBRID", "CLUSTER*TYPE*K8S*INGRESS*CONTROLLER", "CLUSTER*TYPE*CONTROL*PLANE*GROUP", "CLUSTER*TYPE_SERVERLESS"]
	ClusterType pulumi.StringPtrInput
	// CP configuration object for related access endpoints.
	Config GatewayControlPlaneConfigPtrInput
	// The description of the control plane in Konnect.
	Description pulumi.StringPtrInput
	// Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
	Labels pulumi.StringMapInput
	// The name of the control plane.
	Name pulumi.StringPtrInput
	// Array of proxy URLs associated with reaching the data-planes connected to a control-plane.
	ProxyUrls GatewayControlPlaneProxyUrlArrayInput
}

func (GatewayControlPlaneState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayControlPlaneState)(nil)).Elem()
}

type gatewayControlPlaneArgs struct {
	// The auth type value of the cluster associated with the Runtime Group. must be one of ["pinned*client*certs", "pki*client*certs"]
	AuthType *string `pulumi:"authType"`
	// Whether this control-plane can be used for cloud-gateways. Requires replacement if changed.
	CloudGateway *bool `pulumi:"cloudGateway"`
	// The ClusterType value of the cluster associated with the Control Plane. Requires replacement if changed. ; must be one of ["CLUSTER*TYPE*CONTROL*PLANE", "CLUSTER*TYPE*HYBRID", "CLUSTER*TYPE*K8S*INGRESS*CONTROLLER", "CLUSTER*TYPE*CONTROL*PLANE*GROUP", "CLUSTER*TYPE_SERVERLESS"]
	ClusterType *string `pulumi:"clusterType"`
	// The description of the control plane in Konnect.
	Description *string `pulumi:"description"`
	// Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
	Labels map[string]string `pulumi:"labels"`
	// The name of the control plane.
	Name *string `pulumi:"name"`
	// Array of proxy URLs associated with reaching the data-planes connected to a control-plane.
	ProxyUrls []GatewayControlPlaneProxyUrl `pulumi:"proxyUrls"`
}

// The set of arguments for constructing a GatewayControlPlane resource.
type GatewayControlPlaneArgs struct {
	// The auth type value of the cluster associated with the Runtime Group. must be one of ["pinned*client*certs", "pki*client*certs"]
	AuthType pulumi.StringPtrInput
	// Whether this control-plane can be used for cloud-gateways. Requires replacement if changed.
	CloudGateway pulumi.BoolPtrInput
	// The ClusterType value of the cluster associated with the Control Plane. Requires replacement if changed. ; must be one of ["CLUSTER*TYPE*CONTROL*PLANE", "CLUSTER*TYPE*HYBRID", "CLUSTER*TYPE*K8S*INGRESS*CONTROLLER", "CLUSTER*TYPE*CONTROL*PLANE*GROUP", "CLUSTER*TYPE_SERVERLESS"]
	ClusterType pulumi.StringPtrInput
	// The description of the control plane in Konnect.
	Description pulumi.StringPtrInput
	// Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
	Labels pulumi.StringMapInput
	// The name of the control plane.
	Name pulumi.StringPtrInput
	// Array of proxy URLs associated with reaching the data-planes connected to a control-plane.
	ProxyUrls GatewayControlPlaneProxyUrlArrayInput
}

func (GatewayControlPlaneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayControlPlaneArgs)(nil)).Elem()
}

type GatewayControlPlaneInput interface {
	pulumi.Input

	ToGatewayControlPlaneOutput() GatewayControlPlaneOutput
	ToGatewayControlPlaneOutputWithContext(ctx context.Context) GatewayControlPlaneOutput
}

func (*GatewayControlPlane) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayControlPlane)(nil)).Elem()
}

func (i *GatewayControlPlane) ToGatewayControlPlaneOutput() GatewayControlPlaneOutput {
	return i.ToGatewayControlPlaneOutputWithContext(context.Background())
}

func (i *GatewayControlPlane) ToGatewayControlPlaneOutputWithContext(ctx context.Context) GatewayControlPlaneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayControlPlaneOutput)
}

// GatewayControlPlaneArrayInput is an input type that accepts GatewayControlPlaneArray and GatewayControlPlaneArrayOutput values.
// You can construct a concrete instance of `GatewayControlPlaneArrayInput` via:
//
//	GatewayControlPlaneArray{ GatewayControlPlaneArgs{...} }
type GatewayControlPlaneArrayInput interface {
	pulumi.Input

	ToGatewayControlPlaneArrayOutput() GatewayControlPlaneArrayOutput
	ToGatewayControlPlaneArrayOutputWithContext(context.Context) GatewayControlPlaneArrayOutput
}

type GatewayControlPlaneArray []GatewayControlPlaneInput

func (GatewayControlPlaneArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayControlPlane)(nil)).Elem()
}

func (i GatewayControlPlaneArray) ToGatewayControlPlaneArrayOutput() GatewayControlPlaneArrayOutput {
	return i.ToGatewayControlPlaneArrayOutputWithContext(context.Background())
}

func (i GatewayControlPlaneArray) ToGatewayControlPlaneArrayOutputWithContext(ctx context.Context) GatewayControlPlaneArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayControlPlaneArrayOutput)
}

// GatewayControlPlaneMapInput is an input type that accepts GatewayControlPlaneMap and GatewayControlPlaneMapOutput values.
// You can construct a concrete instance of `GatewayControlPlaneMapInput` via:
//
//	GatewayControlPlaneMap{ "key": GatewayControlPlaneArgs{...} }
type GatewayControlPlaneMapInput interface {
	pulumi.Input

	ToGatewayControlPlaneMapOutput() GatewayControlPlaneMapOutput
	ToGatewayControlPlaneMapOutputWithContext(context.Context) GatewayControlPlaneMapOutput
}

type GatewayControlPlaneMap map[string]GatewayControlPlaneInput

func (GatewayControlPlaneMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayControlPlane)(nil)).Elem()
}

func (i GatewayControlPlaneMap) ToGatewayControlPlaneMapOutput() GatewayControlPlaneMapOutput {
	return i.ToGatewayControlPlaneMapOutputWithContext(context.Background())
}

func (i GatewayControlPlaneMap) ToGatewayControlPlaneMapOutputWithContext(ctx context.Context) GatewayControlPlaneMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayControlPlaneMapOutput)
}

type GatewayControlPlaneOutput struct{ *pulumi.OutputState }

func (GatewayControlPlaneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayControlPlane)(nil)).Elem()
}

func (o GatewayControlPlaneOutput) ToGatewayControlPlaneOutput() GatewayControlPlaneOutput {
	return o
}

func (o GatewayControlPlaneOutput) ToGatewayControlPlaneOutputWithContext(ctx context.Context) GatewayControlPlaneOutput {
	return o
}

// The auth type value of the cluster associated with the Runtime Group. must be one of ["pinned*client*certs", "pki*client*certs"]
func (o GatewayControlPlaneOutput) AuthType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayControlPlane) pulumi.StringPtrOutput { return v.AuthType }).(pulumi.StringPtrOutput)
}

// Whether this control-plane can be used for cloud-gateways. Requires replacement if changed.
func (o GatewayControlPlaneOutput) CloudGateway() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GatewayControlPlane) pulumi.BoolPtrOutput { return v.CloudGateway }).(pulumi.BoolPtrOutput)
}

// The ClusterType value of the cluster associated with the Control Plane. Requires replacement if changed. ; must be one of ["CLUSTER*TYPE*CONTROL*PLANE", "CLUSTER*TYPE*HYBRID", "CLUSTER*TYPE*K8S*INGRESS*CONTROLLER", "CLUSTER*TYPE*CONTROL*PLANE*GROUP", "CLUSTER*TYPE_SERVERLESS"]
func (o GatewayControlPlaneOutput) ClusterType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayControlPlane) pulumi.StringPtrOutput { return v.ClusterType }).(pulumi.StringPtrOutput)
}

// CP configuration object for related access endpoints.
func (o GatewayControlPlaneOutput) Config() GatewayControlPlaneConfigOutput {
	return o.ApplyT(func(v *GatewayControlPlane) GatewayControlPlaneConfigOutput { return v.Config }).(GatewayControlPlaneConfigOutput)
}

// The description of the control plane in Konnect.
func (o GatewayControlPlaneOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayControlPlane) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
func (o GatewayControlPlaneOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GatewayControlPlane) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The name of the control plane.
func (o GatewayControlPlaneOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayControlPlane) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Array of proxy URLs associated with reaching the data-planes connected to a control-plane.
func (o GatewayControlPlaneOutput) ProxyUrls() GatewayControlPlaneProxyUrlArrayOutput {
	return o.ApplyT(func(v *GatewayControlPlane) GatewayControlPlaneProxyUrlArrayOutput { return v.ProxyUrls }).(GatewayControlPlaneProxyUrlArrayOutput)
}

type GatewayControlPlaneArrayOutput struct{ *pulumi.OutputState }

func (GatewayControlPlaneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayControlPlane)(nil)).Elem()
}

func (o GatewayControlPlaneArrayOutput) ToGatewayControlPlaneArrayOutput() GatewayControlPlaneArrayOutput {
	return o
}

func (o GatewayControlPlaneArrayOutput) ToGatewayControlPlaneArrayOutputWithContext(ctx context.Context) GatewayControlPlaneArrayOutput {
	return o
}

func (o GatewayControlPlaneArrayOutput) Index(i pulumi.IntInput) GatewayControlPlaneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GatewayControlPlane {
		return vs[0].([]*GatewayControlPlane)[vs[1].(int)]
	}).(GatewayControlPlaneOutput)
}

type GatewayControlPlaneMapOutput struct{ *pulumi.OutputState }

func (GatewayControlPlaneMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayControlPlane)(nil)).Elem()
}

func (o GatewayControlPlaneMapOutput) ToGatewayControlPlaneMapOutput() GatewayControlPlaneMapOutput {
	return o
}

func (o GatewayControlPlaneMapOutput) ToGatewayControlPlaneMapOutputWithContext(ctx context.Context) GatewayControlPlaneMapOutput {
	return o
}

func (o GatewayControlPlaneMapOutput) MapIndex(k pulumi.StringInput) GatewayControlPlaneOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GatewayControlPlane {
		return vs[0].(map[string]*GatewayControlPlane)[vs[1].(string)]
	}).(GatewayControlPlaneOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayControlPlaneInput)(nil)).Elem(), &GatewayControlPlane{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayControlPlaneArrayInput)(nil)).Elem(), GatewayControlPlaneArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayControlPlaneMapInput)(nil)).Elem(), GatewayControlPlaneMap{})
	pulumi.RegisterOutputType(GatewayControlPlaneOutput{})
	pulumi.RegisterOutputType(GatewayControlPlaneArrayOutput{})
	pulumi.RegisterOutputType(GatewayControlPlaneMapOutput{})
}
