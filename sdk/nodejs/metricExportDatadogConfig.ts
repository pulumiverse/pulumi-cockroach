// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * DataDog metric export configuration for a cluster.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cockroach from "@pulumiverse/cockroach";
 *
 * const config = new pulumi.Config();
 * const clusterId = config.require("clusterId");
 * const datadogSite = config.require("datadogSite");
 * const datadogApiKey = config.require("datadogApiKey");
 * const example = new cockroach.MetricExportDatadogConfig("example", {
 *     clusterId: clusterId,
 *     site: datadogSite,
 *     apiKey: datadogApiKey,
 * });
 * ```
 */
export class MetricExportDatadogConfig extends pulumi.CustomResource {
    /**
     * Get an existing MetricExportDatadogConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MetricExportDatadogConfigState, opts?: pulumi.CustomResourceOptions): MetricExportDatadogConfig {
        return new MetricExportDatadogConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cockroach:index/metricExportDatadogConfig:MetricExportDatadogConfig';

    /**
     * Returns true if the given object is an instance of MetricExportDatadogConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MetricExportDatadogConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MetricExportDatadogConfig.__pulumiType;
    }

    /**
     * A Datadog API key.
     */
    public readonly apiKey!: pulumi.Output<string>;
    /**
     * Cluster ID.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * The Datadog region to export to.
     */
    public readonly site!: pulumi.Output<string>;
    /**
     * Encodes the possible states that a metric export configuration can be in as it is created, deployed, and disabled.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Elaborates on the metric export status and hints at how to fix issues that may have occurred during asynchronous operations.
     */
    public /*out*/ readonly userMessage!: pulumi.Output<string>;

    /**
     * Create a MetricExportDatadogConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MetricExportDatadogConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MetricExportDatadogConfigArgs | MetricExportDatadogConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MetricExportDatadogConfigState | undefined;
            resourceInputs["apiKey"] = state ? state.apiKey : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["site"] = state ? state.site : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["userMessage"] = state ? state.userMessage : undefined;
        } else {
            const args = argsOrState as MetricExportDatadogConfigArgs | undefined;
            if ((!args || args.apiKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiKey'");
            }
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.site === undefined) && !opts.urn) {
                throw new Error("Missing required property 'site'");
            }
            resourceInputs["apiKey"] = args?.apiKey ? pulumi.secret(args.apiKey) : undefined;
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["site"] = args ? args.site : undefined;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["userMessage"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["apiKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(MetricExportDatadogConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MetricExportDatadogConfig resources.
 */
export interface MetricExportDatadogConfigState {
    /**
     * A Datadog API key.
     */
    apiKey?: pulumi.Input<string>;
    /**
     * Cluster ID.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * The Datadog region to export to.
     */
    site?: pulumi.Input<string>;
    /**
     * Encodes the possible states that a metric export configuration can be in as it is created, deployed, and disabled.
     */
    status?: pulumi.Input<string>;
    /**
     * Elaborates on the metric export status and hints at how to fix issues that may have occurred during asynchronous operations.
     */
    userMessage?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MetricExportDatadogConfig resource.
 */
export interface MetricExportDatadogConfigArgs {
    /**
     * A Datadog API key.
     */
    apiKey: pulumi.Input<string>;
    /**
     * Cluster ID.
     */
    clusterId: pulumi.Input<string>;
    /**
     * The Datadog region to export to.
     */
    site: pulumi.Input<string>;
}