// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayPluginAiPromptDecoratorConfigPromptsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("appends")]
        private InputList<Inputs.GatewayPluginAiPromptDecoratorConfigPromptsAppendGetArgs>? _appends;

        /// <summary>
        /// Insert chat messages at the end of the chat message array. This array preserves exact order when adding messages.
        /// </summary>
        public InputList<Inputs.GatewayPluginAiPromptDecoratorConfigPromptsAppendGetArgs> Appends
        {
            get => _appends ?? (_appends = new InputList<Inputs.GatewayPluginAiPromptDecoratorConfigPromptsAppendGetArgs>());
            set => _appends = value;
        }

        [Input("prepends")]
        private InputList<Inputs.GatewayPluginAiPromptDecoratorConfigPromptsPrependGetArgs>? _prepends;

        /// <summary>
        /// Insert chat messages at the beginning of the chat message array. This array preserves exact order when adding messages.
        /// </summary>
        public InputList<Inputs.GatewayPluginAiPromptDecoratorConfigPromptsPrependGetArgs> Prepends
        {
            get => _prepends ?? (_prepends = new InputList<Inputs.GatewayPluginAiPromptDecoratorConfigPromptsPrependGetArgs>());
            set => _prepends = value;
        }

        public GatewayPluginAiPromptDecoratorConfigPromptsGetArgs()
        {
        }
        public static new GatewayPluginAiPromptDecoratorConfigPromptsGetArgs Empty => new GatewayPluginAiPromptDecoratorConfigPromptsGetArgs();
    }
}