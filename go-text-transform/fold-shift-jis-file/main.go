package main

import "os"

func main() {
}

// Shift_JISのファイルの全角英数などは半角に、半角カナなどは全角にする
func foldShiftJISFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// TODO
}
