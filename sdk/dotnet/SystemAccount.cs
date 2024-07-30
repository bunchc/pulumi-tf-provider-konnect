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
    /// SystemAccount Resource
    /// </summary>
    [KonnectResourceType("konnect:index/systemAccount:SystemAccount")]
    public partial class SystemAccount : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Timestamp of when the system account was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Description of the system account. Useful when the system account name is not sufficient to differentiate one system account from another.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The system account is managed by Konnect (true/false). Requires replacement if changed.
        /// </summary>
        [Output("konnectManaged")]
        public Output<bool> KonnectManaged { get; private set; } = null!;

        /// <summary>
        /// Name of the system account.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Timestamp of when the system account was last updated.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a SystemAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SystemAccount(string name, SystemAccountArgs args, CustomResourceOptions? options = null)
            : base("konnect:index/systemAccount:SystemAccount", name, args ?? new SystemAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SystemAccount(string name, Input<string> id, SystemAccountState? state = null, CustomResourceOptions? options = null)
            : base("konnect:index/systemAccount:SystemAccount", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SystemAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SystemAccount Get(string name, Input<string> id, SystemAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new SystemAccount(name, id, state, options);
        }
    }

    public sealed class SystemAccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the system account. Useful when the system account name is not sufficient to differentiate one system account from another.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// The system account is managed by Konnect (true/false). Requires replacement if changed.
        /// </summary>
        [Input("konnectManaged")]
        public Input<bool>? KonnectManaged { get; set; }

        /// <summary>
        /// Name of the system account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SystemAccountArgs()
        {
        }
        public static new SystemAccountArgs Empty => new SystemAccountArgs();
    }

    public sealed class SystemAccountState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Timestamp of when the system account was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Description of the system account. Useful when the system account name is not sufficient to differentiate one system account from another.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The system account is managed by Konnect (true/false). Requires replacement if changed.
        /// </summary>
        [Input("konnectManaged")]
        public Input<bool>? KonnectManaged { get; set; }

        /// <summary>
        /// Name of the system account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Timestamp of when the system account was last updated.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public SystemAccountState()
        {
        }
        public static new SystemAccountState Empty => new SystemAccountState();
    }
}
