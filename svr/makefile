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
	@if docker inspect $(DOCKER_IMAGE):$(DOCKER_TAG) >/dev/null 2>&1; then \
        docker rmi $(DOCKER_IMAGE):$(DOCKER_TAG); \
    else \
        echo "$(DOCKER_IMAGE):$(DOCKER_TAG) does not exist"; \
    fi

cbuild: clean build

crun: clean run

.PHONY: build run clean
