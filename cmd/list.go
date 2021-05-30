package cmd

import (
	"github.com/spf13/cobra"
	"go-watcher/watcher"
	"strings"
)

var ListCmd = &cobra.Command{
  Use:   "list",
  Short: "List files in path recursive",
  Run: listCmd,
}

func init() {
	// Flags cmd watch here
	ListCmd.Flags().String("path", "./","Directory to list")
	ListCmd.Flags().String("filters", "","Extension to filter with comma separated")
}

func listCmd(cmd *cobra.Command, args []string) {
	path, _ := cmd.Flags().GetString("path")
  	filters, _ := cmd.Flags().GetString("filters")

	// Watcher builder
	wf := watcher.New().
		SetExtFilter(strings.Split(filters, ",")...).
		Build()

	// Start watcher
	wf.List(path)
}
