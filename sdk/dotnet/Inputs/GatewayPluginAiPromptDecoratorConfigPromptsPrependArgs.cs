// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayPluginAiPromptDecoratorConfigPromptsPrependArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Not Null
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// must be one of ["system", "assistant", "user"]
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public GatewayPluginAiPromptDecoratorConfigPromptsPrependArgs()
        {
        }
        public static new GatewayPluginAiPromptDecoratorConfigPromptsPrependArgs Empty => new GatewayPluginAiPromptDecoratorConfigPromptsPrependArgs();
    }
}
