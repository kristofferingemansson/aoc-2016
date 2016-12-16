package main

import (
	"fmt"
	"crypto/md5"
	"strconv"
)

func main() {
	keyPrefix := "reyedfim"
	code := [8]string{}
	digits := 0
	for counter := 0; digits < 8; counter++ {
		key := fmt.Sprintf("%v%v", keyPrefix, counter)
		hash := fmt.Sprintf("%x", md5.Sum([]byte(key)))
		if hash[0:5] == "00000" {
			pos := hash[5:6]
			if pos >= "0" && pos <= "7" {
				n, _ := strconv.Atoi(pos)
				if code[n] == "" {
					code[n] = hash[6:7]
					digits++
				}
			}
		}
	}
	fmt.Println(code)
}
