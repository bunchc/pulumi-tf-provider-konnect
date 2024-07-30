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
    'GetSystemAccountAccessTokenResult',
    'AwaitableGetSystemAccountAccessTokenResult',
    'get_system_account_access_token',
    'get_system_account_access_token_output',
]

@pulumi.output_type
class GetSystemAccountAccessTokenResult:
    """
    A collection of values returned by getSystemAccountAccessToken.
    """
    def __init__(__self__, account_id=None, created_at=None, expires_at=None, id=None, last_used_at=None, name=None, updated_at=None):
        if account_id and not isinstance(account_id, str):
            raise TypeError("Expected argument 'account_id' to be a str")
        pulumi.set(__self__, "account_id", account_id)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if expires_at and not isinstance(expires_at, str):
            raise TypeError("Expected argument 'expires_at' to be a str")
        pulumi.set(__self__, "expires_at", expires_at)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if last_used_at and not isinstance(last_used_at, str):
            raise TypeError("Expected argument 'last_used_at' to be a str")
        pulumi.set(__self__, "last_used_at", last_used_at)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> str:
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> str:
        return pulumi.get(self, "expires_at")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lastUsedAt")
    def last_used_at(self) -> str:
        return pulumi.get(self, "last_used_at")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        return pulumi.get(self, "updated_at")


class AwaitableGetSystemAccountAccessTokenResult(GetSystemAccountAccessTokenResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemAccountAccessTokenResult(
            account_id=self.account_id,
            created_at=self.created_at,
            expires_at=self.expires_at,
            id=self.id,
            last_used_at=self.last_used_at,
            name=self.name,
            updated_at=self.updated_at)


def get_system_account_access_token(account_id: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemAccountAccessTokenResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['accountId'] = account_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('konnect:index/getSystemAccountAccessToken:getSystemAccountAccessToken', __args__, opts=opts, typ=GetSystemAccountAccessTokenResult).value

    return AwaitableGetSystemAccountAccessTokenResult(
        account_id=pulumi.get(__ret__, 'account_id'),
        created_at=pulumi.get(__ret__, 'created_at'),
        expires_at=pulumi.get(__ret__, 'expires_at'),
        id=pulumi.get(__ret__, 'id'),
        last_used_at=pulumi.get(__ret__, 'last_used_at'),
        name=pulumi.get(__ret__, 'name'),
        updated_at=pulumi.get(__ret__, 'updated_at'))


@_utilities.lift_output_func(get_system_account_access_token)
def get_system_account_access_token_output(account_id: Optional[pulumi.Input[str]] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSystemAccountAccessTokenResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
