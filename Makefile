.PHONY: build test clean static_analysis lint vet fmt chkfmt

GO          = go
GO_BUILD    = $(GO) build
GO_FORMAT   = $(GO) fmt
GOFMT       = gofmt
GOLINT		= staticcheck
GO_LIST     = $(GO) list
GO_TEST     = $(GO) test -v
GO_VET      = $(GO) vet
GO_LDFLAGS  = -ldflags="-s -w"
GOOS        = darwin

EXECUTABLES = bin/main
TARGETS     = $(EXECUTABLES)
GO_PKGROOT  = ./...
GO_PACKAGES = $(shell $(GO_LIST) $(GO_PKGROOT) | grep -v vendor)

build: $(TARGETS)
test:
	env GOOS=$(GOOS) $(GO_TEST) $(GO_PKGROOT)
clean:
	rm -rf $(TARGETS) ./vendor Gopkg.lock

static_analysis: chkfmt lint vet

lint:
	$(GOLINT) $(GO_PKGROOT)
vet:
	$(GO_VET) $(GO_PACKAGES)
fmt:
	$(GO_FORMAT) $(GO_PKGROOT)
chkfmt:
	bash -c "diff -u <(echo -n) <($(GOFMT) -d $(shell git ls-files | grep ".go$$"))"

bin/main:
	env GO111MODULE=on GOOS=$(GOOS) $(GO_BUILD) $(GO_LDFLAGS) -o $@