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
    'GetPortalAuthResult',
    'AwaitableGetPortalAuthResult',
    'get_portal_auth',
    'get_portal_auth_output',
]

@pulumi.output_type
class GetPortalAuthResult:
    """
    A collection of values returned by getPortalAuth.
    """
    def __init__(__self__, basic_auth_enabled=None, id=None, konnect_mapping_enabled=None, oidc_auth_enabled=None, oidc_config=None, oidc_team_mapping_enabled=None, portal_id=None):
        if basic_auth_enabled and not isinstance(basic_auth_enabled, bool):
            raise TypeError("Expected argument 'basic_auth_enabled' to be a bool")
        pulumi.set(__self__, "basic_auth_enabled", basic_auth_enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if konnect_mapping_enabled and not isinstance(konnect_mapping_enabled, bool):
            raise TypeError("Expected argument 'konnect_mapping_enabled' to be a bool")
        pulumi.set(__self__, "konnect_mapping_enabled", konnect_mapping_enabled)
        if oidc_auth_enabled and not isinstance(oidc_auth_enabled, bool):
            raise TypeError("Expected argument 'oidc_auth_enabled' to be a bool")
        pulumi.set(__self__, "oidc_auth_enabled", oidc_auth_enabled)
        if oidc_config and not isinstance(oidc_config, dict):
            raise TypeError("Expected argument 'oidc_config' to be a dict")
        pulumi.set(__self__, "oidc_config", oidc_config)
        if oidc_team_mapping_enabled and not isinstance(oidc_team_mapping_enabled, bool):
            raise TypeError("Expected argument 'oidc_team_mapping_enabled' to be a bool")
        pulumi.set(__self__, "oidc_team_mapping_enabled", oidc_team_mapping_enabled)
        if portal_id and not isinstance(portal_id, str):
            raise TypeError("Expected argument 'portal_id' to be a str")
        pulumi.set(__self__, "portal_id", portal_id)

    @property
    @pulumi.getter(name="basicAuthEnabled")
    def basic_auth_enabled(self) -> bool:
        return pulumi.get(self, "basic_auth_enabled")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="konnectMappingEnabled")
    def konnect_mapping_enabled(self) -> bool:
        return pulumi.get(self, "konnect_mapping_enabled")

    @property
    @pulumi.getter(name="oidcAuthEnabled")
    def oidc_auth_enabled(self) -> bool:
        return pulumi.get(self, "oidc_auth_enabled")

    @property
    @pulumi.getter(name="oidcConfig")
    def oidc_config(self) -> 'outputs.GetPortalAuthOidcConfigResult':
        return pulumi.get(self, "oidc_config")

    @property
    @pulumi.getter(name="oidcTeamMappingEnabled")
    def oidc_team_mapping_enabled(self) -> bool:
        return pulumi.get(self, "oidc_team_mapping_enabled")

    @property
    @pulumi.getter(name="portalId")
    def portal_id(self) -> str:
        return pulumi.get(self, "portal_id")


class AwaitableGetPortalAuthResult(GetPortalAuthResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPortalAuthResult(
            basic_auth_enabled=self.basic_auth_enabled,
            id=self.id,
            konnect_mapping_enabled=self.konnect_mapping_enabled,
            oidc_auth_enabled=self.oidc_auth_enabled,
            oidc_config=self.oidc_config,
            oidc_team_mapping_enabled=self.oidc_team_mapping_enabled,
            portal_id=self.portal_id)


def get_portal_auth(portal_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPortalAuthResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['portalId'] = portal_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('konnect:index/getPortalAuth:getPortalAuth', __args__, opts=opts, typ=GetPortalAuthResult).value

    return AwaitableGetPortalAuthResult(
        basic_auth_enabled=pulumi.get(__ret__, 'basic_auth_enabled'),
        id=pulumi.get(__ret__, 'id'),
        konnect_mapping_enabled=pulumi.get(__ret__, 'konnect_mapping_enabled'),
        oidc_auth_enabled=pulumi.get(__ret__, 'oidc_auth_enabled'),
        oidc_config=pulumi.get(__ret__, 'oidc_config'),
        oidc_team_mapping_enabled=pulumi.get(__ret__, 'oidc_team_mapping_enabled'),
        portal_id=pulumi.get(__ret__, 'portal_id'))


@_utilities.lift_output_func(get_portal_auth)
def get_portal_auth_output(portal_id: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPortalAuthResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
