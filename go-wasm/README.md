# Go Wasm

```console
$ wget 'https://raw.githubusercontent.com/golang/go/master/misc/wasm/wasm_exec.js'
$ GOOS=js GOARCH=wasm go build -o main.wasm
$ serve -a :8000
```

Access http://localhost:8000/ to check.
