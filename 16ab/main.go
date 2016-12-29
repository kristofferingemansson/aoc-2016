package main

import (
	"fmt"
)


func main() {
	input := "10111100110001111"
	size := 272
	size = 35651584

	data := GenerateData(input, size)

	hash := data
	for true {
		hash = CalcHash(hash)
		if !IsEven(len(hash)) {
			fmt.Println("Hash: ", hash)
			break
		}
	}
}

func GenerateData(input string, length int) string {
	a := input
	for len(a) < length {
		b := StringReverse(a)
		b = FlipZeroOne(b)
		a = a + "0" + b
	}
	return a[:length]
}

func StringReverse (s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, r := range s {
		n--
		runes[n] = r
	}
	return string(runes[:])
}

func FlipZeroOne (s string) string {
	o := make([]rune, len(s))
	for i, r := range s {
		switch r {
		case '1':
			o[i] = '0'
		case '0':
			o[i] = '1'
		default:
			o[i] = r
		}
	}
	return string(o[:])
}

func CalcHash(d string) string {
	l := len(d)
	o := make([]rune, l/2)
	for i, l := 0, len(d); i < l; i += 2 {
		if d[i:i+1] == d[i+1:i+2] {
			o[i/2] = '1'
		} else {
			o[i/2] = '0'
		}
	}
	return string(o[:])
}

func IsEven(i int) bool {
	return i & 0x1 == 0
}