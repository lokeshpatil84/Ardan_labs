package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
    name, publicRepos, err := Userinfo("ardanlabs")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Name:", name)
    fmt.Println("Public Repos:", publicRepos)
}

func Userinfo(login string) (string, int, error) {
	url := "https://api.github.com/users/" + login

	resp, err := http.Get(url)

	if err != nil {

		return "", 0, err
	}
	if resp.StatusCode != http.StatusOK {

		return "", 0, fmt.Errorf("%q - bad status getting: %s", url, resp.Status)
	}
	return parseResponse(resp.Body)
}

func parseResponse(r io.Reader) (string, int, error) {
	var reply struct {
		Name        string
		PublicRepos int `json:"Punlic_repos"`
	}
	dec := json.NewDecoder(r)
	if err := dec.Decode(&reply); err != nil {
		fmt.Println("error in decoding json", err)
		return "", 0, err
	}
	return reply.Name, reply.PublicRepos, nil
}
