APP?=go-kafka-libs-comparement
BIN?=./bin/$(APP)
GO?=go
GO_ENVS?=CGO_ENABLED=0 GOOS=linux GOARCH=amd64
LOCAL_BIN:=$(CURDIR)/bin

.PHONY: build
build:
	$(GO_ENVS) $(GO) build -o $(BIN) ./cmd/$(APP)

.PHONY: run
run: build
	$(BIN)
