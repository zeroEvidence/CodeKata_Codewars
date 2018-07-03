package kata_string_repeat_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestStringRepeat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "StringRepeat Suite")
}
