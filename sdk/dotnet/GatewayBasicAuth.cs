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
    /// GatewayBasicAuth Resource
    /// </summary>
    [KonnectResourceType("konnect:index/gatewayBasicAuth:GatewayBasicAuth")]
    public partial class GatewayBasicAuth : global::Pulumi.CustomResource
    {
        [Output("consumer")]
        public Output<Outputs.GatewayBasicAuthConsumer> Consumer { get; private set; } = null!;

        /// <summary>
        /// Consumer ID for nested entities. Requires replacement if changed.
        /// </summary>
        [Output("consumerId")]
        public Output<string> ConsumerId { get; private set; } = null!;

        /// <summary>
        /// The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.
        /// </summary>
        [Output("controlPlaneId")]
        public Output<string> ControlPlaneId { get; private set; } = null!;

        /// <summary>
        /// Unix epoch when the resource was created.
        /// </summary>
        [Output("createdAt")]
        public Output<int> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Requires replacement if changed.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Requires replacement if changed.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Requires replacement if changed.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a GatewayBasicAuth resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GatewayBasicAuth(string name, GatewayBasicAuthArgs args, CustomResourceOptions? options = null)
            : base("konnect:index/gatewayBasicAuth:GatewayBasicAuth", name, args ?? new GatewayBasicAuthArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GatewayBasicAuth(string name, Input<string> id, GatewayBasicAuthState? state = null, CustomResourceOptions? options = null)
            : base("konnect:index/gatewayBasicAuth:GatewayBasicAuth", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing GatewayBasicAuth resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GatewayBasicAuth Get(string name, Input<string> id, GatewayBasicAuthState? state = null, CustomResourceOptions? options = null)
        {
            return new GatewayBasicAuth(name, id, state, options);
        }
    }

    public sealed class GatewayBasicAuthArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Consumer ID for nested entities. Requires replacement if changed.
        /// </summary>
        [Input("consumerId", required: true)]
        public Input<string> ConsumerId { get; set; } = null!;

        /// <summary>
        /// The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.
        /// </summary>
        [Input("controlPlaneId", required: true)]
        public Input<string> ControlPlaneId { get; set; } = null!;

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Requires replacement if changed.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Requires replacement if changed.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Requires replacement if changed.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public GatewayBasicAuthArgs()
        {
        }
        public static new GatewayBasicAuthArgs Empty => new GatewayBasicAuthArgs();
    }

    public sealed class GatewayBasicAuthState : global::Pulumi.ResourceArgs
    {
        [Input("consumer")]
        public Input<Inputs.GatewayBasicAuthConsumerGetArgs>? Consumer { get; set; }

        /// <summary>
        /// Consumer ID for nested entities. Requires replacement if changed.
        /// </summary>
        [Input("consumerId")]
        public Input<string>? ConsumerId { get; set; }

        /// <summary>
        /// The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.
        /// </summary>
        [Input("controlPlaneId")]
        public Input<string>? ControlPlaneId { get; set; }

        /// <summary>
        /// Unix epoch when the resource was created.
        /// </summary>
        [Input("createdAt")]
        public Input<int>? CreatedAt { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Requires replacement if changed.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Requires replacement if changed.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Requires replacement if changed.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public GatewayBasicAuthState()
        {
        }
        public static new GatewayBasicAuthState Empty => new GatewayBasicAuthState();
    }
}