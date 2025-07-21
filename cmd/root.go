package cmd

import (
    "github.com/spf13/cobra"
)

// RootCmd is the main entry point for GitMen CLI.
var RootCmd = &cobra.Command{
    Use:   "gitmen",
    Short: "GitMen - Next-Generation CLI Command Center",
    Long:  `GitMen transforms the CLI into a powerful and delightful assistant for developers.`,
    Run: func(cmd *cobra.Command, args []string) {
        // Entry point logic
    },
}

func Execute() {
    cobra.CheckErr(RootCmd.Execute())
}