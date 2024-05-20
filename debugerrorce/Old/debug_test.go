package debugerrorce

import (
	"testing"
)

const (
	debug int = iota
	debugln
	conddebug
	conddebugln
)

func metaDebug(choice int, msg string) func() {
	return func() {
		switch choice {
		case debug:
			Debug(msg)
		case debugln:
			Debugln(msg)
		case conddebug:
			CondDebug(msg)
		case conddebugln:
			CondDebugln(msg)
		default:
			Debug("THIS SHOULD NEVER HAPPEN")
		}
	}
}

func checkOutputAndError(cmdOut string, expectedOut string, cmdErr string, expectedErr string, t *testing.T) {
	if cmdOut != expectedOut {
		t.Errorf("Stdout error, is:%s, expected:%s\n", cmdOut, expectedOut)
	}
	if cmdErr != expectedErr {
		t.Errorf("Stderr error, is:%s, expected:%s\n", cmdErr, expectedErr)
	}
}

func TestDebug(t *testing.T) {
	const msg2 = "01 demo 00"
	err, out := CaptureOutput(metaDebug(debug, msg2))
	checkOutputAndError(out, "", err, "["+msg2+"]", t)
}

func TestDebugln(t *testing.T) {
	const msg2 = "01 demo 00"
	err, out := CaptureOutput(metaDebug(debugln, msg2))
	checkOutputAndError(out, "", err, "["+msg2+"]\n", t)
}

func TestCondDebugEmpty(t *testing.T) {
	const msg2 = "01 demo 99"
	err, out := CaptureOutput(metaDebug(conddebug, msg2))
	checkOutputAndError(out, "", err, "", t)
}

func TestCondDebugNonEmpty(t *testing.T) {
	const msg2 = "01 demo xx"
	CondDebugSet(true)
	err, out := CaptureOutput(metaDebug(conddebug, msg2))
	CondDebugSet(false)
	checkOutputAndError(out, "", err, "["+msg2+"]", t)
}

func TestCondDebuglnEmpty(t *testing.T) {
	const msg2 = "01 demo z33"
	err, out := CaptureOutput(metaDebug(conddebugln, msg2))
	checkOutputAndError(out, "", err, "", t)
}

func TestCondDebuglnNonEmpty(t *testing.T) {
	const msg2 = "01 demo z444"
	CondDebugSet(true)
	err, out := CaptureOutput(metaDebug(conddebugln, msg2))
	CondDebugSet(false)
	checkOutputAndError(out, "", err, "["+msg2+"]\n", t)
}

// EOF
