package debugerrorce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPlainFile0(t *testing.T) {
	// pure file
	out := IsPlainFile("file.go")
	assert.Equal(t, true, out)
}

func TestIsPlainFile1(t *testing.T) {
	// path
	out := IsPlainFile("./file.go")
	assert.Equal(t, true, out)
}

func TestIsPlainFile2(t *testing.T) {
	// dir
	out := IsPlainFile(".")
	assert.Equal(t, false, out)
}

func TestIsPlainFile3(t *testing.T) {
	// char dev
	out := IsPlainFile("/dev/null")
	assert.Equal(t, false, out)
}

func TestIsPlainFile4(t *testing.T) {
	// non-existing
	out := IsPlainFile("/thisShallNeverExit898908087373")
	assert.Equal(t, false, out)
}

// =================================

func TestIsExisting0(t *testing.T) {
	assert.Equal(t, true, IsExistingFile("file_test.go"))
}

func TestIsExisting1(t *testing.T) {
	assert.Equal(t, true, IsExistingFile("/"))
}

func TestIsExisting2(t *testing.T) {
	assert.Equal(t, true, IsExistingFile(".."))
}

func TestIsExisting3(t *testing.T) {
	assert.Equal(t, true, IsExistingFile("/dev/null"))
}

func TestIsExisting4(t *testing.T) {
	assert.Equal(t, false, IsExistingFile("/dev/null8308409380"))
}

// =================================

func TestIsDirectory0(t *testing.T) {
	assert.Equal(t, false, IsDirectory("file_test.go"))
}

func TestIsDirectory1(t *testing.T) {
	assert.Equal(t, true, IsDirectory("."))
}

func TestIsDirectory2(t *testing.T) {
	assert.Equal(t, true, IsDirectory("/"))
}

func TestIsDirectory3(t *testing.T) {
	assert.Equal(t, false, IsDirectory("/dev/null"))
}

// EOF
