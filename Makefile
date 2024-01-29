include 
gen_js:
	@cp ${GO_PATH}/misc/wasm/wasm_exec.js ./dist/wasm_exec.js

build:
	@GOOS=js GOARCH=wasm go build -o main.wasm main.go
	@rm -rf ./dist/main.wasm
	@cp ./main.wasm ./dist/main.wasm
	@rm -rf ./main.wasm

optimize: build
	@wasm-opt ./dist/main.wasm --enable-bulk-memory -Oz -o ./dist/main.wasm

