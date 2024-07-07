package main

import (
    "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
    "github.com/bridge-yt/terraform-provider-bridge/bridge"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: bridge.Provider,
	})
}
