package pnm

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

const testResourceDir = "testdata"

func OpenGoldenFile(t *testing.T, filename, source string, update bool) string {
	t.Helper()
	reader := Open(t, filename, os.O_RDWR)
	defer reader.(*os.File).Close()
	if update {
		err := reader.(*os.File).Truncate(0)
		CheckErrorf(t, err)

		_, err = reader.(*os.File).WriteString(source)
		CheckErrorf(t, err)
	}
	content, err := ioutil.ReadAll(reader)
	CheckErrorf(t, err)

	return string(content)
}

func Open(t *testing.T, filename string, perm int) io.ReadWriter {
	t.Helper()
	path := filepath.Join(testResourceDir, filename)
	file, err := os.OpenFile(path, perm, 0644)
	if err != nil {
		t.Errorf("Can't open file %v", path)
	}

	return file
}

func CheckErrorf(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Fail wiht %v", err)
	}
}
