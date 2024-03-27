package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	filename := "members.yml"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Open %s: %v.", filename, err)
	}

	var members []string
	if err := yaml.NewDecoder(f).Decode(&members); err != nil {
		log.Fatalf("Decode yaml: %v.", err)
	}

	var pairs [][]string
	length := len(members)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			pairs = append(pairs, []string{members[i], members[j]})
		}
	}

	for len(pairs) > 0 {
		for i := len(pairs) - 1; i > 0; i-- {
			j := rand.Intn(i + 1)
			pairs[i], pairs[j] = pairs[j], pairs[i]
		}

		alreadyMemberMap := map[string]struct{}{}
		for i := len(pairs) - 1; i >= 0; i-- {
			_, ok0 := alreadyMemberMap[pairs[i][0]]
			_, ok1 := alreadyMemberMap[pairs[i][1]]
			if ok0 || ok1 {
				continue
			}

			fmt.Printf("%s on %s\n", pairs[i][0], pairs[i][1])
			alreadyMemberMap[pairs[i][0]] = struct{}{}
			alreadyMemberMap[pairs[i][1]] = struct{}{}
			pairs = append(pairs[:i], pairs[i+1:]...)
		}

		fmt.Println("----------")
	}
}
