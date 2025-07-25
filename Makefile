.PHONY: run build clean help

# Default target
run: build
	./yamato-mcp -http :8080

# Build the binary
build:
	go build -o yamato-mcp main.go

# Clean build artifacts
clean:
	rm -f yamato-mcp

# Show help
help:
	@echo "Available targets:"
	@echo "  run    - Build and run the server on port 8080"
	@echo "  build  - Build the binary"
	@echo "  clean  - Remove build artifacts"
	@echo "  help   - Show this help message"
	@echo ""
	@echo "After running 'make run', connect MCP Inspector:"
	@echo "  1. Run: npx @modelcontextprotocol/inspector"
	@echo "  2. Use URL: http://localhost:8080/mcp"
	@echo "  3. Select transport: Streamable HTTP"
	@echo "  4. Leave proxy token empty"
	@echo "  5. Tools > List Tools to get started"