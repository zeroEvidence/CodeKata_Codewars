package kata_sum_of_positive_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSumOfPositive(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SumOfPositive Suite")
}
