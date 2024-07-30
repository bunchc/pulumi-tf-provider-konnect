// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayPluginIpRestrictionConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("allows")]
        private InputList<string>? _allows;

        /// <summary>
        /// List of IPs or CIDR ranges to allow. One of `config.allow` or `config.deny` must be specified.
        /// </summary>
        public InputList<string> Allows
        {
            get => _allows ?? (_allows = new InputList<string>());
            set => _allows = value;
        }

        [Input("denies")]
        private InputList<string>? _denies;

        /// <summary>
        /// List of IPs or CIDR ranges to deny. One of `config.allow` or `config.deny` must be specified.
        /// </summary>
        public InputList<string> Denies
        {
            get => _denies ?? (_denies = new InputList<string>());
            set => _denies = value;
        }

        /// <summary>
        /// The message to send as a response body to rejected requests.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        /// <summary>
        /// The HTTP status of the requests that will be rejected by the plugin.
        /// </summary>
        [Input("status")]
        public Input<double>? Status { get; set; }

        public GatewayPluginIpRestrictionConfigGetArgs()
        {
        }
        public static new GatewayPluginIpRestrictionConfigGetArgs Empty => new GatewayPluginIpRestrictionConfigGetArgs();
    }
}