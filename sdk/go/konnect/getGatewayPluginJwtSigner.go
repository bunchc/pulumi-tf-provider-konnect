// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package konnect

import (
	"context"
	"reflect"

	"github.com/bunchc/pulumi-tf-provider-konnect/sdk/go/konnect/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGatewayPluginJwtSigner(ctx *pulumi.Context, args *LookupGatewayPluginJwtSignerArgs, opts ...pulumi.InvokeOption) (*LookupGatewayPluginJwtSignerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGatewayPluginJwtSignerResult
	err := ctx.Invoke("konnect:index/getGatewayPluginJwtSigner:getGatewayPluginJwtSigner", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGatewayPluginJwtSigner.
type LookupGatewayPluginJwtSignerArgs struct {
	ControlPlaneId string `pulumi:"controlPlaneId"`
}

// A collection of values returned by getGatewayPluginJwtSigner.
type LookupGatewayPluginJwtSignerResult struct {
	Config         GetGatewayPluginJwtSignerConfig        `pulumi:"config"`
	Consumer       GetGatewayPluginJwtSignerConsumer      `pulumi:"consumer"`
	ConsumerGroup  GetGatewayPluginJwtSignerConsumerGroup `pulumi:"consumerGroup"`
	ControlPlaneId string                                 `pulumi:"controlPlaneId"`
	CreatedAt      int                                    `pulumi:"createdAt"`
	Enabled        bool                                   `pulumi:"enabled"`
	Id             string                                 `pulumi:"id"`
	InstanceName   string                                 `pulumi:"instanceName"`
	Protocols      []string                               `pulumi:"protocols"`
	Route          GetGatewayPluginJwtSignerRoute         `pulumi:"route"`
	Service        GetGatewayPluginJwtSignerService       `pulumi:"service"`
	Tags           []string                               `pulumi:"tags"`
	UpdatedAt      int                                    `pulumi:"updatedAt"`
}

func LookupGatewayPluginJwtSignerOutput(ctx *pulumi.Context, args LookupGatewayPluginJwtSignerOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayPluginJwtSignerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayPluginJwtSignerResult, error) {
			args := v.(LookupGatewayPluginJwtSignerArgs)
			r, err := LookupGatewayPluginJwtSigner(ctx, &args, opts...)
			var s LookupGatewayPluginJwtSignerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayPluginJwtSignerResultOutput)
}

// A collection of arguments for invoking getGatewayPluginJwtSigner.
type LookupGatewayPluginJwtSignerOutputArgs struct {
	ControlPlaneId pulumi.StringInput `pulumi:"controlPlaneId"`
}

func (LookupGatewayPluginJwtSignerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayPluginJwtSignerArgs)(nil)).Elem()
}

// A collection of values returned by getGatewayPluginJwtSigner.
type LookupGatewayPluginJwtSignerResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayPluginJwtSignerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayPluginJwtSignerResult)(nil)).Elem()
}

func (o LookupGatewayPluginJwtSignerResultOutput) ToLookupGatewayPluginJwtSignerResultOutput() LookupGatewayPluginJwtSignerResultOutput {
	return o
}

func (o LookupGatewayPluginJwtSignerResultOutput) ToLookupGatewayPluginJwtSignerResultOutputWithContext(ctx context.Context) LookupGatewayPluginJwtSignerResultOutput {
	return o
}

func (o LookupGatewayPluginJwtSignerResultOutput) Config() GetGatewayPluginJwtSignerConfigOutput {
	return o.ApplyT(func(v LookupGatewayPluginJwtSignerResult) GetGatewayPluginJwtSignerConfig { return v.Config }).(GetGatewayPluginJwtSignerConfigOutput)
}

func (o LookupGatewayPluginJwtSignerResultOutput) Consumer() GetGatewayPluginJwtSignerConsumerOutput {
	return o.ApplyT(func(v LookupGatewayPluginJwtSignerResult) GetGatewayPluginJwtSignerConsumer { return v.Consumer }).(GetGatewayPluginJwtSignerConsumerOutput)
}

func (o LookupGatewayPluginJwtSignerResultOutput) ConsumerGroup() GetGatewayPluginJwtSignerConsumerGroupOutput {
	return o.ApplyT(func(v LookupGatewayPluginJwtSignerResult) GetGatewayPluginJwtSignerConsumerGroup {
		return v.ConsumerGroup
	}).(GetGatewayPluginJwtSignerConsumerGroupOutput)
}

func (o LookupGatewayPluginJwtSignerResultOutput) ControlPlaneId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayPluginJwtSignerResult) string { return v.ControlPlaneId }).(pulumi.StringOutput)
}

func (o LookupGatewayPluginJwtSignerResultOutput) CreatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGatewayPluginJwtSignerResult) int { return v.CreatedAt }).(pulumi.IntOutput)
}

func (o LookupGatewayPluginJwtSignerResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGatewayPluginJwtSignerResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupGatewayPluginJwtSignerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayPluginJwtSignerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGatewayPluginJwtSignerResultOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayPluginJwtSignerResult) string { return v.InstanceName }).(pulumi.StringOutput)
}

func (o LookupGatewayPluginJwtSignerResultOutput) Protocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGatewayPluginJwtSignerResult) []string { return v.Protocols }).(pulumi.StringArrayOutput)
}

func (o LookupGatewayPluginJwtSignerResultOutput) Route() GetGatewayPluginJwtSignerRouteOutput {
	return o.ApplyT(func(v LookupGatewayPluginJwtSignerResult) GetGatewayPluginJwtSignerRoute { return v.Route }).(GetGatewayPluginJwtSignerRouteOutput)
}

func (o LookupGatewayPluginJwtSignerResultOutput) Service() GetGatewayPluginJwtSignerServiceOutput {
	return o.ApplyT(func(v LookupGatewayPluginJwtSignerResult) GetGatewayPluginJwtSignerService { return v.Service }).(GetGatewayPluginJwtSignerServiceOutput)
}

func (o LookupGatewayPluginJwtSignerResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGatewayPluginJwtSignerResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupGatewayPluginJwtSignerResultOutput) UpdatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGatewayPluginJwtSignerResult) int { return v.UpdatedAt }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayPluginJwtSignerResultOutput{})
}
