package main

import (
	"github.com/5h4d0wb0y/terraform-provider-pythonanywhere/pythonanywhere"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: pythonanywhere.Provider})
}
