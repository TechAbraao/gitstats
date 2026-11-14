package cmd

import (
	"gitstats/internal/github"
	"gitstats/internal/output"
	"github.com/spf13/cobra"
)

var commitsCmd = &cobra.Command{
    Use:   "commits [owner/repository]",
    Short: "Exibe a quantidade total de commits em um repositório.",
    Args:  cobra.ExactArgs(1),

    RunE: func(cmd *cobra.Command, args []string) error {
        repository := args[0]

        client := github.NewClient()
        body, err := github.GetCountCommitsInfo(client, repository)
        if err != nil {
            return err
        }

        allComits := output.CommitsInfo{
            Commits: body,
        }

        output.PrintAllCommitsInRepository(allComits)

        return nil
    },
}

func init() {
    rootCmd.AddCommand(commitsCmd)
}