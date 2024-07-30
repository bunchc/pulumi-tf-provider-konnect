// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class PortalAppearanceCustomFontsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the font to render in the browser. Not Null; must be one of ["Roboto", "Inter", "Open Sans", "Lato", "Slabo 27px", "Slabo 13px", "Oswald", "Source Sans Pro", "Montserrat", "Raleway", "PT Sans", "Lora", "Roboto Mono", "Inconsolata", "Source Code Pro", "PT Mono", "Ubuntu Mono", "IBM Plex Mono"]
        /// </summary>
        [Input("base")]
        public Input<string>? Base { get; set; }

        /// <summary>
        /// The name of the font to render in the browser. Not Null; must be one of ["Roboto", "Inter", "Open Sans", "Lato", "Slabo 27px", "Slabo 13px", "Oswald", "Source Sans Pro", "Montserrat", "Raleway", "PT Sans", "Lora", "Roboto Mono", "Inconsolata", "Source Code Pro", "PT Mono", "Ubuntu Mono", "IBM Plex Mono"]
        /// </summary>
        [Input("code")]
        public Input<string>? Code { get; set; }

        /// <summary>
        /// The name of the font to render in the browser. Not Null; must be one of ["Roboto", "Inter", "Open Sans", "Lato", "Slabo 27px", "Slabo 13px", "Oswald", "Source Sans Pro", "Montserrat", "Raleway", "PT Sans", "Lora", "Roboto Mono", "Inconsolata", "Source Code Pro", "PT Mono", "Ubuntu Mono", "IBM Plex Mono"]
        /// </summary>
        [Input("headings")]
        public Input<string>? Headings { get; set; }

        public PortalAppearanceCustomFontsGetArgs()
        {
        }
        public static new PortalAppearanceCustomFontsGetArgs Empty => new PortalAppearanceCustomFontsGetArgs();
    }
}
