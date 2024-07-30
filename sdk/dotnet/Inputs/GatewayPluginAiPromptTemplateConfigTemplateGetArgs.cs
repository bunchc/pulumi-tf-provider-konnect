// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayPluginAiPromptTemplateConfigTemplateGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique name for the template, can be called with `{template://NAME}`. Not Null
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Template string for this request, supports mustache-style `{{placeholders}}`. Not Null
        /// </summary>
        [Input("template")]
        public Input<string>? Template { get; set; }

        public GatewayPluginAiPromptTemplateConfigTemplateGetArgs()
        {
        }
        public static new GatewayPluginAiPromptTemplateConfigTemplateGetArgs Empty => new GatewayPluginAiPromptTemplateConfigTemplateGetArgs();
    }
}
