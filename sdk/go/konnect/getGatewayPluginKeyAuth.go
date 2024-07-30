// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package konnect

import (
	"context"
	"reflect"

	"github.com/bunchc/pulumi-tf-provider-konnect/sdk/go/konnect/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGatewayPluginKeyAuth(ctx *pulumi.Context, args *LookupGatewayPluginKeyAuthArgs, opts ...pulumi.InvokeOption) (*LookupGatewayPluginKeyAuthResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGatewayPluginKeyAuthResult
	err := ctx.Invoke("konnect:index/getGatewayPluginKeyAuth:getGatewayPluginKeyAuth", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGatewayPluginKeyAuth.
type LookupGatewayPluginKeyAuthArgs struct {
	ControlPlaneId string `pulumi:"controlPlaneId"`
}

// A collection of values returned by getGatewayPluginKeyAuth.
type LookupGatewayPluginKeyAuthResult struct {
	Config         GetGatewayPluginKeyAuthConfig        `pulumi:"config"`
	Consumer       GetGatewayPluginKeyAuthConsumer      `pulumi:"consumer"`
	ConsumerGroup  GetGatewayPluginKeyAuthConsumerGroup `pulumi:"consumerGroup"`
	ControlPlaneId string                               `pulumi:"controlPlaneId"`
	CreatedAt      int                                  `pulumi:"createdAt"`
	Enabled        bool                                 `pulumi:"enabled"`
	Id             string                               `pulumi:"id"`
	InstanceName   string                               `pulumi:"instanceName"`
	Protocols      []string                             `pulumi:"protocols"`
	Route          GetGatewayPluginKeyAuthRoute         `pulumi:"route"`
	Service        GetGatewayPluginKeyAuthService       `pulumi:"service"`
	Tags           []string                             `pulumi:"tags"`
	UpdatedAt      int                                  `pulumi:"updatedAt"`
}

func LookupGatewayPluginKeyAuthOutput(ctx *pulumi.Context, args LookupGatewayPluginKeyAuthOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayPluginKeyAuthResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayPluginKeyAuthResult, error) {
			args := v.(LookupGatewayPluginKeyAuthArgs)
			r, err := LookupGatewayPluginKeyAuth(ctx, &args, opts...)
			var s LookupGatewayPluginKeyAuthResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayPluginKeyAuthResultOutput)
}

// A collection of arguments for invoking getGatewayPluginKeyAuth.
type LookupGatewayPluginKeyAuthOutputArgs struct {
	ControlPlaneId pulumi.StringInput `pulumi:"controlPlaneId"`
}

func (LookupGatewayPluginKeyAuthOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayPluginKeyAuthArgs)(nil)).Elem()
}

// A collection of values returned by getGatewayPluginKeyAuth.
type LookupGatewayPluginKeyAuthResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayPluginKeyAuthResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayPluginKeyAuthResult)(nil)).Elem()
}

func (o LookupGatewayPluginKeyAuthResultOutput) ToLookupGatewayPluginKeyAuthResultOutput() LookupGatewayPluginKeyAuthResultOutput {
	return o
}

func (o LookupGatewayPluginKeyAuthResultOutput) ToLookupGatewayPluginKeyAuthResultOutputWithContext(ctx context.Context) LookupGatewayPluginKeyAuthResultOutput {
	return o
}

func (o LookupGatewayPluginKeyAuthResultOutput) Config() GetGatewayPluginKeyAuthConfigOutput {
	return o.ApplyT(func(v LookupGatewayPluginKeyAuthResult) GetGatewayPluginKeyAuthConfig { return v.Config }).(GetGatewayPluginKeyAuthConfigOutput)
}

func (o LookupGatewayPluginKeyAuthResultOutput) Consumer() GetGatewayPluginKeyAuthConsumerOutput {
	return o.ApplyT(func(v LookupGatewayPluginKeyAuthResult) GetGatewayPluginKeyAuthConsumer { return v.Consumer }).(GetGatewayPluginKeyAuthConsumerOutput)
}

func (o LookupGatewayPluginKeyAuthResultOutput) ConsumerGroup() GetGatewayPluginKeyAuthConsumerGroupOutput {
	return o.ApplyT(func(v LookupGatewayPluginKeyAuthResult) GetGatewayPluginKeyAuthConsumerGroup { return v.ConsumerGroup }).(GetGatewayPluginKeyAuthConsumerGroupOutput)
}

func (o LookupGatewayPluginKeyAuthResultOutput) ControlPlaneId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayPluginKeyAuthResult) string { return v.ControlPlaneId }).(pulumi.StringOutput)
}

func (o LookupGatewayPluginKeyAuthResultOutput) CreatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGatewayPluginKeyAuthResult) int { return v.CreatedAt }).(pulumi.IntOutput)
}

func (o LookupGatewayPluginKeyAuthResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGatewayPluginKeyAuthResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupGatewayPluginKeyAuthResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayPluginKeyAuthResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGatewayPluginKeyAuthResultOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayPluginKeyAuthResult) string { return v.InstanceName }).(pulumi.StringOutput)
}

func (o LookupGatewayPluginKeyAuthResultOutput) Protocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGatewayPluginKeyAuthResult) []string { return v.Protocols }).(pulumi.StringArrayOutput)
}

func (o LookupGatewayPluginKeyAuthResultOutput) Route() GetGatewayPluginKeyAuthRouteOutput {
	return o.ApplyT(func(v LookupGatewayPluginKeyAuthResult) GetGatewayPluginKeyAuthRoute { return v.Route }).(GetGatewayPluginKeyAuthRouteOutput)
}

func (o LookupGatewayPluginKeyAuthResultOutput) Service() GetGatewayPluginKeyAuthServiceOutput {
	return o.ApplyT(func(v LookupGatewayPluginKeyAuthResult) GetGatewayPluginKeyAuthService { return v.Service }).(GetGatewayPluginKeyAuthServiceOutput)
}

func (o LookupGatewayPluginKeyAuthResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGatewayPluginKeyAuthResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupGatewayPluginKeyAuthResultOutput) UpdatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGatewayPluginKeyAuthResult) int { return v.UpdatedAt }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayPluginKeyAuthResultOutput{})
}
