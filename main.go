package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"
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

// Need to fix this to not be an infinite loop
func checkResolvers() {
	i := 1
	for i < 3 {
		log.Println("Checking if DNS is working")
		_, err := net.LookupIP("github.com")
		if err != nil {
			log.Println("No DNS yet, trying again in 5s")
			time.Sleep(5 * time.Second)
			i += 1
		} else {
			break
		}
	}
	if i == 3 {
		log.Fatalln("Could not reliably lookup host github.com")
	}
	return
}

func main() {
	username := parseArgs()
	checkResolvers()
	checkUsername(username)
	keys := fetchKeys(username)
	fmt.Print(keys)
	return
}
