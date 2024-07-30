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
    'GetCloudGatewayTransitGatewayResult',
    'AwaitableGetCloudGatewayTransitGatewayResult',
    'get_cloud_gateway_transit_gateway',
    'get_cloud_gateway_transit_gateway_output',
]

@pulumi.output_type
class GetCloudGatewayTransitGatewayResult:
    """
    A collection of values returned by getCloudGatewayTransitGateway.
    """
    def __init__(__self__, cidr_blocks=None, created_at=None, dns_configs=None, entity_version=None, id=None, name=None, network_id=None, state=None, transit_gateway_attachment_config=None, updated_at=None):
        if cidr_blocks and not isinstance(cidr_blocks, list):
            raise TypeError("Expected argument 'cidr_blocks' to be a list")
        pulumi.set(__self__, "cidr_blocks", cidr_blocks)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if dns_configs and not isinstance(dns_configs, list):
            raise TypeError("Expected argument 'dns_configs' to be a list")
        pulumi.set(__self__, "dns_configs", dns_configs)
        if entity_version and not isinstance(entity_version, int):
            raise TypeError("Expected argument 'entity_version' to be a int")
        pulumi.set(__self__, "entity_version", entity_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_id and not isinstance(network_id, str):
            raise TypeError("Expected argument 'network_id' to be a str")
        pulumi.set(__self__, "network_id", network_id)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if transit_gateway_attachment_config and not isinstance(transit_gateway_attachment_config, dict):
            raise TypeError("Expected argument 'transit_gateway_attachment_config' to be a dict")
        pulumi.set(__self__, "transit_gateway_attachment_config", transit_gateway_attachment_config)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="cidrBlocks")
    def cidr_blocks(self) -> Sequence[str]:
        return pulumi.get(self, "cidr_blocks")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="dnsConfigs")
    def dns_configs(self) -> Sequence['outputs.GetCloudGatewayTransitGatewayDnsConfigResult']:
        return pulumi.get(self, "dns_configs")

    @property
    @pulumi.getter(name="entityVersion")
    def entity_version(self) -> int:
        return pulumi.get(self, "entity_version")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> str:
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter
    def state(self) -> str:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="transitGatewayAttachmentConfig")
    def transit_gateway_attachment_config(self) -> 'outputs.GetCloudGatewayTransitGatewayTransitGatewayAttachmentConfigResult':
        return pulumi.get(self, "transit_gateway_attachment_config")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        return pulumi.get(self, "updated_at")


class AwaitableGetCloudGatewayTransitGatewayResult(GetCloudGatewayTransitGatewayResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCloudGatewayTransitGatewayResult(
            cidr_blocks=self.cidr_blocks,
            created_at=self.created_at,
            dns_configs=self.dns_configs,
            entity_version=self.entity_version,
            id=self.id,
            name=self.name,
            network_id=self.network_id,
            state=self.state,
            transit_gateway_attachment_config=self.transit_gateway_attachment_config,
            updated_at=self.updated_at)


def get_cloud_gateway_transit_gateway(network_id: Optional[str] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCloudGatewayTransitGatewayResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['networkId'] = network_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('konnect:index/getCloudGatewayTransitGateway:getCloudGatewayTransitGateway', __args__, opts=opts, typ=GetCloudGatewayTransitGatewayResult).value

    return AwaitableGetCloudGatewayTransitGatewayResult(
        cidr_blocks=pulumi.get(__ret__, 'cidr_blocks'),
        created_at=pulumi.get(__ret__, 'created_at'),
        dns_configs=pulumi.get(__ret__, 'dns_configs'),
        entity_version=pulumi.get(__ret__, 'entity_version'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        network_id=pulumi.get(__ret__, 'network_id'),
        state=pulumi.get(__ret__, 'state'),
        transit_gateway_attachment_config=pulumi.get(__ret__, 'transit_gateway_attachment_config'),
        updated_at=pulumi.get(__ret__, 'updated_at'))


@_utilities.lift_output_func(get_cloud_gateway_transit_gateway)
def get_cloud_gateway_transit_gateway_output(network_id: Optional[pulumi.Input[str]] = None,
                                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCloudGatewayTransitGatewayResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
