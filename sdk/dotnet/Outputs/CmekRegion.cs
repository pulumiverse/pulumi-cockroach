// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Cockroach.Outputs
{

    [OutputType]
    public sealed class CmekRegion
    {
        public readonly Outputs.CmekRegionKey Key;
        /// <summary>
        /// Cloud provider region code.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Describes the status of the current encryption key within the region.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private CmekRegion(
            Outputs.CmekRegionKey key,

            string region,

            string? status)
        {
            Key = key;
            Region = region;
            Status = status;
        }
    }
}