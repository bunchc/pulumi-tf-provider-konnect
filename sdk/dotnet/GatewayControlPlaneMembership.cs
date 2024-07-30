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
    /// GatewayControlPlaneMembership Resource
    /// </summary>
    [KonnectResourceType("konnect:index/gatewayControlPlaneMembership:GatewayControlPlaneMembership")]
    public partial class GatewayControlPlaneMembership : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Requires replacement if changed.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<Outputs.GatewayControlPlaneMembershipMember>> Members { get; private set; } = null!;


        /// <summary>
        /// Create a GatewayControlPlaneMembership resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GatewayControlPlaneMembership(string name, GatewayControlPlaneMembershipArgs? args = null, CustomResourceOptions? options = null)
            : base("konnect:index/gatewayControlPlaneMembership:GatewayControlPlaneMembership", name, args ?? new GatewayControlPlaneMembershipArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GatewayControlPlaneMembership(string name, Input<string> id, GatewayControlPlaneMembershipState? state = null, CustomResourceOptions? options = null)
            : base("konnect:index/gatewayControlPlaneMembership:GatewayControlPlaneMembership", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GatewayControlPlaneMembership resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GatewayControlPlaneMembership Get(string name, Input<string> id, GatewayControlPlaneMembershipState? state = null, CustomResourceOptions? options = null)
        {
            return new GatewayControlPlaneMembership(name, id, state, options);
        }
    }

    public sealed class GatewayControlPlaneMembershipArgs : global::Pulumi.ResourceArgs
    {
        [Input("members")]
        private InputList<Inputs.GatewayControlPlaneMembershipMemberArgs>? _members;

        /// <summary>
        /// Requires replacement if changed.
        /// </summary>
        public InputList<Inputs.GatewayControlPlaneMembershipMemberArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.GatewayControlPlaneMembershipMemberArgs>());
            set => _members = value;
        }

        public GatewayControlPlaneMembershipArgs()
        {
        }
        public static new GatewayControlPlaneMembershipArgs Empty => new GatewayControlPlaneMembershipArgs();
    }

    public sealed class GatewayControlPlaneMembershipState : global::Pulumi.ResourceArgs
    {
        [Input("members")]
        private InputList<Inputs.GatewayControlPlaneMembershipMemberGetArgs>? _members;

        /// <summary>
        /// Requires replacement if changed.
        /// </summary>
        public InputList<Inputs.GatewayControlPlaneMembershipMemberGetArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.GatewayControlPlaneMembershipMemberGetArgs>());
            set => _members = value;
        }

        public GatewayControlPlaneMembershipState()
        {
        }
        public static new GatewayControlPlaneMembershipState Empty => new GatewayControlPlaneMembershipState();
    }
}
