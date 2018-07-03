package kata_string_repeat_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "gitlab.dmdigital.net.au/codekata/8kyu/StringRepeat"
)

var _ = Describe("StringRepeat", func() {
	It("should repeat correctly", func() {
		Expect(StringRepeat(4, "a")).To(Equal("aaaa"))
		Expect(StringRepeat(3, "hello ")).To(Equal("hello hello hello "))
		Expect(StringRepeat(2, "abc")).To(Equal("abcabc"))
	})

})
