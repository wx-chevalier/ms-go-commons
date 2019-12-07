package protoutils

import (
	"testing"

	"github.com/wx-chevalier/go-utils/log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestProtoutil(t *testing.T) {
	RegisterFailHandler(Fail)
	log.DefaultOut = GinkgoWriter
	RunSpecs(t, "Protoutil Suite")
}
