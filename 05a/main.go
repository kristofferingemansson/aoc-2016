package main

import (
	"fmt"
	"crypto/md5"
)

func main() {
	keyPrefix := "reyedfim"
	code := ""
	digits := 0
	for counter := 0; digits < 8; counter++ {
		key := fmt.Sprintf("%v%v", keyPrefix, counter)
		hash := fmt.Sprintf("%x", md5.Sum([]byte(key)))
		if hash[0:5] == "00000" {
			code += hash[5:6]
			digits++
		}
	}
	fmt.Println(code)
}
