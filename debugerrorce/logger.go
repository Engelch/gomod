package debugerrorce

// small helper class for logging + writing to stderr

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"time"
)

var logInitialised = false
var stringLog = false
var stag string

// verifyLogInitialised panics if log is not initialised.
func verifyLogInitialised() {
	if !logInitialised {
		panic("Log was not initialsed")
	}
}

// LogInit tries to initialise the logging service.
func LogInit(tag string) {
	logInitialised = true
	stringLog = false
	stag = tag
}

// LogStringInit does not use syslog (for dockerised environments. Instead, it writes all messages to stderr)
// This is suited for dockerised environments.
func LogStringInit(tag string) {
	logInitialised = true
	stringLog = true
	stag = tag
}

// doLog handles the actual writing of the logging message
func doLog(msg string, prio syslog.Priority) {
	if stringLog {
		_, _ = fmt.Fprintln(os.Stderr, stag+"::"+msg)
	} else {
		syslogger, err := syslog.New(prio, stag)
		if err != nil {
			panic("Cannot initialise syslog")
		}
		log.SetOutput(syslogger)
		log.Println(msg)
		syslogger.Close() // flushing
	}
}

// LogErr creates a message preprended with ERROR to syslog and stderr, but tries to continue execution.
func LogErr(msg string) {
	verifyLogInitialised()
	doLog("ERROR:"+time.Now().UTC().Format(time.RFC3339)+":"+msg, syslog.LOG_ERR)
}

// LogWarn creates a syslog and STDERR message labeled with WARNING.
func LogWarn(msg string) {
	verifyLogInitialised()
	doLog("Warning:"+time.Now().UTC().Format(time.RFC3339)+":"+msg, syslog.LOG_WARNING)
}

// LogInfo creates an info error message to syslog and STDERR.
func LogInfo(msg string) {
	verifyLogInitialised()
	doLog("info:"+time.Now().UTC().Format(time.RFC3339)+":"+msg, syslog.LOG_INFO)
}

// EOF
