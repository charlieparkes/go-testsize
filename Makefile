INSTALL_DIR?=$(HOME)/.local/bin

define announce
	@printf '\033[36m⌁ %s\033[0m\n' "$(1)"
endef

.PHONY: default
default: all

.PHONY: all
all: lint test

.PHONY: lint
lint:
	$(call announce,$@)
	golangci-lint run --fix ./...

.PHONY: test
test:
	$(call announce,$@)
	go test ./...
