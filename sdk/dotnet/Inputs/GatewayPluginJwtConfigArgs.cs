// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayPluginJwtConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional string (consumer UUID or username) value to use as an “anonymous” consumer if authentication fails.
        /// </summary>
        [Input("anonymous")]
        public Input<string>? Anonymous { get; set; }

        [Input("claimsToVerifies")]
        private InputList<string>? _claimsToVerifies;

        /// <summary>
        /// A list of registered claims (according to RFC 7519) that Kong can verify as well. Accepted values: one of exp or nbf.
        /// </summary>
        public InputList<string> ClaimsToVerifies
        {
            get => _claimsToVerifies ?? (_claimsToVerifies = new InputList<string>());
            set => _claimsToVerifies = value;
        }

        [Input("cookieNames")]
        private InputList<string>? _cookieNames;

        /// <summary>
        /// A list of cookie names that Kong will inspect to retrieve JWTs.
        /// </summary>
        public InputList<string> CookieNames
        {
            get => _cookieNames ?? (_cookieNames = new InputList<string>());
            set => _cookieNames = value;
        }

        [Input("headerNames")]
        private InputList<string>? _headerNames;

        /// <summary>
        /// A list of HTTP header names that Kong will inspect to retrieve JWTs.
        /// </summary>
        public InputList<string> HeaderNames
        {
            get => _headerNames ?? (_headerNames = new InputList<string>());
            set => _headerNames = value;
        }

        /// <summary>
        /// The name of the claim in which the key identifying the secret must be passed. The plugin will attempt to read this claim from the JWT payload and the header, in that order.
        /// </summary>
        [Input("keyClaimName")]
        public Input<string>? KeyClaimName { get; set; }

        /// <summary>
        /// A value between 0 and 31536000 (365 days) limiting the lifetime of the JWT to maximum_expiration seconds in the future.
        /// </summary>
        [Input("maximumExpiration")]
        public Input<double>? MaximumExpiration { get; set; }

        /// <summary>
        /// A boolean value that indicates whether the plugin should run (and try to authenticate) on OPTIONS preflight requests. If set to false, then OPTIONS requests will always be allowed.
        /// </summary>
        [Input("runOnPreflight")]
        public Input<bool>? RunOnPreflight { get; set; }

        /// <summary>
        /// If true, the plugin assumes the credential’s secret to be base64 encoded. You will need to create a base64-encoded secret for your Consumer, and sign your JWT with the original secret.
        /// </summary>
        [Input("secretIsBase64")]
        public Input<bool>? SecretIsBase64 { get; set; }

        [Input("uriParamNames")]
        private InputList<string>? _uriParamNames;

        /// <summary>
        /// A list of querystring parameters that Kong will inspect to retrieve JWTs.
        /// </summary>
        public InputList<string> UriParamNames
        {
            get => _uriParamNames ?? (_uriParamNames = new InputList<string>());
            set => _uriParamNames = value;
        }

        public GatewayPluginJwtConfigArgs()
        {
        }
        public static new GatewayPluginJwtConfigArgs Empty => new GatewayPluginJwtConfigArgs();
    }
}
