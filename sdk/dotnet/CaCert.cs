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
    /// Manages client CA certs.
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
    ///     // The X509 certificate in PEM format.
    ///     var clientCertificate = config.Require("clientCertificate");
    ///     var prod = new Cockroach.CaCert("prod", new()
    ///     {
    ///         ClusterId = prodCockroachCluster.Id,
    ///         X509PemCert = clientCertificate,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [CockroachResourceType("cockroach:index/caCert:CaCert")]
    public partial class CaCert : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Status of client CA certs on a cluster.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// X509 certificate in PEM format.
        /// </summary>
        [Output("x509PemCert")]
        public Output<string> X509PemCert { get; private set; } = null!;


        /// <summary>
        /// Create a CaCert resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CaCert(string name, CaCertArgs args, CustomResourceOptions? options = null)
            : base("cockroach:index/caCert:CaCert", name, args ?? new CaCertArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CaCert(string name, Input<string> id, CaCertState? state = null, CustomResourceOptions? options = null)
            : base("cockroach:index/caCert:CaCert", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CaCert resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CaCert Get(string name, Input<string> id, CaCertState? state = null, CustomResourceOptions? options = null)
        {
            return new CaCert(name, id, state, options);
        }
    }

    public sealed class CaCertArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// X509 certificate in PEM format.
        /// </summary>
        [Input("x509PemCert", required: true)]
        public Input<string> X509PemCert { get; set; } = null!;

        public CaCertArgs()
        {
        }
        public static new CaCertArgs Empty => new CaCertArgs();
    }

    public sealed class CaCertState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Status of client CA certs on a cluster.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// X509 certificate in PEM format.
        /// </summary>
        [Input("x509PemCert")]
        public Input<string>? X509PemCert { get; set; }

        public CaCertState()
        {
        }
        public static new CaCertState Empty => new CaCertState();
    }
}
