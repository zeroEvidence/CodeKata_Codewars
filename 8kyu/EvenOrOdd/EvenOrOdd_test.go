package kata_even_or_odd_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "gitlab.dmdigital.net.au/codekata/8kyu/EvenOrOdd"
)

var _ = Describe("EvenOrOdd", func() {
	It("Should test that the solution returns the correct value", func() {
		Expect(EvenOrOdd(1)).To(Equal("Odd"))
		Expect(EvenOrOdd(2)).To(Equal("Even"))
	})
})
