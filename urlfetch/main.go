//go:build !solution

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, j := range args {
		resp, _ := http.Get(j)
		body, _ := io.ReadAll(resp.Body)
		defer resp.Body.Close()
		fmt.Println(string(body))
	}
}
