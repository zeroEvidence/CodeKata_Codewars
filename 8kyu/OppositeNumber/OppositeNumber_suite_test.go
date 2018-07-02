package kata_opposite_number_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOppositeNumber(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OppositeNumber Suite")
}
