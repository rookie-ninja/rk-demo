.PHONY: all
all: fmt swag lint doctoc

.PHONY: lint
lint:
	@echo "[golangci-lint] Running golangci-lint..."
	@golangci-lint run 2>&1

.PHONY: fmt
fmt:
	@echo "[fmt] Formatting go project..."
	@gofmt -s -w . 2>&1

.PHONY: swag
swag:
	@echo "[swag] Running swag... (Looking up at internal/ and output to ./docs)"
	@swag init --parseInternal --propertyStrategy camelcase

.PHONY: doctoc
doctoc:
	@echo "[doctoc] Running doctoc..."
	@doctoc . 2>&1
