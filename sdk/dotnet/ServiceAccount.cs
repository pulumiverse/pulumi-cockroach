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
    /// CockroachDB Cloud service account. A service account represents a non-person user. By default a service account has no access but it can be accompanied by either a cockroach.UserRoleGrants resource or any number of cockroach.UserRoleGrant resources to grant it roles.
    /// 
    /// ## Import
    /// 
    /// format &lt;resource&gt; &lt;service account id&gt;
    /// 
    /// ```sh
    /// $ pulumi import cockroach:index/serviceAccount:ServiceAccount api_service_account 1f69fdd2-600a-4cfc-a9ba-16995df0d77d
    /// ```
    /// </summary>
    [CockroachResourceType("cockroach:index/serviceAccount:ServiceAccount")]
    public partial class ServiceAccount : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Creation time of the service account.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Name of the creator of the service account.
        /// </summary>
        [Output("creatorName")]
        public Output<string> CreatorName { get; private set; } = null!;

        /// <summary>
        /// Description of the service account.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Name of the service account.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceAccount(string name, ServiceAccountArgs? args = null, CustomResourceOptions? options = null)
            : base("cockroach:index/serviceAccount:ServiceAccount", name, args ?? new ServiceAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceAccount(string name, Input<string> id, ServiceAccountState? state = null, CustomResourceOptions? options = null)
            : base("cockroach:index/serviceAccount:ServiceAccount", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceAccount Get(string name, Input<string> id, ServiceAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceAccount(name, id, state, options);
        }
    }

    public sealed class ServiceAccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the service account.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the service account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ServiceAccountArgs()
        {
        }
        public static new ServiceAccountArgs Empty => new ServiceAccountArgs();
    }

    public sealed class ServiceAccountState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creation time of the service account.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Name of the creator of the service account.
        /// </summary>
        [Input("creatorName")]
        public Input<string>? CreatorName { get; set; }

        /// <summary>
        /// Description of the service account.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the service account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ServiceAccountState()
        {
        }
        public static new ServiceAccountState Empty => new ServiceAccountState();
    }
}
