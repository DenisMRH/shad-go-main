//go:build !solution

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	m := make(map[string]int)
	for _, file := range args {
		file, _ := os.Open(file)
		filebuf := bufio.NewScanner(file)
		for filebuf.Scan() {
			m[filebuf.Text()]++
		}
		file.Close()
	}
	for i, j := range m {
		if j != 1 {
			fmt.Printf("%d\t%s\n", j, i)
		}
	}

}
