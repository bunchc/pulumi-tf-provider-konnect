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
    public sealed class GetPortalAppearanceCustomThemeColorsTextLinkResult
    {
        public readonly string Description;
        public readonly string Value;

        [OutputConstructor]
        private GetPortalAppearanceCustomThemeColorsTextLinkResult(
            string description,

            string value)
        {
            Description = description;
            Value = value;
        }
    }
}
