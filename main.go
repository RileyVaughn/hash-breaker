package main

import "fmt"

func main() {
	b := []byte{0}

	for i := 0; i < 10; i++ {
		fmt.Println(b)
		b = incremBit(b)
	}

}

func incremBit(b []byte) []byte {

	for i := len(b) - 1; i >= 0; i-- {
		if b[i] == byte(0) {
			b[i] = byte(1)
			return b
		}
	}
	b = append([]byte{1}, b...)
	return b
}
