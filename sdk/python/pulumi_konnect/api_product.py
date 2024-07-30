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

__all__ = ['ApiProductArgs', 'ApiProduct']

@pulumi.input_type
class ApiProductArgs:
    def __init__(__self__, *,
                 portal_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ApiProduct resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] portal_ids: The list of portal identifiers which this API product should be published to
        :param pulumi.Input[str] description: The description of the API product.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
        :param pulumi.Input[str] name: The name of the API product.
        """
        pulumi.set(__self__, "portal_ids", portal_ids)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="portalIds")
    def portal_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The list of portal identifiers which this API product should be published to
        """
        return pulumi.get(self, "portal_ids")

    @portal_ids.setter
    def portal_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "portal_ids", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the API product.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the API product.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ApiProductState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 portal_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 portals: Optional[pulumi.Input[Sequence[pulumi.Input['ApiProductPortalArgs']]]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None,
                 version_count: Optional[pulumi.Input[float]] = None):
        """
        Input properties used for looking up and filtering ApiProduct resources.
        :param pulumi.Input[str] created_at: An ISO-8601 timestamp representation of entity creation date.
        :param pulumi.Input[str] description: The description of the API product.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
        :param pulumi.Input[str] name: The name of the API product.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] portal_ids: The list of portal identifiers which this API product should be published to
        :param pulumi.Input[Sequence[pulumi.Input['ApiProductPortalArgs']]] portals: The list of portals which this API product is published to
        :param pulumi.Input[str] updated_at: An ISO-8601 timestamp representation of entity update date.
        :param pulumi.Input[float] version_count: The number of product versions attached to this API product
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if portal_ids is not None:
            pulumi.set(__self__, "portal_ids", portal_ids)
        if portals is not None:
            pulumi.set(__self__, "portals", portals)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if version_count is not None:
            pulumi.set(__self__, "version_count", version_count)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        An ISO-8601 timestamp representation of entity creation date.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the API product.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the API product.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="portalIds")
    def portal_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of portal identifiers which this API product should be published to
        """
        return pulumi.get(self, "portal_ids")

    @portal_ids.setter
    def portal_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "portal_ids", value)

    @property
    @pulumi.getter
    def portals(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ApiProductPortalArgs']]]]:
        """
        The list of portals which this API product is published to
        """
        return pulumi.get(self, "portals")

    @portals.setter
    def portals(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ApiProductPortalArgs']]]]):
        pulumi.set(self, "portals", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        An ISO-8601 timestamp representation of entity update date.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)

    @property
    @pulumi.getter(name="versionCount")
    def version_count(self) -> Optional[pulumi.Input[float]]:
        """
        The number of product versions attached to this API product
        """
        return pulumi.get(self, "version_count")

    @version_count.setter
    def version_count(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "version_count", value)


class ApiProduct(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 portal_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        APIProduct Resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the API product.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
        :param pulumi.Input[str] name: The name of the API product.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] portal_ids: The list of portal identifiers which this API product should be published to
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApiProductArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        APIProduct Resource

        :param str resource_name: The name of the resource.
        :param ApiProductArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApiProductArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 portal_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApiProductArgs.__new__(ApiProductArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            if portal_ids is None and not opts.urn:
                raise TypeError("Missing required property 'portal_ids'")
            __props__.__dict__["portal_ids"] = portal_ids
            __props__.__dict__["created_at"] = None
            __props__.__dict__["portals"] = None
            __props__.__dict__["updated_at"] = None
            __props__.__dict__["version_count"] = None
        super(ApiProduct, __self__).__init__(
            'konnect:index/apiProduct:ApiProduct',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            portal_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            portals: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApiProductPortalArgs']]]]] = None,
            updated_at: Optional[pulumi.Input[str]] = None,
            version_count: Optional[pulumi.Input[float]] = None) -> 'ApiProduct':
        """
        Get an existing ApiProduct resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: An ISO-8601 timestamp representation of entity creation date.
        :param pulumi.Input[str] description: The description of the API product.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
        :param pulumi.Input[str] name: The name of the API product.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] portal_ids: The list of portal identifiers which this API product should be published to
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApiProductPortalArgs']]]] portals: The list of portals which this API product is published to
        :param pulumi.Input[str] updated_at: An ISO-8601 timestamp representation of entity update date.
        :param pulumi.Input[float] version_count: The number of product versions attached to this API product
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApiProductState.__new__(_ApiProductState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["portal_ids"] = portal_ids
        __props__.__dict__["portals"] = portals
        __props__.__dict__["updated_at"] = updated_at
        __props__.__dict__["version_count"] = version_count
        return ApiProduct(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        An ISO-8601 timestamp representation of entity creation date.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the API product.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the API product.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="portalIds")
    def portal_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of portal identifiers which this API product should be published to
        """
        return pulumi.get(self, "portal_ids")

    @property
    @pulumi.getter
    def portals(self) -> pulumi.Output[Sequence['outputs.ApiProductPortal']]:
        """
        The list of portals which this API product is published to
        """
        return pulumi.get(self, "portals")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        An ISO-8601 timestamp representation of entity update date.
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="versionCount")
    def version_count(self) -> pulumi.Output[float]:
        """
        The number of product versions attached to this API product
        """
        return pulumi.get(self, "version_count")

