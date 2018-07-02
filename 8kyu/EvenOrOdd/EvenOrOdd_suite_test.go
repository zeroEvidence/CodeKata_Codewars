package kata_even_or_odd_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEvenOrOdd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "EvenOrOdd Suite")
}
