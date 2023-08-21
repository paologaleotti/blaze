# List of all directories containing entrypoints (cmd directories)
CMD_DIRS := $(wildcard cmd/*)

# Build all entrypoints and place artifacts in /bin
build:
	@mkdir -p bin
	@for dir in $(CMD_DIRS); do \
		echo "building $$(basename $$dir)..."; \
		go build -v -o bin/$$(basename $$dir) $$dir/main.go; \
	done

# Clean build artifacts
clean:
	rm -rf bin

.PHONY: build clean
