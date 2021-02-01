build:build_linux
	go build -o bin/ws ./

build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/gateway ./

build_mac:
	go build -o bin/ws ./