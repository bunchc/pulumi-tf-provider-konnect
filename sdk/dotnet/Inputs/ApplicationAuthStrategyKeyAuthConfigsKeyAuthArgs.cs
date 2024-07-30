// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class ApplicationAuthStrategyKeyAuthConfigsKeyAuthArgs : global::Pulumi.ResourceArgs
    {
        [Input("keyNames")]
        private InputList<string>? _keyNames;

        /// <summary>
        /// The names of the headers containing the API key. You can specify multiple header names. Requires replacement if changed.
        /// </summary>
        public InputList<string> KeyNames
        {
            get => _keyNames ?? (_keyNames = new InputList<string>());
            set => _keyNames = value;
        }

        public ApplicationAuthStrategyKeyAuthConfigsKeyAuthArgs()
        {
        }
        public static new ApplicationAuthStrategyKeyAuthConfigsKeyAuthArgs Empty => new ApplicationAuthStrategyKeyAuthConfigsKeyAuthArgs();
    }
}
