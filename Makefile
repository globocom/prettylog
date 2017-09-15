binary_name=prettylog

.PHONY: default install

default:
	@go build -o $(binary_name)

install:
	@go build -o ${GOPATH}/bin/$(binary_name)