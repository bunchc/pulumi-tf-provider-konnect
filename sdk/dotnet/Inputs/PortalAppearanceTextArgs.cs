// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class PortalAppearanceTextArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Not Null
        /// </summary>
        [Input("catalog")]
        public Input<Inputs.PortalAppearanceTextCatalogArgs>? Catalog { get; set; }

        public PortalAppearanceTextArgs()
        {
        }
        public static new PortalAppearanceTextArgs Empty => new PortalAppearanceTextArgs();
    }
}
