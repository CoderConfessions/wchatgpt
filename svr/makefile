# 项目名称
PROJECT_NAME=wchatgpt

# 源文件目录
SRC_DIR=./

# 可执行文件输出目录
BUILD_DIR=./build

# 编译命令
GO_BUILD=go build

# 需要编译的文件列表
# 如果有新的文件，需要手动更新该列表
SOURCES=$(shell find $(SRC_DIR) -type f -name '*.go')

# 可执行文件名称
EXECUTABLE=$(BUILD_DIR)/main

# 配置文件
CONFIG_FILE=./configuration.json

# 编译目标（默认）
build: $(EXECUTABLE)

# 编译可执行文件
$(EXECUTABLE): $(SOURCES)
	$(GO_BUILD) -o $(EXECUTABLE) -ldflags="-s -w" -v ./main.go

# 运行
run: $(EXECUTABLE)
	./main --config-file=$(CONFIG_FILE)

# 清除生成的文件
clean:
	rm -rf $(BUILD_DIR)

.PHONY: build run clean
