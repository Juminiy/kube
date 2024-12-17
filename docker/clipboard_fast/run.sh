#!/bin/bash

docker stop clipboardfast
docker rm clipboardfast
docker run -dit \
  --cpus 1 \
	--memory 256MB \
	--name clipboardfast \
	--restart=always \
	-p 8081:8081 \
	clipboard_fast:latest