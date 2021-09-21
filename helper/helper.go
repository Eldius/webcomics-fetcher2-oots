package helper

import (
	"github.com/Eldius/webcomics-fetcher2-go/plugins"
)

func PluginInfo() *plugins.PluginInfo {
	return &plugins.PluginInfo{
		Name:        "oots",
		Path:        plugins.GetAbsolutePath(),
		Description: "Order of The Stick plugin",
	}
}
