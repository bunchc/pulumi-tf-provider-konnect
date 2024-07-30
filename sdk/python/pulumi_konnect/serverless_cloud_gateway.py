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
from ._inputs import *

__all__ = ['ServerlessCloudGatewayArgs', 'ServerlessCloudGateway']

@pulumi.input_type
class ServerlessCloudGatewayArgs:
    def __init__(__self__, *,
                 cluster_cert: pulumi.Input[str],
                 cluster_cert_key: pulumi.Input[str],
                 control_plane: pulumi.Input['ServerlessCloudGatewayControlPlaneArgs'],
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a ServerlessCloudGateway resource.
        :param pulumi.Input[str] cluster_cert: The cluster certificate (public key). Requires replacement if changed.
        :param pulumi.Input[str] cluster_cert_key: The cluster certificate key (private key). Requires replacement if changed.
        :param pulumi.Input['ServerlessCloudGatewayControlPlaneArgs'] control_plane: Requires replacement if changed.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to facilitate tagged search on serverless cloud gateways. Keys must be of length 1-63 characters, and cannot start with 'kong', 'konnect', 'mesh', 'kic', or '_'. Requires replacement if changed.
        """
        pulumi.set(__self__, "cluster_cert", cluster_cert)
        pulumi.set(__self__, "cluster_cert_key", cluster_cert_key)
        pulumi.set(__self__, "control_plane", control_plane)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)

    @property
    @pulumi.getter(name="clusterCert")
    def cluster_cert(self) -> pulumi.Input[str]:
        """
        The cluster certificate (public key). Requires replacement if changed.
        """
        return pulumi.get(self, "cluster_cert")

    @cluster_cert.setter
    def cluster_cert(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_cert", value)

    @property
    @pulumi.getter(name="clusterCertKey")
    def cluster_cert_key(self) -> pulumi.Input[str]:
        """
        The cluster certificate key (private key). Requires replacement if changed.
        """
        return pulumi.get(self, "cluster_cert_key")

    @cluster_cert_key.setter
    def cluster_cert_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_cert_key", value)

    @property
    @pulumi.getter(name="controlPlane")
    def control_plane(self) -> pulumi.Input['ServerlessCloudGatewayControlPlaneArgs']:
        """
        Requires replacement if changed.
        """
        return pulumi.get(self, "control_plane")

    @control_plane.setter
    def control_plane(self, value: pulumi.Input['ServerlessCloudGatewayControlPlaneArgs']):
        pulumi.set(self, "control_plane", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels to facilitate tagged search on serverless cloud gateways. Keys must be of length 1-63 characters, and cannot start with 'kong', 'konnect', 'mesh', 'kic', or '_'. Requires replacement if changed.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)


@pulumi.input_type
class _ServerlessCloudGatewayState:
    def __init__(__self__, *,
                 cluster_cert: Optional[pulumi.Input[str]] = None,
                 cluster_cert_key: Optional[pulumi.Input[str]] = None,
                 control_plane: Optional[pulumi.Input['ServerlessCloudGatewayControlPlaneArgs']] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 gateway_endpoint: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServerlessCloudGateway resources.
        :param pulumi.Input[str] cluster_cert: The cluster certificate (public key). Requires replacement if changed.
        :param pulumi.Input[str] cluster_cert_key: The cluster certificate key (private key). Requires replacement if changed.
        :param pulumi.Input['ServerlessCloudGatewayControlPlaneArgs'] control_plane: Requires replacement if changed.
        :param pulumi.Input[str] gateway_endpoint: Endpoint for the serverless cloud gateway.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to facilitate tagged search on serverless cloud gateways. Keys must be of length 1-63 characters, and cannot start with 'kong', 'konnect', 'mesh', 'kic', or '_'. Requires replacement if changed.
        """
        if cluster_cert is not None:
            pulumi.set(__self__, "cluster_cert", cluster_cert)
        if cluster_cert_key is not None:
            pulumi.set(__self__, "cluster_cert_key", cluster_cert_key)
        if control_plane is not None:
            pulumi.set(__self__, "control_plane", control_plane)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if gateway_endpoint is not None:
            pulumi.set(__self__, "gateway_endpoint", gateway_endpoint)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="clusterCert")
    def cluster_cert(self) -> Optional[pulumi.Input[str]]:
        """
        The cluster certificate (public key). Requires replacement if changed.
        """
        return pulumi.get(self, "cluster_cert")

    @cluster_cert.setter
    def cluster_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_cert", value)

    @property
    @pulumi.getter(name="clusterCertKey")
    def cluster_cert_key(self) -> Optional[pulumi.Input[str]]:
        """
        The cluster certificate key (private key). Requires replacement if changed.
        """
        return pulumi.get(self, "cluster_cert_key")

    @cluster_cert_key.setter
    def cluster_cert_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_cert_key", value)

    @property
    @pulumi.getter(name="controlPlane")
    def control_plane(self) -> Optional[pulumi.Input['ServerlessCloudGatewayControlPlaneArgs']]:
        """
        Requires replacement if changed.
        """
        return pulumi.get(self, "control_plane")

    @control_plane.setter
    def control_plane(self, value: Optional[pulumi.Input['ServerlessCloudGatewayControlPlaneArgs']]):
        pulumi.set(self, "control_plane", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="gatewayEndpoint")
    def gateway_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        Endpoint for the serverless cloud gateway.
        """
        return pulumi.get(self, "gateway_endpoint")

    @gateway_endpoint.setter
    def gateway_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway_endpoint", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels to facilitate tagged search on serverless cloud gateways. Keys must be of length 1-63 characters, and cannot start with 'kong', 'konnect', 'mesh', 'kic', or '_'. Requires replacement if changed.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class ServerlessCloudGateway(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_cert: Optional[pulumi.Input[str]] = None,
                 cluster_cert_key: Optional[pulumi.Input[str]] = None,
                 control_plane: Optional[pulumi.Input[pulumi.InputType['ServerlessCloudGatewayControlPlaneArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        ServerlessCloudGateway Resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_cert: The cluster certificate (public key). Requires replacement if changed.
        :param pulumi.Input[str] cluster_cert_key: The cluster certificate key (private key). Requires replacement if changed.
        :param pulumi.Input[pulumi.InputType['ServerlessCloudGatewayControlPlaneArgs']] control_plane: Requires replacement if changed.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to facilitate tagged search on serverless cloud gateways. Keys must be of length 1-63 characters, and cannot start with 'kong', 'konnect', 'mesh', 'kic', or '_'. Requires replacement if changed.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServerlessCloudGatewayArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ServerlessCloudGateway Resource

        :param str resource_name: The name of the resource.
        :param ServerlessCloudGatewayArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServerlessCloudGatewayArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_cert: Optional[pulumi.Input[str]] = None,
                 cluster_cert_key: Optional[pulumi.Input[str]] = None,
                 control_plane: Optional[pulumi.Input[pulumi.InputType['ServerlessCloudGatewayControlPlaneArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServerlessCloudGatewayArgs.__new__(ServerlessCloudGatewayArgs)

            if cluster_cert is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_cert'")
            __props__.__dict__["cluster_cert"] = cluster_cert
            if cluster_cert_key is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_cert_key'")
            __props__.__dict__["cluster_cert_key"] = cluster_cert_key
            if control_plane is None and not opts.urn:
                raise TypeError("Missing required property 'control_plane'")
            __props__.__dict__["control_plane"] = control_plane
            __props__.__dict__["labels"] = labels
            __props__.__dict__["created_at"] = None
            __props__.__dict__["gateway_endpoint"] = None
            __props__.__dict__["updated_at"] = None
        super(ServerlessCloudGateway, __self__).__init__(
            'konnect:index/serverlessCloudGateway:ServerlessCloudGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_cert: Optional[pulumi.Input[str]] = None,
            cluster_cert_key: Optional[pulumi.Input[str]] = None,
            control_plane: Optional[pulumi.Input[pulumi.InputType['ServerlessCloudGatewayControlPlaneArgs']]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            gateway_endpoint: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'ServerlessCloudGateway':
        """
        Get an existing ServerlessCloudGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_cert: The cluster certificate (public key). Requires replacement if changed.
        :param pulumi.Input[str] cluster_cert_key: The cluster certificate key (private key). Requires replacement if changed.
        :param pulumi.Input[pulumi.InputType['ServerlessCloudGatewayControlPlaneArgs']] control_plane: Requires replacement if changed.
        :param pulumi.Input[str] gateway_endpoint: Endpoint for the serverless cloud gateway.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to facilitate tagged search on serverless cloud gateways. Keys must be of length 1-63 characters, and cannot start with 'kong', 'konnect', 'mesh', 'kic', or '_'. Requires replacement if changed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServerlessCloudGatewayState.__new__(_ServerlessCloudGatewayState)

        __props__.__dict__["cluster_cert"] = cluster_cert
        __props__.__dict__["cluster_cert_key"] = cluster_cert_key
        __props__.__dict__["control_plane"] = control_plane
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["gateway_endpoint"] = gateway_endpoint
        __props__.__dict__["labels"] = labels
        __props__.__dict__["updated_at"] = updated_at
        return ServerlessCloudGateway(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterCert")
    def cluster_cert(self) -> pulumi.Output[str]:
        """
        The cluster certificate (public key). Requires replacement if changed.
        """
        return pulumi.get(self, "cluster_cert")

    @property
    @pulumi.getter(name="clusterCertKey")
    def cluster_cert_key(self) -> pulumi.Output[str]:
        """
        The cluster certificate key (private key). Requires replacement if changed.
        """
        return pulumi.get(self, "cluster_cert_key")

    @property
    @pulumi.getter(name="controlPlane")
    def control_plane(self) -> pulumi.Output['outputs.ServerlessCloudGatewayControlPlane']:
        """
        Requires replacement if changed.
        """
        return pulumi.get(self, "control_plane")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="gatewayEndpoint")
    def gateway_endpoint(self) -> pulumi.Output[str]:
        """
        Endpoint for the serverless cloud gateway.
        """
        return pulumi.get(self, "gateway_endpoint")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Labels to facilitate tagged search on serverless cloud gateways. Keys must be of length 1-63 characters, and cannot start with 'kong', 'konnect', 'mesh', 'kic', or '_'. Requires replacement if changed.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "updated_at")
