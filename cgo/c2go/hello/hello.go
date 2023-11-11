package hello

// extern void hello();
import "C"
import "fmt"

//export goHello
func goHello() {
	fmt.Println("hello")
}

// Hello main.goから呼び出される関数。
func Hello() {
	// C言語で書かれたhello関数を呼ぶ。
	C.hello()
}
