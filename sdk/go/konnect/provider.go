// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package konnect

import (
	"context"
	"reflect"

	"github.com/bunchc/pulumi-tf-provider-konnect/sdk/go/konnect/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the konnect package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	KonnectAccessToken  pulumi.StringPtrOutput `pulumi:"konnectAccessToken"`
	PersonalAccessToken pulumi.StringPtrOutput `pulumi:"personalAccessToken"`
	// Server URL (defaults to https://global.api.konghq.com)
	ServerUrl                pulumi.StringPtrOutput `pulumi:"serverUrl"`
	SystemAccountAccessToken pulumi.StringPtrOutput `pulumi:"systemAccountAccessToken"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	if args.KonnectAccessToken != nil {
		args.KonnectAccessToken = pulumi.ToSecret(args.KonnectAccessToken).(pulumi.StringPtrInput)
	}
	if args.PersonalAccessToken != nil {
		args.PersonalAccessToken = pulumi.ToSecret(args.PersonalAccessToken).(pulumi.StringPtrInput)
	}
	if args.SystemAccountAccessToken != nil {
		args.SystemAccountAccessToken = pulumi.ToSecret(args.SystemAccountAccessToken).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"konnectAccessToken",
		"personalAccessToken",
		"systemAccountAccessToken",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:konnect", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	KonnectAccessToken  *string `pulumi:"konnectAccessToken"`
	PersonalAccessToken *string `pulumi:"personalAccessToken"`
	// Server URL (defaults to https://global.api.konghq.com)
	ServerUrl                *string `pulumi:"serverUrl"`
	SystemAccountAccessToken *string `pulumi:"systemAccountAccessToken"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	KonnectAccessToken  pulumi.StringPtrInput
	PersonalAccessToken pulumi.StringPtrInput
	// Server URL (defaults to https://global.api.konghq.com)
	ServerUrl                pulumi.StringPtrInput
	SystemAccountAccessToken pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func (o ProviderOutput) KonnectAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.KonnectAccessToken }).(pulumi.StringPtrOutput)
}

func (o ProviderOutput) PersonalAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.PersonalAccessToken }).(pulumi.StringPtrOutput)
}

// Server URL (defaults to https://global.api.konghq.com)
func (o ProviderOutput) ServerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ServerUrl }).(pulumi.StringPtrOutput)
}

func (o ProviderOutput) SystemAccountAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.SystemAccountAccessToken }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
