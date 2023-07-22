package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/text/encoding/japanese"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("A filename is required.")
	}

	if err := printCSV(args[0]); err != nil {
		log.Fatal(err)
	}
}

func printCSV(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Shift_JISとして読み込む
	dec := japanese.ShiftJIS.NewDecoder()
	cr := csv.NewReader(dec.Reader(f))
	for {
		rec, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		// UTF-8に変換されているので表示しても
		// 文字化けしない
		fmt.Println(rec)
	}

	return nil
}
