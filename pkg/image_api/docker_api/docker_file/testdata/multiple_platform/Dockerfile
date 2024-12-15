# syntax=docker/dockerfile:1
FROM --platform=$BUILDPLATFORM golang:alpine AS build
ARG TARGETOS
ARG TARGETARCH
WORKDIR /app
ADD https://github.com/dvdksn/buildme.git#eb6279e0ad8a10003718656c6867539bd9426ad8 .
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o server .

FROM alpine
COPY --from=build /app/server /server
ENTRYPOINT ["/server"]