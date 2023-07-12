// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cockroach;

import com.pulumi.cockroach.MetricExportCloudwatchConfigArgs;
import com.pulumi.cockroach.Utilities;
import com.pulumi.cockroach.inputs.MetricExportCloudwatchConfigState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Metric Export CloudWatch Config Resource
 * 
 */
@ResourceType(type="cockroach:index/metricExportCloudwatchConfig:MetricExportCloudwatchConfig")
public class MetricExportCloudwatchConfig extends com.pulumi.resources.CustomResource {
    /**
     * Cluster ID
     * 
     */
    @Export(name="clusterId", type=String.class, parameters={})
    private Output<String> clusterId;

    /**
     * @return Cluster ID
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }
    /**
     * The customized AWS CloudWatch log group name.
     * 
     */
    @Export(name="logGroupName", type=String.class, parameters={})
    private Output<String> logGroupName;

    /**
     * @return The customized AWS CloudWatch log group name.
     * 
     */
    public Output<String> logGroupName() {
        return this.logGroupName;
    }
    /**
     * The IAM role used to upload metric segments to the target AWS account.
     * 
     */
    @Export(name="roleArn", type=String.class, parameters={})
    private Output<String> roleArn;

    /**
     * @return The IAM role used to upload metric segments to the target AWS account.
     * 
     */
    public Output<String> roleArn() {
        return this.roleArn;
    }
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    public Output<String> status() {
        return this.status;
    }
    /**
     * The specific AWS region that the metrics will be exported to.
     * 
     */
    @Export(name="targetRegion", type=String.class, parameters={})
    private Output<String> targetRegion;

    /**
     * @return The specific AWS region that the metrics will be exported to.
     * 
     */
    public Output<String> targetRegion() {
        return this.targetRegion;
    }
    @Export(name="userMessage", type=String.class, parameters={})
    private Output<String> userMessage;

    public Output<String> userMessage() {
        return this.userMessage;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MetricExportCloudwatchConfig(String name) {
        this(name, MetricExportCloudwatchConfigArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MetricExportCloudwatchConfig(String name, MetricExportCloudwatchConfigArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MetricExportCloudwatchConfig(String name, MetricExportCloudwatchConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cockroach:index/metricExportCloudwatchConfig:MetricExportCloudwatchConfig", name, args == null ? MetricExportCloudwatchConfigArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MetricExportCloudwatchConfig(String name, Output<String> id, @Nullable MetricExportCloudwatchConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("cockroach:index/metricExportCloudwatchConfig:MetricExportCloudwatchConfig", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static MetricExportCloudwatchConfig get(String name, Output<String> id, @Nullable MetricExportCloudwatchConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MetricExportCloudwatchConfig(name, id, state, options);
    }
}
