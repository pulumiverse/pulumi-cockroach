// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Cockroach.Inputs
{

    public sealed class LogExportConfigGroupGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("channels", required: true)]
        private InputList<string>? _channels;

        /// <summary>
        /// A list of CockroachDB log channels to include in this group.
        /// </summary>
        public InputList<string> Channels
        {
            get => _channels ?? (_channels = new InputList<string>());
            set => _channels = value;
        }

        /// <summary>
        /// The name of the group, reflected in the log sink.
        /// </summary>
        [Input("logName", required: true)]
        public Input<string> LogName { get; set; } = null!;

        /// <summary>
        /// The minimum log level to filter to this log group.
        /// </summary>
        [Input("minLevel")]
        public Input<string>? MinLevel { get; set; }

        /// <summary>
        /// Governs whether this log group should aggregate redacted logs if unset.
        /// </summary>
        [Input("redact")]
        public Input<bool>? Redact { get; set; }

        public LogExportConfigGroupGetArgs()
        {
        }
        public static new LogExportConfigGroupGetArgs Empty => new LogExportConfigGroupGetArgs();
    }
}