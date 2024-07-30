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
    public sealed class GetGatewayPluginOauth2ConfigResult
    {
        /// <summary>
        /// Accepts HTTPs requests that have already been terminated by a proxy or load balancer.
        /// </summary>
        public readonly bool AcceptHttpIfAlreadyTerminated;
        /// <summary>
        /// An optional string (consumer UUID or username) value to use as an “anonymous” consumer if authentication fails.
        /// </summary>
        public readonly string Anonymous;
        /// <summary>
        /// The name of the header that is supposed to carry the access token.
        /// </summary>
        public readonly string AuthHeaderName;
        /// <summary>
        /// An optional boolean value to enable the three-legged Authorization Code flow (RFC 6742 Section 4.1).
        /// </summary>
        public readonly bool EnableAuthorizationCode;
        /// <summary>
        /// An optional boolean value to enable the Client Credentials Grant flow (RFC 6742 Section 4.4).
        /// </summary>
        public readonly bool EnableClientCredentials;
        /// <summary>
        /// An optional boolean value to enable the Implicit Grant flow which allows to provision a token as a result of the authorization process (RFC 6742 Section 4.2).
        /// </summary>
        public readonly bool EnableImplicitGrant;
        /// <summary>
        /// An optional boolean value to enable the Resource Owner Password Credentials Grant flow (RFC 6742 Section 4.3).
        /// </summary>
        public readonly bool EnablePasswordGrant;
        /// <summary>
        /// An optional boolean value that allows using the same OAuth credentials generated by the plugin with any other service whose OAuth 2.0 plugin configuration also has `config.global_credentials=true`.
        /// </summary>
        public readonly bool GlobalCredentials;
        /// <summary>
        /// An optional boolean value telling the plugin to show or hide the credential from the upstream service.
        /// </summary>
        public readonly bool HideCredentials;
        /// <summary>
        /// An optional boolean value telling the plugin to require at least one `scope` to be authorized by the end user.
        /// </summary>
        public readonly bool MandatoryScope;
        public readonly bool PersistentRefreshToken;
        /// <summary>
        /// Specifies a mode of how the Proof Key for Code Exchange (PKCE) should be handled by the plugin. must be one of ["none", "lax", "strict"]
        /// </summary>
        public readonly string Pkce;
        /// <summary>
        /// The unique key the plugin has generated when it has been added to the Service.
        /// </summary>
        public readonly string ProvisionKey;
        /// <summary>
        /// Time-to-live value for data
        /// </summary>
        public readonly double RefreshTokenTtl;
        /// <summary>
        /// An optional boolean value that indicates whether an OAuth refresh token is reused when refreshing an access token.
        /// </summary>
        public readonly bool ReuseRefreshToken;
        /// <summary>
        /// Describes an array of scope names that will be available to the end user. If `mandatory_scope` is set to `true`, then `scopes` are required.
        /// </summary>
        public readonly ImmutableArray<string> Scopes;
        /// <summary>
        /// An optional integer value telling the plugin how many seconds a token should last, after which the client will need to refresh the token. Set to `0` to disable the expiration.
        /// </summary>
        public readonly double TokenExpiration;

        [OutputConstructor]
        private GetGatewayPluginOauth2ConfigResult(
            bool acceptHttpIfAlreadyTerminated,

            string anonymous,

            string authHeaderName,

            bool enableAuthorizationCode,

            bool enableClientCredentials,

            bool enableImplicitGrant,

            bool enablePasswordGrant,

            bool globalCredentials,

            bool hideCredentials,

            bool mandatoryScope,

            bool persistentRefreshToken,

            string pkce,

            string provisionKey,

            double refreshTokenTtl,

            bool reuseRefreshToken,

            ImmutableArray<string> scopes,

            double tokenExpiration)
        {
            AcceptHttpIfAlreadyTerminated = acceptHttpIfAlreadyTerminated;
            Anonymous = anonymous;
            AuthHeaderName = authHeaderName;
            EnableAuthorizationCode = enableAuthorizationCode;
            EnableClientCredentials = enableClientCredentials;
            EnableImplicitGrant = enableImplicitGrant;
            EnablePasswordGrant = enablePasswordGrant;
            GlobalCredentials = globalCredentials;
            HideCredentials = hideCredentials;
            MandatoryScope = mandatoryScope;
            PersistentRefreshToken = persistentRefreshToken;
            Pkce = pkce;
            ProvisionKey = provisionKey;
            RefreshTokenTtl = refreshTokenTtl;
            ReuseRefreshToken = reuseRefreshToken;
            Scopes = scopes;
            TokenExpiration = tokenExpiration;
        }
    }
}
