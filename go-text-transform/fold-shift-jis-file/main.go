package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"golang.org/x/text/width"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("A filename is required.")
	}

	if err := foldShiftJISFile(args[0]); err != nil {
		log.Fatal(err)
	}
}

// Shift_JISのファイルの全角英数などは半角に、半角カナなどは全角にする
func foldShiftJISFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Shift_JISからUTF-8に変換してから
	// 全角英数などは半角に、半角カナなどは全角にする
	dec := japanese.ShiftJIS.NewDecoder()
	t := transform.Chain(dec, width.Fold)
	r := transform.NewReader(f, t)

	s := bufio.NewScanner(r)
	for s.Scan() {
		fmt.Println(s.Text())
	}

	if err := s.Err(); err != nil {
		return err
	}

	return nil
}
