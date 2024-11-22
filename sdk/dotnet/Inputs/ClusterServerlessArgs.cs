// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Cockroach.Inputs
{

    public sealed class ClusterServerlessArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster identifier in a connection string.
        /// </summary>
        [Input("routingId")]
        public Input<string>? RoutingId { get; set; }

        /// <summary>
        /// Spend limit in US cents.
        /// </summary>
        [Input("spendLimit")]
        public Input<int>? SpendLimit { get; set; }

        /// <summary>
        /// Dictates the behavior of CockroachDB major version upgrades. Manual upgrades are not supported on CockroachDB Basic. Manual or automatic upgrades are supported on CockroachDB Standard. If you omit the field, it defaults to `AUTOMATIC`. Allowed values are:
        ///   * MANUAL
        ///   * AUTOMATIC
        /// </summary>
        [Input("upgradeType")]
        public Input<string>? UpgradeType { get; set; }

        [Input("usageLimits")]
        public Input<Inputs.ClusterServerlessUsageLimitsArgs>? UsageLimits { get; set; }

        public ClusterServerlessArgs()
        {
        }
        public static new ClusterServerlessArgs Empty => new ClusterServerlessArgs();
    }
}
