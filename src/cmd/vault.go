package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"

	"path/filepath"
)
import "go-pass-keeper/src/actions"

var DatabaseFilePath string

func init() {

	vaultStoreCmd.PersistentFlags().StringVarP(&DatabaseFilePath, "user", "u", "", "... (required)")
	vaultStoreCmd.MarkPersistentFlagRequired("user")

	vaultCmd.AddCommand(vaultStoreCmd)
	vaultCmd.AddCommand(vaultListCmd)
	vaultCmd.AddCommand(vaultFindCmd)

	vaultCmd.PersistentFlags().StringVarP(&DatabaseFilePath, "file", "f", "", "File where the password will be stored. (required)")
	vaultCmd.MarkPersistentFlagRequired("file")

	rootCmd.AddCommand(vaultCmd)
}

var vaultCmd = &cobra.Command{
	Use:   "vault",
	Short: "Sub command to handle stored passwords",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		parent := filepath.Dir(DatabaseFilePath)
		if _, err := os.Stat(parent); os.IsNotExist(err) {
			cmd.Help()
			log.Fatalf("The pwd %s does not exit.", parent)
		}
	},
}

var vaultStoreCmd = &cobra.Command{
	Use:   "store",
	Short: "Save new entry",
	Run: func(cmd *cobra.Command, args []string) {
		user, _ := cmd.Flags().GetString("user")
		actions.StoreEntry(user)
	},
}
var vaultListCmd = &cobra.Command{
	Use:   "list",
	Short: "List All entries",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
var vaultFindCmd = &cobra.Command{
	Use:   "find",
	Short: "Find entry",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
