// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayPluginAiPromptDecoratorConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("prompts")]
        public Input<Inputs.GatewayPluginAiPromptDecoratorConfigPromptsGetArgs>? Prompts { get; set; }

        public GatewayPluginAiPromptDecoratorConfigGetArgs()
        {
        }
        public static new GatewayPluginAiPromptDecoratorConfigGetArgs Empty => new GatewayPluginAiPromptDecoratorConfigGetArgs();
    }
}