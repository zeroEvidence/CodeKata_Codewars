package kata_multiply_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMultiply(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Multiply Suite")
}
