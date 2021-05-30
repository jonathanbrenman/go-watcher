package cmd

import (
	"github.com/spf13/cobra"
	"go-watcher/watcher"
	"log"
	"strings"
	"time"
)

var WatchCmd = &cobra.Command{
  Use:   "watch",
  Short: "List files in path recursive",
  Run: watchCmd,
}

func init() {
	// Flags cmd watch here
	WatchCmd.Flags().String("path", "","Directory to watch")
	WatchCmd.Flags().Duration("delay", 3 * time.Second,"Delay for next execution")
	WatchCmd.Flags().Bool("debug", false,"Debug will give more info about the process")
	WatchCmd.Flags().String("filters", "","Extension to filter with comma separated")
}

func watchCmd(cmd *cobra.Command, args []string) {
	path, _ := cmd.Flags().GetString("path")
	if path == "" {
		log.Fatal("missing path flag.")
	}
  	delay, _ := cmd.Flags().GetDuration("delay")
  	debug, _ := cmd.Flags().GetBool("debug")
  	filters, _ := cmd.Flags().GetString("filters")

	// Watcher builder
	wf := watcher.New().
		SetDelay(delay).
		SetDebug(debug).
		SetExtFilter(strings.Split(filters, ",")...).
		Build()

	// Start watcher
	wf.Add(path)
}
