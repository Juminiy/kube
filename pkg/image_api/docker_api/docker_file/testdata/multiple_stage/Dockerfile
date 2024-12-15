# syntax=docker/dockerfile:1
ARG ARCH
ARG OS

FROM --platform=${OS}/${ARCH} golang:1.23-alpine AS builder
WORKDIR /src
LABEL authors="chisato"

ARG ARCH
ARG OS
ARG GOPROXY

ENV GOARCH=${ARCH}
ENV GOOS=${OS}
ENV GOPROXY=${GOPROXY}
ENV GO111MODULE=on
ENV CGO_ENABLED=0

COPY main.go ./main.go
RUN go mod init netconn
RUN go mod tidy
RUN go build -o /bin/netconn ./main.go

FROM --platform=${OS}/${ARCH} alpine:latest
WORKDIR /bin
LABEL authors="chisato"

COPY --from=builder /bin/netconn /bin/netconn
ENTRYPOINT ["/bin/netconn"]