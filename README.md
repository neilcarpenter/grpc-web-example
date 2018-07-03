# gRPC-web research

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
