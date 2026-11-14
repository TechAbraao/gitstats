package output

import ("fmt")

type CommitsInfo struct {
	Commits int
	StatusCode int
}

func PrintAllCommitsInRepository(info CommitsInfo) {

	if info.StatusCode != 200 {
		fmt.Println("\n[X] Repository Not Found.\n")
		return
	}
	fmt.Println("\n[X] Number of commits in this repository:", info.Commits, "\n")
}
