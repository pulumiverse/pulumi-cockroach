// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Private Endpoint Trusted Owner.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cockroach from "@pulumiverse/cockroach";
 *
 * const config = new pulumi.Config();
 * const clusterId = config.require("clusterId");
 * const example = new cockroach.PrivateEndpointTrustedOwner("example", {
 *     clusterId: clusterId,
 *     type: "AWS_ACCOUNT_ID",
 *     externalOwnerId: "012345678901",
 * });
 * ```
 *
 * ## Import
 *
 * format: <cluster id>:<owner id>
 *
 * ```sh
 * $ pulumi import cockroach:index/privateEndpointTrustedOwner:PrivateEndpointTrustedOwner resource_name 1f69fdd2-600a-4cfc-a9ba-16995df0d77d:e50aa10d-1a16-4be8-85e6-4c18221daa49
 * ```
 */
export class PrivateEndpointTrustedOwner extends pulumi.CustomResource {
    /**
     * Get an existing PrivateEndpointTrustedOwner resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PrivateEndpointTrustedOwnerState, opts?: pulumi.CustomResourceOptions): PrivateEndpointTrustedOwner {
        return new PrivateEndpointTrustedOwner(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cockroach:index/privateEndpointTrustedOwner:PrivateEndpointTrustedOwner';

    /**
     * Returns true if the given object is an instance of PrivateEndpointTrustedOwner.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivateEndpointTrustedOwner {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivateEndpointTrustedOwner.__pulumiType;
    }

    /**
     * UUID of the cluster the private endpoint trusted owner entry belongs to.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Owner ID of the private endpoint connection in the cloud provider.
     */
    public readonly externalOwnerId!: pulumi.Output<string>;
    /**
     * UUID of the private endpoint trusted owner entry.
     */
    public /*out*/ readonly ownerId!: pulumi.Output<string>;
    /**
     * Representation of the externalOwnerId field. Allowed values are: * AWS_ACCOUNT_ID
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a PrivateEndpointTrustedOwner resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivateEndpointTrustedOwnerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PrivateEndpointTrustedOwnerArgs | PrivateEndpointTrustedOwnerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PrivateEndpointTrustedOwnerState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["externalOwnerId"] = state ? state.externalOwnerId : undefined;
            resourceInputs["ownerId"] = state ? state.ownerId : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as PrivateEndpointTrustedOwnerArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.externalOwnerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'externalOwnerId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["externalOwnerId"] = args ? args.externalOwnerId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["ownerId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PrivateEndpointTrustedOwner.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PrivateEndpointTrustedOwner resources.
 */
export interface PrivateEndpointTrustedOwnerState {
    /**
     * UUID of the cluster the private endpoint trusted owner entry belongs to.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Owner ID of the private endpoint connection in the cloud provider.
     */
    externalOwnerId?: pulumi.Input<string>;
    /**
     * UUID of the private endpoint trusted owner entry.
     */
    ownerId?: pulumi.Input<string>;
    /**
     * Representation of the externalOwnerId field. Allowed values are: * AWS_ACCOUNT_ID
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PrivateEndpointTrustedOwner resource.
 */
export interface PrivateEndpointTrustedOwnerArgs {
    /**
     * UUID of the cluster the private endpoint trusted owner entry belongs to.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Owner ID of the private endpoint connection in the cloud provider.
     */
    externalOwnerId: pulumi.Input<string>;
    /**
     * Representation of the externalOwnerId field. Allowed values are: * AWS_ACCOUNT_ID
     */
    type: pulumi.Input<string>;
}
