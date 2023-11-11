package main

/*
#include <math.h>
#cgo LDFLAGS: -lm
*/
import "C"
import "fmt"

func main() {
	// 2が出力される。
	fmt.Println(C.sqrt(4))
}
