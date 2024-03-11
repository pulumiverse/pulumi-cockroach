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

    public sealed class ClusterServerlessUsageLimitsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum number of Request Units that the cluster can consume during the month.
        /// </summary>
        [Input("requestUnitLimit", required: true)]
        public Input<int> RequestUnitLimit { get; set; } = null!;

        /// <summary>
        /// Maximum amount of storage (in MiB) that the cluster can have at any time during the month.
        /// </summary>
        [Input("storageMibLimit", required: true)]
        public Input<int> StorageMibLimit { get; set; } = null!;

        public ClusterServerlessUsageLimitsArgs()
        {
        }
        public static new ClusterServerlessUsageLimitsArgs Empty => new ClusterServerlessUsageLimitsArgs();
    }
}
