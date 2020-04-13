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

## Testing
This project uses https://github.com/agnivade/wasmbrowsertest  
Everything for running should be setup by `go generate ./...` (in wasm-folder) already.  
If you want to run this from shell, please make sure to set  
```bash
export GOOS=js
export GOARCH=wasm
```

## Run the backend
```bash
go run main.go # make sure to use go1.14 for that
```
