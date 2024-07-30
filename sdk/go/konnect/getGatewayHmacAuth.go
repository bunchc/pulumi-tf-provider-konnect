// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package konnect

import (
	"context"
	"reflect"

	"github.com/bunchc/pulumi-tf-provider-konnect/sdk/go/konnect/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGatewayHmacAuth(ctx *pulumi.Context, args *LookupGatewayHmacAuthArgs, opts ...pulumi.InvokeOption) (*LookupGatewayHmacAuthResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGatewayHmacAuthResult
	err := ctx.Invoke("konnect:index/getGatewayHmacAuth:getGatewayHmacAuth", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGatewayHmacAuth.
type LookupGatewayHmacAuthArgs struct {
	ConsumerId     string `pulumi:"consumerId"`
	ControlPlaneId string `pulumi:"controlPlaneId"`
}

// A collection of values returned by getGatewayHmacAuth.
type LookupGatewayHmacAuthResult struct {
	Consumer       GetGatewayHmacAuthConsumer `pulumi:"consumer"`
	ConsumerId     string                     `pulumi:"consumerId"`
	ControlPlaneId string                     `pulumi:"controlPlaneId"`
	CreatedAt      int                        `pulumi:"createdAt"`
	Id             string                     `pulumi:"id"`
	Secret         string                     `pulumi:"secret"`
	Tags           []string                   `pulumi:"tags"`
	Username       string                     `pulumi:"username"`
}

func LookupGatewayHmacAuthOutput(ctx *pulumi.Context, args LookupGatewayHmacAuthOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayHmacAuthResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayHmacAuthResult, error) {
			args := v.(LookupGatewayHmacAuthArgs)
			r, err := LookupGatewayHmacAuth(ctx, &args, opts...)
			var s LookupGatewayHmacAuthResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayHmacAuthResultOutput)
}

// A collection of arguments for invoking getGatewayHmacAuth.
type LookupGatewayHmacAuthOutputArgs struct {
	ConsumerId     pulumi.StringInput `pulumi:"consumerId"`
	ControlPlaneId pulumi.StringInput `pulumi:"controlPlaneId"`
}

func (LookupGatewayHmacAuthOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayHmacAuthArgs)(nil)).Elem()
}

// A collection of values returned by getGatewayHmacAuth.
type LookupGatewayHmacAuthResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayHmacAuthResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayHmacAuthResult)(nil)).Elem()
}

func (o LookupGatewayHmacAuthResultOutput) ToLookupGatewayHmacAuthResultOutput() LookupGatewayHmacAuthResultOutput {
	return o
}

func (o LookupGatewayHmacAuthResultOutput) ToLookupGatewayHmacAuthResultOutputWithContext(ctx context.Context) LookupGatewayHmacAuthResultOutput {
	return o
}

func (o LookupGatewayHmacAuthResultOutput) Consumer() GetGatewayHmacAuthConsumerOutput {
	return o.ApplyT(func(v LookupGatewayHmacAuthResult) GetGatewayHmacAuthConsumer { return v.Consumer }).(GetGatewayHmacAuthConsumerOutput)
}

func (o LookupGatewayHmacAuthResultOutput) ConsumerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayHmacAuthResult) string { return v.ConsumerId }).(pulumi.StringOutput)
}

func (o LookupGatewayHmacAuthResultOutput) ControlPlaneId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayHmacAuthResult) string { return v.ControlPlaneId }).(pulumi.StringOutput)
}

func (o LookupGatewayHmacAuthResultOutput) CreatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGatewayHmacAuthResult) int { return v.CreatedAt }).(pulumi.IntOutput)
}

func (o LookupGatewayHmacAuthResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayHmacAuthResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGatewayHmacAuthResultOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayHmacAuthResult) string { return v.Secret }).(pulumi.StringOutput)
}

func (o LookupGatewayHmacAuthResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGatewayHmacAuthResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupGatewayHmacAuthResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayHmacAuthResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayHmacAuthResultOutput{})
}
