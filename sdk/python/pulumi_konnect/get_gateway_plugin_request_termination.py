# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetGatewayPluginRequestTerminationResult',
    'AwaitableGetGatewayPluginRequestTerminationResult',
    'get_gateway_plugin_request_termination',
    'get_gateway_plugin_request_termination_output',
]

@pulumi.output_type
class GetGatewayPluginRequestTerminationResult:
    """
    A collection of values returned by getGatewayPluginRequestTermination.
    """
    def __init__(__self__, config=None, consumer=None, consumer_group=None, control_plane_id=None, created_at=None, enabled=None, id=None, instance_name=None, protocols=None, route=None, service=None, tags=None, updated_at=None):
        if config and not isinstance(config, dict):
            raise TypeError("Expected argument 'config' to be a dict")
        pulumi.set(__self__, "config", config)
        if consumer and not isinstance(consumer, dict):
            raise TypeError("Expected argument 'consumer' to be a dict")
        pulumi.set(__self__, "consumer", consumer)
        if consumer_group and not isinstance(consumer_group, dict):
            raise TypeError("Expected argument 'consumer_group' to be a dict")
        pulumi.set(__self__, "consumer_group", consumer_group)
        if control_plane_id and not isinstance(control_plane_id, str):
            raise TypeError("Expected argument 'control_plane_id' to be a str")
        pulumi.set(__self__, "control_plane_id", control_plane_id)
        if created_at and not isinstance(created_at, int):
            raise TypeError("Expected argument 'created_at' to be a int")
        pulumi.set(__self__, "created_at", created_at)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_name and not isinstance(instance_name, str):
            raise TypeError("Expected argument 'instance_name' to be a str")
        pulumi.set(__self__, "instance_name", instance_name)
        if protocols and not isinstance(protocols, list):
            raise TypeError("Expected argument 'protocols' to be a list")
        pulumi.set(__self__, "protocols", protocols)
        if route and not isinstance(route, dict):
            raise TypeError("Expected argument 'route' to be a dict")
        pulumi.set(__self__, "route", route)
        if service and not isinstance(service, dict):
            raise TypeError("Expected argument 'service' to be a dict")
        pulumi.set(__self__, "service", service)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if updated_at and not isinstance(updated_at, int):
            raise TypeError("Expected argument 'updated_at' to be a int")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter
    def config(self) -> 'outputs.GetGatewayPluginRequestTerminationConfigResult':
        return pulumi.get(self, "config")

    @property
    @pulumi.getter
    def consumer(self) -> 'outputs.GetGatewayPluginRequestTerminationConsumerResult':
        return pulumi.get(self, "consumer")

    @property
    @pulumi.getter(name="consumerGroup")
    def consumer_group(self) -> 'outputs.GetGatewayPluginRequestTerminationConsumerGroupResult':
        return pulumi.get(self, "consumer_group")

    @property
    @pulumi.getter(name="controlPlaneId")
    def control_plane_id(self) -> str:
        return pulumi.get(self, "control_plane_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> int:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> str:
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter
    def protocols(self) -> Sequence[str]:
        return pulumi.get(self, "protocols")

    @property
    @pulumi.getter
    def route(self) -> 'outputs.GetGatewayPluginRequestTerminationRouteResult':
        return pulumi.get(self, "route")

    @property
    @pulumi.getter
    def service(self) -> 'outputs.GetGatewayPluginRequestTerminationServiceResult':
        return pulumi.get(self, "service")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> int:
        return pulumi.get(self, "updated_at")


class AwaitableGetGatewayPluginRequestTerminationResult(GetGatewayPluginRequestTerminationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGatewayPluginRequestTerminationResult(
            config=self.config,
            consumer=self.consumer,
            consumer_group=self.consumer_group,
            control_plane_id=self.control_plane_id,
            created_at=self.created_at,
            enabled=self.enabled,
            id=self.id,
            instance_name=self.instance_name,
            protocols=self.protocols,
            route=self.route,
            service=self.service,
            tags=self.tags,
            updated_at=self.updated_at)


def get_gateway_plugin_request_termination(control_plane_id: Optional[str] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGatewayPluginRequestTerminationResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['controlPlaneId'] = control_plane_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('konnect:index/getGatewayPluginRequestTermination:getGatewayPluginRequestTermination', __args__, opts=opts, typ=GetGatewayPluginRequestTerminationResult).value

    return AwaitableGetGatewayPluginRequestTerminationResult(
        config=pulumi.get(__ret__, 'config'),
        consumer=pulumi.get(__ret__, 'consumer'),
        consumer_group=pulumi.get(__ret__, 'consumer_group'),
        control_plane_id=pulumi.get(__ret__, 'control_plane_id'),
        created_at=pulumi.get(__ret__, 'created_at'),
        enabled=pulumi.get(__ret__, 'enabled'),
        id=pulumi.get(__ret__, 'id'),
        instance_name=pulumi.get(__ret__, 'instance_name'),
        protocols=pulumi.get(__ret__, 'protocols'),
        route=pulumi.get(__ret__, 'route'),
        service=pulumi.get(__ret__, 'service'),
        tags=pulumi.get(__ret__, 'tags'),
        updated_at=pulumi.get(__ret__, 'updated_at'))


@_utilities.lift_output_func(get_gateway_plugin_request_termination)
def get_gateway_plugin_request_termination_output(control_plane_id: Optional[pulumi.Input[str]] = None,
                                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGatewayPluginRequestTerminationResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...