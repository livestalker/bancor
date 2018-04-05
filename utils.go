package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
)

const (
	MIN_PRECISION = 32
	MAX_PRECISION = 127
)

var ONE = big.NewInt(1)
var MAX_NUM = big.NewInt(0)
var FIXED_1 = big.NewInt(0)
var FIXED_2 = big.NewInt(0)
var LN2_NUMERATOR = big.NewInt(0)
var LN2_DENOMINATOR = big.NewInt(0)

func init() {
	fmt.Println("Init")
	b, _ := hex.DecodeString("01ffffffffffffffffffffffffffffffff")
	MAX_NUM.SetBytes(b)
	b, _ = hex.DecodeString("0080000000000000000000000000000000")
	FIXED_1.SetBytes(b)
	b, _ = hex.DecodeString("0100000000000000000000000000000000")
	FIXED_2.SetBytes(b)
	b, _ = hex.DecodeString("03f80fe03f80fe03f80fe03f80fe03f8")
	LN2_NUMERATOR.SetBytes(b)
	b, _ = hex.DecodeString("05b9de1d10bf4103d647b0955897ba80")
	LN2_DENOMINATOR.SetBytes(b)
}

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
			if x.Cmp(FIXED_2) > 0 {
				x.Rsh(x, 1)
				t := (&big.Int{}).Set(ONE)
				t.Lsh(ONE, uint(i-1))
				res.Add(res, t)
			}
		}
	}
	res.Mul(res, LN2_NUMERATOR)
	res.Div(res, LN2_DENOMINATOR)
	return res, nil
}

//func Power(baseN, baseD *big.Int, expN, expD uint32) (*big.Int, uint8) {
//	lnBaseTimesExp, _ := Ln(baseN, baseD)
//	lnBaseTimesExp.Mul(lnBaseTimesExp, big.NewInt(int64(expN)))
//	lnBaseTimesExp.Div(lnBaseTimesExp, big.NewInt(int64(expD)))
//	uint8 precision = findPositionInMaxExpArray(lnBaseTimesExp);
//	return (fixedExp(lnBaseTimesExp >> (MAX_PRECISION - precision), precision), precision);
//}

func FindPositionInMaxExpArray(x *big.Int) (uint8, error) {
	lo := uint8(MIN_PRECISION)
	hi := uint8(MAX_PRECISION)

	for lo+1 < hi {
		mid := (lo + hi) / 2
		//fmt.Println(MaxExpArray[mid])
		if MaxExpArray[mid].Cmp(x) > 0 {
			lo = mid
		} else {
			hi = mid
		}
	}

	if MaxExpArray[hi].Cmp(x) > 0 {
		return hi, nil
	}
	if MaxExpArray[lo].Cmp(x) > 0 {
		return lo, nil
	}

	return 0, errors.New("Position not found")
}
