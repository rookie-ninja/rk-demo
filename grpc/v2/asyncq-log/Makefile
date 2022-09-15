.PHONY: all
all: fmt buf

.PHONY: fmt
fmt:
	@echo "[fmt] Format go project..."
	@gofmt -s -w . 2>&1
	@echo "------------------------------------[Done]"

.PHONY: buf
buf:
	@echo "[buf] Running buf..."
	@buf generate --path api/v1
