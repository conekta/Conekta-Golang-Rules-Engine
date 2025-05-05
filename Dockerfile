ARG WORKDIR_DEFAULT=/go/src/github.com/conekta/Conekta-Golang-Rules-Engine

# build stage
FROM golang:1.23-alpine as builder
LABEL "com.conekta.vendor"="Conekta"
LABEL "com.conekta.maintainer"=""
LABEL "version"="2022.0.1"

RUN apk update && apk add bash ca-certificates git openssh gcc g++ libc-dev librdkafka-dev pkgconf make curl
RUN mkdir -p -m 0600 ~/.ssh && ssh-keyscan github.com >> ~/.ssh/known_hosts

ARG WORKDIR_DEFAULT
WORKDIR $WORKDIR_DEFAULT

# Copy all the Code and stuff to compile everything
COPY . .

ENV GO111MODULE=on
RUN --mount=type=ssh git config --global url."ssh://git@github.com/conekta".insteadOf https://github.com/conekta \
    && git config --global --add safe.directory $WORKDIR_DEFAULT  \
    && go mod download -x
