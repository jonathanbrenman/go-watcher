package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-watcher/cmd"
	"log"
)

func main() {
	welcome := `                                  __         .__                  
   ____   ____   __  _  _______ _/  |_  ____ |  |__   ___________ 
  / ___\ /  _ \  \ \/ \/ /\__  \\   __\/ ___\|  |  \_/ __ \_  __ \
 / /_/  >  <_> )  \     /  / __ \|  | \  \___|   Y  \  ___/|  | \/
 \___  / \____/    \/\_/  (____  /__|  \___  >___|  /\___  >__|   
/_____/                        \/          \/     \/     \/       `

	fmt.Println(welcome)
	fmt.Println()

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmd.WatchCmd)
	rootCmd.AddCommand(cmd.ListCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}