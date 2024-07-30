# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ApiProductSpecificationArgs', 'ApiProductSpecification']

@pulumi.input_type
class ApiProductSpecificationArgs:
    def __init__(__self__, *,
                 api_product_id: pulumi.Input[str],
                 api_product_version_id: pulumi.Input[str],
                 content: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ApiProductSpecification resource.
        :param pulumi.Input[str] api_product_id: The API product identifier
        :param pulumi.Input[str] api_product_version_id: The API product version identifier
        :param pulumi.Input[str] content: The base64 encoded contents of the API product version specification
        :param pulumi.Input[str] name: The name of the API product version specification
        """
        pulumi.set(__self__, "api_product_id", api_product_id)
        pulumi.set(__self__, "api_product_version_id", api_product_version_id)
        pulumi.set(__self__, "content", content)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="apiProductId")
    def api_product_id(self) -> pulumi.Input[str]:
        """
        The API product identifier
        """
        return pulumi.get(self, "api_product_id")

    @api_product_id.setter
    def api_product_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_product_id", value)

    @property
    @pulumi.getter(name="apiProductVersionId")
    def api_product_version_id(self) -> pulumi.Input[str]:
        """
        The API product version identifier
        """
        return pulumi.get(self, "api_product_version_id")

    @api_product_version_id.setter
    def api_product_version_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_product_version_id", value)

    @property
    @pulumi.getter
    def content(self) -> pulumi.Input[str]:
        """
        The base64 encoded contents of the API product version specification
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: pulumi.Input[str]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the API product version specification
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ApiProductSpecificationState:
    def __init__(__self__, *,
                 api_product_id: Optional[pulumi.Input[str]] = None,
                 api_product_version_id: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ApiProductSpecification resources.
        :param pulumi.Input[str] api_product_id: The API product identifier
        :param pulumi.Input[str] api_product_version_id: The API product version identifier
        :param pulumi.Input[str] content: The base64 encoded contents of the API product version specification
        :param pulumi.Input[str] created_at: An ISO-8601 timestamp representation of entity creation date.
        :param pulumi.Input[str] name: The name of the API product version specification
        :param pulumi.Input[str] updated_at: An ISO-8601 timestamp representation of entity update date.
        """
        if api_product_id is not None:
            pulumi.set(__self__, "api_product_id", api_product_id)
        if api_product_version_id is not None:
            pulumi.set(__self__, "api_product_version_id", api_product_version_id)
        if content is not None:
            pulumi.set(__self__, "content", content)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="apiProductId")
    def api_product_id(self) -> Optional[pulumi.Input[str]]:
        """
        The API product identifier
        """
        return pulumi.get(self, "api_product_id")

    @api_product_id.setter
    def api_product_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_product_id", value)

    @property
    @pulumi.getter(name="apiProductVersionId")
    def api_product_version_id(self) -> Optional[pulumi.Input[str]]:
        """
        The API product version identifier
        """
        return pulumi.get(self, "api_product_version_id")

    @api_product_version_id.setter
    def api_product_version_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_product_version_id", value)

    @property
    @pulumi.getter
    def content(self) -> Optional[pulumi.Input[str]]:
        """
        The base64 encoded contents of the API product version specification
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content", value)

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
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the API product version specification
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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


class ApiProductSpecification(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_product_id: Optional[pulumi.Input[str]] = None,
                 api_product_version_id: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        APIProductSpecification Resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_product_id: The API product identifier
        :param pulumi.Input[str] api_product_version_id: The API product version identifier
        :param pulumi.Input[str] content: The base64 encoded contents of the API product version specification
        :param pulumi.Input[str] name: The name of the API product version specification
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApiProductSpecificationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        APIProductSpecification Resource

        :param str resource_name: The name of the resource.
        :param ApiProductSpecificationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApiProductSpecificationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_product_id: Optional[pulumi.Input[str]] = None,
                 api_product_version_id: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApiProductSpecificationArgs.__new__(ApiProductSpecificationArgs)

            if api_product_id is None and not opts.urn:
                raise TypeError("Missing required property 'api_product_id'")
            __props__.__dict__["api_product_id"] = api_product_id
            if api_product_version_id is None and not opts.urn:
                raise TypeError("Missing required property 'api_product_version_id'")
            __props__.__dict__["api_product_version_id"] = api_product_version_id
            if content is None and not opts.urn:
                raise TypeError("Missing required property 'content'")
            __props__.__dict__["content"] = content
            __props__.__dict__["name"] = name
            __props__.__dict__["created_at"] = None
            __props__.__dict__["updated_at"] = None
        super(ApiProductSpecification, __self__).__init__(
            'konnect:index/apiProductSpecification:ApiProductSpecification',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_product_id: Optional[pulumi.Input[str]] = None,
            api_product_version_id: Optional[pulumi.Input[str]] = None,
            content: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'ApiProductSpecification':
        """
        Get an existing ApiProductSpecification resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_product_id: The API product identifier
        :param pulumi.Input[str] api_product_version_id: The API product version identifier
        :param pulumi.Input[str] content: The base64 encoded contents of the API product version specification
        :param pulumi.Input[str] created_at: An ISO-8601 timestamp representation of entity creation date.
        :param pulumi.Input[str] name: The name of the API product version specification
        :param pulumi.Input[str] updated_at: An ISO-8601 timestamp representation of entity update date.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApiProductSpecificationState.__new__(_ApiProductSpecificationState)

        __props__.__dict__["api_product_id"] = api_product_id
        __props__.__dict__["api_product_version_id"] = api_product_version_id
        __props__.__dict__["content"] = content
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["name"] = name
        __props__.__dict__["updated_at"] = updated_at
        return ApiProductSpecification(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiProductId")
    def api_product_id(self) -> pulumi.Output[str]:
        """
        The API product identifier
        """
        return pulumi.get(self, "api_product_id")

    @property
    @pulumi.getter(name="apiProductVersionId")
    def api_product_version_id(self) -> pulumi.Output[str]:
        """
        The API product version identifier
        """
        return pulumi.get(self, "api_product_version_id")

    @property
    @pulumi.getter
    def content(self) -> pulumi.Output[str]:
        """
        The base64 encoded contents of the API product version specification
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        An ISO-8601 timestamp representation of entity creation date.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the API product version specification
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        An ISO-8601 timestamp representation of entity update date.
        """
        return pulumi.get(self, "updated_at")
