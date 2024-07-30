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
    'GetPortalResult',
    'AwaitableGetPortalResult',
    'get_portal',
    'get_portal_output',
]

@pulumi.output_type
class GetPortalResult:
    """
    A collection of values returned by getPortal.
    """
    def __init__(__self__, application_count=None, auto_approve_applications=None, auto_approve_developers=None, created_at=None, custom_client_domain=None, custom_domain=None, default_application_auth_strategy_id=None, default_domain=None, description=None, developer_count=None, display_name=None, id=None, is_public=None, labels=None, name=None, published_product_count=None, rbac_enabled=None, updated_at=None):
        if application_count and not isinstance(application_count, float):
            raise TypeError("Expected argument 'application_count' to be a float")
        pulumi.set(__self__, "application_count", application_count)
        if auto_approve_applications and not isinstance(auto_approve_applications, bool):
            raise TypeError("Expected argument 'auto_approve_applications' to be a bool")
        pulumi.set(__self__, "auto_approve_applications", auto_approve_applications)
        if auto_approve_developers and not isinstance(auto_approve_developers, bool):
            raise TypeError("Expected argument 'auto_approve_developers' to be a bool")
        pulumi.set(__self__, "auto_approve_developers", auto_approve_developers)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if custom_client_domain and not isinstance(custom_client_domain, str):
            raise TypeError("Expected argument 'custom_client_domain' to be a str")
        pulumi.set(__self__, "custom_client_domain", custom_client_domain)
        if custom_domain and not isinstance(custom_domain, str):
            raise TypeError("Expected argument 'custom_domain' to be a str")
        pulumi.set(__self__, "custom_domain", custom_domain)
        if default_application_auth_strategy_id and not isinstance(default_application_auth_strategy_id, str):
            raise TypeError("Expected argument 'default_application_auth_strategy_id' to be a str")
        pulumi.set(__self__, "default_application_auth_strategy_id", default_application_auth_strategy_id)
        if default_domain and not isinstance(default_domain, str):
            raise TypeError("Expected argument 'default_domain' to be a str")
        pulumi.set(__self__, "default_domain", default_domain)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if developer_count and not isinstance(developer_count, float):
            raise TypeError("Expected argument 'developer_count' to be a float")
        pulumi.set(__self__, "developer_count", developer_count)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_public and not isinstance(is_public, bool):
            raise TypeError("Expected argument 'is_public' to be a bool")
        pulumi.set(__self__, "is_public", is_public)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if published_product_count and not isinstance(published_product_count, float):
            raise TypeError("Expected argument 'published_product_count' to be a float")
        pulumi.set(__self__, "published_product_count", published_product_count)
        if rbac_enabled and not isinstance(rbac_enabled, bool):
            raise TypeError("Expected argument 'rbac_enabled' to be a bool")
        pulumi.set(__self__, "rbac_enabled", rbac_enabled)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="applicationCount")
    def application_count(self) -> float:
        return pulumi.get(self, "application_count")

    @property
    @pulumi.getter(name="autoApproveApplications")
    def auto_approve_applications(self) -> bool:
        return pulumi.get(self, "auto_approve_applications")

    @property
    @pulumi.getter(name="autoApproveDevelopers")
    def auto_approve_developers(self) -> bool:
        return pulumi.get(self, "auto_approve_developers")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="customClientDomain")
    def custom_client_domain(self) -> str:
        return pulumi.get(self, "custom_client_domain")

    @property
    @pulumi.getter(name="customDomain")
    def custom_domain(self) -> str:
        return pulumi.get(self, "custom_domain")

    @property
    @pulumi.getter(name="defaultApplicationAuthStrategyId")
    def default_application_auth_strategy_id(self) -> str:
        return pulumi.get(self, "default_application_auth_strategy_id")

    @property
    @pulumi.getter(name="defaultDomain")
    def default_domain(self) -> str:
        return pulumi.get(self, "default_domain")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="developerCount")
    def developer_count(self) -> float:
        return pulumi.get(self, "developer_count")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isPublic")
    def is_public(self) -> bool:
        return pulumi.get(self, "is_public")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="publishedProductCount")
    def published_product_count(self) -> float:
        return pulumi.get(self, "published_product_count")

    @property
    @pulumi.getter(name="rbacEnabled")
    def rbac_enabled(self) -> bool:
        return pulumi.get(self, "rbac_enabled")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        return pulumi.get(self, "updated_at")


class AwaitableGetPortalResult(GetPortalResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPortalResult(
            application_count=self.application_count,
            auto_approve_applications=self.auto_approve_applications,
            auto_approve_developers=self.auto_approve_developers,
            created_at=self.created_at,
            custom_client_domain=self.custom_client_domain,
            custom_domain=self.custom_domain,
            default_application_auth_strategy_id=self.default_application_auth_strategy_id,
            default_domain=self.default_domain,
            description=self.description,
            developer_count=self.developer_count,
            display_name=self.display_name,
            id=self.id,
            is_public=self.is_public,
            labels=self.labels,
            name=self.name,
            published_product_count=self.published_product_count,
            rbac_enabled=self.rbac_enabled,
            updated_at=self.updated_at)


def get_portal(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPortalResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('konnect:index/getPortal:getPortal', __args__, opts=opts, typ=GetPortalResult).value

    return AwaitableGetPortalResult(
        application_count=pulumi.get(__ret__, 'application_count'),
        auto_approve_applications=pulumi.get(__ret__, 'auto_approve_applications'),
        auto_approve_developers=pulumi.get(__ret__, 'auto_approve_developers'),
        created_at=pulumi.get(__ret__, 'created_at'),
        custom_client_domain=pulumi.get(__ret__, 'custom_client_domain'),
        custom_domain=pulumi.get(__ret__, 'custom_domain'),
        default_application_auth_strategy_id=pulumi.get(__ret__, 'default_application_auth_strategy_id'),
        default_domain=pulumi.get(__ret__, 'default_domain'),
        description=pulumi.get(__ret__, 'description'),
        developer_count=pulumi.get(__ret__, 'developer_count'),
        display_name=pulumi.get(__ret__, 'display_name'),
        id=pulumi.get(__ret__, 'id'),
        is_public=pulumi.get(__ret__, 'is_public'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        published_product_count=pulumi.get(__ret__, 'published_product_count'),
        rbac_enabled=pulumi.get(__ret__, 'rbac_enabled'),
        updated_at=pulumi.get(__ret__, 'updated_at'))


@_utilities.lift_output_func(get_portal)
def get_portal_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPortalResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...