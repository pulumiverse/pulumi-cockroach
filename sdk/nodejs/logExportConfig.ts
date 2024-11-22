// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Log Export configuration for a cluster.
 *
 * ## Import
 *
 * format: <cluster id>
 *
 * ```sh
 * $ pulumi import cockroach:index/logExportConfig:LogExportConfig example 1f69fdd2-600a-4cfc-a9ba-16995df0d77d
 * ```
 */
export class LogExportConfig extends pulumi.CustomResource {
    /**
     * Get an existing LogExportConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LogExportConfigState, opts?: pulumi.CustomResourceOptions): LogExportConfig {
        return new LogExportConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cockroach:index/logExportConfig:LogExportConfig';

    /**
     * Returns true if the given object is an instance of LogExportConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogExportConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogExportConfig.__pulumiType;
    }

    /**
     * Either the AWS Role ARN that identifies a role that the cluster account can assume to write to CloudWatch or the GCP
     * Project ID that the cluster service account has permissions to write to for cloud logging.
     */
    public readonly authPrincipal!: pulumi.Output<string>;
    /**
     * Cluster ID.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Indicates when log export was initially configured.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    public readonly groups!: pulumi.Output<outputs.LogExportConfigGroup[] | undefined>;
    /**
     * An identifier for the logs in the customer's log sink.
     */
    public readonly logName!: pulumi.Output<string>;
    /**
     * Controls what CRDB channels do not get exported.
     */
    public readonly omittedChannels!: pulumi.Output<string[] | undefined>;
    /**
     * Controls whether logs are redacted before forwarding to customer sinks.
     */
    public readonly redact!: pulumi.Output<boolean | undefined>;
    /**
     * Controls whether all logs are sent to a specific region in the customer sink.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Encodes the possible states that a log export configuration can be in as it is created, deployed, and disabled.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The cloud selection being exported to along with the cloud logging platform. Possible values are: * AWS_CLOUDWATCH *
     * GCP_CLOUD_LOGGING * AZURE_LOG_ANALYTICS
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Indicates when the log export configuration was last updated.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * Elaborates on the log export status and hints at how to fix issues that may have occurred during asynchronous
     * operations.
     */
    public /*out*/ readonly userMessage!: pulumi.Output<string>;

    /**
     * Create a LogExportConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LogExportConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LogExportConfigArgs | LogExportConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LogExportConfigState | undefined;
            resourceInputs["authPrincipal"] = state ? state.authPrincipal : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["groups"] = state ? state.groups : undefined;
            resourceInputs["logName"] = state ? state.logName : undefined;
            resourceInputs["omittedChannels"] = state ? state.omittedChannels : undefined;
            resourceInputs["redact"] = state ? state.redact : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["userMessage"] = state ? state.userMessage : undefined;
        } else {
            const args = argsOrState as LogExportConfigArgs | undefined;
            if ((!args || args.authPrincipal === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authPrincipal'");
            }
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.logName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'logName'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["authPrincipal"] = args ? args.authPrincipal : undefined;
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["groups"] = args ? args.groups : undefined;
            resourceInputs["logName"] = args ? args.logName : undefined;
            resourceInputs["omittedChannels"] = args ? args.omittedChannels : undefined;
            resourceInputs["redact"] = args ? args.redact : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
            resourceInputs["userMessage"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LogExportConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LogExportConfig resources.
 */
export interface LogExportConfigState {
    /**
     * Either the AWS Role ARN that identifies a role that the cluster account can assume to write to CloudWatch or the GCP
     * Project ID that the cluster service account has permissions to write to for cloud logging.
     */
    authPrincipal?: pulumi.Input<string>;
    /**
     * Cluster ID.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Indicates when log export was initially configured.
     */
    createdAt?: pulumi.Input<string>;
    groups?: pulumi.Input<pulumi.Input<inputs.LogExportConfigGroup>[]>;
    /**
     * An identifier for the logs in the customer's log sink.
     */
    logName?: pulumi.Input<string>;
    /**
     * Controls what CRDB channels do not get exported.
     */
    omittedChannels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Controls whether logs are redacted before forwarding to customer sinks.
     */
    redact?: pulumi.Input<boolean>;
    /**
     * Controls whether all logs are sent to a specific region in the customer sink.
     */
    region?: pulumi.Input<string>;
    /**
     * Encodes the possible states that a log export configuration can be in as it is created, deployed, and disabled.
     */
    status?: pulumi.Input<string>;
    /**
     * The cloud selection being exported to along with the cloud logging platform. Possible values are: * AWS_CLOUDWATCH *
     * GCP_CLOUD_LOGGING * AZURE_LOG_ANALYTICS
     */
    type?: pulumi.Input<string>;
    /**
     * Indicates when the log export configuration was last updated.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * Elaborates on the log export status and hints at how to fix issues that may have occurred during asynchronous
     * operations.
     */
    userMessage?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LogExportConfig resource.
 */
export interface LogExportConfigArgs {
    /**
     * Either the AWS Role ARN that identifies a role that the cluster account can assume to write to CloudWatch or the GCP
     * Project ID that the cluster service account has permissions to write to for cloud logging.
     */
    authPrincipal: pulumi.Input<string>;
    /**
     * Cluster ID.
     */
    clusterId: pulumi.Input<string>;
    groups?: pulumi.Input<pulumi.Input<inputs.LogExportConfigGroup>[]>;
    /**
     * An identifier for the logs in the customer's log sink.
     */
    logName: pulumi.Input<string>;
    /**
     * Controls what CRDB channels do not get exported.
     */
    omittedChannels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Controls whether logs are redacted before forwarding to customer sinks.
     */
    redact?: pulumi.Input<boolean>;
    /**
     * Controls whether all logs are sent to a specific region in the customer sink.
     */
    region?: pulumi.Input<string>;
    /**
     * The cloud selection being exported to along with the cloud logging platform. Possible values are: * AWS_CLOUDWATCH *
     * GCP_CLOUD_LOGGING * AZURE_LOG_ANALYTICS
     */
    type: pulumi.Input<string>;
}
