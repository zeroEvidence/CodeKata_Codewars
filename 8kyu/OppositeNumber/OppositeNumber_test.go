package kata_opposite_number_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "gitlab.dmdigital.net.au/codekata/8kyu/OppositeNumber"
)

var _ = Describe("OppositeNumber", func() {
	It("Should invert number", func() {
		Expect(OppositeNumber(1)).To(Equal(-1))
	})
})
