package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"strings"
	"strconv"
	"regexp"
)

type Triangle struct {
	a, b, c int
}

func main() {
	triangles := GetTriangles()

	numValid := 0
	numInvalid := 0
	for _, triangle := range triangles {
		if triangle.IsValid() {
			numValid++;
		} else {
			numInvalid++;
		}
	}
	fmt.Printf("Valid: %v\n", numValid)
	fmt.Printf("Invalid: %v\n", numInvalid)
}

func (t *Triangle) IsValid() bool {
	return t.a + t.b > t.c && t.b + t.c > t.a && t.a + t.c > t.b
}

func GetTriangles() []Triangle {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	input, err := ioutil.ReadFile(dir + "/03a/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	rows := strings.Split(strings.Trim(string(input), "\n"), "\n")

	r, _ := regexp.Compile("(\\d+)\\s+(\\d+)\\s+(\\d+)")
	ret := []Triangle{}
	for _, row := range rows {
		m := r.FindStringSubmatch(row);
		if m != nil {
			ret = append(ret, Triangle{CastToInt(m[1]), CastToInt(m[2]), CastToInt(m[3])})
		}
	}
	return ret
}

func CastToInt(s string) int {
	i, _ := strconv.Atoi(strings.Trim(s, " "));
	return i
}