package helper

import (
	"github.com/Eldius/webcomics-fetcher2-go/plugins"
)

func PluginInfo() *plugins.Plugin {
	return &plugins.Plugin{
		Name:        "oots",
		Path:        plugins.GetBinaryPath(),
		Description: "Order of The Stick plugin",
	}
}
