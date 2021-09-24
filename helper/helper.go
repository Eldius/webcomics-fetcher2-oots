package helper

import (
	"github.com/Eldius/webcomics-fetcher2-go/plugins"
)

const (

	// PluginName is the plugin name (doh)
	PluginName = "oots"
)

/*
PluginInfo returns info about plugin to save in the main app
*/
func PluginInfo() *plugins.PluginInfo {
	return &plugins.PluginInfo{
		Name:        PluginName,
		Path:        plugins.GetAbsolutePath(),
		Description: "Order of The Stick plugin",
	}
}
