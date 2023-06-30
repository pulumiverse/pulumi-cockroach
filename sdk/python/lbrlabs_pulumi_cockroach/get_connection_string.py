# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetConnectionStringResult',
    'AwaitableGetConnectionStringResult',
    'get_connection_string',
    'get_connection_string_output',
]

@pulumi.output_type
class GetConnectionStringResult:
    """
    A collection of values returned by getConnectionString.
    """
    def __init__(__self__, connection_params=None, connection_string=None, database=None, id=None, os=None, password=None, sql_user=None):
        if connection_params and not isinstance(connection_params, dict):
            raise TypeError("Expected argument 'connection_params' to be a dict")
        pulumi.set(__self__, "connection_params", connection_params)
        if connection_string and not isinstance(connection_string, str):
            raise TypeError("Expected argument 'connection_string' to be a str")
        pulumi.set(__self__, "connection_string", connection_string)
        if database and not isinstance(database, str):
            raise TypeError("Expected argument 'database' to be a str")
        pulumi.set(__self__, "database", database)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if os and not isinstance(os, str):
            raise TypeError("Expected argument 'os' to be a str")
        pulumi.set(__self__, "os", os)
        if password and not isinstance(password, str):
            raise TypeError("Expected argument 'password' to be a str")
        pulumi.set(__self__, "password", password)
        if sql_user and not isinstance(sql_user, str):
            raise TypeError("Expected argument 'sql_user' to be a str")
        pulumi.set(__self__, "sql_user", sql_user)

    @property
    @pulumi.getter(name="connectionParams")
    def connection_params(self) -> Mapping[str, str]:
        """
        List of individual connection string parameters. Can be used to build nonstandard connection strings.
        """
        return pulumi.get(self, "connection_params")

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> str:
        """
        Fully formatted connection string. Assumes the cluster certificate is stored in the default location.
        """
        return pulumi.get(self, "connection_string")

    @property
    @pulumi.getter
    def database(self) -> str:
        """
        Database to connect to. Defaults to 'defaultdb'.
        """
        return pulumi.get(self, "database")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Cluster ID
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def os(self) -> str:
        return pulumi.get(self, "os")

    @property
    @pulumi.getter
    def password(self) -> Optional[str]:
        """
        Database user password. Must also include `sql_user`.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="sqlUser")
    def sql_user(self) -> Optional[str]:
        """
        Database username.
        """
        return pulumi.get(self, "sql_user")


class AwaitableGetConnectionStringResult(GetConnectionStringResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConnectionStringResult(
            connection_params=self.connection_params,
            connection_string=self.connection_string,
            database=self.database,
            id=self.id,
            os=self.os,
            password=self.password,
            sql_user=self.sql_user)


def get_connection_string(database: Optional[str] = None,
                          id: Optional[str] = None,
                          os: Optional[str] = None,
                          password: Optional[str] = None,
                          sql_user: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConnectionStringResult:
    """
    Generic connection string for a given cluster


    :param str database: Database to connect to. Defaults to 'defaultdb'.
    :param str id: Cluster ID
    :param str password: Database user password. Must also include `sql_user`.
    :param str sql_user: Database username.
    """
    __args__ = dict()
    __args__['database'] = database
    __args__['id'] = id
    __args__['os'] = os
    __args__['password'] = password
    __args__['sqlUser'] = sql_user
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('cockroach:index/getConnectionString:getConnectionString', __args__, opts=opts, typ=GetConnectionStringResult).value

    return AwaitableGetConnectionStringResult(
        connection_params=__ret__.connection_params,
        connection_string=__ret__.connection_string,
        database=__ret__.database,
        id=__ret__.id,
        os=__ret__.os,
        password=__ret__.password,
        sql_user=__ret__.sql_user)


@_utilities.lift_output_func(get_connection_string)
def get_connection_string_output(database: Optional[pulumi.Input[Optional[str]]] = None,
                                 id: Optional[pulumi.Input[str]] = None,
                                 os: Optional[pulumi.Input[Optional[str]]] = None,
                                 password: Optional[pulumi.Input[Optional[str]]] = None,
                                 sql_user: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetConnectionStringResult]:
    """
    Generic connection string for a given cluster


    :param str database: Database to connect to. Defaults to 'defaultdb'.
    :param str id: Cluster ID
    :param str password: Database user password. Must also include `sql_user`.
    :param str sql_user: Database username.
    """
    ...