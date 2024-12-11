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

    public sealed class CmekAdditionalRegionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Internal DNS name of the cluster within the cloud provider's network. Used to connect to the cluster with PrivateLink or VPC peering.
        /// </summary>
        [Input("internalDns")]
        public Input<string>? InternalDns { get; set; }

        /// <summary>
        /// Name of the region. Should match the region code used by the cluster's cloud provider.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Number of nodes in the region. Valid for Advanced clusters only.
        /// </summary>
        [Input("nodeCount")]
        public Input<int>? NodeCount { get; set; }

        /// <summary>
        /// Set to true to mark this region as the primary for a serverless cluster. Exactly one region must be primary. Dedicated clusters expect to have no primary region.
        /// </summary>
        [Input("primary")]
        public Input<bool>? Primary { get; set; }

        /// <summary>
        /// DNS name of the cluster's SQL interface. Used to connect to the cluster with IP allowlisting.
        /// </summary>
        [Input("sqlDns")]
        public Input<string>? SqlDns { get; set; }

        /// <summary>
        /// DNS name used when connecting to the DB Console for the cluster.
        /// </summary>
        [Input("uiDns")]
        public Input<string>? UiDns { get; set; }

        public CmekAdditionalRegionArgs()
        {
        }
        public static new CmekAdditionalRegionArgs Empty => new CmekAdditionalRegionArgs();
    }
}
