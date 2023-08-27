package main

import (
	"fmt"
	"unicode/utf8"

	"github.com/rivo/uniseg"
	"golang.org/x/text/unicode/norm"
)

func main() {
	for i, r := range "世界" {
		fmt.Printf("%d: %c ", i, r)
	}

	fmt.Println("\nエンコード --------------------------------------------------")
	buf := make([]byte, 3)
	n := utf8.EncodeRune(buf, '世')
	fmt.Printf("%v %q %d\n", buf, string(buf), n)

	fmt.Println("デコード --------------------------------------------------")
	b := []byte("世界")
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		fmt.Printf("%q: %d ", r, size)
		b = b[size:]
	}

	fmt.Println("\n書記素クラスタ分解 --------------------------------------------------")
	gr := uniseg.NewGraphemes("Cafe\u0301")
	for gr.Next() {
		fmt.Printf("%s %x ", gr.Str(), gr.Runes())
	}

	fmt.Println("\n正準等価性に基づいて分解・合成 --------------------------------------------------")
	s := "é"
	fmt.Printf("%[1]q %+[1]q\n", s)

	// 分解
	s = norm.NFD.String(s)
	fmt.Printf("%[1]q %+[1]q\n", s)

	// 合成
	s = norm.NFC.String(s)
	fmt.Printf("%[1]q %+[1]q\n", s)
}
