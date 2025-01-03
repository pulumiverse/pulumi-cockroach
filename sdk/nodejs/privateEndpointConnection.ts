// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Private endpoint connections allow customer applications to connect to a CockroachDB Cloud cluster without traversing the public internet. All application-database traffic remains within the cloud-provider network.
 *
 * ## Import
 *
 * format: <cluster id>:<connection id>
 *
 * ```sh
 * $ pulumi import cockroach:index/privateEndpointConnection:PrivateEndpointConnection resource_name 1f69fdd2-600a-4cfc-a9ba-16995df0d77d:vpce-0c1308d7312217abc
 * ```
 */
export class PrivateEndpointConnection extends pulumi.CustomResource {
    /**
     * Get an existing PrivateEndpointConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PrivateEndpointConnectionState, opts?: pulumi.CustomResourceOptions): PrivateEndpointConnection {
        return new PrivateEndpointConnection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cockroach:index/privateEndpointConnection:PrivateEndpointConnection';

    /**
     * Returns true if the given object is an instance of PrivateEndpointConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivateEndpointConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivateEndpointConnection.__pulumiType;
    }

    /**
     * Cloud provider associated with this connection.
     */
    public /*out*/ readonly cloudProvider!: pulumi.Output<string>;
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Client side ID of the Private Endpoint Connection.
     */
    public readonly endpointId!: pulumi.Output<string>;
    /**
     * Cloud provider region code associated with this connection.
     */
    public /*out*/ readonly regionName!: pulumi.Output<string>;
    /**
     * Server side ID of the Private Endpoint Connection.
     */
    public /*out*/ readonly serviceId!: pulumi.Output<string>;

    /**
     * Create a PrivateEndpointConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivateEndpointConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PrivateEndpointConnectionArgs | PrivateEndpointConnectionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PrivateEndpointConnectionState | undefined;
            resourceInputs["cloudProvider"] = state ? state.cloudProvider : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["endpointId"] = state ? state.endpointId : undefined;
            resourceInputs["regionName"] = state ? state.regionName : undefined;
            resourceInputs["serviceId"] = state ? state.serviceId : undefined;
        } else {
            const args = argsOrState as PrivateEndpointConnectionArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.endpointId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpointId'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["endpointId"] = args ? args.endpointId : undefined;
            resourceInputs["cloudProvider"] = undefined /*out*/;
            resourceInputs["regionName"] = undefined /*out*/;
            resourceInputs["serviceId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PrivateEndpointConnection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PrivateEndpointConnection resources.
 */
export interface PrivateEndpointConnectionState {
    /**
     * Cloud provider associated with this connection.
     */
    cloudProvider?: pulumi.Input<string>;
    clusterId?: pulumi.Input<string>;
    /**
     * Client side ID of the Private Endpoint Connection.
     */
    endpointId?: pulumi.Input<string>;
    /**
     * Cloud provider region code associated with this connection.
     */
    regionName?: pulumi.Input<string>;
    /**
     * Server side ID of the Private Endpoint Connection.
     */
    serviceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PrivateEndpointConnection resource.
 */
export interface PrivateEndpointConnectionArgs {
    clusterId: pulumi.Input<string>;
    /**
     * Client side ID of the Private Endpoint Connection.
     */
    endpointId: pulumi.Input<string>;
}
