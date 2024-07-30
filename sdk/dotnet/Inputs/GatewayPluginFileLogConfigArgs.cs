// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class GatewayPluginFileLogConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("customFieldsByLua")]
        private InputMap<string>? _customFieldsByLua;

        /// <summary>
        /// Lua code as a key-value map
        /// </summary>
        public InputMap<string> CustomFieldsByLua
        {
            get => _customFieldsByLua ?? (_customFieldsByLua = new InputMap<string>());
            set => _customFieldsByLua = value;
        }

        /// <summary>
        /// The file path of the output log file. The plugin creates the log file if it doesn't exist yet.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Determines whether the log file is closed and reopened on every request.
        /// </summary>
        [Input("reopen")]
        public Input<bool>? Reopen { get; set; }

        public GatewayPluginFileLogConfigArgs()
        {
        }
        public static new GatewayPluginFileLogConfigArgs Empty => new GatewayPluginFileLogConfigArgs();
    }
}
