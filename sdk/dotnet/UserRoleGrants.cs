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
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// format: &lt;user id&gt;
    /// 
    /// ```sh
    /// $ pulumi import cockroach:index/userRoleGrants:UserRoleGrants service_account 1f69fdd2-600a-4cfc-a9ba-16995df0d77d
    /// ```
    /// </summary>
    [CockroachResourceType("cockroach:index/userRoleGrants:UserRoleGrants")]
    public partial class UserRoleGrants : global::Pulumi.CustomResource
    {
        [Output("roles")]
        public Output<ImmutableArray<Outputs.UserRoleGrantsRole>> Roles { get; private set; } = null!;

        /// <summary>
        /// ID of the user to grant these roles to.
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a UserRoleGrants resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserRoleGrants(string name, UserRoleGrantsArgs args, CustomResourceOptions? options = null)
            : base("cockroach:index/userRoleGrants:UserRoleGrants", name, args ?? new UserRoleGrantsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserRoleGrants(string name, Input<string> id, UserRoleGrantsState? state = null, CustomResourceOptions? options = null)
            : base("cockroach:index/userRoleGrants:UserRoleGrants", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserRoleGrants resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserRoleGrants Get(string name, Input<string> id, UserRoleGrantsState? state = null, CustomResourceOptions? options = null)
        {
            return new UserRoleGrants(name, id, state, options);
        }
    }

    public sealed class UserRoleGrantsArgs : global::Pulumi.ResourceArgs
    {
        [Input("roles", required: true)]
        private InputList<Inputs.UserRoleGrantsRoleArgs>? _roles;
        public InputList<Inputs.UserRoleGrantsRoleArgs> Roles
        {
            get => _roles ?? (_roles = new InputList<Inputs.UserRoleGrantsRoleArgs>());
            set => _roles = value;
        }

        /// <summary>
        /// ID of the user to grant these roles to.
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public UserRoleGrantsArgs()
        {
        }
        public static new UserRoleGrantsArgs Empty => new UserRoleGrantsArgs();
    }

    public sealed class UserRoleGrantsState : global::Pulumi.ResourceArgs
    {
        [Input("roles")]
        private InputList<Inputs.UserRoleGrantsRoleGetArgs>? _roles;
        public InputList<Inputs.UserRoleGrantsRoleGetArgs> Roles
        {
            get => _roles ?? (_roles = new InputList<Inputs.UserRoleGrantsRoleGetArgs>());
            set => _roles = value;
        }

        /// <summary>
        /// ID of the user to grant these roles to.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public UserRoleGrantsState()
        {
        }
        public static new UserRoleGrantsState Empty => new UserRoleGrantsState();
    }
}
