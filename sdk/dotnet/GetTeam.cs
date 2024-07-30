// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect
{
    public static class GetTeam
    {
        public static Task<GetTeamResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTeamResult>("konnect:index/getTeam:getTeam", InvokeArgs.Empty, options.WithDefaults());

        public static Output<GetTeamResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTeamResult>("konnect:index/getTeam:getTeam", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetTeamResult
    {
        public readonly string CreatedAt;
        public readonly string Description;
        public readonly string Id;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly string Name;
        public readonly bool SystemTeam;
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetTeamResult(
            string createdAt,

            string description,

            string id,

            ImmutableDictionary<string, string> labels,

            string name,

            bool systemTeam,

            string updatedAt)
        {
            CreatedAt = createdAt;
            Description = description;
            Id = id;
            Labels = labels;
            Name = name;
            SystemTeam = systemTeam;
            UpdatedAt = updatedAt;
        }
    }
}