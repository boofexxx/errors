package errors

import (
	"bytes"
	"errors"
	"fmt"
)

var (
	Is     = errors.Is
	Unwrap = errors.Unwrap
	As     = errors.As
	New    = errors.New
	Errorf = fmt.Errorf
)

var (
	_ error = (*Error)(nil)
)

type Error struct {
	// Op is the operation being performed, usually the name of the method
	// being invoked (Get, Put, etc.). It should not contain an at sign @.
	Op Op
	// Handwrited message that appears after Op
	Msg string
	// The underlying error that triggered this one, if any.
	Err error

	// info is the location in file where E is called otherwise it's empty
	info string
}

func (e *Error) Error() string {
	b := &bytes.Buffer{}
	if e.Op != "" {
		pad(b, ": ")
		b.WriteString(string(e.Op))
	}
	if e.Msg != "" {
		pad(b, ": ")
		b.WriteString(e.Msg)
	}
	if e.Err != nil {
		// Indent on new line if we are cascading non-empty
		if prevErr, ok := e.Err.(*Error); ok {
			if !prevErr.isZero() {
				if e.info != "" {
					pad(b, ": ")
					b.WriteString(e.info)
				}
				pad(b, Separator)
				b.WriteString(e.Err.Error())
			}
		} else {
			pad(b, ": ")
			b.WriteString(e.Err.Error())
			if e.info != "" {
				pad(b, ": ")
				b.WriteString(e.info)
			}
		}
	}
	return b.String()
}

func (e *Error) Unwrap() error {
	return e.Err
}

func (e *Error) isZero() bool {
	return e.Op == "" && e.Msg == "" && e.info == "" && e.Err == nil
}

// Op describes an operation, usually as the package and method,
// such as "key/server.Lookup".
type Op string

// Separator is the string used to separate nested errors. By
// default, to make errors easier on the eye, nested errors are
// indented on a new line. A server may instead choose to keep each
// error on a single line by modifying the separator string, perhaps
// to ":: ".
var Separator = ":\n\t"

// E builds an error value from its arguments.
// There must be at least one argument or E panics.
// The type of each argument determines its meaning.
// If more than one argument of a given type is presented,
// only the last one is recorded.
func E(args ...interface{}) error {
	if len(args) == 0 {
		panic("call to errors.E with no arguments")
	}
	err := &Error{info: info()}
	for _, arg := range args {
		switch arg := arg.(type) {
		case error:
			err.Err = arg
		case Op:
			err.Op = arg
		case string:
			err.Msg = arg
		default:
			return Errorf("unknown type %T, value %v in error call", arg, arg)
		}
	}
	return err
}

// pad appends str to the buffer if the buffer already has some data.
func pad(b *bytes.Buffer, str string) {
	if b.Len() == 0 {
		return
	}
	b.WriteString(str)
}

// Innermost returns innermost error by unwrapping
func Innermost(err error) error {
	next := err
	for next != nil {
		err = next
		next = Unwrap(err)
	}
	return err
}
