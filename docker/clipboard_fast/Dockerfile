# syntax=docker/dockerfile:1
FROM --platform=$BUILDPLATFORM golang:1.23-alpine AS build
ARG TARGETOS
ARG TARGETARCH
WORKDIR /app
COPY cmd/clipboard/clipboard_fast .

ENV GOPROXY=https://goproxy.cn,direct
ENV CGO_ENABLED=0
ENV GO111MODULE=on

RUN go mod init clipboard_fast
RUN go mod tidy
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o clipboard_fast .

FROM alpine:latest

COPY --from=build /app/clipboard_fast /clipboard_fast
ENTRYPOINT ["/clipboard_fast"]