GO=@go
SOURCE=cmd/gom/main.go
DEST=cmd/gom/gom
BASE_PKG=github.com/bakku/gom

gom:
	$(GO) build -o $(DEST) $(SOURCE)

test:
	$(GO) test $(BASE_PKG)/util
	$(GO) test $(BASE_PKG)/commands

fmt:
	$(GO) fmt $(BASE_PKG)/util
	$(GO) fmt $(BASE_PKG)/commands
	$(GO) fmt $(BASE_PKG)/mocks
	$(GO) fmt $(BASE_PKG)/cmd/gom
