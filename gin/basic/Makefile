.PHONY: all
all: fmt lint doctoc

.PHONY: lint
lint:
	@echo "[golangci-lint] Running golangci-lint..."
	@golangci-lint run 2>&1

.PHONY: fmt
fmt:
	@echo "[fmt] Formatting go project..."
	@gofmt -s -w . 2>&1

.PHONY: doctoc
doctoc:
	@echo "[doctoc] Running doctoc..."
	@doctoc . 2>&1
