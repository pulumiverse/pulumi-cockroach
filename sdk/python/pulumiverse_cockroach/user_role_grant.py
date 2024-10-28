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
from . import outputs
from ._inputs import *

__all__ = ['UserRoleGrantArgs', 'UserRoleGrant']

@pulumi.input_type
class UserRoleGrantArgs:
    def __init__(__self__, *,
                 role: pulumi.Input['UserRoleGrantRoleArgs'],
                 user_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a UserRoleGrant resource.
        :param pulumi.Input[str] user_id: ID of the user to grant these roles to.
        """
        pulumi.set(__self__, "role", role)
        pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input['UserRoleGrantRoleArgs']:
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input['UserRoleGrantRoleArgs']):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Input[str]:
        """
        ID of the user to grant these roles to.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_id", value)


@pulumi.input_type
class _UserRoleGrantState:
    def __init__(__self__, *,
                 role: Optional[pulumi.Input['UserRoleGrantRoleArgs']] = None,
                 user_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering UserRoleGrant resources.
        :param pulumi.Input[str] user_id: ID of the user to grant these roles to.
        """
        if role is not None:
            pulumi.set(__self__, "role", role)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input['UserRoleGrantRoleArgs']]:
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input['UserRoleGrantRoleArgs']]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the user to grant these roles to.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)


class UserRoleGrant(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 role: Optional[pulumi.Input[Union['UserRoleGrantRoleArgs', 'UserRoleGrantRoleArgsDict']]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ## Import

        Cluster and folder level role grants can be imported using:

        <user_id>,<role_name>,<resource_type>,<resource_id>

        ```sh
        $ pulumi import cockroach:index/userRoleGrant:UserRoleGrant admin_grant 1f69fdd2-600a-4cfc-a9ba-16995df0d77d,CLUSTER_ADMIN,CLUSTER,9b9d23fe-3848-40b2-a3c5-d8ccb1c4f831
        ```

        Organization level grants can omit the resource_id

        ```sh
        $ pulumi import cockroach:index/userRoleGrant:UserRoleGrant org_level_grant 1f69fdd2-600a-4cfc-a9ba-16995df0d77d,ORG_ADMIN,ORGANIZATION
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] user_id: ID of the user to grant these roles to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserRoleGrantArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ## Import

        Cluster and folder level role grants can be imported using:

        <user_id>,<role_name>,<resource_type>,<resource_id>

        ```sh
        $ pulumi import cockroach:index/userRoleGrant:UserRoleGrant admin_grant 1f69fdd2-600a-4cfc-a9ba-16995df0d77d,CLUSTER_ADMIN,CLUSTER,9b9d23fe-3848-40b2-a3c5-d8ccb1c4f831
        ```

        Organization level grants can omit the resource_id

        ```sh
        $ pulumi import cockroach:index/userRoleGrant:UserRoleGrant org_level_grant 1f69fdd2-600a-4cfc-a9ba-16995df0d77d,ORG_ADMIN,ORGANIZATION
        ```

        :param str resource_name: The name of the resource.
        :param UserRoleGrantArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserRoleGrantArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 role: Optional[pulumi.Input[Union['UserRoleGrantRoleArgs', 'UserRoleGrantRoleArgsDict']]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserRoleGrantArgs.__new__(UserRoleGrantArgs)

            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            if user_id is None and not opts.urn:
                raise TypeError("Missing required property 'user_id'")
            __props__.__dict__["user_id"] = user_id
        super(UserRoleGrant, __self__).__init__(
            'cockroach:index/userRoleGrant:UserRoleGrant',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            role: Optional[pulumi.Input[Union['UserRoleGrantRoleArgs', 'UserRoleGrantRoleArgsDict']]] = None,
            user_id: Optional[pulumi.Input[str]] = None) -> 'UserRoleGrant':
        """
        Get an existing UserRoleGrant resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] user_id: ID of the user to grant these roles to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserRoleGrantState.__new__(_UserRoleGrantState)

        __props__.__dict__["role"] = role
        __props__.__dict__["user_id"] = user_id
        return UserRoleGrant(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output['outputs.UserRoleGrantRole']:
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[str]:
        """
        ID of the user to grant these roles to.
        """
        return pulumi.get(self, "user_id")

