package fileutils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopyFileErrorIfDirectory(t *testing.T) {
	t.Log("It fails if source is a directory")
	{
		tmpFolder, err := NormalizedOSTempDirPath("_tmp")
		require.NoError(t, err)
		require.EqualError(t, CopyFile(tmpFolder, "./nothing/whatever"), "Source is a directory: "+tmpFolder)
	}
}
