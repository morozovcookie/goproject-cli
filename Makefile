GOBUILD_TARGET = goproject-cli-bin
GOFLAGS = CGO_ENABLED=0

.PHONY: gobuild clean

gobuild:
	$(GOFLAGS) go build -o $(GOBUILD_TARGET) cmd/goproject-cli/main.go

gotest:
	go test -v -cover -coverprofile coverage.out ./...

clean:
	rm $(DOT_TARGET)
	rm $(GOBUILD_TARGET)
