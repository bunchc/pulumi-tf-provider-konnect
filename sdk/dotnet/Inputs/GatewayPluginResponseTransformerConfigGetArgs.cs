// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayPluginResponseTransformerConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("add")]
        public Input<Inputs.GatewayPluginResponseTransformerConfigAddGetArgs>? Add { get; set; }

        [Input("append")]
        public Input<Inputs.GatewayPluginResponseTransformerConfigAppendGetArgs>? Append { get; set; }

        [Input("remove")]
        public Input<Inputs.GatewayPluginResponseTransformerConfigRemoveGetArgs>? Remove { get; set; }

        [Input("rename")]
        public Input<Inputs.GatewayPluginResponseTransformerConfigRenameGetArgs>? Rename { get; set; }

        [Input("replace")]
        public Input<Inputs.GatewayPluginResponseTransformerConfigReplaceGetArgs>? Replace { get; set; }

        public GatewayPluginResponseTransformerConfigGetArgs()
        {
        }
        public static new GatewayPluginResponseTransformerConfigGetArgs Empty => new GatewayPluginResponseTransformerConfigGetArgs();
    }
}
