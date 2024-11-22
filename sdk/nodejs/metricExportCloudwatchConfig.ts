// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Amazon CloudWatch metric export configuration for a cluster.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cockroach from "@pulumiverse/cockroach";
 *
 * const config = new pulumi.Config();
 * const clusterId = config.require("clusterId");
 * const roleArn = config.require("roleArn");
 * const logGroupName = config.require("logGroupName");
 * const awsRegion = config.require("awsRegion");
 * const example = new cockroach.MetricExportCloudwatchConfig("example", {
 *     clusterId: clusterId,
 *     roleArn: roleArn,
 *     logGroupName: logGroupName,
 *     targetRegion: awsRegion,
 * });
 * ```
 *
 * ## Import
 *
 * format: <cluster id>
 *
 * ```sh
 * $ pulumi import cockroach:index/metricExportCloudwatchConfig:MetricExportCloudwatchConfig example 1f69fdd2-600a-4cfc-a9ba-16995df0d77d
 * ```
 */
export class MetricExportCloudwatchConfig extends pulumi.CustomResource {
    /**
     * Get an existing MetricExportCloudwatchConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MetricExportCloudwatchConfigState, opts?: pulumi.CustomResourceOptions): MetricExportCloudwatchConfig {
        return new MetricExportCloudwatchConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cockroach:index/metricExportCloudwatchConfig:MetricExportCloudwatchConfig';

    /**
     * Returns true if the given object is an instance of MetricExportCloudwatchConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MetricExportCloudwatchConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MetricExportCloudwatchConfig.__pulumiType;
    }

    /**
     * Cluster ID.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * The customized AWS CloudWatch log group name.
     */
    public readonly logGroupName!: pulumi.Output<string>;
    /**
     * The IAM role used to upload metric segments to the target AWS account.
     */
    public readonly roleArn!: pulumi.Output<string>;
    /**
     * Encodes the possible states that a metric export configuration can be in as it is created, deployed, and disabled.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The specific AWS region that the metrics will be exported to.
     */
    public readonly targetRegion!: pulumi.Output<string>;
    /**
     * Elaborates on the metric export status and hints at how to fix issues that may have occurred during asynchronous operations.
     */
    public /*out*/ readonly userMessage!: pulumi.Output<string>;

    /**
     * Create a MetricExportCloudwatchConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MetricExportCloudwatchConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MetricExportCloudwatchConfigArgs | MetricExportCloudwatchConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MetricExportCloudwatchConfigState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["logGroupName"] = state ? state.logGroupName : undefined;
            resourceInputs["roleArn"] = state ? state.roleArn : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["targetRegion"] = state ? state.targetRegion : undefined;
            resourceInputs["userMessage"] = state ? state.userMessage : undefined;
        } else {
            const args = argsOrState as MetricExportCloudwatchConfigArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["logGroupName"] = args ? args.logGroupName : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["targetRegion"] = args ? args.targetRegion : undefined;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["userMessage"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MetricExportCloudwatchConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MetricExportCloudwatchConfig resources.
 */
export interface MetricExportCloudwatchConfigState {
    /**
     * Cluster ID.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * The customized AWS CloudWatch log group name.
     */
    logGroupName?: pulumi.Input<string>;
    /**
     * The IAM role used to upload metric segments to the target AWS account.
     */
    roleArn?: pulumi.Input<string>;
    /**
     * Encodes the possible states that a metric export configuration can be in as it is created, deployed, and disabled.
     */
    status?: pulumi.Input<string>;
    /**
     * The specific AWS region that the metrics will be exported to.
     */
    targetRegion?: pulumi.Input<string>;
    /**
     * Elaborates on the metric export status and hints at how to fix issues that may have occurred during asynchronous operations.
     */
    userMessage?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MetricExportCloudwatchConfig resource.
 */
export interface MetricExportCloudwatchConfigArgs {
    /**
     * Cluster ID.
     */
    clusterId: pulumi.Input<string>;
    /**
     * The customized AWS CloudWatch log group name.
     */
    logGroupName?: pulumi.Input<string>;
    /**
     * The IAM role used to upload metric segments to the target AWS account.
     */
    roleArn: pulumi.Input<string>;
    /**
     * The specific AWS region that the metrics will be exported to.
     */
    targetRegion?: pulumi.Input<string>;
}
