GOVERSION=$(shell go version)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

PROJNAME=textAnalyzer
SRC=./src/main.go ./src/reader.go ./src/model.go

show-version:
	@echo $(GOVERSION)
	@echo $(GOOS)
	@echo $(GOARCH)

build:
	@echo build $(PROJNAME)
	@go build -o $(PROJNAME) $(SRC)

exec:
	@echo build and exec $(PROJNAME)
	@go build -o $(PROJNAME) $(SRC)
	@./$(PROJNAME) ./src/data/input.txt ./src/data/master.txt
