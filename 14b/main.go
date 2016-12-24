package main

import (
	"fmt"
	"crypto/md5"
	"strings"
)

func main() {
	input := "zpqevtbw"
	seq := 0
	fmt.Println("Input key:", input)

	offsetLimit := 1000

	hashes := make([]string, offsetLimit * 100)
	for i := seq; i <= offsetLimit; i++ {
		hashes[i] = GenerateKey(input, i)
	}

	fmt.Println("Generated initial set")

	foundKeys := 0
	for foundKeys < 64 {
		k1 := hashes[seq]
		hashes[seq+offsetLimit] = GenerateKey(input, seq+offsetLimit)
		keychar := IsValidKey(k1, 3, "")
		if keychar != "" {
			for i := 1; i <= offsetLimit; i++ {
				kc2 := IsValidKey(hashes[seq + i], 5, keychar)
				if kc2 != "" {
					foundKeys++
					fmt.Println(foundKeys, seq, k1, hashes[seq+i])
					break
				}
			}
		}
		seq++
	}
}

func GenerateKey(input string, seq int) string {
	key := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%v%v", input, seq))))
	for i := 0; i < 2016; i++ {
		key = fmt.Sprintf("%x", md5.Sum([]byte(key)))
	}
	return key
}

func IsValidKey(key string, reps int, check string) string {
	for i, l := 0, len(key) - (reps - 1); i < l; i++ {
		c := check
		if c == "" {
			c = key[i:i+1]
		}
		if strings.Count(key[i:i+reps], c) == reps {
			return c
		}
	}
	return ""
}