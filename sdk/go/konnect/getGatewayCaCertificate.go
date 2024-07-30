// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package konnect

import (
	"context"
	"reflect"

	"github.com/bunchc/pulumi-tf-provider-konnect/sdk/go/konnect/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGatewayCaCertificate(ctx *pulumi.Context, args *LookupGatewayCaCertificateArgs, opts ...pulumi.InvokeOption) (*LookupGatewayCaCertificateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGatewayCaCertificateResult
	err := ctx.Invoke("konnect:index/getGatewayCaCertificate:getGatewayCaCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGatewayCaCertificate.
type LookupGatewayCaCertificateArgs struct {
	ControlPlaneId string `pulumi:"controlPlaneId"`
}

// A collection of values returned by getGatewayCaCertificate.
type LookupGatewayCaCertificateResult struct {
	Cert           string   `pulumi:"cert"`
	CertDigest     string   `pulumi:"certDigest"`
	ControlPlaneId string   `pulumi:"controlPlaneId"`
	CreatedAt      int      `pulumi:"createdAt"`
	Id             string   `pulumi:"id"`
	Tags           []string `pulumi:"tags"`
	UpdatedAt      int      `pulumi:"updatedAt"`
}

func LookupGatewayCaCertificateOutput(ctx *pulumi.Context, args LookupGatewayCaCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayCaCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayCaCertificateResult, error) {
			args := v.(LookupGatewayCaCertificateArgs)
			r, err := LookupGatewayCaCertificate(ctx, &args, opts...)
			var s LookupGatewayCaCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayCaCertificateResultOutput)
}

// A collection of arguments for invoking getGatewayCaCertificate.
type LookupGatewayCaCertificateOutputArgs struct {
	ControlPlaneId pulumi.StringInput `pulumi:"controlPlaneId"`
}

func (LookupGatewayCaCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayCaCertificateArgs)(nil)).Elem()
}

// A collection of values returned by getGatewayCaCertificate.
type LookupGatewayCaCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayCaCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayCaCertificateResult)(nil)).Elem()
}

func (o LookupGatewayCaCertificateResultOutput) ToLookupGatewayCaCertificateResultOutput() LookupGatewayCaCertificateResultOutput {
	return o
}

func (o LookupGatewayCaCertificateResultOutput) ToLookupGatewayCaCertificateResultOutputWithContext(ctx context.Context) LookupGatewayCaCertificateResultOutput {
	return o
}

func (o LookupGatewayCaCertificateResultOutput) Cert() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayCaCertificateResult) string { return v.Cert }).(pulumi.StringOutput)
}

func (o LookupGatewayCaCertificateResultOutput) CertDigest() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayCaCertificateResult) string { return v.CertDigest }).(pulumi.StringOutput)
}

func (o LookupGatewayCaCertificateResultOutput) ControlPlaneId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayCaCertificateResult) string { return v.ControlPlaneId }).(pulumi.StringOutput)
}

func (o LookupGatewayCaCertificateResultOutput) CreatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGatewayCaCertificateResult) int { return v.CreatedAt }).(pulumi.IntOutput)
}

func (o LookupGatewayCaCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayCaCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGatewayCaCertificateResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGatewayCaCertificateResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupGatewayCaCertificateResultOutput) UpdatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGatewayCaCertificateResult) int { return v.UpdatedAt }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayCaCertificateResultOutput{})
}