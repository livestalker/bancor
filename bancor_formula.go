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

func CalculateSaleReturn(supply, connectorBalance *big.Int, connectorWeight uint32, sellAmount *big.Int) (*big.Int, error) {
	// validate input
	if !(supply.Cmp(ZERO) == 1 && connectorBalance.Cmp(ZERO) == 1 && connectorWeight > 0 && connectorWeight <= MAX_WEIGHT && sellAmount.Cmp(supply) <= 0) {
		return nil, errors.New("Wrong input parameters")
	}
	// special case for 0 sell amount
	if sellAmount.Cmp(ZERO) == 0 {
		return big.NewInt(0), nil
	}
	// special case for selling the entire supply
	if sellAmount.Cmp(supply) == 0 {
		return connectorBalance, nil
	}
	// special case if the weight = 100%
	if connectorWeight == MAX_WEIGHT {
		res := big.NewInt(0)
		res.Mul(connectorBalance, sellAmount).Div(res, supply)
		return res, nil
	}
	baseD := big.NewInt(0)
	baseD.Sub(supply, sellAmount)
	res, precision, _ := Power(supply, baseD, MAX_WEIGHT, connectorWeight)
	tmp1 := new(big.Int)
	tmp1.Mul(connectorBalance, res)
	tmp2 := new(big.Int)
	tmp2.Lsh(connectorBalance, uint(precision))
	return tmp1.Sub(tmp1, tmp2).Div(tmp1, res), nil
}
