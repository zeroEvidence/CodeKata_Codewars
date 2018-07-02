package kata_ball_upwards_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBallUpwards(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BallUpwards Suite")
}
