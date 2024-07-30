// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class PortalAppearanceImagesFaviconGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// must be a data URL with base64 image data, e.g., data:image/jpeg;base64,\n\n. Not Null
        /// </summary>
        [Input("data")]
        public Input<string>? Data { get; set; }

        [Input("filename")]
        public Input<string>? Filename { get; set; }

        public PortalAppearanceImagesFaviconGetArgs()
        {
        }
        public static new PortalAppearanceImagesFaviconGetArgs Empty => new PortalAppearanceImagesFaviconGetArgs();
    }
}
