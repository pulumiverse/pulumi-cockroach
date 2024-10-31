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
    /// Configuration to allow external OIDC providers to issue tokens for use with CC API.
    /// </summary>
    [CockroachResourceType("cockroach:index/apiOidcConfig:ApiOidcConfig")]
    public partial class ApiOidcConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The audience that CC API should accept for this API OIDC Configuration.
        /// </summary>
        [Output("audience")]
        public Output<string> Audience { get; private set; } = null!;

        /// <summary>
        /// The JWT claim that should be used as the user identifier. Defaults to the subject.
        /// </summary>
        [Output("claim")]
        public Output<string> Claim { get; private set; } = null!;

        /// <summary>
        /// The mapping rules to convert token user identifiers into a new form.
        /// </summary>
        [Output("identityMaps")]
        public Output<ImmutableArray<Outputs.ApiOidcConfigIdentityMap>> IdentityMaps { get; private set; } = null!;

        /// <summary>
        /// The issuer of tokens for the API OIDC Configuration. Usually this is a url.
        /// </summary>
        [Output("issuer")]
        public Output<string> Issuer { get; private set; } = null!;

        /// <summary>
        /// The JSON Web Key Set used to check the signature of the JWTs.
        /// </summary>
        [Output("jwks")]
        public Output<string> Jwks { get; private set; } = null!;


        /// <summary>
        /// Create a ApiOidcConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApiOidcConfig(string name, ApiOidcConfigArgs args, CustomResourceOptions? options = null)
            : base("cockroach:index/apiOidcConfig:ApiOidcConfig", name, args ?? new ApiOidcConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApiOidcConfig(string name, Input<string> id, ApiOidcConfigState? state = null, CustomResourceOptions? options = null)
            : base("cockroach:index/apiOidcConfig:ApiOidcConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApiOidcConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApiOidcConfig Get(string name, Input<string> id, ApiOidcConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new ApiOidcConfig(name, id, state, options);
        }
    }

    public sealed class ApiOidcConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The audience that CC API should accept for this API OIDC Configuration.
        /// </summary>
        [Input("audience", required: true)]
        public Input<string> Audience { get; set; } = null!;

        /// <summary>
        /// The JWT claim that should be used as the user identifier. Defaults to the subject.
        /// </summary>
        [Input("claim")]
        public Input<string>? Claim { get; set; }

        [Input("identityMaps")]
        private InputList<Inputs.ApiOidcConfigIdentityMapArgs>? _identityMaps;

        /// <summary>
        /// The mapping rules to convert token user identifiers into a new form.
        /// </summary>
        public InputList<Inputs.ApiOidcConfigIdentityMapArgs> IdentityMaps
        {
            get => _identityMaps ?? (_identityMaps = new InputList<Inputs.ApiOidcConfigIdentityMapArgs>());
            set => _identityMaps = value;
        }

        /// <summary>
        /// The issuer of tokens for the API OIDC Configuration. Usually this is a url.
        /// </summary>
        [Input("issuer", required: true)]
        public Input<string> Issuer { get; set; } = null!;

        /// <summary>
        /// The JSON Web Key Set used to check the signature of the JWTs.
        /// </summary>
        [Input("jwks", required: true)]
        public Input<string> Jwks { get; set; } = null!;

        public ApiOidcConfigArgs()
        {
        }
        public static new ApiOidcConfigArgs Empty => new ApiOidcConfigArgs();
    }

    public sealed class ApiOidcConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The audience that CC API should accept for this API OIDC Configuration.
        /// </summary>
        [Input("audience")]
        public Input<string>? Audience { get; set; }

        /// <summary>
        /// The JWT claim that should be used as the user identifier. Defaults to the subject.
        /// </summary>
        [Input("claim")]
        public Input<string>? Claim { get; set; }

        [Input("identityMaps")]
        private InputList<Inputs.ApiOidcConfigIdentityMapGetArgs>? _identityMaps;

        /// <summary>
        /// The mapping rules to convert token user identifiers into a new form.
        /// </summary>
        public InputList<Inputs.ApiOidcConfigIdentityMapGetArgs> IdentityMaps
        {
            get => _identityMaps ?? (_identityMaps = new InputList<Inputs.ApiOidcConfigIdentityMapGetArgs>());
            set => _identityMaps = value;
        }

        /// <summary>
        /// The issuer of tokens for the API OIDC Configuration. Usually this is a url.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        /// <summary>
        /// The JSON Web Key Set used to check the signature of the JWTs.
        /// </summary>
        [Input("jwks")]
        public Input<string>? Jwks { get; set; }

        public ApiOidcConfigState()
        {
        }
        public static new ApiOidcConfigState Empty => new ApiOidcConfigState();
    }
}