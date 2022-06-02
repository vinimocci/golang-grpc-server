GOBUILD   = go build -v 
OS        = linux
ARCH      = amd64 
APPNAME   = golang-grpc-server
VERSION   = 0.1.0

## mudar esse caminho para criar o .pb.go na pasta correta
compile-proto:
	protoc --proto_path=${GOPATH}/src/test --proto_path=. --go_out=plugins=grpc:${GOPATH}/src pb/user.proto

clean:
	@go clean -i -x ./...
	@rm -rf tmp

build:
	@GOOS=$(OS) GOARCH=$(ARCH) $(GOBUILD) -o tmp/$(APPNAME) *.go

tests:
	@go clean -testcache
	@go test -v test/*.go

tag:
	@git tag v$(VERSION) && git push --tags || :
