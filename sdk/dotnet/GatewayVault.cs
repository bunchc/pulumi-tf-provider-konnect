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
    /// GatewayVault Resource
    /// </summary>
    [KonnectResourceType("konnect:index/gatewayVault:GatewayVault")]
    public partial class GatewayVault : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The configuration properties for the Vault which can be found on the vaults' documentation page. Parsed as JSON.
        /// </summary>
        [Output("config")]
        public Output<string> Config { get; private set; } = null!;

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
        /// The description of the Vault entity.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the Vault that's going to be added. Currently, the Vault implementation must be installed in every Kong instance.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The unique prefix (or identifier) for this Vault configuration. The prefix is used to load the right Vault configuration and implementation when referencing secrets with the other entities.
        /// </summary>
        [Output("prefix")]
        public Output<string> Prefix { get; private set; } = null!;

        /// <summary>
        /// An optional set of strings associated with the Vault for grouping and filtering.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Unix epoch when the resource was last updated.
        /// </summary>
        [Output("updatedAt")]
        public Output<int> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a GatewayVault resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GatewayVault(string name, GatewayVaultArgs args, CustomResourceOptions? options = null)
            : base("konnect:index/gatewayVault:GatewayVault", name, args ?? new GatewayVaultArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GatewayVault(string name, Input<string> id, GatewayVaultState? state = null, CustomResourceOptions? options = null)
            : base("konnect:index/gatewayVault:GatewayVault", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GatewayVault resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GatewayVault Get(string name, Input<string> id, GatewayVaultState? state = null, CustomResourceOptions? options = null)
        {
            return new GatewayVault(name, id, state, options);
        }
    }

    public sealed class GatewayVaultArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration properties for the Vault which can be found on the vaults' documentation page. Parsed as JSON.
        /// </summary>
        [Input("config")]
        public Input<string>? Config { get; set; }

        /// <summary>
        /// The UUID of your control plane. This variable is available in the Konnect manager.
        /// </summary>
        [Input("controlPlaneId", required: true)]
        public Input<string> ControlPlaneId { get; set; } = null!;

        /// <summary>
        /// The description of the Vault entity.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the Vault that's going to be added. Currently, the Vault implementation must be installed in every Kong instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The unique prefix (or identifier) for this Vault configuration. The prefix is used to load the right Vault configuration and implementation when referencing secrets with the other entities.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// An optional set of strings associated with the Vault for grouping and filtering.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public GatewayVaultArgs()
        {
        }
        public static new GatewayVaultArgs Empty => new GatewayVaultArgs();
    }

    public sealed class GatewayVaultState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration properties for the Vault which can be found on the vaults' documentation page. Parsed as JSON.
        /// </summary>
        [Input("config")]
        public Input<string>? Config { get; set; }

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
        /// The description of the Vault entity.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the Vault that's going to be added. Currently, the Vault implementation must be installed in every Kong instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The unique prefix (or identifier) for this Vault configuration. The prefix is used to load the right Vault configuration and implementation when referencing secrets with the other entities.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// An optional set of strings associated with the Vault for grouping and filtering.
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

        public GatewayVaultState()
        {
        }
        public static new GatewayVaultState Empty => new GatewayVaultState();
    }
}