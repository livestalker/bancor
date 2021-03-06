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

var (
	BIG_ZERO        = new(big.Int)
	BIG_ONE         = big.NewInt(1)
	MAX_NUM         = new(big.Int)
	FIXED_1         = new(big.Int)
	FIXED_2         = new(big.Int)
	LN2_NUMERATOR   = new(big.Int)
	LN2_DENOMINATOR = new(big.Int)
)

var (
	MaxExpArray [128]*big.Int
	FactorArray [32]*big.Int
)

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
	initMaxExpArray()
	initFactorArray()
}

func initMaxExpArray() {
	var strMaxExpArray [128]string
	strMaxExpArray[0] = "6bffffffffffffffffffffffffffffffff"
	strMaxExpArray[1] = "67ffffffffffffffffffffffffffffffff"
	strMaxExpArray[2] = "637fffffffffffffffffffffffffffffff"
	strMaxExpArray[3] = "5f6fffffffffffffffffffffffffffffff"
	strMaxExpArray[4] = "5b77ffffffffffffffffffffffffffffff"
	strMaxExpArray[5] = "57b3ffffffffffffffffffffffffffffff"
	strMaxExpArray[6] = "5419ffffffffffffffffffffffffffffff"
	strMaxExpArray[7] = "50a2ffffffffffffffffffffffffffffff"
	strMaxExpArray[8] = "4d517fffffffffffffffffffffffffffff"
	strMaxExpArray[9] = "4a233fffffffffffffffffffffffffffff"
	strMaxExpArray[10] = "47165fffffffffffffffffffffffffffff"
	strMaxExpArray[11] = "4429afffffffffffffffffffffffffffff"
	strMaxExpArray[12] = "415bc7ffffffffffffffffffffffffffff"
	strMaxExpArray[13] = "3eab73ffffffffffffffffffffffffffff"
	strMaxExpArray[14] = "3c1771ffffffffffffffffffffffffffff"
	strMaxExpArray[15] = "399e96ffffffffffffffffffffffffffff"
	strMaxExpArray[16] = "373fc47fffffffffffffffffffffffffff"
	strMaxExpArray[17] = "34f9e8ffffffffffffffffffffffffffff"
	strMaxExpArray[18] = "32cbfd5fffffffffffffffffffffffffff"
	strMaxExpArray[19] = "30b5057fffffffffffffffffffffffffff"
	strMaxExpArray[20] = "2eb40f9fffffffffffffffffffffffffff"
	strMaxExpArray[21] = "2cc8340fffffffffffffffffffffffffff"
	strMaxExpArray[22] = "2af09481ffffffffffffffffffffffffff"
	strMaxExpArray[23] = "292c5bddffffffffffffffffffffffffff"
	strMaxExpArray[24] = "277abdcdffffffffffffffffffffffffff"
	strMaxExpArray[25] = "25daf6657fffffffffffffffffffffffff"
	strMaxExpArray[26] = "244c49c65fffffffffffffffffffffffff"
	strMaxExpArray[27] = "22ce03cd5fffffffffffffffffffffffff"
	strMaxExpArray[28] = "215f77c047ffffffffffffffffffffffff"
	strMaxExpArray[29] = "1fffffffffffffffffffffffffffffffff"
	strMaxExpArray[30] = "1eaefdbdabffffffffffffffffffffffff"
	strMaxExpArray[31] = "1d6bd8b2ebffffffffffffffffffffffff"
	strMaxExpArray[32] = "1c35fedd14ffffffffffffffffffffffff"
	strMaxExpArray[33] = "1b0ce43b323fffffffffffffffffffffff"
	strMaxExpArray[34] = "19f0028ec1ffffffffffffffffffffffff"
	strMaxExpArray[35] = "18ded91f0e7fffffffffffffffffffffff"
	strMaxExpArray[36] = "17d8ec7f0417ffffffffffffffffffffff"
	strMaxExpArray[37] = "16ddc6556cdbffffffffffffffffffffff"
	strMaxExpArray[38] = "15ecf52776a1ffffffffffffffffffffff"
	strMaxExpArray[39] = "15060c256cb2ffffffffffffffffffffff"
	strMaxExpArray[40] = "1428a2f98d72ffffffffffffffffffffff"
	strMaxExpArray[41] = "13545598e5c23fffffffffffffffffffff"
	strMaxExpArray[42] = "1288c4161ce1dfffffffffffffffffffff"
	strMaxExpArray[43] = "11c592761c666fffffffffffffffffffff"
	strMaxExpArray[44] = "110a688680a757ffffffffffffffffffff"
	strMaxExpArray[45] = "1056f1b5bedf77ffffffffffffffffffff"
	strMaxExpArray[46] = "0faadceceeff8bffffffffffffffffffff"
	strMaxExpArray[47] = "0f05dc6b27edadffffffffffffffffffff"
	strMaxExpArray[48] = "0e67a5a25da4107fffffffffffffffffff"
	strMaxExpArray[49] = "0dcff115b14eedffffffffffffffffffff"
	strMaxExpArray[50] = "0d3e7a392431239fffffffffffffffffff"
	strMaxExpArray[51] = "0cb2ff529eb71e4fffffffffffffffffff"
	strMaxExpArray[52] = "0c2d415c3db974afffffffffffffffffff"
	strMaxExpArray[53] = "0bad03e7d883f69bffffffffffffffffff"
	strMaxExpArray[54] = "0b320d03b2c343d5ffffffffffffffffff"
	strMaxExpArray[55] = "0abc25204e02828dffffffffffffffffff"
	strMaxExpArray[56] = "0a4b16f74ee4bb207fffffffffffffffff"
	strMaxExpArray[57] = "09deaf736ac1f569ffffffffffffffffff"
	strMaxExpArray[58] = "0976bd9952c7aa957fffffffffffffffff"
	strMaxExpArray[59] = "09131271922eaa606fffffffffffffffff"
	strMaxExpArray[60] = "08b380f3558668c46fffffffffffffffff"
	strMaxExpArray[61] = "0857ddf0117efa215bffffffffffffffff"
	strMaxExpArray[62] = "07ffffffffffffffffffffffffffffffff"
	strMaxExpArray[63] = "07abbf6f6abb9d087fffffffffffffffff"
	strMaxExpArray[64] = "075af62cbac95f7dfa7fffffffffffffff"
	strMaxExpArray[65] = "070d7fb7452e187ac13fffffffffffffff"
	strMaxExpArray[66] = "06c3390ecc8af379295fffffffffffffff"
	strMaxExpArray[67] = "067c00a3b07ffc01fd6fffffffffffffff"
	strMaxExpArray[68] = "0637b647c39cbb9d3d27ffffffffffffff"
	strMaxExpArray[69] = "05f63b1fc104dbd39587ffffffffffffff"
	strMaxExpArray[70] = "05b771955b36e12f7235ffffffffffffff"
	strMaxExpArray[71] = "057b3d49dda84556d6f6ffffffffffffff"
	strMaxExpArray[72] = "054183095b2c8ececf30ffffffffffffff"
	strMaxExpArray[73] = "050a28be635ca2b888f77fffffffffffff"
	strMaxExpArray[74] = "04d5156639708c9db33c3fffffffffffff"
	strMaxExpArray[75] = "04a23105873875bd52dfdfffffffffffff"
	strMaxExpArray[76] = "0471649d87199aa990756fffffffffffff"
	strMaxExpArray[77] = "04429a21a029d4c1457cfbffffffffffff"
	strMaxExpArray[78] = "0415bc6d6fb7dd71af2cb3ffffffffffff"
	strMaxExpArray[79] = "03eab73b3bbfe282243ce1ffffffffffff"
	strMaxExpArray[80] = "03c1771ac9fb6b4c18e229ffffffffffff"
	strMaxExpArray[81] = "0399e96897690418f785257fffffffffff"
	strMaxExpArray[82] = "0373fc456c53bb779bf0ea9fffffffffff"
	strMaxExpArray[83] = "034f9e8e490c48e67e6ab8bfffffffffff"
	strMaxExpArray[84] = "032cbfd4a7adc790560b3337ffffffffff"
	strMaxExpArray[85] = "030b50570f6e5d2acca94613ffffffffff"
	strMaxExpArray[86] = "02eb40f9f620fda6b56c2861ffffffffff"
	strMaxExpArray[87] = "02cc8340ecb0d0f520a6af58ffffffffff"
	strMaxExpArray[88] = "02af09481380a0a35cf1ba02ffffffffff"
	strMaxExpArray[89] = "0292c5bdd3b92ec810287b1b3fffffffff"
	strMaxExpArray[90] = "0277abdcdab07d5a77ac6d6b9fffffffff"
	strMaxExpArray[91] = "025daf6654b1eaa55fd64df5efffffffff"
	strMaxExpArray[92] = "0244c49c648baa98192dce88b7ffffffff"
	strMaxExpArray[93] = "022ce03cd5619a311b2471268bffffffff"
	strMaxExpArray[94] = "0215f77c045fbe885654a44a0fffffffff"
	strMaxExpArray[95] = "01ffffffffffffffffffffffffffffffff"
	strMaxExpArray[96] = "01eaefdbdaaee7421fc4d3ede5ffffffff"
	strMaxExpArray[97] = "01d6bd8b2eb257df7e8ca57b09bfffffff"
	strMaxExpArray[98] = "01c35fedd14b861eb0443f7f133fffffff"
	strMaxExpArray[99] = "01b0ce43b322bcde4a56e8ada5afffffff"
	strMaxExpArray[100] = "019f0028ec1fff007f5a195a39dfffffff"
	strMaxExpArray[101] = "018ded91f0e72ee74f49b15ba527ffffff"
	strMaxExpArray[102] = "017d8ec7f04136f4e5615fd41a63ffffff"
	strMaxExpArray[103] = "016ddc6556cdb84bdc8d12d22e6fffffff"
	strMaxExpArray[104] = "015ecf52776a1155b5bd8395814f7fffff"
	strMaxExpArray[105] = "015060c256cb23b3b3cc3754cf40ffffff"
	strMaxExpArray[106] = "01428a2f98d728ae223ddab715be3fffff"
	strMaxExpArray[107] = "013545598e5c23276ccf0ede68034fffff"
	strMaxExpArray[108] = "01288c4161ce1d6f54b7f61081194fffff"
	strMaxExpArray[109] = "011c592761c666aa641d5a01a40f17ffff"
	strMaxExpArray[110] = "0110a688680a7530515f3e6e6cfdcdffff"
	strMaxExpArray[111] = "01056f1b5bedf75c6bcb2ce8aed428ffff"
	strMaxExpArray[112] = "00faadceceeff8a0890f3875f008277fff"
	strMaxExpArray[113] = "00f05dc6b27edad306388a600f6ba0bfff"
	strMaxExpArray[114] = "00e67a5a25da41063de1495d5b18cdbfff"
	strMaxExpArray[115] = "00dcff115b14eedde6fc3aa5353f2e4fff"
	strMaxExpArray[116] = "00d3e7a3924312399f9aae2e0f868f8fff"
	strMaxExpArray[117] = "00cb2ff529eb71e41582cccd5a1ee26fff"
	strMaxExpArray[118] = "00c2d415c3db974ab32a51840c0b67edff"
	strMaxExpArray[119] = "00bad03e7d883f69ad5b0a186184e06bff"
	strMaxExpArray[120] = "00b320d03b2c343d4829abd6075f0cc5ff"
	strMaxExpArray[121] = "00abc25204e02828d73c6e80bcdb1a95bf"
	strMaxExpArray[122] = "00a4b16f74ee4bb2040a1ec6c15fbbf2df"
	strMaxExpArray[123] = "009deaf736ac1f569deb1b5ae3f36c130f"
	strMaxExpArray[124] = "00976bd9952c7aa957f5937d790ef65037"
	strMaxExpArray[125] = "009131271922eaa6064b73a22d0bd4f2bf"
	strMaxExpArray[126] = "008b380f3558668c46c91c49a2f8e967b9"
	strMaxExpArray[127] = "00857ddf0117efa215952912839f6473e6"
	for i := 0; i < 128; i++ {
		b, _ := hex.DecodeString(strMaxExpArray[i])
		MaxExpArray[i] = (&big.Int{}).SetBytes(b)
	}
}

func initFactorArray() {
	var strFactorArray [32]string
	strFactorArray[0] = "03442c4e6074a82f1797f72ac0000000"  // (33! / 2!)
	strFactorArray[1] = "0116b96f757c380fb287fd0e40000000"  // (33! / 3!)
	strFactorArray[2] = "0045ae5bdd5f0e03eca1ff4390000000"  // (33! / 4!)
	strFactorArray[3] = "000defabf91302cd95b9ffda50000000"  // (33! / 5!)
	strFactorArray[4] = "0002529ca9832b22439efff9b8000000"  // (33! / 6!)
	strFactorArray[5] = "000054f1cf12bd04e516b6da88000000"  // (33! / 7!)
	strFactorArray[6] = "00000a9e39e257a09ca2d6db51000000"  // (33! / 8!)
	strFactorArray[7] = "0000012e066e7b839fa050c309000000"  // (33! / 9!)
	strFactorArray[8] = "0000001e33d7d926c329a1ad1a800000"  // (33! / 10!)
	strFactorArray[9] = "00000002bee513bdb4a6b19b5f800000"  // (33! / 11!)
	strFactorArray[10] = "000000003a9316fa79b88eccf2a00000" // (33! / 12!)
	strFactorArray[11] = "00000000048177ebe1fa812375200000" // (33! / 13!)
	strFactorArray[12] = "00000000005263fe90242dcbacf00000" // (33! / 14!)
	strFactorArray[13] = "0000000000057e22099c030d94100000" // (33! / 15!)
	strFactorArray[14] = "00000000000057e22099c030d9410000" // (33! / 16!)
	strFactorArray[15] = "000000000000052b6b54569976310000" // (33! / 17!)
	strFactorArray[16] = "000000000000004985f67696bf748000" // (33! / 18!)
	strFactorArray[17] = "0000000000000003dea12ea99e498000" // (33! / 19!)
	strFactorArray[18] = "000000000000000031880f2214b6e000" // (33! / 20!)
	strFactorArray[19] = "0000000000000000025bcff56eb36000" // (33! / 21!)
	strFactorArray[20] = "0000000000000000001b722e10ab1000" // (33! / 22!)
	strFactorArray[21] = "00000000000000000001317c70077000" // (33! / 23!)
	strFactorArray[22] = "000000000000000000000cba84aafa00" // (33! / 24!)
	strFactorArray[23] = "000000000000000000000082573a0a00" // (33! / 25!)
	strFactorArray[24] = "000000000000000000000005035ad900" // (33! / 26!)
	strFactorArray[25] = "0000000000000000000000002f881b00" // (33! / 27!)
	strFactorArray[26] = "00000000000000000000000001b29340" // (33! / 28!)
	strFactorArray[27] = "000000000000000000000000000efc40" // (33! / 29!)
	strFactorArray[28] = "00000000000000000000000000007fe0" // (33! / 30!)
	strFactorArray[29] = "00000000000000000000000000000420" // (33! / 31!)
	strFactorArray[30] = "00000000000000000000000000000021" // (33! / 32!)
	strFactorArray[31] = "00000000000000000000000000000001" // (33! / 33!)
	for i := 0; i < 32; i++ {
		b, _ := hex.DecodeString(strFactorArray[i])
		FactorArray[i] = (&big.Int{}).SetBytes(b)
	}
}
