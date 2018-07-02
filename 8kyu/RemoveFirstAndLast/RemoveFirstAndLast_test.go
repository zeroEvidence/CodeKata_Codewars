package kata_remove_first_and_last_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "gitlab.dmdigital.net.au/codekata/8kyu/RemoveFirstAndLast"
)

var _ = Describe("RemoveFirstAndLast", func() {
	Describe("Fixed Tests", func() {
		It("eloquent", func() {
			Expect(RemoveFirstAndLast("eloquent")).To(Equal("loquen"))
		})
		It("country", func() {
			Expect(RemoveFirstAndLast("country")).To(Equal("ountr"))
		})
		It("person", func() {
			Expect(RemoveFirstAndLast("person")).To(Equal("erso"))
		})
		It("place", func() {
			Expect(RemoveFirstAndLast("place")).To(Equal("lac"))
		})
	})
})
