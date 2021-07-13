package main

import (
	"syscall/js"
)

// main is the entrypoint for the wasm
func main() {

	c := make(chan struct{}, 0)
	js.Global().Set("generateDockerfile", js.FuncOf(generateDockerfile))
	<-c
}
