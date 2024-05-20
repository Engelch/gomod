package debugerrorce

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	_, err := GetEnvValue("PATH")
	if err != nil {
		t.Error(CurrentFunctionName() + ":PATH should always be set.")
	}
}

func TestEnvGetValue(t *testing.T) {
	// unset environment variable
	if err := os.Unsetenv("XPATH"); err != nil {
		log.Fatal(err)
	}
	_, err := GetEnvValue("XPATH")
	if err == nil {
		t.Error(CurrentFunctionName() + ":PATH should should not be set.")
	}
}

func TestEnvGetBool0(t *testing.T) {
	_ = os.Setenv("XUID", "true")
	res := GetEnvValueOrDefaultBool("XUID", false)
	if !res {
		t.Error(fmt.Sprintf("%s%v", CurrentFunctionName()+":UID is expected to be positive:", res))

	}
}

func TestEnvGetBool1(t *testing.T) {
	_ = os.Setenv("XUID", "TRUE")
	res := GetEnvValueOrDefaultBool("XUID", false)
	if !res {
		t.Error(fmt.Sprintf("%s%v", CurrentFunctionName()+":UID is expected to be positive:", res))

	}
}

func TestEnvGetBool2(t *testing.T) {
	_ = os.Setenv("XUID", "True")
	res := GetEnvValueOrDefaultBool("XUID", false)
	if !res {
		t.Error(fmt.Sprintf("%s%v", CurrentFunctionName()+":UID is expected to be positive:", res))

	}
}

func TestEnvGetBool3(t *testing.T) {
	_ = os.Setenv("XUID", "false")
	res := GetEnvValueOrDefaultBool("XUID", false)
	if res {
		t.Error(fmt.Sprintf("%s%v", CurrentFunctionName()+":UID is expected to be positive:", res))

	}
}

func TestEnvGetInt(t *testing.T) {
	_ = os.Setenv("XUID", "99")
	val := GetEnvValueOrDefaultInt("XUID", -1)
	if val < 0 {
	}
	if val != 99 {
		t.Error(fmt.Sprintf("%s%d", CurrentFunctionName()+":UID is expected to be 99:", val))
	}
}

func TestEnvGetIntWrongType(t *testing.T) {
	_ = os.Setenv("XUID", "bla")
	val := GetEnvValueOrDefaultInt("XUID", -1)
	if val == -1 {
		return
	}
	t.Error(fmt.Sprintf("%s%d", CurrentFunctionName()+":UID is expected to be positive:", val))
}

// test string variant

func TestGetEnvString(t *testing.T) {
	_ = os.Setenv("XUID", "99")
	res := GetEnvValueOrDefaultString("XUID", "error")
	if res == "99" {
		return
	}
	t.Error(fmt.Sprintf("%s%s\n", CurrentFunctionName()+":expected restult was 99 but is:", res))
}

func TestGetEnvStringNonExisting(t *testing.T) {
	// unset environment variable
	if err := os.Unsetenv("XUID"); err != nil {
		log.Fatal(err)
	}
	res := GetEnvValueOrDefaultString("XUID", "error")
	if res == "error" {
		return
	}
	t.Error(fmt.Sprintf("%s%s\n", CurrentFunctionName()+":expected restult was 99 but is:", res))
}

// only non-fatal test can be done :-)

func TestGetEnvStringNonFatal(t *testing.T) {
	// unset environment variable
	if err := os.Setenv("XUID", "99"); err != nil {
		log.Fatal(err)
	}
	res := FatalGetEnvValue("XUID")
	if res != "99" {
		t.Error(fmt.Sprintf("%s%s\n", CurrentFunctionName()+":expected restult was 99 but is:", res))
	}
}

// EOF
