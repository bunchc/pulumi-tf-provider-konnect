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
    'GetGatewayCustomPluginSchemaResult',
    'AwaitableGetGatewayCustomPluginSchemaResult',
    'get_gateway_custom_plugin_schema',
    'get_gateway_custom_plugin_schema_output',
]

@pulumi.output_type
class GetGatewayCustomPluginSchemaResult:
    """
    A collection of values returned by getGatewayCustomPluginSchema.
    """
    def __init__(__self__, control_plane_id=None, created_at=None, id=None, lua_schema=None, name=None, updated_at=None):
        if control_plane_id and not isinstance(control_plane_id, str):
            raise TypeError("Expected argument 'control_plane_id' to be a str")
        pulumi.set(__self__, "control_plane_id", control_plane_id)
        if created_at and not isinstance(created_at, int):
            raise TypeError("Expected argument 'created_at' to be a int")
        pulumi.set(__self__, "created_at", created_at)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lua_schema and not isinstance(lua_schema, str):
            raise TypeError("Expected argument 'lua_schema' to be a str")
        pulumi.set(__self__, "lua_schema", lua_schema)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if updated_at and not isinstance(updated_at, int):
            raise TypeError("Expected argument 'updated_at' to be a int")
        pulumi.set(__self__, "updated_at", updated_at)

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
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="luaSchema")
    def lua_schema(self) -> str:
        return pulumi.get(self, "lua_schema")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> int:
        return pulumi.get(self, "updated_at")


class AwaitableGetGatewayCustomPluginSchemaResult(GetGatewayCustomPluginSchemaResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGatewayCustomPluginSchemaResult(
            control_plane_id=self.control_plane_id,
            created_at=self.created_at,
            id=self.id,
            lua_schema=self.lua_schema,
            name=self.name,
            updated_at=self.updated_at)


def get_gateway_custom_plugin_schema(control_plane_id: Optional[str] = None,
                                     name: Optional[str] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGatewayCustomPluginSchemaResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['controlPlaneId'] = control_plane_id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('konnect:index/getGatewayCustomPluginSchema:getGatewayCustomPluginSchema', __args__, opts=opts, typ=GetGatewayCustomPluginSchemaResult).value

    return AwaitableGetGatewayCustomPluginSchemaResult(
        control_plane_id=pulumi.get(__ret__, 'control_plane_id'),
        created_at=pulumi.get(__ret__, 'created_at'),
        id=pulumi.get(__ret__, 'id'),
        lua_schema=pulumi.get(__ret__, 'lua_schema'),
        name=pulumi.get(__ret__, 'name'),
        updated_at=pulumi.get(__ret__, 'updated_at'))


@_utilities.lift_output_func(get_gateway_custom_plugin_schema)
def get_gateway_custom_plugin_schema_output(control_plane_id: Optional[pulumi.Input[str]] = None,
                                            name: Optional[pulumi.Input[str]] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGatewayCustomPluginSchemaResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
