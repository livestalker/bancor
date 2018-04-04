package main

import (
	"fmt"
	"math/big"
	"testing"
)

func TestFloorLog2(t *testing.T) {
	for n := 1; n <= 255; n++ {
		input := big.NewInt(2)
		output := big.NewInt(int64(n))
		input.Exp(input, big.NewInt(int64(n)), nil)
		fmt.Println(input, " ", FloorLog2(input), " ", output)
	}
}
