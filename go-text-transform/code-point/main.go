package main

import (
	"fmt"
	"log"
	"unicode"
	"unicode/utf8"

	"github.com/rivo/uniseg"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"golang.org/x/text/width"
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
	s1 := "é"
	fmt.Printf("%[1]q %+[1]q\n", s1)

	// 分解
	s1 = norm.NFD.String(s1)
	fmt.Printf("%[1]q %+[1]q\n", s1)

	// 合成
	s1 = norm.NFC.String(s1)
	fmt.Printf("%[1]q %+[1]q\n", s1)

	fmt.Println("互換等価性に基づいて分解・合成 --------------------------------------------------")
	s2 := "ゴ"
	fmt.Printf("%[1]q %+[1]q\n", s2)

	// 分解
	s2 = norm.NFKD.String(s2)
	fmt.Printf("%[1]q %+[1]q\n", s2)

	// 合成
	s2 = norm.NFKC.String(s2)
	fmt.Printf("%[1]q %+[1]q\n", s2)

	fmt.Println("コードポイントごとの変換 --------------------------------------------------")

	// カタカナであれば全角にする
	t1 := runes.If(runes.In(unicode.Katakana), width.Widen, nil)
	fmt.Println(t1.String("５ｱアAα"))

	fmt.Println("アクサンテギュなどを削除 --------------------------------------------------")
	removeMn := runes.Remove(runes.In(unicode.Mn))
	t2 := transform.Chain(norm.NFD, removeMn, norm.NFC)
	s3, _, err := transform.String(t2, "résumé")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s3)

	fmt.Println("アルファベットを大文字に変換 --------------------------------------------------")
	t3 := runes.Map(func(r rune) rune {
		if 'a' <= r && r <= 'z' {
			return r - 'a' + 'A'
		}
		return r
	})
	fmt.Println(t3.String("Hello, World"))
}
