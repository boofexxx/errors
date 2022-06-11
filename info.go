//go:build !bopt
// +build !bopt

package errors

import (
	"fmt"
	"path"
	"runtime"
)

func info() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return "???:???"
	}
	return fmt.Sprintf("%s:%d", path.Base(file), line)
}
