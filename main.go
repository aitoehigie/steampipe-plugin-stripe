package main

import (
	"github.com/aitoehigie/steampipe-plugin-stripe/stripe"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: stripe.Plugin})
}
