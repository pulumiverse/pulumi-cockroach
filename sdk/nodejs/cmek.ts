// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Customer-managed encryption keys (CMEK) resource for a single cluster.
 */
export class Cmek extends pulumi.CustomResource {
    /**
     * Get an existing Cmek resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CmekState, opts?: pulumi.CustomResourceOptions): Cmek {
        return new Cmek(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cockroach:index/cmek:Cmek';

    /**
     * Returns true if the given object is an instance of Cmek.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cmek {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cmek.__pulumiType;
    }

    /**
     * Once CMEK is enabled for a cluster, no new regions can be added to the cluster resource, since they need encryption key
     * info stored in the CMEK resource. New regions can be added and maintained here instead.
     */
    public readonly additionalRegions!: pulumi.Output<outputs.CmekAdditionalRegion[] | undefined>;
    /**
     * Cluster ID.
     */
    public readonly clusterId!: pulumi.Output<string>;
    public readonly regions!: pulumi.Output<outputs.CmekRegion[]>;
    /**
     * Aggregated status of the cluster's encryption key(s).
     */
    public readonly status!: pulumi.Output<string>;

    /**
     * Create a Cmek resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CmekArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CmekArgs | CmekState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CmekState | undefined;
            resourceInputs["additionalRegions"] = state ? state.additionalRegions : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["regions"] = state ? state.regions : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as CmekArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.regions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'regions'");
            }
            resourceInputs["additionalRegions"] = args ? args.additionalRegions : undefined;
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["regions"] = args ? args.regions : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Cmek.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cmek resources.
 */
export interface CmekState {
    /**
     * Once CMEK is enabled for a cluster, no new regions can be added to the cluster resource, since they need encryption key
     * info stored in the CMEK resource. New regions can be added and maintained here instead.
     */
    additionalRegions?: pulumi.Input<pulumi.Input<inputs.CmekAdditionalRegion>[]>;
    /**
     * Cluster ID.
     */
    clusterId?: pulumi.Input<string>;
    regions?: pulumi.Input<pulumi.Input<inputs.CmekRegion>[]>;
    /**
     * Aggregated status of the cluster's encryption key(s).
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cmek resource.
 */
export interface CmekArgs {
    /**
     * Once CMEK is enabled for a cluster, no new regions can be added to the cluster resource, since they need encryption key
     * info stored in the CMEK resource. New regions can be added and maintained here instead.
     */
    additionalRegions?: pulumi.Input<pulumi.Input<inputs.CmekAdditionalRegion>[]>;
    /**
     * Cluster ID.
     */
    clusterId: pulumi.Input<string>;
    regions: pulumi.Input<pulumi.Input<inputs.CmekRegion>[]>;
    /**
     * Aggregated status of the cluster's encryption key(s).
     */
    status?: pulumi.Input<string>;
}
