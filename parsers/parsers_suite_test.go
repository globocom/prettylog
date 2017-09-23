package parsers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestParsers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Parsers Suite")
}
