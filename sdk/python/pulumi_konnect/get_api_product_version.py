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
    'GetApiProductVersionResult',
    'AwaitableGetApiProductVersionResult',
    'get_api_product_version',
    'get_api_product_version_output',
]

@pulumi.output_type
class GetApiProductVersionResult:
    """
    A collection of values returned by getApiProductVersion.
    """
    def __init__(__self__, api_product_id=None, created_at=None, deprecated=None, gateway_service=None, id=None, name=None, portals=None, updated_at=None):
        if api_product_id and not isinstance(api_product_id, str):
            raise TypeError("Expected argument 'api_product_id' to be a str")
        pulumi.set(__self__, "api_product_id", api_product_id)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if deprecated and not isinstance(deprecated, bool):
            raise TypeError("Expected argument 'deprecated' to be a bool")
        pulumi.set(__self__, "deprecated", deprecated)
        if gateway_service and not isinstance(gateway_service, dict):
            raise TypeError("Expected argument 'gateway_service' to be a dict")
        pulumi.set(__self__, "gateway_service", gateway_service)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if portals and not isinstance(portals, list):
            raise TypeError("Expected argument 'portals' to be a list")
        pulumi.set(__self__, "portals", portals)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="apiProductId")
    def api_product_id(self) -> str:
        return pulumi.get(self, "api_product_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def deprecated(self) -> bool:
        return pulumi.get(self, "deprecated")

    @property
    @pulumi.getter(name="gatewayService")
    def gateway_service(self) -> 'outputs.GetApiProductVersionGatewayServiceResult':
        return pulumi.get(self, "gateway_service")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def portals(self) -> Sequence['outputs.GetApiProductVersionPortalResult']:
        return pulumi.get(self, "portals")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        return pulumi.get(self, "updated_at")


class AwaitableGetApiProductVersionResult(GetApiProductVersionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApiProductVersionResult(
            api_product_id=self.api_product_id,
            created_at=self.created_at,
            deprecated=self.deprecated,
            gateway_service=self.gateway_service,
            id=self.id,
            name=self.name,
            portals=self.portals,
            updated_at=self.updated_at)


def get_api_product_version(api_product_id: Optional[str] = None,
                            id: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApiProductVersionResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['apiProductId'] = api_product_id
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('konnect:index/getApiProductVersion:getApiProductVersion', __args__, opts=opts, typ=GetApiProductVersionResult).value

    return AwaitableGetApiProductVersionResult(
        api_product_id=pulumi.get(__ret__, 'api_product_id'),
        created_at=pulumi.get(__ret__, 'created_at'),
        deprecated=pulumi.get(__ret__, 'deprecated'),
        gateway_service=pulumi.get(__ret__, 'gateway_service'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        portals=pulumi.get(__ret__, 'portals'),
        updated_at=pulumi.get(__ret__, 'updated_at'))


@_utilities.lift_output_func(get_api_product_version)
def get_api_product_version_output(api_product_id: Optional[pulumi.Input[str]] = None,
                                   id: Optional[pulumi.Input[str]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetApiProductVersionResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
