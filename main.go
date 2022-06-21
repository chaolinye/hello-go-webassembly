package main

import "fmt"
import "syscall/js"

func main() {
	// 打印 console
	fmt.Println("Hello go webassembly")
	// 操作 dom
	js.Global().Get("document").Call("getElementById", "content").Set("innerText", "Hello go webassembly")
	// 调用 js 方法
	js.Global().Get("window").Call("jsFun")
	ch := make(chan bool)
	// 提供方法给 js 调用, 只能调用一次
	js.Global().Set("goHello", js.FuncOf(func(this js.Value, args []js.Value) any {
		result := "js call wasm function, args: ("
		for i, arg := range args {
			if i != 0 {
				result += ", "
			}
			result += arg.String()
		}
		result += ")"
		ch <- true
		return result
	}))
	// 阻塞，避免 go 程序结束导致暴露的方法失效
	<-ch
}
