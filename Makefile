export GOPATH := ${CURDIR}/backend/go

GRPC_WEB_REPO_PATH := /add/path/here
GRPC_WEB_PLUGIN_PATH := /add/path/here
CLOSURE_COMPILER_PATH := /add/path/here

all: proto compile-js

proto: proto-go proto-js

proto-go:
	protoc --go_out=plugins=grpc:${GOPATH}/src proto/*/*.proto

proto-js:
	mkdir -p ${CURDIR}/web/src/proto
	protoc --js_out=import_style=closure,binary:${CURDIR}/web/src/proto proto/*/*.proto
	protoc  -I=. --plugin=protoc-gen-grpc-web=$(GRPC_WEB_PLUGIN_PATH) \
		--grpc-web_out=out=./web/src/proto/echo.grpc.pb.js,mode=grpcweb:. proto/*/*.proto

compile-js:
	java \
		-jar ${CLOSURE_COMPILER_PATH} \
		--js ./web/src \
		--js $(GRPC_WEB_REPO_PATH)/javascript \
		--js $(GRPC_WEB_REPO_PATH)/third_party/closure-library \
		--js $(GRPC_WEB_REPO_PATH)/third_party/grpc/third_party/protobuf/js \
		--entry_point=goog:proto.grpc.web.research.EchoServiceClient \
		--dependency_mode=STRICT \
		--js_output_file ./web/dist/js/compiled.js
