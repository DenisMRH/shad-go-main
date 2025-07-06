//go:build !solution

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

func main() {
	urls := os.Args[1:]
	if len(urls) == 0 {
		fmt.Println("Use urls after calling the programs")
	}
	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		url1 := url
		go func() {
			defer wg.Done()
			resp, err := http.Get(url1)
			if err != nil {
				fmt.Println(err)
				return
			}

			defer resp.Body.Close()
			_, _ = io.Copy(io.Discard, resp.Body)
			fmt.Println(resp.Status)
		}()
	}
	wg.Wait()
}
