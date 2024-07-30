// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect
{
    /// <summary>
    /// GatewayService Resource
    /// </summary>
    [KonnectResourceType("konnect:index/gatewayService:GatewayService")]
    public partial class GatewayService : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Array of `CA Certificate` object UUIDs that are used to build the trust store while verifying upstream server's TLS certificate. If set to `null` when Nginx default is respected. If default CA list in Nginx are not specified and TLS verification is enabled, then handshake with upstream server will always fail (because no CA are trusted).
        /// </summary>
        [Output("caCertificates")]
        public Output<ImmutableArray<string>> CaCertificates { get; private set; } = null!;

        /// <summary>
        /// Certificate to be used as client certificate while TLS handshaking to the upstream server.
        /// </summary>
        [Output("clientCertificate")]
        public Output<Outputs.GatewayServiceClientCertificate> ClientCertificate { get; private set; } = null!;

        /// <summary>
        /// The timeout in milliseconds for establishing a connection to the upstream server.
        /// </summary>
        [Output("connectTimeout")]
        public Output<int> ConnectTimeout { get; private set; } = null!;

        /// <summary>
        /// The UUID of your control plane. This variable is available in the Konnect manager.
        /// </summary>
        [Output("controlPlaneId")]
        public Output<string> ControlPlaneId { get; private set; } = null!;

        /// <summary>
        /// Unix epoch when the resource was created.
        /// </summary>
        [Output("createdAt")]
        public Output<int> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Whether the Service is active. If set to `false`, the proxy behavior will be as if any routes attached to it do not exist (404). Default: `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// The host of the upstream server. Note that the host value is case sensitive.
        /// </summary>
        [Output("host")]
        public Output<string> Host { get; private set; } = null!;

        /// <summary>
        /// The Service name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The path to be used in requests to the upstream server.
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;

        /// <summary>
        /// The upstream server port.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// The protocol used to communicate with the upstream. must be one of ["grpc", "grpcs", "http", "https", "tcp", "tls", "tls_passthrough", "udp", "ws", "wss"]
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// The timeout in milliseconds between two successive read operations for transmitting a request to the upstream server.
        /// </summary>
        [Output("readTimeout")]
        public Output<int> ReadTimeout { get; private set; } = null!;

        /// <summary>
        /// The number of retries to execute upon failure to proxy.
        /// </summary>
        [Output("retries")]
        public Output<int> Retries { get; private set; } = null!;

        /// <summary>
        /// An optional set of strings associated with the Service for grouping and filtering.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Whether to enable verification of upstream server TLS certificate. If set to `null`, then the Nginx default is respected.
        /// </summary>
        [Output("tlsVerify")]
        public Output<bool> TlsVerify { get; private set; } = null!;

        /// <summary>
        /// Maximum depth of chain while verifying Upstream server's TLS certificate. If set to `null`, then the Nginx default is respected.
        /// </summary>
        [Output("tlsVerifyDepth")]
        public Output<int> TlsVerifyDepth { get; private set; } = null!;

        /// <summary>
        /// Unix epoch when the resource was last updated.
        /// </summary>
        [Output("updatedAt")]
        public Output<int> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The timeout in milliseconds between two successive write operations for transmitting a request to the upstream server.
        /// </summary>
        [Output("writeTimeout")]
        public Output<int> WriteTimeout { get; private set; } = null!;


        /// <summary>
        /// Create a GatewayService resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GatewayService(string name, GatewayServiceArgs args, CustomResourceOptions? options = null)
            : base("konnect:index/gatewayService:GatewayService", name, args ?? new GatewayServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GatewayService(string name, Input<string> id, GatewayServiceState? state = null, CustomResourceOptions? options = null)
            : base("konnect:index/gatewayService:GatewayService", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing GatewayService resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GatewayService Get(string name, Input<string> id, GatewayServiceState? state = null, CustomResourceOptions? options = null)
        {
            return new GatewayService(name, id, state, options);
        }
    }

    public sealed class GatewayServiceArgs : global::Pulumi.ResourceArgs
    {
        [Input("caCertificates")]
        private InputList<string>? _caCertificates;

        /// <summary>
        /// Array of `CA Certificate` object UUIDs that are used to build the trust store while verifying upstream server's TLS certificate. If set to `null` when Nginx default is respected. If default CA list in Nginx are not specified and TLS verification is enabled, then handshake with upstream server will always fail (because no CA are trusted).
        /// </summary>
        public InputList<string> CaCertificates
        {
            get => _caCertificates ?? (_caCertificates = new InputList<string>());
            set => _caCertificates = value;
        }

        /// <summary>
        /// Certificate to be used as client certificate while TLS handshaking to the upstream server.
        /// </summary>
        [Input("clientCertificate")]
        public Input<Inputs.GatewayServiceClientCertificateArgs>? ClientCertificate { get; set; }

        /// <summary>
        /// The timeout in milliseconds for establishing a connection to the upstream server.
        /// </summary>
        [Input("connectTimeout")]
        public Input<int>? ConnectTimeout { get; set; }

        /// <summary>
        /// The UUID of your control plane. This variable is available in the Konnect manager.
        /// </summary>
        [Input("controlPlaneId", required: true)]
        public Input<string> ControlPlaneId { get; set; } = null!;

        /// <summary>
        /// Whether the Service is active. If set to `false`, the proxy behavior will be as if any routes attached to it do not exist (404). Default: `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The host of the upstream server. Note that the host value is case sensitive.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// The Service name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The path to be used in requests to the upstream server.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The upstream server port.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The protocol used to communicate with the upstream. must be one of ["grpc", "grpcs", "http", "https", "tcp", "tls", "tls_passthrough", "udp", "ws", "wss"]
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The timeout in milliseconds between two successive read operations for transmitting a request to the upstream server.
        /// </summary>
        [Input("readTimeout")]
        public Input<int>? ReadTimeout { get; set; }

        /// <summary>
        /// The number of retries to execute upon failure to proxy.
        /// </summary>
        [Input("retries")]
        public Input<int>? Retries { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// An optional set of strings associated with the Service for grouping and filtering.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Whether to enable verification of upstream server TLS certificate. If set to `null`, then the Nginx default is respected.
        /// </summary>
        [Input("tlsVerify")]
        public Input<bool>? TlsVerify { get; set; }

        /// <summary>
        /// Maximum depth of chain while verifying Upstream server's TLS certificate. If set to `null`, then the Nginx default is respected.
        /// </summary>
        [Input("tlsVerifyDepth")]
        public Input<int>? TlsVerifyDepth { get; set; }

        /// <summary>
        /// The timeout in milliseconds between two successive write operations for transmitting a request to the upstream server.
        /// </summary>
        [Input("writeTimeout")]
        public Input<int>? WriteTimeout { get; set; }

        public GatewayServiceArgs()
        {
        }
        public static new GatewayServiceArgs Empty => new GatewayServiceArgs();
    }

    public sealed class GatewayServiceState : global::Pulumi.ResourceArgs
    {
        [Input("caCertificates")]
        private InputList<string>? _caCertificates;

        /// <summary>
        /// Array of `CA Certificate` object UUIDs that are used to build the trust store while verifying upstream server's TLS certificate. If set to `null` when Nginx default is respected. If default CA list in Nginx are not specified and TLS verification is enabled, then handshake with upstream server will always fail (because no CA are trusted).
        /// </summary>
        public InputList<string> CaCertificates
        {
            get => _caCertificates ?? (_caCertificates = new InputList<string>());
            set => _caCertificates = value;
        }

        /// <summary>
        /// Certificate to be used as client certificate while TLS handshaking to the upstream server.
        /// </summary>
        [Input("clientCertificate")]
        public Input<Inputs.GatewayServiceClientCertificateGetArgs>? ClientCertificate { get; set; }

        /// <summary>
        /// The timeout in milliseconds for establishing a connection to the upstream server.
        /// </summary>
        [Input("connectTimeout")]
        public Input<int>? ConnectTimeout { get; set; }

        /// <summary>
        /// The UUID of your control plane. This variable is available in the Konnect manager.
        /// </summary>
        [Input("controlPlaneId")]
        public Input<string>? ControlPlaneId { get; set; }

        /// <summary>
        /// Unix epoch when the resource was created.
        /// </summary>
        [Input("createdAt")]
        public Input<int>? CreatedAt { get; set; }

        /// <summary>
        /// Whether the Service is active. If set to `false`, the proxy behavior will be as if any routes attached to it do not exist (404). Default: `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The host of the upstream server. Note that the host value is case sensitive.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// The Service name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The path to be used in requests to the upstream server.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The upstream server port.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The protocol used to communicate with the upstream. must be one of ["grpc", "grpcs", "http", "https", "tcp", "tls", "tls_passthrough", "udp", "ws", "wss"]
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The timeout in milliseconds between two successive read operations for transmitting a request to the upstream server.
        /// </summary>
        [Input("readTimeout")]
        public Input<int>? ReadTimeout { get; set; }

        /// <summary>
        /// The number of retries to execute upon failure to proxy.
        /// </summary>
        [Input("retries")]
        public Input<int>? Retries { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// An optional set of strings associated with the Service for grouping and filtering.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Whether to enable verification of upstream server TLS certificate. If set to `null`, then the Nginx default is respected.
        /// </summary>
        [Input("tlsVerify")]
        public Input<bool>? TlsVerify { get; set; }

        /// <summary>
        /// Maximum depth of chain while verifying Upstream server's TLS certificate. If set to `null`, then the Nginx default is respected.
        /// </summary>
        [Input("tlsVerifyDepth")]
        public Input<int>? TlsVerifyDepth { get; set; }

        /// <summary>
        /// Unix epoch when the resource was last updated.
        /// </summary>
        [Input("updatedAt")]
        public Input<int>? UpdatedAt { get; set; }

        /// <summary>
        /// The timeout in milliseconds between two successive write operations for transmitting a request to the upstream server.
        /// </summary>
        [Input("writeTimeout")]
        public Input<int>? WriteTimeout { get; set; }

        public GatewayServiceState()
        {
        }
        public static new GatewayServiceState Empty => new GatewayServiceState();
    }
}