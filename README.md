# WASM-Bootstrap with tinygo

## Installation
1. [Install tinygo](https://tinygo.org/getting-started/)
2. Install/Use go1.13 (see https://github.com/tinygo-org/tinygo/issues/930)

## Build
For building/refreshing the WASM part of the frontend, run
```bash
export GOROOT=`go1.13 env GOROOT` # see https://github.com/tinygo-org/tinygo/issues/930#issuecomment-605597556
go generate ./...
```

## Run the backend
```bash
go run main.go # make sure to use go1.14 for that
```
