package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	// window Objectを取得する。
	window := js.Global()
	// document Objectを取得する。
	document := window.Get("document")
	// bodyのDOMを取得する。
	body := document.Get("body")
	// buttonのDOMを作成する。
	btn := document.Call("createElement", "button")
	// buttonに表示する文字を設定する。
	btn.Set("textContent", "click me!")
	// buttonにclickのEventListenerを設定する。
	btn.Call("addEventListener", "click", js.FuncOf(func(js.Value, []js.Value) interface{} {
		fmt.Println("Hello, WebAssembly!")
		return nil
	}))
	// buttonをbody内に追加する。
	body.Call("appendChild", btn)
	// プログラムが終了しないようにするため待機する。
	select {}
}
