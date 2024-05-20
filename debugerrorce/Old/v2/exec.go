package debugerrorce

import (
	"errors"
	"os/exec"
	"strings"
)

// ExecCmd is a helper to execute an external application. If the exit status of this command is non-zero, then an
// error is returned, else nil.
// Deprecated: ExecCmd is deprecated. Replaced by ExecNoOutputCmd.
func ExecCmd(cmd string, args ...string) error {
	CondDebug("Spawning command:" + cmd + " " + strings.Join(args, " "))
	err := exec.Command(cmd, args...).Run()
	if err != nil {
		return errors.New("Error executing cmd:" + cmd + " " + strings.Join(args, " ") + ":" + err.Error())
	}
	return nil
}

// ExecNoOutputCmd calls a command but does not expect handle output except for the return code of the app.
func ExecNoOutputCmd(cmd string) error {
	err := exec.Command("bash", "-c", cmd).Run()
	if err != nil {
		return errors.New("ERROR:ExecNoOutputCmd:" + cmd + ":" + err.Error())
	}
	return nil
}

// ExecutableReachableByPath checks for all given input if the input is executable and can be found by
// the current setting of the PATH variable.
func ExecutableReachableByPath(cmd ...string) error {
	for _, val := range cmd {
		if _, err := exec.LookPath(val); err != nil {
			return errors.New("ERROR:" + CurrentFunctionName() + ":command not found:" + val)
		}
	}
	return nil
}

func ExecOutputCmd(cmd string) ([]byte, error) {
	if out, err := exec.Command("bash", "-c", cmd).Output(); err == nil {
		return out, nil
	} else {
		return nil, errors.New("ERROR:" + CurrentFunctionName() + ":" + err.Error())
	}
}

// EOF
