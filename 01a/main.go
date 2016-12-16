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
)

func main() {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	input, err := ioutil.ReadFile(dir + "/01a/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	moves := strings.Split(string(input), " ")

	for _, v := range moves {
		turn(v[0:1])
		distance, _ := strconv.Atoi(strings.TrimRight(v[1:], ",\n"))

		mod := directionMod[currentDirection]
		currentPosition.x += mod.x * distance
		currentPosition.y += mod.y * distance
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