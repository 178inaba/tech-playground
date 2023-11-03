package main

import (
	"bytes"
	"compress/gzip"
	"io"
	"syscall/js"
	"time"
)

func main() {
	js.Global().Set("compress", compressFunc)
	select {}
}

func Compress(src io.Reader) (io.Reader, error) {
	var buf bytes.Buffer

	zw := gzip.NewWriter(&buf)
	zw.ModTime = time.Now()

	if _, err := io.Copy(zw, src); err != nil {
		return nil, err
	}

	if err := zw.Close(); err != nil {
		return nil, err
	}

	return &buf, nil
}

func newUint8Array(size int) js.Value {
	ua := js.Global().Get("Uint8Array")
	return ua.New(size)
}

var compressFunc = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	jsSrc := args[0] // Uint8Arrayを受け取る。
	srcLen := jsSrc.Get("length").Int()
	srcBytes := make([]byte, srcLen)
	js.CopyBytesToGo(srcBytes, jsSrc) // JavaScript側のファイルデータをGo側にコピーする。

	src := bytes.NewReader(srcBytes)

	r, err := Compress(src) // 圧縮処理の実行。
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		panic(err)
	}
	ua := newUint8Array(buf.Len())
	js.CopyBytesToJS(ua, buf.Bytes()) // Go側で圧縮されたファイルデータをJavaScript側にコピーする。
	return ua
})
