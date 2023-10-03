.DEFAULT_GOAL := default
# HINT: The line above doesn't work on Make <= 3.80, used by older OSX releases.
# The '.PHONY' target will also be used for setting the default target.

binary_name=prettylog

.PHONY: default setup install test

default:
	@go build -o $(binary_name)

install:
	@go build -o ${GOPATH}/bin/$(binary_name)

setup:
	@go install github.com/onsi/ginkgo/ginkgo@v1.16.4

test:
	@go run github.com/onsi/ginkgo/ginkgo@v1.16.4 -r .

clean:
	@rm -f $(binary_name) ${GOPATH}/bin/$(binary_name)