// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Cockroach
{
    /// <summary>
    /// Metric Export CloudWatch Config Resource
    /// </summary>
    [CockroachResourceType("cockroach:index/metricExportCloudwatchConfig:MetricExportCloudwatchConfig")]
    public partial class MetricExportCloudwatchConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The customized AWS CloudWatch log group name.
        /// </summary>
        [Output("logGroupName")]
        public Output<string> LogGroupName { get; private set; } = null!;

        /// <summary>
        /// The IAM role used to upload metric segments to the target AWS account.
        /// </summary>
        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The specific AWS region that the metrics will be exported to.
        /// </summary>
        [Output("targetRegion")]
        public Output<string> TargetRegion { get; private set; } = null!;

        [Output("userMessage")]
        public Output<string> UserMessage { get; private set; } = null!;


        /// <summary>
        /// Create a MetricExportCloudwatchConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MetricExportCloudwatchConfig(string name, MetricExportCloudwatchConfigArgs args, CustomResourceOptions? options = null)
            : base("cockroach:index/metricExportCloudwatchConfig:MetricExportCloudwatchConfig", name, args ?? new MetricExportCloudwatchConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MetricExportCloudwatchConfig(string name, Input<string> id, MetricExportCloudwatchConfigState? state = null, CustomResourceOptions? options = null)
            : base("cockroach:index/metricExportCloudwatchConfig:MetricExportCloudwatchConfig", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/lbrlabs",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MetricExportCloudwatchConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MetricExportCloudwatchConfig Get(string name, Input<string> id, MetricExportCloudwatchConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new MetricExportCloudwatchConfig(name, id, state, options);
        }
    }

    public sealed class MetricExportCloudwatchConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The customized AWS CloudWatch log group name.
        /// </summary>
        [Input("logGroupName")]
        public Input<string>? LogGroupName { get; set; }

        /// <summary>
        /// The IAM role used to upload metric segments to the target AWS account.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        /// The specific AWS region that the metrics will be exported to.
        /// </summary>
        [Input("targetRegion")]
        public Input<string>? TargetRegion { get; set; }

        public MetricExportCloudwatchConfigArgs()
        {
        }
        public static new MetricExportCloudwatchConfigArgs Empty => new MetricExportCloudwatchConfigArgs();
    }

    public sealed class MetricExportCloudwatchConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The customized AWS CloudWatch log group name.
        /// </summary>
        [Input("logGroupName")]
        public Input<string>? LogGroupName { get; set; }

        /// <summary>
        /// The IAM role used to upload metric segments to the target AWS account.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The specific AWS region that the metrics will be exported to.
        /// </summary>
        [Input("targetRegion")]
        public Input<string>? TargetRegion { get; set; }

        [Input("userMessage")]
        public Input<string>? UserMessage { get; set; }

        public MetricExportCloudwatchConfigState()
        {
        }
        public static new MetricExportCloudwatchConfigState Empty => new MetricExportCloudwatchConfigState();
    }
}