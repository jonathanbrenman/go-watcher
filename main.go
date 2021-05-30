package main

import (
	"github.com/spf13/cobra"
	"go-watcher/cmd"
)

func main() {
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmd.WatchCmd)
	rootCmd.Execute()
}