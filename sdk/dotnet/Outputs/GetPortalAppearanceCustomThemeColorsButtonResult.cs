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
    public sealed class GetPortalAppearanceCustomThemeColorsButtonResult
    {
        public readonly Outputs.GetPortalAppearanceCustomThemeColorsButtonPrimaryFillResult PrimaryFill;
        public readonly Outputs.GetPortalAppearanceCustomThemeColorsButtonPrimaryTextResult PrimaryText;

        [OutputConstructor]
        private GetPortalAppearanceCustomThemeColorsButtonResult(
            Outputs.GetPortalAppearanceCustomThemeColorsButtonPrimaryFillResult primaryFill,

            Outputs.GetPortalAppearanceCustomThemeColorsButtonPrimaryTextResult primaryText)
        {
            PrimaryFill = primaryFill;
            PrimaryText = primaryText;
        }
    }
}