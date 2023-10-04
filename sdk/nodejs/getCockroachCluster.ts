// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * CockroachDB Cloud cluster. Can be Dedicated or Serverless.
 */
export function getCockroachCluster(args: GetCockroachClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetCockroachClusterResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("cockroach:index/getCockroachCluster:getCockroachCluster", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getCockroachCluster.
 */
export interface GetCockroachClusterArgs {
    id: string;
}

/**
 * A collection of values returned by getCockroachCluster.
 */
export interface GetCockroachClusterResult {
    readonly accountId: string;
    readonly cloudProvider: string;
    readonly cockroachVersion: string;
    readonly creatorId: string;
    readonly dedicated: outputs.GetCockroachClusterDedicated;
    readonly id: string;
    readonly name: string;
    readonly operationStatus: string;
    readonly parentId: string;
    readonly plan: string;
    readonly regions: outputs.GetCockroachClusterRegion[];
    readonly serverless: outputs.GetCockroachClusterServerless;
    readonly state: string;
    readonly upgradeStatus: string;
}
/**
 * CockroachDB Cloud cluster. Can be Dedicated or Serverless.
 */
export function getCockroachClusterOutput(args: GetCockroachClusterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCockroachClusterResult> {
    return pulumi.output(args).apply((a: any) => getCockroachCluster(a, opts))
}

/**
 * A collection of arguments for invoking getCockroachCluster.
 */
export interface GetCockroachClusterOutputArgs {
    id: pulumi.Input<string>;
}
