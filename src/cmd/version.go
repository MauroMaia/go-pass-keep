package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("%s - %s - %s", Version, CommitId, BuildDate)
		fmt.Printf("%s - %s - %s", Version, CommitId, BuildDate)
	},
}
