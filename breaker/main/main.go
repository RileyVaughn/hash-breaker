package main

import (
	"breaker/sha256"
	"fmt"
	"log"
	"strconv"
)

func main() {
	msg := "Hello World!"
	hash := sha256.Hash(msg)
	fmt.Println(hash)

	ls := sha256.HashLastStep(msg)
	reverseStep(ls)

}

func stringToHexSlice(hash string) [8]uint32 {

	var H [8]uint32
	for i := 0; i < len(hash); i = i + 8 {
		hx, err := strconv.ParseInt(hash[i:i+8], 16, 64)
		hexHash := uint32(hx)
		if err != nil {
			log.Fatalln("Error converting hash string into hex slice", err)
		}
		H[i/8] = hexHash
	}

	return H
}

func reverseStep(step [8]uint32) {

	var prev [8]uint32

	prev[0] = step[1]
	prev[1] = step[2]
	prev[2] = step[3]
	prev[3] = step[4] - step[0] + sha256.Σ0(prev[0]) + sha256.Maj(prev[0], prev[1], prev[2])
	prev[4] = step[5]
	prev[5] = step[6]
	prev[6] = step[7]
	prev[7] = 0
	fmt.Printf("%x", prev)
}
