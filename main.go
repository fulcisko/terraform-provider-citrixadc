// Package main is the entry point for the Citrix ADC Terraform provider.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/teleivo/terraform-provider-citrixadc/citrixadc"
)

// Generate docs for the provider
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	var debugMode bool
	var debugTimeout time.Duration

	// Default debug to false so the provider behaves like a normal release binary.
	// Pass -debug=true when running locally with delve or similar debuggers.
	// Note: I often run this with dlv using: dlv exec ./terraform-provider-citrixadc -- -debug=true
	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	// Increased default timeout from 10m to 30m since I frequently need more time to step through
	// breakpoints when exploring unfamiliar parts of the codebase.
	flag.DurationVar(&debugTimeout, "debug-timeout", 30*time.Minute, "how long to wait for a debugger to attach before exiting")
	flag.Parse()

	opts := &plugin.ServeOpts{
		ProviderFunc: citrixadc.Provider,
	}

	if debugMode {
		log.Println("Running in debug mode - attach your debugger now")
		log.Println("Provider address: registry.terraform.io/teleivo/citrixadc")
		log.Printf("Debug session will time out after %s", debugTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), debugTimeout)
		defer cancel()
		err := plugin.Debug(ctx, "registry.terraform.io/teleivo/citrixadc", opts)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	plugin.Serve(opts)
}
