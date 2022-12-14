# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['BoundaryAuthMethodArgs', 'BoundaryAuthMethod']

@pulumi.input_type
class BoundaryAuthMethodArgs:
    def __init__(__self__, *,
                 scope_id: pulumi.Input[str],
                 type: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 min_login_name_length: Optional[pulumi.Input[int]] = None,
                 min_password_length: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BoundaryAuthMethod resource.
        :param pulumi.Input[str] scope_id: The scope ID.
        :param pulumi.Input[str] type: The resource type.
        :param pulumi.Input[str] description: The auth method description.
        :param pulumi.Input[int] min_login_name_length: The minimum login name length.
        :param pulumi.Input[int] min_password_length: The minimum password length.
        :param pulumi.Input[str] name: The auth method name. Defaults to the resource name.
        """
        pulumi.set(__self__, "scope_id", scope_id)
        pulumi.set(__self__, "type", type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if min_login_name_length is not None:
            warnings.warn("""Will be removed in favor of using attributes parameter""", DeprecationWarning)
            pulumi.log.warn("""min_login_name_length is deprecated: Will be removed in favor of using attributes parameter""")
        if min_login_name_length is not None:
            pulumi.set(__self__, "min_login_name_length", min_login_name_length)
        if min_password_length is not None:
            warnings.warn("""Will be removed in favor of using attributes parameter""", DeprecationWarning)
            pulumi.log.warn("""min_password_length is deprecated: Will be removed in favor of using attributes parameter""")
        if min_password_length is not None:
            pulumi.set(__self__, "min_password_length", min_password_length)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> pulumi.Input[str]:
        """
        The scope ID.
        """
        return pulumi.get(self, "scope_id")

    @scope_id.setter
    def scope_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "scope_id", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The resource type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The auth method description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="minLoginNameLength")
    def min_login_name_length(self) -> Optional[pulumi.Input[int]]:
        """
        The minimum login name length.
        """
        return pulumi.get(self, "min_login_name_length")

    @min_login_name_length.setter
    def min_login_name_length(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_login_name_length", value)

    @property
    @pulumi.getter(name="minPasswordLength")
    def min_password_length(self) -> Optional[pulumi.Input[int]]:
        """
        The minimum password length.
        """
        return pulumi.get(self, "min_password_length")

    @min_password_length.setter
    def min_password_length(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_password_length", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The auth method name. Defaults to the resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _BoundaryAuthMethodState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 min_login_name_length: Optional[pulumi.Input[int]] = None,
                 min_password_length: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BoundaryAuthMethod resources.
        :param pulumi.Input[str] description: The auth method description.
        :param pulumi.Input[int] min_login_name_length: The minimum login name length.
        :param pulumi.Input[int] min_password_length: The minimum password length.
        :param pulumi.Input[str] name: The auth method name. Defaults to the resource name.
        :param pulumi.Input[str] scope_id: The scope ID.
        :param pulumi.Input[str] type: The resource type.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if min_login_name_length is not None:
            warnings.warn("""Will be removed in favor of using attributes parameter""", DeprecationWarning)
            pulumi.log.warn("""min_login_name_length is deprecated: Will be removed in favor of using attributes parameter""")
        if min_login_name_length is not None:
            pulumi.set(__self__, "min_login_name_length", min_login_name_length)
        if min_password_length is not None:
            warnings.warn("""Will be removed in favor of using attributes parameter""", DeprecationWarning)
            pulumi.log.warn("""min_password_length is deprecated: Will be removed in favor of using attributes parameter""")
        if min_password_length is not None:
            pulumi.set(__self__, "min_password_length", min_password_length)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if scope_id is not None:
            pulumi.set(__self__, "scope_id", scope_id)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The auth method description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="minLoginNameLength")
    def min_login_name_length(self) -> Optional[pulumi.Input[int]]:
        """
        The minimum login name length.
        """
        return pulumi.get(self, "min_login_name_length")

    @min_login_name_length.setter
    def min_login_name_length(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_login_name_length", value)

    @property
    @pulumi.getter(name="minPasswordLength")
    def min_password_length(self) -> Optional[pulumi.Input[int]]:
        """
        The minimum password length.
        """
        return pulumi.get(self, "min_password_length")

    @min_password_length.setter
    def min_password_length(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_password_length", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The auth method name. Defaults to the resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> Optional[pulumi.Input[str]]:
        """
        The scope ID.
        """
        return pulumi.get(self, "scope_id")

    @scope_id.setter
    def scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope_id", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The resource type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class BoundaryAuthMethod(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 min_login_name_length: Optional[pulumi.Input[int]] = None,
                 min_password_length: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The auth method resource allows you to configure a Boundary auth_method.

        ## Example Usage

        ```python
        import pulumi
        import katasec_boundary as boundary

        org = boundary.BoundaryScope("org",
            description="My first scope!",
            scope_id="global",
            auto_create_admin_role=True,
            auto_create_default_role=True)
        password = boundary.BoundaryAuthMethod("password",
            scope_id=org.id,
            type="password")
        ```

        ## Import

        ```sh
         $ pulumi import boundary:index/boundaryAuthMethod:BoundaryAuthMethod foo <my-id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The auth method description.
        :param pulumi.Input[int] min_login_name_length: The minimum login name length.
        :param pulumi.Input[int] min_password_length: The minimum password length.
        :param pulumi.Input[str] name: The auth method name. Defaults to the resource name.
        :param pulumi.Input[str] scope_id: The scope ID.
        :param pulumi.Input[str] type: The resource type.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BoundaryAuthMethodArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The auth method resource allows you to configure a Boundary auth_method.

        ## Example Usage

        ```python
        import pulumi
        import katasec_boundary as boundary

        org = boundary.BoundaryScope("org",
            description="My first scope!",
            scope_id="global",
            auto_create_admin_role=True,
            auto_create_default_role=True)
        password = boundary.BoundaryAuthMethod("password",
            scope_id=org.id,
            type="password")
        ```

        ## Import

        ```sh
         $ pulumi import boundary:index/boundaryAuthMethod:BoundaryAuthMethod foo <my-id>
        ```

        :param str resource_name: The name of the resource.
        :param BoundaryAuthMethodArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BoundaryAuthMethodArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 min_login_name_length: Optional[pulumi.Input[int]] = None,
                 min_password_length: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BoundaryAuthMethodArgs.__new__(BoundaryAuthMethodArgs)

            __props__.__dict__["description"] = description
            if min_login_name_length is not None and not opts.urn:
                warnings.warn("""Will be removed in favor of using attributes parameter""", DeprecationWarning)
                pulumi.log.warn("""min_login_name_length is deprecated: Will be removed in favor of using attributes parameter""")
            __props__.__dict__["min_login_name_length"] = min_login_name_length
            if min_password_length is not None and not opts.urn:
                warnings.warn("""Will be removed in favor of using attributes parameter""", DeprecationWarning)
                pulumi.log.warn("""min_password_length is deprecated: Will be removed in favor of using attributes parameter""")
            __props__.__dict__["min_password_length"] = min_password_length
            __props__.__dict__["name"] = name
            if scope_id is None and not opts.urn:
                raise TypeError("Missing required property 'scope_id'")
            __props__.__dict__["scope_id"] = scope_id
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        super(BoundaryAuthMethod, __self__).__init__(
            'boundary:index/boundaryAuthMethod:BoundaryAuthMethod',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            min_login_name_length: Optional[pulumi.Input[int]] = None,
            min_password_length: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            scope_id: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'BoundaryAuthMethod':
        """
        Get an existing BoundaryAuthMethod resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The auth method description.
        :param pulumi.Input[int] min_login_name_length: The minimum login name length.
        :param pulumi.Input[int] min_password_length: The minimum password length.
        :param pulumi.Input[str] name: The auth method name. Defaults to the resource name.
        :param pulumi.Input[str] scope_id: The scope ID.
        :param pulumi.Input[str] type: The resource type.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BoundaryAuthMethodState.__new__(_BoundaryAuthMethodState)

        __props__.__dict__["description"] = description
        __props__.__dict__["min_login_name_length"] = min_login_name_length
        __props__.__dict__["min_password_length"] = min_password_length
        __props__.__dict__["name"] = name
        __props__.__dict__["scope_id"] = scope_id
        __props__.__dict__["type"] = type
        return BoundaryAuthMethod(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The auth method description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="minLoginNameLength")
    def min_login_name_length(self) -> pulumi.Output[int]:
        """
        The minimum login name length.
        """
        return pulumi.get(self, "min_login_name_length")

    @property
    @pulumi.getter(name="minPasswordLength")
    def min_password_length(self) -> pulumi.Output[int]:
        """
        The minimum password length.
        """
        return pulumi.get(self, "min_password_length")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The auth method name. Defaults to the resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> pulumi.Output[str]:
        """
        The scope ID.
        """
        return pulumi.get(self, "scope_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The resource type.
        """
        return pulumi.get(self, "type")

