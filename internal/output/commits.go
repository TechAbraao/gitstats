package output

import ("fmt")

type CommitsInfo struct {
	Commits int
	StatusCode int
}

func PrintAllCommitsInRepository(info CommitsInfo) {

	if info.StatusCode != 200 {
		fmt.Println("[X] Repository Not Found.")
	}
	fmt.Println("\n[X] Number of commits in this repository:", info.Commits, "\n")
}
