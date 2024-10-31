// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Cockroach
{
    /// <summary>
    /// Prometheus metric export configuration for a cluster. This is only available for dedicated clusters with AWS and GCP cloud providers.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Cockroach = Pulumiverse.Cockroach;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var clusterId = config.Require("clusterId");
    ///     var example = new Cockroach.MetricExportPrometheusConfig("example", new()
    ///     {
    ///         ClusterId = clusterId,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [CockroachResourceType("cockroach:index/metricExportPrometheusConfig:MetricExportPrometheusConfig")]
    public partial class MetricExportPrometheusConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// The current state of the metric export configuration.  Possible values are: [`NOT_DEPLOYED` `DISABLING` `ENABLING` `ENABLED` `ERROR`]
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        [Output("targets")]
        public Output<ImmutableDictionary<string, string>> Targets { get; private set; } = null!;


        /// <summary>
        /// Create a MetricExportPrometheusConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MetricExportPrometheusConfig(string name, MetricExportPrometheusConfigArgs args, CustomResourceOptions? options = null)
            : base("cockroach:index/metricExportPrometheusConfig:MetricExportPrometheusConfig", name, args ?? new MetricExportPrometheusConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MetricExportPrometheusConfig(string name, Input<string> id, MetricExportPrometheusConfigState? state = null, CustomResourceOptions? options = null)
            : base("cockroach:index/metricExportPrometheusConfig:MetricExportPrometheusConfig", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MetricExportPrometheusConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MetricExportPrometheusConfig Get(string name, Input<string> id, MetricExportPrometheusConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new MetricExportPrometheusConfig(name, id, state, options);
        }
    }

    public sealed class MetricExportPrometheusConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        public MetricExportPrometheusConfigArgs()
        {
        }
        public static new MetricExportPrometheusConfigArgs Empty => new MetricExportPrometheusConfigArgs();
    }

    public sealed class MetricExportPrometheusConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// The current state of the metric export configuration.  Possible values are: [`NOT_DEPLOYED` `DISABLING` `ENABLING` `ENABLED` `ERROR`]
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("targets")]
        private InputMap<string>? _targets;
        public InputMap<string> Targets
        {
            get => _targets ?? (_targets = new InputMap<string>());
            set => _targets = value;
        }

        public MetricExportPrometheusConfigState()
        {
        }
        public static new MetricExportPrometheusConfigState Empty => new MetricExportPrometheusConfigState();
    }
}
