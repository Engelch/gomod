package debugerrorce

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecutableReachableByPath0(t *testing.T) {
	// we expect that bash, sed, grep, wc are available
	err := ExecutableReachableByPath("bash", "sed", "grep", "wc")
	assert.Nil(t, err)
}

func TestExecutableReachableByPath1(t *testing.T) {
	// we expect that the binary jfdaslkjd09jlk does not exist
	err := ExecutableReachableByPath("jfdaslkjd09jlk")
	assert.NotNil(t, err)
}

// ===================================================================

// expects that CWD is writable
func TestExecNoOutputCmd0(t *testing.T) {
	// we expect that the binary jfdaslkjd09jlk does not exist
	err := ExecNoOutputCmd("rm -f bla")
	assert.Nil(t, err)
	assert.Equal(t, false, IsExistingFile("bla"))
	err = ExecNoOutputCmd("touch bla")
	assert.Nil(t, err)
	assert.Equal(t, true, IsExistingFile("bla"))
	err = ExecNoOutputCmd("rm -f bla")
	assert.Nil(t, err)
	assert.Equal(t, false, IsExistingFile("bla"))
}

func TestExecNoOutputCmd1(t *testing.T) {
	// we expect that the binary jfdaslkjd09jlk does not exist
	err := ExecNoOutputCmd("/bjalsdl")
	assert.NotNil(t, err)
}

// ===================================================================

func TestExecOutCmd0(t *testing.T) {
	// we expect that the binary jfdaslkjd09jlk does not exist
	out, err := ExecOutputCmd("echo demo")
	assert.Nil(t, err)
	assert.Equal(t, out, []byte("demo\n"))
}

func TestExecOutCmd1(t *testing.T) {
	// we expect that the binary jfdaslkjd09jlk does not exist
	out, err := ExecOutputCmd("echo demo | wc -c")
	assert.Nil(t, err)
	outNormalised := strings.Fields(string(out))[0]
	assert.Equal(t, outNormalised, "5")
}

func TestExecOutCmd2(t *testing.T) {
	// we expect that the binary jfdaslkjd09jlk does not exist
	_, err := ExecOutputCmd("ehcNotExisting jkdajld")
	assert.NotNil(t, err)
}

// EOF
