package main

import (
	"encoding/hex"
	//	"errors"
	"math/big"
)

var ONE = big.NewInt(1)
var MAX_NUM = big.NewInt(0)
var FIXED_1 = big.NewInt(0)
var FIXED_2 = big.NewInt(0)

func init() {
	b, _ := hex.DecodeString("1ffffffffffffffffffffffffffffffff")
	MAX_NUM.SetBytes(b)
	b, _ = hex.DecodeString("080000000000000000000000000000000")
	FIXED_1.SetBytes(b)
	b, _ = hex.DecodeString("100000000000000000000000000000000")
	FIXED_2.SetBytes(b)
}

func FloorLog2(val *big.Int) int {
	res := 0
	n := (&big.Int{}).Set(val)

	if n.Cmp(big.NewInt(256)) == -1 {
		// At most 8 iterations
		for n.Cmp(big.NewInt(1)) == 1 {
			n.Rsh(n, 1)
			res = res + 1
		}
	} else {
		// Exactly 8 iterations
		for s := 128; s > 0; s = s >> 1 {
			t := (&big.Int{}).Set(ONE)
			t.Lsh(t, uint(s))
			if n.Cmp(t) >= 0 {
				n.Rsh(n, uint(s))
				res = res | s
			}
		}
	}
	return res
}

//func Ln(numerator, denominator *big.Int) (*big.Int, error) {
//	if numerator.Cmp(MAX_NUM) == 1 {
//		return nil, errors.New("Numerator greater than MAX_NUM")
//	}
//	res := new(big.Int)
//	x := new(big.Int)
//	x.Mul(numerator, FIXED_1)
//	x.Div(x, denominator)
//
//	// If x >= 2, then we compute the integer part of log2(x), which is larger than 0.
//	if x.Cmp(FIXED_2) >= 0 {
//		t := (&big.Int{}).Set(x)
//		t.Div(t, FIXED_1)
//		count := FloorLog2(t)
//		x.Rsh(x,
//	x >>= count; // now x < 2
//	res = count * FIXED_1;
//	}
//
//	      // If x > 1, then we compute the fraction part of log2(x), which is larger than 0.
//	      if (x > FIXED_1) {
//	          for (uint8 i = MAX_PRECISION; i > 0; --i) {
//	              x = (x * x) / FIXED_1; // now 1 < x < 4
//	              if (x >= FIXED_2) {
//	                  x >>= 1; // now 1 < x < 2
//	                  res += ONE << (i - 1);
//	              }
//	          }
//	      }
//
//	return res.Mul(res, LN2_NUMERATOR).Div(LN2_DENOMINATOR)
//}
