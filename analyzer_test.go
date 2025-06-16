package nozerologmsgf

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/golangci/plugin-module-register/register"
	"github.com/test-go/testify/require"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestPluginExample(t *testing.T) {
	nozerologmsgfPlugin, err := register.GetPlugin("nozerologmsgf")
	require.NoError(t, err)

	plugin, err := nozerologmsgfPlugin(nil)
	require.NoError(t, err)

	analyzers, err := plugin.BuildAnalyzers()
	require.NoError(t, err)

	analysistest.Run(t, testdataDir(t), analyzers[0], "p")
}

func testdataDir(t *testing.T) string {
	t.Helper()

	_, testFilename, _, ok := runtime.Caller(1)
	if !ok {
		require.Fail(t, "unable to get current test filename")
	}

	return filepath.Join(filepath.Dir(testFilename), "testdata")
}
