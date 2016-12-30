package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"strings"
	"regexp"
	"strconv"
)

type Pos struct {
	x int
	y int
}

type Disc struct {
	pos Pos
	size int
	used int
	avail int
}

type Cluster map[Pos]Disc

func main() {
	data := GetInputData()
	cluster := ParseData(data)

	pairs := 0
	for _, disc := range cluster {
		pairs += disc.GetViablePairs(cluster)
	}
	fmt.Println(pairs)
}

func GetInputData() []string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	input, err := ioutil.ReadFile(dir + "/22a/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	return strings.Split(strings.Trim(string(input), "\n"), "\n")
}

func ParseData(data []string) Cluster {
	headers := 2
	cluster := make(Cluster, len(data) - headers)

	regexpDf, _ := regexp.Compile("/dev/grid/node-x(\\d+)-y(\\d+)\\s+(\\d+)T\\s+(\\d+)T\\s+(\\d+)T")
	for i, row := range data {
		if i < headers {
			continue
		}
		m := regexpDf.FindStringSubmatch(row)
		if m != nil {
			pos := Pos{
				x: StringToInt(m[1]),
				y: StringToInt(m[2]),
			}
			disc := Disc{
				pos: pos,
				size: StringToInt(m[3]),
				used: StringToInt(m[4]),
				avail: StringToInt(m[5]),
			}
			cluster[pos] = disc
		}
	}

	return cluster
}

func StringToInt(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}

func (d *Disc) GetViablePairs(cluster Cluster) int {
	if d.used == 0 {
		return 0
	}

	ret := 0
	for _, target := range cluster {
		if target == *d {
			continue
		}

		if d.used <= target.avail {
			ret++
		}
	}

	return ret
}