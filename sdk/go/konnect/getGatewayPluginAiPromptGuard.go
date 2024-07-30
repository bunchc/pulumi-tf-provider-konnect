// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package konnect

import (
	"context"
	"reflect"

	"github.com/bunchc/pulumi-tf-provider-konnect/sdk/go/konnect/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGatewayPluginAiPromptGuard(ctx *pulumi.Context, args *LookupGatewayPluginAiPromptGuardArgs, opts ...pulumi.InvokeOption) (*LookupGatewayPluginAiPromptGuardResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGatewayPluginAiPromptGuardResult
	err := ctx.Invoke("konnect:index/getGatewayPluginAiPromptGuard:getGatewayPluginAiPromptGuard", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGatewayPluginAiPromptGuard.
type LookupGatewayPluginAiPromptGuardArgs struct {
	ControlPlaneId string `pulumi:"controlPlaneId"`
}

// A collection of values returned by getGatewayPluginAiPromptGuard.
type LookupGatewayPluginAiPromptGuardResult struct {
	Config         GetGatewayPluginAiPromptGuardConfig        `pulumi:"config"`
	Consumer       GetGatewayPluginAiPromptGuardConsumer      `pulumi:"consumer"`
	ConsumerGroup  GetGatewayPluginAiPromptGuardConsumerGroup `pulumi:"consumerGroup"`
	ControlPlaneId string                                     `pulumi:"controlPlaneId"`
	CreatedAt      int                                        `pulumi:"createdAt"`
	Enabled        bool                                       `pulumi:"enabled"`
	Id             string                                     `pulumi:"id"`
	InstanceName   string                                     `pulumi:"instanceName"`
	Protocols      []string                                   `pulumi:"protocols"`
	Route          GetGatewayPluginAiPromptGuardRoute         `pulumi:"route"`
	Service        GetGatewayPluginAiPromptGuardService       `pulumi:"service"`
	Tags           []string                                   `pulumi:"tags"`
	UpdatedAt      int                                        `pulumi:"updatedAt"`
}

func LookupGatewayPluginAiPromptGuardOutput(ctx *pulumi.Context, args LookupGatewayPluginAiPromptGuardOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayPluginAiPromptGuardResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayPluginAiPromptGuardResult, error) {
			args := v.(LookupGatewayPluginAiPromptGuardArgs)
			r, err := LookupGatewayPluginAiPromptGuard(ctx, &args, opts...)
			var s LookupGatewayPluginAiPromptGuardResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayPluginAiPromptGuardResultOutput)
}

// A collection of arguments for invoking getGatewayPluginAiPromptGuard.
type LookupGatewayPluginAiPromptGuardOutputArgs struct {
	ControlPlaneId pulumi.StringInput `pulumi:"controlPlaneId"`
}

func (LookupGatewayPluginAiPromptGuardOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayPluginAiPromptGuardArgs)(nil)).Elem()
}

// A collection of values returned by getGatewayPluginAiPromptGuard.
type LookupGatewayPluginAiPromptGuardResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayPluginAiPromptGuardResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayPluginAiPromptGuardResult)(nil)).Elem()
}

func (o LookupGatewayPluginAiPromptGuardResultOutput) ToLookupGatewayPluginAiPromptGuardResultOutput() LookupGatewayPluginAiPromptGuardResultOutput {
	return o
}

func (o LookupGatewayPluginAiPromptGuardResultOutput) ToLookupGatewayPluginAiPromptGuardResultOutputWithContext(ctx context.Context) LookupGatewayPluginAiPromptGuardResultOutput {
	return o
}

func (o LookupGatewayPluginAiPromptGuardResultOutput) Config() GetGatewayPluginAiPromptGuardConfigOutput {
	return o.ApplyT(func(v LookupGatewayPluginAiPromptGuardResult) GetGatewayPluginAiPromptGuardConfig { return v.Config }).(GetGatewayPluginAiPromptGuardConfigOutput)
}

func (o LookupGatewayPluginAiPromptGuardResultOutput) Consumer() GetGatewayPluginAiPromptGuardConsumerOutput {
	return o.ApplyT(func(v LookupGatewayPluginAiPromptGuardResult) GetGatewayPluginAiPromptGuardConsumer {
		return v.Consumer
	}).(GetGatewayPluginAiPromptGuardConsumerOutput)
}

func (o LookupGatewayPluginAiPromptGuardResultOutput) ConsumerGroup() GetGatewayPluginAiPromptGuardConsumerGroupOutput {
	return o.ApplyT(func(v LookupGatewayPluginAiPromptGuardResult) GetGatewayPluginAiPromptGuardConsumerGroup {
		return v.ConsumerGroup
	}).(GetGatewayPluginAiPromptGuardConsumerGroupOutput)
}

func (o LookupGatewayPluginAiPromptGuardResultOutput) ControlPlaneId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayPluginAiPromptGuardResult) string { return v.ControlPlaneId }).(pulumi.StringOutput)
}

func (o LookupGatewayPluginAiPromptGuardResultOutput) CreatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGatewayPluginAiPromptGuardResult) int { return v.CreatedAt }).(pulumi.IntOutput)
}

func (o LookupGatewayPluginAiPromptGuardResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGatewayPluginAiPromptGuardResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupGatewayPluginAiPromptGuardResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayPluginAiPromptGuardResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGatewayPluginAiPromptGuardResultOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayPluginAiPromptGuardResult) string { return v.InstanceName }).(pulumi.StringOutput)
}

func (o LookupGatewayPluginAiPromptGuardResultOutput) Protocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGatewayPluginAiPromptGuardResult) []string { return v.Protocols }).(pulumi.StringArrayOutput)
}

func (o LookupGatewayPluginAiPromptGuardResultOutput) Route() GetGatewayPluginAiPromptGuardRouteOutput {
	return o.ApplyT(func(v LookupGatewayPluginAiPromptGuardResult) GetGatewayPluginAiPromptGuardRoute { return v.Route }).(GetGatewayPluginAiPromptGuardRouteOutput)
}

func (o LookupGatewayPluginAiPromptGuardResultOutput) Service() GetGatewayPluginAiPromptGuardServiceOutput {
	return o.ApplyT(func(v LookupGatewayPluginAiPromptGuardResult) GetGatewayPluginAiPromptGuardService { return v.Service }).(GetGatewayPluginAiPromptGuardServiceOutput)
}

func (o LookupGatewayPluginAiPromptGuardResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGatewayPluginAiPromptGuardResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupGatewayPluginAiPromptGuardResultOutput) UpdatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGatewayPluginAiPromptGuardResult) int { return v.UpdatedAt }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayPluginAiPromptGuardResultOutput{})
}