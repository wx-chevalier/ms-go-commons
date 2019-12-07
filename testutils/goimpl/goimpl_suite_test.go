package goimpl

import (
	"testing"

	"github.com/wx-chevalier/go-utils/testutils"

	. "github.com/onsi/ginkgo"
)

func TestGoImpl(t *testing.T) {

	testutils.RegisterPreFailHandler(
		func() {
			testutils.PrintTrimmedStack()
		})
	testutils.RegisterCommonFailHandlers()
	RunSpecs(t, "Go Impl Suite")
}
