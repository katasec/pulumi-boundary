package shim

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/katasec/terraform-provider-boundary/internal/provider"
)

func Provider() *schema.Provider {
	return provider.New()
}
