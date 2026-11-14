package github

import (
	"net/url"
	"strconv"
	"strings"
	"fmt"
)

func lastPage(link string) (int, error) {
	for _, parte := range strings.Split(link, ",") {
		if !strings.Contains(parte, `rel="last"`) {
			continue
		}

		ini := strings.Index(parte, "<")
		fim := strings.Index(parte, ">")
		if ini == -1 || fim == -1 {
			continue
		}

		u, err := url.Parse(parte[ini+1 : fim])
		if err != nil {
			return 0, err
		}
		return strconv.Atoi(u.Query().Get("page"))
	}
	return 0, fmt.Errorf("rel=last não encontrado")
}

func GetCountCommitsInfo(client *Client, repository string) (int, error) {
	resp, err := client.httpClient.Get("https://api.github.com/repos/" + repository + "/commits?per_page=1")

	if err != nil {
		return 0, err
	}
	
	defer resp.Body.Close()
	commitsRepository, err := lastPage(resp.Header.Get("Link"))

	return commitsRepository, nil
}