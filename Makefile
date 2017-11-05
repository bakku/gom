GO=go
SOURCE=cmd/gom/main.go
DEST=cmd/gom/gom

migrator:
	$(GO) build -o $(DEST) $(SOURCE)

test:
	cd util && $(GO) test && cd ..
