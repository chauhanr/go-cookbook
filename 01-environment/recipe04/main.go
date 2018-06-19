package main

import (
	"os"
	"log"
)

func main(){
	args := os.Args

	if len(args) != 2{
		log.Printf("Usage : lookup env-var")
		return
	}

	key := args[1]
	os.Setenv(key, "postgres://as:as@example.org/pg?sslmode=verify-full")
	val := GetEnvDefault(key, "postgres://as:as@localhost/pg?sslmode=verify-full")
	log.Printf("The value of Env %s is %s \n", key, val)

	os.Unsetenv(key)
	val = GetEnvDefault(key, "postgres://as:as@127.0.0.1/pg?sslmode=verify-full")
	log.Printf("The value of Env %s after unset is %s \n", key, val)
}

func GetEnvDefault(key, defaultValue string) string{
	val, exists := os.LookupEnv(key)
	if !exists{
		return defaultValue
	}
	return val
}