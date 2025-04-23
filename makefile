# Имя приложения
APP_NAME := postsl

# Версия приложения
VERSION := 0.1.0

# Целевые ОС
OS_LIST := linux darwin freebsd

# Целевые архитектуры
ARCH_LIST := amd64 arm64 386

# Директория для бинарников и архивов
DIST_DIR := dist
BUILD_DIR := build

# Флаги сборки
LDFLAGS := -ldflags="-s -w -X main.Version=$(VERSION)"

# Цвета для вывода
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: all clean build release test

all: test build

## Собрать для текущей платформы
build: ## Собрать для текущей платформы
	@echo "${YELLOW}Building for current platform...${RESET}"
	go build $(LDFLAGS) -o $(APP_NAME) ./main.go

## Собрать релизные версии для всех платформ и архитектур
release: clean ## Собрать релизные версии для всех платформ и архитектур
	@echo "${YELLOW}Building releases...${RESET}"
	@mkdir -p $(BUILD_DIR) $(DIST_DIR)
	@for os in $(OS_LIST); do \
		for arch in $(ARCH_LIST); do \
			if [ "$$os" = "darwin" ] && [ "$$arch" = "386" ]; then \
				continue; \
			fi; \
			OUTPUT=$(BUILD_DIR)/$(APP_NAME); \
			echo "${GREEN}Building $$os/$$arch...${RESET}"; \
			GOOS=$$os GOARCH=$$arch go build $(LDFLAGS) -o $$OUTPUT ./main.go; \
			tar -czf $(DIST_DIR)/$(APP_NAME)-$$os-$$arch.tar.gz -C $(BUILD_DIR) $(APP_NAME); \
			rm -f $(BUILD_DIR)/$(APP_NAME); \
		done; \
	done
	@rmdir $(BUILD_DIR)
	@echo "${GREEN}Build complete!${RESET}"
	@ls -lh $(DIST_DIR)/*

## Очистить собранные файлы
clean: ## Очистить собранные файлы
	@echo "${YELLOW}Cleaning...${RESET}"
	@rm -rf $(DIST_DIR) $(BUILD_DIR) $(APP_NAME)

## Запустить тесты
test: ## Запустить тесты
	@echo "${YELLOW}Running tests...${RESET}"
	go test ./...

## Показать помощь
help:  ## Показать помощь
	@echo "Usage:"
	@echo "  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}"
	@echo ""
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  ${GREEN}%-15s${RESET} %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort
