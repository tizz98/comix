build:
	go build .

build-pi:
	GOARM=5 GOOS=linux GOARCH=arm go build -o comix-pi .

test:
	go test -v -cover -race ./...

generate-proto:
	protoc -I cnc cnc/cnc.proto --go_out=plugins=grpc:cnc

build-version:
	GOARM=5 GOOS=linux GOARCH=arm go build -o dist/${PI_VERSION} .
	shasum -a 256 comix-pi | cut -d " " -f1 > dist/${PI_VERSION}.sha
