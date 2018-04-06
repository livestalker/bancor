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
	res, err := Ln(numerator, bigOne)
	if err == nil {
		t.Errorf("Case should return error", res)
	}
}

func TestFindPositionInMaxExpArray(t *testing.T) {
	bigZero := big.NewInt(0)
	bigOne := big.NewInt(1)
	for precision := MIN_PRECISION; precision <= MAX_PRECISION; precision++ {
		maxExp := (&big.Int{}).Set(HelperMaxExpArray[precision])
		shlVal := big.NewInt(2)
		shlVal.Exp(shlVal, big.NewInt(int64(MAX_PRECISION-precision)), nil)

		i1 := (&big.Int{}).Set(maxExp)
		i1.Add(i1, bigZero).Mul(i1, shlVal).Sub(i1, bigOne)
		o1 := big.NewInt(int64(precision))

		i2 := (&big.Int{}).Set(maxExp)
		i2.Add(i2, bigZero).Mul(i2, shlVal).Sub(i2, bigZero)
		o2 := big.NewInt(int64(precision))

		i3 := (&big.Int{}).Set(maxExp)
		i3.Add(i3, bigOne).Mul(i3, shlVal).Sub(i3, bigOne)
		o3 := big.NewInt(int64(precision))

		i4 := (&big.Int{}).Set(maxExp)
		i4.Add(i4, bigOne).Mul(i4, shlVal).Sub(i4, bigZero)
		o4 := big.NewInt(int64(precision - 1))

		tuples := []map[string]*big.Int{
			map[string]*big.Int{"input": i1, "output": o1},
			map[string]*big.Int{"input": i2, "output": o2},
			map[string]*big.Int{"input": i3, "output": o3},
			map[string]*big.Int{"input": i4, "output": o4},
		}
		for _, el := range tuples {
			input := el["input"]
			output := el["output"]
			res, err := FindPositionInMaxExpArray(input)
			if err == nil && big.NewInt(int64(res)).Cmp(output) != 0 {
				t.Errorf("Output should be %s but it is %d", output, res)
			}
			if err == nil && !(precision > MIN_PRECISION || output.Cmp(big.NewInt(int64(precision))) >= 0) {
				t.Error("Passed when it should have failed")
			}
			if err != nil && !(precision == MIN_PRECISION && output.Cmp(big.NewInt(int64(precision))) == -1) {
				t.Error("Failed when it should have passed")
			}
		}
		break
	}
}

func TestFixedExp(t *testing.T) {
	bigOne := big.NewInt(1)
	res := FixedExp(big.NewInt(0), 0)
	if res.Cmp(bigOne) != 0 {
		t.Errorf("Output should be 1 but it is %d", res)
	}
	t.Error(FixedExp(big.NewInt(1), 1))
}
