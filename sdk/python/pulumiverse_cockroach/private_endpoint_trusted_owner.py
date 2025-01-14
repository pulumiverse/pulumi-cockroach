# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['PrivateEndpointTrustedOwnerArgs', 'PrivateEndpointTrustedOwner']

@pulumi.input_type
class PrivateEndpointTrustedOwnerArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 external_owner_id: pulumi.Input[str],
                 type: pulumi.Input[str]):
        """
        The set of arguments for constructing a PrivateEndpointTrustedOwner resource.
        :param pulumi.Input[str] cluster_id: UUID of the cluster the private endpoint trusted owner entry belongs to.
        :param pulumi.Input[str] external_owner_id: Owner ID of the private endpoint connection in the cloud provider.
        :param pulumi.Input[str] type: Representation of the external_owner_id field. Allowed values are: * AWS_ACCOUNT_ID
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "external_owner_id", external_owner_id)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        UUID of the cluster the private endpoint trusted owner entry belongs to.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="externalOwnerId")
    def external_owner_id(self) -> pulumi.Input[str]:
        """
        Owner ID of the private endpoint connection in the cloud provider.
        """
        return pulumi.get(self, "external_owner_id")

    @external_owner_id.setter
    def external_owner_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "external_owner_id", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Representation of the external_owner_id field. Allowed values are: * AWS_ACCOUNT_ID
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _PrivateEndpointTrustedOwnerState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 external_owner_id: Optional[pulumi.Input[str]] = None,
                 owner_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PrivateEndpointTrustedOwner resources.
        :param pulumi.Input[str] cluster_id: UUID of the cluster the private endpoint trusted owner entry belongs to.
        :param pulumi.Input[str] external_owner_id: Owner ID of the private endpoint connection in the cloud provider.
        :param pulumi.Input[str] owner_id: UUID of the private endpoint trusted owner entry.
        :param pulumi.Input[str] type: Representation of the external_owner_id field. Allowed values are: * AWS_ACCOUNT_ID
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if external_owner_id is not None:
            pulumi.set(__self__, "external_owner_id", external_owner_id)
        if owner_id is not None:
            pulumi.set(__self__, "owner_id", owner_id)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        UUID of the cluster the private endpoint trusted owner entry belongs to.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="externalOwnerId")
    def external_owner_id(self) -> Optional[pulumi.Input[str]]:
        """
        Owner ID of the private endpoint connection in the cloud provider.
        """
        return pulumi.get(self, "external_owner_id")

    @external_owner_id.setter
    def external_owner_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "external_owner_id", value)

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> Optional[pulumi.Input[str]]:
        """
        UUID of the private endpoint trusted owner entry.
        """
        return pulumi.get(self, "owner_id")

    @owner_id.setter
    def owner_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner_id", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Representation of the external_owner_id field. Allowed values are: * AWS_ACCOUNT_ID
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class PrivateEndpointTrustedOwner(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 external_owner_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Private Endpoint Trusted Owner.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_cockroach as cockroach

        config = pulumi.Config()
        cluster_id = config.require("clusterId")
        example = cockroach.PrivateEndpointTrustedOwner("example",
            cluster_id=cluster_id,
            type="AWS_ACCOUNT_ID",
            external_owner_id="012345678901")
        ```

        ## Import

        format: <cluster id>:<owner id>

        ```sh
        $ pulumi import cockroach:index/privateEndpointTrustedOwner:PrivateEndpointTrustedOwner resource_name 1f69fdd2-600a-4cfc-a9ba-16995df0d77d:e50aa10d-1a16-4be8-85e6-4c18221daa49
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: UUID of the cluster the private endpoint trusted owner entry belongs to.
        :param pulumi.Input[str] external_owner_id: Owner ID of the private endpoint connection in the cloud provider.
        :param pulumi.Input[str] type: Representation of the external_owner_id field. Allowed values are: * AWS_ACCOUNT_ID
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PrivateEndpointTrustedOwnerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Private Endpoint Trusted Owner.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_cockroach as cockroach

        config = pulumi.Config()
        cluster_id = config.require("clusterId")
        example = cockroach.PrivateEndpointTrustedOwner("example",
            cluster_id=cluster_id,
            type="AWS_ACCOUNT_ID",
            external_owner_id="012345678901")
        ```

        ## Import

        format: <cluster id>:<owner id>

        ```sh
        $ pulumi import cockroach:index/privateEndpointTrustedOwner:PrivateEndpointTrustedOwner resource_name 1f69fdd2-600a-4cfc-a9ba-16995df0d77d:e50aa10d-1a16-4be8-85e6-4c18221daa49
        ```

        :param str resource_name: The name of the resource.
        :param PrivateEndpointTrustedOwnerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrivateEndpointTrustedOwnerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 external_owner_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PrivateEndpointTrustedOwnerArgs.__new__(PrivateEndpointTrustedOwnerArgs)

            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            if external_owner_id is None and not opts.urn:
                raise TypeError("Missing required property 'external_owner_id'")
            __props__.__dict__["external_owner_id"] = external_owner_id
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["owner_id"] = None
        super(PrivateEndpointTrustedOwner, __self__).__init__(
            'cockroach:index/privateEndpointTrustedOwner:PrivateEndpointTrustedOwner',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            external_owner_id: Optional[pulumi.Input[str]] = None,
            owner_id: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'PrivateEndpointTrustedOwner':
        """
        Get an existing PrivateEndpointTrustedOwner resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: UUID of the cluster the private endpoint trusted owner entry belongs to.
        :param pulumi.Input[str] external_owner_id: Owner ID of the private endpoint connection in the cloud provider.
        :param pulumi.Input[str] owner_id: UUID of the private endpoint trusted owner entry.
        :param pulumi.Input[str] type: Representation of the external_owner_id field. Allowed values are: * AWS_ACCOUNT_ID
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PrivateEndpointTrustedOwnerState.__new__(_PrivateEndpointTrustedOwnerState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["external_owner_id"] = external_owner_id
        __props__.__dict__["owner_id"] = owner_id
        __props__.__dict__["type"] = type
        return PrivateEndpointTrustedOwner(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        UUID of the cluster the private endpoint trusted owner entry belongs to.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="externalOwnerId")
    def external_owner_id(self) -> pulumi.Output[str]:
        """
        Owner ID of the private endpoint connection in the cloud provider.
        """
        return pulumi.get(self, "external_owner_id")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Output[str]:
        """
        UUID of the private endpoint trusted owner entry.
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Representation of the external_owner_id field. Allowed values are: * AWS_ACCOUNT_ID
        """
        return pulumi.get(self, "type")

