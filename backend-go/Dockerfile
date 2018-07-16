FROM golang:alpine

ADD go/src /go/src

RUN apk update
RUN apk add ngrep
RUN apk add --no-cache git && cd /go/src/server && go get -v ./...
RUN cd /go/src/server && go build ./
CMD ["/go/bin/server"]
