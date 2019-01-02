GOVERSION=$(shell go version)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

PROJNAME=textAnalyzer

show-version:
	@echo $(GOVERSION)
	@echo $(GOOS)
	@echo $(GOARCH)

build:
	@echo build $(PROJNAME)
	@go build -o $(PROJNAME) ./src/main.go ./src/reader.go
