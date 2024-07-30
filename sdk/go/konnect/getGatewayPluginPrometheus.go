// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package konnect

import (
	"context"
	"reflect"

	"github.com/bunchc/pulumi-tf-provider-konnect/sdk/go/konnect/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGatewayPluginPrometheus(ctx *pulumi.Context, args *LookupGatewayPluginPrometheusArgs, opts ...pulumi.InvokeOption) (*LookupGatewayPluginPrometheusResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGatewayPluginPrometheusResult
	err := ctx.Invoke("konnect:index/getGatewayPluginPrometheus:getGatewayPluginPrometheus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGatewayPluginPrometheus.
type LookupGatewayPluginPrometheusArgs struct {
	ControlPlaneId string `pulumi:"controlPlaneId"`
}

// A collection of values returned by getGatewayPluginPrometheus.
type LookupGatewayPluginPrometheusResult struct {
	Config         GetGatewayPluginPrometheusConfig        `pulumi:"config"`
	Consumer       GetGatewayPluginPrometheusConsumer      `pulumi:"consumer"`
	ConsumerGroup  GetGatewayPluginPrometheusConsumerGroup `pulumi:"consumerGroup"`
	ControlPlaneId string                                  `pulumi:"controlPlaneId"`
	CreatedAt      int                                     `pulumi:"createdAt"`
	Enabled        bool                                    `pulumi:"enabled"`
	Id             string                                  `pulumi:"id"`
	InstanceName   string                                  `pulumi:"instanceName"`
	Protocols      []string                                `pulumi:"protocols"`
	Route          GetGatewayPluginPrometheusRoute         `pulumi:"route"`
	Service        GetGatewayPluginPrometheusService       `pulumi:"service"`
	Tags           []string                                `pulumi:"tags"`
	UpdatedAt      int                                     `pulumi:"updatedAt"`
}

func LookupGatewayPluginPrometheusOutput(ctx *pulumi.Context, args LookupGatewayPluginPrometheusOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayPluginPrometheusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayPluginPrometheusResult, error) {
			args := v.(LookupGatewayPluginPrometheusArgs)
			r, err := LookupGatewayPluginPrometheus(ctx, &args, opts...)
			var s LookupGatewayPluginPrometheusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayPluginPrometheusResultOutput)
}

// A collection of arguments for invoking getGatewayPluginPrometheus.
type LookupGatewayPluginPrometheusOutputArgs struct {
	ControlPlaneId pulumi.StringInput `pulumi:"controlPlaneId"`
}

func (LookupGatewayPluginPrometheusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayPluginPrometheusArgs)(nil)).Elem()
}

// A collection of values returned by getGatewayPluginPrometheus.
type LookupGatewayPluginPrometheusResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayPluginPrometheusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayPluginPrometheusResult)(nil)).Elem()
}

func (o LookupGatewayPluginPrometheusResultOutput) ToLookupGatewayPluginPrometheusResultOutput() LookupGatewayPluginPrometheusResultOutput {
	return o
}

func (o LookupGatewayPluginPrometheusResultOutput) ToLookupGatewayPluginPrometheusResultOutputWithContext(ctx context.Context) LookupGatewayPluginPrometheusResultOutput {
	return o
}

func (o LookupGatewayPluginPrometheusResultOutput) Config() GetGatewayPluginPrometheusConfigOutput {
	return o.ApplyT(func(v LookupGatewayPluginPrometheusResult) GetGatewayPluginPrometheusConfig { return v.Config }).(GetGatewayPluginPrometheusConfigOutput)
}

func (o LookupGatewayPluginPrometheusResultOutput) Consumer() GetGatewayPluginPrometheusConsumerOutput {
	return o.ApplyT(func(v LookupGatewayPluginPrometheusResult) GetGatewayPluginPrometheusConsumer { return v.Consumer }).(GetGatewayPluginPrometheusConsumerOutput)
}

func (o LookupGatewayPluginPrometheusResultOutput) ConsumerGroup() GetGatewayPluginPrometheusConsumerGroupOutput {
	return o.ApplyT(func(v LookupGatewayPluginPrometheusResult) GetGatewayPluginPrometheusConsumerGroup {
		return v.ConsumerGroup
	}).(GetGatewayPluginPrometheusConsumerGroupOutput)
}

func (o LookupGatewayPluginPrometheusResultOutput) ControlPlaneId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayPluginPrometheusResult) string { return v.ControlPlaneId }).(pulumi.StringOutput)
}

func (o LookupGatewayPluginPrometheusResultOutput) CreatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGatewayPluginPrometheusResult) int { return v.CreatedAt }).(pulumi.IntOutput)
}

func (o LookupGatewayPluginPrometheusResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGatewayPluginPrometheusResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupGatewayPluginPrometheusResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayPluginPrometheusResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGatewayPluginPrometheusResultOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayPluginPrometheusResult) string { return v.InstanceName }).(pulumi.StringOutput)
}

func (o LookupGatewayPluginPrometheusResultOutput) Protocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGatewayPluginPrometheusResult) []string { return v.Protocols }).(pulumi.StringArrayOutput)
}

func (o LookupGatewayPluginPrometheusResultOutput) Route() GetGatewayPluginPrometheusRouteOutput {
	return o.ApplyT(func(v LookupGatewayPluginPrometheusResult) GetGatewayPluginPrometheusRoute { return v.Route }).(GetGatewayPluginPrometheusRouteOutput)
}

func (o LookupGatewayPluginPrometheusResultOutput) Service() GetGatewayPluginPrometheusServiceOutput {
	return o.ApplyT(func(v LookupGatewayPluginPrometheusResult) GetGatewayPluginPrometheusService { return v.Service }).(GetGatewayPluginPrometheusServiceOutput)
}

func (o LookupGatewayPluginPrometheusResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGatewayPluginPrometheusResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupGatewayPluginPrometheusResultOutput) UpdatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGatewayPluginPrometheusResult) int { return v.UpdatedAt }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayPluginPrometheusResultOutput{})
}