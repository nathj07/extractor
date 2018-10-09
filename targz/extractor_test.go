package targz

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtractTarGz(t *testing.T) {
	sourcePath := "testdata/PMC3405500.tar.gz"
	dest := "testdata/output"
	nxmlName := "JEM_20101974.nxml"
	formats := map[string]struct{}{".nxml": struct{}{}}
	extractedFiles, err := Extract(sourcePath, dest, formats)
	require.Nil(t, err)
	defer os.RemoveAll(dest)

	// check only one file
	files, _ := ioutil.ReadDir(filepath.Join(dest, "PMC3405500"))
	assert.Equal(t, 1, len(files))

	// check the file is the right one
	_, err = os.Stat(filepath.Join(dest, "PMC3405500", nxmlName))
	require.Nil(t, err)
	assert.Equal(t, nxmlName, filepath.Base(extractedFiles[0]))
}

func TestExtractNoFile(t *testing.T) {
	sourcePath := "testdata/NotHere.tar.gz"
	dest := "testdata/output"
	formats := map[string]struct{}{".txt": struct{}{}}
	_, err := Extract(sourcePath, dest, formats)
	require.NotNil(t, err)
}
