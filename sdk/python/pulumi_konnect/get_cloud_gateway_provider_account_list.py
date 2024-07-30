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
    'GetCloudGatewayProviderAccountListResult',
    'AwaitableGetCloudGatewayProviderAccountListResult',
    'get_cloud_gateway_provider_account_list',
    'get_cloud_gateway_provider_account_list_output',
]

@pulumi.output_type
class GetCloudGatewayProviderAccountListResult:
    """
    A collection of values returned by getCloudGatewayProviderAccountList.
    """
    def __init__(__self__, datas=None, id=None, meta=None, page_number=None, page_size=None):
        if datas and not isinstance(datas, list):
            raise TypeError("Expected argument 'datas' to be a list")
        pulumi.set(__self__, "datas", datas)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if meta and not isinstance(meta, dict):
            raise TypeError("Expected argument 'meta' to be a dict")
        pulumi.set(__self__, "meta", meta)
        if page_number and not isinstance(page_number, int):
            raise TypeError("Expected argument 'page_number' to be a int")
        pulumi.set(__self__, "page_number", page_number)
        if page_size and not isinstance(page_size, int):
            raise TypeError("Expected argument 'page_size' to be a int")
        pulumi.set(__self__, "page_size", page_size)

    @property
    @pulumi.getter
    def datas(self) -> Sequence['outputs.GetCloudGatewayProviderAccountListDataResult']:
        return pulumi.get(self, "datas")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def meta(self) -> 'outputs.GetCloudGatewayProviderAccountListMetaResult':
        return pulumi.get(self, "meta")

    @property
    @pulumi.getter(name="pageNumber")
    def page_number(self) -> Optional[int]:
        return pulumi.get(self, "page_number")

    @property
    @pulumi.getter(name="pageSize")
    def page_size(self) -> Optional[int]:
        return pulumi.get(self, "page_size")


class AwaitableGetCloudGatewayProviderAccountListResult(GetCloudGatewayProviderAccountListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCloudGatewayProviderAccountListResult(
            datas=self.datas,
            id=self.id,
            meta=self.meta,
            page_number=self.page_number,
            page_size=self.page_size)


def get_cloud_gateway_provider_account_list(page_number: Optional[int] = None,
                                            page_size: Optional[int] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCloudGatewayProviderAccountListResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['pageNumber'] = page_number
    __args__['pageSize'] = page_size
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('konnect:index/getCloudGatewayProviderAccountList:getCloudGatewayProviderAccountList', __args__, opts=opts, typ=GetCloudGatewayProviderAccountListResult).value

    return AwaitableGetCloudGatewayProviderAccountListResult(
        datas=pulumi.get(__ret__, 'datas'),
        id=pulumi.get(__ret__, 'id'),
        meta=pulumi.get(__ret__, 'meta'),
        page_number=pulumi.get(__ret__, 'page_number'),
        page_size=pulumi.get(__ret__, 'page_size'))


@_utilities.lift_output_func(get_cloud_gateway_provider_account_list)
def get_cloud_gateway_provider_account_list_output(page_number: Optional[pulumi.Input[Optional[int]]] = None,
                                                   page_size: Optional[pulumi.Input[Optional[int]]] = None,
                                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCloudGatewayProviderAccountListResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
