package versionutils_test

import (
	"testing"

	"github.com/wx-chevalier/go-utils/testutils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestVersionUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	testutils.RegisterPreFailHandler(
		func() {
			testutils.PrintTrimmedStack()
		})
	testutils.RegisterCommonFailHandlers()
	RunSpecs(t, "Versionutils Suite")
}
