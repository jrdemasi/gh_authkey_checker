package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func parseArgs() (string, error) {
	// We can only accept one argument
	if len(os.Args) != 2 {
		return "", fmt.Errorf("Usage: gh_authkey_checker <username>")
	}

	return os.Args[1], nil
}

func fetchKeys(username string) (string, error) {
	url := fmt.Sprintf("https://github.com/%s.keys", username)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return "", fmt.Errorf("%s is an invalid user", username)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Expected http 200 but got %d instead", resp.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

func main() {
	username, err := parseArgs()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Fetching keys for user %s", username)
	keys, err := fetchKeys(username)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(keys)
}
