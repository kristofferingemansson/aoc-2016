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

	input, err := ioutil.ReadFile(dir + "/03b/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	rows := strings.Split(strings.Trim(string(input), "\n"), "\n")

	tmp := [3][3]int{}

	r, _ := regexp.Compile("(\\d+)\\s+(\\d+)\\s+(\\d+)")
	ret := []Triangle{}
	n := 0
	for _, row := range rows {
		m := r.FindStringSubmatch(row);
		if m != nil {
			tmp[0][n] = CastToInt(m[1])
			tmp[1][n] = CastToInt(m[2])
			tmp[2][n] = CastToInt(m[3])
			n++
			if n == 3 {
				n = 0
				ret = append(
					ret,
					Triangle{tmp[0][0], tmp[0][1], tmp[0][2]},
					Triangle{tmp[1][0], tmp[1][1], tmp[1][2]},
					Triangle{tmp[2][0], tmp[2][1], tmp[2][2]},
				)
			}
		}
	}
	return ret
}

func CastToInt(s string) int {
	i, _ := strconv.Atoi(strings.Trim(s, " "));
	return i
}