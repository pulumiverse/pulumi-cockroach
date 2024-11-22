// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * CockroachDB Cloud cluster.
 *
 * ## Import
 *
 * format: <cluster id>
 *
 * ```sh
 * $ pulumi import cockroach:index/cluster:Cluster my_cluster 1f69fdd2-600a-4cfc-a9ba-16995df0d77d
 * ```
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cockroach:index/cluster:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * The cloud provider account ID that hosts the cluster. Needed for CMEK and other advanced features.
     */
    public /*out*/ readonly accountId!: pulumi.Output<string>;
    /**
     * The backup settings for a cluster. Each cluster has backup settings that determine if backups are enabled, how
     * frequently they are taken, and how long they are retained for. Use this attribute to manage those settings.
     */
    public readonly backupConfig!: pulumi.Output<outputs.ClusterBackupConfig>;
    /**
     * Cloud provider used to host the cluster. Allowed values are: * GCP * AWS * AZURE
     */
    public readonly cloudProvider!: pulumi.Output<string>;
    /**
     * Major version of CockroachDB running on the cluster.
     */
    public readonly cockroachVersion!: pulumi.Output<string>;
    /**
     * ID of the user who created the cluster.
     */
    public /*out*/ readonly creatorId!: pulumi.Output<string>;
    public readonly dedicated!: pulumi.Output<outputs.ClusterDedicated | undefined>;
    /**
     * Set to true to enable delete protection on the cluster. If unset, the server chooses the value on cluster creation, and
     * preserves the value on cluster update.
     */
    public readonly deleteProtection!: pulumi.Output<boolean>;
    /**
     * Name of the cluster.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Describes the current long-running operation, if any.
     */
    public /*out*/ readonly operationStatus!: pulumi.Output<string>;
    /**
     * The ID of the cluster's parent folder. 'root' is used for a cluster at the root level.
     */
    public readonly parentId!: pulumi.Output<string>;
    /**
     * Denotes cluster plan type: 'BASIC' or 'STANDARD' or 'ADVANCED'.
     */
    public readonly plan!: pulumi.Output<string>;
    public readonly regions!: pulumi.Output<outputs.ClusterRegion[]>;
    public readonly serverless!: pulumi.Output<outputs.ClusterServerless | undefined>;
    /**
     * Describes whether the cluster is being created, updated, deleted, etc.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Describes the status of any in-progress CockroachDB upgrade or rollback.
     */
    public /*out*/ readonly upgradeStatus!: pulumi.Output<string>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClusterState | undefined;
            resourceInputs["accountId"] = state ? state.accountId : undefined;
            resourceInputs["backupConfig"] = state ? state.backupConfig : undefined;
            resourceInputs["cloudProvider"] = state ? state.cloudProvider : undefined;
            resourceInputs["cockroachVersion"] = state ? state.cockroachVersion : undefined;
            resourceInputs["creatorId"] = state ? state.creatorId : undefined;
            resourceInputs["dedicated"] = state ? state.dedicated : undefined;
            resourceInputs["deleteProtection"] = state ? state.deleteProtection : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["operationStatus"] = state ? state.operationStatus : undefined;
            resourceInputs["parentId"] = state ? state.parentId : undefined;
            resourceInputs["plan"] = state ? state.plan : undefined;
            resourceInputs["regions"] = state ? state.regions : undefined;
            resourceInputs["serverless"] = state ? state.serverless : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["upgradeStatus"] = state ? state.upgradeStatus : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if ((!args || args.cloudProvider === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cloudProvider'");
            }
            if ((!args || args.regions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'regions'");
            }
            resourceInputs["backupConfig"] = args ? args.backupConfig : undefined;
            resourceInputs["cloudProvider"] = args ? args.cloudProvider : undefined;
            resourceInputs["cockroachVersion"] = args ? args.cockroachVersion : undefined;
            resourceInputs["dedicated"] = args ? args.dedicated : undefined;
            resourceInputs["deleteProtection"] = args ? args.deleteProtection : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parentId"] = args ? args.parentId : undefined;
            resourceInputs["plan"] = args ? args.plan : undefined;
            resourceInputs["regions"] = args ? args.regions : undefined;
            resourceInputs["serverless"] = args ? args.serverless : undefined;
            resourceInputs["accountId"] = undefined /*out*/;
            resourceInputs["creatorId"] = undefined /*out*/;
            resourceInputs["operationStatus"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["upgradeStatus"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Cluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * The cloud provider account ID that hosts the cluster. Needed for CMEK and other advanced features.
     */
    accountId?: pulumi.Input<string>;
    /**
     * The backup settings for a cluster. Each cluster has backup settings that determine if backups are enabled, how
     * frequently they are taken, and how long they are retained for. Use this attribute to manage those settings.
     */
    backupConfig?: pulumi.Input<inputs.ClusterBackupConfig>;
    /**
     * Cloud provider used to host the cluster. Allowed values are: * GCP * AWS * AZURE
     */
    cloudProvider?: pulumi.Input<string>;
    /**
     * Major version of CockroachDB running on the cluster.
     */
    cockroachVersion?: pulumi.Input<string>;
    /**
     * ID of the user who created the cluster.
     */
    creatorId?: pulumi.Input<string>;
    dedicated?: pulumi.Input<inputs.ClusterDedicated>;
    /**
     * Set to true to enable delete protection on the cluster. If unset, the server chooses the value on cluster creation, and
     * preserves the value on cluster update.
     */
    deleteProtection?: pulumi.Input<boolean>;
    /**
     * Name of the cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * Describes the current long-running operation, if any.
     */
    operationStatus?: pulumi.Input<string>;
    /**
     * The ID of the cluster's parent folder. 'root' is used for a cluster at the root level.
     */
    parentId?: pulumi.Input<string>;
    /**
     * Denotes cluster plan type: 'BASIC' or 'STANDARD' or 'ADVANCED'.
     */
    plan?: pulumi.Input<string>;
    regions?: pulumi.Input<pulumi.Input<inputs.ClusterRegion>[]>;
    serverless?: pulumi.Input<inputs.ClusterServerless>;
    /**
     * Describes whether the cluster is being created, updated, deleted, etc.
     */
    state?: pulumi.Input<string>;
    /**
     * Describes the status of any in-progress CockroachDB upgrade or rollback.
     */
    upgradeStatus?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * The backup settings for a cluster. Each cluster has backup settings that determine if backups are enabled, how
     * frequently they are taken, and how long they are retained for. Use this attribute to manage those settings.
     */
    backupConfig?: pulumi.Input<inputs.ClusterBackupConfig>;
    /**
     * Cloud provider used to host the cluster. Allowed values are: * GCP * AWS * AZURE
     */
    cloudProvider: pulumi.Input<string>;
    /**
     * Major version of CockroachDB running on the cluster.
     */
    cockroachVersion?: pulumi.Input<string>;
    dedicated?: pulumi.Input<inputs.ClusterDedicated>;
    /**
     * Set to true to enable delete protection on the cluster. If unset, the server chooses the value on cluster creation, and
     * preserves the value on cluster update.
     */
    deleteProtection?: pulumi.Input<boolean>;
    /**
     * Name of the cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the cluster's parent folder. 'root' is used for a cluster at the root level.
     */
    parentId?: pulumi.Input<string>;
    /**
     * Denotes cluster plan type: 'BASIC' or 'STANDARD' or 'ADVANCED'.
     */
    plan?: pulumi.Input<string>;
    regions: pulumi.Input<pulumi.Input<inputs.ClusterRegion>[]>;
    serverless?: pulumi.Input<inputs.ClusterServerless>;
}
