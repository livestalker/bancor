package main

//import (
//	"encoding/hex"
//	"errors"
//	"math/big"
//)
//
//const (
//	MAX_WEIGHT = 1000000
//)
//
//var ONE = big.NewInt(0)
//var FIXED_1 = big.NewInt(0)
//var FIXED_2 = big.NewInt(0)
//var MAX_NUM = big.NewInt(0)
//var LN2_NUMERATOR = big.NewInt(0)
//var LN2_DENOMINATOR = big.NewInt(0)

func main() {
}

//func init() {
//	b, _ := hex.DecodeString("080000000000000000000000000000000")
//	FIXED_1.SetBytes(b)
//	b, _ = hex.DecodeString("100000000000000000000000000000000")
//	FIXED_2.SetBytes(b)
//	b, _ = hex.DecodeString("1ffffffffffffffffffffffffffffffff")
//	MAX_NUM.SetBytes(b)
//	b, _ = hex.DecodeString("3f80fe03f80fe03f80fe03f80fe03f8")
//	LN2_NUMERATOR.SetBytes(b)
//	b, _ = hex.DecodeString("5b9de1d10bf4103d647b0955897ba80")
//	LN2_DENOMINATOR.SetBytes(b)
//}
//
//func CalculatePurchaseReturn(supply, connectorBalance, connectorWeight, depositAmount *big.Int) (*big.Int, error) {
//	zero := big.NewInt(0)
//	// validate input
//	if !(supply.Cmp(zero) == 1 && connectorBalance.Cmp(zero) == 1 && connectorWeight.Cmp(zero) == 1 && (connectorWeight.Cmp(zero) == -1 || connectorWeight.Cmp(zero) == 0)) {
//		return nil, errors.New("Wrong imput parameters")
//	}
//	// special case for 0 deposit amount
//	if depositAmount.Cmp(zero) == 0 {
//		return big.NewInt(0), nil
//	}
//	// special case if the weight = 100%
//	if connectorWeight.Cmp(big.NewInt(MAX_WEIGHT)) == 0 {
//		// TODO ?
//		return big.NewInt(0), nil
//	}
//	result := new(big.Int)
//	precision := new(big.Int)
//	baseN := safeAdd(depositAmount, connectorBalance)
//	//(result, precision) = power(baseN, _connectorBalance, _connectorWeight, MAX_WEIGHT);
//	//uint256 temp = safeMul(_supply, result) >> precision;
//	//return temp - _supply;
//	return nil, nil
//}
//
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
////func power(baseN, baseD, expN, expD *big.Int) (*big.Int, *big.Int) {
////		// TODO
////        uint256 lnBaseTimesExp = ln(_baseN, _baseD) * _expN / _expD;
////        uint8 precision = findPositionInMaxExpArray(lnBaseTimesExp);
////        return (fixedExp(lnBaseTimesExp >> (MAX_PRECISION - precision), precision), precision);
////}
//
//func ln(numerator, denominator *big.Int) *big.Int {
//        // TODO assert(_numerator <= MAX_NUM);
//
//		res := new(big.Int)
//		x := new(big.Int)
//		x.Mul(numerator, FIXED_1).Div(denominator)
//
//        // If x >= 2, then we compute the integer part of log2(x), which is larger than 0.
//        if x.Cmp(FIXED_2) >= 0 {
//            uint8 count = floorLog2(x / FIXED_1);
//            x >>= count; // now x < 2
//            res = count * FIXED_1;
//        }
//
//        // If x > 1, then we compute the fraction part of log2(x), which is larger than 0.
//        if (x > FIXED_1) {
//            for (uint8 i = MAX_PRECISION; i > 0; --i) {
//                x = (x * x) / FIXED_1; // now 1 < x < 4
//                if (x >= FIXED_2) {
//                    x >>= 1; // now 1 < x < 2
//                    res += ONE << (i - 1);
//                }
//            }
//        }
//
//	return res.Mul(res, LN2_NUMERATOR).Div(LN2_DENOMINATOR)
//}
//
//func floorLog2(n *big.Int) *big.Int {
//	res := new(big.Int)
//
//	if res.Cmp(big.NewInt(256)) == -1 {
//		// At most 8 iterations
//		for n.Cmp(big.NewInt(1)) == 1 {
//			n.Rsh(n, 1)
//			res.Add(res, big.NewInt(1))
//		}
//	} else {
//		// Exactly 8 iterations
//		for s := 128; s > 0; s = s >> 1 {
//			t := big.NewInt(0)
//			t.Set(ONE)
//			t.Lsh(t, uint(s))
//			if n.Cmp(t) >= 0 {
//				n.Rsh(n, uint(s))
//				res.Or(res, big.NewInt(int64(s)))
//			}
//		}
//	}
//	return res
//}
