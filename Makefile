
GO_PKG=github.com/aneeshkp/silicon-pod-operator
export CGO_ENABLED:=0
export GO111MODULE:=on
export GOPROXY?=https://proxy.golang.org/


	
 

.PHONY: format
.PHONY: vet 
.PHONY: tidy
.PHONY: clean
.PHONY: build
.PHONY: all

all: clean vet format build

format:

		go fmt ./...


vet:
		go vet ./...


tidy:
		go mod tidy -v


clean:
		rm -rf build


build:  
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -o build/_output/bin/silicon-pod-operator github.com/aneeshkp/silicon-pod-operator/cmd/manager


 