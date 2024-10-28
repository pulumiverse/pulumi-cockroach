// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Information about the organization associated with the user's API key.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cockroach from "@pulumi/cockroach";
 *
 * const prod = cockroach.getOrganization({});
 * ```
 */
export function getOrganization(opts?: pulumi.InvokeOptions): Promise<GetOrganizationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("cockroach:index/getOrganization:getOrganization", {
    }, opts);
}

/**
 * A collection of values returned by getOrganization.
 */
export interface GetOrganizationResult {
    /**
     * Indicates when the organization was created.
     */
    readonly createdAt: string;
    /**
     * Organization ID.
     */
    readonly id: string;
    /**
     * A short ID used by CockroachDB Support.
     */
    readonly label: string;
    /**
     * Name of the organization.
     */
    readonly name: string;
}
/**
 * Information about the organization associated with the user's API key.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cockroach from "@pulumi/cockroach";
 *
 * const prod = cockroach.getOrganization({});
 * ```
 */
export function getOrganizationOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetOrganizationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("cockroach:index/getOrganization:getOrganization", {
    }, opts);
}
