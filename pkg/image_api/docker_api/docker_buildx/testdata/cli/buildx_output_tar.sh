docker buildx build \
  --builder multiple-x \
  --debug \
  --file Dockerfile \
  --platform linux/amd64 \
  --tag minimal:latest \
  --output