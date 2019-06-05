package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func parseArgs() string {
	// Check for at least one arg, bail if none
	if len(os.Args) < 2 {
		log.Fatalln("You must provide exactly one GitHub username.")
	}

	// We have one arg possible comma separated
	if len(os.Args) == 2 {
		return os.Args[1]
	} else {
		log.Fatalln("You have provided too many arguments")
	}
	return ""
}

func fetchKeys(username string) string {
	log.Printf("Fetching keys for user %s", username)

	url := fmt.Sprintf("https://github.com/%s.keys", username)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		bodyString := string(bodyBytes)
		return bodyString
	}
	return ""
}

func checkUsername(username string) {
	log.Printf("Checking for GitHub user %s", username)

	url := fmt.Sprintf("https://github.com/%s.keys", username)
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	if response.StatusCode != http.StatusOK {
		log.Fatalf("%s is an invalid user", username)
	}

	log.Printf("Found valid user %s", username)

	return
}

func main() {
	username := parseArgs()
	checkUsername(username)
	keys := fetchKeys(username)
	fmt.Print(keys)
	return
}
