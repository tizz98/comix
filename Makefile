build:
	go build .

build-pi:
	GOARM=5 GOOS=linux GOARCH=arm go build -o comix-pi .
