FROM golang:1.20.0-alpine

# ARG TARGET_DIR=/cmd/server

# ARG GOBIN=/.bin

# FROM golang:1.20.0-alpine

# ARG TARGET_DIR
# ARG GOBIN

# RUN apk add --update make git musl-dev gcc libc-dev binutils-gold

# ENV GO111MODULE=on
# ENV GOPATH=/go
# WORKDIR /go/src/github.com/testovoe/

# COPY go.mod go.mod
# COPY go.sum go.sum

# COPY . .

# RUN GOBIN=$GOBIN make install-tools
# RUN TARGET_DIR=$TARGET_DIR make build

# ENTRYPOINT make watch