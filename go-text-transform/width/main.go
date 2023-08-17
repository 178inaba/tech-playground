package main

import (
	"fmt"

	"golang.org/x/text/width"
)

func main() {
	rs := []rune{'５', 'ｱ', 'ア', 'A', 'α'}
	fmt.Println("rune\tWide\tNarrow\tFolded\tKind")
	fmt.Println("--------------------------------------------------")
	for _, r := range rs {
		p := width.LookupRune(r)
		w, n, f, k := p.Wide(), p.Narrow(), p.Folded(), p.Kind()
		fmt.Printf("%2c\t%2c\t%3c\t%3c\t%s\n", r, w, n, f, k)
	}
}
