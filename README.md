# Hello Go WebAssembly

该工程使用 Go 编写 WebAssembly 程序，实现 wasm 与 js 的交互

## 构建命令

```bash
GOOS=js GOARCH=wasm go build -o assets/hello.wasm
```

## References

- [WebAssembly: Introduction to WebAssembly using Go](https://golangbot.com/webassembly-using-go/)
