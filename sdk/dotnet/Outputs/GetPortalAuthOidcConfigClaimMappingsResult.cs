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
    public sealed class GetPortalAuthOidcConfigClaimMappingsResult
    {
        public readonly string Email;
        public readonly string Groups;
        public readonly string Name;

        [OutputConstructor]
        private GetPortalAuthOidcConfigClaimMappingsResult(
            string email,

            string groups,

            string name)
        {
            Email = email;
            Groups = groups;
            Name = name;
        }
    }
}
