package kata_multiple_of_index_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "gitlab.dmdigital.net.au/codekata/8kyu/MultipleOfIndex"
)

var _ = Describe("MultipleOfIndex", func() {
	It("should return the correct values", func() {
		Expect(MultipleOfIndex([]int{22, -6, 32, 82, 9, 25})).To(ConsistOf(-6, 32, 25))
		Expect(MultipleOfIndex([]int{68, -1, 1, -7, 10, 10})).To(ConsistOf(-1, 10))
		Expect(MultipleOfIndex([]int{11, -11})).To(ConsistOf(-11))
		Expect(MultipleOfIndex([]int{-56, -85, 72, -26, -14, 76, -27, 72, 35, -21, -67, 87, 0, 21, 59, 27, -92, 68})).To(ConsistOf(-85, 72, 0, 68))
		Expect(MultipleOfIndex([]int{28, 38, -44, -99, -13, -54, 77, -51})).To(ConsistOf(38, -44, -99))
		Expect(MultipleOfIndex([]int{-1, -49, -1, 67, 8, -60, 39, 35})).To(ConsistOf(-49, 8, -60, 35))
	})
})
