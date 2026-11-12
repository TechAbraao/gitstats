package output

import "fmt"

type UserProfileInfo struct {
	Username string `json:"login"`
	Name string `json:"name"`
	PublicRepos int `json:"public_repos"`
	Followers int `json:"followers"`
	Following int `json:"following"` 
	StatusCode string `json:"status"`
}	


func PrintUserProfileInfo(info UserProfileInfo) {
	if info.StatusCode == "404" {

		fmt.Println("\n[X] User Not Found\n")
		return
	}

	fmt.Println()
	fmt.Println("[X] Username:", info.Username)
	fmt.Println("[X] Name:", info.Name)
	fmt.Println("[X] Public Repos:", info.PublicRepos)
	fmt.Println("[X] Followers:", info.Followers)
	fmt.Println("[X] Following:", info.Following)
	fmt.Println()

}