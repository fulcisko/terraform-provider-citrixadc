// Package main is the entry point for the Citrix ADC Terraform provider.
package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/teleivo/terraform-provider-citrixadc/citrixadc"
)

// Generate docs for the provider
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	var debugMode bool

	// Default debug to true for local development convenience
	flag.BoolVar(&debugMode, "debug", true, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		ProviderFunc: citrixadc.Provider,
	}

	if debugMode {
		err := plugin.Debug(context.Background(), "registry.terraform.io/teleivo/citrixadc", opts)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	plugin.Serve(opts)
}
