.PHONY:build test
build:
	go build -o upc

clean:
	rm upc

test:
	go test .

arm:
	mkdir -p build
	env GOOS=linux GOARCH=arm GOARM=5 go build -o build/upc-arm

mac:
	mkdir -p build
	env GOOS=darwin GOARCH=amd64 go build -o build/upc-osx