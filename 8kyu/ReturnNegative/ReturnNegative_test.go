package kata_return_negative_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "gitlab.dmdigital.net.au/codekata/8kyu/ReturnNegative"
)

var _ = Describe("return negative", func() {
	It("should multiply integers", func() {
		Expect(ReturnNegative(42)).To(Equal(-42))
	})
})
