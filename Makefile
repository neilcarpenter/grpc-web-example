export GOPATH := ${CURDIR}/backend/go

proto: proto-go proto-js

proto-go:
	protoc --go_out=plugins=grpc:${GOPATH}/src proto/*/*.proto

proto-js:
	protoc --js_out=import_style=commonjs,binary:${CURDIR}/web proto/*/*.proto
