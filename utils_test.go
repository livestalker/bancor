package main

import (
	"math/big"
	"testing"
)

func TestFloorLog2(t *testing.T) {
	bigOne := big.NewInt(1)
	for n := 1; n <= 255; n++ {
		input1 := big.NewInt(2)
		input1.Exp(input1, big.NewInt(int64(n)), nil)

		input2 := big.NewInt(2)
		input2.Exp(input2, big.NewInt(int64(n)), nil)
		input2.Add(input2, bigOne)

		input3 := big.NewInt(2)
		input3.Exp(input2, big.NewInt(int64(n+1)), nil)
		input3.Sub(input2, bigOne)

		output := n
		res := FloorLog2(input1)
		if res != output {
			t.Errorf("Error FloorLog2: expected %s, got %s", output, res)
		}
		res = FloorLog2(input2)
		if res != output {
			t.Errorf("Error FloorLog2: expected %s, got %s", output, res)
		}
		res = FloorLog2(input3)
		if res != output {
			t.Errorf("Error FloorLog2: expected %s, got %s", output, res)
		}
	}
}
