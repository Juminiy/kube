ARG PLATFORM="linux/amd64"
ARG BINARYCLI="timeoutcli"

FROM --platform=${PLATFORM} golang:1.23-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY . .

RUN go mod tidy \
#    && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest \
#    && go install golang.org/x/vuln/cmd/govulncheck@latest \
#    && govulncheck ./... \
#    && golangci-lint run \
#    && go test ./... \
    && go build -o ${BINARYCLI} main.go

FROM --platform=${PLATFORM} alpine:latest

WORKDIR /app

COPY --from=builder /build/${BINARYCLI} /app/

ENTRYPOINT ["/app/timeoutcli"]