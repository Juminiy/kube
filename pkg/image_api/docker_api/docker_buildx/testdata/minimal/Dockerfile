# syntax=docker/dockerfile:1
FROM --platform=$BUILDPLATFORM golang:1.23-alpine AS build
ARG TARGETOS
ARG TARGETARCH
WORKDIR /app
COPY main.go .

ENV GOPROXY=https://goproxy.cn,direct
ENV CGO_ENABLED=0
ENV GO111MODULE=on

RUN go mod init minimal
RUN go mod tidy
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o minimal .

FROM alpine:latest

COPY --from=build /app/minimal /minimal
ENTRYPOINT ["/minimal"]