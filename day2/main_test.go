// main_test.go

package main

import (
	"testing"
)

func BenchmarkPart1(b *testing.B) {

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		Part1()
	}
}
