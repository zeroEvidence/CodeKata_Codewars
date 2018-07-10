package kata_drying_potatoes_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "gitlab.dmdigital.net.au/codekata/7kyu/DryingPotatoes"
)

func dotest(p0, w0, p1 int, exp int) {
	var ans = Potatoes(p0, w0, p1)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("DryingPotatoes", func() {
	It("Should handle basic cases", func() {
		dotest(99, 100, 98, 50)
		dotest(82, 127, 80, 114)
	})
})
