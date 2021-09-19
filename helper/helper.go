package helper

import (
	"plugin"

	"github.com/Eldius/webcomics-fetcher2-go/plugins"
)

func PluginInfo() *plugins.Plugin {
	return &plugins.Plugin{
		Name:        "oots",
		Path:        plugin.GetBinaryPath(),
		Description: "Order of The Stick plugin",
	}
}
