package kata_square_or_square_root_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSquareOrSquareRoot(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SquareOrSquareRoot Suite")
}
