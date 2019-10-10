binary_name=prettylog

.PHONY: default install test

default:
	@go build -o $(binary_name)

install:
	@go build -o ${GOPATH}/bin/$(binary_name)

test:
ifeq (, $(shell which ginkgo))
	go get github.com/onsi/ginkgo/ginkgo
endif
	@ginkgo -r .
