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
    public sealed class GetPortalAppearanceCustomThemeColorsTextResult
    {
        public readonly Outputs.GetPortalAppearanceCustomThemeColorsTextAccentResult Accent;
        public readonly Outputs.GetPortalAppearanceCustomThemeColorsTextFooterResult Footer;
        public readonly Outputs.GetPortalAppearanceCustomThemeColorsTextHeaderResult Header;
        public readonly Outputs.GetPortalAppearanceCustomThemeColorsTextHeadingsResult Headings;
        public readonly Outputs.GetPortalAppearanceCustomThemeColorsTextHeroResult Hero;
        public readonly Outputs.GetPortalAppearanceCustomThemeColorsTextLinkResult Link;
        public readonly Outputs.GetPortalAppearanceCustomThemeColorsTextPrimaryResult Primary;
        public readonly Outputs.GetPortalAppearanceCustomThemeColorsTextSecondaryResult Secondary;

        [OutputConstructor]
        private GetPortalAppearanceCustomThemeColorsTextResult(
            Outputs.GetPortalAppearanceCustomThemeColorsTextAccentResult accent,

            Outputs.GetPortalAppearanceCustomThemeColorsTextFooterResult footer,

            Outputs.GetPortalAppearanceCustomThemeColorsTextHeaderResult header,

            Outputs.GetPortalAppearanceCustomThemeColorsTextHeadingsResult headings,

            Outputs.GetPortalAppearanceCustomThemeColorsTextHeroResult hero,

            Outputs.GetPortalAppearanceCustomThemeColorsTextLinkResult link,

            Outputs.GetPortalAppearanceCustomThemeColorsTextPrimaryResult primary,

            Outputs.GetPortalAppearanceCustomThemeColorsTextSecondaryResult secondary)
        {
            Accent = accent;
            Footer = footer;
            Header = header;
            Headings = headings;
            Hero = hero;
            Link = link;
            Primary = primary;
            Secondary = secondary;
        }
    }
}
