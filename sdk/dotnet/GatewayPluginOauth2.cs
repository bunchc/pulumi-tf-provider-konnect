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
    /// GatewayPluginOauth2 Resource
    /// </summary>
    [KonnectResourceType("konnect:index/gatewayPluginOauth2:GatewayPluginOauth2")]
    public partial class GatewayPluginOauth2 : global::Pulumi.CustomResource
    {
        [Output("config")]
        public Output<Outputs.GatewayPluginOauth2Config> Config { get; private set; } = null!;

        /// <summary>
        /// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
        /// </summary>
        [Output("consumer")]
        public Output<Outputs.GatewayPluginOauth2Consumer> Consumer { get; private set; } = null!;

        [Output("consumerGroup")]
        public Output<Outputs.GatewayPluginOauth2ConsumerGroup> ConsumerGroup { get; private set; } = null!;

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
        /// Whether the plugin is applied.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        [Output("instanceName")]
        public Output<string> InstanceName { get; private set; } = null!;

        /// <summary>
        /// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
        /// </summary>
        [Output("protocols")]
        public Output<ImmutableArray<string>> Protocols { get; private set; } = null!;

        /// <summary>
        /// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
        /// </summary>
        [Output("route")]
        public Output<Outputs.GatewayPluginOauth2Route> Route { get; private set; } = null!;

        /// <summary>
        /// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
        /// </summary>
        [Output("service")]
        public Output<Outputs.GatewayPluginOauth2Service> Service { get; private set; } = null!;

        /// <summary>
        /// An optional set of strings associated with the Plugin for grouping and filtering.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Unix epoch when the resource was last updated.
        /// </summary>
        [Output("updatedAt")]
        public Output<int> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a GatewayPluginOauth2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GatewayPluginOauth2(string name, GatewayPluginOauth2Args args, CustomResourceOptions? options = null)
            : base("konnect:index/gatewayPluginOauth2:GatewayPluginOauth2", name, args ?? new GatewayPluginOauth2Args(), MakeResourceOptions(options, ""))
        {
        }

        private GatewayPluginOauth2(string name, Input<string> id, GatewayPluginOauth2State? state = null, CustomResourceOptions? options = null)
            : base("konnect:index/gatewayPluginOauth2:GatewayPluginOauth2", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GatewayPluginOauth2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GatewayPluginOauth2 Get(string name, Input<string> id, GatewayPluginOauth2State? state = null, CustomResourceOptions? options = null)
        {
            return new GatewayPluginOauth2(name, id, state, options);
        }
    }

    public sealed class GatewayPluginOauth2Args : global::Pulumi.ResourceArgs
    {
        [Input("config")]
        public Input<Inputs.GatewayPluginOauth2ConfigArgs>? Config { get; set; }

        /// <summary>
        /// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
        /// </summary>
        [Input("consumer")]
        public Input<Inputs.GatewayPluginOauth2ConsumerArgs>? Consumer { get; set; }

        [Input("consumerGroup")]
        public Input<Inputs.GatewayPluginOauth2ConsumerGroupArgs>? ConsumerGroup { get; set; }

        /// <summary>
        /// The UUID of your control plane. This variable is available in the Konnect manager.
        /// </summary>
        [Input("controlPlaneId", required: true)]
        public Input<string> ControlPlaneId { get; set; } = null!;

        /// <summary>
        /// Whether the plugin is applied.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        [Input("protocols")]
        private InputList<string>? _protocols;

        /// <summary>
        /// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
        /// </summary>
        public InputList<string> Protocols
        {
            get => _protocols ?? (_protocols = new InputList<string>());
            set => _protocols = value;
        }

        /// <summary>
        /// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
        /// </summary>
        [Input("route")]
        public Input<Inputs.GatewayPluginOauth2RouteArgs>? Route { get; set; }

        /// <summary>
        /// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
        /// </summary>
        [Input("service")]
        public Input<Inputs.GatewayPluginOauth2ServiceArgs>? Service { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// An optional set of strings associated with the Plugin for grouping and filtering.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public GatewayPluginOauth2Args()
        {
        }
        public static new GatewayPluginOauth2Args Empty => new GatewayPluginOauth2Args();
    }

    public sealed class GatewayPluginOauth2State : global::Pulumi.ResourceArgs
    {
        [Input("config")]
        public Input<Inputs.GatewayPluginOauth2ConfigGetArgs>? Config { get; set; }

        /// <summary>
        /// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
        /// </summary>
        [Input("consumer")]
        public Input<Inputs.GatewayPluginOauth2ConsumerGetArgs>? Consumer { get; set; }

        [Input("consumerGroup")]
        public Input<Inputs.GatewayPluginOauth2ConsumerGroupGetArgs>? ConsumerGroup { get; set; }

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
        /// Whether the plugin is applied.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        [Input("protocols")]
        private InputList<string>? _protocols;

        /// <summary>
        /// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
        /// </summary>
        public InputList<string> Protocols
        {
            get => _protocols ?? (_protocols = new InputList<string>());
            set => _protocols = value;
        }

        /// <summary>
        /// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
        /// </summary>
        [Input("route")]
        public Input<Inputs.GatewayPluginOauth2RouteGetArgs>? Route { get; set; }

        /// <summary>
        /// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
        /// </summary>
        [Input("service")]
        public Input<Inputs.GatewayPluginOauth2ServiceGetArgs>? Service { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// An optional set of strings associated with the Plugin for grouping and filtering.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Unix epoch when the resource was last updated.
        /// </summary>
        [Input("updatedAt")]
        public Input<int>? UpdatedAt { get; set; }

        public GatewayPluginOauth2State()
        {
        }
        public static new GatewayPluginOauth2State Empty => new GatewayPluginOauth2State();
    }
}