build:
	go build -o bin/website

run: build
	./bin/website

test:
	go test -v ./... -count=1

install-wasmtime:
	curl https://wasmtime.dev/install.sh -sSf | bash

dev:
	t1 generate --watch --proxy="http://localhost:3020" --cmd="make up"

up: generate run

generate:
	t1 generate

connect:
	ssh senforsce_a@ssh-senforsce.alwaysdata.net

build-wasm:
	t1 generate
	GOOS=js GOARCH=wasm go build -o cmd/main.wasm .
	cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./static
	