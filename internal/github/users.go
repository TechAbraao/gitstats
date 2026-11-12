package github

import "io"

func GetUserPrfileInfo(client *Client, username string) ([]byte, error) {
	resp, err := client.httpClient.Get("https://api.github.com/users/" + username)
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