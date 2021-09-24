package main

import "fmt"

func main() {
	b := []byte{0}

	for i := 0; i < 1000; i++ {

		fmt.Printf("%08b", b)
		println()
		incremBit(&b)
	}

}

func incremBit(b *[]byte) {

	for i := len(*b) - 1; i >= 0; i-- {
		if (*b)[i] != byte(255) {
			(*b)[i]++
			for j := i + 1; j < len(*b); j++ {
				(*b)[j] = byte(0)
			}
			return
		}
	}

	// Append new byte to front of slice. Only happens slice is if full of 255's
	(*b)[0] = byte(1)
	for i := 1; i < len(*b); i++ {
		(*b)[i] = byte(0)
	}
	*b = append(*b, byte(0))
	return
}
