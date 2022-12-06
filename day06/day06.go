package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func findMarker(input string, markerSize int) int {
	for i := range input {
		buffer := input[i : i+markerSize]
		isMarker := true

		for _, letter := range buffer {
			lCount := strings.Count(buffer, string(letter))

			if lCount > 1 {
				isMarker = false
			}
		}

		if isMarker {
			return i + markerSize
		}
	}

	return -1
}

func part1(input string) int {
	return findMarker(input, 4)
}

func part2(input string) int {
	return findMarker(input, 14)
}

func main() {
	fileData, errRead := os.ReadFile("input.txt")
	dataString := string(fileData)

	if errRead != nil {
		log.Fatal(errRead)
	}

	fmt.Printf("Part 1 : %d\n", part1(dataString))
	fmt.Printf("Part 2 : %d\n", part2(dataString))
}
