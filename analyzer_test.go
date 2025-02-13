package linter

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/test-go/testify/require"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestPluginExample(t *testing.T) {
	analysistest.Run(t, testdataDir(t), NoZeroLogMsgfAnalyzer, "p")
}

func testdataDir(t *testing.T) string {
	t.Helper()

	_, testFilename, _, ok := runtime.Caller(1)
	if !ok {
		require.Fail(t, "unable to get current test filename")
	}

	return filepath.Join(filepath.Dir(testFilename), "testdata")
}
