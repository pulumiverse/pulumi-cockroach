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

__all__ = ['FolderArgs', 'Folder']

@pulumi.input_type
class FolderArgs:
    def __init__(__self__, *,
                 parent_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Folder resource.
        :param pulumi.Input[str] parent_id: ID of the parent folder. Use 'root' for the root level (no parent folder).
        :param pulumi.Input[str] name: Name of the folder.
        """
        pulumi.set(__self__, "parent_id", parent_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> pulumi.Input[str]:
        """
        ID of the parent folder. Use 'root' for the root level (no parent folder).
        """
        return pulumi.get(self, "parent_id")

    @parent_id.setter
    def parent_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "parent_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the folder.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _FolderState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Folder resources.
        :param pulumi.Input[str] name: Name of the folder.
        :param pulumi.Input[str] parent_id: ID of the parent folder. Use 'root' for the root level (no parent folder).
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parent_id is not None:
            pulumi.set(__self__, "parent_id", parent_id)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the folder.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the parent folder. Use 'root' for the root level (no parent folder).
        """
        return pulumi.get(self, "parent_id")

    @parent_id.setter
    def parent_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_id", value)


class Folder(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        CockroachDB Cloud folder.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_cockroach as cockroach

        a_team = cockroach.Folder("a_team",
            name="a-team",
            parent_id="root")
        a_team_dev = cockroach.Folder("a_team_dev",
            name="dev",
            parent_id=a_team.id)
        ```

        ## Import

        format: <folder id>

        ```sh
        $ pulumi import cockroach:index/folder:Folder my_folder 1f69fdd2-600a-4cfc-a9ba-16995df0d77d
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name of the folder.
        :param pulumi.Input[str] parent_id: ID of the parent folder. Use 'root' for the root level (no parent folder).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FolderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        CockroachDB Cloud folder.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_cockroach as cockroach

        a_team = cockroach.Folder("a_team",
            name="a-team",
            parent_id="root")
        a_team_dev = cockroach.Folder("a_team_dev",
            name="dev",
            parent_id=a_team.id)
        ```

        ## Import

        format: <folder id>

        ```sh
        $ pulumi import cockroach:index/folder:Folder my_folder 1f69fdd2-600a-4cfc-a9ba-16995df0d77d
        ```

        :param str resource_name: The name of the resource.
        :param FolderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FolderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FolderArgs.__new__(FolderArgs)

            __props__.__dict__["name"] = name
            if parent_id is None and not opts.urn:
                raise TypeError("Missing required property 'parent_id'")
            __props__.__dict__["parent_id"] = parent_id
        super(Folder, __self__).__init__(
            'cockroach:index/folder:Folder',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            parent_id: Optional[pulumi.Input[str]] = None) -> 'Folder':
        """
        Get an existing Folder resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name of the folder.
        :param pulumi.Input[str] parent_id: ID of the parent folder. Use 'root' for the root level (no parent folder).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FolderState.__new__(_FolderState)

        __props__.__dict__["name"] = name
        __props__.__dict__["parent_id"] = parent_id
        return Folder(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the folder.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> pulumi.Output[str]:
        """
        ID of the parent folder. Use 'root' for the root level (no parent folder).
        """
        return pulumi.get(self, "parent_id")
