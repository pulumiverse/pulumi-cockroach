// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Cockroach.Inputs
{

    public sealed class PrivateEndpointServicesServiceArgs : global::Pulumi.ResourceArgs
    {
        [Input("aws")]
        public Input<Inputs.PrivateEndpointServicesServiceAwsArgs>? Aws { get; set; }

        /// <summary>
        /// Cloud provider associated with this service.
        /// </summary>
        [Input("cloudProvider")]
        public Input<string>? CloudProvider { get; set; }

        /// <summary>
        /// Cloud provider region code associated with this service.
        /// </summary>
        [Input("regionName")]
        public Input<string>? RegionName { get; set; }

        /// <summary>
        /// Operation status of the service.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public PrivateEndpointServicesServiceArgs()
        {
        }
        public static new PrivateEndpointServicesServiceArgs Empty => new PrivateEndpointServicesServiceArgs();
    }
}
