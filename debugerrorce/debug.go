// Copyright (c) 2021 Christian Engel
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package debugerrorce

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync/atomic"
)

// This debug version offers conditional and unconditional debug calls.
// All debug output is written to stderr, so that is is compatible with
// - Docker
// - Kubernetes
// - Linux systemd

// threat-safe implementation
// todo hide variable by closure, PRIO: low
// globalDebug is used to store the state if debugging is on/off. This variable should
// never be set directly.
var globalDebug atomic.Value

// OutputWriter defines the default output channel. It can be changed if required.
var OutputWriter io.Writer = os.Stderr

// init() is always executed at the startup of the application. It makes sure that
// the global debug functionality has a defined state: off ⇔ false.
func init() {
	globalDebug.Store(false)
}

// CondDebugln is the implementation of a global debug function. If it was turned on using
// CondDebugSet(true), then the string is shown to stderr. Else, no output is created.
func CondDebugln(msg ...string) {
	if globalDebug.Load().(bool) {
		fmt.Fprintln(OutputWriter, msg)
	}
}

// CondDebug outputs if debug is set without an added newline at the EOL.
func CondDebug(msg ...string) {
	if globalDebug.Load().(bool) {
		fmt.Fprint(OutputWriter, msg)
	}
}

// CondDebugSet allows us to turn debug on/off.
func CondDebugSet(val bool) {
	globalDebug.Store(val)
}

// Debug outputs a message without adding a newline at the EOL
func Debug(msg ...string) {
	fmt.Fprint(OutputWriter, msg)
}

// Debugln outputs a message with adding a newline at the EOL
func Debugln(msg ...string) {
	fmt.Fprintln(OutputWriter, msg)
}

// CondDebugStatus allows to check if debug is turned on/off.
func CondDebugStatus() bool {
	return globalDebug.Load().(bool)
}

// CaptureOutput get a function as its argument. It executes the function and returns the output (stderr and stdout) created by this function.
// While capturing this output, this output is not written to default stdout or stderr.
func CaptureOutput(f func()) (stderr string, stdout string) {
	//fmt.Println("in captureOutput")
	rerr, werr, err := os.Pipe()
	if err != nil {
		fmt.Fprintf(OutputWriter, "error creating error pipe\n")
		os.Exit(1)
	}
	rout, wout, err := os.Pipe()
	if err != nil {
		fmt.Fprintf(OutputWriter, "error creating output pipe\n")
		os.Exit(1)
	}

	outbuf := bytes.NewBuffer(nil)
	errbuf := bytes.NewBuffer(nil)

	olderr := os.Stderr
	oldOutputWriter := OutputWriter
	oldout := os.Stdout

	os.Stderr = werr
	OutputWriter = werr // we also have to reset OutputWriter, so that CaptureOutput also works for the above routines.
	os.Stdout = wout
	f()
	werr.Close()
	wout.Close()

	os.Stderr = olderr
	OutputWriter = oldOutputWriter
	os.Stdout = oldout
	io.Copy(errbuf, rerr)
	io.Copy(outbuf, rout)

	rerr.Close()
	rout.Close()
	return string(errbuf.Bytes()), string(outbuf.Bytes())
}

// CurrentFunctionName returns the name of the current function being executed.
func CurrentFunctionName() string {
	pc := make([]uintptr, 1) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0]).Name()
	f = filepath.Base(f) // strip the filename and remove the package
	return f
}

// EOF
