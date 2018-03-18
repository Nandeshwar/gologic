package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			for _, word := range strings.Split(line, " ") {
				if word == "" {
					continue
				}
				counts[word]++
			}
		}
	}

	for word, n := range counts {
		if n >= 1 {
			fmt.Printf("%d\t%s\n", n, word)
		}
	}
}
