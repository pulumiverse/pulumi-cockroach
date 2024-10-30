// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Cockroach
{
    /// <summary>
    /// CockroachDB Cloud cluster. Can be Dedicated or Serverless.
    /// </summary>
    [CockroachResourceType("cockroach:index/cluster:Cluster")]
    public partial class Cluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The cloud provider account ID that hosts the cluster. Needed for CMEK and other advanced features.
        /// </summary>
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        /// <summary>
        /// Cloud provider used to host the cluster. Allowed values are: * GCP * AWS * AZURE
        /// </summary>
        [Output("cloudProvider")]
        public Output<string> CloudProvider { get; private set; } = null!;

        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Major version of CockroachDB running on the cluster.
        /// </summary>
        [Output("cockroachVersion")]
        public Output<string> CockroachVersion { get; private set; } = null!;

        /// <summary>
        /// ID of the user who created the cluster.
        /// </summary>
        [Output("creatorId")]
        public Output<string> CreatorId { get; private set; } = null!;

        [Output("dedicated")]
        public Output<Outputs.ClusterDedicated?> Dedicated { get; private set; } = null!;

        /// <summary>
        /// Set to true to enable delete protection on the cluster. If unset, the server chooses the value on cluster creation, and
        /// preserves the value on cluster update.
        /// </summary>
        [Output("deleteProtection")]
        public Output<bool> DeleteProtection { get; private set; } = null!;

        /// <summary>
        /// Name of the cluster.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Describes the current long-running operation, if any.
        /// </summary>
        [Output("operationStatus")]
        public Output<string> OperationStatus { get; private set; } = null!;

        /// <summary>
        /// The ID of the cluster's parent folder. 'root' is used for a cluster at the root level.
        /// </summary>
        [Output("parentId")]
        public Output<string> ParentId { get; private set; } = null!;

        /// <summary>
        /// Denotes cluster deployment type: 'DEDICATED' or 'SERVERLESS'.
        /// </summary>
        [Output("plan")]
        public Output<string> Plan { get; private set; } = null!;

        [Output("regions")]
        public Output<ImmutableArray<Outputs.ClusterRegion>> Regions { get; private set; } = null!;

        [Output("serverless")]
        public Output<Outputs.ClusterServerless?> Serverless { get; private set; } = null!;

        /// <summary>
        /// Describes whether the cluster is being created, updated, deleted, etc.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Describes the status of any in-progress CockroachDB upgrade or rollback.
        /// </summary>
        [Output("upgradeStatus")]
        public Output<string> UpgradeStatus { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs args, CustomResourceOptions? options = null)
            : base("cockroach:index/cluster:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
            : base("cockroach:index/cluster:Cluster", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Cluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cluster Get(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new Cluster(name, id, state, options);
        }
    }

    public sealed class ClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cloud provider used to host the cluster. Allowed values are: * GCP * AWS * AZURE
        /// </summary>
        [Input("cloudProvider", required: true)]
        public Input<string> CloudProvider { get; set; } = null!;

        /// <summary>
        /// Major version of CockroachDB running on the cluster.
        /// </summary>
        [Input("cockroachVersion")]
        public Input<string>? CockroachVersion { get; set; }

        [Input("dedicated")]
        public Input<Inputs.ClusterDedicatedArgs>? Dedicated { get; set; }

        /// <summary>
        /// Set to true to enable delete protection on the cluster. If unset, the server chooses the value on cluster creation, and
        /// preserves the value on cluster update.
        /// </summary>
        [Input("deleteProtection")]
        public Input<bool>? DeleteProtection { get; set; }

        /// <summary>
        /// Name of the cluster.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The ID of the cluster's parent folder. 'root' is used for a cluster at the root level.
        /// </summary>
        [Input("parentId")]
        public Input<string>? ParentId { get; set; }

        [Input("regions", required: true)]
        private InputList<Inputs.ClusterRegionArgs>? _regions;
        public InputList<Inputs.ClusterRegionArgs> Regions
        {
            get => _regions ?? (_regions = new InputList<Inputs.ClusterRegionArgs>());
            set => _regions = value;
        }

        [Input("serverless")]
        public Input<Inputs.ClusterServerlessArgs>? Serverless { get; set; }

        public ClusterArgs()
        {
        }
        public static new ClusterArgs Empty => new ClusterArgs();
    }

    public sealed class ClusterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cloud provider account ID that hosts the cluster. Needed for CMEK and other advanced features.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// Cloud provider used to host the cluster. Allowed values are: * GCP * AWS * AZURE
        /// </summary>
        [Input("cloudProvider")]
        public Input<string>? CloudProvider { get; set; }

        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Major version of CockroachDB running on the cluster.
        /// </summary>
        [Input("cockroachVersion")]
        public Input<string>? CockroachVersion { get; set; }

        /// <summary>
        /// ID of the user who created the cluster.
        /// </summary>
        [Input("creatorId")]
        public Input<string>? CreatorId { get; set; }

        [Input("dedicated")]
        public Input<Inputs.ClusterDedicatedGetArgs>? Dedicated { get; set; }

        /// <summary>
        /// Set to true to enable delete protection on the cluster. If unset, the server chooses the value on cluster creation, and
        /// preserves the value on cluster update.
        /// </summary>
        [Input("deleteProtection")]
        public Input<bool>? DeleteProtection { get; set; }

        /// <summary>
        /// Name of the cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Describes the current long-running operation, if any.
        /// </summary>
        [Input("operationStatus")]
        public Input<string>? OperationStatus { get; set; }

        /// <summary>
        /// The ID of the cluster's parent folder. 'root' is used for a cluster at the root level.
        /// </summary>
        [Input("parentId")]
        public Input<string>? ParentId { get; set; }

        /// <summary>
        /// Denotes cluster deployment type: 'DEDICATED' or 'SERVERLESS'.
        /// </summary>
        [Input("plan")]
        public Input<string>? Plan { get; set; }

        [Input("regions")]
        private InputList<Inputs.ClusterRegionGetArgs>? _regions;
        public InputList<Inputs.ClusterRegionGetArgs> Regions
        {
            get => _regions ?? (_regions = new InputList<Inputs.ClusterRegionGetArgs>());
            set => _regions = value;
        }

        [Input("serverless")]
        public Input<Inputs.ClusterServerlessGetArgs>? Serverless { get; set; }

        /// <summary>
        /// Describes whether the cluster is being created, updated, deleted, etc.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Describes the status of any in-progress CockroachDB upgrade or rollback.
        /// </summary>
        [Input("upgradeStatus")]
        public Input<string>? UpgradeStatus { get; set; }

        public ClusterState()
        {
        }
        public static new ClusterState Empty => new ClusterState();
    }
}
