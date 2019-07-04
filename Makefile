REPO ?= docker.io/autonomy
GOLANG_IMAGE ?= golang:1.12.6

TAG := $(shell gitmeta image tag)

COMMON_ARGS := -f ./Dockerfile --build-arg GOLANG_IMAGE=$(GOLANG_IMAGE) .

all: enforce clean build image

.PHONY: enforce
enforce:
	@conform enforce

.PHONY: build
build:
	@DOCKER_BUILDKIT=1 docker build \
		-t gitmeta/$@:$(TAG) \
		--target $@ \
		$(COMMON_ARGS)
	@docker run --rm -v $(PWD)/build:/build gitmeta/$@:$(TAG) cp /gitmeta-linux-amd64 /build
	@docker run --rm -v $(PWD)/build:/build gitmeta/$@:$(TAG) cp /gitmeta-darwin-amd64 /build

.PHONY: image
image: build
	@DOCKER_BUILDKIT=1 docker build \
		-t autonomy/gitmeta:$(TAG) \
		--target=$@ \
		$(COMMON_ARGS)

clean:
	rm -rf ./build
