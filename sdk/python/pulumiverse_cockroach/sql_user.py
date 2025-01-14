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

__all__ = ['SqlUserArgs', 'SqlUser']

@pulumi.input_type
class SqlUserArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SqlUser resource.
        :param pulumi.Input[str] name: SQL user name.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        SQL user name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)


@pulumi.input_type
class _SqlUserState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SqlUser resources.
        :param pulumi.Input[str] name: SQL user name.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        SQL user name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)


class SqlUser(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        CockroachDB SQL user.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_cockroach as cockroach

        config = pulumi.Config()
        cluster_id = config.require("clusterId")
        sql_user_password = config.require("sqlUserPassword")
        cockroach = cockroach.SqlUser("cockroach",
            name="example-sql-user",
            password=sql_user_password,
            cluster_id=cluster_id)
        ```

        ## Import

        format: <cluster id>:<sql user name>

        ```sh
        $ pulumi import cockroach:index/sqlUser:SqlUser bill 1f69fdd2-600a-4cfc-a9ba-16995df0d77d:bill
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: SQL user name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SqlUserArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        CockroachDB SQL user.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_cockroach as cockroach

        config = pulumi.Config()
        cluster_id = config.require("clusterId")
        sql_user_password = config.require("sqlUserPassword")
        cockroach = cockroach.SqlUser("cockroach",
            name="example-sql-user",
            password=sql_user_password,
            cluster_id=cluster_id)
        ```

        ## Import

        format: <cluster id>:<sql user name>

        ```sh
        $ pulumi import cockroach:index/sqlUser:SqlUser bill 1f69fdd2-600a-4cfc-a9ba-16995df0d77d:bill
        ```

        :param str resource_name: The name of the resource.
        :param SqlUserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SqlUserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SqlUserArgs.__new__(SqlUserArgs)

            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            __props__.__dict__["name"] = name
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(SqlUser, __self__).__init__(
            'cockroach:index/sqlUser:SqlUser',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None) -> 'SqlUser':
        """
        Get an existing SqlUser resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: SQL user name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SqlUserState.__new__(_SqlUserState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        return SqlUser(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        SQL user name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "password")

