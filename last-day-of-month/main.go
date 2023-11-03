package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	flag.Parse()
	args := flag.Args()

	n := time.Now()
	month := n.Month()

	if len(args) == 1 {
		m, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}

		month = time.Month(m)
	}

	t := time.Date(n.Year(), month+1, 1, 0, 0, 0, 0, n.Location()).AddDate(0, 0, -1)

	fmt.Println(t)
	fmt.Println(t.Unix())
}
