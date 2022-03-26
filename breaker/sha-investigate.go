package main

import (
	"fmt"
	"hashFunctions/sha256"
)

func main() {
	msg := "Hello World!"
	hash := sha256.Hash(msg)
	fmt.Println(hash)
}
