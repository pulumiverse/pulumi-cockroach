// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Utility resource that represents the one-time action of finalizing a cluster's pending CockroachDB version upgrade.
 */
export class FinalizeVersionUpgrade extends pulumi.CustomResource {
    /**
     * Get an existing FinalizeVersionUpgrade resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FinalizeVersionUpgradeState, opts?: pulumi.CustomResourceOptions): FinalizeVersionUpgrade {
        return new FinalizeVersionUpgrade(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cockroach:index/finalizeVersionUpgrade:FinalizeVersionUpgrade';

    /**
     * Returns true if the given object is an instance of FinalizeVersionUpgrade.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FinalizeVersionUpgrade {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FinalizeVersionUpgrade.__pulumiType;
    }

    /**
     * Cluster ID.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Major version of the cluster to be finalized.
     */
    public readonly cockroachVersion!: pulumi.Output<string>;

    /**
     * Create a FinalizeVersionUpgrade resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FinalizeVersionUpgradeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FinalizeVersionUpgradeArgs | FinalizeVersionUpgradeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FinalizeVersionUpgradeState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["cockroachVersion"] = state ? state.cockroachVersion : undefined;
        } else {
            const args = argsOrState as FinalizeVersionUpgradeArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.cockroachVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cockroachVersion'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["cockroachVersion"] = args ? args.cockroachVersion : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FinalizeVersionUpgrade.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FinalizeVersionUpgrade resources.
 */
export interface FinalizeVersionUpgradeState {
    /**
     * Cluster ID.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Major version of the cluster to be finalized.
     */
    cockroachVersion?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FinalizeVersionUpgrade resource.
 */
export interface FinalizeVersionUpgradeArgs {
    /**
     * Cluster ID.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Major version of the cluster to be finalized.
     */
    cockroachVersion: pulumi.Input<string>;
}
