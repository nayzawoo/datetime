package datetime

import (
	"fmt"
	"runtime"
	"testing"
)

var errStart = "\n\033[0;31m"
var errEnd = "\033[0m\n"

func assertTrue(t *testing.T, result bool, message string) {
	if result == true {
		return
	}

	_, fn, line, _ := runtime.Caller(1)

	t.Errorf("%s[Error] %s%s:%d%s", errStart, message+"\n", fn, line, errEnd)
}

func assertDateTime(t *testing.T, dt *DateTime, datetimeString string) {
	_, f, line, _ := runtime.Caller(1)

	if dt.DateTimeString() == datetimeString {
		return
	}

	message := "\nExpected: " + datetimeString + "\nActual: " + dt.DateTimeString() + "\n"

	t.Errorf(formatError(message, f, line))
}

func formatError(msg, f string, line int) string {
	return fmt.Sprintf("%s[Error] %s \n[File]%s:%d%s", errStart, msg, f, line, errEnd)
}
