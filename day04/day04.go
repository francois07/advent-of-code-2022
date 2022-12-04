package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type overlapChecker func(e1 [2]int, e2 [2]int) bool

func parseInput(input string) [][][2]int {
	var res [][][2]int
	inputSplit := strings.Split(input, "\n")

	for _, line := range inputSplit {
		lineSplit := strings.Split(line, ",")
		var group [][2]int

		for _, elf := range lineSplit {
			elfSplit := strings.Split(elf, "-")

			if len(elfSplit) != 2 {
				log.Fatal("Error when splitting elf sections")
			}

			fstSctn, errFstConv := strconv.Atoi(elfSplit[0])
			scdSctn, errScdConv := strconv.Atoi(elfSplit[1])

			if errFstConv != nil || errScdConv != nil {
				log.Fatal("Error when converting elf sections to integers")
			}

			sections := [2]int{fstSctn, scdSctn}

			group = append(group, sections)
		}

		res = append(res, group)
	}

	return res
}

func isFullyOverlapping(e1 [2]int, e2 [2]int) bool {
	return (e1[0] <= e2[0] && e1[1] >= e2[1]) || (e1[0] >= e2[0] && e1[1] <= e2[1])
}

func isOverlapping(e1 [2]int, e2 [2]int) bool {
	return e1[1] >= e2[0] && e1[0] <= e2[1]
}

func countOverlaps(input [][][2]int, countFunc overlapChecker) int {
	overlapCount := 0

	for _, pair := range input {
		if countFunc(pair[0], pair[1]) {
			overlapCount += 1
		}
	}

	return overlapCount
}

func part1(input string) int {
	parsedInput := parseInput(input)

	return countOverlaps(parsedInput, isFullyOverlapping)
}

func part2(input string) int {
	parsedInput := parseInput(input)

	return countOverlaps(parsedInput, isOverlapping)
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
