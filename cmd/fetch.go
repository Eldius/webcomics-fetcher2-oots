package cmd

import (
	"github.com/Eldius/webcomics-fetcher2-go/plugins"
	"github.com/Eldius/webcomics-fetcher2-oots/oots"
	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetches oots strips info",
	Long:  `Fetches oots strips info.`,
	Run: func(_ *cobra.Command, _ []string) {
		plugins.ToOutputFile(output, oots.Fetch())
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
