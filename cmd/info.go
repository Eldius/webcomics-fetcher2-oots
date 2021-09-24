package cmd

import (
	"github.com/Eldius/webcomics-fetcher2-go/plugins"
	"github.com/Eldius/webcomics-fetcher2-oots/helper"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Returns info about this plugin",
	Long:  `Returns info about this plugin.`,
	Run: func(_ *cobra.Command, _ []string) {
		plugins.ToOutputFile(output, helper.PluginInfo())
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
