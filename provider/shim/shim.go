package shim

import (
	tfpf "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/kong/terraform-provider-konnect/internal/provider"
)

func NewProvider() tfpf.Provider {
	return provider.New("dev")()
}
