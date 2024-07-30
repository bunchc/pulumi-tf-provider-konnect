// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Outputs
{

    [OutputType]
    public sealed class ApplicationAuthStrategyOpenidConnectDcrProvider
    {
        /// <summary>
        /// The display name of the DCR provider. This is used to identify the DCR provider in the Portal UI.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// Contains a unique identifier used for this resource.
        /// </summary>
        public readonly string? Id;
        public readonly string? Name;
        /// <summary>
        /// The type of DCR provider. Can be one of the following - auth0, azureAd, curity, okta, http. must be one of ["auth0", "azureAd", "curity", "okta", "http"]
        /// </summary>
        public readonly string? ProviderType;

        [OutputConstructor]
        private ApplicationAuthStrategyOpenidConnectDcrProvider(
            string? displayName,

            string? id,

            string? name,

            string? providerType)
        {
            DisplayName = displayName;
            Id = id;
            Name = name;
            ProviderType = providerType;
        }
    }
}
