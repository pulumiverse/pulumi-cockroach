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
    public sealed class GetCockroachClusterServerlessResult
    {
        /// <summary>
        /// Cluster identifier in a connection string.
        /// </summary>
        public readonly string RoutingId;
        /// <summary>
        /// Spend limit in US cents.
        /// </summary>
        public readonly int SpendLimit;
        /// <summary>
        /// Dictates the behavior of cockroach major version upgrades.
        /// </summary>
        public readonly string UpgradeType;
        public readonly Outputs.GetCockroachClusterServerlessUsageLimitsResult UsageLimits;

        [OutputConstructor]
        private GetCockroachClusterServerlessResult(
            string routingId,

            int spendLimit,

            string upgradeType,

            Outputs.GetCockroachClusterServerlessUsageLimitsResult usageLimits)
        {
            RoutingId = routingId;
            SpendLimit = spendLimit;
            UpgradeType = upgradeType;
            UsageLimits = usageLimits;
        }
    }
}
