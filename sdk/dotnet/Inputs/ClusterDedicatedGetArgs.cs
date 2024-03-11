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

    public sealed class ClusterDedicatedGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of disk I/O operations per second that are permitted on each node in the cluster. Zero indicates the cloud provider-specific default.
        /// </summary>
        [Input("diskIops")]
        public Input<int>? DiskIops { get; set; }

        /// <summary>
        /// Machine type identifier within the given cloud provider, e.g., m6.xlarge, n2-standard-4.
        /// </summary>
        [Input("machineType")]
        public Input<string>? MachineType { get; set; }

        /// <summary>
        /// Memory per node in GiB.
        /// </summary>
        [Input("memoryGib")]
        public Input<double>? MemoryGib { get; set; }

        /// <summary>
        /// Number of virtual CPUs per node in the cluster.
        /// </summary>
        [Input("numVirtualCpus")]
        public Input<int>? NumVirtualCpus { get; set; }

        /// <summary>
        /// Set to true to assign private IP addresses to nodes. Required for CMEK and other advanced networking features.
        /// </summary>
        [Input("privateNetworkVisibility")]
        public Input<bool>? PrivateNetworkVisibility { get; set; }

        /// <summary>
        /// Storage amount per node in GiB.
        /// </summary>
        [Input("storageGib")]
        public Input<int>? StorageGib { get; set; }

        public ClusterDedicatedGetArgs()
        {
        }
        public static new ClusterDedicatedGetArgs Empty => new ClusterDedicatedGetArgs();
    }
}
