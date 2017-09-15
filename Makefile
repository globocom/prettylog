binary_name=prettylog

.PHONY: default install test

default:
	@go build -o $(binary_name)

install:
	@go build -o ${GOPATH}/bin/$(binary_name)

test:
	@ginkgo -r .