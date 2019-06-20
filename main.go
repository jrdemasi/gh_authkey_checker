package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

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
	// Ensure we have the correct number of arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: gh_authkey_checker <username>")
		os.Exit(1)
	}

	username := os.Args[1]

	log.Printf("Fetching keys for user %s", username)
	keys, err := fetchKeys(username)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(keys)
}
