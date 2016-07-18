.PHONY: run

# Golang Flags
GOPATH ?= $(GOPATH:):./vendor
GOFLAGS ?= $(GOFLAGS:)
GO=go

.prebuild:
	@echo Preparation for running go code
	go get $(GOFLAGS) github.com/Masterminds/glide

	touch $@

run: .prebuild
	$(GO) run $(GOFLAGS) $(GO_LINKER_FLAGS) *.go

