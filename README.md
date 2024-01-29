# Go Wasm

Working with Go Wasm to run Go code on the browser.

## Running

First prepare `wasm_exec.js` by running:

```bash
make gen_js GO_PATH=$(go env GOPATH)
```

Then run the server:

```bash
cd server
go run .
```

> Access the served site at `http://localhost:5000/`.

## Development

To generate new `wasm` code, run:

```bash
make build
# optionally optimize the wasm file
make optimize
```
