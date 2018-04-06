package main

import (
	"encoding/hex"
	"math/big"
)

var (
	HelperMaxExpArray [128]*big.Int
	HelperMaxValArray [128]*big.Int
)

func init() {
	initHelperMaxExpArray()
	initHelperMaxValArray()
}

func initHelperMaxExpArray() {
	var HelperStrMaxExpArray [128]string
	HelperStrMaxExpArray[0] = "d7"
	HelperStrMaxExpArray[1] = "019f"
	HelperStrMaxExpArray[2] = "031b"
	HelperStrMaxExpArray[3] = "05f6"
	HelperStrMaxExpArray[4] = "0b6e"
	HelperStrMaxExpArray[5] = "15ec"
	HelperStrMaxExpArray[6] = "2a0c"
	HelperStrMaxExpArray[7] = "50a2"
	HelperStrMaxExpArray[8] = "9aa2"
	HelperStrMaxExpArray[9] = "01288c"
	HelperStrMaxExpArray[10] = "0238b2"
	HelperStrMaxExpArray[11] = "04429a"
	HelperStrMaxExpArray[12] = "082b78"
	HelperStrMaxExpArray[13] = "0faadc"
	HelperStrMaxExpArray[14] = "1e0bb8"
	HelperStrMaxExpArray[15] = "399e96"
	HelperStrMaxExpArray[16] = "6e7f88"
	HelperStrMaxExpArray[17] = "d3e7a3"
	HelperStrMaxExpArray[18] = "01965fea"
	HelperStrMaxExpArray[19] = "030b5057"
	HelperStrMaxExpArray[20] = "05d681f3"
	HelperStrMaxExpArray[21] = "0b320d03"
	HelperStrMaxExpArray[22] = "15784a40"
	HelperStrMaxExpArray[23] = "292c5bdd"
	HelperStrMaxExpArray[24] = "4ef57b9b"
	HelperStrMaxExpArray[25] = "976bd995"
	HelperStrMaxExpArray[26] = "0122624e32"
	HelperStrMaxExpArray[27] = "022ce03cd5"
	HelperStrMaxExpArray[28] = "042beef808"
	HelperStrMaxExpArray[29] = "07ffffffff"
	HelperStrMaxExpArray[30] = "0f577eded5"
	HelperStrMaxExpArray[31] = "1d6bd8b2eb"
	HelperStrMaxExpArray[32] = "386bfdba29"
	HelperStrMaxExpArray[33] = "6c3390ecc8"
	HelperStrMaxExpArray[34] = "cf8014760f"
	HelperStrMaxExpArray[35] = "018ded91f0e7"
	HelperStrMaxExpArray[36] = "02fb1d8fe082"
	HelperStrMaxExpArray[37] = "05b771955b36"
	HelperStrMaxExpArray[38] = "0af67a93bb50"
	HelperStrMaxExpArray[39] = "15060c256cb2"
	HelperStrMaxExpArray[40] = "285145f31ae5"
	HelperStrMaxExpArray[41] = "4d5156639708"
	HelperStrMaxExpArray[42] = "944620b0e70e"
	HelperStrMaxExpArray[43] = "011c592761c666"
	HelperStrMaxExpArray[44] = "02214d10d014ea"
	HelperStrMaxExpArray[45] = "0415bc6d6fb7dd"
	HelperStrMaxExpArray[46] = "07d56e76777fc5"
	HelperStrMaxExpArray[47] = "0f05dc6b27edad"
	HelperStrMaxExpArray[48] = "1ccf4b44bb4820"
	HelperStrMaxExpArray[49] = "373fc456c53bb7"
	HelperStrMaxExpArray[50] = "69f3d1c921891c"
	HelperStrMaxExpArray[51] = "cb2ff529eb71e4"
	HelperStrMaxExpArray[52] = "0185a82b87b72e95"
	HelperStrMaxExpArray[53] = "02eb40f9f620fda6"
	HelperStrMaxExpArray[54] = "05990681d961a1ea"
	HelperStrMaxExpArray[55] = "0abc25204e02828d"
	HelperStrMaxExpArray[56] = "14962dee9dc97640"
	HelperStrMaxExpArray[57] = "277abdcdab07d5a7"
	HelperStrMaxExpArray[58] = "4bb5ecca963d54ab"
	HelperStrMaxExpArray[59] = "9131271922eaa606"
	HelperStrMaxExpArray[60] = "0116701e6ab0cd188d"
	HelperStrMaxExpArray[61] = "0215f77c045fbe8856"
	HelperStrMaxExpArray[62] = "03ffffffffffffffff"
	HelperStrMaxExpArray[63] = "07abbf6f6abb9d087f"
	HelperStrMaxExpArray[64] = "0eb5ec597592befbf4"
	HelperStrMaxExpArray[65] = "1c35fedd14b861eb04"
	HelperStrMaxExpArray[66] = "3619c87664579bc94a"
	HelperStrMaxExpArray[67] = "67c00a3b07ffc01fd6"
	HelperStrMaxExpArray[68] = "c6f6c8f8739773a7a4"
	HelperStrMaxExpArray[69] = "017d8ec7f04136f4e561"
	HelperStrMaxExpArray[70] = "02dbb8caad9b7097b91a"
	HelperStrMaxExpArray[71] = "057b3d49dda84556d6f6"
	HelperStrMaxExpArray[72] = "0a830612b6591d9d9e61"
	HelperStrMaxExpArray[73] = "1428a2f98d728ae223dd"
	HelperStrMaxExpArray[74] = "26a8ab31cb8464ed99e1"
	HelperStrMaxExpArray[75] = "4a23105873875bd52dfd"
	HelperStrMaxExpArray[76] = "8e2c93b0e33355320ead"
	HelperStrMaxExpArray[77] = "0110a688680a7530515f3e"
	HelperStrMaxExpArray[78] = "020ade36b7dbeeb8d79659"
	HelperStrMaxExpArray[79] = "03eab73b3bbfe282243ce1"
	HelperStrMaxExpArray[80] = "0782ee3593f6d69831c453"
	HelperStrMaxExpArray[81] = "0e67a5a25da41063de1495"
	HelperStrMaxExpArray[82] = "1b9fe22b629ddbbcdf8754"
	HelperStrMaxExpArray[83] = "34f9e8e490c48e67e6ab8b"
	HelperStrMaxExpArray[84] = "6597fa94f5b8f20ac16666"
	HelperStrMaxExpArray[85] = "c2d415c3db974ab32a5184"
	HelperStrMaxExpArray[86] = "0175a07cfb107ed35ab61430"
	HelperStrMaxExpArray[87] = "02cc8340ecb0d0f520a6af58"
	HelperStrMaxExpArray[88] = "055e129027014146b9e37405"
	HelperStrMaxExpArray[89] = "0a4b16f74ee4bb2040a1ec6c"
	HelperStrMaxExpArray[90] = "13bd5ee6d583ead3bd636b5c"
	HelperStrMaxExpArray[91] = "25daf6654b1eaa55fd64df5e"
	HelperStrMaxExpArray[92] = "4898938c9175530325b9d116"
	HelperStrMaxExpArray[93] = "8b380f3558668c46c91c49a2"
	HelperStrMaxExpArray[94] = "010afbbe022fdf442b2a522507"
	HelperStrMaxExpArray[95] = "01ffffffffffffffffffffffff"
	HelperStrMaxExpArray[96] = "03d5dfb7b55dce843f89a7dbcb"
	HelperStrMaxExpArray[97] = "075af62cbac95f7dfa3295ec26"
	HelperStrMaxExpArray[98] = "0e1aff6e8a5c30f58221fbf899"
	HelperStrMaxExpArray[99] = "1b0ce43b322bcde4a56e8ada5a"
	HelperStrMaxExpArray[100] = "33e0051d83ffe00feb432b473b"
	HelperStrMaxExpArray[101] = "637b647c39cbb9d3d26c56e949"
	HelperStrMaxExpArray[102] = "bec763f8209b7a72b0afea0d31"
	HelperStrMaxExpArray[103] = "016ddc6556cdb84bdc8d12d22e6f"
	HelperStrMaxExpArray[104] = "02bd9ea4eed422ab6b7b072b029e"
	HelperStrMaxExpArray[105] = "054183095b2c8ececf30dd533d03"
	HelperStrMaxExpArray[106] = "0a14517cc6b9457111eed5b8adf1"
	HelperStrMaxExpArray[107] = "13545598e5c23276ccf0ede68034"
	HelperStrMaxExpArray[108] = "2511882c39c3adea96fec2102329"
	HelperStrMaxExpArray[109] = "471649d87199aa990756806903c5"
	HelperStrMaxExpArray[110] = "88534434053a9828af9f37367ee6"
	HelperStrMaxExpArray[111] = "01056f1b5bedf75c6bcb2ce8aed428"
	HelperStrMaxExpArray[112] = "01f55b9d9ddff141121e70ebe0104e"
	HelperStrMaxExpArray[113] = "03c1771ac9fb6b4c18e229803dae82"
	HelperStrMaxExpArray[114] = "0733d2d12ed20831ef0a4aead8c66d"
	HelperStrMaxExpArray[115] = "0dcff115b14eedde6fc3aa5353f2e4"
	HelperStrMaxExpArray[116] = "1a7cf47248624733f355c5c1f0d1f1"
	HelperStrMaxExpArray[117] = "32cbfd4a7adc790560b3335687b89b"
	HelperStrMaxExpArray[118] = "616a0ae1edcba5599528c20605b3f6"
	HelperStrMaxExpArray[119] = "bad03e7d883f69ad5b0a186184e06b"
	HelperStrMaxExpArray[120] = "016641a07658687a905357ac0ebe198b"
	HelperStrMaxExpArray[121] = "02af09481380a0a35cf1ba02f36c6a56"
	HelperStrMaxExpArray[122] = "05258b7ba7725d902050f6360afddf96"
	HelperStrMaxExpArray[123] = "09deaf736ac1f569deb1b5ae3f36c130"
	HelperStrMaxExpArray[124] = "12ed7b32a58f552afeb26faf21deca06"
	HelperStrMaxExpArray[125] = "244c49c648baa98192dce88b42f53caf"
	HelperStrMaxExpArray[126] = "459c079aac334623648e24d17c74b3dc"
	HelperStrMaxExpArray[127] = "857ddf0117efa215952912839f6473e6"
	for i := 0; i < 128; i++ {
		b, _ := hex.DecodeString(HelperStrMaxExpArray[i])
		HelperMaxExpArray[i] = (&big.Int{}).SetBytes(b)
	}
}

func initHelperMaxValArray() {
	var HelperStrMaxValArray [128]string
	HelperMaxValArray[0] = "2550a7d99147ce113d27f304d24a422c3d"
	HelperMaxValArray[1] = "1745f7d567fdd8c93da354496cf4dddf34"
	HelperMaxValArray[2] = "0b5301cf4bf20167721bcdbe218a66f1e0"
	HelperMaxValArray[3] = "05e2d2ca56fae9ef2e524ba4d0f75b8754"
	HelperMaxValArray[4] = "2f45acad795bce6dcd748391bb828dcea"
	HelperMaxValArray[5] = "17f631b6609d1047920e1a1f9613f870d"
	HelperMaxValArray[6] = "c29d4a7745ae89ef20a05db656441649"
	HelperMaxValArray[7] = "6242dea9277cf2d473468985313625bb"
	HelperMaxValArray[8] = "31aef9b37fbc57d1ca51c53eb472c345"
	HelperMaxValArray[9] = "1923b23c38638957faeb8b4fe57b5ead"
	HelperMaxValArray[10] = "cb919ec79bf364210433b9b9680eadd"
	HelperMaxValArray[11] = "67186c63186761709a96a91d44ff2bf"
	HelperMaxValArray[12] = "343e6242f854acd626b78022c4a8002"
	HelperMaxValArray[13] = "1a7efb7b1b687ccb2bb413b92d5e413"
	HelperMaxValArray[14] = "d72d0627fadb6aa6e0f3c994a5592a"
	HelperMaxValArray[15] = "6d4f32a7dcd0924c122312b7522049"
	HelperMaxValArray[16] = "37947990f145344d736c1e7e5cff2f"
	HelperMaxValArray[17] = "1c49d8ceb31e3ef3e98703e0e656cc"
	HelperMaxValArray[18] = "e69cb6255a180e2ead170f676fa3c"
	HelperMaxValArray[19] = "75a24620898b4a19aafdfa67d23e8"
	HelperMaxValArray[20] = "3c1419351dd33d49e1ce203728e25"
	HelperMaxValArray[21] = "1eb97e709f819575e656eefb8bd98"
	HelperMaxValArray[22] = "fbc4a1f867f03d4c057d522b6523"
	HelperMaxValArray[23] = "812507c14867d2237468ba955def"
	HelperMaxValArray[24] = "425b9d8ca5a58142d5172c3eb2b5"
	HelperMaxValArray[25] = "2228e76a368b75ea80882c9f6010"
	HelperMaxValArray[26] = "119ed9f43c52cdd38348ee8d7b23"
	HelperMaxValArray[27] = "91bfcff5e91c7f115393af54bad"
	HelperMaxValArray[28] = "4b8845f19f7b4a93653588ce846"
	HelperMaxValArray[29] = "273fa600431f30b0f21b619c797"
	HelperMaxValArray[30] = "1474840ba4069691110ff1bb823"
	HelperMaxValArray[31] = "ab212322b671a11d3647e3ecaf"
	HelperMaxValArray[32] = "59ce8876bf3a3b1b396ae19c95"
	HelperMaxValArray[33] = "2f523e50d3b0d68a3e39f2f06e"
	HelperMaxValArray[34] = "190c4f51698c5ee5c3b34928a0"
	HelperMaxValArray[35] = "d537c5d5647f2a79965d56f94"
	HelperMaxValArray[36] = "72169649d403b5b512b40d5c2"
	HelperMaxValArray[37] = "3d713a141a21a93a218c980c1"
	HelperMaxValArray[38] = "215544c77538e6de9275431a6"
	HelperMaxValArray[39] = "123c0edc8bf784d147024b7df"
	HelperMaxValArray[40] = "a11eada236d9ccb5d9a46757"
	HelperMaxValArray[41] = "59f185464ae514ade263ef14"
	HelperMaxValArray[42] = "32d507935c586248656e95cb"
	HelperMaxValArray[43] = "1d2270a4f18efd8eab5a27d7"
	HelperMaxValArray[44] = "10f7bfaf758e3c1010bead08"
	HelperMaxValArray[45] = "a101f6bc5df6cc4cf4cb56d"
	HelperMaxValArray[46] = "61773c45cb6403833991e6e"
	HelperMaxValArray[47] = "3c5f563f3abca8034b91c7d"
	HelperMaxValArray[48] = "265cd2a70d374397f75a844"
	HelperMaxValArray[49] = "1911bbf62c34780ee22ce8e"
	HelperMaxValArray[50] = "10e3053085e97a7710c2e6d"
	HelperMaxValArray[51] = "bbfc0e61443560740fa601"
	HelperMaxValArray[52] = "874f16aa407949aebced14"
	HelperMaxValArray[53] = "64df208d66f55c59261f5d"
	HelperMaxValArray[54] = "4dee90487e19a58fbf52e9"
	HelperMaxValArray[55] = "3e679f9e3b2f65e9d9b0db"
	HelperMaxValArray[56] = "33c719b34c57f9f7a922f6"
	HelperMaxValArray[57] = "2c7c090c36927c216fe17c"
	HelperMaxValArray[58] = "2789fc1ccdbd02af70650f"
	HelperMaxValArray[59] = "2451aae7a1741e150c6ae0"
	HelperMaxValArray[60] = "22700f74722225e8c308e6"
	HelperMaxValArray[61] = "21aae2600cf1170129eb92"
	HelperMaxValArray[62] = "21e552192ec12eccaa1d44"
	HelperMaxValArray[63] = "231a0b6c2a250a15897b8a"
	HelperMaxValArray[64] = "255901ff2640b9b00fef5e"
	HelperMaxValArray[65] = "28c842993fe2877ca68b09"
	HelperMaxValArray[66] = "2da7b7138200abf065bc12"
	HelperMaxValArray[67] = "34584e19c1677771772dbf"
	HelperMaxValArray[68] = "3d678fd12af3f51aa5828a"
	HelperMaxValArray[69] = "49a16c994ca36bb50c32c9"
	HelperMaxValArray[70] = "5a2b2d67887520aacedab6"
	HelperMaxValArray[71] = "70ac191abaee2a72987db6"
	HelperMaxValArray[72] = "8f8afbb1a74e96379df7b1"
	HelperMaxValArray[73] = "ba4bd6d86b43467101fd6c"
	HelperMaxValArray[74] = "f61f8e0679ef553e95c271"
	HelperMaxValArray[75] = "14ac1e3b06c9771ad8f351c"
	HelperMaxValArray[76] = "1c3d320c47b0e10030f080e"
	HelperMaxValArray[77] = "272f678a02b5bd5dcc145a7"
	HelperMaxValArray[78] = "3732bb25f4914992758a3aa"
	HelperMaxValArray[79] = "4ee25a85a30b4e758af15a0"
	HelperMaxValArray[80] = "724dbc7344a886ed20dbae2"
	HelperMaxValArray[81] = "a7d64de739a14a222daf692"
	HelperMaxValArray[82] = "f99876906cf6526b6b82ecc"
	HelperMaxValArray[83] = "177bbaca105a36b48757a319"
	HelperMaxValArray[84] = "23c442370233418f33964a65"
	HelperMaxValArray[85] = "3716c05776b217ecbb587d11"
	HelperMaxValArray[86] = "55c42bb597ed985a9d69778e"
	HelperMaxValArray[87] = "86e8f9efa6efeba9e16b0a90"
	HelperMaxValArray[88] = "d651f2e547d194ee8b6d9a69"
	HelperMaxValArray[89] = "157b681e454d31a35819b1989"
	HelperMaxValArray[90] = "22c414309a2b397b4f8e0eb28"
	HelperMaxValArray[91] = "38c1a2330fcf634a5db1378a0"
	HelperMaxValArray[92] = "5d6efaaf8133556840468bbbb"
	HelperMaxValArray[93] = "9b0c82dee2e1f20d0a157a7ae"
	HelperMaxValArray[94] = "10347bdd997b95a7905d850436"
	HelperMaxValArray[95] = "1b4c902e273a586783055cede8"
	HelperMaxValArray[96] = "2e50642e85a0b7c589bac2651b"
	HelperMaxValArray[97] = "4f1b7f75028232ad3258b8b742"
	HelperMaxValArray[98] = "880028111c381b5279db2271c3"
	HelperMaxValArray[99] = "eb454460fe475acef6b927865e"
	HelperMaxValArray[100] = "1996fab0c95ac4a2b5cfa8f555d"
	HelperMaxValArray[101] = "2cc9f3994685c8d3224acb9fea1"
	HelperMaxValArray[102] = "4ed2e079d693966878c7149351a"
	HelperMaxValArray[103] = "8b740d663b523dad8b67451d8fc"
	HelperMaxValArray[104] = "f7f73c5d826e196ff66a259204c"
	HelperMaxValArray[105] = "1bb0d7eb2857065dcad087986fa6"
	HelperMaxValArray[106] = "31b4dfa1eedd2bd17d3504820344"
	HelperMaxValArray[107] = "599fae8ac47c48cf034887f489bb"
	HelperMaxValArray[108] = "a249948898a0e444bffa21361f42"
	HelperMaxValArray[109] = "12711786051c98ca2acc4adf7ba6a"
	HelperMaxValArray[110] = "21a98821bf01e72cc3f724b65a121"
	HelperMaxValArray[111] = "3dad0dd7c71f7b443dddd56fede23"
	HelperMaxValArray[112] = "716933ca69ac1b439f976665fafdf"
	HelperMaxValArray[113] = "d143a4beebca9707458aad7b22dcd"
	HelperMaxValArray[114] = "18369cb4cd8522c1b28abc22a3e805"
	HelperMaxValArray[115] = "2cf816f46d1971ec18f0ffb6922e86"
	HelperMaxValArray[116] = "53c58e5a59ee4d9fd7f747f67a3aac"
	HelperMaxValArray[117] = "9c833e3c0364561037250933eab9a9"
	HelperMaxValArray[118] = "1253c9d983f03e6a0955355049411cb"
	HelperMaxValArray[119] = "226e05852615979ea99f6ef68dbab51"
	HelperMaxValArray[120] = "40d8c81134ee9e16db1e0108defbb9f"
	HelperMaxValArray[121] = "7a70173a27075f4b9482d36deadc951"
	HelperMaxValArray[122] = "e7b966d76665f99c3fb1791404f62c6"
	HelperMaxValArray[123] = "1b78e22c38ae6aa69d36b8ccfade23fd"
	HelperMaxValArray[124] = "3439aeef615a970c9678397b6ad71179"
	HelperMaxValArray[125] = "637d37d6cb204d7419ac094d7e89f0dd"
	HelperMaxValArray[126] = "bde80a98943810876a7852209de22be2"
	HelperMaxValArray[127] = "16b3160a3c604c6667ff40ff1882b0fcf"
	for i := 0; i < 128; i++ {
		b, _ := hex.DecodeString(HelperStrMaxValArray[i])
		HelperMaxValArray[i] = (&big.Int{}).SetBytes(b)
	}
}
