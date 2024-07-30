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

__all__ = ['GatewayPluginAiProxyArgs', 'GatewayPluginAiProxy']

@pulumi.input_type
class GatewayPluginAiProxyArgs:
    def __init__(__self__, *,
                 control_plane_id: pulumi.Input[str],
                 config: Optional[pulumi.Input['GatewayPluginAiProxyConfigArgs']] = None,
                 consumer: Optional[pulumi.Input['GatewayPluginAiProxyConsumerArgs']] = None,
                 consumer_group: Optional[pulumi.Input['GatewayPluginAiProxyConsumerGroupArgs']] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 protocols: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 route: Optional[pulumi.Input['GatewayPluginAiProxyRouteArgs']] = None,
                 service: Optional[pulumi.Input['GatewayPluginAiProxyServiceArgs']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a GatewayPluginAiProxy resource.
        :param pulumi.Input[str] control_plane_id: The UUID of your control plane. This variable is available in the Konnect manager.
        :param pulumi.Input['GatewayPluginAiProxyConsumerArgs'] consumer: If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
        :param pulumi.Input[bool] enabled: Whether the plugin is applied.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] protocols: A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
        :param pulumi.Input['GatewayPluginAiProxyRouteArgs'] route: If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
        :param pulumi.Input['GatewayPluginAiProxyServiceArgs'] service: If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: An optional set of strings associated with the Plugin for grouping and filtering.
        """
        pulumi.set(__self__, "control_plane_id", control_plane_id)
        if config is not None:
            pulumi.set(__self__, "config", config)
        if consumer is not None:
            pulumi.set(__self__, "consumer", consumer)
        if consumer_group is not None:
            pulumi.set(__self__, "consumer_group", consumer_group)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if instance_name is not None:
            pulumi.set(__self__, "instance_name", instance_name)
        if protocols is not None:
            pulumi.set(__self__, "protocols", protocols)
        if route is not None:
            pulumi.set(__self__, "route", route)
        if service is not None:
            pulumi.set(__self__, "service", service)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="controlPlaneId")
    def control_plane_id(self) -> pulumi.Input[str]:
        """
        The UUID of your control plane. This variable is available in the Konnect manager.
        """
        return pulumi.get(self, "control_plane_id")

    @control_plane_id.setter
    def control_plane_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "control_plane_id", value)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input['GatewayPluginAiProxyConfigArgs']]:
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input['GatewayPluginAiProxyConfigArgs']]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter
    def consumer(self) -> Optional[pulumi.Input['GatewayPluginAiProxyConsumerArgs']]:
        """
        If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
        """
        return pulumi.get(self, "consumer")

    @consumer.setter
    def consumer(self, value: Optional[pulumi.Input['GatewayPluginAiProxyConsumerArgs']]):
        pulumi.set(self, "consumer", value)

    @property
    @pulumi.getter(name="consumerGroup")
    def consumer_group(self) -> Optional[pulumi.Input['GatewayPluginAiProxyConsumerGroupArgs']]:
        return pulumi.get(self, "consumer_group")

    @consumer_group.setter
    def consumer_group(self, value: Optional[pulumi.Input['GatewayPluginAiProxyConsumerGroupArgs']]):
        pulumi.set(self, "consumer_group", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the plugin is applied.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "instance_name")

    @instance_name.setter
    def instance_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_name", value)

    @property
    @pulumi.getter
    def protocols(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
        """
        return pulumi.get(self, "protocols")

    @protocols.setter
    def protocols(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "protocols", value)

    @property
    @pulumi.getter
    def route(self) -> Optional[pulumi.Input['GatewayPluginAiProxyRouteArgs']]:
        """
        If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
        """
        return pulumi.get(self, "route")

    @route.setter
    def route(self, value: Optional[pulumi.Input['GatewayPluginAiProxyRouteArgs']]):
        pulumi.set(self, "route", value)

    @property
    @pulumi.getter
    def service(self) -> Optional[pulumi.Input['GatewayPluginAiProxyServiceArgs']]:
        """
        If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
        """
        return pulumi.get(self, "service")

    @service.setter
    def service(self, value: Optional[pulumi.Input['GatewayPluginAiProxyServiceArgs']]):
        pulumi.set(self, "service", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An optional set of strings associated with the Plugin for grouping and filtering.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _GatewayPluginAiProxyState:
    def __init__(__self__, *,
                 config: Optional[pulumi.Input['GatewayPluginAiProxyConfigArgs']] = None,
                 consumer: Optional[pulumi.Input['GatewayPluginAiProxyConsumerArgs']] = None,
                 consumer_group: Optional[pulumi.Input['GatewayPluginAiProxyConsumerGroupArgs']] = None,
                 control_plane_id: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[int]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 protocols: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 route: Optional[pulumi.Input['GatewayPluginAiProxyRouteArgs']] = None,
                 service: Optional[pulumi.Input['GatewayPluginAiProxyServiceArgs']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 updated_at: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering GatewayPluginAiProxy resources.
        :param pulumi.Input['GatewayPluginAiProxyConsumerArgs'] consumer: If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
        :param pulumi.Input[str] control_plane_id: The UUID of your control plane. This variable is available in the Konnect manager.
        :param pulumi.Input[int] created_at: Unix epoch when the resource was created.
        :param pulumi.Input[bool] enabled: Whether the plugin is applied.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] protocols: A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
        :param pulumi.Input['GatewayPluginAiProxyRouteArgs'] route: If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
        :param pulumi.Input['GatewayPluginAiProxyServiceArgs'] service: If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: An optional set of strings associated with the Plugin for grouping and filtering.
        :param pulumi.Input[int] updated_at: Unix epoch when the resource was last updated.
        """
        if config is not None:
            pulumi.set(__self__, "config", config)
        if consumer is not None:
            pulumi.set(__self__, "consumer", consumer)
        if consumer_group is not None:
            pulumi.set(__self__, "consumer_group", consumer_group)
        if control_plane_id is not None:
            pulumi.set(__self__, "control_plane_id", control_plane_id)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if instance_name is not None:
            pulumi.set(__self__, "instance_name", instance_name)
        if protocols is not None:
            pulumi.set(__self__, "protocols", protocols)
        if route is not None:
            pulumi.set(__self__, "route", route)
        if service is not None:
            pulumi.set(__self__, "service", service)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input['GatewayPluginAiProxyConfigArgs']]:
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input['GatewayPluginAiProxyConfigArgs']]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter
    def consumer(self) -> Optional[pulumi.Input['GatewayPluginAiProxyConsumerArgs']]:
        """
        If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
        """
        return pulumi.get(self, "consumer")

    @consumer.setter
    def consumer(self, value: Optional[pulumi.Input['GatewayPluginAiProxyConsumerArgs']]):
        pulumi.set(self, "consumer", value)

    @property
    @pulumi.getter(name="consumerGroup")
    def consumer_group(self) -> Optional[pulumi.Input['GatewayPluginAiProxyConsumerGroupArgs']]:
        return pulumi.get(self, "consumer_group")

    @consumer_group.setter
    def consumer_group(self, value: Optional[pulumi.Input['GatewayPluginAiProxyConsumerGroupArgs']]):
        pulumi.set(self, "consumer_group", value)

    @property
    @pulumi.getter(name="controlPlaneId")
    def control_plane_id(self) -> Optional[pulumi.Input[str]]:
        """
        The UUID of your control plane. This variable is available in the Konnect manager.
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
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the plugin is applied.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "instance_name")

    @instance_name.setter
    def instance_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_name", value)

    @property
    @pulumi.getter
    def protocols(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
        """
        return pulumi.get(self, "protocols")

    @protocols.setter
    def protocols(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "protocols", value)

    @property
    @pulumi.getter
    def route(self) -> Optional[pulumi.Input['GatewayPluginAiProxyRouteArgs']]:
        """
        If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
        """
        return pulumi.get(self, "route")

    @route.setter
    def route(self, value: Optional[pulumi.Input['GatewayPluginAiProxyRouteArgs']]):
        pulumi.set(self, "route", value)

    @property
    @pulumi.getter
    def service(self) -> Optional[pulumi.Input['GatewayPluginAiProxyServiceArgs']]:
        """
        If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
        """
        return pulumi.get(self, "service")

    @service.setter
    def service(self, value: Optional[pulumi.Input['GatewayPluginAiProxyServiceArgs']]):
        pulumi.set(self, "service", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An optional set of strings associated with the Plugin for grouping and filtering.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[int]]:
        """
        Unix epoch when the resource was last updated.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "updated_at", value)


class GatewayPluginAiProxy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[pulumi.InputType['GatewayPluginAiProxyConfigArgs']]] = None,
                 consumer: Optional[pulumi.Input[pulumi.InputType['GatewayPluginAiProxyConsumerArgs']]] = None,
                 consumer_group: Optional[pulumi.Input[pulumi.InputType['GatewayPluginAiProxyConsumerGroupArgs']]] = None,
                 control_plane_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 protocols: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 route: Optional[pulumi.Input[pulumi.InputType['GatewayPluginAiProxyRouteArgs']]] = None,
                 service: Optional[pulumi.Input[pulumi.InputType['GatewayPluginAiProxyServiceArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        GatewayPluginAIProxy Resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['GatewayPluginAiProxyConsumerArgs']] consumer: If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
        :param pulumi.Input[str] control_plane_id: The UUID of your control plane. This variable is available in the Konnect manager.
        :param pulumi.Input[bool] enabled: Whether the plugin is applied.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] protocols: A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
        :param pulumi.Input[pulumi.InputType['GatewayPluginAiProxyRouteArgs']] route: If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
        :param pulumi.Input[pulumi.InputType['GatewayPluginAiProxyServiceArgs']] service: If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: An optional set of strings associated with the Plugin for grouping and filtering.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GatewayPluginAiProxyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        GatewayPluginAIProxy Resource

        :param str resource_name: The name of the resource.
        :param GatewayPluginAiProxyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GatewayPluginAiProxyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[pulumi.InputType['GatewayPluginAiProxyConfigArgs']]] = None,
                 consumer: Optional[pulumi.Input[pulumi.InputType['GatewayPluginAiProxyConsumerArgs']]] = None,
                 consumer_group: Optional[pulumi.Input[pulumi.InputType['GatewayPluginAiProxyConsumerGroupArgs']]] = None,
                 control_plane_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 protocols: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 route: Optional[pulumi.Input[pulumi.InputType['GatewayPluginAiProxyRouteArgs']]] = None,
                 service: Optional[pulumi.Input[pulumi.InputType['GatewayPluginAiProxyServiceArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GatewayPluginAiProxyArgs.__new__(GatewayPluginAiProxyArgs)

            __props__.__dict__["config"] = config
            __props__.__dict__["consumer"] = consumer
            __props__.__dict__["consumer_group"] = consumer_group
            if control_plane_id is None and not opts.urn:
                raise TypeError("Missing required property 'control_plane_id'")
            __props__.__dict__["control_plane_id"] = control_plane_id
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["instance_name"] = instance_name
            __props__.__dict__["protocols"] = protocols
            __props__.__dict__["route"] = route
            __props__.__dict__["service"] = service
            __props__.__dict__["tags"] = tags
            __props__.__dict__["created_at"] = None
            __props__.__dict__["updated_at"] = None
        super(GatewayPluginAiProxy, __self__).__init__(
            'konnect:index/gatewayPluginAiProxy:GatewayPluginAiProxy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config: Optional[pulumi.Input[pulumi.InputType['GatewayPluginAiProxyConfigArgs']]] = None,
            consumer: Optional[pulumi.Input[pulumi.InputType['GatewayPluginAiProxyConsumerArgs']]] = None,
            consumer_group: Optional[pulumi.Input[pulumi.InputType['GatewayPluginAiProxyConsumerGroupArgs']]] = None,
            control_plane_id: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[int]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            instance_name: Optional[pulumi.Input[str]] = None,
            protocols: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            route: Optional[pulumi.Input[pulumi.InputType['GatewayPluginAiProxyRouteArgs']]] = None,
            service: Optional[pulumi.Input[pulumi.InputType['GatewayPluginAiProxyServiceArgs']]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            updated_at: Optional[pulumi.Input[int]] = None) -> 'GatewayPluginAiProxy':
        """
        Get an existing GatewayPluginAiProxy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['GatewayPluginAiProxyConsumerArgs']] consumer: If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
        :param pulumi.Input[str] control_plane_id: The UUID of your control plane. This variable is available in the Konnect manager.
        :param pulumi.Input[int] created_at: Unix epoch when the resource was created.
        :param pulumi.Input[bool] enabled: Whether the plugin is applied.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] protocols: A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
        :param pulumi.Input[pulumi.InputType['GatewayPluginAiProxyRouteArgs']] route: If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
        :param pulumi.Input[pulumi.InputType['GatewayPluginAiProxyServiceArgs']] service: If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: An optional set of strings associated with the Plugin for grouping and filtering.
        :param pulumi.Input[int] updated_at: Unix epoch when the resource was last updated.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GatewayPluginAiProxyState.__new__(_GatewayPluginAiProxyState)

        __props__.__dict__["config"] = config
        __props__.__dict__["consumer"] = consumer
        __props__.__dict__["consumer_group"] = consumer_group
        __props__.__dict__["control_plane_id"] = control_plane_id
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["instance_name"] = instance_name
        __props__.__dict__["protocols"] = protocols
        __props__.__dict__["route"] = route
        __props__.__dict__["service"] = service
        __props__.__dict__["tags"] = tags
        __props__.__dict__["updated_at"] = updated_at
        return GatewayPluginAiProxy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output['outputs.GatewayPluginAiProxyConfig']:
        return pulumi.get(self, "config")

    @property
    @pulumi.getter
    def consumer(self) -> pulumi.Output['outputs.GatewayPluginAiProxyConsumer']:
        """
        If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
        """
        return pulumi.get(self, "consumer")

    @property
    @pulumi.getter(name="consumerGroup")
    def consumer_group(self) -> pulumi.Output['outputs.GatewayPluginAiProxyConsumerGroup']:
        return pulumi.get(self, "consumer_group")

    @property
    @pulumi.getter(name="controlPlaneId")
    def control_plane_id(self) -> pulumi.Output[str]:
        """
        The UUID of your control plane. This variable is available in the Konnect manager.
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
    def enabled(self) -> pulumi.Output[bool]:
        """
        Whether the plugin is applied.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter
    def protocols(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
        """
        return pulumi.get(self, "protocols")

    @property
    @pulumi.getter
    def route(self) -> pulumi.Output['outputs.GatewayPluginAiProxyRoute']:
        """
        If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
        """
        return pulumi.get(self, "route")

    @property
    @pulumi.getter
    def service(self) -> pulumi.Output['outputs.GatewayPluginAiProxyService']:
        """
        If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
        """
        return pulumi.get(self, "service")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Sequence[str]]:
        """
        An optional set of strings associated with the Plugin for grouping and filtering.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[int]:
        """
        Unix epoch when the resource was last updated.
        """
        return pulumi.get(self, "updated_at")

