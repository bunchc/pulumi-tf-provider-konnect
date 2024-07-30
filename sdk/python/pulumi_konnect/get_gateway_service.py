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
    'GetGatewayServiceResult',
    'AwaitableGetGatewayServiceResult',
    'get_gateway_service',
    'get_gateway_service_output',
]

@pulumi.output_type
class GetGatewayServiceResult:
    """
    A collection of values returned by getGatewayService.
    """
    def __init__(__self__, ca_certificates=None, client_certificate=None, connect_timeout=None, control_plane_id=None, created_at=None, enabled=None, host=None, id=None, name=None, path=None, port=None, protocol=None, read_timeout=None, retries=None, tags=None, tls_verify=None, tls_verify_depth=None, updated_at=None, write_timeout=None):
        if ca_certificates and not isinstance(ca_certificates, list):
            raise TypeError("Expected argument 'ca_certificates' to be a list")
        pulumi.set(__self__, "ca_certificates", ca_certificates)
        if client_certificate and not isinstance(client_certificate, dict):
            raise TypeError("Expected argument 'client_certificate' to be a dict")
        pulumi.set(__self__, "client_certificate", client_certificate)
        if connect_timeout and not isinstance(connect_timeout, int):
            raise TypeError("Expected argument 'connect_timeout' to be a int")
        pulumi.set(__self__, "connect_timeout", connect_timeout)
        if control_plane_id and not isinstance(control_plane_id, str):
            raise TypeError("Expected argument 'control_plane_id' to be a str")
        pulumi.set(__self__, "control_plane_id", control_plane_id)
        if created_at and not isinstance(created_at, int):
            raise TypeError("Expected argument 'created_at' to be a int")
        pulumi.set(__self__, "created_at", created_at)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if host and not isinstance(host, str):
            raise TypeError("Expected argument 'host' to be a str")
        pulumi.set(__self__, "host", host)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)
        if port and not isinstance(port, int):
            raise TypeError("Expected argument 'port' to be a int")
        pulumi.set(__self__, "port", port)
        if protocol and not isinstance(protocol, str):
            raise TypeError("Expected argument 'protocol' to be a str")
        pulumi.set(__self__, "protocol", protocol)
        if read_timeout and not isinstance(read_timeout, int):
            raise TypeError("Expected argument 'read_timeout' to be a int")
        pulumi.set(__self__, "read_timeout", read_timeout)
        if retries and not isinstance(retries, int):
            raise TypeError("Expected argument 'retries' to be a int")
        pulumi.set(__self__, "retries", retries)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if tls_verify and not isinstance(tls_verify, bool):
            raise TypeError("Expected argument 'tls_verify' to be a bool")
        pulumi.set(__self__, "tls_verify", tls_verify)
        if tls_verify_depth and not isinstance(tls_verify_depth, int):
            raise TypeError("Expected argument 'tls_verify_depth' to be a int")
        pulumi.set(__self__, "tls_verify_depth", tls_verify_depth)
        if updated_at and not isinstance(updated_at, int):
            raise TypeError("Expected argument 'updated_at' to be a int")
        pulumi.set(__self__, "updated_at", updated_at)
        if write_timeout and not isinstance(write_timeout, int):
            raise TypeError("Expected argument 'write_timeout' to be a int")
        pulumi.set(__self__, "write_timeout", write_timeout)

    @property
    @pulumi.getter(name="caCertificates")
    def ca_certificates(self) -> Sequence[str]:
        return pulumi.get(self, "ca_certificates")

    @property
    @pulumi.getter(name="clientCertificate")
    def client_certificate(self) -> 'outputs.GetGatewayServiceClientCertificateResult':
        return pulumi.get(self, "client_certificate")

    @property
    @pulumi.getter(name="connectTimeout")
    def connect_timeout(self) -> int:
        return pulumi.get(self, "connect_timeout")

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
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def host(self) -> str:
        return pulumi.get(self, "host")

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
    def path(self) -> str:
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def port(self) -> int:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="readTimeout")
    def read_timeout(self) -> int:
        return pulumi.get(self, "read_timeout")

    @property
    @pulumi.getter
    def retries(self) -> int:
        return pulumi.get(self, "retries")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tlsVerify")
    def tls_verify(self) -> bool:
        return pulumi.get(self, "tls_verify")

    @property
    @pulumi.getter(name="tlsVerifyDepth")
    def tls_verify_depth(self) -> int:
        return pulumi.get(self, "tls_verify_depth")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> int:
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="writeTimeout")
    def write_timeout(self) -> int:
        return pulumi.get(self, "write_timeout")


class AwaitableGetGatewayServiceResult(GetGatewayServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGatewayServiceResult(
            ca_certificates=self.ca_certificates,
            client_certificate=self.client_certificate,
            connect_timeout=self.connect_timeout,
            control_plane_id=self.control_plane_id,
            created_at=self.created_at,
            enabled=self.enabled,
            host=self.host,
            id=self.id,
            name=self.name,
            path=self.path,
            port=self.port,
            protocol=self.protocol,
            read_timeout=self.read_timeout,
            retries=self.retries,
            tags=self.tags,
            tls_verify=self.tls_verify,
            tls_verify_depth=self.tls_verify_depth,
            updated_at=self.updated_at,
            write_timeout=self.write_timeout)


def get_gateway_service(control_plane_id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGatewayServiceResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['controlPlaneId'] = control_plane_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('konnect:index/getGatewayService:getGatewayService', __args__, opts=opts, typ=GetGatewayServiceResult).value

    return AwaitableGetGatewayServiceResult(
        ca_certificates=pulumi.get(__ret__, 'ca_certificates'),
        client_certificate=pulumi.get(__ret__, 'client_certificate'),
        connect_timeout=pulumi.get(__ret__, 'connect_timeout'),
        control_plane_id=pulumi.get(__ret__, 'control_plane_id'),
        created_at=pulumi.get(__ret__, 'created_at'),
        enabled=pulumi.get(__ret__, 'enabled'),
        host=pulumi.get(__ret__, 'host'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        path=pulumi.get(__ret__, 'path'),
        port=pulumi.get(__ret__, 'port'),
        protocol=pulumi.get(__ret__, 'protocol'),
        read_timeout=pulumi.get(__ret__, 'read_timeout'),
        retries=pulumi.get(__ret__, 'retries'),
        tags=pulumi.get(__ret__, 'tags'),
        tls_verify=pulumi.get(__ret__, 'tls_verify'),
        tls_verify_depth=pulumi.get(__ret__, 'tls_verify_depth'),
        updated_at=pulumi.get(__ret__, 'updated_at'),
        write_timeout=pulumi.get(__ret__, 'write_timeout'))


@_utilities.lift_output_func(get_gateway_service)
def get_gateway_service_output(control_plane_id: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGatewayServiceResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...