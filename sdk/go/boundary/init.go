// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package boundary

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "boundary:index/boundaryAccount:BoundaryAccount":
		r = &BoundaryAccount{}
	case "boundary:index/boundaryAccountOidc:BoundaryAccountOidc":
		r = &BoundaryAccountOidc{}
	case "boundary:index/boundaryAccountPassword:BoundaryAccountPassword":
		r = &BoundaryAccountPassword{}
	case "boundary:index/boundaryAuthMethod:BoundaryAuthMethod":
		r = &BoundaryAuthMethod{}
	case "boundary:index/boundaryAuthMethodOidc:BoundaryAuthMethodOidc":
		r = &BoundaryAuthMethodOidc{}
	case "boundary:index/boundaryAuthMethodPassword:BoundaryAuthMethodPassword":
		r = &BoundaryAuthMethodPassword{}
	case "boundary:index/boundaryCredentialLibraryVault:BoundaryCredentialLibraryVault":
		r = &BoundaryCredentialLibraryVault{}
	case "boundary:index/boundaryCredentialSshPrivateKey:BoundaryCredentialSshPrivateKey":
		r = &BoundaryCredentialSshPrivateKey{}
	case "boundary:index/boundaryCredentialStoreStatic:BoundaryCredentialStoreStatic":
		r = &BoundaryCredentialStoreStatic{}
	case "boundary:index/boundaryCredentialStoreVault:BoundaryCredentialStoreVault":
		r = &BoundaryCredentialStoreVault{}
	case "boundary:index/boundaryCredentialUserName:BoundaryCredentialUserName":
		r = &BoundaryCredentialUserName{}
	case "boundary:index/boundaryGroup:BoundaryGroup":
		r = &BoundaryGroup{}
	case "boundary:index/boundaryHost:BoundaryHost":
		r = &BoundaryHost{}
	case "boundary:index/boundaryHostCatalog:BoundaryHostCatalog":
		r = &BoundaryHostCatalog{}
	case "boundary:index/boundaryHostCatalogPlugin:BoundaryHostCatalogPlugin":
		r = &BoundaryHostCatalogPlugin{}
	case "boundary:index/boundaryHostCatalogStatic:BoundaryHostCatalogStatic":
		r = &BoundaryHostCatalogStatic{}
	case "boundary:index/boundaryHostSet:BoundaryHostSet":
		r = &BoundaryHostSet{}
	case "boundary:index/boundaryHostSetPlugin:BoundaryHostSetPlugin":
		r = &BoundaryHostSetPlugin{}
	case "boundary:index/boundaryHostSetStatic:BoundaryHostSetStatic":
		r = &BoundaryHostSetStatic{}
	case "boundary:index/boundaryHostStatic:BoundaryHostStatic":
		r = &BoundaryHostStatic{}
	case "boundary:index/boundaryManagedGroup:BoundaryManagedGroup":
		r = &BoundaryManagedGroup{}
	case "boundary:index/boundaryRole:BoundaryRole":
		r = &BoundaryRole{}
	case "boundary:index/boundaryScope:BoundaryScope":
		r = &BoundaryScope{}
	case "boundary:index/boundaryTarget:BoundaryTarget":
		r = &BoundaryTarget{}
	case "boundary:index/boundaryUser:BoundaryUser":
		r = &BoundaryUser{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:boundary" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, _ := PkgVersion()
	pulumi.RegisterResourceModule(
		"boundary",
		"index/boundaryAccount",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/boundaryAccountOidc",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/boundaryAccountPassword",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/boundaryAuthMethod",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/boundaryAuthMethodOidc",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/boundaryAuthMethodPassword",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/boundaryCredentialLibraryVault",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/boundaryCredentialSshPrivateKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/boundaryCredentialStoreStatic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/boundaryCredentialStoreVault",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/boundaryCredentialUserName",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/boundaryGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/boundaryHost",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/boundaryHostCatalog",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/boundaryHostCatalogPlugin",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/boundaryHostCatalogStatic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/boundaryHostSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/boundaryHostSetPlugin",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/boundaryHostSetStatic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/boundaryHostStatic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/boundaryManagedGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/boundaryRole",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/boundaryScope",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/boundaryTarget",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/boundaryUser",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"boundary",
		&pkg{version},
	)
}