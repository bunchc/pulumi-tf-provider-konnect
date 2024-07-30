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
    public sealed class ApplicationAuthStrategyOpenidConnectConfigs
    {
        /// <summary>
        /// A more advanced mode to configure an API Product Version’s Application Auth Strategy.
        /// Using this mode will allow developers to use API credentials issued from an external IdP that will authenticate their application requests.
        /// Once authenticated, an application will be granted access to any Product Version it is registered for that is configured for the same Auth Strategy.
        /// An OIDC strategy may be used in conjunction with a DCR provider to automatically create the IdP application.
        /// </summary>
        public readonly Outputs.ApplicationAuthStrategyOpenidConnectConfigsOpenidConnect? OpenidConnect;

        [OutputConstructor]
        private ApplicationAuthStrategyOpenidConnectConfigs(Outputs.ApplicationAuthStrategyOpenidConnectConfigsOpenidConnect? openidConnect)
        {
            OpenidConnect = openidConnect;
        }
    }
}
