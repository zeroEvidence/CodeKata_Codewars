package kata_grasshopper_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "gitlab.dmdigital.net.au/codekata/8kyu/Grasshopper"
)

var _ = Describe("Damaged but alive", func() {
	It("Should return the correct positive value", func() {
		Expect(combat(100.0, 50.0)).To(Equal(50.0))
	})
})
var _ = Describe("Dead"), func() {
	It("Should return 0 rather than negative health", func() {
		Expect(combat(1.0, 100)).To(Equal(0))
	})
}
