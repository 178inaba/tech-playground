package checkdigit_test

import (
	"fmt"
	"log"
	"strconv"

	"github.com/178inaba/tech-playground/checkdigit"
)

func ExampleNewLuhn() {
	p := checkdigit.NewLuhn()

	const seed = "411111111111111"
	cd, err := p.Generate(seed)
	if err != nil {
		log.Fatalln("failed to generate check digit")
	}

	ok := p.Verify(seed + strconv.Itoa(cd))
	fmt.Printf("seed: %s, check digit: %d, verify: %t\n", seed, cd, ok)

	// Output:
	// seed: 411111111111111, check digit: 1, verify: true
}
