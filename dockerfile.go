package main

import (
	"fmt"
	"github.com/autamus/chef/container"
	"strings"
	"syscall/js"
)

// returnResult back to the browser, in the innerHTML of the result element
func returnResult(output string, divid string) {
	js.Global().Get("document").
		Call("getElementById", divid).
		Set("innerHTML", output)
}

// generateDockerfile is linked with the JavaScript function of the same name.
func generateDockerfile(this js.Value, val []js.Value) interface{} {

	fmt.Println("The packages are:", val)

	// Split single string into packages
	packages := strings.Split(val[0].String(), ",")
	dockerfile := container.Dockerfile(packages, false)

	// Replace newlines with <br> for browser
	dockerfile = strings.ReplaceAll(dockerfile, "\n", "<br>")

	// Send result back to browser, key is div id, content is string
	returnResult(dockerfile, "dockerfile")
	return nil
}
