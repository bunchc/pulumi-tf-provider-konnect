// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package konnect

import (
	"context"
	"reflect"

	"github.com/bunchc/pulumi-tf-provider-konnect/sdk/go/konnect/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGatewayConsumerGroup(ctx *pulumi.Context, args *LookupGatewayConsumerGroupArgs, opts ...pulumi.InvokeOption) (*LookupGatewayConsumerGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGatewayConsumerGroupResult
	err := ctx.Invoke("konnect:index/getGatewayConsumerGroup:getGatewayConsumerGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGatewayConsumerGroup.
type LookupGatewayConsumerGroupArgs struct {
	ControlPlaneId string `pulumi:"controlPlaneId"`
}

// A collection of values returned by getGatewayConsumerGroup.
type LookupGatewayConsumerGroupResult struct {
	ControlPlaneId string   `pulumi:"controlPlaneId"`
	CreatedAt      int      `pulumi:"createdAt"`
	Id             string   `pulumi:"id"`
	Name           string   `pulumi:"name"`
	Tags           []string `pulumi:"tags"`
	UpdatedAt      int      `pulumi:"updatedAt"`
}

func LookupGatewayConsumerGroupOutput(ctx *pulumi.Context, args LookupGatewayConsumerGroupOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayConsumerGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayConsumerGroupResult, error) {
			args := v.(LookupGatewayConsumerGroupArgs)
			r, err := LookupGatewayConsumerGroup(ctx, &args, opts...)
			var s LookupGatewayConsumerGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayConsumerGroupResultOutput)
}

// A collection of arguments for invoking getGatewayConsumerGroup.
type LookupGatewayConsumerGroupOutputArgs struct {
	ControlPlaneId pulumi.StringInput `pulumi:"controlPlaneId"`
}

func (LookupGatewayConsumerGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayConsumerGroupArgs)(nil)).Elem()
}

// A collection of values returned by getGatewayConsumerGroup.
type LookupGatewayConsumerGroupResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayConsumerGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayConsumerGroupResult)(nil)).Elem()
}

func (o LookupGatewayConsumerGroupResultOutput) ToLookupGatewayConsumerGroupResultOutput() LookupGatewayConsumerGroupResultOutput {
	return o
}

func (o LookupGatewayConsumerGroupResultOutput) ToLookupGatewayConsumerGroupResultOutputWithContext(ctx context.Context) LookupGatewayConsumerGroupResultOutput {
	return o
}

func (o LookupGatewayConsumerGroupResultOutput) ControlPlaneId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayConsumerGroupResult) string { return v.ControlPlaneId }).(pulumi.StringOutput)
}

func (o LookupGatewayConsumerGroupResultOutput) CreatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGatewayConsumerGroupResult) int { return v.CreatedAt }).(pulumi.IntOutput)
}

func (o LookupGatewayConsumerGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayConsumerGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGatewayConsumerGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayConsumerGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGatewayConsumerGroupResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGatewayConsumerGroupResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupGatewayConsumerGroupResultOutput) UpdatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGatewayConsumerGroupResult) int { return v.UpdatedAt }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayConsumerGroupResultOutput{})
}
