package main

/*
Benchmark tests for the command argument functions.
To run it: $ go test -bench=.
*/

import (
	"io"
	"testing"
)

func BenchmarkShowCommandArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		showCommandArgs(io.Discard)
	}
}

func BenchmarkShowCommandArgs2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		showCommandArgs2(io.Discard)
	}
}

func BenchmarkShowCommandArgs2b(b *testing.B) {
	for i := 0; i < b.N; i++ {
		showCommandArgs2b(io.Discard)
	}
}

func BenchmarkShowCommandArgs3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		showCommandArgs3(io.Discard)
	}
}
