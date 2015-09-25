package minhash

import (
	"math"
	"math/big"
	"math/rand"
	"fmt"
)

var bitMask = uint32(0x1)
var hashFuncNum = 10

// finding the minimum key
func minKey(l map[string]uint32) (string, uint32) {
	var result string
	m := uint32(math.MaxUint32)
	//fmt.Println(m)
	//fmt.Println(l)
	for k := range l {
		//fmt.Println("l[k] ",l[k])
		if m > l[k] {
			//fmt.Println(m," > ", l[k])
			m = l[k]
			result = k
		}
	}

	return result, m
}

func minHash(data []string, seed uint32) uint32 {
	// fmt.Println(data)
	vector := make(map[string]uint32)
	// fmt.Println(vector)
	for k := range data {
		vector[data[k]] = Murmurhash3_32(data[k], seed)
	}
	_, value := minKey(vector)
	// fmt.Println(value)
	return value
}

func signature(data []string) *big.Int {
	rand.Seed(1)
	sigBig := big.NewInt(0)
	for i := 0; i < hashFuncNum; i++ {
		// fmt.Println(uint(minHash(data, rand.Uint32())&bitMask))
		sigBig.SetBit(sigBig, i, uint(minHash(data, rand.Uint32())&bitMask))
	}
	return sigBig
}

func popCount(bits uint32) uint32 {
	bits = (bits & 0x55555555) + (bits >> 1 & 0x55555555)
	bits = (bits & 0x33333333) + (bits >> 2 & 0x33333333)
	bits = (bits & 0x0f0f0f0f) + (bits >> 4 & 0x0f0f0f0f)
	bits = (bits & 0x00ff00ff) + (bits >> 8 & 0x00ff00ff)
	return (bits & 0x0000ffff) + (bits >> 16 & 0x0000ffff)
}

func popCountBig(bits *big.Int) int {
	result := 0
	// fmt.Println(bits)
	// fmt.Println(bits.Bytes())
	for _, v := range bits.Bytes() {
		fmt.Println(v)
		fmt.Println("==========")
		fmt.Println(uint32(v))
		fmt.Println(popCount(uint32(v)))
		result += int(popCount(uint32(v)))
	}
	fmt.Println("000000000000")
	return result
}

func minhashFromSignature(sig1, sig2 *big.Int) float32 {
	commonBig := big.NewInt(0)
	// fmt.Println(sig1, "====", sig2)
	commonBig.Xor(sig1, sig2)
	// fmt.Println(commonBig)
	return 2.0 * (float32(hashFuncNum-popCountBig(commonBig))/float32(hashFuncNum) - 0.5)
}

func Minhash(v1, v2 []string) float32 {
	return minhashFromSignature(signature(v1), signature(v2))
}
