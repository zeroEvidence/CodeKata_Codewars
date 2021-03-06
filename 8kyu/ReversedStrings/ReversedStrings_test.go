package kata_reversed_strings_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "gitlab.dmdigital.net.au/codekata/8kyu/ReversedStrings"
)

var _ = Describe("ReversedStrings", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(ReversedStrings("world")).To(Equal("dlrow"))
	})
})
