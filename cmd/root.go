package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var v = "1.0.0"
var rootCmd = &cobra.Command{
	Use:     "chain-cmd",
	Short:   "Blockchain commands",
	Long:    "This is a blockchain cmd ",
	Version: v,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
