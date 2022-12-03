package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	ROCK     string = "A"
	PAPER           = "B"
	SCISSORS        = "C"
)

func parseInput(input string) [][2]string {
	var res [][2]string
	inputSplit := strings.Split(input, "\n")

	for _, line := range inputSplit {
		var lineSplit [2]string
		copy(lineSplit[:], strings.Split(line, " "))

		if len(lineSplit) != 2 {
			log.Fatal("Input is not a couple")
		}

		res = append(res, lineSplit)
	}

	return res
}

func part1(input string, winScore int, drawScore int, winMap map[string]string, scoreMap map[string]int) int {
	parsedInput := parseInput(input)
	translateMap := map[string]string{
		"X": ROCK,
		"Y": PAPER,
		"Z": SCISSORS,
	}
	finalScore := 0

	for _, round := range parsedInput {
		opp, player := round[0], translateMap[round[1]]

		if player == winMap[opp] {
			finalScore += winScore
		} else if player == opp {
			finalScore += drawScore
		}

		finalScore += scoreMap[player]
	}

	return finalScore
}

func part2(input string, winScore int, drawScore int, winMap map[string]string, scoreMap map[string]int) int {
	parsedInput := parseInput(input)
	finalScore := 0
	loseMap := make(map[string]string)

	for k, v := range winMap {
		loseMap[v] = k
	}

	for _, round := range parsedInput {
		opp, action := round[0], round[1]
		var player string

		switch action {
		case "X":
			player = loseMap[opp]
		case "Y":
			player = opp
			finalScore += drawScore
		case "Z":
			player = winMap[opp]
			finalScore += winScore
		}

		finalScore += scoreMap[player]
	}

	return finalScore
}

func main() {
	winMap := map[string]string{
		ROCK:     PAPER,
		PAPER:    SCISSORS,
		SCISSORS: ROCK,
	}
	scoreMap := map[string]int{
		ROCK:     1,
		PAPER:    2,
		SCISSORS: 3,
	}
	winScore, drawScore := 6, 3

	fileData, errRead := os.ReadFile("input.txt")
	dataString := string(fileData)

	if errRead != nil {
		log.Fatal(errRead)
	}

	fmt.Printf("Part 1 : %d\n", part1(dataString, winScore, drawScore, winMap, scoreMap))
	fmt.Printf("Part 2 : %d\n", part2(dataString, winScore, drawScore, winMap, scoreMap))
}
