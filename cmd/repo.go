package cmd

import (
    "github.com/spf13/cobra"
    "gitstats/internal/github"
    "gitstats/internal/output"
    "encoding/json"
)

var repoCmd = &cobra.Command{
    Use:   "repo [owner/repository]",
    Short: "Exibe informações sobre um repositório do GitHub",
    Args:  cobra.ExactArgs(1),

    RunE: func(cmd *cobra.Command, args []string) error {
        repository := args[0]
        
        client := github.NewClient()
        body, err := github.GetRepositoryInfo(client, repository)
        if err != nil {
            return err
        }

        info := output.RepositoryInfo{}       
        if err := json.Unmarshal(body, &info); err != nil {
	        return err
        }       
        
        output.PrintRepositoryInfo(info)

        return nil
    },
}

func init() {
    rootCmd.AddCommand(repoCmd)
}