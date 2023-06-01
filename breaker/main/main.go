package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var BITS uint8 = 8

func main() {

	var a string = "a"
	var b string = "b"

	sum := add8_2zn(a, b)

	//Test(BITS)

	err := os.WriteFile("sum.txt", []byte(sum), 0644)
	if err != nil {
		log.Fatal(err)
	}

}

func Test(size uint8) {

	num := uint8(math.Pow(2, float64(size)))
	pass := true

	for i := uint8(0); i < num; i++ {
		for j := uint8(0); j < num; j++ {

			a := add_bit_8_2(i, j) % num
			b := (i + j) % num

			if a != b {
				pass = false
				fmt.Printf("(%v) (%v) %v %v\n", i, j, a, b)

			}

		}
	}
	fmt.Println(pass)
}

func add8(a string, b string) string {

	for i := 0; i < 3; i++ {

		tempa := a
		tempb := b

		a = tempa + "^" + tempb
		b = "((" + tempa + "<<1)&(" + tempb + "<<1))"

	}
	return a
}

func add8_2zn(a string, b string) string {

	a = a + "+" + b
	b = strings.Replace(a, "+", "_1", 1) + "_1"

	for i := 1; i < int(BITS); i++ {

		tempa := a
		tempb := b

		// Add tempa and tempb
		a = tempa + "+" + tempb

		//don't waste computation on calculatin unused b
		if i+1 < int(BITS) {

			b = ""

			// Multiply each monomial in tempa with the monomial in tempb

			tempa2 := strings.Split(tempa, "+")
			tempb2 := strings.Split(tempb, "+")

			for _, mona := range tempa2 {
				for _, monb := range tempb2 {
					b = b + "+" + mona + monb
				}
			}
			//remove extra + sign
			b = b[1:]

			//Insert _1
			b = strings.ReplaceAll(b, "b", "b_1")
			b = strings.ReplaceAll(b, "a", "a_1")

			// Replace _1_X with _(X+1)
			for j := 1; j < int(BITS); j++ {
				val := "1_" + strconv.Itoa(j)
				for strings.Contains(b, val) {
					b = strings.ReplaceAll(b, val, strconv.Itoa(j+1))
				}
			}

			//remove repeating multiplications\

			bsplit := strings.Split(b, "+")
			for k, mon := range bsplit {
				coefs := []string{}
				for strings.Contains(mon, "_") {
					index := strings.LastIndex(mon, "_") - 1
					coefs = append(coefs, mon[index:])
					mon = mon[0:index]
				}

				coefs = removeDuplicateValues(coefs)
				bsplit[k] = strings.Join(coefs, "")
			}

			b = strings.Join(bsplit, "+")
		}

	}
	return a
}

func add_bit_8(a uint8, b uint8) uint8 {

	return a ^ b ^ ((a) & b << 1) ^ ((a ^ b) & ((a) & b << 1) << 1) ^ ((a ^ b ^ ((a) & b << 1)) & ((a ^ b) & ((a) & b << 1) << 1) << 1) ^ ((a ^ b ^ ((a) & b << 1) ^ ((a ^ b) & ((a) & b << 1) << 1)) & ((a ^ b ^ ((a) & b << 1)) & ((a ^ b) & ((a) & b << 1) << 1) << 1) << 1) ^ ((a ^ b ^ ((a) & b << 1) ^ ((a ^ b) & ((a) & b << 1) << 1) ^ ((a ^ b ^ ((a) & b << 1)) & ((a ^ b) & ((a) & b << 1) << 1) << 1)) & ((a ^ b ^ ((a) & b << 1) ^ ((a ^ b) & ((a) & b << 1) << 1)) & ((a ^ b ^ ((a) & b << 1)) & ((a ^ b) & ((a) & b << 1) << 1) << 1) << 1) << 1) ^ ((a ^ b ^ ((a) & b << 1) ^ ((a ^ b) & ((a) & b << 1) << 1) ^ ((a ^ b ^ ((a) & b << 1)) & ((a ^ b) & ((a) & b << 1) << 1) << 1) ^ ((a ^ b ^ ((a) & b << 1) ^ ((a ^ b) & ((a) & b << 1) << 1)) & ((a ^ b ^ ((a) & b << 1)) & ((a ^ b) & ((a) & b << 1) << 1) << 1) << 1)) & ((a ^ b ^ ((a) & b << 1) ^ ((a ^ b) & ((a) & b << 1) << 1) ^ ((a ^ b ^ ((a) & b << 1)) & ((a ^ b) & ((a) & b << 1) << 1) << 1)) & ((a ^ b ^ ((a) & b << 1) ^ ((a ^ b) & ((a) & b << 1) << 1)) & ((a ^ b ^ ((a) & b << 1)) & ((a ^ b) & ((a) & b << 1) << 1) << 1) << 1) << 1) << 1) ^ ((a ^ b ^ ((a) & b << 1) ^ ((a ^ b) & ((a) & b << 1) << 1) ^ ((a ^ b ^ ((a) & b << 1)) & ((a ^ b) & ((a) & b << 1) << 1) << 1) ^ ((a ^ b ^ ((a) & b << 1) ^ ((a ^ b) & ((a) & b << 1) << 1)) & ((a ^ b ^ ((a) & b << 1)) & ((a ^ b) & ((a) & b << 1) << 1) << 1) << 1) ^ ((a ^ b ^ ((a) & b << 1) ^ ((a ^ b) & ((a) & b << 1) << 1) ^ ((a ^ b ^ ((a) & b << 1)) & ((a ^ b) & ((a) & b << 1) << 1) << 1)) & ((a ^ b ^ ((a) & b << 1) ^ ((a ^ b) & ((a) & b << 1) << 1)) & ((a ^ b ^ ((a) & b << 1)) & ((a ^ b) & ((a) & b << 1) << 1) << 1) << 1) << 1)) & ((a ^ b ^ ((a) & b << 1) ^ ((a ^ b) & ((a) & b << 1) << 1) ^ ((a ^ b ^ ((a) & b << 1)) & ((a ^ b) & ((a) & b << 1) << 1) << 1) ^ ((a ^ b ^ ((a) & b << 1) ^ ((a ^ b) & ((a) & b << 1) << 1)) & ((a ^ b ^ ((a) & b << 1)) & ((a ^ b) & ((a) & b << 1) << 1) << 1) << 1)) & ((a ^ b ^ ((a) & b << 1) ^ ((a ^ b) & ((a) & b << 1) << 1) ^ ((a ^ b ^ ((a) & b << 1)) & ((a ^ b) & ((a) & b << 1) << 1) << 1)) & ((a ^ b ^ ((a) & b << 1) ^ ((a ^ b) & ((a) & b << 1) << 1)) & ((a ^ b ^ ((a) & b << 1)) & ((a ^ b) & ((a) & b << 1) << 1) << 1) << 1) << 1) << 1) << 1)
}

func add_bit_8_2(a uint8, b uint8) uint8 {
	return (((a ^ b) ^ ((a << 1) & (b << 1))) ^ (((a ^ b) << 1) & (((a << 1) & (b << 1)) << 1)))

	//a ^ b ^ ((a << 1) & (b << 1)) ^ ((a<<1) & ((a<<2) & (b<<2)) ^ (b<<1) & ((a<<2) & (b<<2)))

}

func removeDuplicateValues(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
