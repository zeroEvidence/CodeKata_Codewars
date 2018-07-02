package kata_sum_of_positive_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "gitlab.dmdigital.net.au/codekata/8kyu/SumOfPositive"
)

var _ = Describe("SumOfPositive", func() {
	It("Should equal 20", func() {
		Expect(SumOfPositive([]int{1, -4, 7, 12})).To(Equal(20))
	})
})
