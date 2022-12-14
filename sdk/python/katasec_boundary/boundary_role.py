# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['BoundaryRoleArgs', 'BoundaryRole']

@pulumi.input_type
class BoundaryRoleArgs:
    def __init__(__self__, *,
                 scope_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 grant_scope_id: Optional[pulumi.Input[str]] = None,
                 grant_strings: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 principal_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a BoundaryRole resource.
        :param pulumi.Input[str] scope_id: The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
        :param pulumi.Input[str] description: The role description.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] grant_strings: A list of stringified grants for the role.
        :param pulumi.Input[str] name: The role name. Defaults to the resource name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] principal_ids: A list of principal (user or group) IDs to add as principals on the role.
        """
        pulumi.set(__self__, "scope_id", scope_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if grant_scope_id is not None:
            pulumi.set(__self__, "grant_scope_id", grant_scope_id)
        if grant_strings is not None:
            pulumi.set(__self__, "grant_strings", grant_strings)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if principal_ids is not None:
            pulumi.set(__self__, "principal_ids", principal_ids)

    @property
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> pulumi.Input[str]:
        """
        The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
        """
        return pulumi.get(self, "scope_id")

    @scope_id.setter
    def scope_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "scope_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The role description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="grantScopeId")
    def grant_scope_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "grant_scope_id")

    @grant_scope_id.setter
    def grant_scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "grant_scope_id", value)

    @property
    @pulumi.getter(name="grantStrings")
    def grant_strings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of stringified grants for the role.
        """
        return pulumi.get(self, "grant_strings")

    @grant_strings.setter
    def grant_strings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "grant_strings", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The role name. Defaults to the resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="principalIds")
    def principal_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of principal (user or group) IDs to add as principals on the role.
        """
        return pulumi.get(self, "principal_ids")

    @principal_ids.setter
    def principal_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "principal_ids", value)


@pulumi.input_type
class _BoundaryRoleState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 grant_scope_id: Optional[pulumi.Input[str]] = None,
                 grant_strings: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 principal_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BoundaryRole resources.
        :param pulumi.Input[str] description: The role description.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] grant_strings: A list of stringified grants for the role.
        :param pulumi.Input[str] name: The role name. Defaults to the resource name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] principal_ids: A list of principal (user or group) IDs to add as principals on the role.
        :param pulumi.Input[str] scope_id: The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if grant_scope_id is not None:
            pulumi.set(__self__, "grant_scope_id", grant_scope_id)
        if grant_strings is not None:
            pulumi.set(__self__, "grant_strings", grant_strings)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if principal_ids is not None:
            pulumi.set(__self__, "principal_ids", principal_ids)
        if scope_id is not None:
            pulumi.set(__self__, "scope_id", scope_id)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The role description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="grantScopeId")
    def grant_scope_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "grant_scope_id")

    @grant_scope_id.setter
    def grant_scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "grant_scope_id", value)

    @property
    @pulumi.getter(name="grantStrings")
    def grant_strings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of stringified grants for the role.
        """
        return pulumi.get(self, "grant_strings")

    @grant_strings.setter
    def grant_strings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "grant_strings", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The role name. Defaults to the resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="principalIds")
    def principal_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of principal (user or group) IDs to add as principals on the role.
        """
        return pulumi.get(self, "principal_ids")

    @principal_ids.setter
    def principal_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "principal_ids", value)

    @property
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> Optional[pulumi.Input[str]]:
        """
        The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
        """
        return pulumi.get(self, "scope_id")

    @scope_id.setter
    def scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope_id", value)


class BoundaryRole(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 grant_scope_id: Optional[pulumi.Input[str]] = None,
                 grant_strings: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 principal_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The role resource allows you to configure a Boundary role.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import katasec_boundary as boundary

        org = boundary.BoundaryScope("org",
            description="My first scope!",
            scope_id="global",
            auto_create_admin_role=True,
            auto_create_default_role=True)
        example = boundary.BoundaryRole("example",
            description="My first role!",
            scope_id=org.id)
        ```

        Usage with a user resource:

        ```python
        import pulumi
        import katasec_boundary as boundary

        org = boundary.BoundaryScope("org",
            description="My first scope!",
            scope_id="global",
            auto_create_admin_role=True,
            auto_create_default_role=True)
        foo = boundary.BoundaryUser("foo", scope_id=org.id)
        bar = boundary.BoundaryUser("bar", scope_id=org.id)
        example = boundary.BoundaryRole("example",
            description="My first role!",
            principal_ids=[
                foo.id,
                bar.id,
            ],
            scope_id=org.id)
        ```

        Usage with user and grants resource:

        ```python
        import pulumi
        import katasec_boundary as boundary

        org = boundary.BoundaryScope("org",
            description="My first scope!",
            scope_id="global",
            auto_create_admin_role=True,
            auto_create_default_role=True)
        readonly_boundary_user = boundary.BoundaryUser("readonlyBoundaryUser",
            description="A readonly user",
            scope_id=org.id)
        readonly_boundary_role = boundary.BoundaryRole("readonlyBoundaryRole",
            description="A readonly role",
            principal_ids=[readonly_boundary_user.id],
            grant_strings=["id=*;type=*;actions=read"],
            scope_id=org.id)
        ```

        Usage for a project-specific role:

        ```python
        import pulumi
        import katasec_boundary as boundary

        org = boundary.BoundaryScope("org",
            description="My first scope!",
            scope_id="global",
            auto_create_admin_role=True,
            auto_create_default_role=True)
        project = boundary.BoundaryScope("project",
            description="My first scope!",
            scope_id=org.id,
            auto_create_admin_role=True)
        readonly_boundary_user = boundary.BoundaryUser("readonlyBoundaryUser",
            description="A readonly user",
            scope_id=org.id)
        readonly_boundary_role = boundary.BoundaryRole("readonlyBoundaryRole",
            description="A readonly role",
            principal_ids=[readonly_boundary_user.id],
            grant_strings=["id=*;type=*;actions=read"],
            scope_id=project.id)
        ```

        ## Import

        ```sh
         $ pulumi import boundary:index/boundaryRole:BoundaryRole foo <my-id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The role description.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] grant_strings: A list of stringified grants for the role.
        :param pulumi.Input[str] name: The role name. Defaults to the resource name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] principal_ids: A list of principal (user or group) IDs to add as principals on the role.
        :param pulumi.Input[str] scope_id: The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BoundaryRoleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The role resource allows you to configure a Boundary role.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import katasec_boundary as boundary

        org = boundary.BoundaryScope("org",
            description="My first scope!",
            scope_id="global",
            auto_create_admin_role=True,
            auto_create_default_role=True)
        example = boundary.BoundaryRole("example",
            description="My first role!",
            scope_id=org.id)
        ```

        Usage with a user resource:

        ```python
        import pulumi
        import katasec_boundary as boundary

        org = boundary.BoundaryScope("org",
            description="My first scope!",
            scope_id="global",
            auto_create_admin_role=True,
            auto_create_default_role=True)
        foo = boundary.BoundaryUser("foo", scope_id=org.id)
        bar = boundary.BoundaryUser("bar", scope_id=org.id)
        example = boundary.BoundaryRole("example",
            description="My first role!",
            principal_ids=[
                foo.id,
                bar.id,
            ],
            scope_id=org.id)
        ```

        Usage with user and grants resource:

        ```python
        import pulumi
        import katasec_boundary as boundary

        org = boundary.BoundaryScope("org",
            description="My first scope!",
            scope_id="global",
            auto_create_admin_role=True,
            auto_create_default_role=True)
        readonly_boundary_user = boundary.BoundaryUser("readonlyBoundaryUser",
            description="A readonly user",
            scope_id=org.id)
        readonly_boundary_role = boundary.BoundaryRole("readonlyBoundaryRole",
            description="A readonly role",
            principal_ids=[readonly_boundary_user.id],
            grant_strings=["id=*;type=*;actions=read"],
            scope_id=org.id)
        ```

        Usage for a project-specific role:

        ```python
        import pulumi
        import katasec_boundary as boundary

        org = boundary.BoundaryScope("org",
            description="My first scope!",
            scope_id="global",
            auto_create_admin_role=True,
            auto_create_default_role=True)
        project = boundary.BoundaryScope("project",
            description="My first scope!",
            scope_id=org.id,
            auto_create_admin_role=True)
        readonly_boundary_user = boundary.BoundaryUser("readonlyBoundaryUser",
            description="A readonly user",
            scope_id=org.id)
        readonly_boundary_role = boundary.BoundaryRole("readonlyBoundaryRole",
            description="A readonly role",
            principal_ids=[readonly_boundary_user.id],
            grant_strings=["id=*;type=*;actions=read"],
            scope_id=project.id)
        ```

        ## Import

        ```sh
         $ pulumi import boundary:index/boundaryRole:BoundaryRole foo <my-id>
        ```

        :param str resource_name: The name of the resource.
        :param BoundaryRoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BoundaryRoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 grant_scope_id: Optional[pulumi.Input[str]] = None,
                 grant_strings: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 principal_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BoundaryRoleArgs.__new__(BoundaryRoleArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["grant_scope_id"] = grant_scope_id
            __props__.__dict__["grant_strings"] = grant_strings
            __props__.__dict__["name"] = name
            __props__.__dict__["principal_ids"] = principal_ids
            if scope_id is None and not opts.urn:
                raise TypeError("Missing required property 'scope_id'")
            __props__.__dict__["scope_id"] = scope_id
        super(BoundaryRole, __self__).__init__(
            'boundary:index/boundaryRole:BoundaryRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            grant_scope_id: Optional[pulumi.Input[str]] = None,
            grant_strings: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            principal_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            scope_id: Optional[pulumi.Input[str]] = None) -> 'BoundaryRole':
        """
        Get an existing BoundaryRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The role description.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] grant_strings: A list of stringified grants for the role.
        :param pulumi.Input[str] name: The role name. Defaults to the resource name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] principal_ids: A list of principal (user or group) IDs to add as principals on the role.
        :param pulumi.Input[str] scope_id: The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BoundaryRoleState.__new__(_BoundaryRoleState)

        __props__.__dict__["description"] = description
        __props__.__dict__["grant_scope_id"] = grant_scope_id
        __props__.__dict__["grant_strings"] = grant_strings
        __props__.__dict__["name"] = name
        __props__.__dict__["principal_ids"] = principal_ids
        __props__.__dict__["scope_id"] = scope_id
        return BoundaryRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The role description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="grantScopeId")
    def grant_scope_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "grant_scope_id")

    @property
    @pulumi.getter(name="grantStrings")
    def grant_strings(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of stringified grants for the role.
        """
        return pulumi.get(self, "grant_strings")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The role name. Defaults to the resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="principalIds")
    def principal_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of principal (user or group) IDs to add as principals on the role.
        """
        return pulumi.get(self, "principal_ids")

    @property
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> pulumi.Output[str]:
        """
        The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
        """
        return pulumi.get(self, "scope_id")

