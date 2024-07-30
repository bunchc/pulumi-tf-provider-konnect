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
    public sealed class GetGatewayPluginAwsLambdaConfigResult
    {
        /// <summary>
        /// The target AWS IAM role ARN used to invoke the Lambda function.
        /// </summary>
        public readonly string AwsAssumeRoleArn;
        /// <summary>
        /// Identifier to select the IMDS protocol version to use: `v1` or `v2`. must be one of ["v1", "v2"]
        /// </summary>
        public readonly string AwsImdsProtocolVersion;
        /// <summary>
        /// The AWS key credential to be used when invoking the function.
        /// </summary>
        public readonly string AwsKey;
        /// <summary>
        /// A string representing a host name, such as example.com.
        /// </summary>
        public readonly string AwsRegion;
        /// <summary>
        /// The identifier of the assumed role session.
        /// </summary>
        public readonly string AwsRoleSessionName;
        /// <summary>
        /// The AWS secret credential to be used when invoking the function.
        /// </summary>
        public readonly string AwsSecret;
        /// <summary>
        /// An optional value that defines whether the plugin should wrap requests into the Amazon API gateway.
        /// </summary>
        public readonly bool AwsgatewayCompatible;
        /// <summary>
        /// An optional value that Base64-encodes the request body.
        /// </summary>
        public readonly bool Base64EncodeBody;
        public readonly bool DisableHttps;
        /// <summary>
        /// An optional value that defines whether the request body is sent in the request_body field of the JSON-encoded request. If the body arguments can be parsed, they are sent in the separate request_body_args field of the request.
        /// </summary>
        public readonly bool ForwardRequestBody;
        /// <summary>
        /// An optional value that defines whether the original HTTP request headers are sent as a map in the request_headers field of the JSON-encoded request.
        /// </summary>
        public readonly bool ForwardRequestHeaders;
        /// <summary>
        /// An optional value that defines whether the original HTTP request method verb is sent in the request_method field of the JSON-encoded request.
        /// </summary>
        public readonly bool ForwardRequestMethod;
        /// <summary>
        /// An optional value that defines whether the original HTTP request URI is sent in the request_uri field of the JSON-encoded request.
        /// </summary>
        public readonly bool ForwardRequestUri;
        /// <summary>
        /// The AWS Lambda function to invoke. Both function name and function ARN (including partial) are supported.
        /// </summary>
        public readonly string FunctionName;
        /// <summary>
        /// A string representing a host name, such as example.com.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// The InvocationType to use when invoking the function. Available types are RequestResponse, Event, DryRun. must be one of ["RequestResponse", "Event", "DryRun"]
        /// </summary>
        public readonly string InvocationType;
        /// <summary>
        /// An optional value that defines whether the response format to receive from the Lambda to this format.
        /// </summary>
        public readonly bool IsProxyIntegration;
        /// <summary>
        /// An optional value in milliseconds that defines how long an idle connection lives before being closed.
        /// </summary>
        public readonly double Keepalive;
        /// <summary>
        /// The LogType to use when invoking the function. By default, None and Tail are supported. must be one of ["Tail", "None"]
        /// </summary>
        public readonly string LogType;
        /// <summary>
        /// An integer representing a port number between 0 and 65535, inclusive.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// A string representing a URL, such as https://example.com/path/to/resource?q=search.
        /// </summary>
        public readonly string ProxyUrl;
        /// <summary>
        /// The qualifier to use when invoking the function.
        /// </summary>
        public readonly string Qualifier;
        /// <summary>
        /// An optional value that defines whether Kong should send large bodies that are buffered to disk
        /// </summary>
        public readonly bool SkipLargeBodies;
        /// <summary>
        /// An optional timeout in milliseconds when invoking the function.
        /// </summary>
        public readonly double Timeout;
        /// <summary>
        /// The response status code to use (instead of the default 200, 202, or 204) in the case of an Unhandled Function Error.
        /// </summary>
        public readonly int UnhandledStatus;

        [OutputConstructor]
        private GetGatewayPluginAwsLambdaConfigResult(
            string awsAssumeRoleArn,

            string awsImdsProtocolVersion,

            string awsKey,

            string awsRegion,

            string awsRoleSessionName,

            string awsSecret,

            bool awsgatewayCompatible,

            bool base64EncodeBody,

            bool disableHttps,

            bool forwardRequestBody,

            bool forwardRequestHeaders,

            bool forwardRequestMethod,

            bool forwardRequestUri,

            string functionName,

            string host,

            string invocationType,

            bool isProxyIntegration,

            double keepalive,

            string logType,

            int port,

            string proxyUrl,

            string qualifier,

            bool skipLargeBodies,

            double timeout,

            int unhandledStatus)
        {
            AwsAssumeRoleArn = awsAssumeRoleArn;
            AwsImdsProtocolVersion = awsImdsProtocolVersion;
            AwsKey = awsKey;
            AwsRegion = awsRegion;
            AwsRoleSessionName = awsRoleSessionName;
            AwsSecret = awsSecret;
            AwsgatewayCompatible = awsgatewayCompatible;
            Base64EncodeBody = base64EncodeBody;
            DisableHttps = disableHttps;
            ForwardRequestBody = forwardRequestBody;
            ForwardRequestHeaders = forwardRequestHeaders;
            ForwardRequestMethod = forwardRequestMethod;
            ForwardRequestUri = forwardRequestUri;
            FunctionName = functionName;
            Host = host;
            InvocationType = invocationType;
            IsProxyIntegration = isProxyIntegration;
            Keepalive = keepalive;
            LogType = logType;
            Port = port;
            ProxyUrl = proxyUrl;
            Qualifier = qualifier;
            SkipLargeBodies = skipLargeBodies;
            Timeout = timeout;
            UnhandledStatus = unhandledStatus;
        }
    }
}