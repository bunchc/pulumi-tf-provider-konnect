// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayPluginRequestTransformerConfigRemoveArgs : global::Pulumi.ResourceArgs
    {
        [Input("bodies")]
        private InputList<string>? _bodies;
        public InputList<string> Bodies
        {
            get => _bodies ?? (_bodies = new InputList<string>());
            set => _bodies = value;
        }

        [Input("headers")]
        private InputList<string>? _headers;
        public InputList<string> Headers
        {
            get => _headers ?? (_headers = new InputList<string>());
            set => _headers = value;
        }

        [Input("querystrings")]
        private InputList<string>? _querystrings;
        public InputList<string> Querystrings
        {
            get => _querystrings ?? (_querystrings = new InputList<string>());
            set => _querystrings = value;
        }

        public GatewayPluginRequestTransformerConfigRemoveArgs()
        {
        }
        public static new GatewayPluginRequestTransformerConfigRemoveArgs Empty => new GatewayPluginRequestTransformerConfigRemoveArgs();
    }
}
