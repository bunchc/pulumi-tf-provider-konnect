// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class PortalAppearanceCustomThemeColorsTextSecondaryArgs : global::Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Not Null
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public PortalAppearanceCustomThemeColorsTextSecondaryArgs()
        {
        }
        public static new PortalAppearanceCustomThemeColorsTextSecondaryArgs Empty => new PortalAppearanceCustomThemeColorsTextSecondaryArgs();
    }
}
