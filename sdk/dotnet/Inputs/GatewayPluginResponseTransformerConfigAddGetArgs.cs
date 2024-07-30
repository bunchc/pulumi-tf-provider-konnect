// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayPluginResponseTransformerConfigAddGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("headers")]
        private InputList<string>? _headers;
        public InputList<string> Headers
        {
            get => _headers ?? (_headers = new InputList<string>());
            set => _headers = value;
        }

        [Input("jsonTypes")]
        private InputList<string>? _jsonTypes;

        /// <summary>
        /// List of JSON type names. Specify the types of the JSON values returned when appending
        /// JSON properties. Each string element can be one of: boolean, number, or string.
        /// </summary>
        public InputList<string> JsonTypes
        {
            get => _jsonTypes ?? (_jsonTypes = new InputList<string>());
            set => _jsonTypes = value;
        }

        [Input("jsons")]
        private InputList<string>? _jsons;
        public InputList<string> Jsons
        {
            get => _jsons ?? (_jsons = new InputList<string>());
            set => _jsons = value;
        }

        public GatewayPluginResponseTransformerConfigAddGetArgs()
        {
        }
        public static new GatewayPluginResponseTransformerConfigAddGetArgs Empty => new GatewayPluginResponseTransformerConfigAddGetArgs();
    }
}
