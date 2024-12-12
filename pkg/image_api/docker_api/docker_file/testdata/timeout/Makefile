linux/amd64:
	docker build --build-arg PLATFORM=$@ -t timeoutcli:latest -f Dockerfile .

linux/arm64:
	docker build --build-arg PLATFORM=$@ -t timeoutcli:latest -f Dockerfile .

darwin/arm64:
	docker build --build-arg PLATFORM=$@ -t timeoutcli:latest -f Dockerfile .

tar:
	chmod 777 Dockerfile go.mod main.go
	tar -czvf timeout.tar Dockerfile go.mod main.go
	chmod 777 timeout.tar