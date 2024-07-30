// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayPluginAwsLambdaConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The target AWS IAM role ARN used to invoke the Lambda function.
        /// </summary>
        [Input("awsAssumeRoleArn")]
        public Input<string>? AwsAssumeRoleArn { get; set; }

        /// <summary>
        /// Identifier to select the IMDS protocol version to use: `v1` or `v2`. must be one of ["v1", "v2"]
        /// </summary>
        [Input("awsImdsProtocolVersion")]
        public Input<string>? AwsImdsProtocolVersion { get; set; }

        /// <summary>
        /// The AWS key credential to be used when invoking the function.
        /// </summary>
        [Input("awsKey")]
        public Input<string>? AwsKey { get; set; }

        /// <summary>
        /// A string representing a host name, such as example.com.
        /// </summary>
        [Input("awsRegion")]
        public Input<string>? AwsRegion { get; set; }

        /// <summary>
        /// The identifier of the assumed role session.
        /// </summary>
        [Input("awsRoleSessionName")]
        public Input<string>? AwsRoleSessionName { get; set; }

        /// <summary>
        /// The AWS secret credential to be used when invoking the function.
        /// </summary>
        [Input("awsSecret")]
        public Input<string>? AwsSecret { get; set; }

        /// <summary>
        /// An optional value that defines whether the plugin should wrap requests into the Amazon API gateway.
        /// </summary>
        [Input("awsgatewayCompatible")]
        public Input<bool>? AwsgatewayCompatible { get; set; }

        /// <summary>
        /// An optional value that Base64-encodes the request body.
        /// </summary>
        [Input("base64EncodeBody")]
        public Input<bool>? Base64EncodeBody { get; set; }

        [Input("disableHttps")]
        public Input<bool>? DisableHttps { get; set; }

        /// <summary>
        /// An optional value that defines whether the request body is sent in the request*body field of the JSON-encoded request. If the body arguments can be parsed, they are sent in the separate request*body_args field of the request.
        /// </summary>
        [Input("forwardRequestBody")]
        public Input<bool>? ForwardRequestBody { get; set; }

        /// <summary>
        /// An optional value that defines whether the original HTTP request headers are sent as a map in the request_headers field of the JSON-encoded request.
        /// </summary>
        [Input("forwardRequestHeaders")]
        public Input<bool>? ForwardRequestHeaders { get; set; }

        /// <summary>
        /// An optional value that defines whether the original HTTP request method verb is sent in the request_method field of the JSON-encoded request.
        /// </summary>
        [Input("forwardRequestMethod")]
        public Input<bool>? ForwardRequestMethod { get; set; }

        /// <summary>
        /// An optional value that defines whether the original HTTP request URI is sent in the request_uri field of the JSON-encoded request.
        /// </summary>
        [Input("forwardRequestUri")]
        public Input<bool>? ForwardRequestUri { get; set; }

        /// <summary>
        /// The AWS Lambda function to invoke. Both function name and function ARN (including partial) are supported.
        /// </summary>
        [Input("functionName")]
        public Input<string>? FunctionName { get; set; }

        /// <summary>
        /// A string representing a host name, such as example.com.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// The InvocationType to use when invoking the function. Available types are RequestResponse, Event, DryRun. must be one of ["RequestResponse", "Event", "DryRun"]
        /// </summary>
        [Input("invocationType")]
        public Input<string>? InvocationType { get; set; }

        /// <summary>
        /// An optional value that defines whether the response format to receive from the Lambda to this format.
        /// </summary>
        [Input("isProxyIntegration")]
        public Input<bool>? IsProxyIntegration { get; set; }

        /// <summary>
        /// An optional value in milliseconds that defines how long an idle connection lives before being closed.
        /// </summary>
        [Input("keepalive")]
        public Input<double>? Keepalive { get; set; }

        /// <summary>
        /// The LogType to use when invoking the function. By default, None and Tail are supported. must be one of ["Tail", "None"]
        /// </summary>
        [Input("logType")]
        public Input<string>? LogType { get; set; }

        /// <summary>
        /// An integer representing a port number between 0 and 65535, inclusive.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// A string representing a URL, such as https://example.com/path/to/resource?q=search.
        /// </summary>
        [Input("proxyUrl")]
        public Input<string>? ProxyUrl { get; set; }

        /// <summary>
        /// The qualifier to use when invoking the function.
        /// </summary>
        [Input("qualifier")]
        public Input<string>? Qualifier { get; set; }

        /// <summary>
        /// An optional value that defines whether Kong should send large bodies that are buffered to disk
        /// </summary>
        [Input("skipLargeBodies")]
        public Input<bool>? SkipLargeBodies { get; set; }

        /// <summary>
        /// An optional timeout in milliseconds when invoking the function.
        /// </summary>
        [Input("timeout")]
        public Input<double>? Timeout { get; set; }

        /// <summary>
        /// The response status code to use (instead of the default 200, 202, or 204) in the case of an Unhandled Function Error.
        /// </summary>
        [Input("unhandledStatus")]
        public Input<int>? UnhandledStatus { get; set; }

        public GatewayPluginAwsLambdaConfigArgs()
        {
        }
        public static new GatewayPluginAwsLambdaConfigArgs Empty => new GatewayPluginAwsLambdaConfigArgs();
    }
}
