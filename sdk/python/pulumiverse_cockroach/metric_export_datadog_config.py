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

__all__ = ['MetricExportDatadogConfigArgs', 'MetricExportDatadogConfig']

@pulumi.input_type
class MetricExportDatadogConfigArgs:
    def __init__(__self__, *,
                 api_key: pulumi.Input[str],
                 cluster_id: pulumi.Input[str],
                 site: pulumi.Input[str]):
        """
        The set of arguments for constructing a MetricExportDatadogConfig resource.
        :param pulumi.Input[str] api_key: A Datadog API key.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] site: The Datadog region to export to.
        """
        pulumi.set(__self__, "api_key", api_key)
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "site", site)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> pulumi.Input[str]:
        """
        A Datadog API key.
        """
        return pulumi.get(self, "api_key")

    @api_key.setter
    def api_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_key", value)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def site(self) -> pulumi.Input[str]:
        """
        The Datadog region to export to.
        """
        return pulumi.get(self, "site")

    @site.setter
    def site(self, value: pulumi.Input[str]):
        pulumi.set(self, "site", value)


@pulumi.input_type
class _MetricExportDatadogConfigState:
    def __init__(__self__, *,
                 api_key: Optional[pulumi.Input[str]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 site: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 user_message: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MetricExportDatadogConfig resources.
        :param pulumi.Input[str] api_key: A Datadog API key.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] site: The Datadog region to export to.
        :param pulumi.Input[str] status: Encodes the possible states that a metric export configuration can be in as it is created, deployed, and disabled.
        :param pulumi.Input[str] user_message: Elaborates on the metric export status and hints at how to fix issues that may have occurred during asynchronous operations.
        """
        if api_key is not None:
            pulumi.set(__self__, "api_key", api_key)
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if site is not None:
            pulumi.set(__self__, "site", site)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if user_message is not None:
            pulumi.set(__self__, "user_message", user_message)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> Optional[pulumi.Input[str]]:
        """
        A Datadog API key.
        """
        return pulumi.get(self, "api_key")

    @api_key.setter
    def api_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_key", value)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def site(self) -> Optional[pulumi.Input[str]]:
        """
        The Datadog region to export to.
        """
        return pulumi.get(self, "site")

    @site.setter
    def site(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "site", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Encodes the possible states that a metric export configuration can be in as it is created, deployed, and disabled.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="userMessage")
    def user_message(self) -> Optional[pulumi.Input[str]]:
        """
        Elaborates on the metric export status and hints at how to fix issues that may have occurred during asynchronous operations.
        """
        return pulumi.get(self, "user_message")

    @user_message.setter
    def user_message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_message", value)


class MetricExportDatadogConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_key: Optional[pulumi.Input[str]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 site: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        DataDog metric export configuration for a cluster.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_cockroach as cockroach

        config = pulumi.Config()
        cluster_id = config.require("clusterId")
        datadog_site = config.require("datadogSite")
        datadog_api_key = config.require("datadogApiKey")
        example = cockroach.MetricExportDatadogConfig("example",
            cluster_id=cluster_id,
            site=datadog_site,
            api_key=datadog_api_key)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_key: A Datadog API key.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] site: The Datadog region to export to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MetricExportDatadogConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        DataDog metric export configuration for a cluster.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_cockroach as cockroach

        config = pulumi.Config()
        cluster_id = config.require("clusterId")
        datadog_site = config.require("datadogSite")
        datadog_api_key = config.require("datadogApiKey")
        example = cockroach.MetricExportDatadogConfig("example",
            cluster_id=cluster_id,
            site=datadog_site,
            api_key=datadog_api_key)
        ```

        :param str resource_name: The name of the resource.
        :param MetricExportDatadogConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MetricExportDatadogConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_key: Optional[pulumi.Input[str]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 site: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MetricExportDatadogConfigArgs.__new__(MetricExportDatadogConfigArgs)

            if api_key is None and not opts.urn:
                raise TypeError("Missing required property 'api_key'")
            __props__.__dict__["api_key"] = None if api_key is None else pulumi.Output.secret(api_key)
            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            if site is None and not opts.urn:
                raise TypeError("Missing required property 'site'")
            __props__.__dict__["site"] = site
            __props__.__dict__["status"] = None
            __props__.__dict__["user_message"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["apiKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(MetricExportDatadogConfig, __self__).__init__(
            'cockroach:index/metricExportDatadogConfig:MetricExportDatadogConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_key: Optional[pulumi.Input[str]] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            site: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            user_message: Optional[pulumi.Input[str]] = None) -> 'MetricExportDatadogConfig':
        """
        Get an existing MetricExportDatadogConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_key: A Datadog API key.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] site: The Datadog region to export to.
        :param pulumi.Input[str] status: Encodes the possible states that a metric export configuration can be in as it is created, deployed, and disabled.
        :param pulumi.Input[str] user_message: Elaborates on the metric export status and hints at how to fix issues that may have occurred during asynchronous operations.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MetricExportDatadogConfigState.__new__(_MetricExportDatadogConfigState)

        __props__.__dict__["api_key"] = api_key
        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["site"] = site
        __props__.__dict__["status"] = status
        __props__.__dict__["user_message"] = user_message
        return MetricExportDatadogConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> pulumi.Output[str]:
        """
        A Datadog API key.
        """
        return pulumi.get(self, "api_key")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def site(self) -> pulumi.Output[str]:
        """
        The Datadog region to export to.
        """
        return pulumi.get(self, "site")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Encodes the possible states that a metric export configuration can be in as it is created, deployed, and disabled.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="userMessage")
    def user_message(self) -> pulumi.Output[str]:
        """
        Elaborates on the metric export status and hints at how to fix issues that may have occurred during asynchronous operations.
        """
        return pulumi.get(self, "user_message")
