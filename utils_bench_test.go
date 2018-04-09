package main

import (
	"math/big"
	"testing"
)

func BenchmarkFixedExp(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FixedExp(big.NewInt(10), 127)
	}
}
