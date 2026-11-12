package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
    Use:   "repo [owner/repository]",
    Short: "Display repository information",
    Args:  cobra.ExactArgs(1),

    RunE: func(cmd *cobra.Command, args []string) error {
        repository := args[0]

        fmt.Println("Repository:", repository)

        return nil
    },
}

func init() {
    rootCmd.AddCommand(repoCmd)
}