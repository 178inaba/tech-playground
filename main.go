package main

import (
	"fmt"

	"github.com/oklog/ulid/v2"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(ulid.Make())
	}
}
