package debugerrorce

import (
	"errors"
	"os/exec"
	"strings"
)

// ExecCmd is a helper to execute an external application. If the exit status of this command is non-zero, then an
// error is returned, else nil.
func ExecCmd(cmd string, args ...string) error {
	CondDebug("Spawning command:" + cmd + " " + strings.Join(args, " "))
	err := exec.Command(cmd, args...).Run()
	if err != nil {
		return errors.New("Error executing cmd:" + cmd + " " + strings.Join(args, " ") + ":" + err.Error())
	}
	return nil
}

// EOF
