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

// GatewayUpstream Resource
type GatewayUpstream struct {
	pulumi.CustomResourceState

	// Which load balancing algorithm to use. must be one of ["consistent-hashing", "least-connections", "round-robin", "latency"]
	Algorithm pulumi.StringOutput `pulumi:"algorithm"`
	// If set, the certificate to be used as client certificate while TLS handshaking to the upstream server.
	ClientCertificate GatewayUpstreamClientCertificateOutput `pulumi:"clientCertificate"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId pulumi.StringOutput `pulumi:"controlPlaneId"`
	// Unix epoch when the resource was created.
	CreatedAt pulumi.IntOutput `pulumi:"createdAt"`
	// What to use as hashing input if the primary `hashOn` does not return a hash (eg. header is missing, or no Consumer identified). Not available if `hashOn` is set to `cookie`. must be one of ["none", "consumer", "ip", "header", "cookie", "path", "query*arg", "uri*capture"]
	HashFallback pulumi.StringOutput `pulumi:"hashFallback"`
	// The header name to take the value from as hash input. Only required when `hashFallback` is set to `header`.
	HashFallbackHeader pulumi.StringOutput `pulumi:"hashFallbackHeader"`
	// The name of the query string argument to take the value from as hash input. Only required when `hashFallback` is set to `queryArg`.
	HashFallbackQueryArg pulumi.StringOutput `pulumi:"hashFallbackQueryArg"`
	// The name of the route URI capture to take the value from as hash input. Only required when `hashFallback` is set to `uriCapture`.
	HashFallbackUriCapture pulumi.StringOutput `pulumi:"hashFallbackUriCapture"`
	// What to use as hashing input. Using `none` results in a weighted-round-robin scheme with no hashing. must be one of ["none", "consumer", "ip", "header", "cookie", "path", "query*arg", "uri*capture"]
	HashOn pulumi.StringOutput `pulumi:"hashOn"`
	// The cookie name to take the value from as hash input. Only required when `hashOn` or `hashFallback` is set to `cookie`. If the specified cookie is not in the request, Kong will generate a value and set the cookie in the response.
	HashOnCookie pulumi.StringOutput `pulumi:"hashOnCookie"`
	// The cookie path to set in the response headers. Only required when `hashOn` or `hashFallback` is set to `cookie`.
	HashOnCookiePath pulumi.StringOutput `pulumi:"hashOnCookiePath"`
	// The header name to take the value from as hash input. Only required when `hashOn` is set to `header`.
	HashOnHeader pulumi.StringOutput `pulumi:"hashOnHeader"`
	// The name of the query string argument to take the value from as hash input. Only required when `hashOn` is set to `queryArg`.
	HashOnQueryArg pulumi.StringOutput `pulumi:"hashOnQueryArg"`
	// The name of the route URI capture to take the value from as hash input. Only required when `hashOn` is set to `uriCapture`.
	HashOnUriCapture pulumi.StringOutput               `pulumi:"hashOnUriCapture"`
	Healthchecks     GatewayUpstreamHealthchecksOutput `pulumi:"healthchecks"`
	// The hostname to be used as `Host` header when proxying requests through Kong.
	HostHeader pulumi.StringOutput `pulumi:"hostHeader"`
	// This is a hostname, which must be equal to the `host` of a Service.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of slots in the load balancer algorithm. If `algorithm` is set to `round-robin`, this setting determines the maximum number of slots. If `algorithm` is set to `consistent-hashing`, this setting determines the actual number of slots in the algorithm. Accepts an integer in the range `10`-`65536`.
	Slots pulumi.IntOutput `pulumi:"slots"`
	// An optional set of strings associated with the Upstream for grouping and filtering.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Unix epoch when the resource was last updated.
	UpdatedAt pulumi.IntOutput `pulumi:"updatedAt"`
	// If set, the balancer will use SRV hostname(if DNS Answer has SRV record) as the proxy upstream `Host`.
	UseSrvName pulumi.BoolOutput `pulumi:"useSrvName"`
}

// NewGatewayUpstream registers a new resource with the given unique name, arguments, and options.
func NewGatewayUpstream(ctx *pulumi.Context,
	name string, args *GatewayUpstreamArgs, opts ...pulumi.ResourceOption) (*GatewayUpstream, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ControlPlaneId == nil {
		return nil, errors.New("invalid value for required argument 'ControlPlaneId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GatewayUpstream
	err := ctx.RegisterResource("konnect:index/gatewayUpstream:GatewayUpstream", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayUpstream gets an existing GatewayUpstream resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayUpstream(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayUpstreamState, opts ...pulumi.ResourceOption) (*GatewayUpstream, error) {
	var resource GatewayUpstream
	err := ctx.ReadResource("konnect:index/gatewayUpstream:GatewayUpstream", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayUpstream resources.
type gatewayUpstreamState struct {
	// Which load balancing algorithm to use. must be one of ["consistent-hashing", "least-connections", "round-robin", "latency"]
	Algorithm *string `pulumi:"algorithm"`
	// If set, the certificate to be used as client certificate while TLS handshaking to the upstream server.
	ClientCertificate *GatewayUpstreamClientCertificate `pulumi:"clientCertificate"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId *string `pulumi:"controlPlaneId"`
	// Unix epoch when the resource was created.
	CreatedAt *int `pulumi:"createdAt"`
	// What to use as hashing input if the primary `hashOn` does not return a hash (eg. header is missing, or no Consumer identified). Not available if `hashOn` is set to `cookie`. must be one of ["none", "consumer", "ip", "header", "cookie", "path", "query*arg", "uri*capture"]
	HashFallback *string `pulumi:"hashFallback"`
	// The header name to take the value from as hash input. Only required when `hashFallback` is set to `header`.
	HashFallbackHeader *string `pulumi:"hashFallbackHeader"`
	// The name of the query string argument to take the value from as hash input. Only required when `hashFallback` is set to `queryArg`.
	HashFallbackQueryArg *string `pulumi:"hashFallbackQueryArg"`
	// The name of the route URI capture to take the value from as hash input. Only required when `hashFallback` is set to `uriCapture`.
	HashFallbackUriCapture *string `pulumi:"hashFallbackUriCapture"`
	// What to use as hashing input. Using `none` results in a weighted-round-robin scheme with no hashing. must be one of ["none", "consumer", "ip", "header", "cookie", "path", "query*arg", "uri*capture"]
	HashOn *string `pulumi:"hashOn"`
	// The cookie name to take the value from as hash input. Only required when `hashOn` or `hashFallback` is set to `cookie`. If the specified cookie is not in the request, Kong will generate a value and set the cookie in the response.
	HashOnCookie *string `pulumi:"hashOnCookie"`
	// The cookie path to set in the response headers. Only required when `hashOn` or `hashFallback` is set to `cookie`.
	HashOnCookiePath *string `pulumi:"hashOnCookiePath"`
	// The header name to take the value from as hash input. Only required when `hashOn` is set to `header`.
	HashOnHeader *string `pulumi:"hashOnHeader"`
	// The name of the query string argument to take the value from as hash input. Only required when `hashOn` is set to `queryArg`.
	HashOnQueryArg *string `pulumi:"hashOnQueryArg"`
	// The name of the route URI capture to take the value from as hash input. Only required when `hashOn` is set to `uriCapture`.
	HashOnUriCapture *string                      `pulumi:"hashOnUriCapture"`
	Healthchecks     *GatewayUpstreamHealthchecks `pulumi:"healthchecks"`
	// The hostname to be used as `Host` header when proxying requests through Kong.
	HostHeader *string `pulumi:"hostHeader"`
	// This is a hostname, which must be equal to the `host` of a Service.
	Name *string `pulumi:"name"`
	// The number of slots in the load balancer algorithm. If `algorithm` is set to `round-robin`, this setting determines the maximum number of slots. If `algorithm` is set to `consistent-hashing`, this setting determines the actual number of slots in the algorithm. Accepts an integer in the range `10`-`65536`.
	Slots *int `pulumi:"slots"`
	// An optional set of strings associated with the Upstream for grouping and filtering.
	Tags []string `pulumi:"tags"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int `pulumi:"updatedAt"`
	// If set, the balancer will use SRV hostname(if DNS Answer has SRV record) as the proxy upstream `Host`.
	UseSrvName *bool `pulumi:"useSrvName"`
}

type GatewayUpstreamState struct {
	// Which load balancing algorithm to use. must be one of ["consistent-hashing", "least-connections", "round-robin", "latency"]
	Algorithm pulumi.StringPtrInput
	// If set, the certificate to be used as client certificate while TLS handshaking to the upstream server.
	ClientCertificate GatewayUpstreamClientCertificatePtrInput
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId pulumi.StringPtrInput
	// Unix epoch when the resource was created.
	CreatedAt pulumi.IntPtrInput
	// What to use as hashing input if the primary `hashOn` does not return a hash (eg. header is missing, or no Consumer identified). Not available if `hashOn` is set to `cookie`. must be one of ["none", "consumer", "ip", "header", "cookie", "path", "query*arg", "uri*capture"]
	HashFallback pulumi.StringPtrInput
	// The header name to take the value from as hash input. Only required when `hashFallback` is set to `header`.
	HashFallbackHeader pulumi.StringPtrInput
	// The name of the query string argument to take the value from as hash input. Only required when `hashFallback` is set to `queryArg`.
	HashFallbackQueryArg pulumi.StringPtrInput
	// The name of the route URI capture to take the value from as hash input. Only required when `hashFallback` is set to `uriCapture`.
	HashFallbackUriCapture pulumi.StringPtrInput
	// What to use as hashing input. Using `none` results in a weighted-round-robin scheme with no hashing. must be one of ["none", "consumer", "ip", "header", "cookie", "path", "query*arg", "uri*capture"]
	HashOn pulumi.StringPtrInput
	// The cookie name to take the value from as hash input. Only required when `hashOn` or `hashFallback` is set to `cookie`. If the specified cookie is not in the request, Kong will generate a value and set the cookie in the response.
	HashOnCookie pulumi.StringPtrInput
	// The cookie path to set in the response headers. Only required when `hashOn` or `hashFallback` is set to `cookie`.
	HashOnCookiePath pulumi.StringPtrInput
	// The header name to take the value from as hash input. Only required when `hashOn` is set to `header`.
	HashOnHeader pulumi.StringPtrInput
	// The name of the query string argument to take the value from as hash input. Only required when `hashOn` is set to `queryArg`.
	HashOnQueryArg pulumi.StringPtrInput
	// The name of the route URI capture to take the value from as hash input. Only required when `hashOn` is set to `uriCapture`.
	HashOnUriCapture pulumi.StringPtrInput
	Healthchecks     GatewayUpstreamHealthchecksPtrInput
	// The hostname to be used as `Host` header when proxying requests through Kong.
	HostHeader pulumi.StringPtrInput
	// This is a hostname, which must be equal to the `host` of a Service.
	Name pulumi.StringPtrInput
	// The number of slots in the load balancer algorithm. If `algorithm` is set to `round-robin`, this setting determines the maximum number of slots. If `algorithm` is set to `consistent-hashing`, this setting determines the actual number of slots in the algorithm. Accepts an integer in the range `10`-`65536`.
	Slots pulumi.IntPtrInput
	// An optional set of strings associated with the Upstream for grouping and filtering.
	Tags pulumi.StringArrayInput
	// Unix epoch when the resource was last updated.
	UpdatedAt pulumi.IntPtrInput
	// If set, the balancer will use SRV hostname(if DNS Answer has SRV record) as the proxy upstream `Host`.
	UseSrvName pulumi.BoolPtrInput
}

func (GatewayUpstreamState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayUpstreamState)(nil)).Elem()
}

type gatewayUpstreamArgs struct {
	// Which load balancing algorithm to use. must be one of ["consistent-hashing", "least-connections", "round-robin", "latency"]
	Algorithm *string `pulumi:"algorithm"`
	// If set, the certificate to be used as client certificate while TLS handshaking to the upstream server.
	ClientCertificate *GatewayUpstreamClientCertificate `pulumi:"clientCertificate"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId string `pulumi:"controlPlaneId"`
	// What to use as hashing input if the primary `hashOn` does not return a hash (eg. header is missing, or no Consumer identified). Not available if `hashOn` is set to `cookie`. must be one of ["none", "consumer", "ip", "header", "cookie", "path", "query*arg", "uri*capture"]
	HashFallback *string `pulumi:"hashFallback"`
	// The header name to take the value from as hash input. Only required when `hashFallback` is set to `header`.
	HashFallbackHeader *string `pulumi:"hashFallbackHeader"`
	// The name of the query string argument to take the value from as hash input. Only required when `hashFallback` is set to `queryArg`.
	HashFallbackQueryArg *string `pulumi:"hashFallbackQueryArg"`
	// The name of the route URI capture to take the value from as hash input. Only required when `hashFallback` is set to `uriCapture`.
	HashFallbackUriCapture *string `pulumi:"hashFallbackUriCapture"`
	// What to use as hashing input. Using `none` results in a weighted-round-robin scheme with no hashing. must be one of ["none", "consumer", "ip", "header", "cookie", "path", "query*arg", "uri*capture"]
	HashOn *string `pulumi:"hashOn"`
	// The cookie name to take the value from as hash input. Only required when `hashOn` or `hashFallback` is set to `cookie`. If the specified cookie is not in the request, Kong will generate a value and set the cookie in the response.
	HashOnCookie *string `pulumi:"hashOnCookie"`
	// The cookie path to set in the response headers. Only required when `hashOn` or `hashFallback` is set to `cookie`.
	HashOnCookiePath *string `pulumi:"hashOnCookiePath"`
	// The header name to take the value from as hash input. Only required when `hashOn` is set to `header`.
	HashOnHeader *string `pulumi:"hashOnHeader"`
	// The name of the query string argument to take the value from as hash input. Only required when `hashOn` is set to `queryArg`.
	HashOnQueryArg *string `pulumi:"hashOnQueryArg"`
	// The name of the route URI capture to take the value from as hash input. Only required when `hashOn` is set to `uriCapture`.
	HashOnUriCapture *string                      `pulumi:"hashOnUriCapture"`
	Healthchecks     *GatewayUpstreamHealthchecks `pulumi:"healthchecks"`
	// The hostname to be used as `Host` header when proxying requests through Kong.
	HostHeader *string `pulumi:"hostHeader"`
	// This is a hostname, which must be equal to the `host` of a Service.
	Name *string `pulumi:"name"`
	// The number of slots in the load balancer algorithm. If `algorithm` is set to `round-robin`, this setting determines the maximum number of slots. If `algorithm` is set to `consistent-hashing`, this setting determines the actual number of slots in the algorithm. Accepts an integer in the range `10`-`65536`.
	Slots *int `pulumi:"slots"`
	// An optional set of strings associated with the Upstream for grouping and filtering.
	Tags []string `pulumi:"tags"`
	// If set, the balancer will use SRV hostname(if DNS Answer has SRV record) as the proxy upstream `Host`.
	UseSrvName *bool `pulumi:"useSrvName"`
}

// The set of arguments for constructing a GatewayUpstream resource.
type GatewayUpstreamArgs struct {
	// Which load balancing algorithm to use. must be one of ["consistent-hashing", "least-connections", "round-robin", "latency"]
	Algorithm pulumi.StringPtrInput
	// If set, the certificate to be used as client certificate while TLS handshaking to the upstream server.
	ClientCertificate GatewayUpstreamClientCertificatePtrInput
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId pulumi.StringInput
	// What to use as hashing input if the primary `hashOn` does not return a hash (eg. header is missing, or no Consumer identified). Not available if `hashOn` is set to `cookie`. must be one of ["none", "consumer", "ip", "header", "cookie", "path", "query*arg", "uri*capture"]
	HashFallback pulumi.StringPtrInput
	// The header name to take the value from as hash input. Only required when `hashFallback` is set to `header`.
	HashFallbackHeader pulumi.StringPtrInput
	// The name of the query string argument to take the value from as hash input. Only required when `hashFallback` is set to `queryArg`.
	HashFallbackQueryArg pulumi.StringPtrInput
	// The name of the route URI capture to take the value from as hash input. Only required when `hashFallback` is set to `uriCapture`.
	HashFallbackUriCapture pulumi.StringPtrInput
	// What to use as hashing input. Using `none` results in a weighted-round-robin scheme with no hashing. must be one of ["none", "consumer", "ip", "header", "cookie", "path", "query*arg", "uri*capture"]
	HashOn pulumi.StringPtrInput
	// The cookie name to take the value from as hash input. Only required when `hashOn` or `hashFallback` is set to `cookie`. If the specified cookie is not in the request, Kong will generate a value and set the cookie in the response.
	HashOnCookie pulumi.StringPtrInput
	// The cookie path to set in the response headers. Only required when `hashOn` or `hashFallback` is set to `cookie`.
	HashOnCookiePath pulumi.StringPtrInput
	// The header name to take the value from as hash input. Only required when `hashOn` is set to `header`.
	HashOnHeader pulumi.StringPtrInput
	// The name of the query string argument to take the value from as hash input. Only required when `hashOn` is set to `queryArg`.
	HashOnQueryArg pulumi.StringPtrInput
	// The name of the route URI capture to take the value from as hash input. Only required when `hashOn` is set to `uriCapture`.
	HashOnUriCapture pulumi.StringPtrInput
	Healthchecks     GatewayUpstreamHealthchecksPtrInput
	// The hostname to be used as `Host` header when proxying requests through Kong.
	HostHeader pulumi.StringPtrInput
	// This is a hostname, which must be equal to the `host` of a Service.
	Name pulumi.StringPtrInput
	// The number of slots in the load balancer algorithm. If `algorithm` is set to `round-robin`, this setting determines the maximum number of slots. If `algorithm` is set to `consistent-hashing`, this setting determines the actual number of slots in the algorithm. Accepts an integer in the range `10`-`65536`.
	Slots pulumi.IntPtrInput
	// An optional set of strings associated with the Upstream for grouping and filtering.
	Tags pulumi.StringArrayInput
	// If set, the balancer will use SRV hostname(if DNS Answer has SRV record) as the proxy upstream `Host`.
	UseSrvName pulumi.BoolPtrInput
}

func (GatewayUpstreamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayUpstreamArgs)(nil)).Elem()
}

type GatewayUpstreamInput interface {
	pulumi.Input

	ToGatewayUpstreamOutput() GatewayUpstreamOutput
	ToGatewayUpstreamOutputWithContext(ctx context.Context) GatewayUpstreamOutput
}

func (*GatewayUpstream) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayUpstream)(nil)).Elem()
}

func (i *GatewayUpstream) ToGatewayUpstreamOutput() GatewayUpstreamOutput {
	return i.ToGatewayUpstreamOutputWithContext(context.Background())
}

func (i *GatewayUpstream) ToGatewayUpstreamOutputWithContext(ctx context.Context) GatewayUpstreamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayUpstreamOutput)
}

// GatewayUpstreamArrayInput is an input type that accepts GatewayUpstreamArray and GatewayUpstreamArrayOutput values.
// You can construct a concrete instance of `GatewayUpstreamArrayInput` via:
//
//	GatewayUpstreamArray{ GatewayUpstreamArgs{...} }
type GatewayUpstreamArrayInput interface {
	pulumi.Input

	ToGatewayUpstreamArrayOutput() GatewayUpstreamArrayOutput
	ToGatewayUpstreamArrayOutputWithContext(context.Context) GatewayUpstreamArrayOutput
}

type GatewayUpstreamArray []GatewayUpstreamInput

func (GatewayUpstreamArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayUpstream)(nil)).Elem()
}

func (i GatewayUpstreamArray) ToGatewayUpstreamArrayOutput() GatewayUpstreamArrayOutput {
	return i.ToGatewayUpstreamArrayOutputWithContext(context.Background())
}

func (i GatewayUpstreamArray) ToGatewayUpstreamArrayOutputWithContext(ctx context.Context) GatewayUpstreamArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayUpstreamArrayOutput)
}

// GatewayUpstreamMapInput is an input type that accepts GatewayUpstreamMap and GatewayUpstreamMapOutput values.
// You can construct a concrete instance of `GatewayUpstreamMapInput` via:
//
//	GatewayUpstreamMap{ "key": GatewayUpstreamArgs{...} }
type GatewayUpstreamMapInput interface {
	pulumi.Input

	ToGatewayUpstreamMapOutput() GatewayUpstreamMapOutput
	ToGatewayUpstreamMapOutputWithContext(context.Context) GatewayUpstreamMapOutput
}

type GatewayUpstreamMap map[string]GatewayUpstreamInput

func (GatewayUpstreamMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayUpstream)(nil)).Elem()
}

func (i GatewayUpstreamMap) ToGatewayUpstreamMapOutput() GatewayUpstreamMapOutput {
	return i.ToGatewayUpstreamMapOutputWithContext(context.Background())
}

func (i GatewayUpstreamMap) ToGatewayUpstreamMapOutputWithContext(ctx context.Context) GatewayUpstreamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayUpstreamMapOutput)
}

type GatewayUpstreamOutput struct{ *pulumi.OutputState }

func (GatewayUpstreamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayUpstream)(nil)).Elem()
}

func (o GatewayUpstreamOutput) ToGatewayUpstreamOutput() GatewayUpstreamOutput {
	return o
}

func (o GatewayUpstreamOutput) ToGatewayUpstreamOutputWithContext(ctx context.Context) GatewayUpstreamOutput {
	return o
}

// Which load balancing algorithm to use. must be one of ["consistent-hashing", "least-connections", "round-robin", "latency"]
func (o GatewayUpstreamOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayUpstream) pulumi.StringOutput { return v.Algorithm }).(pulumi.StringOutput)
}

// If set, the certificate to be used as client certificate while TLS handshaking to the upstream server.
func (o GatewayUpstreamOutput) ClientCertificate() GatewayUpstreamClientCertificateOutput {
	return o.ApplyT(func(v *GatewayUpstream) GatewayUpstreamClientCertificateOutput { return v.ClientCertificate }).(GatewayUpstreamClientCertificateOutput)
}

// The UUID of your control plane. This variable is available in the Konnect manager.
func (o GatewayUpstreamOutput) ControlPlaneId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayUpstream) pulumi.StringOutput { return v.ControlPlaneId }).(pulumi.StringOutput)
}

// Unix epoch when the resource was created.
func (o GatewayUpstreamOutput) CreatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayUpstream) pulumi.IntOutput { return v.CreatedAt }).(pulumi.IntOutput)
}

// What to use as hashing input if the primary `hashOn` does not return a hash (eg. header is missing, or no Consumer identified). Not available if `hashOn` is set to `cookie`. must be one of ["none", "consumer", "ip", "header", "cookie", "path", "query*arg", "uri*capture"]
func (o GatewayUpstreamOutput) HashFallback() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayUpstream) pulumi.StringOutput { return v.HashFallback }).(pulumi.StringOutput)
}

// The header name to take the value from as hash input. Only required when `hashFallback` is set to `header`.
func (o GatewayUpstreamOutput) HashFallbackHeader() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayUpstream) pulumi.StringOutput { return v.HashFallbackHeader }).(pulumi.StringOutput)
}

// The name of the query string argument to take the value from as hash input. Only required when `hashFallback` is set to `queryArg`.
func (o GatewayUpstreamOutput) HashFallbackQueryArg() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayUpstream) pulumi.StringOutput { return v.HashFallbackQueryArg }).(pulumi.StringOutput)
}

// The name of the route URI capture to take the value from as hash input. Only required when `hashFallback` is set to `uriCapture`.
func (o GatewayUpstreamOutput) HashFallbackUriCapture() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayUpstream) pulumi.StringOutput { return v.HashFallbackUriCapture }).(pulumi.StringOutput)
}

// What to use as hashing input. Using `none` results in a weighted-round-robin scheme with no hashing. must be one of ["none", "consumer", "ip", "header", "cookie", "path", "query*arg", "uri*capture"]
func (o GatewayUpstreamOutput) HashOn() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayUpstream) pulumi.StringOutput { return v.HashOn }).(pulumi.StringOutput)
}

// The cookie name to take the value from as hash input. Only required when `hashOn` or `hashFallback` is set to `cookie`. If the specified cookie is not in the request, Kong will generate a value and set the cookie in the response.
func (o GatewayUpstreamOutput) HashOnCookie() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayUpstream) pulumi.StringOutput { return v.HashOnCookie }).(pulumi.StringOutput)
}

// The cookie path to set in the response headers. Only required when `hashOn` or `hashFallback` is set to `cookie`.
func (o GatewayUpstreamOutput) HashOnCookiePath() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayUpstream) pulumi.StringOutput { return v.HashOnCookiePath }).(pulumi.StringOutput)
}

// The header name to take the value from as hash input. Only required when `hashOn` is set to `header`.
func (o GatewayUpstreamOutput) HashOnHeader() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayUpstream) pulumi.StringOutput { return v.HashOnHeader }).(pulumi.StringOutput)
}

// The name of the query string argument to take the value from as hash input. Only required when `hashOn` is set to `queryArg`.
func (o GatewayUpstreamOutput) HashOnQueryArg() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayUpstream) pulumi.StringOutput { return v.HashOnQueryArg }).(pulumi.StringOutput)
}

// The name of the route URI capture to take the value from as hash input. Only required when `hashOn` is set to `uriCapture`.
func (o GatewayUpstreamOutput) HashOnUriCapture() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayUpstream) pulumi.StringOutput { return v.HashOnUriCapture }).(pulumi.StringOutput)
}

func (o GatewayUpstreamOutput) Healthchecks() GatewayUpstreamHealthchecksOutput {
	return o.ApplyT(func(v *GatewayUpstream) GatewayUpstreamHealthchecksOutput { return v.Healthchecks }).(GatewayUpstreamHealthchecksOutput)
}

// The hostname to be used as `Host` header when proxying requests through Kong.
func (o GatewayUpstreamOutput) HostHeader() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayUpstream) pulumi.StringOutput { return v.HostHeader }).(pulumi.StringOutput)
}

// This is a hostname, which must be equal to the `host` of a Service.
func (o GatewayUpstreamOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayUpstream) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The number of slots in the load balancer algorithm. If `algorithm` is set to `round-robin`, this setting determines the maximum number of slots. If `algorithm` is set to `consistent-hashing`, this setting determines the actual number of slots in the algorithm. Accepts an integer in the range `10`-`65536`.
func (o GatewayUpstreamOutput) Slots() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayUpstream) pulumi.IntOutput { return v.Slots }).(pulumi.IntOutput)
}

// An optional set of strings associated with the Upstream for grouping and filtering.
func (o GatewayUpstreamOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayUpstream) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Unix epoch when the resource was last updated.
func (o GatewayUpstreamOutput) UpdatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayUpstream) pulumi.IntOutput { return v.UpdatedAt }).(pulumi.IntOutput)
}

// If set, the balancer will use SRV hostname(if DNS Answer has SRV record) as the proxy upstream `Host`.
func (o GatewayUpstreamOutput) UseSrvName() pulumi.BoolOutput {
	return o.ApplyT(func(v *GatewayUpstream) pulumi.BoolOutput { return v.UseSrvName }).(pulumi.BoolOutput)
}

type GatewayUpstreamArrayOutput struct{ *pulumi.OutputState }

func (GatewayUpstreamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayUpstream)(nil)).Elem()
}

func (o GatewayUpstreamArrayOutput) ToGatewayUpstreamArrayOutput() GatewayUpstreamArrayOutput {
	return o
}

func (o GatewayUpstreamArrayOutput) ToGatewayUpstreamArrayOutputWithContext(ctx context.Context) GatewayUpstreamArrayOutput {
	return o
}

func (o GatewayUpstreamArrayOutput) Index(i pulumi.IntInput) GatewayUpstreamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GatewayUpstream {
		return vs[0].([]*GatewayUpstream)[vs[1].(int)]
	}).(GatewayUpstreamOutput)
}

type GatewayUpstreamMapOutput struct{ *pulumi.OutputState }

func (GatewayUpstreamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayUpstream)(nil)).Elem()
}

func (o GatewayUpstreamMapOutput) ToGatewayUpstreamMapOutput() GatewayUpstreamMapOutput {
	return o
}

func (o GatewayUpstreamMapOutput) ToGatewayUpstreamMapOutputWithContext(ctx context.Context) GatewayUpstreamMapOutput {
	return o
}

func (o GatewayUpstreamMapOutput) MapIndex(k pulumi.StringInput) GatewayUpstreamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GatewayUpstream {
		return vs[0].(map[string]*GatewayUpstream)[vs[1].(string)]
	}).(GatewayUpstreamOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayUpstreamInput)(nil)).Elem(), &GatewayUpstream{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayUpstreamArrayInput)(nil)).Elem(), GatewayUpstreamArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayUpstreamMapInput)(nil)).Elem(), GatewayUpstreamMap{})
	pulumi.RegisterOutputType(GatewayUpstreamOutput{})
	pulumi.RegisterOutputType(GatewayUpstreamArrayOutput{})
	pulumi.RegisterOutputType(GatewayUpstreamMapOutput{})
}
