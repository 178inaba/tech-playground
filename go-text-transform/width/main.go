package main

import (
	"fmt"

	"golang.org/x/text/width"
)

func main() {
	fmt.Println("p.Wide(), p.Narrow(), p.Folded(), p.Kind() ------------------------------------------------------------")
	rs := []rune{'５', 'ｱ', 'ア', 'A', 'α'}
	fmt.Println("rune\tWide\tNarrow\tFolded\tKind")
	fmt.Println("--------------------------------------------------")
	for _, r := range rs {
		p := width.LookupRune(r)
		w, n, f, k := p.Wide(), p.Narrow(), p.Folded(), p.Kind()
		fmt.Printf("%2c\t%2c\t%3c\t%3c\t%s\n", r, w, n, f, k)
	}

	fmt.Println("")
	fmt.Println("width.Fold ------------------------------------------------------------")
	for _, r := range width.Fold.String(string(rs)) {
		p := width.LookupRune(r)
		fmt.Printf("%c: %s\n", r, p.Kind())
	}

	fmt.Println("")
	fmt.Println("width.Narrow ------------------------------------------------------------")
	for _, r := range width.Narrow.String(string(rs)) {
		p := width.LookupRune(r)
		fmt.Printf("%c: %s\n", r, p.Kind())
	}

	fmt.Println("")
	fmt.Println("width.Widen ------------------------------------------------------------")
	for _, r := range width.Widen.String(string(rs)) {
		p := width.LookupRune(r)
		fmt.Printf("%c: %s\n", r, p.Kind())
	}
}
