package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func fetch(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("X Error fetching URL:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("X Error reading body: ", err)
		return
	}
	fmt.Printf("Fetched %d bytes from %s\n", len(body), url)

}

func main() {
	var wg sync.WaitGroup

	urls := []string{
		"https://golang.org",
		"https://example.com",
		"https://github.com",
		"https://stackoverflow.com",
		"https://medium.com",
	}
	for _, url := range urls {
		wg.Add(1)
		go fetch(url, &wg)
	}
	wg.Wait()
	fmt.Println("All fetches completed")
}
