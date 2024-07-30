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

__all__ = ['ApplicationAuthStrategyArgs', 'ApplicationAuthStrategy']

@pulumi.input_type
class ApplicationAuthStrategyArgs:
    def __init__(__self__, *,
                 key_auth: Optional[pulumi.Input['ApplicationAuthStrategyKeyAuthArgs']] = None,
                 openid_connect: Optional[pulumi.Input['ApplicationAuthStrategyOpenidConnectArgs']] = None):
        """
        The set of arguments for constructing a ApplicationAuthStrategy resource.
        :param pulumi.Input['ApplicationAuthStrategyKeyAuthArgs'] key_auth: Request for creating a Key Auth Application Auth Strategy. Requires replacement if changed.
        :param pulumi.Input['ApplicationAuthStrategyOpenidConnectArgs'] openid_connect: Payload for creating an OIDC Application Auth Strategy. Requires replacement if changed.
        """
        if key_auth is not None:
            pulumi.set(__self__, "key_auth", key_auth)
        if openid_connect is not None:
            pulumi.set(__self__, "openid_connect", openid_connect)

    @property
    @pulumi.getter(name="keyAuth")
    def key_auth(self) -> Optional[pulumi.Input['ApplicationAuthStrategyKeyAuthArgs']]:
        """
        Request for creating a Key Auth Application Auth Strategy. Requires replacement if changed.
        """
        return pulumi.get(self, "key_auth")

    @key_auth.setter
    def key_auth(self, value: Optional[pulumi.Input['ApplicationAuthStrategyKeyAuthArgs']]):
        pulumi.set(self, "key_auth", value)

    @property
    @pulumi.getter(name="openidConnect")
    def openid_connect(self) -> Optional[pulumi.Input['ApplicationAuthStrategyOpenidConnectArgs']]:
        """
        Payload for creating an OIDC Application Auth Strategy. Requires replacement if changed.
        """
        return pulumi.get(self, "openid_connect")

    @openid_connect.setter
    def openid_connect(self, value: Optional[pulumi.Input['ApplicationAuthStrategyOpenidConnectArgs']]):
        pulumi.set(self, "openid_connect", value)


@pulumi.input_type
class _ApplicationAuthStrategyState:
    def __init__(__self__, *,
                 active: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 key_auth: Optional[pulumi.Input['ApplicationAuthStrategyKeyAuthArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 openid_connect: Optional[pulumi.Input['ApplicationAuthStrategyOpenidConnectArgs']] = None):
        """
        Input properties used for looking up and filtering ApplicationAuthStrategy resources.
        :param pulumi.Input[bool] active: At least one published product version is using this auth strategy.
        :param pulumi.Input[str] display_name: The display name of the Auth strategy. This is used to identify the Auth strategy in the Portal UI.
        :param pulumi.Input['ApplicationAuthStrategyKeyAuthArgs'] key_auth: Request for creating a Key Auth Application Auth Strategy. Requires replacement if changed.
        :param pulumi.Input[str] name: The name of the auth strategy. This is used to identify the auth strategy in the Konnect UI.
        :param pulumi.Input['ApplicationAuthStrategyOpenidConnectArgs'] openid_connect: Payload for creating an OIDC Application Auth Strategy. Requires replacement if changed.
        """
        if active is not None:
            pulumi.set(__self__, "active", active)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if key_auth is not None:
            pulumi.set(__self__, "key_auth", key_auth)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if openid_connect is not None:
            pulumi.set(__self__, "openid_connect", openid_connect)

    @property
    @pulumi.getter
    def active(self) -> Optional[pulumi.Input[bool]]:
        """
        At least one published product version is using this auth strategy.
        """
        return pulumi.get(self, "active")

    @active.setter
    def active(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "active", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of the Auth strategy. This is used to identify the Auth strategy in the Portal UI.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="keyAuth")
    def key_auth(self) -> Optional[pulumi.Input['ApplicationAuthStrategyKeyAuthArgs']]:
        """
        Request for creating a Key Auth Application Auth Strategy. Requires replacement if changed.
        """
        return pulumi.get(self, "key_auth")

    @key_auth.setter
    def key_auth(self, value: Optional[pulumi.Input['ApplicationAuthStrategyKeyAuthArgs']]):
        pulumi.set(self, "key_auth", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the auth strategy. This is used to identify the auth strategy in the Konnect UI.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="openidConnect")
    def openid_connect(self) -> Optional[pulumi.Input['ApplicationAuthStrategyOpenidConnectArgs']]:
        """
        Payload for creating an OIDC Application Auth Strategy. Requires replacement if changed.
        """
        return pulumi.get(self, "openid_connect")

    @openid_connect.setter
    def openid_connect(self, value: Optional[pulumi.Input['ApplicationAuthStrategyOpenidConnectArgs']]):
        pulumi.set(self, "openid_connect", value)


class ApplicationAuthStrategy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key_auth: Optional[pulumi.Input[pulumi.InputType['ApplicationAuthStrategyKeyAuthArgs']]] = None,
                 openid_connect: Optional[pulumi.Input[pulumi.InputType['ApplicationAuthStrategyOpenidConnectArgs']]] = None,
                 __props__=None):
        """
        ApplicationAuthStrategy Resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ApplicationAuthStrategyKeyAuthArgs']] key_auth: Request for creating a Key Auth Application Auth Strategy. Requires replacement if changed.
        :param pulumi.Input[pulumi.InputType['ApplicationAuthStrategyOpenidConnectArgs']] openid_connect: Payload for creating an OIDC Application Auth Strategy. Requires replacement if changed.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ApplicationAuthStrategyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ApplicationAuthStrategy Resource

        :param str resource_name: The name of the resource.
        :param ApplicationAuthStrategyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationAuthStrategyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key_auth: Optional[pulumi.Input[pulumi.InputType['ApplicationAuthStrategyKeyAuthArgs']]] = None,
                 openid_connect: Optional[pulumi.Input[pulumi.InputType['ApplicationAuthStrategyOpenidConnectArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationAuthStrategyArgs.__new__(ApplicationAuthStrategyArgs)

            __props__.__dict__["key_auth"] = key_auth
            __props__.__dict__["openid_connect"] = openid_connect
            __props__.__dict__["active"] = None
            __props__.__dict__["display_name"] = None
            __props__.__dict__["name"] = None
        super(ApplicationAuthStrategy, __self__).__init__(
            'konnect:index/applicationAuthStrategy:ApplicationAuthStrategy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            active: Optional[pulumi.Input[bool]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            key_auth: Optional[pulumi.Input[pulumi.InputType['ApplicationAuthStrategyKeyAuthArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            openid_connect: Optional[pulumi.Input[pulumi.InputType['ApplicationAuthStrategyOpenidConnectArgs']]] = None) -> 'ApplicationAuthStrategy':
        """
        Get an existing ApplicationAuthStrategy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active: At least one published product version is using this auth strategy.
        :param pulumi.Input[str] display_name: The display name of the Auth strategy. This is used to identify the Auth strategy in the Portal UI.
        :param pulumi.Input[pulumi.InputType['ApplicationAuthStrategyKeyAuthArgs']] key_auth: Request for creating a Key Auth Application Auth Strategy. Requires replacement if changed.
        :param pulumi.Input[str] name: The name of the auth strategy. This is used to identify the auth strategy in the Konnect UI.
        :param pulumi.Input[pulumi.InputType['ApplicationAuthStrategyOpenidConnectArgs']] openid_connect: Payload for creating an OIDC Application Auth Strategy. Requires replacement if changed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationAuthStrategyState.__new__(_ApplicationAuthStrategyState)

        __props__.__dict__["active"] = active
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["key_auth"] = key_auth
        __props__.__dict__["name"] = name
        __props__.__dict__["openid_connect"] = openid_connect
        return ApplicationAuthStrategy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def active(self) -> pulumi.Output[bool]:
        """
        At least one published product version is using this auth strategy.
        """
        return pulumi.get(self, "active")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The display name of the Auth strategy. This is used to identify the Auth strategy in the Portal UI.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="keyAuth")
    def key_auth(self) -> pulumi.Output['outputs.ApplicationAuthStrategyKeyAuth']:
        """
        Request for creating a Key Auth Application Auth Strategy. Requires replacement if changed.
        """
        return pulumi.get(self, "key_auth")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the auth strategy. This is used to identify the auth strategy in the Konnect UI.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="openidConnect")
    def openid_connect(self) -> pulumi.Output['outputs.ApplicationAuthStrategyOpenidConnect']:
        """
        Payload for creating an OIDC Application Auth Strategy. Requires replacement if changed.
        """
        return pulumi.get(self, "openid_connect")

