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

__all__ = ['GatewayHmacAuthArgs', 'GatewayHmacAuth']

@pulumi.input_type
class GatewayHmacAuthArgs:
    def __init__(__self__, *,
                 consumer_id: pulumi.Input[str],
                 control_plane_id: pulumi.Input[str],
                 secret: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GatewayHmacAuth resource.
        :param pulumi.Input[str] consumer_id: Consumer ID for nested entities. Requires replacement if changed.
        :param pulumi.Input[str] control_plane_id: The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.
        :param pulumi.Input[str] secret: Requires replacement if changed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Requires replacement if changed.
        :param pulumi.Input[str] username: Requires replacement if changed.
        """
        pulumi.set(__self__, "consumer_id", consumer_id)
        pulumi.set(__self__, "control_plane_id", control_plane_id)
        if secret is not None:
            pulumi.set(__self__, "secret", secret)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="consumerId")
    def consumer_id(self) -> pulumi.Input[str]:
        """
        Consumer ID for nested entities. Requires replacement if changed.
        """
        return pulumi.get(self, "consumer_id")

    @consumer_id.setter
    def consumer_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "consumer_id", value)

    @property
    @pulumi.getter(name="controlPlaneId")
    def control_plane_id(self) -> pulumi.Input[str]:
        """
        The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.
        """
        return pulumi.get(self, "control_plane_id")

    @control_plane_id.setter
    def control_plane_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "control_plane_id", value)

    @property
    @pulumi.getter
    def secret(self) -> Optional[pulumi.Input[str]]:
        """
        Requires replacement if changed.
        """
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Requires replacement if changed.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        Requires replacement if changed.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class _GatewayHmacAuthState:
    def __init__(__self__, *,
                 consumer: Optional[pulumi.Input['GatewayHmacAuthConsumerArgs']] = None,
                 consumer_id: Optional[pulumi.Input[str]] = None,
                 control_plane_id: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[int]] = None,
                 secret: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GatewayHmacAuth resources.
        :param pulumi.Input[str] consumer_id: Consumer ID for nested entities. Requires replacement if changed.
        :param pulumi.Input[str] control_plane_id: The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.
        :param pulumi.Input[int] created_at: Unix epoch when the resource was created.
        :param pulumi.Input[str] secret: Requires replacement if changed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Requires replacement if changed.
        :param pulumi.Input[str] username: Requires replacement if changed.
        """
        if consumer is not None:
            pulumi.set(__self__, "consumer", consumer)
        if consumer_id is not None:
            pulumi.set(__self__, "consumer_id", consumer_id)
        if control_plane_id is not None:
            pulumi.set(__self__, "control_plane_id", control_plane_id)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if secret is not None:
            pulumi.set(__self__, "secret", secret)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def consumer(self) -> Optional[pulumi.Input['GatewayHmacAuthConsumerArgs']]:
        return pulumi.get(self, "consumer")

    @consumer.setter
    def consumer(self, value: Optional[pulumi.Input['GatewayHmacAuthConsumerArgs']]):
        pulumi.set(self, "consumer", value)

    @property
    @pulumi.getter(name="consumerId")
    def consumer_id(self) -> Optional[pulumi.Input[str]]:
        """
        Consumer ID for nested entities. Requires replacement if changed.
        """
        return pulumi.get(self, "consumer_id")

    @consumer_id.setter
    def consumer_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "consumer_id", value)

    @property
    @pulumi.getter(name="controlPlaneId")
    def control_plane_id(self) -> Optional[pulumi.Input[str]]:
        """
        The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.
        """
        return pulumi.get(self, "control_plane_id")

    @control_plane_id.setter
    def control_plane_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "control_plane_id", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[int]]:
        """
        Unix epoch when the resource was created.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def secret(self) -> Optional[pulumi.Input[str]]:
        """
        Requires replacement if changed.
        """
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Requires replacement if changed.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        Requires replacement if changed.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class GatewayHmacAuth(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consumer_id: Optional[pulumi.Input[str]] = None,
                 control_plane_id: Optional[pulumi.Input[str]] = None,
                 secret: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        GatewayHMACAuth Resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] consumer_id: Consumer ID for nested entities. Requires replacement if changed.
        :param pulumi.Input[str] control_plane_id: The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.
        :param pulumi.Input[str] secret: Requires replacement if changed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Requires replacement if changed.
        :param pulumi.Input[str] username: Requires replacement if changed.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GatewayHmacAuthArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        GatewayHMACAuth Resource

        :param str resource_name: The name of the resource.
        :param GatewayHmacAuthArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GatewayHmacAuthArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consumer_id: Optional[pulumi.Input[str]] = None,
                 control_plane_id: Optional[pulumi.Input[str]] = None,
                 secret: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GatewayHmacAuthArgs.__new__(GatewayHmacAuthArgs)

            if consumer_id is None and not opts.urn:
                raise TypeError("Missing required property 'consumer_id'")
            __props__.__dict__["consumer_id"] = consumer_id
            if control_plane_id is None and not opts.urn:
                raise TypeError("Missing required property 'control_plane_id'")
            __props__.__dict__["control_plane_id"] = control_plane_id
            __props__.__dict__["secret"] = secret
            __props__.__dict__["tags"] = tags
            __props__.__dict__["username"] = username
            __props__.__dict__["consumer"] = None
            __props__.__dict__["created_at"] = None
        super(GatewayHmacAuth, __self__).__init__(
            'konnect:index/gatewayHmacAuth:GatewayHmacAuth',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            consumer: Optional[pulumi.Input[pulumi.InputType['GatewayHmacAuthConsumerArgs']]] = None,
            consumer_id: Optional[pulumi.Input[str]] = None,
            control_plane_id: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[int]] = None,
            secret: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'GatewayHmacAuth':
        """
        Get an existing GatewayHmacAuth resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] consumer_id: Consumer ID for nested entities. Requires replacement if changed.
        :param pulumi.Input[str] control_plane_id: The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.
        :param pulumi.Input[int] created_at: Unix epoch when the resource was created.
        :param pulumi.Input[str] secret: Requires replacement if changed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Requires replacement if changed.
        :param pulumi.Input[str] username: Requires replacement if changed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GatewayHmacAuthState.__new__(_GatewayHmacAuthState)

        __props__.__dict__["consumer"] = consumer
        __props__.__dict__["consumer_id"] = consumer_id
        __props__.__dict__["control_plane_id"] = control_plane_id
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["secret"] = secret
        __props__.__dict__["tags"] = tags
        __props__.__dict__["username"] = username
        return GatewayHmacAuth(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def consumer(self) -> pulumi.Output['outputs.GatewayHmacAuthConsumer']:
        return pulumi.get(self, "consumer")

    @property
    @pulumi.getter(name="consumerId")
    def consumer_id(self) -> pulumi.Output[str]:
        """
        Consumer ID for nested entities. Requires replacement if changed.
        """
        return pulumi.get(self, "consumer_id")

    @property
    @pulumi.getter(name="controlPlaneId")
    def control_plane_id(self) -> pulumi.Output[str]:
        """
        The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.
        """
        return pulumi.get(self, "control_plane_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[int]:
        """
        Unix epoch when the resource was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def secret(self) -> pulumi.Output[str]:
        """
        Requires replacement if changed.
        """
        return pulumi.get(self, "secret")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Sequence[str]]:
        """
        Requires replacement if changed.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        Requires replacement if changed.
        """
        return pulumi.get(self, "username")

