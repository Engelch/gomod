package debugerrorce

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIPv4_00(t *testing.T) {
	assert.Nil(t, ValidIPv4Address("0.0.0.0"), "valid IP address not detected as such 00")
}

func TestIPv4_01(t *testing.T) {
	assert.Nil(t, ValidIPv4Address("192.168.255.255"), "valid IP address not detected as such 00")
}

func TestIPv4_02(t *testing.T) {
	assert.Nil(t, ValidIPv4Address("255.255.255.255"), "valid IP address not detected as such 00")
}

func TestIPv4_03(t *testing.T) {
	assert.Nil(t, ValidIPv4Address("1.1.1.255"), "valid IP address not detected as such 00")
}

func TestIPv4_99(t *testing.T) {
	assert.Error(t, ValidIPv4Address("1.1.1.256"), "valid IP address not detected as such 00")
}

func TestIPv4_98(t *testing.T) {
	assert.Error(t, ValidIPv4Address(" 1.1.1.25"), "valid IP address not detected as such 00")
}

func TestIPv4_97(t *testing.T) {
	assert.Error(t, ValidIPv4Address(" 1.1.1.25 "), "valid IP address not detected as such 00")
}

func TestIPv4_96(t *testing.T) {
	assert.Error(t, ValidIPv4Address("1.1.1,25"), "valid IP address not detected as such 00")
}

func TestIPv4_95(t *testing.T) {
	assert.Error(t, ValidIPv4Address("::1"), "valid IP address not detected as such 00")
}

func TestIPv4_94(t *testing.T) {
	assert.Error(t, ValidIPv4Address("1.256.0.0"), "valid IP address not detected as such 00")
}

func TestIPv4_93(t *testing.T) {
	assert.Error(t, ValidIPv4Address("300.1.0.0"), "valid IP address not detected as such 00")
}

func TestIPv4_92(t *testing.T) {
	assert.Error(t, ValidIPv4Address("1.1.256.0"), "valid IP address not detected as such 00")
}

func TestIPv4_91(t *testing.T) {
	assert.Error(t, ValidIPv4Address("1.1.0.256"), "valid IP address not detected as such 00")
}

func TestIPv4_90(t *testing.T) {
	assert.Error(t, ValidIPv4Address("-1.1.0.0"), "valid IP address not detected as such 00")
}
