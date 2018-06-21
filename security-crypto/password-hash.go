package main

import (
	"fmt"
	"os"
	"encoding/base64"
	"crypto/hmac"
	"crypto/sha256"
	"io"
	"encoding/hex"
	"crypto/rand"
)

func printUsageP(){
	fmt.Printf("Usage: %s <pasword>\n", os.Args[0])
	fmt.Printf("Example: %s pasword1!", os.Args[0])
}

func checkArgsP() string{
	if len(os.Args) < 2{
		printUsageP()
		os.Exit(1)
	}
	return os.Args[1]
}

/* keeping a secret key for HMAC
   HMAC stands for - hash-based message authentication code - this adds additional secret key into hashing algorithm.
*/
var secretKey = "neictr98y85klfgneghre"

func generateSalt() string {
	randomByte := make([]byte,32)
	_, err := rand.Read(randomByte)
	if err != nil{
		return ""
	}
	return base64.URLEncoding.EncodeToString(randomByte)
}

func hashPassword(plainText string, salt string) string{
	hash := hmac.New(sha256.New, []byte(secretKey))
	io.WriteString(hash, plainText+salt)
    hashedValue := hash.Sum(nil)
    return hex.EncodeToString(hashedValue)
}

func main(){
	pass := checkArgsP()
	salt := generateSalt()

	hmacPass := hashPassword(pass, salt)
	fmt.Printf("Hashed Password %s salt %s is : %s\n",pass, salt, hmacPass)
}