package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		log.Fatal("Please specify two file names to compare as arguments.")
	}

	aText, err := os.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}

	bText, err := os.ReadFile(args[1])
	if err != nil {
		log.Fatal(err)
	}

	var diffLines []string
	for _, a := range strings.Split(string(aText), "\n") {
		var isFound bool
		for _, b := range strings.Split(string(bText), "\n") {
			if a == b {
				isFound = true
				break
			}
		}
		if !isFound {
			diffLines = append(diffLines, a)
		}
	}

	for _, line := range diffLines {
		fmt.Println(line)
	}
}
