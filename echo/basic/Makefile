.PHONY: all
all: fmt lint doctoc swag

.PHONY: lint
lint:
	@echo "[golangci-lint] Running golangci-lint..."
	@golangci-lint run 2>&1
	@echo "------------------------------------[Done]"

.PHONY: fmt
fmt:
	@echo "[fmt] Formatting go project..."
	@gofmt -s -w . 2>&1
	@echo "------------------------------------[Done]"

.PHONY: doctoc
doctoc:
	@echo "[doctoc] Running doctoc..."
	@doctoc . 2>&1
	@echo "------------------------------------[Done]"

.PHONY: swag
swag:
	@echo "[swag] Running swag... (Looking up at internal/ and output to ./docs)"
	@swag init --parseInternal --propertyStrategy camelcase
	@echo "------------------------------------[Done]"
