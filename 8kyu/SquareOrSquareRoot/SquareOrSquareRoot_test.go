package kata_square_or_square_root_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "gitlab.dmdigital.net.au/codekata/8kyu/SquareOrSquareRoot"
)

var _ = Describe("SquareOrSquareRoot", func() {
	tests := [][][]int{
		{{4, 3, 9, 7, 2, 1}, {2, 9, 3, 49, 4, 1}},
		{{100, 101, 5, 5, 1, 1}, {10, 10201, 25, 25, 1, 1}},
		{{1, 2, 3, 4, 5, 6}, {1, 4, 9, 2, 25, 36}},
	}

	It("should test that the solution returns the correct value", func() {
		for _, t := range tests {
			Expect(SquareOrSquareRoot(t[0])).To(Equal(t[1]))
		}
	})
})
