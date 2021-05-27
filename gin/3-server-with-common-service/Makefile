.PHONY: all
all: readme fmt swag

.PHONY: fmt
fmt:
	@echo "format go project..."
	@gofmt -s -w . 2>&1

.PHONY: readme
readme:
	@echo "running doctoc..."
	@doctoc . 2>&1

.PHONY: swag
swag:
	@echo "running swag..."
	@swag init --generalInfo main.go --propertyStrategy camelcase

