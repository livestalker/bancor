package main

import (
	"errors"
	"math/big"
)

func main() {
}

func CalculatePurchaseReturn(supply, connectorBalance *big.Int, connectorWeight uint32, depositAmount *big.Int) (*big.Int, error) {
	zero := big.NewInt(0)
	// validate input
	if !(supply.Cmp(zero) == 1 && connectorBalance.Cmp(zero) == 1 && connectorWeight > 0 && connectorWeight <= MAX_WEIGHT) {
		return nil, errors.New("Wrong input parameters")
	}
	// special case for 0 deposit amount
	if depositAmount.Cmp(zero) == 0 {
		return big.NewInt(0), nil
	}
	// special case if the weight = 100%
	if connectorWeight == MAX_WEIGHT {
		res := (&big.Int{}).Set(supply)
		res.Mul(res, depositAmount).Div(res, connectorBalance)
		return res, nil
	}
	baseN := new(big.Int)
	baseN.Add(depositAmount, connectorBalance)
	res, precision, _ := Power(baseN, connectorBalance, connectorWeight, MAX_WEIGHT)
	tmp := new(big.Int)
	tmp.Mul(supply, res).Rsh(tmp, uint(precision)).Sub(tmp, supply)
	return tmp, nil
}

//func CalculateSaleReturn(supply, connectorBalance, connectorWeight, sellAmount *big.Int) *big.Int {
//	return big.NewInt(0)
//}
//
//func safeAdd(x, y *big.Int) *big.Int {
//	z := new(big.Int)
//	z.Add(x, y)
//	// TODO assert(z >= _x);
//	return z
//}
//
//func safeMul(x, y *big.Int) *big.Int {
//	z := new(big.Int)
//	z.Mul(x, y)
//	// TODO assert(_x == 0 || z / _x == _y);
//	return z
//}
//
