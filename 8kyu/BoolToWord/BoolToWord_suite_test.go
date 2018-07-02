package kata_bool_to_word_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBoolToWord(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BoolToWord Suite")
}
