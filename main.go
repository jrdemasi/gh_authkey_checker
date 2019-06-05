package main

import (
    "fmt"
    "log"
    "net"
    "os"
    "net/http"
    "io/ioutil"
    "time"
)

func parseArgs() string {
	// Check for at least one arg, bail if none
	if len(os.Args) < 2 {
		log.Fatalln("You must provide exactly one GitHub username.")
	}

	// We have one arg possible comma separated
	if len(os.Args) == 2 {
		return(os.Args[1])
	} else {
        log.Fatalln("You have provided too many arguments")
    }
    return("")
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
            log.Fatal(err)
        }
        bodyString := string(bodyBytes)
        return(bodyString)
    }
    return("")
}

// Need to fix this to not be an infinite loop
func checkResolvers() {
    i := 0
	for i < 3 {
		_, err := net.LookupIP("github.com")
		if err != nil {
			log.Println("No DNS yet, trying again in 5s")
            time.Sleep(5 * time.Second)
            i += 1
		} else {
			break
		}
	}
	return
}

func main() {
    username := parseArgs()
	checkResolvers()
    keys := fetchKeys(username)
    fmt.Print(keys)
    return
}
