package errors

import (
	"fmt"
	"testing"
)

func BenchmarkError(b *testing.B) {
	op := Op("Op")
	err := New("nothing")
	msg := "BenchmarkError"
	for i := 0; i < b.N; i++ {
		_ = &Error{Err: err, Op: op, Msg: msg, info: info()}
	}
}

func BenchmarkE(b *testing.B) {
	op := Op("Op")
	err := New("nothing")
	msg := "BenchmarkError"
	for i := 0; i < b.N; i++ {
		_ = E(err, op, msg)
	}
}

func BenchmarkErrorf(b *testing.B) {
	op := Op("Op")
	err := New("nothing")
	msg := "BenchmarkError"
	for i := 0; i < b.N; i++ {
		_ = Errorf("%s: %s: %w", op, msg, err)
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = New("nothing")
	}
}

func BenchmarkErrorWrapRec(b *testing.B) {
	op := Op("BenchmarkError")
	err := New("nothing")
	msg := "BenchmarkError"
	for i := 0; i < b.N; i++ {
		err = &Error{Err: err, Op: op, Msg: msg, info: info()}
	}
}

func BenchmarkEWrapRec(b *testing.B) {
	op := Op("BenchmarkError")
	err := New("nothing")
	msg := "BenchmarkError"
	for i := 0; i < b.N; i++ {
		err = E(err, op, msg)
	}
}

func BenchmarkErrorfWrapRec(b *testing.B) {
	op := Op("BenchmarkError")
	err := New("nothing")
	msg := "BenchmarkError"
	for i := 0; i < b.N; i++ {
		err = fmt.Errorf("%s: %s: %w", op, msg, err)
	}
}
