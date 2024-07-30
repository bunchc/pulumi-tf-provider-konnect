// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package konnect

import (
	"context"
	"reflect"

	"github.com/bunchc/pulumi-tf-provider-konnect/sdk/go/konnect/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplicationAuthStrategy(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*LookupApplicationAuthStrategyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupApplicationAuthStrategyResult
	err := ctx.Invoke("konnect:index/getApplicationAuthStrategy:getApplicationAuthStrategy", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getApplicationAuthStrategy.
type LookupApplicationAuthStrategyResult struct {
	Active        bool                                    `pulumi:"active"`
	DisplayName   string                                  `pulumi:"displayName"`
	Id            string                                  `pulumi:"id"`
	KeyAuth       GetApplicationAuthStrategyKeyAuth       `pulumi:"keyAuth"`
	Name          string                                  `pulumi:"name"`
	OpenidConnect GetApplicationAuthStrategyOpenidConnect `pulumi:"openidConnect"`
}

func LookupApplicationAuthStrategyOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) LookupApplicationAuthStrategyResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (LookupApplicationAuthStrategyResult, error) {
		r, err := LookupApplicationAuthStrategy(ctx, opts...)
		var s LookupApplicationAuthStrategyResult
		if r != nil {
			s = *r
		}
		return s, err
	}).(LookupApplicationAuthStrategyResultOutput)
}

// A collection of values returned by getApplicationAuthStrategy.
type LookupApplicationAuthStrategyResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationAuthStrategyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationAuthStrategyResult)(nil)).Elem()
}

func (o LookupApplicationAuthStrategyResultOutput) ToLookupApplicationAuthStrategyResultOutput() LookupApplicationAuthStrategyResultOutput {
	return o
}

func (o LookupApplicationAuthStrategyResultOutput) ToLookupApplicationAuthStrategyResultOutputWithContext(ctx context.Context) LookupApplicationAuthStrategyResultOutput {
	return o
}

func (o LookupApplicationAuthStrategyResultOutput) Active() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApplicationAuthStrategyResult) bool { return v.Active }).(pulumi.BoolOutput)
}

func (o LookupApplicationAuthStrategyResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationAuthStrategyResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupApplicationAuthStrategyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationAuthStrategyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApplicationAuthStrategyResultOutput) KeyAuth() GetApplicationAuthStrategyKeyAuthOutput {
	return o.ApplyT(func(v LookupApplicationAuthStrategyResult) GetApplicationAuthStrategyKeyAuth { return v.KeyAuth }).(GetApplicationAuthStrategyKeyAuthOutput)
}

func (o LookupApplicationAuthStrategyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationAuthStrategyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApplicationAuthStrategyResultOutput) OpenidConnect() GetApplicationAuthStrategyOpenidConnectOutput {
	return o.ApplyT(func(v LookupApplicationAuthStrategyResult) GetApplicationAuthStrategyOpenidConnect {
		return v.OpenidConnect
	}).(GetApplicationAuthStrategyOpenidConnectOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationAuthStrategyResultOutput{})
}
