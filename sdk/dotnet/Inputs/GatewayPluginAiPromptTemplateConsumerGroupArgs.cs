// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayPluginAiPromptTemplateConsumerGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        public GatewayPluginAiPromptTemplateConsumerGroupArgs()
        {
        }
        public static new GatewayPluginAiPromptTemplateConsumerGroupArgs Empty => new GatewayPluginAiPromptTemplateConsumerGroupArgs();
    }
}
