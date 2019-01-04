build:
	go build .

build-pi:
	GOARM=5 GOOS=linux GOARCH=arm go build -o comix-pi .

test:
	go test -v ./...

generate-proto:
	protoc -I cnc cnc/cnc.proto --go_out=plugins=grpc:cnc
