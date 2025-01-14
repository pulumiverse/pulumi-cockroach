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

    public sealed class CmekRegionKeyGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Principal to authenticate as in order to access the key.
        /// </summary>
        [Input("authPrincipal", required: true)]
        public Input<string> AuthPrincipal { get; set; } = null!;

        /// <summary>
        /// Indicates when the key was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Current status of this key.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Type of encryption key. Current allowed values are:
        ///   * AWS_KMS
        ///   * GCP_CLOUD_KMS
        ///   * NULL_KMS
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Indicates when the key was last updated.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// Provider-specific URI pointing to the encryption key.
        /// </summary>
        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        /// <summary>
        /// Elaborates on the key's status and hints at how to fix issues that may have occurred during asynchronous key operations.
        /// </summary>
        [Input("userMessage")]
        public Input<string>? UserMessage { get; set; }

        public CmekRegionKeyGetArgs()
        {
        }
        public static new CmekRegionKeyGetArgs Empty => new CmekRegionKeyGetArgs();
    }
}