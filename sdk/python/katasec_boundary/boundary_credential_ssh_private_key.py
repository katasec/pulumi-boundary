# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['BoundaryCredentialSshPrivateKeyArgs', 'BoundaryCredentialSshPrivateKey']

@pulumi.input_type
class BoundaryCredentialSshPrivateKeyArgs:
    def __init__(__self__, *,
                 credential_store_id: pulumi.Input[str],
                 private_key: pulumi.Input[str],
                 username: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_key_passphrase: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BoundaryCredentialSshPrivateKey resource.
        :param pulumi.Input[str] credential_store_id: ID of the credential store this credential belongs to.
        :param pulumi.Input[str] private_key: The private key associated with the credential.
        :param pulumi.Input[str] username: The username associated with the credential.
        :param pulumi.Input[str] description: The description of the credential.
        :param pulumi.Input[str] name: The name of the credential. Defaults to the resource name.
        :param pulumi.Input[str] private_key_passphrase: The passphrase of the private key associated with the credential.
        """
        pulumi.set(__self__, "credential_store_id", credential_store_id)
        pulumi.set(__self__, "private_key", private_key)
        pulumi.set(__self__, "username", username)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if private_key_passphrase is not None:
            pulumi.set(__self__, "private_key_passphrase", private_key_passphrase)

    @property
    @pulumi.getter(name="credentialStoreId")
    def credential_store_id(self) -> pulumi.Input[str]:
        """
        ID of the credential store this credential belongs to.
        """
        return pulumi.get(self, "credential_store_id")

    @credential_store_id.setter
    def credential_store_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "credential_store_id", value)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Input[str]:
        """
        The private key associated with the credential.
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "private_key", value)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        """
        The username associated with the credential.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the credential.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the credential. Defaults to the resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="privateKeyPassphrase")
    def private_key_passphrase(self) -> Optional[pulumi.Input[str]]:
        """
        The passphrase of the private key associated with the credential.
        """
        return pulumi.get(self, "private_key_passphrase")

    @private_key_passphrase.setter
    def private_key_passphrase(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key_passphrase", value)


@pulumi.input_type
class _BoundaryCredentialSshPrivateKeyState:
    def __init__(__self__, *,
                 credential_store_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 private_key_hmac: Optional[pulumi.Input[str]] = None,
                 private_key_passphrase: Optional[pulumi.Input[str]] = None,
                 private_key_passphrase_hmac: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BoundaryCredentialSshPrivateKey resources.
        :param pulumi.Input[str] credential_store_id: ID of the credential store this credential belongs to.
        :param pulumi.Input[str] description: The description of the credential.
        :param pulumi.Input[str] name: The name of the credential. Defaults to the resource name.
        :param pulumi.Input[str] private_key: The private key associated with the credential.
        :param pulumi.Input[str] private_key_hmac: The private key hmac.
        :param pulumi.Input[str] private_key_passphrase: The passphrase of the private key associated with the credential.
        :param pulumi.Input[str] private_key_passphrase_hmac: The private key passphrase hmac.
        :param pulumi.Input[str] username: The username associated with the credential.
        """
        if credential_store_id is not None:
            pulumi.set(__self__, "credential_store_id", credential_store_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if private_key is not None:
            pulumi.set(__self__, "private_key", private_key)
        if private_key_hmac is not None:
            pulumi.set(__self__, "private_key_hmac", private_key_hmac)
        if private_key_passphrase is not None:
            pulumi.set(__self__, "private_key_passphrase", private_key_passphrase)
        if private_key_passphrase_hmac is not None:
            pulumi.set(__self__, "private_key_passphrase_hmac", private_key_passphrase_hmac)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="credentialStoreId")
    def credential_store_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the credential store this credential belongs to.
        """
        return pulumi.get(self, "credential_store_id")

    @credential_store_id.setter
    def credential_store_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "credential_store_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the credential.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the credential. Defaults to the resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[pulumi.Input[str]]:
        """
        The private key associated with the credential.
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key", value)

    @property
    @pulumi.getter(name="privateKeyHmac")
    def private_key_hmac(self) -> Optional[pulumi.Input[str]]:
        """
        The private key hmac.
        """
        return pulumi.get(self, "private_key_hmac")

    @private_key_hmac.setter
    def private_key_hmac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key_hmac", value)

    @property
    @pulumi.getter(name="privateKeyPassphrase")
    def private_key_passphrase(self) -> Optional[pulumi.Input[str]]:
        """
        The passphrase of the private key associated with the credential.
        """
        return pulumi.get(self, "private_key_passphrase")

    @private_key_passphrase.setter
    def private_key_passphrase(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key_passphrase", value)

    @property
    @pulumi.getter(name="privateKeyPassphraseHmac")
    def private_key_passphrase_hmac(self) -> Optional[pulumi.Input[str]]:
        """
        The private key passphrase hmac.
        """
        return pulumi.get(self, "private_key_passphrase_hmac")

    @private_key_passphrase_hmac.setter
    def private_key_passphrase_hmac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key_passphrase_hmac", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The username associated with the credential.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class BoundaryCredentialSshPrivateKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 credential_store_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 private_key_passphrase: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The SSH private key credential resource allows you to configure a credential using a username, private key and optional passphrase.

        ## Example Usage

        ```python
        import pulumi
        import katasec_boundary as boundary

        org = boundary.BoundaryScope("org",
            description="global scope",
            scope_id="global",
            auto_create_admin_role=True,
            auto_create_default_role=True)
        project = boundary.BoundaryScope("project",
            description="My first scope!",
            scope_id=org.id,
            auto_create_admin_role=True)
        example_boundary_credential_store_static = boundary.BoundaryCredentialStoreStatic("exampleBoundaryCredentialStoreStatic",
            description="My first static credential store!",
            scope_id=project.id)
        example_boundary_credential_ssh_private_key = boundary.BoundaryCredentialSshPrivateKey("exampleBoundaryCredentialSshPrivateKey",
            description="My first ssh private key credential!",
            credential_store_id=example_boundary_credential_store_static.id,
            username="my-username",
            private_key=(lambda path: open(path).read())("~/.ssh/id_rsa"),
            private_key_passphrase="optional-passphrase")
        # change to the passphrase of the Private Key if required
        ```

        ## Import

        ```sh
         $ pulumi import boundary:index/boundaryCredentialSshPrivateKey:BoundaryCredentialSshPrivateKey example_ssh_private_key <my-id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] credential_store_id: ID of the credential store this credential belongs to.
        :param pulumi.Input[str] description: The description of the credential.
        :param pulumi.Input[str] name: The name of the credential. Defaults to the resource name.
        :param pulumi.Input[str] private_key: The private key associated with the credential.
        :param pulumi.Input[str] private_key_passphrase: The passphrase of the private key associated with the credential.
        :param pulumi.Input[str] username: The username associated with the credential.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BoundaryCredentialSshPrivateKeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The SSH private key credential resource allows you to configure a credential using a username, private key and optional passphrase.

        ## Example Usage

        ```python
        import pulumi
        import katasec_boundary as boundary

        org = boundary.BoundaryScope("org",
            description="global scope",
            scope_id="global",
            auto_create_admin_role=True,
            auto_create_default_role=True)
        project = boundary.BoundaryScope("project",
            description="My first scope!",
            scope_id=org.id,
            auto_create_admin_role=True)
        example_boundary_credential_store_static = boundary.BoundaryCredentialStoreStatic("exampleBoundaryCredentialStoreStatic",
            description="My first static credential store!",
            scope_id=project.id)
        example_boundary_credential_ssh_private_key = boundary.BoundaryCredentialSshPrivateKey("exampleBoundaryCredentialSshPrivateKey",
            description="My first ssh private key credential!",
            credential_store_id=example_boundary_credential_store_static.id,
            username="my-username",
            private_key=(lambda path: open(path).read())("~/.ssh/id_rsa"),
            private_key_passphrase="optional-passphrase")
        # change to the passphrase of the Private Key if required
        ```

        ## Import

        ```sh
         $ pulumi import boundary:index/boundaryCredentialSshPrivateKey:BoundaryCredentialSshPrivateKey example_ssh_private_key <my-id>
        ```

        :param str resource_name: The name of the resource.
        :param BoundaryCredentialSshPrivateKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BoundaryCredentialSshPrivateKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 credential_store_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 private_key_passphrase: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BoundaryCredentialSshPrivateKeyArgs.__new__(BoundaryCredentialSshPrivateKeyArgs)

            if credential_store_id is None and not opts.urn:
                raise TypeError("Missing required property 'credential_store_id'")
            __props__.__dict__["credential_store_id"] = credential_store_id
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if private_key is None and not opts.urn:
                raise TypeError("Missing required property 'private_key'")
            __props__.__dict__["private_key"] = private_key
            __props__.__dict__["private_key_passphrase"] = private_key_passphrase
            if username is None and not opts.urn:
                raise TypeError("Missing required property 'username'")
            __props__.__dict__["username"] = username
            __props__.__dict__["private_key_hmac"] = None
            __props__.__dict__["private_key_passphrase_hmac"] = None
        super(BoundaryCredentialSshPrivateKey, __self__).__init__(
            'boundary:index/boundaryCredentialSshPrivateKey:BoundaryCredentialSshPrivateKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            credential_store_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            private_key: Optional[pulumi.Input[str]] = None,
            private_key_hmac: Optional[pulumi.Input[str]] = None,
            private_key_passphrase: Optional[pulumi.Input[str]] = None,
            private_key_passphrase_hmac: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'BoundaryCredentialSshPrivateKey':
        """
        Get an existing BoundaryCredentialSshPrivateKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] credential_store_id: ID of the credential store this credential belongs to.
        :param pulumi.Input[str] description: The description of the credential.
        :param pulumi.Input[str] name: The name of the credential. Defaults to the resource name.
        :param pulumi.Input[str] private_key: The private key associated with the credential.
        :param pulumi.Input[str] private_key_hmac: The private key hmac.
        :param pulumi.Input[str] private_key_passphrase: The passphrase of the private key associated with the credential.
        :param pulumi.Input[str] private_key_passphrase_hmac: The private key passphrase hmac.
        :param pulumi.Input[str] username: The username associated with the credential.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BoundaryCredentialSshPrivateKeyState.__new__(_BoundaryCredentialSshPrivateKeyState)

        __props__.__dict__["credential_store_id"] = credential_store_id
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["private_key"] = private_key
        __props__.__dict__["private_key_hmac"] = private_key_hmac
        __props__.__dict__["private_key_passphrase"] = private_key_passphrase
        __props__.__dict__["private_key_passphrase_hmac"] = private_key_passphrase_hmac
        __props__.__dict__["username"] = username
        return BoundaryCredentialSshPrivateKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="credentialStoreId")
    def credential_store_id(self) -> pulumi.Output[str]:
        """
        ID of the credential store this credential belongs to.
        """
        return pulumi.get(self, "credential_store_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the credential.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the credential. Defaults to the resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Output[str]:
        """
        The private key associated with the credential.
        """
        return pulumi.get(self, "private_key")

    @property
    @pulumi.getter(name="privateKeyHmac")
    def private_key_hmac(self) -> pulumi.Output[str]:
        """
        The private key hmac.
        """
        return pulumi.get(self, "private_key_hmac")

    @property
    @pulumi.getter(name="privateKeyPassphrase")
    def private_key_passphrase(self) -> pulumi.Output[Optional[str]]:
        """
        The passphrase of the private key associated with the credential.
        """
        return pulumi.get(self, "private_key_passphrase")

    @property
    @pulumi.getter(name="privateKeyPassphraseHmac")
    def private_key_passphrase_hmac(self) -> pulumi.Output[str]:
        """
        The private key passphrase hmac.
        """
        return pulumi.get(self, "private_key_passphrase_hmac")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        The username associated with the credential.
        """
        return pulumi.get(self, "username")

