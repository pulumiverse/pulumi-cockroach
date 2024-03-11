// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Cockroach.Outputs
{

    [OutputType]
    public sealed class CmekRegionKey
    {
        /// <summary>
        /// Principal to authenticate as in order to access the key.
        /// </summary>
        public readonly string AuthPrincipal;
        /// <summary>
        /// Indicates when the key was created.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// Current status of this key.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// Type of encryption key. Current allowed values are:
        ///   * AWS_KMS
        ///   * GCP_CLOUD_KMS
        ///   * NULL_KMS
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Indicates when the key was last updated.
        /// </summary>
        public readonly string? UpdatedAt;
        /// <summary>
        /// Provider-specific URI pointing to the encryption key.
        /// </summary>
        public readonly string Uri;
        /// <summary>
        /// Elaborates on the key's status and hints at how to fix issues that may have occurred during asynchronous key operations.
        /// </summary>
        public readonly string? UserMessage;

        [OutputConstructor]
        private CmekRegionKey(
            string authPrincipal,

            string? createdAt,

            string? status,

            string type,

            string? updatedAt,

            string uri,

            string? userMessage)
        {
            AuthPrincipal = authPrincipal;
            CreatedAt = createdAt;
            Status = status;
            Type = type;
            UpdatedAt = updatedAt;
            Uri = uri;
            UserMessage = userMessage;
        }
    }
}
