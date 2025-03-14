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

__all__ = ['ClusterArgs', 'Cluster']

@pulumi.input_type
class ClusterArgs:
    def __init__(__self__, *,
                 cloud_provider: pulumi.Input[str],
                 regions: pulumi.Input[Sequence[pulumi.Input['ClusterRegionArgs']]],
                 backup_config: Optional[pulumi.Input['ClusterBackupConfigArgs']] = None,
                 cockroach_version: Optional[pulumi.Input[str]] = None,
                 dedicated: Optional[pulumi.Input['ClusterDedicatedArgs']] = None,
                 delete_protection: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_id: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input[str]] = None,
                 serverless: Optional[pulumi.Input['ClusterServerlessArgs']] = None):
        """
        The set of arguments for constructing a Cluster resource.
        :param pulumi.Input[str] cloud_provider: Cloud provider used to host the cluster. Allowed values are: * GCP * AWS * AZURE
        :param pulumi.Input['ClusterBackupConfigArgs'] backup_config: The backup settings for a cluster. Each cluster has backup settings that determine if backups are enabled, how
               frequently they are taken, and how long they are retained for. Use this attribute to manage those settings.
        :param pulumi.Input[str] cockroach_version: Major version of CockroachDB running on the cluster. This value can be used to orchestrate version upgrades. Supported
               for ADVANCED and STANDARD clusters (when `serverless.upgrade_type` set to 'MANUAL').
        :param pulumi.Input[bool] delete_protection: Set to true to enable delete protection on the cluster. If unset, the server chooses the value on cluster creation, and
               preserves the value on cluster update.
        :param pulumi.Input[str] name: Name of the cluster.
        :param pulumi.Input[str] parent_id: The ID of the cluster's parent folder. 'root' is used for a cluster at the root level.
        :param pulumi.Input[str] plan: Denotes cluster plan type: 'BASIC' or 'STANDARD' or 'ADVANCED'.
        """
        pulumi.set(__self__, "cloud_provider", cloud_provider)
        pulumi.set(__self__, "regions", regions)
        if backup_config is not None:
            pulumi.set(__self__, "backup_config", backup_config)
        if cockroach_version is not None:
            pulumi.set(__self__, "cockroach_version", cockroach_version)
        if dedicated is not None:
            pulumi.set(__self__, "dedicated", dedicated)
        if delete_protection is not None:
            pulumi.set(__self__, "delete_protection", delete_protection)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parent_id is not None:
            pulumi.set(__self__, "parent_id", parent_id)
        if plan is not None:
            pulumi.set(__self__, "plan", plan)
        if serverless is not None:
            pulumi.set(__self__, "serverless", serverless)

    @property
    @pulumi.getter(name="cloudProvider")
    def cloud_provider(self) -> pulumi.Input[str]:
        """
        Cloud provider used to host the cluster. Allowed values are: * GCP * AWS * AZURE
        """
        return pulumi.get(self, "cloud_provider")

    @cloud_provider.setter
    def cloud_provider(self, value: pulumi.Input[str]):
        pulumi.set(self, "cloud_provider", value)

    @property
    @pulumi.getter
    def regions(self) -> pulumi.Input[Sequence[pulumi.Input['ClusterRegionArgs']]]:
        return pulumi.get(self, "regions")

    @regions.setter
    def regions(self, value: pulumi.Input[Sequence[pulumi.Input['ClusterRegionArgs']]]):
        pulumi.set(self, "regions", value)

    @property
    @pulumi.getter(name="backupConfig")
    def backup_config(self) -> Optional[pulumi.Input['ClusterBackupConfigArgs']]:
        """
        The backup settings for a cluster. Each cluster has backup settings that determine if backups are enabled, how
        frequently they are taken, and how long they are retained for. Use this attribute to manage those settings.
        """
        return pulumi.get(self, "backup_config")

    @backup_config.setter
    def backup_config(self, value: Optional[pulumi.Input['ClusterBackupConfigArgs']]):
        pulumi.set(self, "backup_config", value)

    @property
    @pulumi.getter(name="cockroachVersion")
    def cockroach_version(self) -> Optional[pulumi.Input[str]]:
        """
        Major version of CockroachDB running on the cluster. This value can be used to orchestrate version upgrades. Supported
        for ADVANCED and STANDARD clusters (when `serverless.upgrade_type` set to 'MANUAL').
        """
        return pulumi.get(self, "cockroach_version")

    @cockroach_version.setter
    def cockroach_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cockroach_version", value)

    @property
    @pulumi.getter
    def dedicated(self) -> Optional[pulumi.Input['ClusterDedicatedArgs']]:
        return pulumi.get(self, "dedicated")

    @dedicated.setter
    def dedicated(self, value: Optional[pulumi.Input['ClusterDedicatedArgs']]):
        pulumi.set(self, "dedicated", value)

    @property
    @pulumi.getter(name="deleteProtection")
    def delete_protection(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to true to enable delete protection on the cluster. If unset, the server chooses the value on cluster creation, and
        preserves the value on cluster update.
        """
        return pulumi.get(self, "delete_protection")

    @delete_protection.setter
    def delete_protection(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_protection", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the cluster.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the cluster's parent folder. 'root' is used for a cluster at the root level.
        """
        return pulumi.get(self, "parent_id")

    @parent_id.setter
    def parent_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_id", value)

    @property
    @pulumi.getter
    def plan(self) -> Optional[pulumi.Input[str]]:
        """
        Denotes cluster plan type: 'BASIC' or 'STANDARD' or 'ADVANCED'.
        """
        return pulumi.get(self, "plan")

    @plan.setter
    def plan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plan", value)

    @property
    @pulumi.getter
    def serverless(self) -> Optional[pulumi.Input['ClusterServerlessArgs']]:
        return pulumi.get(self, "serverless")

    @serverless.setter
    def serverless(self, value: Optional[pulumi.Input['ClusterServerlessArgs']]):
        pulumi.set(self, "serverless", value)


@pulumi.input_type
class _ClusterState:
    def __init__(__self__, *,
                 account_id: Optional[pulumi.Input[str]] = None,
                 backup_config: Optional[pulumi.Input['ClusterBackupConfigArgs']] = None,
                 cloud_provider: Optional[pulumi.Input[str]] = None,
                 cockroach_version: Optional[pulumi.Input[str]] = None,
                 creator_id: Optional[pulumi.Input[str]] = None,
                 dedicated: Optional[pulumi.Input['ClusterDedicatedArgs']] = None,
                 delete_protection: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 operation_status: Optional[pulumi.Input[str]] = None,
                 parent_id: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input[str]] = None,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterRegionArgs']]]] = None,
                 serverless: Optional[pulumi.Input['ClusterServerlessArgs']] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 upgrade_status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Cluster resources.
        :param pulumi.Input[str] account_id: The cloud provider account ID that hosts the cluster. Needed for CMEK and other advanced features.
        :param pulumi.Input['ClusterBackupConfigArgs'] backup_config: The backup settings for a cluster. Each cluster has backup settings that determine if backups are enabled, how
               frequently they are taken, and how long they are retained for. Use this attribute to manage those settings.
        :param pulumi.Input[str] cloud_provider: Cloud provider used to host the cluster. Allowed values are: * GCP * AWS * AZURE
        :param pulumi.Input[str] cockroach_version: Major version of CockroachDB running on the cluster. This value can be used to orchestrate version upgrades. Supported
               for ADVANCED and STANDARD clusters (when `serverless.upgrade_type` set to 'MANUAL').
        :param pulumi.Input[str] creator_id: ID of the user who created the cluster.
        :param pulumi.Input[bool] delete_protection: Set to true to enable delete protection on the cluster. If unset, the server chooses the value on cluster creation, and
               preserves the value on cluster update.
        :param pulumi.Input[str] name: Name of the cluster.
        :param pulumi.Input[str] operation_status: Describes the current long-running operation, if any.
        :param pulumi.Input[str] parent_id: The ID of the cluster's parent folder. 'root' is used for a cluster at the root level.
        :param pulumi.Input[str] plan: Denotes cluster plan type: 'BASIC' or 'STANDARD' or 'ADVANCED'.
        :param pulumi.Input[str] state: Describes whether the cluster is being created, updated, deleted, etc.
        :param pulumi.Input[str] upgrade_status: Describes the status of any in-progress CockroachDB upgrade or rollback.
        """
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if backup_config is not None:
            pulumi.set(__self__, "backup_config", backup_config)
        if cloud_provider is not None:
            pulumi.set(__self__, "cloud_provider", cloud_provider)
        if cockroach_version is not None:
            pulumi.set(__self__, "cockroach_version", cockroach_version)
        if creator_id is not None:
            pulumi.set(__self__, "creator_id", creator_id)
        if dedicated is not None:
            pulumi.set(__self__, "dedicated", dedicated)
        if delete_protection is not None:
            pulumi.set(__self__, "delete_protection", delete_protection)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if operation_status is not None:
            pulumi.set(__self__, "operation_status", operation_status)
        if parent_id is not None:
            pulumi.set(__self__, "parent_id", parent_id)
        if plan is not None:
            pulumi.set(__self__, "plan", plan)
        if regions is not None:
            pulumi.set(__self__, "regions", regions)
        if serverless is not None:
            pulumi.set(__self__, "serverless", serverless)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if upgrade_status is not None:
            pulumi.set(__self__, "upgrade_status", upgrade_status)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The cloud provider account ID that hosts the cluster. Needed for CMEK and other advanced features.
        """
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter(name="backupConfig")
    def backup_config(self) -> Optional[pulumi.Input['ClusterBackupConfigArgs']]:
        """
        The backup settings for a cluster. Each cluster has backup settings that determine if backups are enabled, how
        frequently they are taken, and how long they are retained for. Use this attribute to manage those settings.
        """
        return pulumi.get(self, "backup_config")

    @backup_config.setter
    def backup_config(self, value: Optional[pulumi.Input['ClusterBackupConfigArgs']]):
        pulumi.set(self, "backup_config", value)

    @property
    @pulumi.getter(name="cloudProvider")
    def cloud_provider(self) -> Optional[pulumi.Input[str]]:
        """
        Cloud provider used to host the cluster. Allowed values are: * GCP * AWS * AZURE
        """
        return pulumi.get(self, "cloud_provider")

    @cloud_provider.setter
    def cloud_provider(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cloud_provider", value)

    @property
    @pulumi.getter(name="cockroachVersion")
    def cockroach_version(self) -> Optional[pulumi.Input[str]]:
        """
        Major version of CockroachDB running on the cluster. This value can be used to orchestrate version upgrades. Supported
        for ADVANCED and STANDARD clusters (when `serverless.upgrade_type` set to 'MANUAL').
        """
        return pulumi.get(self, "cockroach_version")

    @cockroach_version.setter
    def cockroach_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cockroach_version", value)

    @property
    @pulumi.getter(name="creatorId")
    def creator_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the user who created the cluster.
        """
        return pulumi.get(self, "creator_id")

    @creator_id.setter
    def creator_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creator_id", value)

    @property
    @pulumi.getter
    def dedicated(self) -> Optional[pulumi.Input['ClusterDedicatedArgs']]:
        return pulumi.get(self, "dedicated")

    @dedicated.setter
    def dedicated(self, value: Optional[pulumi.Input['ClusterDedicatedArgs']]):
        pulumi.set(self, "dedicated", value)

    @property
    @pulumi.getter(name="deleteProtection")
    def delete_protection(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to true to enable delete protection on the cluster. If unset, the server chooses the value on cluster creation, and
        preserves the value on cluster update.
        """
        return pulumi.get(self, "delete_protection")

    @delete_protection.setter
    def delete_protection(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_protection", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the cluster.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="operationStatus")
    def operation_status(self) -> Optional[pulumi.Input[str]]:
        """
        Describes the current long-running operation, if any.
        """
        return pulumi.get(self, "operation_status")

    @operation_status.setter
    def operation_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "operation_status", value)

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the cluster's parent folder. 'root' is used for a cluster at the root level.
        """
        return pulumi.get(self, "parent_id")

    @parent_id.setter
    def parent_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_id", value)

    @property
    @pulumi.getter
    def plan(self) -> Optional[pulumi.Input[str]]:
        """
        Denotes cluster plan type: 'BASIC' or 'STANDARD' or 'ADVANCED'.
        """
        return pulumi.get(self, "plan")

    @plan.setter
    def plan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plan", value)

    @property
    @pulumi.getter
    def regions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ClusterRegionArgs']]]]:
        return pulumi.get(self, "regions")

    @regions.setter
    def regions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterRegionArgs']]]]):
        pulumi.set(self, "regions", value)

    @property
    @pulumi.getter
    def serverless(self) -> Optional[pulumi.Input['ClusterServerlessArgs']]:
        return pulumi.get(self, "serverless")

    @serverless.setter
    def serverless(self, value: Optional[pulumi.Input['ClusterServerlessArgs']]):
        pulumi.set(self, "serverless", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        Describes whether the cluster is being created, updated, deleted, etc.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="upgradeStatus")
    def upgrade_status(self) -> Optional[pulumi.Input[str]]:
        """
        Describes the status of any in-progress CockroachDB upgrade or rollback.
        """
        return pulumi.get(self, "upgrade_status")

    @upgrade_status.setter
    def upgrade_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "upgrade_status", value)


class Cluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_config: Optional[pulumi.Input[Union['ClusterBackupConfigArgs', 'ClusterBackupConfigArgsDict']]] = None,
                 cloud_provider: Optional[pulumi.Input[str]] = None,
                 cockroach_version: Optional[pulumi.Input[str]] = None,
                 dedicated: Optional[pulumi.Input[Union['ClusterDedicatedArgs', 'ClusterDedicatedArgsDict']]] = None,
                 delete_protection: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_id: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input[str]] = None,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ClusterRegionArgs', 'ClusterRegionArgsDict']]]]] = None,
                 serverless: Optional[pulumi.Input[Union['ClusterServerlessArgs', 'ClusterServerlessArgsDict']]] = None,
                 __props__=None):
        """
        CockroachDB Cloud cluster.

        ## Import

        format: <cluster id>

        ```sh
        $ pulumi import cockroach:index/cluster:Cluster my_cluster 1f69fdd2-600a-4cfc-a9ba-16995df0d77d
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['ClusterBackupConfigArgs', 'ClusterBackupConfigArgsDict']] backup_config: The backup settings for a cluster. Each cluster has backup settings that determine if backups are enabled, how
               frequently they are taken, and how long they are retained for. Use this attribute to manage those settings.
        :param pulumi.Input[str] cloud_provider: Cloud provider used to host the cluster. Allowed values are: * GCP * AWS * AZURE
        :param pulumi.Input[str] cockroach_version: Major version of CockroachDB running on the cluster. This value can be used to orchestrate version upgrades. Supported
               for ADVANCED and STANDARD clusters (when `serverless.upgrade_type` set to 'MANUAL').
        :param pulumi.Input[bool] delete_protection: Set to true to enable delete protection on the cluster. If unset, the server chooses the value on cluster creation, and
               preserves the value on cluster update.
        :param pulumi.Input[str] name: Name of the cluster.
        :param pulumi.Input[str] parent_id: The ID of the cluster's parent folder. 'root' is used for a cluster at the root level.
        :param pulumi.Input[str] plan: Denotes cluster plan type: 'BASIC' or 'STANDARD' or 'ADVANCED'.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClusterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        CockroachDB Cloud cluster.

        ## Import

        format: <cluster id>

        ```sh
        $ pulumi import cockroach:index/cluster:Cluster my_cluster 1f69fdd2-600a-4cfc-a9ba-16995df0d77d
        ```

        :param str resource_name: The name of the resource.
        :param ClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_config: Optional[pulumi.Input[Union['ClusterBackupConfigArgs', 'ClusterBackupConfigArgsDict']]] = None,
                 cloud_provider: Optional[pulumi.Input[str]] = None,
                 cockroach_version: Optional[pulumi.Input[str]] = None,
                 dedicated: Optional[pulumi.Input[Union['ClusterDedicatedArgs', 'ClusterDedicatedArgsDict']]] = None,
                 delete_protection: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_id: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input[str]] = None,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ClusterRegionArgs', 'ClusterRegionArgsDict']]]]] = None,
                 serverless: Optional[pulumi.Input[Union['ClusterServerlessArgs', 'ClusterServerlessArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClusterArgs.__new__(ClusterArgs)

            __props__.__dict__["backup_config"] = backup_config
            if cloud_provider is None and not opts.urn:
                raise TypeError("Missing required property 'cloud_provider'")
            __props__.__dict__["cloud_provider"] = cloud_provider
            __props__.__dict__["cockroach_version"] = cockroach_version
            __props__.__dict__["dedicated"] = dedicated
            __props__.__dict__["delete_protection"] = delete_protection
            __props__.__dict__["name"] = name
            __props__.__dict__["parent_id"] = parent_id
            __props__.__dict__["plan"] = plan
            if regions is None and not opts.urn:
                raise TypeError("Missing required property 'regions'")
            __props__.__dict__["regions"] = regions
            __props__.__dict__["serverless"] = serverless
            __props__.__dict__["account_id"] = None
            __props__.__dict__["creator_id"] = None
            __props__.__dict__["operation_status"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["upgrade_status"] = None
        super(Cluster, __self__).__init__(
            'cockroach:index/cluster:Cluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_id: Optional[pulumi.Input[str]] = None,
            backup_config: Optional[pulumi.Input[Union['ClusterBackupConfigArgs', 'ClusterBackupConfigArgsDict']]] = None,
            cloud_provider: Optional[pulumi.Input[str]] = None,
            cockroach_version: Optional[pulumi.Input[str]] = None,
            creator_id: Optional[pulumi.Input[str]] = None,
            dedicated: Optional[pulumi.Input[Union['ClusterDedicatedArgs', 'ClusterDedicatedArgsDict']]] = None,
            delete_protection: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            operation_status: Optional[pulumi.Input[str]] = None,
            parent_id: Optional[pulumi.Input[str]] = None,
            plan: Optional[pulumi.Input[str]] = None,
            regions: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ClusterRegionArgs', 'ClusterRegionArgsDict']]]]] = None,
            serverless: Optional[pulumi.Input[Union['ClusterServerlessArgs', 'ClusterServerlessArgsDict']]] = None,
            state: Optional[pulumi.Input[str]] = None,
            upgrade_status: Optional[pulumi.Input[str]] = None) -> 'Cluster':
        """
        Get an existing Cluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: The cloud provider account ID that hosts the cluster. Needed for CMEK and other advanced features.
        :param pulumi.Input[Union['ClusterBackupConfigArgs', 'ClusterBackupConfigArgsDict']] backup_config: The backup settings for a cluster. Each cluster has backup settings that determine if backups are enabled, how
               frequently they are taken, and how long they are retained for. Use this attribute to manage those settings.
        :param pulumi.Input[str] cloud_provider: Cloud provider used to host the cluster. Allowed values are: * GCP * AWS * AZURE
        :param pulumi.Input[str] cockroach_version: Major version of CockroachDB running on the cluster. This value can be used to orchestrate version upgrades. Supported
               for ADVANCED and STANDARD clusters (when `serverless.upgrade_type` set to 'MANUAL').
        :param pulumi.Input[str] creator_id: ID of the user who created the cluster.
        :param pulumi.Input[bool] delete_protection: Set to true to enable delete protection on the cluster. If unset, the server chooses the value on cluster creation, and
               preserves the value on cluster update.
        :param pulumi.Input[str] name: Name of the cluster.
        :param pulumi.Input[str] operation_status: Describes the current long-running operation, if any.
        :param pulumi.Input[str] parent_id: The ID of the cluster's parent folder. 'root' is used for a cluster at the root level.
        :param pulumi.Input[str] plan: Denotes cluster plan type: 'BASIC' or 'STANDARD' or 'ADVANCED'.
        :param pulumi.Input[str] state: Describes whether the cluster is being created, updated, deleted, etc.
        :param pulumi.Input[str] upgrade_status: Describes the status of any in-progress CockroachDB upgrade or rollback.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClusterState.__new__(_ClusterState)

        __props__.__dict__["account_id"] = account_id
        __props__.__dict__["backup_config"] = backup_config
        __props__.__dict__["cloud_provider"] = cloud_provider
        __props__.__dict__["cockroach_version"] = cockroach_version
        __props__.__dict__["creator_id"] = creator_id
        __props__.__dict__["dedicated"] = dedicated
        __props__.__dict__["delete_protection"] = delete_protection
        __props__.__dict__["name"] = name
        __props__.__dict__["operation_status"] = operation_status
        __props__.__dict__["parent_id"] = parent_id
        __props__.__dict__["plan"] = plan
        __props__.__dict__["regions"] = regions
        __props__.__dict__["serverless"] = serverless
        __props__.__dict__["state"] = state
        __props__.__dict__["upgrade_status"] = upgrade_status
        return Cluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Output[str]:
        """
        The cloud provider account ID that hosts the cluster. Needed for CMEK and other advanced features.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="backupConfig")
    def backup_config(self) -> pulumi.Output['outputs.ClusterBackupConfig']:
        """
        The backup settings for a cluster. Each cluster has backup settings that determine if backups are enabled, how
        frequently they are taken, and how long they are retained for. Use this attribute to manage those settings.
        """
        return pulumi.get(self, "backup_config")

    @property
    @pulumi.getter(name="cloudProvider")
    def cloud_provider(self) -> pulumi.Output[str]:
        """
        Cloud provider used to host the cluster. Allowed values are: * GCP * AWS * AZURE
        """
        return pulumi.get(self, "cloud_provider")

    @property
    @pulumi.getter(name="cockroachVersion")
    def cockroach_version(self) -> pulumi.Output[str]:
        """
        Major version of CockroachDB running on the cluster. This value can be used to orchestrate version upgrades. Supported
        for ADVANCED and STANDARD clusters (when `serverless.upgrade_type` set to 'MANUAL').
        """
        return pulumi.get(self, "cockroach_version")

    @property
    @pulumi.getter(name="creatorId")
    def creator_id(self) -> pulumi.Output[str]:
        """
        ID of the user who created the cluster.
        """
        return pulumi.get(self, "creator_id")

    @property
    @pulumi.getter
    def dedicated(self) -> pulumi.Output[Optional['outputs.ClusterDedicated']]:
        return pulumi.get(self, "dedicated")

    @property
    @pulumi.getter(name="deleteProtection")
    def delete_protection(self) -> pulumi.Output[bool]:
        """
        Set to true to enable delete protection on the cluster. If unset, the server chooses the value on cluster creation, and
        preserves the value on cluster update.
        """
        return pulumi.get(self, "delete_protection")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the cluster.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="operationStatus")
    def operation_status(self) -> pulumi.Output[str]:
        """
        Describes the current long-running operation, if any.
        """
        return pulumi.get(self, "operation_status")

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> pulumi.Output[str]:
        """
        The ID of the cluster's parent folder. 'root' is used for a cluster at the root level.
        """
        return pulumi.get(self, "parent_id")

    @property
    @pulumi.getter
    def plan(self) -> pulumi.Output[str]:
        """
        Denotes cluster plan type: 'BASIC' or 'STANDARD' or 'ADVANCED'.
        """
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter
    def regions(self) -> pulumi.Output[Sequence['outputs.ClusterRegion']]:
        return pulumi.get(self, "regions")

    @property
    @pulumi.getter
    def serverless(self) -> pulumi.Output[Optional['outputs.ClusterServerless']]:
        return pulumi.get(self, "serverless")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        Describes whether the cluster is being created, updated, deleted, etc.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="upgradeStatus")
    def upgrade_status(self) -> pulumi.Output[str]:
        """
        Describes the status of any in-progress CockroachDB upgrade or rollback.
        """
        return pulumi.get(self, "upgrade_status")

