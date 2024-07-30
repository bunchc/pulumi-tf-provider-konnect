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
    'GetApiProductResult',
    'AwaitableGetApiProductResult',
    'get_api_product',
    'get_api_product_output',
]

@pulumi.output_type
class GetApiProductResult:
    """
    A collection of values returned by getApiProduct.
    """
    def __init__(__self__, created_at=None, description=None, id=None, labels=None, name=None, portal_ids=None, portals=None, updated_at=None, version_count=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if portal_ids and not isinstance(portal_ids, list):
            raise TypeError("Expected argument 'portal_ids' to be a list")
        pulumi.set(__self__, "portal_ids", portal_ids)
        if portals and not isinstance(portals, list):
            raise TypeError("Expected argument 'portals' to be a list")
        pulumi.set(__self__, "portals", portals)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)
        if version_count and not isinstance(version_count, float):
            raise TypeError("Expected argument 'version_count' to be a float")
        pulumi.set(__self__, "version_count", version_count)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="portalIds")
    def portal_ids(self) -> Sequence[str]:
        return pulumi.get(self, "portal_ids")

    @property
    @pulumi.getter
    def portals(self) -> Sequence['outputs.GetApiProductPortalResult']:
        return pulumi.get(self, "portals")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="versionCount")
    def version_count(self) -> float:
        return pulumi.get(self, "version_count")


class AwaitableGetApiProductResult(GetApiProductResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApiProductResult(
            created_at=self.created_at,
            description=self.description,
            id=self.id,
            labels=self.labels,
            name=self.name,
            portal_ids=self.portal_ids,
            portals=self.portals,
            updated_at=self.updated_at,
            version_count=self.version_count)


def get_api_product(id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApiProductResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('konnect:index/getApiProduct:getApiProduct', __args__, opts=opts, typ=GetApiProductResult).value

    return AwaitableGetApiProductResult(
        created_at=pulumi.get(__ret__, 'created_at'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        portal_ids=pulumi.get(__ret__, 'portal_ids'),
        portals=pulumi.get(__ret__, 'portals'),
        updated_at=pulumi.get(__ret__, 'updated_at'),
        version_count=pulumi.get(__ret__, 'version_count'))


@_utilities.lift_output_func(get_api_product)
def get_api_product_output(id: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetApiProductResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...