# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['BoundaryHostCatalogPluginArgs', 'BoundaryHostCatalogPlugin']

@pulumi.input_type
class BoundaryHostCatalogPluginArgs:
    def __init__(__self__, *,
                 scope_id: pulumi.Input[str],
                 attributes_json: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 internal_force_update: Optional[pulumi.Input[str]] = None,
                 internal_hmac_used_for_secrets_config_hmac: Optional[pulumi.Input[str]] = None,
                 internal_secrets_config_hmac: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 plugin_id: Optional[pulumi.Input[str]] = None,
                 plugin_name: Optional[pulumi.Input[str]] = None,
                 secrets_hmac: Optional[pulumi.Input[str]] = None,
                 secrets_json: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BoundaryHostCatalogPlugin resource.
        :param pulumi.Input[str] scope_id: The scope ID in which the resource is created.
        :param pulumi.Input[str] attributes_json: The attributes for the host catalog. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host catalog.
        :param pulumi.Input[str] description: The host catalog description.
        :param pulumi.Input[str] internal_force_update: Internal only. Used to force update so that we can always check the value of secrets.
        :param pulumi.Input[str] internal_hmac_used_for_secrets_config_hmac: Internal only. The Boundary-provided HMAC used to calculate the current value of the HMAC'd config. Used for drift detection.
        :param pulumi.Input[str] internal_secrets_config_hmac: Internal only. HMAC of (serverSecretsHmac + config secrets). Used for proper secrets handling.
        :param pulumi.Input[str] name: The host catalog name. Defaults to the resource name.
        :param pulumi.Input[str] plugin_id: The ID of the plugin that should back the resource. This or plugin_name must be defined.
        :param pulumi.Input[str] plugin_name: The name of the plugin that should back the resource. This or plugin_id must be defined.
        :param pulumi.Input[str] secrets_hmac: The HMAC'd secrets value returned from the server.
        :param pulumi.Input[str] secrets_json: The secrets for the host catalog. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" to clear any existing values. NOTE: Unlike "attributes_json", removing this block will NOT clear secrets from the host catalog; this allows injecting secrets for one call, then removing them for storage.
        """
        pulumi.set(__self__, "scope_id", scope_id)
        if attributes_json is not None:
            pulumi.set(__self__, "attributes_json", attributes_json)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if internal_force_update is not None:
            pulumi.set(__self__, "internal_force_update", internal_force_update)
        if internal_hmac_used_for_secrets_config_hmac is not None:
            pulumi.set(__self__, "internal_hmac_used_for_secrets_config_hmac", internal_hmac_used_for_secrets_config_hmac)
        if internal_secrets_config_hmac is not None:
            pulumi.set(__self__, "internal_secrets_config_hmac", internal_secrets_config_hmac)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if plugin_id is not None:
            pulumi.set(__self__, "plugin_id", plugin_id)
        if plugin_name is not None:
            pulumi.set(__self__, "plugin_name", plugin_name)
        if secrets_hmac is not None:
            pulumi.set(__self__, "secrets_hmac", secrets_hmac)
        if secrets_json is not None:
            pulumi.set(__self__, "secrets_json", secrets_json)

    @property
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> pulumi.Input[str]:
        """
        The scope ID in which the resource is created.
        """
        return pulumi.get(self, "scope_id")

    @scope_id.setter
    def scope_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "scope_id", value)

    @property
    @pulumi.getter(name="attributesJson")
    def attributes_json(self) -> Optional[pulumi.Input[str]]:
        """
        The attributes for the host catalog. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host catalog.
        """
        return pulumi.get(self, "attributes_json")

    @attributes_json.setter
    def attributes_json(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "attributes_json", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The host catalog description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="internalForceUpdate")
    def internal_force_update(self) -> Optional[pulumi.Input[str]]:
        """
        Internal only. Used to force update so that we can always check the value of secrets.
        """
        return pulumi.get(self, "internal_force_update")

    @internal_force_update.setter
    def internal_force_update(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internal_force_update", value)

    @property
    @pulumi.getter(name="internalHmacUsedForSecretsConfigHmac")
    def internal_hmac_used_for_secrets_config_hmac(self) -> Optional[pulumi.Input[str]]:
        """
        Internal only. The Boundary-provided HMAC used to calculate the current value of the HMAC'd config. Used for drift detection.
        """
        return pulumi.get(self, "internal_hmac_used_for_secrets_config_hmac")

    @internal_hmac_used_for_secrets_config_hmac.setter
    def internal_hmac_used_for_secrets_config_hmac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internal_hmac_used_for_secrets_config_hmac", value)

    @property
    @pulumi.getter(name="internalSecretsConfigHmac")
    def internal_secrets_config_hmac(self) -> Optional[pulumi.Input[str]]:
        """
        Internal only. HMAC of (serverSecretsHmac + config secrets). Used for proper secrets handling.
        """
        return pulumi.get(self, "internal_secrets_config_hmac")

    @internal_secrets_config_hmac.setter
    def internal_secrets_config_hmac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internal_secrets_config_hmac", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The host catalog name. Defaults to the resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="pluginId")
    def plugin_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the plugin that should back the resource. This or plugin_name must be defined.
        """
        return pulumi.get(self, "plugin_id")

    @plugin_id.setter
    def plugin_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plugin_id", value)

    @property
    @pulumi.getter(name="pluginName")
    def plugin_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the plugin that should back the resource. This or plugin_id must be defined.
        """
        return pulumi.get(self, "plugin_name")

    @plugin_name.setter
    def plugin_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plugin_name", value)

    @property
    @pulumi.getter(name="secretsHmac")
    def secrets_hmac(self) -> Optional[pulumi.Input[str]]:
        """
        The HMAC'd secrets value returned from the server.
        """
        return pulumi.get(self, "secrets_hmac")

    @secrets_hmac.setter
    def secrets_hmac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secrets_hmac", value)

    @property
    @pulumi.getter(name="secretsJson")
    def secrets_json(self) -> Optional[pulumi.Input[str]]:
        """
        The secrets for the host catalog. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" to clear any existing values. NOTE: Unlike "attributes_json", removing this block will NOT clear secrets from the host catalog; this allows injecting secrets for one call, then removing them for storage.
        """
        return pulumi.get(self, "secrets_json")

    @secrets_json.setter
    def secrets_json(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secrets_json", value)


@pulumi.input_type
class _BoundaryHostCatalogPluginState:
    def __init__(__self__, *,
                 attributes_json: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 internal_force_update: Optional[pulumi.Input[str]] = None,
                 internal_hmac_used_for_secrets_config_hmac: Optional[pulumi.Input[str]] = None,
                 internal_secrets_config_hmac: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 plugin_id: Optional[pulumi.Input[str]] = None,
                 plugin_name: Optional[pulumi.Input[str]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None,
                 secrets_hmac: Optional[pulumi.Input[str]] = None,
                 secrets_json: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BoundaryHostCatalogPlugin resources.
        :param pulumi.Input[str] attributes_json: The attributes for the host catalog. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host catalog.
        :param pulumi.Input[str] description: The host catalog description.
        :param pulumi.Input[str] internal_force_update: Internal only. Used to force update so that we can always check the value of secrets.
        :param pulumi.Input[str] internal_hmac_used_for_secrets_config_hmac: Internal only. The Boundary-provided HMAC used to calculate the current value of the HMAC'd config. Used for drift detection.
        :param pulumi.Input[str] internal_secrets_config_hmac: Internal only. HMAC of (serverSecretsHmac + config secrets). Used for proper secrets handling.
        :param pulumi.Input[str] name: The host catalog name. Defaults to the resource name.
        :param pulumi.Input[str] plugin_id: The ID of the plugin that should back the resource. This or plugin_name must be defined.
        :param pulumi.Input[str] plugin_name: The name of the plugin that should back the resource. This or plugin_id must be defined.
        :param pulumi.Input[str] scope_id: The scope ID in which the resource is created.
        :param pulumi.Input[str] secrets_hmac: The HMAC'd secrets value returned from the server.
        :param pulumi.Input[str] secrets_json: The secrets for the host catalog. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" to clear any existing values. NOTE: Unlike "attributes_json", removing this block will NOT clear secrets from the host catalog; this allows injecting secrets for one call, then removing them for storage.
        """
        if attributes_json is not None:
            pulumi.set(__self__, "attributes_json", attributes_json)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if internal_force_update is not None:
            pulumi.set(__self__, "internal_force_update", internal_force_update)
        if internal_hmac_used_for_secrets_config_hmac is not None:
            pulumi.set(__self__, "internal_hmac_used_for_secrets_config_hmac", internal_hmac_used_for_secrets_config_hmac)
        if internal_secrets_config_hmac is not None:
            pulumi.set(__self__, "internal_secrets_config_hmac", internal_secrets_config_hmac)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if plugin_id is not None:
            pulumi.set(__self__, "plugin_id", plugin_id)
        if plugin_name is not None:
            pulumi.set(__self__, "plugin_name", plugin_name)
        if scope_id is not None:
            pulumi.set(__self__, "scope_id", scope_id)
        if secrets_hmac is not None:
            pulumi.set(__self__, "secrets_hmac", secrets_hmac)
        if secrets_json is not None:
            pulumi.set(__self__, "secrets_json", secrets_json)

    @property
    @pulumi.getter(name="attributesJson")
    def attributes_json(self) -> Optional[pulumi.Input[str]]:
        """
        The attributes for the host catalog. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host catalog.
        """
        return pulumi.get(self, "attributes_json")

    @attributes_json.setter
    def attributes_json(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "attributes_json", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The host catalog description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="internalForceUpdate")
    def internal_force_update(self) -> Optional[pulumi.Input[str]]:
        """
        Internal only. Used to force update so that we can always check the value of secrets.
        """
        return pulumi.get(self, "internal_force_update")

    @internal_force_update.setter
    def internal_force_update(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internal_force_update", value)

    @property
    @pulumi.getter(name="internalHmacUsedForSecretsConfigHmac")
    def internal_hmac_used_for_secrets_config_hmac(self) -> Optional[pulumi.Input[str]]:
        """
        Internal only. The Boundary-provided HMAC used to calculate the current value of the HMAC'd config. Used for drift detection.
        """
        return pulumi.get(self, "internal_hmac_used_for_secrets_config_hmac")

    @internal_hmac_used_for_secrets_config_hmac.setter
    def internal_hmac_used_for_secrets_config_hmac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internal_hmac_used_for_secrets_config_hmac", value)

    @property
    @pulumi.getter(name="internalSecretsConfigHmac")
    def internal_secrets_config_hmac(self) -> Optional[pulumi.Input[str]]:
        """
        Internal only. HMAC of (serverSecretsHmac + config secrets). Used for proper secrets handling.
        """
        return pulumi.get(self, "internal_secrets_config_hmac")

    @internal_secrets_config_hmac.setter
    def internal_secrets_config_hmac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internal_secrets_config_hmac", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The host catalog name. Defaults to the resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="pluginId")
    def plugin_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the plugin that should back the resource. This or plugin_name must be defined.
        """
        return pulumi.get(self, "plugin_id")

    @plugin_id.setter
    def plugin_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plugin_id", value)

    @property
    @pulumi.getter(name="pluginName")
    def plugin_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the plugin that should back the resource. This or plugin_id must be defined.
        """
        return pulumi.get(self, "plugin_name")

    @plugin_name.setter
    def plugin_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plugin_name", value)

    @property
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> Optional[pulumi.Input[str]]:
        """
        The scope ID in which the resource is created.
        """
        return pulumi.get(self, "scope_id")

    @scope_id.setter
    def scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope_id", value)

    @property
    @pulumi.getter(name="secretsHmac")
    def secrets_hmac(self) -> Optional[pulumi.Input[str]]:
        """
        The HMAC'd secrets value returned from the server.
        """
        return pulumi.get(self, "secrets_hmac")

    @secrets_hmac.setter
    def secrets_hmac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secrets_hmac", value)

    @property
    @pulumi.getter(name="secretsJson")
    def secrets_json(self) -> Optional[pulumi.Input[str]]:
        """
        The secrets for the host catalog. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" to clear any existing values. NOTE: Unlike "attributes_json", removing this block will NOT clear secrets from the host catalog; this allows injecting secrets for one call, then removing them for storage.
        """
        return pulumi.get(self, "secrets_json")

    @secrets_json.setter
    def secrets_json(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secrets_json", value)


class BoundaryHostCatalogPlugin(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attributes_json: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 internal_force_update: Optional[pulumi.Input[str]] = None,
                 internal_hmac_used_for_secrets_config_hmac: Optional[pulumi.Input[str]] = None,
                 internal_secrets_config_hmac: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 plugin_id: Optional[pulumi.Input[str]] = None,
                 plugin_name: Optional[pulumi.Input[str]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None,
                 secrets_hmac: Optional[pulumi.Input[str]] = None,
                 secrets_json: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The host catalog resource allows you to configure a Boundary plugin-type host catalog. Host catalogs are always part of a project, so a project resource should be used inline or you should have the project ID in hand to successfully configure a host catalog.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] attributes_json: The attributes for the host catalog. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host catalog.
        :param pulumi.Input[str] description: The host catalog description.
        :param pulumi.Input[str] internal_force_update: Internal only. Used to force update so that we can always check the value of secrets.
        :param pulumi.Input[str] internal_hmac_used_for_secrets_config_hmac: Internal only. The Boundary-provided HMAC used to calculate the current value of the HMAC'd config. Used for drift detection.
        :param pulumi.Input[str] internal_secrets_config_hmac: Internal only. HMAC of (serverSecretsHmac + config secrets). Used for proper secrets handling.
        :param pulumi.Input[str] name: The host catalog name. Defaults to the resource name.
        :param pulumi.Input[str] plugin_id: The ID of the plugin that should back the resource. This or plugin_name must be defined.
        :param pulumi.Input[str] plugin_name: The name of the plugin that should back the resource. This or plugin_id must be defined.
        :param pulumi.Input[str] scope_id: The scope ID in which the resource is created.
        :param pulumi.Input[str] secrets_hmac: The HMAC'd secrets value returned from the server.
        :param pulumi.Input[str] secrets_json: The secrets for the host catalog. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" to clear any existing values. NOTE: Unlike "attributes_json", removing this block will NOT clear secrets from the host catalog; this allows injecting secrets for one call, then removing them for storage.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BoundaryHostCatalogPluginArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The host catalog resource allows you to configure a Boundary plugin-type host catalog. Host catalogs are always part of a project, so a project resource should be used inline or you should have the project ID in hand to successfully configure a host catalog.

        :param str resource_name: The name of the resource.
        :param BoundaryHostCatalogPluginArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BoundaryHostCatalogPluginArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attributes_json: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 internal_force_update: Optional[pulumi.Input[str]] = None,
                 internal_hmac_used_for_secrets_config_hmac: Optional[pulumi.Input[str]] = None,
                 internal_secrets_config_hmac: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 plugin_id: Optional[pulumi.Input[str]] = None,
                 plugin_name: Optional[pulumi.Input[str]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None,
                 secrets_hmac: Optional[pulumi.Input[str]] = None,
                 secrets_json: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BoundaryHostCatalogPluginArgs.__new__(BoundaryHostCatalogPluginArgs)

            __props__.__dict__["attributes_json"] = attributes_json
            __props__.__dict__["description"] = description
            __props__.__dict__["internal_force_update"] = internal_force_update
            __props__.__dict__["internal_hmac_used_for_secrets_config_hmac"] = internal_hmac_used_for_secrets_config_hmac
            __props__.__dict__["internal_secrets_config_hmac"] = internal_secrets_config_hmac
            __props__.__dict__["name"] = name
            __props__.__dict__["plugin_id"] = plugin_id
            __props__.__dict__["plugin_name"] = plugin_name
            if scope_id is None and not opts.urn:
                raise TypeError("Missing required property 'scope_id'")
            __props__.__dict__["scope_id"] = scope_id
            __props__.__dict__["secrets_hmac"] = secrets_hmac
            __props__.__dict__["secrets_json"] = secrets_json
        super(BoundaryHostCatalogPlugin, __self__).__init__(
            'boundary:index/boundaryHostCatalogPlugin:BoundaryHostCatalogPlugin',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attributes_json: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            internal_force_update: Optional[pulumi.Input[str]] = None,
            internal_hmac_used_for_secrets_config_hmac: Optional[pulumi.Input[str]] = None,
            internal_secrets_config_hmac: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            plugin_id: Optional[pulumi.Input[str]] = None,
            plugin_name: Optional[pulumi.Input[str]] = None,
            scope_id: Optional[pulumi.Input[str]] = None,
            secrets_hmac: Optional[pulumi.Input[str]] = None,
            secrets_json: Optional[pulumi.Input[str]] = None) -> 'BoundaryHostCatalogPlugin':
        """
        Get an existing BoundaryHostCatalogPlugin resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] attributes_json: The attributes for the host catalog. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host catalog.
        :param pulumi.Input[str] description: The host catalog description.
        :param pulumi.Input[str] internal_force_update: Internal only. Used to force update so that we can always check the value of secrets.
        :param pulumi.Input[str] internal_hmac_used_for_secrets_config_hmac: Internal only. The Boundary-provided HMAC used to calculate the current value of the HMAC'd config. Used for drift detection.
        :param pulumi.Input[str] internal_secrets_config_hmac: Internal only. HMAC of (serverSecretsHmac + config secrets). Used for proper secrets handling.
        :param pulumi.Input[str] name: The host catalog name. Defaults to the resource name.
        :param pulumi.Input[str] plugin_id: The ID of the plugin that should back the resource. This or plugin_name must be defined.
        :param pulumi.Input[str] plugin_name: The name of the plugin that should back the resource. This or plugin_id must be defined.
        :param pulumi.Input[str] scope_id: The scope ID in which the resource is created.
        :param pulumi.Input[str] secrets_hmac: The HMAC'd secrets value returned from the server.
        :param pulumi.Input[str] secrets_json: The secrets for the host catalog. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" to clear any existing values. NOTE: Unlike "attributes_json", removing this block will NOT clear secrets from the host catalog; this allows injecting secrets for one call, then removing them for storage.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BoundaryHostCatalogPluginState.__new__(_BoundaryHostCatalogPluginState)

        __props__.__dict__["attributes_json"] = attributes_json
        __props__.__dict__["description"] = description
        __props__.__dict__["internal_force_update"] = internal_force_update
        __props__.__dict__["internal_hmac_used_for_secrets_config_hmac"] = internal_hmac_used_for_secrets_config_hmac
        __props__.__dict__["internal_secrets_config_hmac"] = internal_secrets_config_hmac
        __props__.__dict__["name"] = name
        __props__.__dict__["plugin_id"] = plugin_id
        __props__.__dict__["plugin_name"] = plugin_name
        __props__.__dict__["scope_id"] = scope_id
        __props__.__dict__["secrets_hmac"] = secrets_hmac
        __props__.__dict__["secrets_json"] = secrets_json
        return BoundaryHostCatalogPlugin(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="attributesJson")
    def attributes_json(self) -> pulumi.Output[Optional[str]]:
        """
        The attributes for the host catalog. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host catalog.
        """
        return pulumi.get(self, "attributes_json")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The host catalog description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="internalForceUpdate")
    def internal_force_update(self) -> pulumi.Output[str]:
        """
        Internal only. Used to force update so that we can always check the value of secrets.
        """
        return pulumi.get(self, "internal_force_update")

    @property
    @pulumi.getter(name="internalHmacUsedForSecretsConfigHmac")
    def internal_hmac_used_for_secrets_config_hmac(self) -> pulumi.Output[str]:
        """
        Internal only. The Boundary-provided HMAC used to calculate the current value of the HMAC'd config. Used for drift detection.
        """
        return pulumi.get(self, "internal_hmac_used_for_secrets_config_hmac")

    @property
    @pulumi.getter(name="internalSecretsConfigHmac")
    def internal_secrets_config_hmac(self) -> pulumi.Output[str]:
        """
        Internal only. HMAC of (serverSecretsHmac + config secrets). Used for proper secrets handling.
        """
        return pulumi.get(self, "internal_secrets_config_hmac")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The host catalog name. Defaults to the resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pluginId")
    def plugin_id(self) -> pulumi.Output[str]:
        """
        The ID of the plugin that should back the resource. This or plugin_name must be defined.
        """
        return pulumi.get(self, "plugin_id")

    @property
    @pulumi.getter(name="pluginName")
    def plugin_name(self) -> pulumi.Output[str]:
        """
        The name of the plugin that should back the resource. This or plugin_id must be defined.
        """
        return pulumi.get(self, "plugin_name")

    @property
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> pulumi.Output[str]:
        """
        The scope ID in which the resource is created.
        """
        return pulumi.get(self, "scope_id")

    @property
    @pulumi.getter(name="secretsHmac")
    def secrets_hmac(self) -> pulumi.Output[str]:
        """
        The HMAC'd secrets value returned from the server.
        """
        return pulumi.get(self, "secrets_hmac")

    @property
    @pulumi.getter(name="secretsJson")
    def secrets_json(self) -> pulumi.Output[Optional[str]]:
        """
        The secrets for the host catalog. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" to clear any existing values. NOTE: Unlike "attributes_json", removing this block will NOT clear secrets from the host catalog; this allows injecting secrets for one call, then removing them for storage.
        """
        return pulumi.get(self, "secrets_json")
