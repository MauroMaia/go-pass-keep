package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go-pass-keeper/src/utils"
)

var Verbose bool
var UseColor bool
var Version string
var CommitId string
var BuildDate string

var rootCmd = &cobra.Command{
	Use:   "go-pass-keeper",
	Short: "go-pass-keeper is a very fast offline password keeper",
	Long: `go-pass-keeper is a very fast offline password keeper.
You should be able to save and search your password form you cli.

MIT License
Copyright (c) 2021 MauroMaia
`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
		os.Exit(1)
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		utils.SetLoggerOptions(UseColor, Verbose)
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&UseColor, "color", "c", false, "color output")
}

// Execute executes the root command.
func Execute(version string, commitId string, buildDate string) {
	Version = version
	CommitId = commitId
	BuildDate = buildDate

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
