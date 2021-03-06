package main

import (
	"encoding/hex"
	"errors"
	"math/big"
)

// Compute the largest integer smaller than or equal to the binary logarithm of the input.
func FloorLog2(val *big.Int) uint {
	var res uint
	n := (&big.Int{}).Set(val)

	if n.Cmp(big.NewInt(256)) == -1 {
		// At most 8 iterations
		for n.Cmp(big.NewInt(1)) == 1 {
			n.Rsh(n, 1)
			res = res + 1
		}
	} else {
		// Exactly 8 iterations
		for s := uint(128); s > 0; s = s >> 1 {
			t := (&big.Int{}).Set(BIG_ONE)
			t.Lsh(t, uint(s))
			if n.Cmp(t) >= 0 {
				n.Rsh(n, uint(s))
				res = res | s
			}
		}
	}
	return res
}

// Return floor(ln(numerator / denominator) * 2 ^ MAX_PRECISION), where:
// - The numerator   is a value between 1 and 2 ^ (256 - MAX_PRECISION) - 1
// - The denominator is a value between 1 and 2 ^ (256 - MAX_PRECISION) - 1
// - The output      is a value between 0 and floor(ln(2 ^ (256 - MAX_PRECISION) - 1) * 2 ^ MAX_PRECISION)
// This functions assumes that the numerator is larger than or equal to the denominator, because the output would be negative otherwise.
func Ln(numerator, denominator *big.Int) (*big.Int, error) {
	if numerator.Cmp(MAX_NUM) == 1 {
		return nil, errors.New("Numerator greater than MAX_NUM")
	}
	res := big.NewInt(0)
	x := big.NewInt(0)
	x.Mul(numerator, FIXED_1)
	x.Div(x, denominator)

	// If x >= 2, then we compute the integer part of log2(x), which is larger than 0.
	if x.Cmp(FIXED_2) >= 0 {
		t := (&big.Int{}).Set(x)
		t.Div(t, FIXED_1)
		count := FloorLog2(t)
		x.Rsh(x, count)
		res.Mul(big.NewInt(int64(count)), FIXED_1)
	}

	// If x > 1, then we compute the fraction part of log2(x), which is larger than 0.
	if x.Cmp(FIXED_1) == 1 {
		for i := MAX_PRECISION - 1; i > 0; i-- {
			x.Mul(x, x)
			x.Div(x, FIXED_1)
			if x.Cmp(FIXED_2) >= 0 {
				x.Rsh(x, 1)
				t := (&big.Int{}).Set(BIG_ONE)
				t.Lsh(BIG_ONE, uint(i-1))
				res.Add(res, t)
			}
		}
	}
	res.Mul(res, LN2_NUMERATOR)
	res.Div(res, LN2_DENOMINATOR)
	return res, nil
}

// General Description:
//         Determine a value of precision.
//         Calculate an integer approximation of (baseN / baseD) ^ (expN / expD) * 2 ^ precision.
//         Return the result along with the precision used.
//
// Detailed Description:
//     Instead of calculating "base ^ exp", we calculate "e ^ (ln(base) * exp)".
//     The value of "ln(base)" is represented with an integer slightly smaller than "ln(base) * 2 ^ precision".
//     The larger "precision" is, the more accurately this value represents the real value.
//     However, the larger "precision" is, the more bits are required in order to store this value.
//     And the exponentiation function, which takes "x" and calculates "e ^ x", is limited to a maximum exponent (maximum value of "x").
//     This maximum exponent depends on the "precision" used, and it is given by "maxExpArray[precision] >> (MAX_PRECISION - precision)".
//     Hence we need to determine the highest precision which can be used for the given input, before calling the exponentiation function.
//     This allows us to compute "base ^ exp" with maximum accuracy and without exceeding 256 bits in any of the intermediate computations.
//     This functions assumes that "_expN < (1 << 256) / ln(MAX_NUM, 1)", otherwise the multiplication should be replaced with a "safeMul".
func Power(baseN, baseD *big.Int, expN, expD uint32) (*big.Int, uint8, error) {
	lnBaseTimesExp, _ := Ln(baseN, baseD)
	lnBaseTimesExp.Mul(lnBaseTimesExp, big.NewInt(int64(expN)))
	lnBaseTimesExp.Div(lnBaseTimesExp, big.NewInt(int64(expD)))
	precision, err := FindPositionInMaxExpArray(lnBaseTimesExp)
	if err != nil {
		return nil, 0, errors.New("Error")
	}
	return FixedExp(lnBaseTimesExp.Rsh(lnBaseTimesExp, uint(MAX_PRECISION-precision)), precision), precision, nil
}

// The global "maxExpArray" is sorted in descending order, and therefore the following statements are equivalent:
// - This function finds the position of [the smallest value in "maxExpArray" larger than or equal to "x"]
// - This function finds the highest position of [a value in "maxExpArray" larger than or equal to "x"]
func FindPositionInMaxExpArray(x *big.Int) (uint8, error) {
	lo := uint8(MIN_PRECISION)
	hi := uint8(MAX_PRECISION)

	for lo+1 < hi {
		mid := (lo + hi) / 2
		if MaxExpArray[mid].Cmp(x) >= 0 {
			lo = mid
		} else {
			hi = mid
		}
	}

	if MaxExpArray[hi].Cmp(x) >= 0 {
		return hi, nil
	}
	if MaxExpArray[lo].Cmp(x) >= 0 {
		return lo, nil
	}

	return 0, errors.New("Position not found")
}

// This function can be auto-generated by the script 'PrintFunctionFixedExp.py'.
//  It approximates "e ^ x" via maclaurin summation: "(x^0)/0! + (x^1)/1! + ... + (x^n)/n!".
//  It returns "e ^ (x / 2 ^ precision) * 2 ^ precision", that is, the result is upshifted for accuracy.
//  The global "maxExpArray" maps each "precision" to "((maximumExponent + 1) << (MAX_PRECISION - precision)) - 1".
//  The maximum permitted value for "x" is therefore given by "maxExpArray[precision] >> (MAX_PRECISION - precision)".
func FixedExp(x *big.Int, precision uint8) *big.Int {
	s33, _ := hex.DecodeString("0688589cc0e9505e2f2fee5580000000")
	f33 := (&big.Int{}).SetBytes(s33)
	xi := (&big.Int{}).Set(x)
	res := big.NewInt(0)
	for _, el := range FactorArray {
		xi.Mul(xi, x).Rsh(xi, uint(precision)).Mul(xi, el)
		res.Add(res, xi)
	}
	t := (&big.Int{}).Set(BIG_ONE)
	t.Lsh(t, uint(precision))
	res.Div(res, f33).Add(res, x).Add(res, t)
	return res
}
