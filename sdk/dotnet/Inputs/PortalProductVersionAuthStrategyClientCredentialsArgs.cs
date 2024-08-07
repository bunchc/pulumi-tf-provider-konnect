// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class PortalProductVersionAuthStrategyClientCredentialsArgs : global::Pulumi.ResourceArgs
    {
        [Input("authMethods")]
        private InputList<string>? _authMethods;
        public InputList<string> AuthMethods
        {
            get => _authMethods ?? (_authMethods = new InputList<string>());
            set => _authMethods = value;
        }

        /// <summary>
        /// must be one of ["client*credentials", "self*managed*client*credentials"]
        /// </summary>
        [Input("credentialType")]
        public Input<string>? CredentialType { get; set; }

        /// <summary>
        /// The Application Auth Strategy ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public PortalProductVersionAuthStrategyClientCredentialsArgs()
        {
        }
        public static new PortalProductVersionAuthStrategyClientCredentialsArgs Empty => new PortalProductVersionAuthStrategyClientCredentialsArgs();
    }
}
