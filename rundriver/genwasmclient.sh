#!/usr/bin/env bash


rm wasmclient.wasm

echo "GOOS=js GOARCH=wasm go build -o wasmclient.wasm"
GOOS=js GOARCH=wasm go build -o wasmclient.wasm wasmclient.go

echo "move files"
mv wasmclient.wasm ./clientdata/

##########################

rm wasmclientgl.wasm

echo "GOOS=js GOARCH=wasm go build -o wasmclientgl.wasm"
GOOS=js GOARCH=wasm go build -o wasmclientgl.wasm wasmclientgl.go

echo "move files"
mv wasmclientgl.wasm ./clientdata/
