package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

const VERSION = "0.0.1"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("v%s", VERSION)
		fmt.Printf("v%s", VERSION)
	},
}
