package debugerrorce

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// isPlainFile is a predicate returning true if the supplied argument is an existing, plain file (no directory, device-file,...)
func IsPlainFile(filename string) bool {
	if stat, err := os.Stat(filename); err == nil && strings.HasPrefix(stat.Mode().String(), "-") {
		// fmt.Printf("filename mode: %v\n", stat.Mode())
		return true
	}
	return false
}

// isExistingFile predicate that sometimes can make the code easier (if we do not are about the error value)
func IsExistingFile(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	}
	return false
}

func IsDirectory(filename string) bool {
	if stat, err := os.Stat(filename); err == nil && stat.IsDir() {
		return true
	}
	return false
}

func FilenameWithoutSuffix(filename string) string {
	extension := filepath.Ext(filename)
	return filename[0 : len(filename)-len(extension)]
}

// ByteArray2File writes a byte array into a file. If required it does so in multiple steps.
// If all succeeds then nil is returned, otherwise an error.
func ByteArray2File(file *os.File, bytes []byte) error {
	nsum := 0
	n := 0
	var err error
	for ; nsum < len(bytes); nsum += n {
		n, err = file.Write(bytes[nsum:])
		if err != nil { // should also never happen
			return errors.New(CurrentFunctionName() + ":" + err.Error())
		}
	}
	return nil
}

// ByteArray2ReponseWriter writes a byte array into a file. If required it does so in multiple steps.
// If all succeeds then nil is returned, otherwise an error.
func ByteArray2ReponseWriter(file http.ResponseWriter, bytes []byte) error {
	nsum := 0
	n := 0
	var err error
	for ; nsum < len(bytes); nsum += n {
		n, err = file.Write(bytes[nsum:])
		if err != nil { // should also never happen
			return errors.New(CurrentFunctionName() + ":" + err.Error())
		}
	}
	return nil
}

// IsExecAny is called by isExecutableCmd and is a predicate to check if the file is executable by anyone.
// TODO: more complex if we do not run as root; currently, we run as root, not complete as FS could be mounted with noexec :-)
func isExecAny(fileinfo os.FileInfo) bool {
	CondDebug(fmt.Sprintf(CurrentFunctionName()+" fileinfo:%v", fileinfo))
	CondDebug(fmt.Sprintf(CurrentFunctionName()+" perms:%#v", fileinfo.Mode().Perm()))
	return fileinfo.Mode().Perm()&0111 != 0
}

// IsExecutableCmd is a predicate checking if the given parameter represents an executable file.
// It returns nil in the positive case. Otherwise, an error is returned.
func IsExecutableCmd(cmd string) error {
	stat, err := os.Stat(cmd)
	if err != nil {
		return errors.New("Error with supplied command:" + cmd + ":" + err.Error())
	}
	if stat.Mode().IsDir() {
		return errors.New("Error command is a directory:" + cmd)
	}
	if !stat.Mode().IsRegular() {
		return errors.New("Error command is not a regular file:" + cmd)
	}
	if !isExecAny(stat) {
		return errors.New("Error command is not executable:" + cmd)
	}
	return nil
}

// eof
