// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class ApplicationAuthStrategyOpenidConnectDcrProviderGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The display name of the DCR provider. This is used to identify the DCR provider in the Portal UI.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Contains a unique identifier used for this resource.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The type of DCR provider. Can be one of the following - auth0, azureAd, curity, okta, http. must be one of ["auth0", "azureAd", "curity", "okta", "http"]
        /// </summary>
        [Input("providerType")]
        public Input<string>? ProviderType { get; set; }

        public ApplicationAuthStrategyOpenidConnectDcrProviderGetArgs()
        {
        }
        public static new ApplicationAuthStrategyOpenidConnectDcrProviderGetArgs Empty => new ApplicationAuthStrategyOpenidConnectDcrProviderGetArgs();
    }
}