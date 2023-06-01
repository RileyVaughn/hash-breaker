package main

import (
	"breaker/sha256"
	official "crypto/sha256"
	"fmt"
	"math/rand"
	"time"
)

// The goal of this function is to test the self-implemented sha256 vs the built in golang version.
func main() {

	rand.Seed(time.Now().UnixNano())

	inputs := generateInput(10000)

	if Test(inputs) {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}

}

func generateString() string {

	numBytes := rand.Intn(5000)

	token := make([]byte, numBytes)
	rand.Read(token)

	return string(token)
}

func generateInput(testSize int) []string {

	var inputs []string

	for i := 0; i < testSize; i++ {
		inputs = append(inputs, generateString())
	}

	return inputs
}

func Test(inputs []string) bool {

	flag := true

	for _, v := range inputs {
		test1 := sha256.Hash(v)
		test2 := official.Sum256([]byte(v))
		if test1 != fmt.Sprintf("%x", test2) {
			fmt.Printf("%x\n", test2)
			fmt.Println(test1)
			flag = false
			break
		}
		//fmt.Printf("%v\n%x\n\n", test1, test2)

	}
	return flag
}
