build:
	go build -o example ./main.go

run: build
	./example ./config.json /

test:
	go test -v -race $(shell go list ./... | grep -v /vendor/)

