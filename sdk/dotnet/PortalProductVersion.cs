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
    /// PortalProductVersion Resource
    /// </summary>
    [KonnectResourceType("konnect:index/portalProductVersion:PortalProductVersion")]
    public partial class PortalProductVersion : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether the application registration on this portal for the api product version is enabled
        /// </summary>
        [Output("applicationRegistrationEnabled")]
        public Output<bool> ApplicationRegistrationEnabled { get; private set; } = null!;

        /// <summary>
        /// A list of authentication strategies
        /// </summary>
        [Output("authStrategies")]
        public Output<ImmutableArray<Outputs.PortalProductVersionAuthStrategy>> AuthStrategies { get; private set; } = null!;

        /// <summary>
        /// A list of authentication strategy IDs
        /// </summary>
        [Output("authStrategyIds")]
        public Output<ImmutableArray<string>> AuthStrategyIds { get; private set; } = null!;

        /// <summary>
        /// Whether the application registration auto approval on this portal for the api product version is enabled
        /// </summary>
        [Output("autoApproveRegistration")]
        public Output<bool> AutoApproveRegistration { get; private set; } = null!;

        /// <summary>
        /// An ISO-8601 timestamp representation of entity creation date.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Whether the api product version on the portal is deprecated
        /// </summary>
        [Output("deprecated")]
        public Output<bool> Deprecated { get; private set; } = null!;

        /// <summary>
        /// Whether to notify developers who are affected by this change
        /// </summary>
        [Output("notifyDevelopers")]
        public Output<bool?> NotifyDevelopers { get; private set; } = null!;

        /// <summary>
        /// portal identifier
        /// </summary>
        [Output("portalId")]
        public Output<string> PortalId { get; private set; } = null!;

        /// <summary>
        /// API product version identifier
        /// </summary>
        [Output("productVersionId")]
        public Output<string> ProductVersionId { get; private set; } = null!;

        /// <summary>
        /// Publication status of the API product version on the portal. must be one of ["published", "unpublished"]
        /// </summary>
        [Output("publishStatus")]
        public Output<string> PublishStatus { get; private set; } = null!;

        /// <summary>
        /// An ISO-8601 timestamp representation of entity update date.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a PortalProductVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PortalProductVersion(string name, PortalProductVersionArgs args, CustomResourceOptions? options = null)
            : base("konnect:index/portalProductVersion:PortalProductVersion", name, args ?? new PortalProductVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PortalProductVersion(string name, Input<string> id, PortalProductVersionState? state = null, CustomResourceOptions? options = null)
            : base("konnect:index/portalProductVersion:PortalProductVersion", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PortalProductVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PortalProductVersion Get(string name, Input<string> id, PortalProductVersionState? state = null, CustomResourceOptions? options = null)
        {
            return new PortalProductVersion(name, id, state, options);
        }
    }

    public sealed class PortalProductVersionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the application registration on this portal for the api product version is enabled
        /// </summary>
        [Input("applicationRegistrationEnabled", required: true)]
        public Input<bool> ApplicationRegistrationEnabled { get; set; } = null!;

        [Input("authStrategyIds", required: true)]
        private InputList<string>? _authStrategyIds;

        /// <summary>
        /// A list of authentication strategy IDs
        /// </summary>
        public InputList<string> AuthStrategyIds
        {
            get => _authStrategyIds ?? (_authStrategyIds = new InputList<string>());
            set => _authStrategyIds = value;
        }

        /// <summary>
        /// Whether the application registration auto approval on this portal for the api product version is enabled
        /// </summary>
        [Input("autoApproveRegistration", required: true)]
        public Input<bool> AutoApproveRegistration { get; set; } = null!;

        /// <summary>
        /// Whether the api product version on the portal is deprecated
        /// </summary>
        [Input("deprecated", required: true)]
        public Input<bool> Deprecated { get; set; } = null!;

        /// <summary>
        /// Whether to notify developers who are affected by this change
        /// </summary>
        [Input("notifyDevelopers")]
        public Input<bool>? NotifyDevelopers { get; set; }

        /// <summary>
        /// portal identifier
        /// </summary>
        [Input("portalId", required: true)]
        public Input<string> PortalId { get; set; } = null!;

        /// <summary>
        /// API product version identifier
        /// </summary>
        [Input("productVersionId", required: true)]
        public Input<string> ProductVersionId { get; set; } = null!;

        /// <summary>
        /// Publication status of the API product version on the portal. must be one of ["published", "unpublished"]
        /// </summary>
        [Input("publishStatus", required: true)]
        public Input<string> PublishStatus { get; set; } = null!;

        public PortalProductVersionArgs()
        {
        }
        public static new PortalProductVersionArgs Empty => new PortalProductVersionArgs();
    }

    public sealed class PortalProductVersionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the application registration on this portal for the api product version is enabled
        /// </summary>
        [Input("applicationRegistrationEnabled")]
        public Input<bool>? ApplicationRegistrationEnabled { get; set; }

        [Input("authStrategies")]
        private InputList<Inputs.PortalProductVersionAuthStrategyGetArgs>? _authStrategies;

        /// <summary>
        /// A list of authentication strategies
        /// </summary>
        public InputList<Inputs.PortalProductVersionAuthStrategyGetArgs> AuthStrategies
        {
            get => _authStrategies ?? (_authStrategies = new InputList<Inputs.PortalProductVersionAuthStrategyGetArgs>());
            set => _authStrategies = value;
        }

        [Input("authStrategyIds")]
        private InputList<string>? _authStrategyIds;

        /// <summary>
        /// A list of authentication strategy IDs
        /// </summary>
        public InputList<string> AuthStrategyIds
        {
            get => _authStrategyIds ?? (_authStrategyIds = new InputList<string>());
            set => _authStrategyIds = value;
        }

        /// <summary>
        /// Whether the application registration auto approval on this portal for the api product version is enabled
        /// </summary>
        [Input("autoApproveRegistration")]
        public Input<bool>? AutoApproveRegistration { get; set; }

        /// <summary>
        /// An ISO-8601 timestamp representation of entity creation date.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Whether the api product version on the portal is deprecated
        /// </summary>
        [Input("deprecated")]
        public Input<bool>? Deprecated { get; set; }

        /// <summary>
        /// Whether to notify developers who are affected by this change
        /// </summary>
        [Input("notifyDevelopers")]
        public Input<bool>? NotifyDevelopers { get; set; }

        /// <summary>
        /// portal identifier
        /// </summary>
        [Input("portalId")]
        public Input<string>? PortalId { get; set; }

        /// <summary>
        /// API product version identifier
        /// </summary>
        [Input("productVersionId")]
        public Input<string>? ProductVersionId { get; set; }

        /// <summary>
        /// Publication status of the API product version on the portal. must be one of ["published", "unpublished"]
        /// </summary>
        [Input("publishStatus")]
        public Input<string>? PublishStatus { get; set; }

        /// <summary>
        /// An ISO-8601 timestamp representation of entity update date.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public PortalProductVersionState()
        {
        }
        public static new PortalProductVersionState Empty => new PortalProductVersionState();
    }
}
