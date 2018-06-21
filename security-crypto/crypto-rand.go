package main

import (
	"math"
	"crypto/rand"
	"math/big"
	"log"
	"fmt"
	"encoding/binary"
)

/**
	the example produces 3 random values
	1. random bytes
    2. random integers
    3. usigned type integer
*/

func main(){
	limit := int64(math.MaxInt64)
	randInt, err := rand.Int(rand.Reader,big.NewInt(limit))
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Random integer using crypto: ", randInt)
	// next we generate a unsinged integer 32
	var num uint32
	err = binary.Read(rand.Reader,binary.BigEndian,&num)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Random unsigned integer: ", num)
	numBytes := 4
	randomBytes := make([]byte, numBytes)
	rand.Read(randomBytes)
	fmt.Println("Random bytes generated: ", randomBytes)
}
