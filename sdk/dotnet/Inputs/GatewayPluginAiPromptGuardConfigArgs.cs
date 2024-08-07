// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayPluginAiPromptGuardConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If true, will ignore all previous chat prompts from the conversation history.
        /// </summary>
        [Input("allowAllConversationHistory")]
        public Input<bool>? AllowAllConversationHistory { get; set; }

        [Input("allowPatterns")]
        private InputList<string>? _allowPatterns;

        /// <summary>
        /// Array of valid regex patterns, or valid questions from the 'user' role in chat.
        /// </summary>
        public InputList<string> AllowPatterns
        {
            get => _allowPatterns ?? (_allowPatterns = new InputList<string>());
            set => _allowPatterns = value;
        }

        [Input("denyPatterns")]
        private InputList<string>? _denyPatterns;

        /// <summary>
        /// Array of invalid regex patterns, or invalid questions from the 'user' role in chat.
        /// </summary>
        public InputList<string> DenyPatterns
        {
            get => _denyPatterns ?? (_denyPatterns = new InputList<string>());
            set => _denyPatterns = value;
        }

        public GatewayPluginAiPromptGuardConfigArgs()
        {
        }
        public static new GatewayPluginAiPromptGuardConfigArgs Empty => new GatewayPluginAiPromptGuardConfigArgs();
    }
}
