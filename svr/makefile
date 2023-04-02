PROJECT_NAME=wchatgpt
SRC_DIR=./
BUILD_DIR=./build
# go compile 
GO_BUILD=go build
# go source file
SOURCES=$(find $(SRC_DIR) -type f -name '*.go')

# bin name
EXECUTABLE=$(BUILD_DIR)/main

# configuration
CONFIG_FILE=./configuration.json

# DOCKER
DOCKER_IMAGE=wchatgpt
DOCKER_TAG=latest
DOCKERFILE_PATH=./Dockerfile

docker-build:
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) -f $(DOCKERFILE_PATH) .

docker-run:docker-build
	docker run -it -rm -p 8080:8080 $(DOCKER_IMAGE):$(DOCKER_TAG)

build: $(EXECUTABLE)

$(EXECUTABLE): $(SOURCES)
	$(GO_BUILD) -o $(EXECUTABLE) -ldflags="-s -w" -v ./main.go

run: $(EXECUTABLE)
	$(BUILD_DIR)/main --config-file=$(CONFIG_FILE)


clean:
	rm -rf $(BUILD_DIR)
	docker rmi $(DOCKER_IMAGE):$(DOCKER_TAG)

.PHONY: build run clean
