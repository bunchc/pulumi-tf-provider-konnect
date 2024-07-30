// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Outputs
{

    [OutputType]
    public sealed class GatewayPluginAiPromptTemplateConfigTemplate
    {
        /// <summary>
        /// Unique name for the template, can be called with `{template://NAME}`. Not Null
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Template string for this request, supports mustache-style `{{placeholders}}`. Not Null
        /// </summary>
        public readonly string? Template;

        [OutputConstructor]
        private GatewayPluginAiPromptTemplateConfigTemplate(
            string? name,

            string? template)
        {
            Name = name;
            Template = template;
        }
    }
}
