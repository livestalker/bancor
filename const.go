package main

import (
	"encoding/hex"
	"math/big"
)

const (
	MIN_PRECISION = 32
	MAX_PRECISION = 127
	MAX_WEIGHT    = 1000000
)

var ZERO = big.NewInt(0)
var ONE = big.NewInt(1)
var MAX_NUM = big.NewInt(0)
var FIXED_1 = big.NewInt(0)
var FIXED_2 = big.NewInt(0)
var LN2_NUMERATOR = big.NewInt(0)
var LN2_DENOMINATOR = big.NewInt(0)

func init() {
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
