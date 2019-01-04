build: generate
	go build .

build-pi: generate
	GOARM=5 GOOS=linux GOARCH=arm go build -o comix-pi .

test: generate
	go test -v ./...

generate:
	protoc -I cnc cnc/cnc.proto --go_out=plugins=grpc:cnc
