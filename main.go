package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://golang.org"

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
