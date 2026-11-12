package cmd

import (
    "github.com/spf13/cobra"
    "gitstats/internal/github"
    "gitstats/internal/output"
    "encoding/json"
)

var usersCmd = &cobra.Command{
    Use:   "user [GitHub Username]",
    Short: "Exibe informações sobre um usuário do GitHub.",
    Args:  cobra.ExactArgs(1),

    RunE: func(cmd *cobra.Command, args []string) error {
		username := args[0]


		client := github.NewClient()
		body, err := github.GetUserPrfileInfo(client, username)

		if err != nil {
			return nil
		}

		info := output.UserProfileInfo{}     
        if err := json.Unmarshal(body, &info); err != nil {
	        return err
        }     

		output.PrintUserProfileInfo(info)

		return nil
	},
}

func init() {
    rootCmd.AddCommand(usersCmd)
}