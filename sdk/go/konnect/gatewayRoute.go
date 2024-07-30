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

// GatewayRoute Resource
type GatewayRoute struct {
	pulumi.CustomResourceState

	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId pulumi.StringOutput `pulumi:"controlPlaneId"`
	// Unix epoch when the resource was created.
	CreatedAt pulumi.IntOutput `pulumi:"createdAt"`
	// A list of IP destinations of incoming connections that match this Route when using stream routing. Each entry is an object with fields "ip" (optionally in CIDR range notation) and/or "port".
	Destinations GatewayRouteDestinationArrayOutput `pulumi:"destinations"`
	// One or more lists of values indexed by header name that will cause this Route to match if present in the request. The `Host` header cannot be used with this attribute: hosts should be specified using the `hosts` attribute. When `headers` contains only one value and that value starts with the special prefix `~*`, the value is interpreted as a regular expression.
	Headers pulumi.StringMapOutput `pulumi:"headers"`
	// A list of domain names that match this Route. Note that the hosts value is case sensitive.
	Hosts pulumi.StringArrayOutput `pulumi:"hosts"`
	// The status code Kong responds with when all properties of a Route match except the protocol i.e. if the protocol of the request is `HTTP` instead of `HTTPS`. `Location` header is injected by Kong if the field is set to 301, 302, 307 or 308. Note: This config applies only if the Route is configured to only accept the `https` protocol. must be one of ["426", "301", "302", "307", "308"]
	HttpsRedirectStatusCode pulumi.IntOutput `pulumi:"httpsRedirectStatusCode"`
	// A list of HTTP methods that match this Route.
	Methods pulumi.StringArrayOutput `pulumi:"methods"`
	// The name of the Route. Route names must be unique, and they are case sensitive. For example, there can be two different Routes named "test" and "Test".
	Name pulumi.StringOutput `pulumi:"name"`
	// Controls how the Service path, Route path and requested path are combined when sending a request to the upstream. See above for a detailed description of each behavior. must be one of ["v0", "v1"]
	PathHandling pulumi.StringOutput `pulumi:"pathHandling"`
	// A list of paths that match this Route.
	Paths pulumi.StringArrayOutput `pulumi:"paths"`
	// When matching a Route via one of the `hosts` domain names, use the request `Host` header in the upstream request headers. If set to `false`, the upstream `Host` header will be that of the Service's `host`.
	PreserveHost pulumi.BoolOutput `pulumi:"preserveHost"`
	// An array of the protocols this Route should allow. See the Route Object section for a list of accepted protocols. When set to only `"https"`, HTTP requests are answered with an upgrade error. When set to only `"http"`, HTTPS requests are answered with an error.
	Protocols pulumi.StringArrayOutput `pulumi:"protocols"`
	// A number used to choose which route resolves a given request when several routes match it using regexes simultaneously. When two routes match the path and have the same `regexPriority`, the older one (lowest `createdAt`) is used. Note that the priority for non-regex routes is different (longer non-regex routes are matched before shorter ones).
	RegexPriority pulumi.IntOutput `pulumi:"regexPriority"`
	// Whether to enable request body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that receive data with chunked transfer encoding.
	RequestBuffering pulumi.BoolOutput `pulumi:"requestBuffering"`
	// Whether to enable response body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that send data with chunked transfer encoding.
	ResponseBuffering pulumi.BoolOutput `pulumi:"responseBuffering"`
	// The Service this Route is associated to. This is where the Route proxies traffic to.
	Service GatewayRouteServiceOutput `pulumi:"service"`
	// A list of SNIs that match this Route when using stream routing.
	Snis pulumi.StringArrayOutput `pulumi:"snis"`
	// A list of IP sources of incoming connections that match this Route when using stream routing. Each entry is an object with fields "ip" (optionally in CIDR range notation) and/or "port".
	Sources GatewayRouteSourceArrayOutput `pulumi:"sources"`
	// When matching a Route via one of the `paths`, strip the matching prefix from the upstream request URL.
	StripPath pulumi.BoolOutput `pulumi:"stripPath"`
	// An optional set of strings associated with the Route for grouping and filtering.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Unix epoch when the resource was last updated.
	UpdatedAt pulumi.IntOutput `pulumi:"updatedAt"`
}

// NewGatewayRoute registers a new resource with the given unique name, arguments, and options.
func NewGatewayRoute(ctx *pulumi.Context,
	name string, args *GatewayRouteArgs, opts ...pulumi.ResourceOption) (*GatewayRoute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ControlPlaneId == nil {
		return nil, errors.New("invalid value for required argument 'ControlPlaneId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GatewayRoute
	err := ctx.RegisterResource("konnect:index/gatewayRoute:GatewayRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayRoute gets an existing GatewayRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayRouteState, opts ...pulumi.ResourceOption) (*GatewayRoute, error) {
	var resource GatewayRoute
	err := ctx.ReadResource("konnect:index/gatewayRoute:GatewayRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayRoute resources.
type gatewayRouteState struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId *string `pulumi:"controlPlaneId"`
	// Unix epoch when the resource was created.
	CreatedAt *int `pulumi:"createdAt"`
	// A list of IP destinations of incoming connections that match this Route when using stream routing. Each entry is an object with fields "ip" (optionally in CIDR range notation) and/or "port".
	Destinations []GatewayRouteDestination `pulumi:"destinations"`
	// One or more lists of values indexed by header name that will cause this Route to match if present in the request. The `Host` header cannot be used with this attribute: hosts should be specified using the `hosts` attribute. When `headers` contains only one value and that value starts with the special prefix `~*`, the value is interpreted as a regular expression.
	Headers map[string]string `pulumi:"headers"`
	// A list of domain names that match this Route. Note that the hosts value is case sensitive.
	Hosts []string `pulumi:"hosts"`
	// The status code Kong responds with when all properties of a Route match except the protocol i.e. if the protocol of the request is `HTTP` instead of `HTTPS`. `Location` header is injected by Kong if the field is set to 301, 302, 307 or 308. Note: This config applies only if the Route is configured to only accept the `https` protocol. must be one of ["426", "301", "302", "307", "308"]
	HttpsRedirectStatusCode *int `pulumi:"httpsRedirectStatusCode"`
	// A list of HTTP methods that match this Route.
	Methods []string `pulumi:"methods"`
	// The name of the Route. Route names must be unique, and they are case sensitive. For example, there can be two different Routes named "test" and "Test".
	Name *string `pulumi:"name"`
	// Controls how the Service path, Route path and requested path are combined when sending a request to the upstream. See above for a detailed description of each behavior. must be one of ["v0", "v1"]
	PathHandling *string `pulumi:"pathHandling"`
	// A list of paths that match this Route.
	Paths []string `pulumi:"paths"`
	// When matching a Route via one of the `hosts` domain names, use the request `Host` header in the upstream request headers. If set to `false`, the upstream `Host` header will be that of the Service's `host`.
	PreserveHost *bool `pulumi:"preserveHost"`
	// An array of the protocols this Route should allow. See the Route Object section for a list of accepted protocols. When set to only `"https"`, HTTP requests are answered with an upgrade error. When set to only `"http"`, HTTPS requests are answered with an error.
	Protocols []string `pulumi:"protocols"`
	// A number used to choose which route resolves a given request when several routes match it using regexes simultaneously. When two routes match the path and have the same `regexPriority`, the older one (lowest `createdAt`) is used. Note that the priority for non-regex routes is different (longer non-regex routes are matched before shorter ones).
	RegexPriority *int `pulumi:"regexPriority"`
	// Whether to enable request body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that receive data with chunked transfer encoding.
	RequestBuffering *bool `pulumi:"requestBuffering"`
	// Whether to enable response body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that send data with chunked transfer encoding.
	ResponseBuffering *bool `pulumi:"responseBuffering"`
	// The Service this Route is associated to. This is where the Route proxies traffic to.
	Service *GatewayRouteService `pulumi:"service"`
	// A list of SNIs that match this Route when using stream routing.
	Snis []string `pulumi:"snis"`
	// A list of IP sources of incoming connections that match this Route when using stream routing. Each entry is an object with fields "ip" (optionally in CIDR range notation) and/or "port".
	Sources []GatewayRouteSource `pulumi:"sources"`
	// When matching a Route via one of the `paths`, strip the matching prefix from the upstream request URL.
	StripPath *bool `pulumi:"stripPath"`
	// An optional set of strings associated with the Route for grouping and filtering.
	Tags []string `pulumi:"tags"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int `pulumi:"updatedAt"`
}

type GatewayRouteState struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId pulumi.StringPtrInput
	// Unix epoch when the resource was created.
	CreatedAt pulumi.IntPtrInput
	// A list of IP destinations of incoming connections that match this Route when using stream routing. Each entry is an object with fields "ip" (optionally in CIDR range notation) and/or "port".
	Destinations GatewayRouteDestinationArrayInput
	// One or more lists of values indexed by header name that will cause this Route to match if present in the request. The `Host` header cannot be used with this attribute: hosts should be specified using the `hosts` attribute. When `headers` contains only one value and that value starts with the special prefix `~*`, the value is interpreted as a regular expression.
	Headers pulumi.StringMapInput
	// A list of domain names that match this Route. Note that the hosts value is case sensitive.
	Hosts pulumi.StringArrayInput
	// The status code Kong responds with when all properties of a Route match except the protocol i.e. if the protocol of the request is `HTTP` instead of `HTTPS`. `Location` header is injected by Kong if the field is set to 301, 302, 307 or 308. Note: This config applies only if the Route is configured to only accept the `https` protocol. must be one of ["426", "301", "302", "307", "308"]
	HttpsRedirectStatusCode pulumi.IntPtrInput
	// A list of HTTP methods that match this Route.
	Methods pulumi.StringArrayInput
	// The name of the Route. Route names must be unique, and they are case sensitive. For example, there can be two different Routes named "test" and "Test".
	Name pulumi.StringPtrInput
	// Controls how the Service path, Route path and requested path are combined when sending a request to the upstream. See above for a detailed description of each behavior. must be one of ["v0", "v1"]
	PathHandling pulumi.StringPtrInput
	// A list of paths that match this Route.
	Paths pulumi.StringArrayInput
	// When matching a Route via one of the `hosts` domain names, use the request `Host` header in the upstream request headers. If set to `false`, the upstream `Host` header will be that of the Service's `host`.
	PreserveHost pulumi.BoolPtrInput
	// An array of the protocols this Route should allow. See the Route Object section for a list of accepted protocols. When set to only `"https"`, HTTP requests are answered with an upgrade error. When set to only `"http"`, HTTPS requests are answered with an error.
	Protocols pulumi.StringArrayInput
	// A number used to choose which route resolves a given request when several routes match it using regexes simultaneously. When two routes match the path and have the same `regexPriority`, the older one (lowest `createdAt`) is used. Note that the priority for non-regex routes is different (longer non-regex routes are matched before shorter ones).
	RegexPriority pulumi.IntPtrInput
	// Whether to enable request body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that receive data with chunked transfer encoding.
	RequestBuffering pulumi.BoolPtrInput
	// Whether to enable response body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that send data with chunked transfer encoding.
	ResponseBuffering pulumi.BoolPtrInput
	// The Service this Route is associated to. This is where the Route proxies traffic to.
	Service GatewayRouteServicePtrInput
	// A list of SNIs that match this Route when using stream routing.
	Snis pulumi.StringArrayInput
	// A list of IP sources of incoming connections that match this Route when using stream routing. Each entry is an object with fields "ip" (optionally in CIDR range notation) and/or "port".
	Sources GatewayRouteSourceArrayInput
	// When matching a Route via one of the `paths`, strip the matching prefix from the upstream request URL.
	StripPath pulumi.BoolPtrInput
	// An optional set of strings associated with the Route for grouping and filtering.
	Tags pulumi.StringArrayInput
	// Unix epoch when the resource was last updated.
	UpdatedAt pulumi.IntPtrInput
}

func (GatewayRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayRouteState)(nil)).Elem()
}

type gatewayRouteArgs struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId string `pulumi:"controlPlaneId"`
	// A list of IP destinations of incoming connections that match this Route when using stream routing. Each entry is an object with fields "ip" (optionally in CIDR range notation) and/or "port".
	Destinations []GatewayRouteDestination `pulumi:"destinations"`
	// One or more lists of values indexed by header name that will cause this Route to match if present in the request. The `Host` header cannot be used with this attribute: hosts should be specified using the `hosts` attribute. When `headers` contains only one value and that value starts with the special prefix `~*`, the value is interpreted as a regular expression.
	Headers map[string]string `pulumi:"headers"`
	// A list of domain names that match this Route. Note that the hosts value is case sensitive.
	Hosts []string `pulumi:"hosts"`
	// The status code Kong responds with when all properties of a Route match except the protocol i.e. if the protocol of the request is `HTTP` instead of `HTTPS`. `Location` header is injected by Kong if the field is set to 301, 302, 307 or 308. Note: This config applies only if the Route is configured to only accept the `https` protocol. must be one of ["426", "301", "302", "307", "308"]
	HttpsRedirectStatusCode *int `pulumi:"httpsRedirectStatusCode"`
	// A list of HTTP methods that match this Route.
	Methods []string `pulumi:"methods"`
	// The name of the Route. Route names must be unique, and they are case sensitive. For example, there can be two different Routes named "test" and "Test".
	Name *string `pulumi:"name"`
	// Controls how the Service path, Route path and requested path are combined when sending a request to the upstream. See above for a detailed description of each behavior. must be one of ["v0", "v1"]
	PathHandling *string `pulumi:"pathHandling"`
	// A list of paths that match this Route.
	Paths []string `pulumi:"paths"`
	// When matching a Route via one of the `hosts` domain names, use the request `Host` header in the upstream request headers. If set to `false`, the upstream `Host` header will be that of the Service's `host`.
	PreserveHost *bool `pulumi:"preserveHost"`
	// An array of the protocols this Route should allow. See the Route Object section for a list of accepted protocols. When set to only `"https"`, HTTP requests are answered with an upgrade error. When set to only `"http"`, HTTPS requests are answered with an error.
	Protocols []string `pulumi:"protocols"`
	// A number used to choose which route resolves a given request when several routes match it using regexes simultaneously. When two routes match the path and have the same `regexPriority`, the older one (lowest `createdAt`) is used. Note that the priority for non-regex routes is different (longer non-regex routes are matched before shorter ones).
	RegexPriority *int `pulumi:"regexPriority"`
	// Whether to enable request body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that receive data with chunked transfer encoding.
	RequestBuffering *bool `pulumi:"requestBuffering"`
	// Whether to enable response body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that send data with chunked transfer encoding.
	ResponseBuffering *bool `pulumi:"responseBuffering"`
	// The Service this Route is associated to. This is where the Route proxies traffic to.
	Service *GatewayRouteService `pulumi:"service"`
	// A list of SNIs that match this Route when using stream routing.
	Snis []string `pulumi:"snis"`
	// A list of IP sources of incoming connections that match this Route when using stream routing. Each entry is an object with fields "ip" (optionally in CIDR range notation) and/or "port".
	Sources []GatewayRouteSource `pulumi:"sources"`
	// When matching a Route via one of the `paths`, strip the matching prefix from the upstream request URL.
	StripPath *bool `pulumi:"stripPath"`
	// An optional set of strings associated with the Route for grouping and filtering.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a GatewayRoute resource.
type GatewayRouteArgs struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneId pulumi.StringInput
	// A list of IP destinations of incoming connections that match this Route when using stream routing. Each entry is an object with fields "ip" (optionally in CIDR range notation) and/or "port".
	Destinations GatewayRouteDestinationArrayInput
	// One or more lists of values indexed by header name that will cause this Route to match if present in the request. The `Host` header cannot be used with this attribute: hosts should be specified using the `hosts` attribute. When `headers` contains only one value and that value starts with the special prefix `~*`, the value is interpreted as a regular expression.
	Headers pulumi.StringMapInput
	// A list of domain names that match this Route. Note that the hosts value is case sensitive.
	Hosts pulumi.StringArrayInput
	// The status code Kong responds with when all properties of a Route match except the protocol i.e. if the protocol of the request is `HTTP` instead of `HTTPS`. `Location` header is injected by Kong if the field is set to 301, 302, 307 or 308. Note: This config applies only if the Route is configured to only accept the `https` protocol. must be one of ["426", "301", "302", "307", "308"]
	HttpsRedirectStatusCode pulumi.IntPtrInput
	// A list of HTTP methods that match this Route.
	Methods pulumi.StringArrayInput
	// The name of the Route. Route names must be unique, and they are case sensitive. For example, there can be two different Routes named "test" and "Test".
	Name pulumi.StringPtrInput
	// Controls how the Service path, Route path and requested path are combined when sending a request to the upstream. See above for a detailed description of each behavior. must be one of ["v0", "v1"]
	PathHandling pulumi.StringPtrInput
	// A list of paths that match this Route.
	Paths pulumi.StringArrayInput
	// When matching a Route via one of the `hosts` domain names, use the request `Host` header in the upstream request headers. If set to `false`, the upstream `Host` header will be that of the Service's `host`.
	PreserveHost pulumi.BoolPtrInput
	// An array of the protocols this Route should allow. See the Route Object section for a list of accepted protocols. When set to only `"https"`, HTTP requests are answered with an upgrade error. When set to only `"http"`, HTTPS requests are answered with an error.
	Protocols pulumi.StringArrayInput
	// A number used to choose which route resolves a given request when several routes match it using regexes simultaneously. When two routes match the path and have the same `regexPriority`, the older one (lowest `createdAt`) is used. Note that the priority for non-regex routes is different (longer non-regex routes are matched before shorter ones).
	RegexPriority pulumi.IntPtrInput
	// Whether to enable request body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that receive data with chunked transfer encoding.
	RequestBuffering pulumi.BoolPtrInput
	// Whether to enable response body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that send data with chunked transfer encoding.
	ResponseBuffering pulumi.BoolPtrInput
	// The Service this Route is associated to. This is where the Route proxies traffic to.
	Service GatewayRouteServicePtrInput
	// A list of SNIs that match this Route when using stream routing.
	Snis pulumi.StringArrayInput
	// A list of IP sources of incoming connections that match this Route when using stream routing. Each entry is an object with fields "ip" (optionally in CIDR range notation) and/or "port".
	Sources GatewayRouteSourceArrayInput
	// When matching a Route via one of the `paths`, strip the matching prefix from the upstream request URL.
	StripPath pulumi.BoolPtrInput
	// An optional set of strings associated with the Route for grouping and filtering.
	Tags pulumi.StringArrayInput
}

func (GatewayRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayRouteArgs)(nil)).Elem()
}

type GatewayRouteInput interface {
	pulumi.Input

	ToGatewayRouteOutput() GatewayRouteOutput
	ToGatewayRouteOutputWithContext(ctx context.Context) GatewayRouteOutput
}

func (*GatewayRoute) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayRoute)(nil)).Elem()
}

func (i *GatewayRoute) ToGatewayRouteOutput() GatewayRouteOutput {
	return i.ToGatewayRouteOutputWithContext(context.Background())
}

func (i *GatewayRoute) ToGatewayRouteOutputWithContext(ctx context.Context) GatewayRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayRouteOutput)
}

// GatewayRouteArrayInput is an input type that accepts GatewayRouteArray and GatewayRouteArrayOutput values.
// You can construct a concrete instance of `GatewayRouteArrayInput` via:
//
//	GatewayRouteArray{ GatewayRouteArgs{...} }
type GatewayRouteArrayInput interface {
	pulumi.Input

	ToGatewayRouteArrayOutput() GatewayRouteArrayOutput
	ToGatewayRouteArrayOutputWithContext(context.Context) GatewayRouteArrayOutput
}

type GatewayRouteArray []GatewayRouteInput

func (GatewayRouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayRoute)(nil)).Elem()
}

func (i GatewayRouteArray) ToGatewayRouteArrayOutput() GatewayRouteArrayOutput {
	return i.ToGatewayRouteArrayOutputWithContext(context.Background())
}

func (i GatewayRouteArray) ToGatewayRouteArrayOutputWithContext(ctx context.Context) GatewayRouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayRouteArrayOutput)
}

// GatewayRouteMapInput is an input type that accepts GatewayRouteMap and GatewayRouteMapOutput values.
// You can construct a concrete instance of `GatewayRouteMapInput` via:
//
//	GatewayRouteMap{ "key": GatewayRouteArgs{...} }
type GatewayRouteMapInput interface {
	pulumi.Input

	ToGatewayRouteMapOutput() GatewayRouteMapOutput
	ToGatewayRouteMapOutputWithContext(context.Context) GatewayRouteMapOutput
}

type GatewayRouteMap map[string]GatewayRouteInput

func (GatewayRouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayRoute)(nil)).Elem()
}

func (i GatewayRouteMap) ToGatewayRouteMapOutput() GatewayRouteMapOutput {
	return i.ToGatewayRouteMapOutputWithContext(context.Background())
}

func (i GatewayRouteMap) ToGatewayRouteMapOutputWithContext(ctx context.Context) GatewayRouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayRouteMapOutput)
}

type GatewayRouteOutput struct{ *pulumi.OutputState }

func (GatewayRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayRoute)(nil)).Elem()
}

func (o GatewayRouteOutput) ToGatewayRouteOutput() GatewayRouteOutput {
	return o
}

func (o GatewayRouteOutput) ToGatewayRouteOutputWithContext(ctx context.Context) GatewayRouteOutput {
	return o
}

// The UUID of your control plane. This variable is available in the Konnect manager.
func (o GatewayRouteOutput) ControlPlaneId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringOutput { return v.ControlPlaneId }).(pulumi.StringOutput)
}

// Unix epoch when the resource was created.
func (o GatewayRouteOutput) CreatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.IntOutput { return v.CreatedAt }).(pulumi.IntOutput)
}

// A list of IP destinations of incoming connections that match this Route when using stream routing. Each entry is an object with fields "ip" (optionally in CIDR range notation) and/or "port".
func (o GatewayRouteOutput) Destinations() GatewayRouteDestinationArrayOutput {
	return o.ApplyT(func(v *GatewayRoute) GatewayRouteDestinationArrayOutput { return v.Destinations }).(GatewayRouteDestinationArrayOutput)
}

// One or more lists of values indexed by header name that will cause this Route to match if present in the request. The `Host` header cannot be used with this attribute: hosts should be specified using the `hosts` attribute. When `headers` contains only one value and that value starts with the special prefix `~*`, the value is interpreted as a regular expression.
func (o GatewayRouteOutput) Headers() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringMapOutput { return v.Headers }).(pulumi.StringMapOutput)
}

// A list of domain names that match this Route. Note that the hosts value is case sensitive.
func (o GatewayRouteOutput) Hosts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringArrayOutput { return v.Hosts }).(pulumi.StringArrayOutput)
}

// The status code Kong responds with when all properties of a Route match except the protocol i.e. if the protocol of the request is `HTTP` instead of `HTTPS`. `Location` header is injected by Kong if the field is set to 301, 302, 307 or 308. Note: This config applies only if the Route is configured to only accept the `https` protocol. must be one of ["426", "301", "302", "307", "308"]
func (o GatewayRouteOutput) HttpsRedirectStatusCode() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.IntOutput { return v.HttpsRedirectStatusCode }).(pulumi.IntOutput)
}

// A list of HTTP methods that match this Route.
func (o GatewayRouteOutput) Methods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringArrayOutput { return v.Methods }).(pulumi.StringArrayOutput)
}

// The name of the Route. Route names must be unique, and they are case sensitive. For example, there can be two different Routes named "test" and "Test".
func (o GatewayRouteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Controls how the Service path, Route path and requested path are combined when sending a request to the upstream. See above for a detailed description of each behavior. must be one of ["v0", "v1"]
func (o GatewayRouteOutput) PathHandling() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringOutput { return v.PathHandling }).(pulumi.StringOutput)
}

// A list of paths that match this Route.
func (o GatewayRouteOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringArrayOutput { return v.Paths }).(pulumi.StringArrayOutput)
}

// When matching a Route via one of the `hosts` domain names, use the request `Host` header in the upstream request headers. If set to `false`, the upstream `Host` header will be that of the Service's `host`.
func (o GatewayRouteOutput) PreserveHost() pulumi.BoolOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.BoolOutput { return v.PreserveHost }).(pulumi.BoolOutput)
}

// An array of the protocols this Route should allow. See the Route Object section for a list of accepted protocols. When set to only `"https"`, HTTP requests are answered with an upgrade error. When set to only `"http"`, HTTPS requests are answered with an error.
func (o GatewayRouteOutput) Protocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringArrayOutput { return v.Protocols }).(pulumi.StringArrayOutput)
}

// A number used to choose which route resolves a given request when several routes match it using regexes simultaneously. When two routes match the path and have the same `regexPriority`, the older one (lowest `createdAt`) is used. Note that the priority for non-regex routes is different (longer non-regex routes are matched before shorter ones).
func (o GatewayRouteOutput) RegexPriority() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.IntOutput { return v.RegexPriority }).(pulumi.IntOutput)
}

// Whether to enable request body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that receive data with chunked transfer encoding.
func (o GatewayRouteOutput) RequestBuffering() pulumi.BoolOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.BoolOutput { return v.RequestBuffering }).(pulumi.BoolOutput)
}

// Whether to enable response body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that send data with chunked transfer encoding.
func (o GatewayRouteOutput) ResponseBuffering() pulumi.BoolOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.BoolOutput { return v.ResponseBuffering }).(pulumi.BoolOutput)
}

// The Service this Route is associated to. This is where the Route proxies traffic to.
func (o GatewayRouteOutput) Service() GatewayRouteServiceOutput {
	return o.ApplyT(func(v *GatewayRoute) GatewayRouteServiceOutput { return v.Service }).(GatewayRouteServiceOutput)
}

// A list of SNIs that match this Route when using stream routing.
func (o GatewayRouteOutput) Snis() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringArrayOutput { return v.Snis }).(pulumi.StringArrayOutput)
}

// A list of IP sources of incoming connections that match this Route when using stream routing. Each entry is an object with fields "ip" (optionally in CIDR range notation) and/or "port".
func (o GatewayRouteOutput) Sources() GatewayRouteSourceArrayOutput {
	return o.ApplyT(func(v *GatewayRoute) GatewayRouteSourceArrayOutput { return v.Sources }).(GatewayRouteSourceArrayOutput)
}

// When matching a Route via one of the `paths`, strip the matching prefix from the upstream request URL.
func (o GatewayRouteOutput) StripPath() pulumi.BoolOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.BoolOutput { return v.StripPath }).(pulumi.BoolOutput)
}

// An optional set of strings associated with the Route for grouping and filtering.
func (o GatewayRouteOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Unix epoch when the resource was last updated.
func (o GatewayRouteOutput) UpdatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.IntOutput { return v.UpdatedAt }).(pulumi.IntOutput)
}

type GatewayRouteArrayOutput struct{ *pulumi.OutputState }

func (GatewayRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayRoute)(nil)).Elem()
}

func (o GatewayRouteArrayOutput) ToGatewayRouteArrayOutput() GatewayRouteArrayOutput {
	return o
}

func (o GatewayRouteArrayOutput) ToGatewayRouteArrayOutputWithContext(ctx context.Context) GatewayRouteArrayOutput {
	return o
}

func (o GatewayRouteArrayOutput) Index(i pulumi.IntInput) GatewayRouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GatewayRoute {
		return vs[0].([]*GatewayRoute)[vs[1].(int)]
	}).(GatewayRouteOutput)
}

type GatewayRouteMapOutput struct{ *pulumi.OutputState }

func (GatewayRouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayRoute)(nil)).Elem()
}

func (o GatewayRouteMapOutput) ToGatewayRouteMapOutput() GatewayRouteMapOutput {
	return o
}

func (o GatewayRouteMapOutput) ToGatewayRouteMapOutputWithContext(ctx context.Context) GatewayRouteMapOutput {
	return o
}

func (o GatewayRouteMapOutput) MapIndex(k pulumi.StringInput) GatewayRouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GatewayRoute {
		return vs[0].(map[string]*GatewayRoute)[vs[1].(string)]
	}).(GatewayRouteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayRouteInput)(nil)).Elem(), &GatewayRoute{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayRouteArrayInput)(nil)).Elem(), GatewayRouteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayRouteMapInput)(nil)).Elem(), GatewayRouteMap{})
	pulumi.RegisterOutputType(GatewayRouteOutput{})
	pulumi.RegisterOutputType(GatewayRouteArrayOutput{})
	pulumi.RegisterOutputType(GatewayRouteMapOutput{})
}