package main

import (
	"math/big"
)

var ONE = big.NewInt(1)

func init() {
}

func FloorLog2(val *big.Int) *big.Int {
	res := new(big.Int)
	n := (&big.Int{}).Set(val)

	if n.Cmp(big.NewInt(256)) == -1 {
		// At most 8 iterations
		for n.Cmp(big.NewInt(1)) == 1 {
			n.Rsh(n, 1)
			res.Add(res, big.NewInt(1))
		}
	} else {
		// Exactly 8 iterations
		for s := 128; s > 0; s = s >> 1 {
			t := (&big.Int{}).Set(ONE)
			t.Lsh(t, uint(s))
			if n.Cmp(t) >= 0 {
				n.Rsh(n, uint(s))
				res.Or(res, big.NewInt(int64(s)))
			}
		}
	}
	return res
}