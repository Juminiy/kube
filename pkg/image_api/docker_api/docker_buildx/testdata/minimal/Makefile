output_tar:
	docker buildx build \
      --builder multiple-x \
      --debug \
      --file Dockerfile \
      --platform linux/amd64,linux/arm64 \
      --tag minimal:latest \
      --output type=oci,dest=./minimal.tar .

output_fs:
	docker buildx build \
          --builder multiple-x \
          --debug \
          --file Dockerfile \
          --platform linux/amd64,linux/arm64 \
          --tag minimal:latest \
          --output type=local,dest=./localfs .

output_docker:
	docker buildx build \
          --builder multiple-x \
          --debug \
          --file Dockerfile \
          --platform linux/amd64,linux/arm64 \
          --tag minimal:compression-v20 \
          --output type=docker,name=minimal:compression-v20,tar=true,compression=zstd,cmpression-level=20 .