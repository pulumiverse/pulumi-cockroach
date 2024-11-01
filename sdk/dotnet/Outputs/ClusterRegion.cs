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
    public sealed class ClusterRegion
    {
        /// <summary>
        /// Internal DNS name of the cluster within the cloud provider's network. Used to connect to the cluster with PrivateLink or VPC peering.
        /// </summary>
        public readonly string? InternalDns;
        /// <summary>
        /// Name of the region. Should match the region code used by the cluster's cloud provider.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Number of nodes in the region. Will always be 0 for serverless clusters.
        /// </summary>
        public readonly int? NodeCount;
        /// <summary>
        /// Set to true to mark this region as the primary for a serverless cluster. Exactly one region must be primary. Dedicated clusters expect to have no primary region.
        /// </summary>
        public readonly bool? Primary;
        /// <summary>
        /// DNS name of the cluster's SQL interface. Used to connect to the cluster with IP allowlisting.
        /// </summary>
        public readonly string? SqlDns;
        /// <summary>
        /// DNS name used when connecting to the DB Console for the cluster.
        /// </summary>
        public readonly string? UiDns;

        [OutputConstructor]
        private ClusterRegion(
            string? internalDns,

            string name,

            int? nodeCount,

            bool? primary,

            string? sqlDns,

            string? uiDns)
        {
            InternalDns = internalDns;
            Name = name;
            NodeCount = nodeCount;
            Primary = primary;
            SqlDns = sqlDns;
            UiDns = uiDns;
        }
    }
}
