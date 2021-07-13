all:
	gofmt -s -w .
	GOOS=js GOARCH=wasm go build -o chef.wasm; mv chef.wasm docs/
