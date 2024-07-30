// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect
{
    public static class GetApplicationAuthStrategy
    {
        public static Task<GetApplicationAuthStrategyResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApplicationAuthStrategyResult>("konnect:index/getApplicationAuthStrategy:getApplicationAuthStrategy", InvokeArgs.Empty, options.WithDefaults());

        public static Output<GetApplicationAuthStrategyResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApplicationAuthStrategyResult>("konnect:index/getApplicationAuthStrategy:getApplicationAuthStrategy", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetApplicationAuthStrategyResult
    {
        public readonly bool Active;
        public readonly string DisplayName;
        public readonly string Id;
        public readonly Outputs.GetApplicationAuthStrategyKeyAuthResult KeyAuth;
        public readonly string Name;
        public readonly Outputs.GetApplicationAuthStrategyOpenidConnectResult OpenidConnect;

        [OutputConstructor]
        private GetApplicationAuthStrategyResult(
            bool active,

            string displayName,

            string id,

            Outputs.GetApplicationAuthStrategyKeyAuthResult keyAuth,

            string name,

            Outputs.GetApplicationAuthStrategyOpenidConnectResult openidConnect)
        {
            Active = active;
            DisplayName = displayName;
            Id = id;
            KeyAuth = keyAuth;
            Name = name;
            OpenidConnect = openidConnect;
        }
    }
}
