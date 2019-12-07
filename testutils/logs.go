package testutils

import (
	"github.com/fgrosse/zaptest"
	"github.com/wx-chevalier/go-utils/contextutils"

	. "github.com/onsi/ginkgo"
)

func SetupLog() {
	logger := zaptest.LoggerWriter(GinkgoWriter)
	contextutils.SetFallbackLogger(logger.Sugar())
}
