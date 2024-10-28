// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * CockroachDB Cloud folder.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cockroach from "@pulumiverse/cockroach";
 *
 * const aTeam = new cockroach.Folder("a_team", {
 *     name: "a-team",
 *     parentId: "root",
 * });
 * const aTeamDev = new cockroach.Folder("a_team_dev", {
 *     name: "dev",
 *     parentId: aTeam.id,
 * });
 * ```
 */
export class Folder extends pulumi.CustomResource {
    /**
     * Get an existing Folder resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FolderState, opts?: pulumi.CustomResourceOptions): Folder {
        return new Folder(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cockroach:index/folder:Folder';

    /**
     * Returns true if the given object is an instance of Folder.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Folder {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Folder.__pulumiType;
    }

    /**
     * Name of the folder.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the parent folder. Use 'root' for the root level (no parent folder).
     */
    public readonly parentId!: pulumi.Output<string>;

    /**
     * Create a Folder resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FolderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FolderArgs | FolderState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FolderState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parentId"] = state ? state.parentId : undefined;
        } else {
            const args = argsOrState as FolderArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.parentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'parentId'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parentId"] = args ? args.parentId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Folder.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Folder resources.
 */
export interface FolderState {
    /**
     * Name of the folder.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the parent folder. Use 'root' for the root level (no parent folder).
     */
    parentId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Folder resource.
 */
export interface FolderArgs {
    /**
     * Name of the folder.
     */
    name: pulumi.Input<string>;
    /**
     * ID of the parent folder. Use 'root' for the root level (no parent folder).
     */
    parentId: pulumi.Input<string>;
}
