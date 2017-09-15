package prettifiers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPrettifiers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Prettifiers Suite")
}
