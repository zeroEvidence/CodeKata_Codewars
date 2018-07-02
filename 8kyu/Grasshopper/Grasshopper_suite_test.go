package kata_grasshopper_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGrasshopper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Grasshopper Suite")
}
