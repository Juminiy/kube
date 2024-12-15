arch=amd64
os=linux
DOCKERRMI=docker rmi netconn:$(os)_$(arch)
DOCKERBLD=docker buildx build \
          		--platform=$(os)/$(arch) \
          		--build-arg ARCH=$(arch) \
          		--build-arg OS=$(os) \
          		--build-arg GOPROXY=https://goproxy.cn,direct \
          		--build-arg HTTP_PROXY=http://192.168.3.37:7890 \
          		--build-arg HTTPS_PROXY=http://192.168.3.37:7890 \
          		--file Dockerfile \
          		--tag netconn:$(os)_$(arch) \
          		--debug \
          		--no-cache .

docker-amd64-linux:
	arch=amd64
	os=linux
	#$(DOCKERRMI)
	$(DOCKERBLD)

docker-amd64-windows:


docker-arm64-linux:


docker-arm64-darwin:


