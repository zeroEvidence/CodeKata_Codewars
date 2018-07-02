package kata_remove_first_and_last_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRemoveFirstAndLast(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RemoveFirstAndLast Suite")
}
