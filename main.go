package main

import (
	p "github.com/Mongey/terraform-provider-middesk/internal/provider"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: p.Provider})
}
