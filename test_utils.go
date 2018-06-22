package datetime

import (
	"runtime"
	"testing"
)

var errStart = "\n\033[0;31m"
var errEnd = "\033[0m\n"
var br = "\n"

func assertTrue(t *testing.T, result bool, message string) {
	if result == true {
		return
	}

	_, fn, line, _ := runtime.Caller(1)

	t.Errorf("%s[Error] %s%s:%d%s", errStart, message+"\n", fn, line, errEnd)
}
