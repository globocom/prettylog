binary_name=prettylog

.PHONY: default setup install test

default:
	go build -o $(binary_name)

install:
	go build -o ${GOPATH}/bin/$(binary_name)

setup:
	go install github.com/onsi/ginkgo/ginkgo@v1.16.4

test:
	go run github.com/onsi/ginkgo/ginkgo@v1.16.4 -r .
