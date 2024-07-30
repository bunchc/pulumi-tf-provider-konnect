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
    public sealed class GetGatewayPluginRequestTerminationConfigResult
    {
        /// <summary>
        /// The raw response body to send. This is mutually exclusive with the `config.message` field.
        /// </summary>
        public readonly string Body;
        /// <summary>
        /// Content type of the raw response configured with `config.body`.
        /// </summary>
        public readonly string ContentType;
        /// <summary>
        /// When set, the plugin will echo a copy of the request back to the client. The main usecase for this is debugging. It can be combined with `trigger` in order to debug requests on live systems without disturbing real traffic.
        /// </summary>
        public readonly bool Echo;
        /// <summary>
        /// The message to send, if using the default response generator.
        /// </summary>
        public readonly string Message;
        /// <summary>
        /// The response code to send. Must be an integer between 100 and 599.
        /// </summary>
        public readonly int StatusCode;
        /// <summary>
        /// A string representing an HTTP header name.
        /// </summary>
        public readonly string Trigger;

        [OutputConstructor]
        private GetGatewayPluginRequestTerminationConfigResult(
            string body,

            string contentType,

            bool echo,

            string message,

            int statusCode,

            string trigger)
        {
            Body = body;
            ContentType = contentType;
            Echo = echo;
            Message = message;
            StatusCode = statusCode;
            Trigger = trigger;
        }
    }
}
