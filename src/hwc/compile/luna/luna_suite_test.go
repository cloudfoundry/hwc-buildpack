package luna_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLuna(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Luna Suite")
}
