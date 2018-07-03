# gRPC-web example - Echo service

Simple demonstration of gRPC-web, consisting of 3 services:
1. Go server implementing Echo gRPC service
2. Static HTML/JS app implementing Echo gRPC client
3. Envoy proxy

## Setup

Ensure you have:
- [grpc/grpc-web](https://github.com/grpc/grpc-web) repo cloned locally, and followed the [installation instructions](https://github.com/grpc/grpc-web/blob/master/INSTALL.md).
- Downloaded [closure compiler](https://github.com/google/closure-compiler).
- Then update relevant references for `GRPC_WEB_REPO_PATH`, `GRPC_WEB_PLUGIN_PATH` and `CLOSURE_COMPILER_PATH` inside `Makefile`.

## Run the example

From repo root:
```
make all
docker-compose up
```
Then visit `http://localhost:8000`.
