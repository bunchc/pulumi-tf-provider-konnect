// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Outputs
{

    [OutputType]
    public sealed class PortalAppearanceCustomThemeColorsSectionStroke
    {
        public readonly string? Description;
        /// <summary>
        /// Not Null
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private PortalAppearanceCustomThemeColorsSectionStroke(
            string? description,

            string? value)
        {
            Description = description;
            Value = value;
        }
    }
}
