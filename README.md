# gRPC-web example - Echo / Reverse services

Simple demonstration of gRPC-web, consisting of 4 services:
1. Go server implementing Echo gRPC service
2. NodeJS server implementing Reverse gRPC service
3. Static HTML/JS app implementing Echo / Reverse gRPC clients
4. Envoy proxy

## Setup

Ensure you have:
- [grpc/grpc-web](https://github.com/grpc/grpc-web) repo cloned locally, and followed the [installation instructions](https://github.com/grpc/grpc-web/blob/master/INSTALL.md).
- Make sure you have compiled the grpc-web protoc plugin (run `make plugin` from `grpc/grpc-web` repo root), and have `protoc-gen-go` plugin installed and available on your path.
- Downloaded [closure compiler](https://github.com/google/closure-compiler).
- Updated references for `GRPC_WEB_REPO_PATH` (path to grpc/grpc-web repo locally), `GRPC_WEB_PLUGIN_PATH` (path to grpc-web protoc plugin, compiled during the grpc/grpc-web installation steps) and `CLOSURE_COMPILER_PATH` (path to downloaded closure compiler .jar) inside `Makefile`.

## Run the example

From repo root:
```
make all
docker-compose up
```
Then visit `http://localhost:8000`.
