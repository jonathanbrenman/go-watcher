package cmd

import (
	"github.com/spf13/cobra"
	"go-watcher/watcher"
	"time"
)

var WatchCmd = &cobra.Command{
  Use:   "watch",
  Short: "List files in path recursive",
  Run: watchCmd,
}

func init() {
	// Flags cmd watch here
	WatchCmd.Flags().String("path", "./","Directory to watch")
	WatchCmd.Flags().Duration("delay", 3,"Delay for next execution in seconds")
	WatchCmd.Flags().Bool("debug", false,"Debug will give more info about the process")
	WatchCmd.Flags().StringArray("filters", []string{},"Extension to filter with comma separated")
}

func watchCmd(cmd *cobra.Command, args []string) {
	path, _ := cmd.Flags().GetString("path")
  	delay, _ := cmd.Flags().GetDuration("delay")
  	debug, _ := cmd.Flags().GetBool("debug")
  	filters, _ := cmd.Flags().GetStringArray("filters")

  	// Watcher builder
	wf := watcher.New().
		SetDelay(delay * time.Second).
		SetDebug(debug).
		SetExtFilter(filters...).
		Build()

	// Start watcher
	wf.Add(path)
}
