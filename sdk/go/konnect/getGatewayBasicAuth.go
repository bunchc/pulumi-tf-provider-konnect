// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package konnect

import (
	"context"
	"reflect"

	"github.com/bunchc/pulumi-tf-provider-konnect/sdk/go/konnect/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGatewayBasicAuth(ctx *pulumi.Context, args *LookupGatewayBasicAuthArgs, opts ...pulumi.InvokeOption) (*LookupGatewayBasicAuthResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGatewayBasicAuthResult
	err := ctx.Invoke("konnect:index/getGatewayBasicAuth:getGatewayBasicAuth", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGatewayBasicAuth.
type LookupGatewayBasicAuthArgs struct {
	ConsumerId     string `pulumi:"consumerId"`
	ControlPlaneId string `pulumi:"controlPlaneId"`
}

// A collection of values returned by getGatewayBasicAuth.
type LookupGatewayBasicAuthResult struct {
	Consumer       GetGatewayBasicAuthConsumer `pulumi:"consumer"`
	ConsumerId     string                      `pulumi:"consumerId"`
	ControlPlaneId string                      `pulumi:"controlPlaneId"`
	CreatedAt      int                         `pulumi:"createdAt"`
	Id             string                      `pulumi:"id"`
	Tags           []string                    `pulumi:"tags"`
	Username       string                      `pulumi:"username"`
}

func LookupGatewayBasicAuthOutput(ctx *pulumi.Context, args LookupGatewayBasicAuthOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayBasicAuthResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayBasicAuthResult, error) {
			args := v.(LookupGatewayBasicAuthArgs)
			r, err := LookupGatewayBasicAuth(ctx, &args, opts...)
			var s LookupGatewayBasicAuthResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayBasicAuthResultOutput)
}

// A collection of arguments for invoking getGatewayBasicAuth.
type LookupGatewayBasicAuthOutputArgs struct {
	ConsumerId     pulumi.StringInput `pulumi:"consumerId"`
	ControlPlaneId pulumi.StringInput `pulumi:"controlPlaneId"`
}

func (LookupGatewayBasicAuthOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayBasicAuthArgs)(nil)).Elem()
}

// A collection of values returned by getGatewayBasicAuth.
type LookupGatewayBasicAuthResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayBasicAuthResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayBasicAuthResult)(nil)).Elem()
}

func (o LookupGatewayBasicAuthResultOutput) ToLookupGatewayBasicAuthResultOutput() LookupGatewayBasicAuthResultOutput {
	return o
}

func (o LookupGatewayBasicAuthResultOutput) ToLookupGatewayBasicAuthResultOutputWithContext(ctx context.Context) LookupGatewayBasicAuthResultOutput {
	return o
}

func (o LookupGatewayBasicAuthResultOutput) Consumer() GetGatewayBasicAuthConsumerOutput {
	return o.ApplyT(func(v LookupGatewayBasicAuthResult) GetGatewayBasicAuthConsumer { return v.Consumer }).(GetGatewayBasicAuthConsumerOutput)
}

func (o LookupGatewayBasicAuthResultOutput) ConsumerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayBasicAuthResult) string { return v.ConsumerId }).(pulumi.StringOutput)
}

func (o LookupGatewayBasicAuthResultOutput) ControlPlaneId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayBasicAuthResult) string { return v.ControlPlaneId }).(pulumi.StringOutput)
}

func (o LookupGatewayBasicAuthResultOutput) CreatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGatewayBasicAuthResult) int { return v.CreatedAt }).(pulumi.IntOutput)
}

func (o LookupGatewayBasicAuthResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayBasicAuthResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGatewayBasicAuthResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGatewayBasicAuthResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupGatewayBasicAuthResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayBasicAuthResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayBasicAuthResultOutput{})
}