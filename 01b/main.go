package main

import (
	"io/ioutil"
	"os"
	"strings"
	"strconv"
	"fmt"
	"log"
)

type position struct {
	x, y int
}

var (
	currentDirection = 0
	currentPosition = position{0, 0}

	directionMod = [4]position{
		position{0, 1},
		position{1, 0},
		position{0, -1},
		position{-1, 0},
	}

	positionUsage = map[int]map[int]bool{}
)

func main() {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	input, err := ioutil.ReadFile(dir + "/01b/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	moves := strings.Split(string(input), " ")

	setVisited(currentPosition)

	Loop:
		for _, v := range moves {
			turn(v[0:1])
			distance, _ := strconv.Atoi(strings.TrimRight(v[1:], ",\n"))

			mod := directionMod[currentDirection]

			for i := 0; i < distance; i++ {
				currentPosition.x += mod.x
				currentPosition.y += mod.y
				if isVisited(currentPosition) {
					break Loop
				}
				setVisited(currentPosition)
			}
		}

	fmt.Printf("Distance: %v", abs(currentPosition.x) + abs(currentPosition.y))
}

func turn(turn string) (int) {
	switch turn {
	case "R":
		currentDirection += 1
	case "L":
		currentDirection -= 1
	}

	if currentDirection < 0 {
		currentDirection = 3
	} else if currentDirection > 3 {
		currentDirection = 0
	}

	return currentDirection
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isVisited(pos position) bool {
	return positionUsage[pos.y][pos.x]
}

func setVisited(pos position) {
	_, exist := positionUsage[pos.y]
	if !exist {
		positionUsage[pos.y] = map[int]bool{}
	}
	positionUsage[pos.y][pos.x] = true
}