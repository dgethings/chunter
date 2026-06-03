.PHONY: all lsp clean release snapshot release-dry-run

TS ?= tree-sitter
TS_DIR := grammars/tree-sitter-cisco_ios
TS_GRAMMAR := $(TS_DIR)/grammar.js
TS_STAMP := .ts-gen-stamp
BIN := bin/chunter

SRCS := $(wildcard main.go cmd/*.go internal/**/*.go)

VERSION := $(shell svu current 2>/dev/null || echo "0.0.0")
NEXT    := $(shell svu next)

all: $(BIN)

$(TS_STAMP): $(TS_GRAMMAR)
	cd $(TS_DIR) && $(TS) generate grammar.js
	@touch $@

$(BIN): $(TS_STAMP) $(SRCS)
	go clean -cache
	CGO_ENABLED=1 go build -o $(BIN) .

lsp:
	CGO_ENABLED=1 go build -o $(BIN) .

clean:
	go clean -cache
	rm -f $(BIN) $(TS_STAMP)
	$(MAKE) -C $(TS_DIR) clean

release:
	@if git tag -l "$(NEXT)" | grep -q "$(NEXT)"; then \
		echo "Tag $(NEXT) already exists, running goreleaser only"; \
	elif [ "$(NEXT)" = "$(VERSION)" ]; then \
		echo "No new conventional commits since $(VERSION). Commit feat:/fix: changes first."; \
		exit 1; \
	else \
		echo "Releasing $(NEXT) (current: $(VERSION))"; \
		git tag -a "$(NEXT)" -m "Release $(NEXT)"; \
		git push origin "$(NEXT)"; \
	fi
	goreleaser release --clean

snapshot:
	goreleaser release --snapshot --clean

release-dry-run:
	@echo "Current: $(VERSION) -> Next: $(NEXT)"
	goreleaser release --snapshot --clean
