package kata_return_negative_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestReturnNegative(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Return Negative Suite")
}
