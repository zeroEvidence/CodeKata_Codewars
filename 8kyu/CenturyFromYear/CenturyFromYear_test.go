package kata_century_from_year_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "gitlab.dmdigital.net.au/codekata/8kyu/CenturyFromYear"
)

var _ = Describe("CenturyFromYear", func() {
	It("Should return the century", func() {
		Expect(CenturyFromYear(1705)).To(Equal(18))
		Expect(CenturyFromYear(1900)).To(Equal(19))
		Expect(CenturyFromYear(1601)).To(Equal(17))
		Expect(CenturyFromYear(2000)).To(Equal(20))
	})
})
