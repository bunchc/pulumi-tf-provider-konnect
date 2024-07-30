// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayPluginAiPromptTemplateConfigGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set true to allow requests that don't call or match any template.
        /// </summary>
        [Input("allowUntemplatedRequests")]
        public Input<bool>? AllowUntemplatedRequests { get; set; }

        /// <summary>
        /// Set true to add the original request to the Kong log plugin(s) output.
        /// </summary>
        [Input("logOriginalRequest")]
        public Input<bool>? LogOriginalRequest { get; set; }

        [Input("templates")]
        private InputList<Inputs.GatewayPluginAiPromptTemplateConfigTemplateGetArgs>? _templates;

        /// <summary>
        /// Array of templates available to the request context.
        /// </summary>
        public InputList<Inputs.GatewayPluginAiPromptTemplateConfigTemplateGetArgs> Templates
        {
            get => _templates ?? (_templates = new InputList<Inputs.GatewayPluginAiPromptTemplateConfigTemplateGetArgs>());
            set => _templates = value;
        }

        public GatewayPluginAiPromptTemplateConfigGetArgs()
        {
        }
        public static new GatewayPluginAiPromptTemplateConfigGetArgs Empty => new GatewayPluginAiPromptTemplateConfigGetArgs();
    }
}