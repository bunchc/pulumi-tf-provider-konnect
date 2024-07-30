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
    'GetGatewayRouteResult',
    'AwaitableGetGatewayRouteResult',
    'get_gateway_route',
    'get_gateway_route_output',
]

@pulumi.output_type
class GetGatewayRouteResult:
    """
    A collection of values returned by getGatewayRoute.
    """
    def __init__(__self__, control_plane_id=None, created_at=None, destinations=None, headers=None, hosts=None, https_redirect_status_code=None, id=None, methods=None, name=None, path_handling=None, paths=None, preserve_host=None, protocols=None, regex_priority=None, request_buffering=None, response_buffering=None, service=None, snis=None, sources=None, strip_path=None, tags=None, updated_at=None):
        if control_plane_id and not isinstance(control_plane_id, str):
            raise TypeError("Expected argument 'control_plane_id' to be a str")
        pulumi.set(__self__, "control_plane_id", control_plane_id)
        if created_at and not isinstance(created_at, int):
            raise TypeError("Expected argument 'created_at' to be a int")
        pulumi.set(__self__, "created_at", created_at)
        if destinations and not isinstance(destinations, list):
            raise TypeError("Expected argument 'destinations' to be a list")
        pulumi.set(__self__, "destinations", destinations)
        if headers and not isinstance(headers, dict):
            raise TypeError("Expected argument 'headers' to be a dict")
        pulumi.set(__self__, "headers", headers)
        if hosts and not isinstance(hosts, list):
            raise TypeError("Expected argument 'hosts' to be a list")
        pulumi.set(__self__, "hosts", hosts)
        if https_redirect_status_code and not isinstance(https_redirect_status_code, int):
            raise TypeError("Expected argument 'https_redirect_status_code' to be a int")
        pulumi.set(__self__, "https_redirect_status_code", https_redirect_status_code)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if methods and not isinstance(methods, list):
            raise TypeError("Expected argument 'methods' to be a list")
        pulumi.set(__self__, "methods", methods)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if path_handling and not isinstance(path_handling, str):
            raise TypeError("Expected argument 'path_handling' to be a str")
        pulumi.set(__self__, "path_handling", path_handling)
        if paths and not isinstance(paths, list):
            raise TypeError("Expected argument 'paths' to be a list")
        pulumi.set(__self__, "paths", paths)
        if preserve_host and not isinstance(preserve_host, bool):
            raise TypeError("Expected argument 'preserve_host' to be a bool")
        pulumi.set(__self__, "preserve_host", preserve_host)
        if protocols and not isinstance(protocols, list):
            raise TypeError("Expected argument 'protocols' to be a list")
        pulumi.set(__self__, "protocols", protocols)
        if regex_priority and not isinstance(regex_priority, int):
            raise TypeError("Expected argument 'regex_priority' to be a int")
        pulumi.set(__self__, "regex_priority", regex_priority)
        if request_buffering and not isinstance(request_buffering, bool):
            raise TypeError("Expected argument 'request_buffering' to be a bool")
        pulumi.set(__self__, "request_buffering", request_buffering)
        if response_buffering and not isinstance(response_buffering, bool):
            raise TypeError("Expected argument 'response_buffering' to be a bool")
        pulumi.set(__self__, "response_buffering", response_buffering)
        if service and not isinstance(service, dict):
            raise TypeError("Expected argument 'service' to be a dict")
        pulumi.set(__self__, "service", service)
        if snis and not isinstance(snis, list):
            raise TypeError("Expected argument 'snis' to be a list")
        pulumi.set(__self__, "snis", snis)
        if sources and not isinstance(sources, list):
            raise TypeError("Expected argument 'sources' to be a list")
        pulumi.set(__self__, "sources", sources)
        if strip_path and not isinstance(strip_path, bool):
            raise TypeError("Expected argument 'strip_path' to be a bool")
        pulumi.set(__self__, "strip_path", strip_path)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
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
    def destinations(self) -> Sequence['outputs.GetGatewayRouteDestinationResult']:
        return pulumi.get(self, "destinations")

    @property
    @pulumi.getter
    def headers(self) -> Mapping[str, str]:
        return pulumi.get(self, "headers")

    @property
    @pulumi.getter
    def hosts(self) -> Sequence[str]:
        return pulumi.get(self, "hosts")

    @property
    @pulumi.getter(name="httpsRedirectStatusCode")
    def https_redirect_status_code(self) -> int:
        return pulumi.get(self, "https_redirect_status_code")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def methods(self) -> Sequence[str]:
        return pulumi.get(self, "methods")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pathHandling")
    def path_handling(self) -> str:
        return pulumi.get(self, "path_handling")

    @property
    @pulumi.getter
    def paths(self) -> Sequence[str]:
        return pulumi.get(self, "paths")

    @property
    @pulumi.getter(name="preserveHost")
    def preserve_host(self) -> bool:
        return pulumi.get(self, "preserve_host")

    @property
    @pulumi.getter
    def protocols(self) -> Sequence[str]:
        return pulumi.get(self, "protocols")

    @property
    @pulumi.getter(name="regexPriority")
    def regex_priority(self) -> int:
        return pulumi.get(self, "regex_priority")

    @property
    @pulumi.getter(name="requestBuffering")
    def request_buffering(self) -> bool:
        return pulumi.get(self, "request_buffering")

    @property
    @pulumi.getter(name="responseBuffering")
    def response_buffering(self) -> bool:
        return pulumi.get(self, "response_buffering")

    @property
    @pulumi.getter
    def service(self) -> 'outputs.GetGatewayRouteServiceResult':
        return pulumi.get(self, "service")

    @property
    @pulumi.getter
    def snis(self) -> Sequence[str]:
        return pulumi.get(self, "snis")

    @property
    @pulumi.getter
    def sources(self) -> Sequence['outputs.GetGatewayRouteSourceResult']:
        return pulumi.get(self, "sources")

    @property
    @pulumi.getter(name="stripPath")
    def strip_path(self) -> bool:
        return pulumi.get(self, "strip_path")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> int:
        return pulumi.get(self, "updated_at")


class AwaitableGetGatewayRouteResult(GetGatewayRouteResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGatewayRouteResult(
            control_plane_id=self.control_plane_id,
            created_at=self.created_at,
            destinations=self.destinations,
            headers=self.headers,
            hosts=self.hosts,
            https_redirect_status_code=self.https_redirect_status_code,
            id=self.id,
            methods=self.methods,
            name=self.name,
            path_handling=self.path_handling,
            paths=self.paths,
            preserve_host=self.preserve_host,
            protocols=self.protocols,
            regex_priority=self.regex_priority,
            request_buffering=self.request_buffering,
            response_buffering=self.response_buffering,
            service=self.service,
            snis=self.snis,
            sources=self.sources,
            strip_path=self.strip_path,
            tags=self.tags,
            updated_at=self.updated_at)


def get_gateway_route(control_plane_id: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGatewayRouteResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['controlPlaneId'] = control_plane_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('konnect:index/getGatewayRoute:getGatewayRoute', __args__, opts=opts, typ=GetGatewayRouteResult).value

    return AwaitableGetGatewayRouteResult(
        control_plane_id=pulumi.get(__ret__, 'control_plane_id'),
        created_at=pulumi.get(__ret__, 'created_at'),
        destinations=pulumi.get(__ret__, 'destinations'),
        headers=pulumi.get(__ret__, 'headers'),
        hosts=pulumi.get(__ret__, 'hosts'),
        https_redirect_status_code=pulumi.get(__ret__, 'https_redirect_status_code'),
        id=pulumi.get(__ret__, 'id'),
        methods=pulumi.get(__ret__, 'methods'),
        name=pulumi.get(__ret__, 'name'),
        path_handling=pulumi.get(__ret__, 'path_handling'),
        paths=pulumi.get(__ret__, 'paths'),
        preserve_host=pulumi.get(__ret__, 'preserve_host'),
        protocols=pulumi.get(__ret__, 'protocols'),
        regex_priority=pulumi.get(__ret__, 'regex_priority'),
        request_buffering=pulumi.get(__ret__, 'request_buffering'),
        response_buffering=pulumi.get(__ret__, 'response_buffering'),
        service=pulumi.get(__ret__, 'service'),
        snis=pulumi.get(__ret__, 'snis'),
        sources=pulumi.get(__ret__, 'sources'),
        strip_path=pulumi.get(__ret__, 'strip_path'),
        tags=pulumi.get(__ret__, 'tags'),
        updated_at=pulumi.get(__ret__, 'updated_at'))


@_utilities.lift_output_func(get_gateway_route)
def get_gateway_route_output(control_plane_id: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGatewayRouteResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
