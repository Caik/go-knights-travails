# Usage:
# make                      # build go source code, build go source code for windows
# make test                 # run unit tests
# make build                # build go source code
# make build_windows        # build go source code for windows

all: test build build_windows
.PHONY: all test build build_windows

test:
	@echo ""
	@echo "############################"
	@echo "##   Running unit tests   ##"
	@echo "############################"
	@echo ""
	@go clean -testcache && go test ./...

build: ./cmd/knights-travails/main.go
	@echo ""
	@echo "############################"
	@echo "##   Building for Linux   ##"
	@echo "############################"
	@echo ""
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static" -s -w' -o dist/knights-travails $<

build_windows: ./cmd/knights-travails/main.go
	@echo ""
	@echo "############################"
	@echo "##  Building for Windows  ##"
	@echo "############################"
	@echo ""
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -ldflags '-extldflags "-static" -s -w' -o dist/knights-travails.exe $<
