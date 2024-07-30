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
    'GetGatewayAclResult',
    'AwaitableGetGatewayAclResult',
    'get_gateway_acl',
    'get_gateway_acl_output',
]

@pulumi.output_type
class GetGatewayAclResult:
    """
    A collection of values returned by getGatewayAcl.
    """
    def __init__(__self__, consumer=None, consumer_id=None, control_plane_id=None, created_at=None, group=None, id=None, tags=None):
        if consumer and not isinstance(consumer, dict):
            raise TypeError("Expected argument 'consumer' to be a dict")
        pulumi.set(__self__, "consumer", consumer)
        if consumer_id and not isinstance(consumer_id, str):
            raise TypeError("Expected argument 'consumer_id' to be a str")
        pulumi.set(__self__, "consumer_id", consumer_id)
        if control_plane_id and not isinstance(control_plane_id, str):
            raise TypeError("Expected argument 'control_plane_id' to be a str")
        pulumi.set(__self__, "control_plane_id", control_plane_id)
        if created_at and not isinstance(created_at, int):
            raise TypeError("Expected argument 'created_at' to be a int")
        pulumi.set(__self__, "created_at", created_at)
        if group and not isinstance(group, str):
            raise TypeError("Expected argument 'group' to be a str")
        pulumi.set(__self__, "group", group)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def consumer(self) -> 'outputs.GetGatewayAclConsumerResult':
        return pulumi.get(self, "consumer")

    @property
    @pulumi.getter(name="consumerId")
    def consumer_id(self) -> str:
        return pulumi.get(self, "consumer_id")

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
    def group(self) -> str:
        return pulumi.get(self, "group")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        return pulumi.get(self, "tags")


class AwaitableGetGatewayAclResult(GetGatewayAclResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGatewayAclResult(
            consumer=self.consumer,
            consumer_id=self.consumer_id,
            control_plane_id=self.control_plane_id,
            created_at=self.created_at,
            group=self.group,
            id=self.id,
            tags=self.tags)


def get_gateway_acl(consumer_id: Optional[str] = None,
                    control_plane_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGatewayAclResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['consumerId'] = consumer_id
    __args__['controlPlaneId'] = control_plane_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('konnect:index/getGatewayAcl:getGatewayAcl', __args__, opts=opts, typ=GetGatewayAclResult).value

    return AwaitableGetGatewayAclResult(
        consumer=pulumi.get(__ret__, 'consumer'),
        consumer_id=pulumi.get(__ret__, 'consumer_id'),
        control_plane_id=pulumi.get(__ret__, 'control_plane_id'),
        created_at=pulumi.get(__ret__, 'created_at'),
        group=pulumi.get(__ret__, 'group'),
        id=pulumi.get(__ret__, 'id'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_gateway_acl)
def get_gateway_acl_output(consumer_id: Optional[pulumi.Input[str]] = None,
                           control_plane_id: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGatewayAclResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
