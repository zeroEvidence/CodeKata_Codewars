package kata_century_from_year_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCenturyFromYear(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CenturyFromYear Suite")
}
