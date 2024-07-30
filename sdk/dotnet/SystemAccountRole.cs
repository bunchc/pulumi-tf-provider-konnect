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
    /// SystemAccountRole Resource
    /// </summary>
    [KonnectResourceType("konnect:index/systemAccountRole:SystemAccountRole")]
    public partial class SystemAccountRole : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the system account. Requires replacement if changed.
        /// </summary>
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        /// <summary>
        /// The ID of the entity. Requires replacement if changed.
        /// </summary>
        [Output("entityId")]
        public Output<string> EntityId { get; private set; } = null!;

        /// <summary>
        /// The region of the team. Requires replacement if changed. ; must be one of ["us", "eu", "au", "*"]
        /// </summary>
        [Output("entityRegion")]
        public Output<string> EntityRegion { get; private set; } = null!;

        /// <summary>
        /// The type of entity. Requires replacement if changed. ; must be one of ["API Products", "Application Auth Strategies", "Audit Logs", "Control Planes", "DCR Providers", "Identity", "Mesh Control Planes", "Networks", "Portals", "Service Hub"]
        /// </summary>
        [Output("entityTypeName")]
        public Output<string> EntityTypeName { get; private set; } = null!;

        /// <summary>
        /// The desired role. Requires replacement if changed. ; must be one of ["Admin", "Appearance Maintainer", "Application Registration", "Certificate Admin", "Cloud Gateway Cluster Admin", "Cloud Gateway Cluster Viewer", "Consumer Admin", "Creator", "Deployer", "Discovery Admin", "Discovery Viewer", "Gateway Service Admin", "Integration Admin", "Integration Viewer", "Key Admin", "Maintainer", "Network Admin", "Network Creator", "Network Viewer", "Plugin Admin", "Plugins Admin", "Product Publisher", "Publisher", "Route Admin", "SNI Admin", "Service Admin", "Service Creator", "Service Viewer", "Upstream Admin", "Vault Admin", "Viewer"]
        /// </summary>
        [Output("roleName")]
        public Output<string> RoleName { get; private set; } = null!;


        /// <summary>
        /// Create a SystemAccountRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SystemAccountRole(string name, SystemAccountRoleArgs args, CustomResourceOptions? options = null)
            : base("konnect:index/systemAccountRole:SystemAccountRole", name, args ?? new SystemAccountRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SystemAccountRole(string name, Input<string> id, SystemAccountRoleState? state = null, CustomResourceOptions? options = null)
            : base("konnect:index/systemAccountRole:SystemAccountRole", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SystemAccountRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SystemAccountRole Get(string name, Input<string> id, SystemAccountRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new SystemAccountRole(name, id, state, options);
        }
    }

    public sealed class SystemAccountRoleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the system account. Requires replacement if changed.
        /// </summary>
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        /// <summary>
        /// The ID of the entity. Requires replacement if changed.
        /// </summary>
        [Input("entityId")]
        public Input<string>? EntityId { get; set; }

        /// <summary>
        /// The region of the team. Requires replacement if changed. ; must be one of ["us", "eu", "au", "*"]
        /// </summary>
        [Input("entityRegion")]
        public Input<string>? EntityRegion { get; set; }

        /// <summary>
        /// The type of entity. Requires replacement if changed. ; must be one of ["API Products", "Application Auth Strategies", "Audit Logs", "Control Planes", "DCR Providers", "Identity", "Mesh Control Planes", "Networks", "Portals", "Service Hub"]
        /// </summary>
        [Input("entityTypeName")]
        public Input<string>? EntityTypeName { get; set; }

        /// <summary>
        /// The desired role. Requires replacement if changed. ; must be one of ["Admin", "Appearance Maintainer", "Application Registration", "Certificate Admin", "Cloud Gateway Cluster Admin", "Cloud Gateway Cluster Viewer", "Consumer Admin", "Creator", "Deployer", "Discovery Admin", "Discovery Viewer", "Gateway Service Admin", "Integration Admin", "Integration Viewer", "Key Admin", "Maintainer", "Network Admin", "Network Creator", "Network Viewer", "Plugin Admin", "Plugins Admin", "Product Publisher", "Publisher", "Route Admin", "SNI Admin", "Service Admin", "Service Creator", "Service Viewer", "Upstream Admin", "Vault Admin", "Viewer"]
        /// </summary>
        [Input("roleName")]
        public Input<string>? RoleName { get; set; }

        public SystemAccountRoleArgs()
        {
        }
        public static new SystemAccountRoleArgs Empty => new SystemAccountRoleArgs();
    }

    public sealed class SystemAccountRoleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the system account. Requires replacement if changed.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// The ID of the entity. Requires replacement if changed.
        /// </summary>
        [Input("entityId")]
        public Input<string>? EntityId { get; set; }

        /// <summary>
        /// The region of the team. Requires replacement if changed. ; must be one of ["us", "eu", "au", "*"]
        /// </summary>
        [Input("entityRegion")]
        public Input<string>? EntityRegion { get; set; }

        /// <summary>
        /// The type of entity. Requires replacement if changed. ; must be one of ["API Products", "Application Auth Strategies", "Audit Logs", "Control Planes", "DCR Providers", "Identity", "Mesh Control Planes", "Networks", "Portals", "Service Hub"]
        /// </summary>
        [Input("entityTypeName")]
        public Input<string>? EntityTypeName { get; set; }

        /// <summary>
        /// The desired role. Requires replacement if changed. ; must be one of ["Admin", "Appearance Maintainer", "Application Registration", "Certificate Admin", "Cloud Gateway Cluster Admin", "Cloud Gateway Cluster Viewer", "Consumer Admin", "Creator", "Deployer", "Discovery Admin", "Discovery Viewer", "Gateway Service Admin", "Integration Admin", "Integration Viewer", "Key Admin", "Maintainer", "Network Admin", "Network Creator", "Network Viewer", "Plugin Admin", "Plugins Admin", "Product Publisher", "Publisher", "Route Admin", "SNI Admin", "Service Admin", "Service Creator", "Service Viewer", "Upstream Admin", "Vault Admin", "Viewer"]
        /// </summary>
        [Input("roleName")]
        public Input<string>? RoleName { get; set; }

        public SystemAccountRoleState()
        {
        }
        public static new SystemAccountRoleState Empty => new SystemAccountRoleState();
    }
}