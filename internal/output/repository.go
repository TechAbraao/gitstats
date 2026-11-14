package output

import "fmt"

type RepositoryInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stars       int    `json:"stargazers_count"`
	StatusCode  string `json:"status"`
	Forks	    int    `json:"forks"`
	OpenIssues  int    `json:"open_issues"`
}

func PrintRepositoryInfo(info RepositoryInfo) {

	if info.StatusCode == "404" {
		PrintRepositoryInfoNotFound()
		return
	}

	fmt.Println()
	fmt.Println("[X] Repository Name:", info.Name)
	fmt.Println("[X] Description:", info.Description)
	fmt.Println("[X] Stars:", info.Stars)
	fmt.Println("[X] Forks:", info.Forks)
	fmt.Println("[X] Open Issues:", info.OpenIssues)
	fmt.Println()
}

func PrintRepositoryInfoNotFound() {
	fmt.Println()
	fmt.Println("[X] There is no instance of this repository for this user." )
	fmt.Println()
}