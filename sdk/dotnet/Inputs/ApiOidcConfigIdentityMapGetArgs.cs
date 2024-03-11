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

    public sealed class ApiOidcConfigIdentityMapGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The username (email or service account id) of the CC user that the token should map to.
        /// </summary>
        [Input("ccIdentity", required: true)]
        public Input<string> CcIdentity { get; set; } = null!;

        /// <summary>
        /// Indicates that the token_principal field is a regex value.
        /// </summary>
        [Input("isRegex")]
        public Input<bool>? IsRegex { get; set; }

        /// <summary>
        /// The token value that needs to be mapped.
        /// </summary>
        [Input("tokenIdentity", required: true)]
        public Input<string> TokenIdentity { get; set; } = null!;

        public ApiOidcConfigIdentityMapGetArgs()
        {
        }
        public static new ApiOidcConfigIdentityMapGetArgs Empty => new ApiOidcConfigIdentityMapGetArgs();
    }
}