//go:build !bopt
// +build !bopt

package errors

import (
	"fmt"
	"path"
	"runtime"
)

func info() string {
	_, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf("%s:%d", path.Base(file), line)
}
