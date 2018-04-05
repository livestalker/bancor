package main

import (
	"fmt"
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

		output := uint(n)
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

func TestLn(t *testing.T) {
	bigOne := big.NewInt(1)
	ILLEGAL_VALUE := big.NewInt(2)
	ILLEGAL_VALUE.Exp(ILLEGAL_VALUE, big.NewInt(256), nil)
	MAX_EXPONENT := 1000000
	MAX_NUMERATOR := big.NewInt(2)
	MAX_NUMERATOR.Exp(MAX_NUMERATOR, big.NewInt(256-MAX_PRECISION), nil)
	MAX_NUMERATOR.Sub(MAX_NUMERATOR, bigOne)

	numerator := (&big.Int{}).Set(MAX_NUMERATOR)
	denominator := (&big.Int{}).Set(MAX_NUMERATOR)
	denominator.Sub(denominator, bigOne)
	res, _ := Ln(numerator, denominator)
	res.Mul(res, big.NewInt(int64(MAX_EXPONENT)))
	if res.Cmp(ILLEGAL_VALUE) == 1 {
		t.Errorf("%s output is too large", res)
	}

	res, _ = Ln(numerator, bigOne)
	res.Mul(res, big.NewInt(int64(MAX_EXPONENT)))
	if res.Cmp(ILLEGAL_VALUE) == 1 {
		t.Errorf("%s output is too large", res)
	}

	numerator = (&big.Int{}).Set(MAX_NUMERATOR)
	numerator.Add(numerator, bigOne)
	fmt.Println(numerator)
	res, err := Ln(numerator, bigOne)
	if err == nil {
		t.Errorf("Case should return error", res)
	}
}
