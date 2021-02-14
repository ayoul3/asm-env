NAME = asm-env
IMAGE_NAME = asm-env
IMAGE_PREFIX = ayoul3
BUILD=go build -ldflags="-s -w"
export GO111MODULE=on

test:
	go test -v ./... -cover

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(BUILD) -o asm-env main.go

docker: build
	docker build --no-cache -t $(IMAGE_PREFIX)/$(IMAGE_NAME):latest .

push: docker
	docker push $(IMAGE_PREFIX)/$(IMAGE_NAME):latest

