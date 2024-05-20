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
	"os"
	"strings"
)

// ErrorExit exits the application with the specified error code. The output is
// written to the assigned output writer, by default stderr.
func ErrorExit(errorCode uint8, msg ...string) {
	// join string array
	str := strings.Join(msg, " ")
	if !strings.HasSuffix(str, "\n") {
		str = str + "\n"
	}
	Debug("*ERROR*:" + str)
	os.Exit(int(errorCode))
}

// ExitIfError exists using ErrorExit if the supplied err is not nil. In such a case,
// the error message of err will be added to the message.
func ExitIfError(err error, exitcode uint8, msg string) {
	if err != nil {
		ErrorExit(exitcode, msg+":"+err.Error())
	}
}

// EOF
