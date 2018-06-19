package main

import (
	"math/rand"
	"fmt"
	"math/big"
	crypto "crypto/rand"
)

func main(){
	sec1 := rand.New(rand.NewSource(10))
	sec2 := rand.New(rand.NewSource(10))

	for i := 0; i<5; i++{
		rnd1 := sec1.Int()
		rnd2 := sec2.Int()
		if rnd1 != rnd2 {
			fmt.Printf("Rand generated non-equal sequences")
			break
		}else{
			fmt.Printf("Math/Rand1: %d , Math/Rand2: %d \n", rnd1, rnd2)
		}
	}

	for i := 0; i<5; i++{
		safeRand1 := NewCryptoRand()
		safeRand2 := NewCryptoRand()

		if safeRand1 == safeRand2 {
			fmt.Printf("Crypto Rand generated equal sequences- rare")
			break
		}else{
			fmt.Printf("Crypto/Rand1: %d , Crypto/Rand2: %d \n", safeRand1, safeRand2)
		}
	}
}

func NewCryptoRand() int64{
	safeNum, err := crypto.Int(crypto.Reader, big.NewInt(100234))
	if err != nil{
		panic(err)
	}
	return safeNum.Int64()
}
