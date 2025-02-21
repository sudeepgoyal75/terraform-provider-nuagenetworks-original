package main

import (
	"github.com/hashicorp/terraform/plugin"
	nuagenetworks "github.com/tpretz/terraform-provider-nuagenetworks/5.0.2.1/nuagenetworks"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: nuagenetworks.Provider})
}
