// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class ApiProductVersionPortalArgs : global::Pulumi.ResourceArgs
    {
        [Input("applicationRegistrationEnabled")]
        public Input<bool>? ApplicationRegistrationEnabled { get; set; }

        [Input("authStrategies")]
        private InputList<Inputs.ApiProductVersionPortalAuthStrategyArgs>? _authStrategies;
        public InputList<Inputs.ApiProductVersionPortalAuthStrategyArgs> AuthStrategies
        {
            get => _authStrategies ?? (_authStrategies = new InputList<Inputs.ApiProductVersionPortalAuthStrategyArgs>());
            set => _authStrategies = value;
        }

        [Input("autoApproveRegistration")]
        public Input<bool>? AutoApproveRegistration { get; set; }

        [Input("deprecated")]
        public Input<bool>? Deprecated { get; set; }

        [Input("portalId")]
        public Input<string>? PortalId { get; set; }

        [Input("portalName")]
        public Input<string>? PortalName { get; set; }

        [Input("portalProductVersionId")]
        public Input<string>? PortalProductVersionId { get; set; }

        /// <summary>
        /// must be one of ["published", "unpublished"]
        /// </summary>
        [Input("publishStatus")]
        public Input<string>? PublishStatus { get; set; }

        public ApiProductVersionPortalArgs()
        {
        }
        public static new ApiProductVersionPortalArgs Empty => new ApiProductVersionPortalArgs();
    }
}
