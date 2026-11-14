package output

import ("fmt")

type CommitsInfo struct {
	Commits int
}

func PrintAllCommitsInRepository(info CommitsInfo) {
	fmt.Println("\n[X] Number of commits in this repository:", info.Commits, "\n")
}