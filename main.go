package main

import (
	"github.com/arykalin/terraform-example-plugin/example"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: example.Provider,
	})
}
