package main

import (
	"errors"
	"math/big"
)

func main() {
}

// Given a token supply, connector balance, weight and a deposit amount (in the connector token),
// calculates the return for a given conversion (in the main token)
//
// Formula:
// Return = supply * ((1 + depositAmount / connectorBalance) ^ (connectorWeight / 1000000) - 1)
//
// supply           *big.Int   token total supply
// connectorBalance uint32     total connector balance
// connectorWeight  *big.Int   connector weight, represented in ppm, 1-1000000
// depositAmount    *big.Int   deposit amount, in connector token
//
// return *big.Int purchase return amount
func CalculatePurchaseReturn(supply, connectorBalance *big.Int, connectorWeight uint32, depositAmount *big.Int) (*big.Int, error) {
	// validate input
	if !(supply.Cmp(BIG_ZERO) == 1 && connectorBalance.Cmp(BIG_ZERO) == 1 && connectorWeight > 0 && connectorWeight <= MAX_WEIGHT) {
		return nil, errors.New("Wrong input parameters")
	}
	// special case for 0 deposit amount
	if depositAmount.Cmp(BIG_ZERO) == 0 {
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

// Given a token supply, connector balance, weight and a sell amount (in the main token),
// calculates the return for a given conversion (in the connector token)
//
// Formula:
// Return = connectorBalance * (1 - (1 - sellAmount / supply) ^ (1 / (connectorWeight / 1000000)))
//
// supply           *big.Int   token total supply
// connectorBalance *big.Int   total connector
// connectorWeight  uint32     constant connector Weight, represented in ppm, 1-1000000
// sellAmount       *big.Int   sell amount, in the token itself
//
// return *big.Int sale return amount
func CalculateSaleReturn(supply, connectorBalance *big.Int, connectorWeight uint32, sellAmount *big.Int) (*big.Int, error) {
	// validate input
	if !(supply.Cmp(BIG_ZERO) == 1 && connectorBalance.Cmp(BIG_ZERO) == 1 && connectorWeight > 0 && connectorWeight <= MAX_WEIGHT && sellAmount.Cmp(supply) <= 0) {
		return nil, errors.New("Wrong input parameters")
	}
	// special case for 0 sell amount
	if sellAmount.Cmp(BIG_ZERO) == 0 {
		return big.NewInt(0), nil
	}
	// special case for selling the entire supply
	if sellAmount.Cmp(supply) == 0 {
		return connectorBalance, nil
	}
	// special case if the weight = 100%
	if connectorWeight == MAX_WEIGHT {
		res := new(big.Int)
		res.Mul(connectorBalance, sellAmount).Div(res, supply)
		return res, nil
	}
	baseD := new(big.Int)
	baseD.Sub(supply, sellAmount)
	res, precision, _ := Power(supply, baseD, MAX_WEIGHT, connectorWeight)
	tmp1 := new(big.Int)
	tmp1.Mul(connectorBalance, res)
	tmp2 := new(big.Int)
	tmp2.Lsh(connectorBalance, uint(precision))
	return tmp1.Sub(tmp1, tmp2).Div(tmp1, res), nil
}
