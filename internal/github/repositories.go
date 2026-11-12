package github

import (
	"io"
)

func GetRepositoryInfo(client *Client, repository string) ([]byte, error) {
	resp, err := client.httpClient.Get("https://api.github.com/repos/" + repository)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
