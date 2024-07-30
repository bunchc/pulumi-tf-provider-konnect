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
    /// CloudGatewayNetwork Resource
    /// </summary>
    [KonnectResourceType("konnect:index/cloudGatewayNetwork:CloudGatewayNetwork")]
    public partial class CloudGatewayNetwork : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of availability zones that the network is attached to. Requires replacement if changed.
        /// </summary>
        [Output("availabilityZones")]
        public Output<ImmutableArray<string>> AvailabilityZones { get; private set; } = null!;

        /// <summary>
        /// CIDR block configuration for the network. Requires replacement if changed.
        /// </summary>
        [Output("cidrBlock")]
        public Output<string> CidrBlock { get; private set; } = null!;

        /// <summary>
        /// Requires replacement if changed.
        /// </summary>
        [Output("cloudGatewayProviderAccountId")]
        public Output<string> CloudGatewayProviderAccountId { get; private set; } = null!;

        /// <summary>
        /// The number of configurations that reference this network.
        /// </summary>
        [Output("configurationReferenceCount")]
        public Output<int> ConfigurationReferenceCount { get; private set; } = null!;

        /// <summary>
        /// An RFC-3339 timestamp representation of network creation date.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Whether DDOS protection is enabled for the network. Requires replacement if changed.
        /// </summary>
        [Output("ddosProtection")]
        public Output<bool> DdosProtection { get; private set; } = null!;

        /// <summary>
        /// Whether the network is a default network or not. Default networks are Networks that are created
        /// automatically by Konnect when an organization is linked to a provider account.
        /// </summary>
        [Output("default")]
        public Output<bool> Default { get; private set; } = null!;

        /// <summary>
        /// Monotonically-increasing version count of the network, to indicate the order of updates to the network.
        /// </summary>
        [Output("entityVersion")]
        public Output<int> EntityVersion { get; private set; } = null!;

        /// <summary>
        /// Firewall configuration for a network.
        /// </summary>
        [Output("firewall")]
        public Output<Outputs.CloudGatewayNetworkFirewall> Firewall { get; private set; } = null!;

        /// <summary>
        /// Human-readable name of the network.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Metadata describing attributes returned by cloud-provider for the network.
        /// </summary>
        [Output("providerMetadata")]
        public Output<Outputs.CloudGatewayNetworkProviderMetadata> ProviderMetadata { get; private set; } = null!;

        /// <summary>
        /// Region ID for cloud provider region. Requires replacement if changed.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// State of the network. must be one of ["created", "initializing", "offline", "ready", "terminating", "terminated"]
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The number of transit gateways attached to this network.
        /// </summary>
        [Output("transitGatewayCount")]
        public Output<int> TransitGatewayCount { get; private set; } = null!;

        /// <summary>
        /// An RFC-3339 timestamp representation of network update date.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a CloudGatewayNetwork resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CloudGatewayNetwork(string name, CloudGatewayNetworkArgs args, CustomResourceOptions? options = null)
            : base("konnect:index/cloudGatewayNetwork:CloudGatewayNetwork", name, args ?? new CloudGatewayNetworkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CloudGatewayNetwork(string name, Input<string> id, CloudGatewayNetworkState? state = null, CustomResourceOptions? options = null)
            : base("konnect:index/cloudGatewayNetwork:CloudGatewayNetwork", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CloudGatewayNetwork resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CloudGatewayNetwork Get(string name, Input<string> id, CloudGatewayNetworkState? state = null, CustomResourceOptions? options = null)
        {
            return new CloudGatewayNetwork(name, id, state, options);
        }
    }

    public sealed class CloudGatewayNetworkArgs : global::Pulumi.ResourceArgs
    {
        [Input("availabilityZones", required: true)]
        private InputList<string>? _availabilityZones;

        /// <summary>
        /// List of availability zones that the network is attached to. Requires replacement if changed.
        /// </summary>
        public InputList<string> AvailabilityZones
        {
            get => _availabilityZones ?? (_availabilityZones = new InputList<string>());
            set => _availabilityZones = value;
        }

        /// <summary>
        /// CIDR block configuration for the network. Requires replacement if changed.
        /// </summary>
        [Input("cidrBlock", required: true)]
        public Input<string> CidrBlock { get; set; } = null!;

        /// <summary>
        /// Requires replacement if changed.
        /// </summary>
        [Input("cloudGatewayProviderAccountId", required: true)]
        public Input<string> CloudGatewayProviderAccountId { get; set; } = null!;

        /// <summary>
        /// Whether DDOS protection is enabled for the network. Requires replacement if changed.
        /// </summary>
        [Input("ddosProtection")]
        public Input<bool>? DdosProtection { get; set; }

        /// <summary>
        /// Firewall configuration for a network.
        /// </summary>
        [Input("firewall")]
        public Input<Inputs.CloudGatewayNetworkFirewallArgs>? Firewall { get; set; }

        /// <summary>
        /// Human-readable name of the network.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Region ID for cloud provider region. Requires replacement if changed.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public CloudGatewayNetworkArgs()
        {
        }
        public static new CloudGatewayNetworkArgs Empty => new CloudGatewayNetworkArgs();
    }

    public sealed class CloudGatewayNetworkState : global::Pulumi.ResourceArgs
    {
        [Input("availabilityZones")]
        private InputList<string>? _availabilityZones;

        /// <summary>
        /// List of availability zones that the network is attached to. Requires replacement if changed.
        /// </summary>
        public InputList<string> AvailabilityZones
        {
            get => _availabilityZones ?? (_availabilityZones = new InputList<string>());
            set => _availabilityZones = value;
        }

        /// <summary>
        /// CIDR block configuration for the network. Requires replacement if changed.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// Requires replacement if changed.
        /// </summary>
        [Input("cloudGatewayProviderAccountId")]
        public Input<string>? CloudGatewayProviderAccountId { get; set; }

        /// <summary>
        /// The number of configurations that reference this network.
        /// </summary>
        [Input("configurationReferenceCount")]
        public Input<int>? ConfigurationReferenceCount { get; set; }

        /// <summary>
        /// An RFC-3339 timestamp representation of network creation date.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Whether DDOS protection is enabled for the network. Requires replacement if changed.
        /// </summary>
        [Input("ddosProtection")]
        public Input<bool>? DdosProtection { get; set; }

        /// <summary>
        /// Whether the network is a default network or not. Default networks are Networks that are created
        /// automatically by Konnect when an organization is linked to a provider account.
        /// </summary>
        [Input("default")]
        public Input<bool>? Default { get; set; }

        /// <summary>
        /// Monotonically-increasing version count of the network, to indicate the order of updates to the network.
        /// </summary>
        [Input("entityVersion")]
        public Input<int>? EntityVersion { get; set; }

        /// <summary>
        /// Firewall configuration for a network.
        /// </summary>
        [Input("firewall")]
        public Input<Inputs.CloudGatewayNetworkFirewallGetArgs>? Firewall { get; set; }

        /// <summary>
        /// Human-readable name of the network.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Metadata describing attributes returned by cloud-provider for the network.
        /// </summary>
        [Input("providerMetadata")]
        public Input<Inputs.CloudGatewayNetworkProviderMetadataGetArgs>? ProviderMetadata { get; set; }

        /// <summary>
        /// Region ID for cloud provider region. Requires replacement if changed.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// State of the network. must be one of ["created", "initializing", "offline", "ready", "terminating", "terminated"]
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The number of transit gateways attached to this network.
        /// </summary>
        [Input("transitGatewayCount")]
        public Input<int>? TransitGatewayCount { get; set; }

        /// <summary>
        /// An RFC-3339 timestamp representation of network update date.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public CloudGatewayNetworkState()
        {
        }
        public static new CloudGatewayNetworkState Empty => new CloudGatewayNetworkState();
    }
}